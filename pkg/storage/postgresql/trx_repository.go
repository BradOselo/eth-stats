package postgresql

import (
	"context"
	"time"

	"github.com/cardenasrjl/eth-stats/pkg/utils"

	v1 "github.com/cardenasrjl/eth-stats/pkg/api/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//FeeRepo ..
type FeeRepo interface {
	GetFeesFromDateTimeRange(ctx context.Context, dateTimeStart string, dateTimeEnd string, resolution int) (result []*v1.Fee, totalFees float32, err error)
}

//FeeDB
type FeeDB struct {
	TotalFee uint64    `db:"total_fees"`
	Datetime time.Time `db:"datetime"`
}

//GetFeesFromDateTimeRange ...
func (db *DB) GetFeesFromDateTimeRange(ctx context.Context, dateTimeStart string, dateTimeEnd string, resolution int) (result []*v1.Fee, totalFees float32, err error) {
	// get SQL connection from pool
	c, err := db.Connect(ctx)
	if err != nil {
		return
	}
	defer c.Close()

	query := `
		select sum(tx.gas_used * tx.gas_price) as total_fees,  to_timestamp(floor((extract('epoch' from block_time) / $1 )) * $1)  AS datetime
		from public.transactions tx
		where tx.status = 'true'
		  and tx.from <> '0x0000000000000000000000000000000000000000'
		  and tx."to" <> '0x0000000000000000000000000000000000000000'
		  and not exists(
				select 'x' from public.contracts ct where (ct.address = tx."from")
			)
		  and not exists(
				select 'x' from public.contracts ct where (tx."to" = ct.address)
			)
		  and block_time between $2 AND $3
		group by  datetime;
	`
	rows, err := c.QueryContext(ctx, query, resolution, dateTimeStart, dateTimeEnd)
	if err != nil {
		return nil, 0, status.Error(codes.Unknown, "failed to select from fees-> "+err.Error())
	}
	defer rows.Close()

	result = []*v1.Fee{}
	for rows.Next() {
		feeDB := FeeDB{}
		if err := rows.Scan(&feeDB.TotalFee, &feeDB.Datetime); err != nil {
			return nil, 0, status.Error(codes.Unknown, "failed to retrieve field values row-> "+err.Error())
		}
		fee := v1.Fee{
			T: feeDB.Datetime.Unix(),
			V: utils.WeiToEther(feeDB.TotalFee),
		}
		result = append(result, &fee)
		totalFees += fee.V
	}

	return
}

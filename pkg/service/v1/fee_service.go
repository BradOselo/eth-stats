package v1

import (
	"context"
	"database/sql"

	"github.com/cardenasrjl/eth-stats/pkg/utils"

	"github.com/cardenasrjl/eth-stats/pkg/storage/postgresql"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	v1 "github.com/cardenasrjl/eth-stats/pkg/api/v1"
	_ "github.com/cardenasrjl/eth-stats/pkg/storage/postgresql"
)

// Resolution in seconds
const DefaultResolution = 60

// feesServiceServer...
type feesServiceServer struct {
	FeeRepo postgresql.FeeRepo
}

// NewFeeServiceServer ...
func NewFeeServiceServer(db *sql.DB) v1.FeeServiceServer {
	st := postgresql.NewDB(db)
	return &feesServiceServer{FeeRepo: st}
}

//validateRequest ...
func validateRequest(m *v1.GetRequest) error {
	if m.Start == 0 || m.End == 0 || m.End < m.Start {
		return status.Error(codes.Unknown, "Invalid timestamp provided ")
	}

	if m.Resolution <= 0 {
		m.Resolution = DefaultResolution
	}

	if int64(m.Resolution) > (m.End - m.Start) {
		return status.Error(codes.Unknown, "The resolution is to big, provide a wider timestamp range")
	}

	//TODO: other validations here to ensure that will not compromise the request
	return nil
}

// GetFees ...
//example: 1599476400 - 1599479999
func (i feesServiceServer) GetFees(ctx context.Context, req *v1.GetRequest) (*v1.GetResponse, error) {
	err := validateRequest(req)
	if err != nil {
		return nil, err
	}

	//todo: would be good to cache this
	fees, totalFees, err := i.FeeRepo.GetFeesFromDateTimeRange(ctx, utils.TimestampToDateTime(req.Start), utils.TimestampToDateTime(req.End), int(req.Resolution))
	if err != nil {
		return nil, err
	}

	//return the response
	return &v1.GetResponse{
		Fees:  fees,
		Total: totalFees,
	}, nil
}

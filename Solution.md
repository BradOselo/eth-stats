## To execute the solution:  

```
$Make up
```
or

```
$docker-compose up
```

#### URL for test: 

```
http://0.0.0.0:8080/v1/eth-stats/fees?start=1599476400&end=1599479999&Resolution=600
```

#### Test Result 

```json
{
    "fees": [
        {
            "t": "1599476400",
            "v": 28.476894
        }
    ],
    "total": 28.476894
}
```

## Solution

I propose to have an endpoint `/v1/eth-stats/fees` with 3 parameters: `Start`,
 `End` and `Resolution`. The `Start` and `End` are unix timestamps and are be the boundaries
 in time to query the database. The resolution is the amount of seconds to be used 
 to group the transactions and get the total fees. 
 
 With this solution we can have hourly grouping (as requested in the task) with a resolution 
 of 3600. If we need more granularity lets say per minute, then we can set a resolution 
 of 60 or if we want per second then we set a resolution of 1. If we don't provide the resolution, the default value is 60.   

## GRPC 

 I decided to create it using GRPC since from my experience allows you to have the service ready
if GRPC is needed in the future bringing all it's benefits. .

## Main structure:

Entry point for initiating the service: 
````
cmd/server/server.go 
````
Proto Files: 

```
api/proto/v1/
```

Service fee implementation: 
````
pkg/service/v1/fee_service.go
````

Database package for managing repositories:

````
pkg/storage/postgresql/trx_repository.go
````

  
## Things to improve with more time: 

1. More and proper testing. 
2. Would consider adding an index to the block_time column in transactions table.. 
2. Depending on the circumstances I would like to cache the results, because this is information that doesn't change once it is stored. Actually would be interesting to try a non transactional database to store the source data. 
3. There is no authentication so it should be implemented. 
4. Proper configuration management. 
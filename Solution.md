## Solution for eth-stats

I propose to have an api endpoint `/v1/eth-stats/fees`. with 3 parameters: `Start`,
 `End` and `Resolution`. The `Start` and `End` are unix timestamps and will be the boundaries
 in time to query the database (limits have to be set to avoid issues). The resolution will tell us the amount of seconds to be used 
 to group the transactions and get the total fees. 
 
 With this solution we can have hourly grouping (as requested in the task), and for that we would need a resolution 
 of 3600. IF we would like to have more granularity lets say per minute, then we can have a resolution 
 of 60 or if we want per second then we have a resolution of 1. The default resolution if not provided is 60.   

## GRPC 

 I decided to create it using GRPC since from my experience allows you to have the service ready
if GRPC is needed in the future bringing all it's benefits. .

The /cmd/server/server.go is the entry point from were the two services 
 /pkg/grpc/server.go and /pkg/rest/server.go are initiated.
 
  
## Things to improve with more time: 

1. More and proper testing. 
2. Would consider adding an index to the block_time column in transactions table.. 
2. Depending on the circumstances I would cache the results, since this is information that doesn't change once it is stored. Actually would be interesting to try a non transactional database to have this data. 
3. There is no authentication so it should be implemented. 

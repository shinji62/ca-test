# ca-test
## Description

Just wait `WAITING_TIME` for each request. This could be used to simulate socket exhaustion.


## Push to CF

Edit `manifest.yml` if needed then `cf push`


`WAITING_TIME` should be in `time.Duration` format https://golang.org/pkg/time/#Duration ex: 1m is 1 minutes.


## Load Testing

Just use the software you use and send many request to the defined endpoint.

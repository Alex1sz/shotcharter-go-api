## Build
![alt text](https://travis-ci.org/Alex1sz/shotcharter-go-api.svg?branch=master "Build status")


## Dev environment setup

1. Install glide `curl https://glide.sh/get | sh`
1. Install vendor dependencies `glide install`
1. Create test db `createdb shotcharter_go_test`
1. Create tables/ load db schema `psql shotcharter_go_test -f db/schema_setup.sql`
1. `go build main.go`
1. Run on `PORT=8080 go run main.go`

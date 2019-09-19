# Property Service
---
## Installation
Install Go (preferably version 1.13+)

Make sure your GOPATH and GOBIN paths are properly configured

### Checkout code

`git clone https://github.com/Sutheres/property-service.git && cd property`

### Install dependencies

`make setup`

### Start required containers

`docker-compose up -d`

### Run server/worker
#### Command Line

To run migrations:
`go run main.go migrate` 

To start HTTP server:
`go run main.go server`

To start background worker:
`go run main.go worker` (coming soong)

#### VSCode Launch.json
"server" == `go run main.go server`

"worker" == `go run main.go worker`

## Development
### Swagger JSON editor
[stoplight](https://stoplight.io/)

### To Regenerate Swagger files

`make gen`

## Testing

### To re-generate mocks against our interfaces:
`go generate ./...`

### To run tests
`go test -v ./...`

run individual tests through `run test | debug test` inline links in VSCode

If you want to see full coverage profile, run:
`make coverage` (coming soon)


## ENV VARS:

### Please refer to sample.env for environment variables that should be set.

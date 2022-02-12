# Time sender and receiver

This project shows basic knowledge of:
- Golang
  - API server with graceful shutdown
  - Client with configurable periodical calls to API server (env: REQUESTS_FREQUENCY)
- gRPC server-client communication
- Docker multistage build
- CI/CD with Github Actions
- docker-compose (with networking)

## Setup and Run
To pull sender and receiver images and run them both copy-paste and run following command in your terminal:

`
docker-compose -f deployments/docker-compose.yml pull && docker-compose -f deployments/docker-compose.yml up
`

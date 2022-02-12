# go-e

## Deliverables

- a public repository in any git management system of your choice.
- docker-compose manifest with 2 services:
    - Services communicate with Golang Protobuf messages over gRPC
  interfaces.
    - First service generates data (e.g current time).
    - Second service receives and represents data in the log output.
    - According log output in the terminal, for eg:
    
      `Service 172.0.first-service.ip: sending 12356789
      Service 172.0.second-service.ip: received 12356789`
      
## Setup and Run

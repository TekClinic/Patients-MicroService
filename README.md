# Patients-Microservice

This repository contains a gRPC service for managing patient information within TekClinic. 
The service is implemented in Go and uses Protobuf for defining message types and service methods.

Please note that the provided code assumes the existence of a `TekClinic/MicroService-Lib `
library for authentication and environment variable handling, 
and setting up the environment variables found in `TekClinic/MicroService-Lib` is a prerequisite.

## Table of Contents

- [Installation](#installation)
- [gRPC Functions](docs/grpc.md#grpc-functions)
    - [GetPatient](docs/grpc.md#getpatient)
    - [CreatePatient](docs/grpc.md#createpatient)
    - [GetPatientsIDs](docs/grpc.md#getpatientsids)

## Installation

1. Clone the repository:

```bash
git clone https://github.com/TekClinic/Patients-MicroService.git
```

2. Set up the required environment variables for database connection:

```
DB_ADDR=<database_address>
DB_USER=<database_user>
DB_PASSWORD=<database_password>
DB_DATABASE=<database_name>
```

3. Run the server:

```bash
go run server.go
```
# xeneelectronic

## Setup

### DI and build

This module uses Wire for dependency injection. Add dependencies to wire.go and then run

```
make wire
```

Build the app as an executable

```
make build
```

### Running migrations

Run all migrations

```
make migrate
# or
./dist/xeneelectronic migrate
# or
go run --mod=vendor ./ migrate
# or
make migration-run
```

## Run the server

Starting the server

```
make run
# or
go run --mod=vendor ./ serve
```

- Run migrations

Once test environment is ready, run the tests:

```
# Run tests under a spcific package
# go test --mod=vendor ./path/to/package/to/be/tested/

# Run tests under models package
go test --mod=vendor ./internal/models

# Run all tests
go test --mod=vendor ./
```

## Development

### Project structure

```
- /xeneelectronic
    - wire.go                   // setup DI here
    - wire_gen.go               // generate DI file
    - server.go
    - main.go
    - /config
        - config.go
        - config.yaml           // define config variables here
    - /db
        - db_client.go
        - *_client.go           // specific DB cient implementation
    - /handlers
        - base_handler.go
        _ *_handler.go          // specific route handler implementation
    - /internals
        - /models
            - base_model.go
            - models_test.go    // contains TestMain() where models package level test dependencies are defined
            - repository.go
            - *_model.go        // model and repository implementation for a specific entity
            - *_model_test.go   // class specific tests
        - /services
            - base_service.go
            - *_service.go      // specific service implementation
    - /swagger
        - swagger.yaml          // base swagger file
        - /components
            - *.yaml            // referenced swagger file per Definition
```

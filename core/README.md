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

Run up to specific migration

```
make migrate-one ID=<migration_id>
# or
./dist/xeneelectronic migrate --id=<migration_id>
# or
go run --mod=vendor ./ migrate --id=<migration_id>
# or
make migration-target id=<migration_id>
```

Rollback down to specific migration

```
./dist/xeneelectronic rollback --id=<migration_id>
# or
go run --mod=vendor ./ rollback --id=<migration_id>
# or
make migration-rollback id=<migration_id>
```

## Run the server

Starting the server

```
make run
# or
go run --mod=vendor ./ serve
```

## Testing

Before testing make sure test environment is ready. Some preparations to consider:

- Create database, and load necessary database extensions (ie. UUID extension)

```
...
psql> CREATE DATABASE "tcc_core";
psql> \c "tcc_core"
psql> CREATE EXTENSION "uuid-ossp";
...
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

### Migrations

Create a new migration file, `migration_<unique_migration_id>.go`, under `./db/migrations` and add the new migration struct to the constructor in `migration.go`.

It is best to copy the struct into the migrations instead of referencing to the struct in the models so the state of the struct will remain consistent during migration.

### Models and repositories

Models and repositories are created under `./internal/models`. Model and repositories per entity can be defined in a single file.

### Writing tests

Create a separate `<package>_test.go` file to define dependecy setup of the package's tests. Dependency setup of tests should be written in the package's `TestMain()` function. Refer to `./internals/models/models_test.go` for sample implementation.

For writing tests for a specific file, create an equivalent `<class_file>_test.go`. Organize the test functions using Go's subtests structure. Using this structure the setup and tear-down code specific to the class to be tested can be written once, and the tests function are called subsequently. Refer to `./internals/models/project_model_test.go` for sample implementation.

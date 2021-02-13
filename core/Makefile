NAME = xeneelectronic-core
FILES = main.go server.go wire_gen.go
ID = ""

wire:
	rm -f wire_gen.go
	./bin/wire

build:
	go mod vendor
	go build -mod=vendor -o dist/$(NAME) $(FILES)

run:
	./dist/${NAME} serve

migrate: build
	./dist/xeneelectronic-core migrate

migrate-one: build
	./dist/xeneelectronic-core migrate --id=$(ID)

rollback: build
	./dist/xeneelectronic-core rollback

rollback-one: build
	./dist/xeneelectronic-core rollback --id=$(ID)

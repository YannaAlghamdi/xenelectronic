NAME = xenelectronic-core
FILES = main.go server.go wire_gen.go
ID = ""

wire:
	rm -f wire_gen.go
	./bin/wire

build:
	go mod vendor
	env GOOS=linux GOARCH=amd64 GOARM=7 go build -mod=vendor -o dist/$(NAME) $(FILES)

run:
	./dist/${NAME} serve

migrate: build
	./dist/xenelectronic-core migrate

migrate-one: build
	./dist/xenelectronic-core migrate --id=$(ID)

rollback: build
	./dist/xenelectronic-core rollback

rollback-one: build
	./dist/xenelectronic-core rollback --id=$(ID)

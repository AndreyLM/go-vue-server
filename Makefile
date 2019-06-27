GCC=go
GCMD=run
GPATH=./cmd/server/main.go

run:
	${GCC} ${GCMD} ${GPATH}

build:
	make build_db

build_db:
	go run ./pkg/db/generator/main.go -json=./pkg/db/generator/config.json
	mv ./db_structs.go ./pkg/models
NAME?=ngcp-cdr-db

all:
	go build -ldflags "-s -w"  -o $(NAME) cmd/ngcp-cdr-db/*.go


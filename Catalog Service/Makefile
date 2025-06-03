build:
	go build -o bin/icom main.go

run:
	go run main.go

buildandrun: build
	./bin/icom

omnidb:
	mkdir -p /run/postgresql
	/usr/bin/postgres -D ../~/pgdata/ -p 5433 -i &

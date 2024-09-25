run-api:
	go build -o backend ./cmd/http/ && ./backend

migration-up:
	go build -o backend ./cmd/http/ && ./backend --migration-up

migration-down:
	go build -o backend ./cmd/http/ && ./backend --migration-down

dry-run:
	go build -o backend ./cmd/http/ && ./backend -d
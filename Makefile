.PHONY: install test build

install:
	cd backend && go mod download
	cd frontend && npm ci

test:
	cd backend && go test ./...
	cd frontend && npm test -- --run

build:
	cd backend && go build -o pascalcalc-api ./cmd/api
	cd frontend && npm run buil
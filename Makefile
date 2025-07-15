.PHONY: prod
prod:
	make build
	./bin/app

.PHONY: dev
dev: 
	@make prepare
	make -j3 templ-watch tailwind-watch air

.PHONY: prepare
prepare:
	mkdir -p ./bin && chmod -R 777 ./public


.PHONY: templ
templ:
	TEMPL_EXPERIMENT=rawgo templ generate -keep-orphaned-files


.PHONY: tailwind
tailwind:
	bunx --yes tailwindcss -i ./public/assets/styles/tailwind.css -o ./public/assets/styles/style.css --minify

.PHONY: tailwind-watch
tailwind-watch:
	bunx --yes tailwindcss -i ./public/assets/styles/tailwind.css -o ./public/assets/styles/style.css --watch


.PHONY: build
build:
	make prepare
	go mod tidy
	make templ tailwind vet staticcheck test
	GOFLAGS="-mod=readonly -modcacherw" go build -o .-v -x -race -trimpath ./bin/app ./cmd/app/main.go

.PHONY: air

# Database Migration-Seeding-Parsing Commands
.PHONY: migrate
migrate:
	go run ./cmd/databases/postgres/migrate/main.go

.PHONY: drop
drop:
	go run ./cmd/databases/postgres/migrate/drop/main.go

.PHONY: seed
seed:
	go run ./cmd/databases/postgres/seed/main.go


.PHONY: vet
vet:
	go vet ./...

.PHONY: staticcheck
staticcheck: 
	staticcheck ./...
	
.PHONY: test
test:
	go test -race -v -timeout 30s ./...


.PHONY: env
env:
	cp -r .example.env .env

.PHONY: update
update:
	go get -u ./...
	go mod tidy
	go mod download
	go mod verify
	go mod tidy

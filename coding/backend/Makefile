.PHONY: generate
generate:
	go generate ./ent

.PHONY: run-watch
run-watch: generate
	go run github.com/gofiber/cli/fiber@latest dev

.PHONY: run
run: generate
	go run main.go

.PHONY: reset-db
reset-db:
	rm -f app.db:ent

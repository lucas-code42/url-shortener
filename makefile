start-redis:
	docker compose up -d


run:
	go mod tidy && go build main.go && ./main
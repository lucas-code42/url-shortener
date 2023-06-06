start-redis:
	docker compose up -d


run:
	go build main.go && ./main
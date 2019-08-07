.PHONY: hello hello2

hello:
	docker-compose up -d hello
	docker-compose exec hello go get -v .
	docker-compose exec hello go run main.go

hello2:
	docker-compose up -d hello2
	docker-compose exec hello2 go get -v .
	docker-compose exec hello2 go run main.go

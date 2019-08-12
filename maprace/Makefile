.PHONY: fmt race1 race2 rwmutex syncmap atomicvalue

fmt:
	docker-compose up -d develop
	docker-compose exec develop gofmt -w -s .

race1:
	docker-compose up -d develop
	docker-compose exec develop go run main.go

race2:
	docker-compose up -d develop
	docker-compose exec develop go run main_using_range.go

rwmutex:
	docker-compose up -d develop
	docker-compose exec develop go run main_rw_mutex.go

syncmap:
	docker-compose up -d develop
	docker-compose exec develop go run main_sync_map.go

atomicvalue:
	docker-compose up -d develop
	docker-compose exec develop go run main_atomic_value.go

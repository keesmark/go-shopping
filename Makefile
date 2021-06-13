up:
	docker-compose up -d
stop:
	docker-compose stop
api:
	docker exec -it go_api sh
front:
	docker exec -it go_frontend sh
run:
	docker exec -i go_api go run main.go
down:
	docker-compose down
up:
	docker-compose up -d
stop:
	docker-compose stop
api:
	docker exec -it api sh
run:
	docker exec -i api go run main.go
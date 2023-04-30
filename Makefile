.PHONY: rundb run stop

# Определение переменных
IMAGE_NAME=my-docker-image
CONTAINER_NAME=my-docker-container

# Команды для сборки Docker образа
rundb:
	docker compose up

# Команды для запуска Docker контейнера и запуска Go-приложения
filldb:
	go run main.go

# Команды для остановки и удаления Docker контейнера
stop:
	docker-compose down

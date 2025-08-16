package main

import (
	"log"

	"github.com/Kilril312/tasks-service/internal/database"
	"github.com/Kilril312/tasks-service/internal/task"
	"github.com/Kilril312/tasks-service/internal/transport"
)

func main() {

	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Ошибка инициализации БД: %v", err)
	}

	repo := task.NewRepository(db)
	svc := task.NewService(repo)

	// 3. Клиент к Users-сервису
	userClient, conn, err := transport.NewUserClient("localhost:50052")
	if err != nil {
		log.Fatalf("failed to connect to users: %v", err)
	}
	defer conn.Close()

	// 4. Запуск gRPC Tasks-сервиса
	if err := transport.RunGRPC(svc, userClient); err != nil {
		log.Fatalf("Tasks gRPC server error: %v", err)
	}
}

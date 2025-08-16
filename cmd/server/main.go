package main

import (
	"log"
	"os"

	"github.com/Kilril312/tasks-service/internal/database"
	"github.com/Kilril312/tasks-service/internal/task"
	"github.com/Kilril312/tasks-service/internal/transport/grpc"
)

func main() {

	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Ошибка инициализации БД: %v", err)
	}

	repo := task.NewRepository(db)
	svc := task.NewService(repo)

	userServiceAddr := os.Getenv("USER_SERVICE_ADDR")
	if userServiceAddr == "" {
		userServiceAddr = "localhost:50051"
	}

	// 3. Клиент к Users-сервису
	userClient, conn, err := grpc.NewUserClient("localhost:50051")
	if err != nil {
		log.Fatalf("failed to connect to users: %v", err)
	}
	defer conn.Close()

	// 4. Запуск gRPC Tasks-сервиса
	if err := grpc.RunGRPC(svc, userClient); err != nil {
		log.Fatalf("Tasks gRPC server error: %v", err)
	}
}

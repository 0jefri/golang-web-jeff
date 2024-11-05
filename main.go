package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/golang-web/database"
	"github.com/golang-web/handler"
	"github.com/golang-web/middleware"
	"github.com/golang-web/repository"
	"github.com/golang-web/service"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Gagal connect database!!!!", err)
	} else {
		log.Println("Success connect to database !!!")
	}
	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(&userRepo)
	userHandler := handler.NewUserHandler(userService)

	taskRepo := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(&taskRepo)
	taskHandler := handler.NewTaskHandler(taskService)

	// ROUTER

	//cara manggil endpoint nya POST http://localhost:8080/register
	http.HandleFunc("/register", userHandler.RegisterHandler)

	//cara manggil endpoint nya POST http://localhost:8080/login
	http.HandleFunc("/login", middleware.TokenMiddleware(userHandler.LoginHandler))
	// http.HandleFunc("/login", userHandler.LoginHandler)

	// cara manggil endpoint nya GET : http://localhost:8080/user?id=1
	http.HandleFunc("/user", middleware.TokenMiddleware(userHandler.GetUserByID))

	//cara manggil endpoint get all users GET : http://localhost:8080/users
	http.HandleFunc("/users", middleware.TokenMiddleware(userHandler.GetAllUsersHandler))

	//cara manggil endpoint create task POST : http://localhost:8080/task
	http.HandleFunc("/task", middleware.TokenMiddleware(taskHandler.RegisterHandlerTask))

	//cara manggil endpoint list task GET : http://localhost:8080/tasks
	http.HandleFunc("/tasks", middleware.TokenMiddleware(taskHandler.GetAllTaskHandler))

	fmt.Println("Server started on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}

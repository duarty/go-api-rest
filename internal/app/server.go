package app

import (
	"goapirest/internal/controller"
	"goapirest/internal/infra/database"
	"goapirest/internal/usecase"
	"log"
	"net/http"
)

func Server() {
	db, err := database.DBConnection()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	gymRepo := database.NewGymRepository(db)
	gymUseCase := usecase.NewGymUseCase(gymRepo)
	gymController := controller.NewGymController(gymUseCase)

	http.HandleFunc("/create-gym", gymController.CreateGymController)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

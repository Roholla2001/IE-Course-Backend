package boot

import (
	"fmt"
	"log"

	"github.com/Roholla2001/ie-course-backend/internal/infra/datastore"
	"github.com/Roholla2001/ie-course-backend/internal/infra/router"
	"github.com/joho/godotenv"
)

func BootServer() error {
	//loading environment variables from .env file
	fmt.Println("Loading Env...")
	err := godotenv.Load()
	if err != nil {
		log.Fatal(fmt.Errorf("Error Loading .env file, %w", err))
	}

	//getting a new connection pool to database
	db, err := datastore.NewDBConn()
	if err != nil {
		log.Fatal(fmt.Errorf("Error connecting to DB: %w", err))
	}

	//create a new app controller
	ac := new(router.AppController)

	//create a new server controller
	ac.ServerController, err = ac.NewServerController(db)

	//create a new user controller
	ac.UserController, err = ac.NewUserController(db)
	if err != nil {
		return err
	}

	//create new url controllers
	ac.URLController, err = ac.NewURLController(db)
	if err != nil {
		return err
	}


	router := router.InitRouter(ac)

	if err := router.Run("localhost:8080"); err != nil {
		log.Fatalf("server could not start")
	}

	return nil
}

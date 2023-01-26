package boot

import (
	"fmt"
	"log"

	"github.com/Roholla2001/ie-course-backend/internal/infra/datastore"
	"github.com/joho/godotenv"
)

func BootServer() error {
	fmt.Println("Loading Env...")
	err := godotenv.Load()
	if err != nil {
		log.Fatal(fmt.Errorf("Error Loading .env file, %w", err))
	}

	db, err := datastore.NewDBConn()
	if err != nil {
		log.Fatal(err)
	}

}

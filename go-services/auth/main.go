package auth

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// чтобы работало добавить докерфайл для гошки
	dbConfig := Config{
		Host:     "postgres",
		Port:     "5432",
		User:     "postgres",
		Password: "ubuntu2!",
		DBName:   "lessonsdb",
	}

	db, err := NewPostgres(dbConfig)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	userRepo := NewUserRepository(db)

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET env is empty")
	}
	authHandler := NewAuthHandler(userRepo, jwtSecret)

	http.HandleFunc("/register", authHandler.registerHandler)
	//http.HandleFunc("/login", loginHandler)
	//http.HandleFunc("/verify", verifyHandler)

	log.Println("Server starting")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	// Database setup
	dsn := os.Getenv("DB_DSN")
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Create a router
	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/api/home", HomeHandler).Methods("GET")
	r.HandleFunc("/api/about", AboutHandler).Methods("GET")
	r.HandleFunc("/api/contact", ContactHandler).Methods("POST")

	// Wrap router with CORS middleware
	corsRouter := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)(r)

	// Start the server
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", corsRouter))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"message": "Welcome to Home!"})
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"message": "Welcome to About!"})
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		Email string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Send email to Telegram channel via bot (implement Telegram API request)
	sendToTelegram(requestData.Email)

	json.NewEncoder(w).Encode(map[string]string{"message": "Email sent!"})
}

func sendToTelegram(email string) {
	botToken := os.Getenv("BOT_KEY")
	chatID := os.Getenv("MANAGER_CHAT_ID")
	text := fmt.Sprintf("New request from: %s", email)

	telegramURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=%s", botToken, chatID, text)
	http.Get(telegramURL)
}

package server

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"go-gemma/db"
	"go-gemma/model"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
)

type Request struct {
	Command string `json:"command"`
	Token   string `json:"token"`
	Text    string `json:"text"`
}

var ctx = context.Background()

func StartServer() {
	database := db.InitDB()
	db.SeedToken(database)

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	http.HandleFunc("/askGemma", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
			return
		}

		var req Request

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var token db.Token

		var commandAndText = req.Command + req.Text

		result := database.Where("token = ?", req.Token).First(&token)

		if result.Error != nil {
			http.Error(w, "Token not authorized", http.StatusUnauthorized)
			return
		}

		h := sha256.New()
		h.Write([]byte(commandAndText))
		hashedText := hex.EncodeToString(h.Sum(nil))

		val, err := rdb.Get(ctx, hashedText).Result()

		if err == redis.Nil {
			output, err := model.AskGemma(commandAndText)

			log.Printf("Command: %s, Text: %s, Result: %s", req.Command, req.Text, string(output))

			if err != nil {
				log.Print(err.Error())

				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			err = rdb.Set(ctx, hashedText, output, 0).Err()

			if err != nil {
				log.Printf("Failed to save result in Redis: %v", err)
			}

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"result": string(output),
			})
		} else if err != nil {
			log.Printf("Redis error: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"result": val,
			})
		}
	})

	log.Println("Server is running on http://localhost:8081")

	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}

}

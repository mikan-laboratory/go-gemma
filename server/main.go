package server

import (
	"encoding/json"
	"go-get-gemma/db"
	"go-get-gemma/model"
	"log"
	"net/http"
)

type Request struct {
	Token string `json:"token"`
	Text  string `json:"text"`
}

func StartServer() {
	database := db.InitDB()
	db.SeedToken(database)

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

		result := database.Where("token = ?", req.Token).First(&token)

		if result.Error != nil {
			http.Error(w, "Token not authorized", http.StatusUnauthorized)
			return
		}

		output, err := model.AskGemma(req.Text)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Text processed",
			"result":  string(output),
		})
	})

	log.Println("Server is running on http://localhost:3000")

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}

}

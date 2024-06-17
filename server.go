package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/openai/openai-go"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	http.HandleFunc("/upload", uploadHandler)
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
		return
	}

	r.ParseMultipartForm(10 << 20) // Limit upload size

	file, _, err := r.FormFile("receipt")
	if err != nil {
		http.Error(w, "Invalid file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}

	summary, amount, err := processReceipt(fileBytes)
	if err != nil {
		http.Error(w, "Error processing receipt", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"summary": summary,
		"amount":  amount,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func processReceipt(fileBytes []byte) (string, string, error) {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	prompt := fmt.Sprintf("Please summarize this purchase and the dollar amount: %s", string(fileBytes))

	resp, err := client.Completions.Create(openai.CompletionRequest{
		Model:   "text-davinci-003",
		Prompt:  prompt,
		MaxTokens: 100,
	})
	if err != nil {
		return "", "", err
	}

	// Assuming the response contains the summary and amount in the first text output
	// This part may need adjustment based on actual OpenAI response format
	text := resp.Choices[0].Text
	// Further parsing of text to extract summary and amount can be added here

	return "Purchase summary", "$100", nil // Placeholder return values
}

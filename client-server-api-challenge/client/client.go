package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Exchange struct {
	Bid float64 `json:"bid"`
}

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}
	host := os.Getenv("HOST")
	fileRepository := os.Getenv("FILE_REPOSITORY")

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/cotacao", host), nil)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var exc Exchange
	err = json.Unmarshal(body, &exc)
	if err != nil {
		log.Fatal(err)
	}

	err = createFile(&exc, fileRepository)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Exchange: %v\n", exc.Bid)
}

func createFile(exchange *Exchange, fileRepository string) error {
	filePath := fmt.Sprintf("%s/cotacao.txt", fileRepository)
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	content := fmt.Sprintf("Dólar: %v", exchange.Bid)
	size, err := f.Write([]byte(content))
	if err != nil {
		return err
	}
	log.Printf("file %s created! Size: %d\n", filePath, size)

	return nil
}

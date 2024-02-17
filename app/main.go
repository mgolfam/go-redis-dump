package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

func main() {
	// Replace these values with your Redis server information
	redisAddr := "localhost:6379"
	redisPassword := ""
	redisDB := 0

	// Connect to Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       redisDB,
	})

	// Check if Redis connection is successful
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("Failed to connect to Redis:", err)
		return
	}
	fmt.Println("Connected to Redis:", pong)

	// Replace 'yourKeyPattern' with the actual key pattern you want to retrieve
	keyPattern := "*airline*"
	keys, err := client.Keys(context.Background(), keyPattern).Result()
	if err != nil {
		fmt.Println("Error retrieving keys:", err)
		return
	}

	// Open CSV file for writing
	file, err := os.Create("redis_data.csv")
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer file.Close()

	// Create CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header to CSV file
	header := []string{"Key", "airline_code", "img", "title", "title_fa", "type"}
	if err := writer.Write(header); err != nil {
		fmt.Println("Error writing header to CSV:", err)
		return
	}

	// Iterate through keys and retrieve values
	for _, key := range keys {
		value, err := client.Get(context.Background(), key).Result()
		if err != nil {
			fmt.Printf("Error retrieving value for key %s: %v\n", key, err)
			continue
		}

		values, err := JSONStringToValues(value)
		if err != nil {
			fmt.Println(err)
		}

		// Write key and value to CSV file
		row := append([]string{key}, values...)
		if err := writer.Write(row); err != nil {
			fmt.Printf("Error writing row to CSV for key %s: %v\n", key, err)
			continue
		}
	}

	fmt.Println("CSV conversion complete. Check 'redis_data.csv' file.")
}

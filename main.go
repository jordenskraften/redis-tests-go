package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	fmt.Println("heya")

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ctx := context.Background()

	session := map[string]string{"name": "John", "surname": "Smith", "company": "Redis", "age": "29"}
	for k, v := range session {
		err := client.HSet(ctx, "user-session:123", k, v).Err()
		if err != nil {
			panic(err)
		}
	}

	userSession := client.HGetAll(ctx, "user-session:123").Val()
	for k, v := range userSession {
		fmt.Printf("user-session:123 key %s val %s\n", k, v)
	}
	//==================================
	type TestStruct struct {
		Name    string `json:"name"`
		Company string `json:"company"`
		Value   int    `json:"value"`
	}

	data := TestStruct{Name: "ya", Company: "dva", Value: 1414}
	dataJSON, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	err = client.HSet(ctx, "user-session:444", "data", dataJSON).Err()
	if err != nil {
		panic(err)
	}

	dataRaw := client.HGet(ctx, "user-session:444", "data").Val()
	var data2 TestStruct
	err = json.Unmarshal([]byte(dataRaw), &data2)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Name: %s, Company: %s, Value: %d\n", data2.Name, data2.Company, data2.Value)
}

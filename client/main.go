package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	config := Config{}

	err := config.ReadConfig("config.json")
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return
	}

	fmt.Printf(
		"Making %d requests to %s with %d TPS and %d users\n",
		config.RequestCount,
		config.BaseUrl,
		config.Tps,
		config.UserCount,
	)

	userIDs := GetUserIDs(config.UserCount)

	ticker := time.NewTicker(time.Second / time.Duration(config.Tps))
	defer ticker.Stop()

	var wg sync.WaitGroup

	for i := 0; i < config.RequestCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range ticker.C {
				userID := PickUserID(userIDs)
				makeRequest(
					config.BaseUrl,
					userID,
				)
			}
		}()
	}

	wg.Wait()
}

func makeRequest(
	baseUrl string,
	userID int,
) {
	url := fmt.Sprintf(baseUrl, userID)
	_, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(url)
}

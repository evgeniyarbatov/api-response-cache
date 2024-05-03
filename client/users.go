package main

import (
	"math/rand"
)

func GetUserIDs(userCount int) []int {
	var userIDs []int
	for i := 0; i < userCount; i++ {
		userID := rand.Intn(9000000000) + 1000000000
		userIDs = append(userIDs, userID)
	}
	return userIDs
}

func PickUserID(userIDs []int) int {
	idx := rand.Intn(len(userIDs))
	return userIDs[idx]
}

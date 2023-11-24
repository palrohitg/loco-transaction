package main

import (
	"fmt"
)

func main() {

	newData := map[string]interface{}{
		"7": map[string]interface{}{
			"user_action_performed": true,
		},
		"14": map[string]interface{}{
			"user_action_performed": true,
		},
		"21": map[string]interface{}{
			"user_action_performed": true,
		},
	}

	streakDay := "7"
	quizId := "23443242342"
	newMap := map[string]interface{}{}

	for key, value := range newData {
		if key == streakDay {
			newKey := streakDay + "_" + quizId
			newMap[newKey] = value
		} else {
			newKey := key + "_" + "old"
			newMap[newKey] = value
		}

	}
	fmt.Println(newMap)
	return
}

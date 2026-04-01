package storage

import (
	"encoding/json"
	"fmt"
	"health-tracker/models"
	"os"
)

func LoadOrRegister(filename string) (models.User, bool) {
	data, err := os.ReadFile(filename)
	if err == nil {
		var u models.User
		err := json.Unmarshal(data, &u)
		if err != nil {
			fmt.Println("Error while reading data", err)
		}
		return u, true
	}
	return models.User{}, false
}

func Save(u models.User, filename string) {
	jsonData, err := json.MarshalIndent(u, "", " ")
	if err != nil {
		fmt.Println("Error", err)
	}
	errSave := os.WriteFile(filename, jsonData, 0644)

	if errSave != nil {
		fmt.Println("Error saving data to file:", errSave)
	}

}

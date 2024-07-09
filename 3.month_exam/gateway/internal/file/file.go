package file

import (
	"encoding/json"
	"fmt"
	"gateway/auth/hash"
	"gateway/internal/models"
	"log"
	"os"
)

func WriteNewUserToFile(filepath string, cLogin models.ClientLogin) error {
	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Println("failed to open file while reading :", err)
	}
	defer f.Close()

	var clients []models.ClientLogin
	err = json.NewDecoder(f).Decode(&clients)
	if err != nil {
		clients = []models.ClientLogin{}
	}

	for _, client := range clients {
		if client.Id == cLogin.Id {
			return fmt.Errorf("user with ID '%s' already exists", cLogin.Id)
		}
	}
	clients = append(clients, cLogin)
	f, err = os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Printf("failed to open file while writing: %v", err)
	}
	defer f.Close()
	err = json.NewEncoder(f).Encode(clients)
	if err != nil {
		log.Printf("failed to encode users to JSON: %v", err)
	}
	return nil
}

func CheckClientFromFile(filepath string, req models.ClientLogin) error {
	f, err := os.OpenFile(filepath, os.O_RDWR ,  0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer f.Close()

	var clients []models.ClientLogin
	err = json.NewDecoder(f).Decode(&clients)
	if err != nil {
		return fmt.Errorf("failed to decode file: %v", err)
	}

	for _, client := range clients {
		if client.Id == req.Id {
			return nil
		}
	}
	return fmt.Errorf("no client found with this id and email ")
}

func DeleteClientFromFile(filepath string, id string) error {
	fmt.Println("started")
	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer f.Close()

	var clients []models.ClientLogin
	err = json.NewDecoder(f).Decode(&clients)
	if err != nil {
		return fmt.Errorf("failed to decode file: %v", err)
	}
	fmt.Println(clients)
	for i, client := range clients {
		if client.Id == id {
			clients = append(clients[:i], clients[i+1:]...)
			f, err = os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
			if err != nil {
				log.Printf("failed to open file while writing: %v", err)
			}
			defer f.Close()
			err = json.NewEncoder(f).Encode(clients)
			if err != nil {
				log.Printf("failed to encode users to JSON: %v", err)
			}
			return nil
		}
	}
	return fmt.Errorf("client does not exists ")
}
func WriteNewAdminToFile(filepath string, aLogin models.AdminLogin) error {
	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Println("failed to open file while reading :", err)
	}
	defer f.Close()

	var admins []models.AdminLogin
	err = json.NewDecoder(f).Decode(&admins)
	if err != nil {
		admins = []models.AdminLogin{}
	}

	for _, admin := range admins {
		if admin.Id == aLogin.Id {
			return fmt.Errorf("admin with ID '%s' already exists", aLogin.Id)
		}
	}
	admins = append(admins, aLogin)
	f, err = os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Printf("failed to open file while writing: %v", err)
	}
	defer f.Close()
	err = json.NewEncoder(f).Encode(admins)
	if err != nil {
		log.Printf("failed to encode users to JSON: %v", err)
	}
	return nil
}

func CheckAdminfromFile(filepath string, id, password string) error {
	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer f.Close()

	var admins []models.AdminLogin
	err = json.NewDecoder(f).Decode(&admins)
	if err != nil {
		return fmt.Errorf("failed to decode file: %v", err)
	}

	for _, admin := range admins {
		if admin.Id == id && hash.ValidateHashPassword(admin.HashPassword, password) {
			return nil
		}
	}
	return fmt.Errorf("no admin found with  Id: %v  and password: %v", id, password)
}

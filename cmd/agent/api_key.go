package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func checkAPIKey() (string, error) {
  key, err := readAPIKey()
  if err != nil {
    return "", err
  }

  if len(key) != 0 {
    return key, nil
  }

  key, err = createNewAPIKey()  
  fmt.Println("Your new API key is:", key)
  return key, nil
}

func createNewAPIKey() (string, error) {
  newKey, keyStr := generateHashedAPIKey()

  err := storeAPIKey(newKey)
  if err != nil {
    return "", err
  }

  return keyStr, nil
}

func storeAPIKey(key []byte) error {
  return os.WriteFile("data", key, 0644)
}

func generateHashedAPIKey() ([]byte, string) {
  rand.Seed(time.Now().UnixNano())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, API_KEY_LENGTH)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}

	keyStr := string(s)

	hashedKey, _ := bcrypt.GenerateFromPassword([]byte(keyStr), bcrypt.DefaultCost)

	return hashedKey, keyStr
}

func readAPIKey() (string, error) {
  err := createDataFile()
    if err != nil {
      return "", err
    }

    content, err := ioutil.ReadFile("data")
    if err != nil {
      return "", err
    }

  return string(content), nil
}

func createDataFile() error {
if _, err := os.Stat("data"); errors.Is(err, os.ErrNotExist) {
		_, err := os.Create("data")
		if err != nil {
      return err
		}
	}

  return nil
}
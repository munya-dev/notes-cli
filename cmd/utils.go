package cmd

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"net/http"
	"os"
)

func addFile(f string) (result string, e error) {
	data, err := os.ReadFile(f)
	if err != nil {
		return "", fmt.Errorf("Error reading from file: ", err)

	}
	fmt.Print(data)
	// contents := []byte("New file contents")
	if err = os.WriteFile(fmt.Sprintf("note%d.txt", rand.IntN(100)), data, 0777); err != nil {
		return "", fmt.Errorf("Error writing to file: ", err)

	}
	return fmt.Sprintf("%s", string(data)), nil
}

func addRemoteFile(url string) (result string, e error) {

	res, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("Error fetching remote note", err)
	}
	var data string
	// json.Unmarshal([]byte(res.Body),)
	json.NewDecoder(res.Body).Decode(&data)
	if err = os.WriteFile(fmt.Sprintf("note%d.txt", rand.IntN(100)), []byte(data), 0777); err != nil {
		return "", fmt.Errorf("Error writing remote note", err)
	}
	return data, nil

}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// User is the info of some GitHub user.
type User struct {
	Name        string `json:"login"`
	PublicRepos int    `json:"public_repos"`
}

// userInfo retrieves the info of a GitHub user if exists.
func userInfo(login string) (*User, error) {
	uri := fmt.Sprintf("https://api.github.com/users/%s", login)
	response, err := http.Get(uri)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("error retriving info for %s", login)
	}

	dec := json.NewDecoder(response.Body)
	user := &User{}

	if err := dec.Decode(user); err != nil {
		return nil, err
	}

	return user, nil
}

func main() {
	login := "facebook"
	user, err := userInfo(login)

	if err != nil {
		log.Fatalln("error:", err)
		return
	}

	fmt.Println(*user)
}

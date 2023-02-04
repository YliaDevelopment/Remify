package utils

import (
	"encoding/json"
	"os"

	"github.com/sandertv/gophertunnel/minecraft/auth"
	"golang.org/x/oauth2"
)

func FetchToken() (*oauth2.Token, error) {
	if _, err := os.Stat("auth.json"); err != nil {
		token, err := auth.RequestLiveToken()

		if err != nil {
			return nil, err
		}

		f, err := os.Create("auth.json")
		defer f.Close()
		if err != nil {
			return nil, err
		}

		json, err := json.Marshal(token)

		if err != nil {
			return nil, err
		}

		if _, err = f.Write(json); err != nil {
			return nil, err
		}

		return token, nil
	} else {
		var token oauth2.Token
		content, err := os.ReadFile("auth.json")

		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(content, &token)

		if err != nil {
			return nil, err
		}

		return &token, nil
	}

}

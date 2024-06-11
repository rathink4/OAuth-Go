package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetUserInfo(googleAccessToken string) (map[string]any, error) {
	userEndPoint := "https://www.googleapis.com/oauth2/v2/userinfo"

	response, err := http.Get(fmt.Sprintf("%s?access_token=%s", userEndPoint, googleAccessToken))
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var userInfo map[string]any
	if err := json.NewDecoder(response.Body).Decode(&userInfo); err != nil {
		return nil, err
	}

	return userInfo, nil

}

package github

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"time"
)

func getAuthToken() (string, error) {
	token := viper.GetString("token")

	if token != "" {
		return token, nil
	}

	return "", errors.New("error getting auth token from config")
}

func IsAuthenticated() (bool, error) {
	if !viper.IsSet("token") {
		return false, errors.New("no token found in config")
	}

	if viper.GetString("token") == "" {
		return false, errors.New("no token found in config")
	}

	if viper.IsSet("authenticated") && !viper.GetBool("authenticated") {
		_, err := getInfo("https://api.github.com", viper.GetString("token"))

		if err != nil {
			return false, err
		}
	}

	return true, nil
}

func WhoAmI() (string, error) {
	var nURL = "https://api.github.com/user"
	resp, err := getInfo(nURL, "")

	if err != nil {
		fmt.Println("There was an error getting notification list.", err)
	}
	return resp, err
}

func GetUserNotifications() (string, error) {
	var nURL = "https://api.github.com/notifications"
	resp, err := getInfo(nURL, "")

	if err != nil {
		fmt.Println("There was an error getting notification list.", err)
	}
	return resp, err
}

func getInfo(url string, authToken string) (string, error) {

	if authToken == "" {
		newAuthToken, err := getAuthToken()
		if err != nil {
			fmt.Println("There was an error getting auth token.", err)
			return "", err
		}
		authToken = newAuthToken
	}

	tr := &http.Transport{
		MaxIdleConns:          10,
		IdleConnTimeout:       30 * time.Second,
		DisableCompression:    true,
		Proxy:                 http.ProxyFromEnvironment,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		MaxIdleConnsPerHost:   10,
		DisableKeepAlives:     false,
	}

	client := &http.Client{Transport: tr}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("There was an error creating get request", err)
	}
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", "Bearer "+authToken)
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("There was an error getting notification list.", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	statusCode := resp.StatusCode

	if err != nil {
		fmt.Println("There was an error reading notification list.", err)
	}

	var bodyJSON = struct {
		message string `json:"message"`
	}{}

	json.Unmarshal(body, &bodyJSON)

	if statusCode != 200 {
		return "", errors.New(bodyJSON.message)
	}

	return string(body), nil
}

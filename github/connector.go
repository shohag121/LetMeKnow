package github

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shohag121/LetMeKnow/notification"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"strings"
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

func WhoAmI() ([]byte, error) {
	var nURL = "https://api.github.com/user"
	resp, err := getInfo(nURL, "")

	if err != nil {
		fmt.Println("There was an error getting notification list.", err)
	}
	return resp, err
}

func GetUserNotifications() ([]notification.Notification, error) {
	var nURL = "https://api.github.com/notifications"
	resp, err := getInfo(nURL, "")
	var list []notification.Notification

	if err != nil {
		fmt.Println("There was an error getting notification list.", err)
		return list, err
	}

	err = json.Unmarshal(resp, &list)
	if err != nil {
		fmt.Println(err)
	}
	return list, err
}

func getInfo(url string, authToken string) ([]byte, error) {
	urlFractions := strings.Split(url, "https://api.github.com/")
	urlFraction := ""
	hasURLFraction := false

	if len(urlFractions) > 1 {
		urlFraction = urlFractions[1]
	}
	if urlFraction != "" {
		hasURLFraction = true
	}

	if authToken == "" {
		newAuthToken, err := getAuthToken()
		if err != nil {
			fmt.Println("There was an error getting auth token.", err)
			return []byte(""), err
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

	if hasURLFraction && viper.IsSet("last_"+urlFraction) && viper.GetString("last_"+urlFraction) != "" {
		req.Header.Add("If-Modified-Since", viper.GetString("last_"+urlFraction))
	}

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

	if hasURLFraction {
		lastModified := resp.Header.Get("Last-Modified")
		if lastModified != "" {
			viper.Set("last_"+urlFraction, lastModified)
		}
		if statusCode == 200 {
			viper.Set("last_result_"+urlFraction, string(body))
		}
		viper.WriteConfig()
	}

	var bodyJSON = struct {
		message string `json:"message"`
	}{}

	json.Unmarshal(body, &bodyJSON)

	if statusCode == 304 {
		lastResult := ""
		if viper.IsSet("last_result_" + urlFraction) {
			lastResult = viper.GetString("last_result_" + urlFraction)
		}

		return []byte(lastResult), nil
	}

	if statusCode != 200 {
		return []byte(""), errors.New(bodyJSON.message)
	}

	return body, nil
}

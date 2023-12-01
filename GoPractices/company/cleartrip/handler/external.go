package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// executeAPIURL call get url and return UnderlyingAsset response.
func executeAPIURL(URL string) error {
	resp, err := http.Get(URL)
	if err != nil {
		return fmt.Errorf("error while requesting")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error into response")
	}

	var data string
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Println(resp.Body)
		return fmt.Errorf("not able to parse underlying asset price response" + err.Error())
	}
	return nil
}

// executeAPIURL call get url and return UnderlyingAsset response.
func executePostURL(URL string) error {

	//resp, err := http.Post(URL, "Content-Type", nil)
	resp, err := http.Client{}
	if err != nil {
		return fmt.Errorf("error while requesting")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error into response")
	}

	var data string
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Println(resp.Body)
		return fmt.Errorf("not able to parse underlying asset price response" + err.Error())
	}
	return nil
}

func ExecuteGetUrl(URL string) error {
	fmt.Println(URL)
	res := executeAPIURL(URL)
	fmt.Println(res)
	return res
}

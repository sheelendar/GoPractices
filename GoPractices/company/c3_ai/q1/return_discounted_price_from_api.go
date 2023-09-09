package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Barcode struct {
	Page       int    `json:"page"`
	PerPage    int    `json:"per_page"`
	Total      int    `json:"total"`
	TotalPages int    `json:"total_pages"`
	Data       []Data `json:"data"`
}

type Data struct {
	Barcode   string `json:"barcode"`
	Item      string `json:"item"`
	Category  string `json:"category"`
	Price     int    `json:"price"`
	Discount  int    `json:"discount"`
	Available int    `json:"available"`
}

func main() {
	fmt.Println(getDiscountedPrice(int32(74002314)))
}

func getDiscountedPrice(barcode int32) int32 {
	var httpClint = http.DefaultClient
	baseURL := "https://jsonmock.hackerrank.com/api/inventory?barcode="
	fullURL := fmt.Sprintf("%s%d", baseURL, barcode)

	resp, err := httpClint.Get(fullURL)
	if err != nil {
		return 0
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return -1
	}
	barCodeRes := &Barcode{}
	err = json.NewDecoder(resp.Body).Decode(&barCodeRes)
	if err != nil {
		return -1
	}
	if len(barCodeRes.Data) == 1 && barCodeRes.Data[0].Barcode != "" {
		price := barCodeRes.Data[0]
		dPrice := price.Price - (price.Discount*price.Price)/100
		return int32(dPrice)
	}
	return int32(-1)
}

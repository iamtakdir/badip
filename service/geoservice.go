package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)


type DataModel struct {
	Ip       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Postal   string `json:"postal"`
	Timezone string `json:"timezone"`
}

var geoData DataModel

func GetGeoLocation( ip string)  {
	token := "36b4e33420fc30"
	url := fmt.Sprintf("https://ipinfo.io/%s?token=%s", ip, token)

	resp, err := http.Get(url)

	if err!=nil{
		log.Fatal("Error in fetching", err)
	}
//reading from response
	data,err := ioutil.ReadAll(resp.Body)
	if err != nil{
		log.Fatal("Error in Reading data")
	}
	//binding data with Data model
	err = json.Unmarshal([]byte(data), &geoData)
	if err != nil {
		log.Fatal("error in marshaling")
	}
//Printing Location details
	fmt.Println("IP : ", geoData.Ip)
	fmt.Println("City : ", geoData.City)
	fmt.Println("Region : ", geoData.Region)
	fmt.Println("Country : ", geoData.Country)
	fmt.Println("Location : ", geoData.Loc)
	fmt.Println("Company : ", geoData.Org)
	fmt.Println("Postal Code : ", geoData.Postal)
	fmt.Println("TimeZone : ", geoData.Timezone)
}

package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

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
	PrintGeoLocation(geoData)
}

func PrintGeoLocation(geolocation DataModel)  {
	fmt.Println("IP : ", geoData.Ip)
	fmt.Println("City : ", geoData.City)
	fmt.Println("Region : ", geoData.Region)
	fmt.Println("Country : ", geoData.Country)
	fmt.Println("Location : ", geoData.Loc)
	fmt.Println("ISP : ", geoData.Org)
	fmt.Println("Postal Code : ", geoData.Postal)
	fmt.Println("TimeZone : ", geoData.Timezone)
}
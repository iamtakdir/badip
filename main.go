package main

import (
	"fmt"
	"github.com/iamtakdir/badip/service"
	"os"
)

func main() {

	if len(os.Args) !=1 {
		ipAddress := os.Args[1]
		service.GetGeoLocation(ipAddress)
	}else {
		fmt.Println("Ip address can't be empty, badip<space>ipaddress ")
	}
}


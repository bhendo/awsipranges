package main

import (
	"fmt"
	"net/http"

	"./awsipranges"
)


func main() {

	data, err := awsipranges.New(http.DefaultClient)

	if err != nil { fmt.Println(err) } else {
		
		/*byRegion, _ := data.PrefixesByRegion("us-east-1")

		for i, ip_range := range byRegion {
			fmt.Printf("%d: %s\n", i, ip_range)
		}

		byService, _ := data.PrefixesByService("amazon")

		for i, ip_range := range byService {
			fmt.Printf("%d: %s\n", i, ip_range)
		}*/

		byRegionAndService, _ := data.PrefixesByRegionAndService("us-east-1", "amazon")

		for i, ip_range := range byRegionAndService {
			fmt.Printf("%d: %s\n", i, ip_range)
		}

	}
}
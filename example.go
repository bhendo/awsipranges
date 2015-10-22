package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type AWSIPRanges struct {
	SyncToken string
	CreateDate string
	Prefixes []Prefix
}

type Prefix struct {
	IP_Prefix string
	Region string
	Service string
}

func main() {

	resp, err := http.Get("https://ip-ranges.amazonaws.com/ip-ranges.json")
	if err != nil { fmt.Println(err) } else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		
		if err != nil { fmt.Println(err) } else {
			var data AWSIPRanges
			err = json.Unmarshal(body, &data)
			if err != nil {
				fmt.Printf("%T\n%s\n%#v\n",err, err, err)
		        switch v := err.(type){
		            case *json.SyntaxError:
		                fmt.Println(string(body[v.Offset-40:v.Offset]))
		        }
			}
			for i, prefix := range data.Prefixes {
				fmt.Printf("%d: %s %s %s\n", i, prefix.IP_Prefix, prefix.Region, prefix.Service)
			}
		}
	}
}
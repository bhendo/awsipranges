# awsipranges
A simple library for parsing and searching https://ip-ranges.amazonaws.com/ip-ranges.json

# Usage
```go
package main

import (
	"fmt"
	"net/http"

	"github.com/bhendo/awsipranges"
)

func main() {
	a, err := awsipranges.New(http.DefaultClient)
	if err != nil { fmt.Println(err) }
	
	fmt.Printf("SyncToken: %s and CreateDate: %s\n", a.SyncToken, a.CreateDate)
	fmt.Printf("First Prefix: %s is service: %s in region: %s\n", a.Prefixes[0].IP_Prefix, a.Prefixes[0].Service, a.Prefixes[0].Region)		

	region, err := a.PrefixesByRegion("us-east-1")
	if err != nil { fmt.Println(err) }
	
	for _, prefix := range region {
		fmt.Println(prefix)
	}
}
```

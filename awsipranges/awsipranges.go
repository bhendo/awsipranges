package awsipranges

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"strings"
)

const (
	url string = "https://ip-ranges.amazonaws.com/ip-ranges.json"
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

func New(client *http.Client) (*AWSIPRanges, error) {
	res, err := client.Get(url)

	if err != nil { return nil, err }
	
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil { return nil, err }

	var data AWSIPRanges
	err = json.Unmarshal(body, &data)

	if err != nil { return nil, err } else {
		return &data, nil
	}
}

func (a *AWSIPRanges) PrefixesByRegion(r string) ([]string, error) {
	var results []string
	for _, prefix := range a.Prefixes {
		if strings.EqualFold(prefix.Region, r) { results = append(results, prefix.IP_Prefix) }
	}
	return results, nil
}

func (a *AWSIPRanges) PrefixesByService(s string) ([]string, error) {
	var results []string
	for _, prefix := range a.Prefixes {
		if strings.EqualFold(prefix.Service, s) { results = append(results, prefix.IP_Prefix) }
	}
	return results, nil
}

func (a *AWSIPRanges) PrefixesByRegionAndService(r, s string) ([]string, error) {
	var results []string
	for _, prefix := range a.Prefixes {
		if strings.EqualFold(prefix.Service, s) && strings.EqualFold(prefix.Region, r) { results = append(results, prefix.IP_Prefix) }
	}
	return results, nil
}
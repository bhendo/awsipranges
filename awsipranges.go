// Package awsipranges provides simple parsing and searching capabilities for AWS's ip-ranges.json.
package awsipranges

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	url string = "https://ip-ranges.amazonaws.com/ip-ranges.json"
)

// AWSIPRanges contains two strings, SyncToken and CreateDate, and a slice of Prefixes.
type AWSIPRanges struct {
	SyncToken  string
	CreateDate string
	Prefixes   []Prefix
}

// Prefix contains three strings; IP_Prefix, Region, and Service.
type Prefix struct {
	IP_Prefix string
	Region    string
	Service   string
}

//New requests the latest IP Ranges from AWS using client.
//It returns a struct containing parsed json and any error encountered.
func New(client *http.Client) (*AWSIPRanges, error) {
	res, err := client.Get(url)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	var data AWSIPRanges
	err = json.Unmarshal(body, &data)

	if err != nil {
		return nil, err
	} else {
		return &data, nil
	}
}

//PrefixesByRegion operates on an AWSIPRanges to search for all prefixes in region r.
//It returns a slice of strings containing the prefix for each match.
func (a *AWSIPRanges) PrefixesByRegion(r string) ([]string, error) {
	var results []string
	for _, prefix := range a.Prefixes {
		if strings.EqualFold(prefix.Region, r) {
			results = append(results, prefix.IP_Prefix)
		}
	}
	return results, nil
}

//PrefixesByService operates on an AWSIPRanges to search for all prefixes for service s.
//It returns a slice of strings containing the prefix for each match.
func (a *AWSIPRanges) PrefixesByService(s string) ([]string, error) {
	var results []string
	for _, prefix := range a.Prefixes {
		if strings.EqualFold(prefix.Service, s) {
			results = append(results, prefix.IP_Prefix)
		}
	}
	return results, nil
}

//PrefixesByRegion operates on an AWSIPRanges to search for all prefixes in region r and for service s.
//It returns a slice of strings containing the prefix for each match.
func (a *AWSIPRanges) PrefixesByRegionAndService(r, s string) ([]string, error) {
	var results []string
	for _, prefix := range a.Prefixes {
		if strings.EqualFold(prefix.Service, s) && strings.EqualFold(prefix.Region, r) {
			results = append(results, prefix.IP_Prefix)
		}
	}
	return results, nil
}

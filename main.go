package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

const (
	baseURL = "https://iwantmyname.com/basicauth/ddns"
)

func main() {
	username := os.Getenv("IWANTMYNAME_USERNAME")
	password := os.Getenv("IWANTMYNAME_PASSWORD")

	if username == "" || password == "" {
		fmt.Println("username or password not set in environment variables")
		os.Exit(1)
	}

	if len(os.Args) < 4 {
		fmt.Println("Usage: iwantmyname hostname recordType value [ttl]")
		os.Exit(1)
	}

	hostname := os.Args[1]
	recordType := os.Args[2]
	value := os.Args[3]

	var ttl string
	if len(os.Args) > 4 {
		ttl = os.Args[4]
	} else {
		ttl = ""
	}

	err := updateRecord(username, password, hostname, recordType, value, ttl)
	if err != nil {
		fmt.Printf("Error updating %s record: %v\n", recordType, err)
		return
	}

	fmt.Printf("%s record updated successfully\n", recordType)
}

func updateRecord(username, password, hostname, recordType, value, ttl string) error {
	client := &http.Client{}
	req, err := http.NewRequest("POST", baseURL, nil)
	if err != nil {
		return err
	}

	req.SetBasicAuth(username, password)

	q := url.Values{}
	q.Add("hostname", hostname)
	q.Add("type", recordType)
	q.Add("value", value)
	if ttl != "" {
		q.Add("ttl", ttl)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("failed to update record, status code: %d, response: %s", resp.StatusCode, string(bodyBytes))
	}

	return nil
}

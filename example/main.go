package main

import (
	"fmt"
	"io/ioutil"

	"github.com/shindakun/envy"
	"github.com/shindakun/mailgunner"
)

const apiurl = "https://api.mailgun.net/v3/YOURMGACCOUNT"

func main() {
	mgKey, err := envy.Get("MGKEY")
	if err != nil {
		panic(err)
	}

	mgc := mailgunner.New(apiurl, mgKey)

	// Update these with your testing info etc.
	req, err := mgc.FormatEmailRequest("from", "to", "subject", "text")
	if err != nil {
		panic(err)
	}

	res, err := mgc.Client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}

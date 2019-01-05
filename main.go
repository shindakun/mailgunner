package mailgunner

import (
	"net/http"
	"net/url"
	"strings"
)

// MgClient struct holds our URL and API key
type MgClient struct {
	MgAPIURL string
	MgAPIKey string
	Client   *http.Client
}

// New returns our MgClient with the proper settings
func New(apiurl, apikey string, client *http.Client) *MgClient {
	if client == nil {
		client = http.DefaultClient
	}
	return &MgClient{
		apiurl,
		apikey,
		client,
	}
}

// FormatEmailRequest puts everything together for sending
func (mgc *MgClient) FormatEmailRequest(from, to, subject, body string) (r *http.Request, err error) {
	data := url.Values{}
	data.Add("from", from)
	data.Add("to", to)
	data.Add("subject", subject)
	data.Add("text", body)

	r, err = http.NewRequest(http.MethodPost, mgc.MgAPIURL+"/messages", strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	r.SetBasicAuth("api", mgc.MgAPIKey)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return r, nil
}

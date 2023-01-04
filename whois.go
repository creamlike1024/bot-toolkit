package toolkit

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"
)

type WhoisQuery struct {
	Key    string
	Domain string
}

type whoisResponse struct {
	Domain      string    `json:"domain"`
	DomainID    string    `json:"domain_id"`
	Status      string    `json:"status"`
	CreateDate  time.Time `json:"create_date"`
	UpdateDate  time.Time `json:"update_date"`
	ExpireDate  time.Time `json:"expire_date"`
	DomainAge   int       `json:"domain_age"`
	WhoisServer string    `json:"whois_server"`
	Registrar   struct {
		IanaID string `json:"iana_id"`
		Name   string `json:"name"`
		URL    string `json:"url"`
	} `json:"registrar"`
	Registrant struct {
		Name          string `json:"name"`
		Organization  string `json:"organization"`
		StreetAddress string `json:"street_address"`
		City          string `json:"city"`
		Region        string `json:"region"`
		ZipCode       string `json:"zip_code"`
		Country       string `json:"country"`
		Phone         string `json:"phone"`
		Fax           string `json:"fax"`
		Email         string `json:"email"`
	} `json:"registrant"`
	Admin struct {
		Name          string `json:"name"`
		Organization  string `json:"organization"`
		StreetAddress string `json:"street_address"`
		City          string `json:"city"`
		Region        string `json:"region"`
		ZipCode       string `json:"zip_code"`
		Country       string `json:"country"`
		Phone         string `json:"phone"`
		Fax           string `json:"fax"`
		Email         string `json:"email"`
	} `json:"admin"`
	Tech struct {
		Name          string `json:"name"`
		Organization  string `json:"organization"`
		StreetAddress string `json:"street_address"`
		City          string `json:"city"`
		Region        string `json:"region"`
		ZipCode       string `json:"zip_code"`
		Country       string `json:"country"`
		Phone         string `json:"phone"`
		Fax           string `json:"fax"`
		Email         string `json:"email"`
	} `json:"tech"`
	Billing struct {
		Name          string `json:"name"`
		Organization  string `json:"organization"`
		StreetAddress string `json:"street_address"`
		City          string `json:"city"`
		Region        string `json:"region"`
		ZipCode       string `json:"zip_code"`
		Country       string `json:"country"`
		Phone         string `json:"phone"`
		Fax           string `json:"fax"`
		Email         string `json:"email"`
	} `json:"billing"`
	Nameservers []string `json:"nameservers"`
}

func Whois(whoisQuery *WhoisQuery) (string, error) {
	endpoint := "https://api.ip2whois.com/v2?key=" + whoisQuery.Key + "&domain=" + whoisQuery.Domain
	resp, err := http.Get(endpoint)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var respBody whoisResponse
	err = json.Unmarshal(body, &respBody)
	if err != nil {
		return "", err
	}
	result := MergeText("Domain: "+respBody.Domain,
		"Domain ID: "+respBody.DomainID,
		"Status: "+respBody.Status,
		"Create Date: "+respBody.CreateDate.Format("2006-01-02 15:04:05"),
		"Update Date: "+respBody.UpdateDate.Format("2006-01-02 15:04:05"),
		"Expire Date: "+respBody.ExpireDate.Format("2006-01-02 15:04:05"),
		"Domain Age: "+strconv.Itoa(respBody.DomainAge)+" days",
		"Whois Server: "+respBody.WhoisServer,
		"Registrar: "+respBody.Registrar.Name,
		"Registrant: "+respBody.Registrant.Name,
		"Admin: "+respBody.Admin.Name,
		"Tech: "+respBody.Tech.Name,
		"Billing: "+respBody.Billing.Name,
		"Nameservers: "+StringSliceToText(respBody.Nameservers))
	return result, nil
}

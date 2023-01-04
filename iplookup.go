package toolkit

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

type IpInfo struct {
	Query         string  `json:"query"`
	Status        string  `json:"status"`
	Continent     string  `json:"continent"`
	ContinentCode string  `json:"continentCode"`
	Country       string  `json:"country"`
	CountryCode   string  `json:"countryCode"`
	Region        string  `json:"region"`
	RegionName    string  `json:"regionName"`
	City          string  `json:"city"`
	District      string  `json:"district"`
	Zip           string  `json:"zip"`
	Lat           float64 `json:"lat"`
	Lon           float64 `json:"lon"`
	Timezone      string  `json:"timezone"`
	Offset        int     `json:"offset"`
	Currency      string  `json:"currency"`
	Isp           string  `json:"isp"`
	Org           string  `json:"org"`
	As            string  `json:"as"`
	Asname        string  `json:"asname"`
	Mobile        bool    `json:"mobile"`
	Proxy         bool    `json:"proxy"`
	Hosting       bool    `json:"hosting"`
}

func GetIpInfo(ip string) (string, error) {
	// 通过 ip-api.com 查询 IP 信息
	resp, err := http.Get("http://ip-api.com/json/" + ip)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var ipInfo IpInfo
	respBody, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(respBody, &ipInfo) // 解析 json
	if err != nil {
		return "", err
	}
	if ipInfo.Status != "success" {
		return "", nil
	}
	infoText := MergeText("IP: "+ipInfo.Query,
		"运营商: "+ipInfo.Isp,
		"组织: "+ipInfo.Org,
		"ASN: "+ipInfo.As,
		"国家: "+ipInfo.Country,
		"地区: "+ipInfo.RegionName,
		"城市: "+ipInfo.City,
		"邮编: "+ipInfo.Zip,
		"时区: "+ipInfo.Timezone,
		"经度: "+strconv.FormatFloat(ipInfo.Lon, 'f', 6, 64),
		"纬度: "+strconv.FormatFloat(ipInfo.Lat, 'f', 6, 64),
	)
	return infoText, nil
}

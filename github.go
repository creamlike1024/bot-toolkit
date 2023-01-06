package toolkit

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"
)

type GithubUserInfo struct {
	Login             string      `json:"login"`
	ID                int         `json:"id"`
	NodeID            string      `json:"node_id"`
	AvatarURL         string      `json:"avatar_url"`
	GravatarID        string      `json:"gravatar_id"`
	URL               string      `json:"url"`
	HTMLURL           string      `json:"html_url"`
	FollowersURL      string      `json:"followers_url"`
	FollowingURL      string      `json:"following_url"`
	GistsURL          string      `json:"gists_url"`
	StarredURL        string      `json:"starred_url"`
	SubscriptionsURL  string      `json:"subscriptions_url"`
	OrganizationsURL  string      `json:"organizations_url"`
	ReposURL          string      `json:"repos_url"`
	EventsURL         string      `json:"events_url"`
	ReceivedEventsURL string      `json:"received_events_url"`
	Type              string      `json:"type"`
	SiteAdmin         bool        `json:"site_admin"`
	Name              string      `json:"name"`
	Company           interface{} `json:"company"`
	Blog              string      `json:"blog"`
	Location          interface{} `json:"location"`
	Email             interface{} `json:"email"`
	Hireable          interface{} `json:"hireable"`
	Bio               string      `json:"bio"`
	TwitterUsername   interface{} `json:"twitter_username"`
	PublicRepos       int         `json:"public_repos"`
	PublicGists       int         `json:"public_gists"`
	Followers         int         `json:"followers"`
	Following         int         `json:"following"`
	CreatedAt         time.Time   `json:"created_at"`
	UpdatedAt         time.Time   `json:"updated_at"`
}

func GetGithubUserInfo(username string) string {
	endpoint := "https://api.github.com/users/" + username
	resp, err := http.Get(endpoint)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	var userInfo GithubUserInfo
	err = json.Unmarshal(respBody, &userInfo)
	if err != nil {
		return ""
	}
	// 需要忽略错误，因为有些字段可能为空
	email, _ := userInfo.Email.(string)
	location, _ := userInfo.Location.(string)
	company, _ := userInfo.Company.(string)
	text := MergeText(
		"Name: "+userInfo.Name,
		"Username: "+userInfo.Login,
		"ID: "+strconv.Itoa(userInfo.ID),
		"Node ID: "+userInfo.NodeID,
		"Avatar URL: "+userInfo.AvatarURL,
		"Gravatar ID: "+userInfo.GravatarID,
		"URL: "+userInfo.URL,
		"Bio: "+userInfo.Bio,
		"Email: "+email,
		"Location: "+location,
		"Company: "+company,
		"Blog: "+userInfo.Blog,
		"Followers: "+strconv.Itoa(userInfo.Followers),
		"Following: "+strconv.Itoa(userInfo.Following),
		"Public Repos: "+strconv.Itoa(userInfo.PublicRepos),
		"Public Gists: "+strconv.Itoa(userInfo.PublicGists),
		"Created At: "+userInfo.CreatedAt.String(),
		"Updated At: "+userInfo.UpdatedAt.String(),
	)
	return text
}

func GetGithubUserSshKeys(username string) string {
	sshKeyEndpoint := "https://github.com/" + username + ".keys"
	resp, err := http.Get(sshKeyEndpoint)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return string(respBody)
}

func GetGithubUserGpgKeys(username string) string {
	gpgKeyEndpoint := "https://github.com/" + username + ".gpg"
	resp, err := http.Get(gpgKeyEndpoint)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return string(respBody)
}

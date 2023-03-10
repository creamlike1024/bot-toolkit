package toolkit

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type MisskeyAuth struct {
	URL   string
	Token string
}

type MisskeyStats struct {
	NotesCount         int `json:"notesCount"`
	OriginalNotesCount int `json:"originalNotesCount"`
	UsersCount         int `json:"usersCount"`
	OriginalUsersCount int `json:"originalUsersCount"`
	ReactionsCount     int `json:"reactionsCount"`
	Instances          int `json:"instances"`
	DriveUsageLocal    int `json:"driveUsageLocal"`
	DriveUsageRemote   int `json:"driveUsageRemote"`
}

type MisskeyInstanceInfo struct {
	Machine string `json:"machine"`
	CPU     struct {
		Model string `json:"model"`
		Cores int    `json:"cores"`
	} `json:"cpu"`
	Mem struct {
		Total int64 `json:"total"`
	} `json:"mem"`
	Fs struct {
		Total int64 `json:"total"`
		Used  int64 `json:"used"`
	} `json:"fs"`
}

type MisskeyMetaInfo struct {
	MaintainerName               string   `json:"maintainerName"`
	MaintainerEmail              string   `json:"maintainerEmail"`
	Version                      string   `json:"version"`
	Name                         string   `json:"name"`
	URI                          string   `json:"uri"`
	Description                  string   `json:"description"`
	Langs                        []string `json:"langs"`
	TosURL                       string   `json:"tosUrl"`
	RepositoryURL                string   `json:"repositoryUrl"`
	FeedbackURL                  string   `json:"feedbackUrl"`
	DefaultDarkTheme             string   `json:"defaultDarkTheme"`
	DefaultLightTheme            string   `json:"defaultLightTheme"`
	DisableRegistration          bool     `json:"disableRegistration"`
	DisableLocalTimeline         bool     `json:"disableLocalTimeline"`
	DisableGlobalTimeline        bool     `json:"disableGlobalTimeline"`
	DriveCapacityPerLocalUserMb  int      `json:"driveCapacityPerLocalUserMb"`
	DriveCapacityPerRemoteUserMb int      `json:"driveCapacityPerRemoteUserMb"`
	CacheRemoteFiles             bool     `json:"cacheRemoteFiles"`
	EmailRequiredForSignup       bool     `json:"emailRequiredForSignup"`
	EnableHcaptcha               bool     `json:"enableHcaptcha"`
	HcaptchaSiteKey              string   `json:"hcaptchaSiteKey"`
	EnableRecaptcha              bool     `json:"enableRecaptcha"`
	RecaptchaSiteKey             string   `json:"recaptchaSiteKey"`
	SwPublickey                  string   `json:"swPublickey"`
	ThemeColor                   string   `json:"themeColor"`
	MascotImageURL               string   `json:"mascotImageUrl"`
	BannerURL                    string   `json:"bannerUrl"`
	ErrorImageURL                string   `json:"errorImageUrl"`
	IconURL                      string   `json:"iconUrl"`
	BackgroundImageURL           string   `json:"backgroundImageUrl"`
	MaxNoteTextLength            int      `json:"maxNoteTextLength"`
	Emojis                       []struct {
		ID       string   `json:"id"`
		Aliases  []string `json:"aliases"`
		Category string   `json:"category"`
		Host     string   `json:"host"`
		URL      string   `json:"url"`
	} `json:"emojis"`
	Ads []struct {
		Place    string `json:"place"`
		URL      string `json:"url"`
		ImageURL string `json:"imageUrl"`
	} `json:"ads"`
	RequireSetup             bool   `json:"requireSetup"`
	EnableEmail              bool   `json:"enableEmail"`
	EnableTwitterIntegration bool   `json:"enableTwitterIntegration"`
	EnableGithubIntegration  bool   `json:"enableGithubIntegration"`
	EnableDiscordIntegration bool   `json:"enableDiscordIntegration"`
	EnableServiceWorker      bool   `json:"enableServiceWorker"`
	TranslatorAvailable      bool   `json:"translatorAvailable"`
	ProxyAccountName         string `json:"proxyAccountName"`
	Features                 struct {
		Registration   bool `json:"registration"`
		LocalTimeLine  bool `json:"localTimeLine"`
		GlobalTimeLine bool `json:"globalTimeLine"`
		Elasticsearch  bool `json:"elasticsearch"`
		Hcaptcha       bool `json:"hcaptcha"`
		Recaptcha      bool `json:"recaptcha"`
		ObjectStorage  bool `json:"objectStorage"`
		Twitter        bool `json:"twitter"`
		Github         bool `json:"github"`
		Discord        bool `json:"discord"`
		ServiceWorker  bool `json:"serviceWorker"`
		Miauth         bool `json:"miauth"`
	} `json:"features"`
}

func GetMisskeyInstanceInfo(host string) string {
	// ???????????????????????? URL???????????????????????????????????????????????? URL
	_, err := url.ParseRequestURI(host)
	if err != nil {
		host = "https://" + host
		// ?????????????????????????????? URL
		_, err = url.ParseRequestURI(host)
		if err != nil {
			return ""
		}
	}
	// Get /api/server-info
	resp, err := http.Post(host+"/api/server-info", "application/json", strings.NewReader("{}"))
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	bodyBytes, _ := io.ReadAll(resp.Body)
	var bodyMapInfo MisskeyInstanceInfo
	err = json.Unmarshal(bodyBytes, &bodyMapInfo)
	if err != nil {
		return ""
	}
	// Get /api/stats
	resp, err = http.Post(host+"/api/stats", "application/json", strings.NewReader("{}"))
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	bodyBytes, _ = io.ReadAll(resp.Body)
	var bodyMapStats MisskeyStats
	err = json.Unmarshal(bodyBytes, &bodyMapStats)
	if err != nil {
		return ""
	}

	// Get /api/meta
	resp, err = http.Post(host+"/api/meta", "application/json", strings.NewReader("{}"))
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	bodyBytes, _ = io.ReadAll(resp.Body)
	var bodyMapMeta MisskeyMetaInfo
	err = json.Unmarshal(bodyBytes, &bodyMapMeta)
	if err != nil {
		return ""
	}

	info := MergeText(
		"Misskey ????????????????????????",
		"????????????????????????"+bodyMapInfo.Machine,
		"CPU: "+bodyMapInfo.CPU.Model,
		"????????????"+strconv.Itoa(bodyMapInfo.CPU.Cores),
		"?????????"+strconv.FormatInt(bodyMapInfo.Mem.Total/1024/1024, 10)+"MB",
		"???????????????"+strconv.FormatInt(bodyMapInfo.Fs.Total/1024/1024, 10)+"MB",
		"???????????????"+strconv.FormatInt(bodyMapInfo.Fs.Used/1024/1024, 10)+"MB",
		"",
		"Misskey ???????????????",
		"????????????"+strconv.Itoa(bodyMapStats.UsersCount),
		"??????????????????"+strconv.Itoa(bodyMapStats.OriginalUsersCount),
		"????????????"+strconv.Itoa(bodyMapStats.NotesCount),
		"??????????????????"+strconv.Itoa(bodyMapStats.OriginalNotesCount),
		"??????????????????"+strconv.Itoa(bodyMapStats.Instances),
		"???????????????????????????"+strconv.FormatInt(int64(bodyMapStats.DriveUsageLocal)/1024/1024, 10)+"MB",
		"???????????????????????????"+strconv.FormatInt(int64(bodyMapStats.DriveUsageRemote)/1024/1024, 10)+"MB",
		"",
		"Misskey ??????????????????",
		"????????????"+bodyMapMeta.Name,
		"?????????"+bodyMapMeta.Description,
		"URI: "+bodyMapMeta.URI,
		"?????????"+bodyMapMeta.Version,
		"???????????? "+bodyMapMeta.MaintainerName,
		"???????????????????????? "+bodyMapMeta.MaintainerEmail,
		"?????? URL???"+bodyMapMeta.FeedbackURL,
		"???????????????"+bodyMapMeta.TosURL,
		"???????????????"+bodyMapMeta.RepositoryURL,
		"?????????"+bodyMapMeta.IconURL,
		"?????????"+strings.Join(bodyMapMeta.Langs, ", "),
		"???????????????"+strconv.FormatBool(!bodyMapMeta.DisableRegistration),
		"????????????????????????"+strconv.FormatBool(bodyMapMeta.Features.LocalTimeLine),
		"????????????????????????"+strconv.FormatBool(bodyMapMeta.Features.GlobalTimeLine),
		"????????????????????????????????????"+strconv.FormatInt(int64(bodyMapMeta.DriveCapacityPerLocalUserMb), 10)+"MB",
		"????????????????????????????????????"+strconv.FormatInt(int64(bodyMapMeta.DriveCapacityPerRemoteUserMb), 10)+"MB",
		"???????????????????????????"+strconv.FormatBool(bodyMapMeta.CacheRemoteFiles),
		"?????????????????????"+strconv.FormatBool(bodyMapMeta.EmailRequiredForSignup),
		"?????? Hcaptcha???"+strconv.FormatBool(bodyMapMeta.EnableHcaptcha),
		"Hcaptcha Site Key???"+bodyMapMeta.HcaptchaSiteKey,
		"?????? ReCaptcha???"+strconv.FormatBool(bodyMapMeta.EnableRecaptcha),
		"ReCaptcha Site Key???"+bodyMapMeta.RecaptchaSiteKey,
		"swPublickey???"+bodyMapMeta.SwPublickey,
		"????????????"+bodyMapMeta.ThemeColor,
		"??????????????? URL???"+bodyMapMeta.MascotImageURL,
		"?????? URL???"+bodyMapMeta.BannerURL,
		"???????????? URL???"+bodyMapMeta.ErrorImageURL,
		"???????????? URL???"+bodyMapMeta.BackgroundImageURL,
		"?????? URL???"+bodyMapMeta.IconURL,
		"?????????????????????"+strconv.Itoa(bodyMapMeta.MaxNoteTextLength),
		"?????? elasticsearch???"+strconv.FormatBool(bodyMapMeta.Features.Elasticsearch),
		"?????????????????????"+strconv.FormatBool(bodyMapMeta.Features.ObjectStorage),
		"?????? Twitter ?????????"+strconv.FormatBool(bodyMapMeta.Features.Twitter),
		"?????? GitHub ?????????"+strconv.FormatBool(bodyMapMeta.Features.Github),
		"?????? Discord ?????????"+strconv.FormatBool(bodyMapMeta.Features.Discord),
		"?????? Service Worker???"+strconv.FormatBool(bodyMapMeta.Features.ServiceWorker),
		"?????? Miauth???"+strconv.FormatBool(bodyMapMeta.Features.Miauth),
	)
	return info
}

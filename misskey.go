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
	// 检查是否为有效的 URL，如果不是，尝试将其转换为有效的 URL
	_, err := url.ParseRequestURI(host)
	if err != nil {
		host = "https://" + host
		// 再次检查是否为有效的 URL
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
		"Misskey 实例服务器信息：",
		"主机名／容器名："+bodyMapInfo.Machine,
		"CPU: "+bodyMapInfo.CPU.Model,
		"核心数："+strconv.Itoa(bodyMapInfo.CPU.Cores),
		"内存："+strconv.FormatInt(bodyMapInfo.Mem.Total/1024/1024, 10)+"MB",
		"存储容量："+strconv.FormatInt(bodyMapInfo.Fs.Total/1024/1024, 10)+"MB",
		"存储已用："+strconv.FormatInt(bodyMapInfo.Fs.Used/1024/1024, 10)+"MB",
		"",
		"Misskey 实例信息：",
		"用户数："+strconv.Itoa(bodyMapStats.UsersCount),
		"本地用户数："+strconv.Itoa(bodyMapStats.OriginalUsersCount),
		"帖子数："+strconv.Itoa(bodyMapStats.NotesCount),
		"本地帖子数："+strconv.Itoa(bodyMapStats.OriginalNotesCount),
		"联合实例数："+strconv.Itoa(bodyMapStats.Instances),
		"本地网盘已用容量："+strconv.FormatInt(int64(bodyMapStats.DriveUsageLocal)/1024/1024, 10)+"MB",
		"远程网盘已用容量："+strconv.FormatInt(int64(bodyMapStats.DriveUsageRemote)/1024/1024, 10)+"MB",
		"",
		"Misskey 实例元数据：",
		"实例名："+bodyMapMeta.Name,
		"描述："+bodyMapMeta.Description,
		"URI: "+bodyMapMeta.URI,
		"版本："+bodyMapMeta.Version,
		"管理员： "+bodyMapMeta.MaintainerName,
		"管理员联系方式： "+bodyMapMeta.MaintainerEmail,
		"反馈 URL："+bodyMapMeta.FeedbackURL,
		"用户协议："+bodyMapMeta.TosURL,
		"源码地址："+bodyMapMeta.RepositoryURL,
		"图标："+bodyMapMeta.IconURL,
		"语言："+strings.Join(bodyMapMeta.Langs, ", "),
		"开放注册："+strconv.FormatBool(!bodyMapMeta.DisableRegistration),
		"启用本地时间线："+strconv.FormatBool(bodyMapMeta.Features.LocalTimeLine),
		"启用全局时间线："+strconv.FormatBool(bodyMapMeta.Features.GlobalTimeLine),
		"每个本地用户的网盘容量："+strconv.FormatInt(int64(bodyMapMeta.DriveCapacityPerLocalUserMb), 10)+"MB",
		"每个远程用户的网盘容量："+strconv.FormatInt(int64(bodyMapMeta.DriveCapacityPerRemoteUserMb), 10)+"MB",
		"是否缓存远程文件："+strconv.FormatBool(bodyMapMeta.CacheRemoteFiles),
		"注册需要邮箱："+strconv.FormatBool(bodyMapMeta.EmailRequiredForSignup),
		"使用 Hcaptcha："+strconv.FormatBool(bodyMapMeta.EnableHcaptcha),
		"Hcaptcha Site Key："+bodyMapMeta.HcaptchaSiteKey,
		"使用 ReCaptcha："+strconv.FormatBool(bodyMapMeta.EnableRecaptcha),
		"ReCaptcha Site Key："+bodyMapMeta.RecaptchaSiteKey,
		"swPublickey："+bodyMapMeta.SwPublickey,
		"主题色："+bodyMapMeta.ThemeColor,
		"吉祥物图像 URL："+bodyMapMeta.MascotImageURL,
		"横幅 URL："+bodyMapMeta.BannerURL,
		"错误图像 URL："+bodyMapMeta.ErrorImageURL,
		"背景图像 URL："+bodyMapMeta.BackgroundImageURL,
		"图标 URL："+bodyMapMeta.IconURL,
		"最大贴文字数："+strconv.Itoa(bodyMapMeta.MaxNoteTextLength),
		"使用 elasticsearch："+strconv.FormatBool(bodyMapMeta.Features.Elasticsearch),
		"使用对象存储："+strconv.FormatBool(bodyMapMeta.Features.ObjectStorage),
		"支持 Twitter 登录："+strconv.FormatBool(bodyMapMeta.Features.Twitter),
		"支持 GitHub 登录："+strconv.FormatBool(bodyMapMeta.Features.Github),
		"支持 Discord 登录："+strconv.FormatBool(bodyMapMeta.Features.Discord),
		"支持 Service Worker："+strconv.FormatBool(bodyMapMeta.Features.ServiceWorker),
		"支持 Miauth："+strconv.FormatBool(bodyMapMeta.Features.Miauth),
	)
	return info
}

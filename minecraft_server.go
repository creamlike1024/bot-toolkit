package toolkit

import (
	"encoding/json"
	"net/http"
	"strconv"
)

const (
	JavaEditionEndpoint    = "https://api.mcsrvstat.us/2/"
	BedrockEditionEndpoint = "https://api.mcsrvstat.us/bedrock/2/"
)

func GetMinecraftServerInfo(ip string, isBedrock bool) (string, error) {
	// 通过 api.mcsrvstat.us 查询 Minecraft 服务器信息
	var endpoint string
	// 根据 isBedrock 判断是 Java 还是 Bedrock 服务器
	if isBedrock {
		endpoint = BedrockEditionEndpoint + ip
	} else {
		endpoint = JavaEditionEndpoint + ip
	}
	resp, err := http.Get(endpoint)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	response := make(map[string]interface{})
	// 解析 json 到 response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return "", err
	}
	if response["online"].(bool) {
		// 服务器在线
		return MergeText(
			"服务器在线",
			"版本: "+response["version"].(string),
			"在线玩家: "+strconv.FormatFloat(response["players"].(map[string]interface{})["online"].(float64), 'f', -1, 64),
			"最大玩家: "+strconv.FormatFloat(response["players"].(map[string]interface{})["max"].(float64), 'f', -1, 64),
			"服务器 MOTD: \n"+getMotd(response),
		), nil
	} else {
		// 服务器离线
		return "", nil
	}
}

func getMotd(response map[string]interface{}) string {
	var motd string
	// 从 response 中取出 motd 作为 string 返回
	for _, v := range response["motd"].(map[string]interface{})["clean"].([]interface{}) {
		motd += v.(string) + "\n"
	}
	return motd
}

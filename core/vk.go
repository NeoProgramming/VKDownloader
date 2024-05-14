package core

import (
	"github.com/SevereCloud/vksdk/v2/api"
	"net/url"
	"vkdownloader/cfmt"
)

type RecordAttrs int

func InitVK() {
	cfmt.PrintlnFunc("VK API library initializing...")
	App.vk = api.NewVK(App.config.AccessToken)
	cfmt.PrintlnLine("VK API library initialized")
}

func extractAccessToken(urlStr string) string {
	u, _ := url.Parse(urlStr)
	parameters, _ := url.ParseQuery(u.Fragment)
	accessToken := parameters.Get("access_token")
	return accessToken
}

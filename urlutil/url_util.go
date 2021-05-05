package urlutil

import (
	"net/url"
	"strings"
)

//判断url是否属于指定站点
func IsOwnSiteUrl(requestUrl string, siteUrl string) bool {
	return strings.ToLower(GetUrlTopDomain(requestUrl)) == strings.ToLower(GetUrlTopDomain(siteUrl))
}

//给url增加query参数
func UrlAddQueryParam(urlStr string, key string, val string) string {
	symbol := "?"
	if strings.Contains(urlStr, "?") {
		symbol = "&"
	}
	if !strings.Contains(urlStr, symbol+key+"="+val) {
		urlStr += symbol + key + val
	}

	return urlStr
}

//获取url的主机
func GetUrlHost(urlStr string) string {
	u, err := url.Parse(urlStr)
	if err != nil {
		panic(err)
	}

	return u.Host
}

//是否为https
func UrlIsHttps(urlStr string) bool {
	u, err := url.Parse(urlStr)
	if err != nil {
		panic(err)
	}

	return u.Scheme == "https" || u.Scheme == "HTTPS"
}

//获取url的顶级域名
func GetUrlTopDomain(urlStr string) string {
	host := GetUrlHost(urlStr)
	arr := strings.Split(host, ".")

	return strings.Join(arr[len(arr)-2:], ".")
}

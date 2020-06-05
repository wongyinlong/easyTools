package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	IntranetV6 = "[::1]"
)

/**
 * 获取IP地址
 *
 * 使用Nginx等反向代理软件， Ctx.getRemoteAddr()获取IP地址
 * 如果使用了多级反向代理的话，X-Forwarded-For的值并不止一个，而是一串IP地址，X-Forwarded-For中第一个非unknown的有效IP字符串，则为真实IP地址
 */
func RemoteIp(r *http.Request) string {
	ipAddr := r.Header.Get("x-forwarded-for")
	if len(ipAddr) <= 0 || "unknown" == strings.ToLower(ipAddr) {
		ipAddr = r.Header.Get("X-Forwarded-For")
	}

	if len(ipAddr) <= 0 || "unknown" == strings.ToLower(ipAddr) {
		ipAddr = r.Header.Get("X-Real-Ip")
	}

	if len(ipAddr) <= 0 || "unknown" == strings.ToLower(ipAddr) {
		ipAddr = r.Header.Get("x-real-ip")
	}

	if len(ipAddr) <= 0 || "unknown" == strings.ToLower(ipAddr) {
		ipAddr = r.Header.Get("HTTP_X_FORWARDED_FOR")
	}

	if len(ipAddr) <= 0 || "unknown" == strings.ToLower(ipAddr) {
		ipAddr = r.Header.Get("HTTP_CLIENT_IP")
	}

	if len(ipAddr) <= 0 || "unknown" == strings.ToLower(ipAddr) {
		ipAddr = r.Header.Get("WL-Proxy-Client-IP")
	}

	if len(ipAddr) <= 0 || "unknown" == strings.ToLower(ipAddr) {
		ipAddr = r.RemoteAddr
	}

	return ipAddr
}

//过滤网段
func IpSelected(ip string) string {

	if strings.Contains(ip, IntranetV6) {
		return ""
	}

	var ips []string
	if strings.Contains(ip, ",") {
		ips = strings.Split(ip, ",")
	}

	if len(ips) <= 0 {
		return ""
	}

	return ips[0]
}

func GetAddressOfMobile(mobile string) string {
	url := "http://api.online-service.vip/phone?number=" + mobile

	res, err := http.Get(url)
	if nil != err {
		log.Println("发送请求失败：", err)
		return ""
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if nil != err {
		log.Println("解析返回值失败：", err)
		return ""
	}

	type response struct {
		Province   string `json:"province"`
		City       string `json:"city"`
		Postcode   string `json:"postcode"`
		Areacode   string `json:"areacode"`
		Mobiletype string `json:"mobiletype"`
	}

	resp := response{}
	if err = json.Unmarshal(body, &resp); err != nil {
		log.Println("解析返回值失败：", err)
		return ""
	}

	return resp.City + "市"
}

func GetAddressOfIp(ip string) (string, string) {

	if len(ip) <= 0 {
		return "", ""
	}

	//1.查询IP地址的位置
	url := "http://ip-api.com/json/" + ip + "?lang=zh-CN"

	res, err := http.Get(url)
	if nil != err {
		log.Println("发送请求失败：", err)
		return "", ""
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if nil != err {
		log.Println("解析返回值失败：", err)
		return "", ""
	}

	type response struct {
		As          string  `json:"as"`
		City        string  `json:"city"`
		Country     string  `json:"country"`
		CountryCode string  `json:"countryCode"`
		Isp         string  `json:"isp"`
		Lat         float64 `json:"lat"`
		Lon         float64 `json:"lon"`
		Org         string  `json:"org"`
		Query       string  `json:"query"`
		Region      string  `json:"region"`
		RegionName  string  `json:"regionName"`
		Status      string  `json:"status"`
		Timezone    string  `json:"timezone"`
		Zip         string  `json:"zip"`
	}

	resp := response{}
	if err = json.Unmarshal(body, &resp); err != nil {
		log.Println("解析返回值失败：", err)
		return "", ""
	}

	return resp.Country, resp.City
}

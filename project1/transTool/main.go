package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

type RESPONSE struct {
	From         string              `json:"from"`
	To           string              `json:"to"`
	Trans_result []map[string]string `json:"trans_result"`
}

func main() {
	file := "./files/file.txt"
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("read file error")
		return
	}
	for i := 0; i < len(content); i++ {
		if content[i] == 13 && i+1 < len(content) && content[i+1] == 10 {
			temp := content[i+2:]
			content = append(append(content[:i], 32), temp...)
		}
	}
	container := string(content)
	rand.Seed(time.Now().UnixNano())
	random := strconv.Itoa(rand.Int())
	params := map[string]string{
		"q":     url.QueryEscape(container),
		"from":  "en",
		"to":    "zh",
		"appid": "20201009000584455",
		"salt":  random,
		"sign":  fmt.Sprintf("%x", md5.Sum([]byte("20201009000584455"+container+random+"4ISQsHwS68UABPBSJzKe"))),
	}
	requestStr := HTTPparam2str(params)
	resp, err := http.Get("http://api.fanyi.baidu.com/api/trans/vip/translate?" + requestStr)
	if err != nil {
		fmt.Println("get error")
		return
	}
	fmt.Println("1")
	fmt.Println(resp)
	robots, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Println("read error")
		return
	}
	var res RESPONSE
	json.Unmarshal(robots, &res)
	ans := []byte(res.Trans_result[0]["dst"])
	ioutil.WriteFile(file, ans, os.ModeAppend)
}

func HTTPparam2str(params map[string]string) string {
	str := ""
	for k, v := range params {
		str += (k + "=" + v) + "&"
	}
	return str
}

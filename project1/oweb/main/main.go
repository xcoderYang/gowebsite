package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, _ := http.Get("http://www.baidu.com")
	byteStatus := resp.Status

	rebots, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(rebots))
	fmt.Println(byteStatus)
	defer resp.Body.Close()
}

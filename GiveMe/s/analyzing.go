package s

import (
	"bufio"
	"encoding/base64"
	"log"
	"net/http"
	"strings"
)

func Analyzing(postDate string) {
	//postDateBytes, _ := url.QueryUnescape(postDate)
	postDateByte, err := base64.StdEncoding.DecodeString(postDate)
	//log.Println(string(postDateByte))
	if err != nil {
		log.Println("数据包base64解码不成功")
		return
	}
	request, err := http.ReadRequest(bufio.NewReader(strings.NewReader(string(postDateByte))))
	if err != nil {
		log.Println("数据包解析不成功")
		return
	}
	log.Println(request.Method, request.Host, request.URL)
	host := request.Host
	if CheckExclude(host) == true {
		return
	}
	HostPath = host
	if strings.Contains(HostPath, ":") {
		HostPath = strings.Replace(HostPath, ":", "_", -1)
	}
	//检查存放结果的文件夹,如果不存在就创建
	checkDir(HostPath)
	//处理url,将get请求和post请求的url分开存放
	//fmt.Println("请求方式:", request.Method)
	switch strings.ToUpper(request.Method) {
	case "GET":
		Get(request)
		ProcessHeader("GET", request)
		ProcessGetParameter(request)
	case "POST":
		Post(request)
		ProcessHeader("POST", request)
	default:
		Other(request)
	}

	Wg.Wait()
}

package s

import (
	"net/http"
	"strings"
)

func ProcessHeader(method string, req *http.Request) {
	//fmt.Println("调用ProcessHeader")
	switch method {
	case "GET":
		ProcessGetHeader(req)
	case "POST":
		ProcessPostHeader(req)
	}

}
func ProcessPostHeader(req *http.Request) {
	//fmt.Println("调用ProcessPostHeader")
	Write(HwaderList(req.Header), HostPath+"/headers.txt")
	//获取body内容
	body := make([]byte, req.ContentLength)
	_, _ = req.Body.Read(body)
	bodys := string(body)
	va := determineCT(req.Header)
	switch va {
	case "JSON":
		ProcessJSONParameter(bodys)
	case "XML":
		ProcessXMLParameter(bodys)
	case "FILE":
		ProcessFILEParameter(bodys)
	default:
		ProcessParameter(bodys)
	}
}

// 判断是否有Content-Type
func determineCT(header http.Header) string {
	ct := strings.ToLower(header.Get("Content-Type"))
	if strings.Contains(ct, "application/json") {
		return "JSON"
	} else if strings.Contains(ct, "application/xml") || strings.Contains(ct, "text/xml") {
		return "XML"
	} else if strings.Contains(ct, "multipart/form-data") {
		return "FILE"
	} else {
		return ""
	}
}

func ProcessGetHeader(req *http.Request) {
	Write(HwaderList(req.Header), HostPath+"/headers.txt")
}
func HwaderList(header http.Header) string {
	var content, contents string
	for parameter, va := range header {
		var value string
		for _, v := range va {
			value = value + v
		}
		content = parameter + ": " + value + "\n"
		contents = contents + content
	}
	return contents
}

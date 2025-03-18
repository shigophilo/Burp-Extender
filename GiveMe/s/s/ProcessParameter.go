package s

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"net/http"
	"strings"
)

// ===========================================================================解析json
func ProcessJSONParameter(body string) {
	log.Println("Processing JSON data")
	Write(body, HostPath+"/josnParameterANDvalue.txt")
	var data map[string]interface{}
	err := json.Unmarshal([]byte(body), &data)
	if err != nil {
		log.Println("Error parsing JSON:", err)
		return
	}
	// 遍历 map 中的键值对
	parseData(data)
}
func parseData(data map[string]interface{}) {
	for key, value := range data {
		// 如果值是 map 类型，继续递归解析
		if subMap, ok := value.(map[string]interface{}); ok {
			parseData(subMap)
		} else {
			// 打印参数名和参数值
			log.Printf("parameter：%s, value：%v\n", key, value)
			Write(key, HostPath+"/parameter.txt")
		}
	}
}

// ===========================================================================解析json结束
// ===========================================================================解析xml
func ProcessXMLParameter(body string) {
	log.Println("Processing XML data")
	//log.Println(body)
	Write(body, HostPath+"/xmlParameterANDvalue.txt")
	decoder := xml.NewDecoder(strings.NewReader(body))
	for {
		token, err := decoder.Token()
		if err != nil {
			break
		}
		switch se := token.(type) {
		case xml.StartElement:
			log.Printf("parameter: %s\n", se.Name.Local)
			Write(se.Name.Local, HostPath+"/parameter.txt")
		}
	}
}

// ===========================================================================解析xml结束
func ProcessParameter(body string) {
	//log.Println("处理default")
	//log.Println(body)
	params := strings.Split(string(body), "&")
	for _, param := range params {
		pair := strings.Split(param, "=")
		if len(pair) == 2 {
			log.Printf("Param: %s, Value: %s\n", pair[0], pair[1])
			Write(pair[0], HostPath+"/parameter.txt")
			Write(pair[0]+"="+pair[1], HostPath+"/postParameterANDvalue.txt")
		}
	}
}

// ===========================================================================解析GET参数
func ProcessGetParameter(req *http.Request) {
	//requestLine := req.Method + " " + req.URL.String() + " " + req.Proto
	//log.Println("请求行:", requestLine)

	// 提取查询字符串（从第一个?到最后一个空格）
	query := req.URL.RawQuery
	//log.Println("查询字符串:", query)
	Write(query, HostPath+"/getParameterANDvalue.txt")

	// 使用&分割查询字符串
	params := strings.Split(query, "&")

	// 提取参数名（不是值）
	paramNames := make([]string, 0, len(params))
	for _, param := range params {
		pair := strings.Split(param, "=")
		paramNames = append(paramNames, pair[0])
	}

	// 打印参数名
	//log.Println("参数名:", paramNames)
	for _, v := range paramNames {
		Write(v, HostPath+"/parameter.txt")
	}
}

// ===========================================================================解析文件上传
func ProcessFILEParameter(bodys string) {
	log.Println("Processing UPDATE data")
	Write(bodys, HostPath+"/update.txt")
}

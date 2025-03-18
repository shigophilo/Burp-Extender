package s

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

func CheckExclude(host string) bool {
	for _, url := range ExcludeHost {
		if url == host {
			return true
		}
	}
	return false
}

func ListExcludeHost() {
	url_file, err := os.Open("exclude.txt")
	if err != nil {
		log.Println("打开排除主机文件\"exclude.txt\"失败")
		return
	}
	defer url_file.Close()
	reader_Url := bufio.NewReader(url_file)
	for {
		url, err := reader_Url.ReadString('\n')
		url = strings.Replace(url, " ", "", -1)
		url = strings.Replace(url, "\n", "", -1)
		url = strings.Replace(url, "\r", "", -1)
		url = strings.Replace(url, "https://", "", -1)
		url = strings.Replace(url, "http://", "", -1)
		ExcludeHost = append(ExcludeHost, url)
		if err == io.EOF {
			break
		}
	}

}

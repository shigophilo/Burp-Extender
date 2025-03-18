package s

import (
	"log"
	"os"
)

func checkDir(host string) {
	if fileExists(host) == false {
		CreateResultsFolder(host)
	}
}

func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false // 文件不存在
	}
	return true // 文件存在
}

func CreateResultsFolder(host string) {

	err := os.MkdirAll(host, os.ModePerm)
	if err != nil {
		log.Println("创建用于保存结果的文件夹失败")
	}
}

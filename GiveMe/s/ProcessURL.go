package s

import (
	"net/http"
	"os"
)

func Get(req *http.Request) {
	url := req.RequestURI
	Write(url, HostPath+"/uris_get.txt")
}
func Post(req *http.Request) {
	url := req.RequestURI
	Write(url, HostPath+"/uris_post.txt")
}

func Other(req *http.Request) {
	url := req.RequestURI
	Write(url, HostPath+"/uris_other.txt")
}

func Write(content, path string) {
	lock.Lock()
	var ok *os.File
	ok, _ = os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	ok.Write([]byte(content + "\r\n"))
	defer ok.Close()
	lock.Unlock()
}

// func Write(path string, content ...string) {
// 	var contents string
// 	for _, v := range content {
// 		contents = contents + v
// 	}
// 	lock.Lock()
// 	var ok *os.File
// 	ok, _ = os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
// 	ok.Write([]byte(contents + "\r\n"))
// 	defer ok.Close()
// 	lock.Unlock()
// }

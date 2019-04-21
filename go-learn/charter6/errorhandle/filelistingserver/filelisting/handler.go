package filelisting

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

type userError string

func (e userError) Error() string  {
	return e.Message()
}

func (e userError) Message() string  {
	return string(e)
}

func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	urlPath := request.URL.Path
	if strings.Index(urlPath, prefix) != 0 {
		return userError("path must start with " + prefix)
	}
	path := urlPath[len(prefix):]
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	writer.Write(all)
	return nil
}

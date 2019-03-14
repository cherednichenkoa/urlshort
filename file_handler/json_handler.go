package file_handler

import (
	"encoding/json"
	"fmt"
	"gopherex/urlshort/config"
	"io/ioutil"
)

type JsonHandler struct {
	Settings config.Settings
}

func (handler JsonHandler) GetFileContent() [] UrlRewriteInterface {
	file, err := ioutil.ReadFile(handler.Settings.GetFilePath())
	if err != nil {
		fmt.Println("Error during json file reading.")
		panic(err)
	}
	data, err := handler.unpackJson(file)
	if err != nil {
		fmt.Println("Json content is invalid.")
		panic(err)
	}
	return data
}


func (handler *JsonHandler ) unpackJson(jsonData []byte) ( [] UrlRewriteInterface, error) {
	var out [] UrlMappingJson
	if err := json.Unmarshal(jsonData, &out); err != nil {
		return nil, err
	}
	// Re-convert items due to this https://github.com/golang/go/wiki/InterfaceSlice
	items := make([]UrlRewriteInterface,len(out))
	for i, item := range out {
		items[i] = item
	}
	return items, nil
}

type UrlMappingJson struct  {
	Path string `json:"url"`
	Url string  `json:"path"`
}

func (url UrlMappingJson) GetUrl() string {
	return url.Url
}

func (url UrlMappingJson) GetPath() string {
	return url.Path
}

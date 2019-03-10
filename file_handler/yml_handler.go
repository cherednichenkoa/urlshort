package file_handler

import (
	"fmt"
	"gopherex/urlshort/config"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type YmlHandler struct {
	Settings config.Settings
}

func (handler YmlHandler) GetFileContent() [] UrlRewriteInterface {
	return handler.parseContent()
}

func (handler *YmlHandler) parseContent() [] UrlRewriteInterface {
	file, err := ioutil.ReadFile(handler.Settings.GetFilePath())
	if err != nil {
		fmt.Println("Error during yml file reading.")
		panic(err)
	}
	data, err := handler.unpackYml(file)
	if err != nil {
		fmt.Println("Yml content is invalid.")
		panic(err)
	}
	return data
}

func (handler *YmlHandler ) unpackYml(yml []byte) ( [] UrlRewriteInterface, error) {
	var out [] UrlMappingYaml
	if err := yaml.Unmarshal(yml,&out); err != nil {
		return nil, err
	}
	// Re-convert items due to this https://github.com/golang/go/wiki/InterfaceSlice
	items := make([]UrlRewriteInterface,len(out))
	for i, item := range out {
		items[i] = item
	}
	return items, nil
}

type UrlMappingYaml struct  {
	Path string `yaml:"url"`
	Url string  `yaml:"path"`
}

func (url UrlMappingYaml) GetUrl() string {
	return url.Url
}

func (url UrlMappingYaml) GetPath() string {
	return url.Path
}

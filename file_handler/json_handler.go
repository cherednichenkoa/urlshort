package file_handler

import "gopherex/urlshort/config"

type JsonHandler struct {
	Settings config.Settings
}

func (s JsonHandler) GetFileContent() [] UrlRewriteInterface {
	// TODO implement json content parser
	var data [] UrlRewriteInterface
	return data
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

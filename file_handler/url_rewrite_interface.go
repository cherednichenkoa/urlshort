package file_handler

type UrlRewriteInterface interface {
	GetUrl() string
	GetPath() string
}
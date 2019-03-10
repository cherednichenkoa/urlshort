package main

import (
	"flag"
	"fmt"
	"gopherex/urlshort/file_handler"
	"net/http"
	"gopherex/urlshort"
	"gopherex/urlshort/config"
)

var (
	filePath = flag.String("filePath","url_data.yml","path to the yml file")
	handlerType = flag.String("handlerType", config.TYPE_YML_FILE,"type of the url handler")
)


func main() {
	flag.Parse()
	settings := config.Settings{*filePath, *handlerType}
	handler := getHandlerByType(settings)
	content := handler.GetFileContent()
	mux := defaultMux()
	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	fileHandler, err := urlshort.UrlRewriteHandler(content,mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", fileHandler)
}


func getHandlerByType(settings config.Settings) file_handler.FileHandler {
	if settings.GetHandlerType() == config.TYPE_YML_FILE {
		return file_handler.YmlHandler{settings}
	}
	return file_handler.JsonHandler{settings}
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}


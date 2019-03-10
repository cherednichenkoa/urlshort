package config

const (
	TYPE_YML_FILE = "yml"
	TYPE_JSON_FILE = "json"
)

type Settings struct {
	FilePath string
	HandlerType string
}

func (s *Settings) GetFilePath() string {
	return s.FilePath
}

func (s *Settings) GetHandlerType() string {
	return s.HandlerType
}
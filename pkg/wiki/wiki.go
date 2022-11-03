package wiki

type Wiki interface {
	Group() string
	RelativePath() string
	Description() string
	HelpLink() string
	Handle(path, group string) error
	Score() float64
}

func HasWiki(relativePath string, group string) bool {
	return true
}

func GetWiki(relativePath string, group string) []Wiki {
	return nil
}

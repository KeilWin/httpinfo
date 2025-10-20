package defaults

const (
	spaDirTempaltePath = "./dist"
	spaTemplateFile    = "index.html"
)

func GetSpaDirTemplatePath() string {
	return spaDirTempaltePath
}

func GetSpaTemplateFile() string {
	return spaTemplateFile
}

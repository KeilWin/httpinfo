package defaults

const (
	indexTemplatePath     = "web/template/index.html"
	appTemplatePath       = "web/template/app.html"
	headerTemplatePath    = "web/template/header.html"
	contentTemplatePath   = "web/template/content.html"
	footerTemplatePath    = "web/template/footer.html"
	leftSideTemplatePath  = "web/template/leftSide.html"
	rightSideTemplatePath = "web/template/rightSide.html"
)

func GetIndexTemplatePath() string {
	return indexTemplatePath
}

func GetAppTemplatePath() string {
	return appTemplatePath
}

func GetHeaderTemplatePath() string {
	return headerTemplatePath
}

func GetContentTemplatePath() string {
	return contentTemplatePath
}

func GetFooterTemplatePath() string {
	return footerTemplatePath
}

func GetLeftSideTemplatePath() string {
	return leftSideTemplatePath
}

func GetRightSideTemplatePath() string {
	return rightSideTemplatePath
}

package defaults

const (
	indexTemplate     = "web/template/index.html"
	appTemplate       = "web/template/app.html"
	headerTemplate    = "web/template/header.html"
	contentTemplate   = "web/template/content.html"
	footerTemplate    = "web/template/footer.html"
	leftSideTemplate  = "web/template/leftSide.html"
	rightSideTemplate = "web/template/rightSide.html"
)

func GetIndexTemplate() string {
	return indexTemplate
}

func GetAppTemplate() string {
	return appTemplate
}

func GetHeaderTemplate() string {
	return headerTemplate
}

func GetContentTemplate() string {
	return contentTemplate
}

func GetFooterTemplate() string {
	return footerTemplate
}

func GetLeftSideTemplate() string {
	return leftSideTemplate
}

func GetRightSideTemplate() string {
	return rightSideTemplate
}

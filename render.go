package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func getTemplateHeader() string {
	return fmt.Sprintf(loadedConfig.TemplatesPath + "/header.html")
}

func getTemplateInclude() string {
	return fmt.Sprintf(loadedConfig.TemplatesPath + "/include.html")
}

func getTemplateFooter() string {
	return fmt.Sprintf(loadedConfig.TemplatesPath + "/footer.html")
}

func getTemplatePage(templateName string) string {
	return fmt.Sprintf(loadedConfig.TemplatesPath+"/%s.html", templateName)
}

func renderHTMLTemplate(w http.ResponseWriter, templateName string, pageData interface{}) error {

	// Preparing html templates
	templateHeader, err := template.ParseFiles(getTemplateHeader())
	if err != nil {
		return err
	}

	templateFooter, err := template.ParseFiles(getTemplateFooter())
	if err != nil {
		return err
	}

	templateMain, err := template.ParseFiles(getTemplatePage(templateName), getTemplateInclude())
	if err != nil {
		return err
	}

	// Executing HTML rendering
	err = templateHeader.Execute(w, pageData)
	if err != nil {
		return err
	}

	err = templateMain.Execute(w, pageData)
	if err != nil {
		return err
	}

	err = templateFooter.Execute(w, pageData)
	if err != nil {
		return err
	}

	return nil
}

// ErrorPage - shows error page
func ErrorPage(w http.ResponseWriter, errHeader string, errText string) {
	var errData ErrorPageData
	log.Println(errText)
	errData.ErrorHeader = errHeader
	errData.ErrorMessage = errText
	err := renderHTMLTemplate(w, "error", errData)
	if err != nil {
		log.Fatal(err.Error())
	}
}

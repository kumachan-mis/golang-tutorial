package main

import (
	"errors"
	"html/template"
	"net/http"
	"path"
	"regexp"
)

var templates = template.Must(template.ParseFiles(
	path.Join("templates", "view.html"), path.Join("templates", "edit.html")))
var validPathRegex = regexp.MustCompile("^/(view|edit|save)/([A-Za-z][A-Za-z0-9]*)$")

func MakeHandler(handler func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		title, err := getPageTitle(responseWriter, request)
		if err != nil {
			return
		}
		handler(responseWriter, request, title)
	}
}

func ViewHandler(responseWriter http.ResponseWriter, request *http.Request, title string) {
	wikipage, err := LoadWikiPage(title)
	if err != nil {
		http.Redirect(responseWriter, request, "/edit/"+title, http.StatusFound)
	}
	renderTemplate(responseWriter, "view", wikipage)
}

func EditHandler(responseWriter http.ResponseWriter, request *http.Request, title string) {
	wikipage, err := LoadWikiPage(title)
	if err != nil {
		wikipage = &WikiPage{Title: title}
	}
	renderTemplate(responseWriter, "edit", wikipage)
}

func SaveHandler(responseWriter http.ResponseWriter, request *http.Request, title string) {
	body := request.FormValue("body")
	wikipage := &WikiPage{Title: title, Body: []byte(body)}
	err := wikipage.Save()
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(responseWriter, request, "/view/"+title, http.StatusFound)
}

func getPageTitle(responseWriter http.ResponseWriter, request *http.Request) (string, error) {
	match := validPathRegex.FindStringSubmatch(request.URL.Path)
	if match == nil {
		return "", errors.New("invalid Page Title")
	}
	return match[2], nil
}

func renderTemplate(responseWriter http.ResponseWriter, templateName string, wikipage *WikiPage) {
	err := templates.ExecuteTemplate(responseWriter, templateName+".html", wikipage)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
}

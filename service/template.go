package service

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
)

var activationMail *template.Template

var baseTemplateDir string

func init() {
	wd, _ := os.Getwd()
	baseTemplateDir = filepath.Join(wd, "templates")
	setActivationMailTemplate()
}

func setActivationMailTemplate() {
	filename := filepath.Join(baseTemplateDir, "ActivationMail.html")
	t, err := template.ParseFiles(filename)
	if err != nil {
		panic(err)
	}
	activationMail = t
}

func GetActivationMail(name, url string) (string, error) {
	b := new(bytes.Buffer)

	data := struct {
		Name string
		URL  string
	}{
		Name: name,
		URL:  url,
	}

	if err := activationMail.Execute(b, data); err != nil {
		return "", err
	}
	by, err := ioutil.ReadAll(b)
	if err != nil {
		return "", err
	}
	return string(by), nil
}

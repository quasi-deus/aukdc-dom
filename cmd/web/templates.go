package main

import (
	"html/template"
	"io/fs"
	"path/filepath"
	"time"
	

	"aukdc.dom.com/ui"
	"aukdc.dom.com/internal/models"
)
type templateData struct {
	CurrentYear int
	Honorarium *models.Honorarium
	Honoraria []*models.Honorarium
	QPK *models.QPK
	VP *models.ValuedPaper
	Course *models.Course
	Courses []*models.Course
	Faculty *models.Faculty
	Faculties []*models.Faculty
	BankDetails *models.BankDetails
	Form any
	Flash string
	IsAuthenticated bool
	IsAuthorized bool
	CSRFToken string
}

func humanDate(t time.Time)string{
	if t.IsZero(){
		return ""
	}

	return t.UTC().Format("15:05 02/Jan/2006")
}
var functions=template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache()  (map[string]*template.Template, error){
	cache:=map[string]*template.Template{}
	
	pages, err := fs.Glob(ui.Files,"html/pages/*.tmpl")
	if err != nil {
		return nil, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		
		patterns:=[]string{
			"html/base.tmpl",
			"html/printb.tmpl",
			"html/partials/*.tmpl",
			page,
		}
		ts,err:=template.New(name).Funcs(functions).ParseFS(ui.Files, patterns...)
		if err != nil {
			return nil, err
		}
		
		cache[name] = ts
	}
	return cache, nil
}




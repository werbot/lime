package server

import (
	"fmt"
	"html/template"
	"time"
)

func formatAsDate(t time.Time) string {
	return fmt.Sprintf("%d.%02d.%2d %02d:%02d:%02d",
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second())
}

func columnStatus(status bool) template.HTML {
	resultat := ""
	if status {
		resultat = "<span class=\"column-green\">true</span>"
	} else {
		resultat = "<span class=\"column-red\">false</span>"
	}
	return template.HTML(resultat)
}

func keyBytesToString(data []byte) template.HTML {
	return template.HTML(string(data[:]))
}

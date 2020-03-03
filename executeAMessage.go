package itswizard_handlingerrors

import (
	"fmt"
	"github.com/itslearninggermany/itswizard_basic"
	"github.com/jinzhu/gorm"
	"html/template"
	"net/http"
)

/*
Wenn true ausgegeben wird, war err != nil.
Es wird eine Seite ausgegebn wenn es einen Fehler gibt. mit den entsprechenden Nachrichten.
Zur Handhabung. Es muss nach der Ausführung ein "return" folgen, damit die Seite funktioniert!
*/
func ExecuteAMessage(tpl *template.Template, dbWebserver *gorm.DB, language string, sitename string, w http.ResponseWriter) {
	var dbHtmlContent itswizard_basic.DbHtmlContent
	var site itswizard_basic.Site
	if dbWebserver.Where("name = ? AND language = ?", sitename, language).Find(&dbHtmlContent).RecordNotFound() {
		site = itswizard_basic.Site{
			Sitename: sitename,
			Special: itswizard_basic.MessageStruct{
				Headline:     "headline",
				Message:      "message",
				TargetClose:  "targelclose",
				TargetSubmit: "targetSubmit",
				Buttontext:   "buttontext",
			},
		}
	} else {
		site = itswizard_basic.Site{
			Sitename: sitename,
			Special: itswizard_basic.MessageStruct{
				Headline:     dbHtmlContent.Field0,
				Message:      dbHtmlContent.Field1,
				TargetClose:  dbHtmlContent.Field2,
				TargetSubmit: dbHtmlContent.Field3,
				Buttontext:   dbHtmlContent.Field4,
			},
		}
	}
	er := tpl.ExecuteTemplate(w, "message.html", site)
	if er != nil {
		WritingToErrorLog(dbWebserver, "", fmt.Sprint(er))
	}
}

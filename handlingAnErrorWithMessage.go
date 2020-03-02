package itswizard_handlingerrors

import (
	"github.com/itslearninggermany/itswizard_basic"
	"github.com/jinzhu/gorm"
	"net/http"
	"html/template"
)

/*
Wenn true ausgegeben wird, war err != nil.
Es wird eine Seite ausgegebn wenn es einen Fehler gibt. mit den entsprechenden Nachrichten.
Zur Handhabung. Es muss nach der Ausführung ein "return" folgen, damit die Seite funktioniert!
 */
func handlingAnErrorWithMessage (tpl *template.Template, webserverdatabase *gorm.DB, err error, sitename string, headline string, message string, targelclose string, targetSubmit string, buttontext string, user string, w http.ResponseWriter) bool {
	if err != nil {
		site := itswizard_basic.Site{
			Sitename:    sitename,
			Special:     itswizard_basic.MessageStruct{
				Headline:     headline,
				Message:      message,
				TargetClose:  targelclose,
				TargetSubmit: targetSubmit,
				Buttontext:   buttontext,
			},
		}

		er := tpl.ExecuteTemplate(w, "message.html", site)
		if er != nil {
			writingToErrorLog(webserverdatabase,"", err.Error())
		}
		return true
	} else {
		return false
	}
}


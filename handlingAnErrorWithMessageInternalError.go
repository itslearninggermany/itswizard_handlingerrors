package itswizard_handlingerrors

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"net/http"
	"html/template"
)

/*
Es wird eine Seite ausgegeben mit dem Hinweis auf einen internen Fehler! Wird in ThereISAnErrorHandkeIt asugeführt.
 */
func handlingAnErrorWithMessageInternalError (tpl *template.Template, webserverdatabase *gorm.DB, err error, w http.ResponseWriter) bool{
	return handlingAnErrorWithMessage(tpl, webserverdatabase, err, "Interner Fehler", "Interner Fehler", "Leider ist ein interner Fehler aufgetreten. Er wird dokumentiert und untersucht.", "/client/startpage", "/client/startpage", "Zur Startseite", "admin", w)
}

/*
Aufzurufen statt if err != nil. Es wird der Fehler geloggt und eine Nachricht für den Client angezeigt.
z.B.:
if ThereIsAnErrorHandleIt {
	return
}
*/
func ThereIsAnErrorHandleIt (err error, tpl *template.Template, webserverdatabase *gorm.DB, w http.ResponseWriter) bool {
	if err != nil {
		writingToErrorLog(webserverdatabase, "", fmt.Sprint(err))
		handlingAnErrorWithMessageInternalError(tpl, webserverdatabase, err, w)
		return true
	}
	return false
}

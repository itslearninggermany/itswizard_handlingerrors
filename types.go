package itswizard_handlingerrors

import "github.com/jinzhu/gorm"

/*
For the Database.
Ist für allgemeine Informationen
*/
type Log struct {
	gorm.Model
	Username string `gorm:"not null"`
	Content  string `gorm:"type:MEDIUMTEXT"`
}

/*
For the Database. ITs für Fehler.
 */
type DbErrorLog struct {
	gorm.Model
	InstanceID string `gorm:"not null"`
	User       string `gorm:"type:MEDIUMTEXT"`
	Content    string `gorm:"not null"`
}
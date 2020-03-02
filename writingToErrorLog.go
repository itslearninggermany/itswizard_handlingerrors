package itswizard_handlingerrors

import (
	"fmt"
	"github.com/itslearninggermany/itswizard_aws"
	"github.com/jinzhu/gorm"
)

/*
The content will be stored in the errorLog.
An error in this function will be printed out in standartoutput
*/
func WritingToErrorLog(db *gorm.DB, user string, content string) {

	err := db.Save(&DbErrorLog{
		InstanceID: itswizard_aws.GetInstance(),
		User:       user,
		Content:    content,
	}).Error
	if err != nil {
		fmt.Println(err)
	}
}

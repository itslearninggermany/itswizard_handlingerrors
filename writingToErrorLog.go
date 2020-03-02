package itswizard_handlingerrors

import (
	"github.com/itslearninggermany/itswizard_aws"
	"github.com/jinzhu/gorm"
	"fmt"
	)

/*
The content will be stored in the errorLog.
 */
func writingToErrorLog (db *gorm.DB, user string, content string) {

	err := db.Save(&DbErrorLog{
		InstanceID: itswizard_aws.GetInstance(),
		User:       user,
		Content:    content,
	}).Error
	if err != nil {
		fmt.Println(err)
	}
}
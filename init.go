package authservices

import (
	"github.com/Dparty/dao/auth"
	"gorm.io/gorm"
)

func Init(inject *gorm.DB) {
	auth.Init(inject)
}

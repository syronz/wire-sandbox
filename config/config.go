package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

// CFG is used as a global struct but it is injected
type CFG struct {
	DB     *gorm.DB
	Log    *logrus.Logger
	Logapi *logrus.Logger
}

// Debug print struct with details with logrus ability
func (c *CFG) Debug(objs ...interface{}) {
	for _, v := range objs {
		c.Log.Debug(fmt.Sprintf("%T :: %+[1]v", v))
	}
}
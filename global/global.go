package global

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"xx/config"
	"xx/forms"
)

var MysqlConn *gorm.DB
var RedisClient *redis.Client
var ServerConfig = &config.ServerConfig{}
var Trans ut.Translator
var SugarLogger *zap.SugaredLogger
var StatusUpdates = make(chan *forms.UserStatus, 2000)
var Exception = make(chan *forms.Exception, 1000)

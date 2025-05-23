package global

import (
	"github.com/gin-gonic/gin"
	"github.com/qiniu/qmgo"
	"meet_directly/config"
	wchatClient "meet_directly/pkg/min_app"
	"sync"

	"github.com/songzhibin97/gkit/cache/local_cache"

	"golang.org/x/sync/singleflight"

	"go.uber.org/zap"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	GVA_DB                  *gorm.DB
	GVA_DBList              map[string]*gorm.DB
	GVA_REDIS               redis.UniversalClient
	GVA_MONGO               *qmgo.QmgoClient
	GVA_CONFIG              config.Server
	GVA_VP                  *viper.Viper
	GVA_LOG                 *zap.Logger
	GVA_Concurrency_Control = &singleflight.Group{}
	GVA_ROUTERS             gin.RoutesInfo
	BlackCache              local_cache.Cache
	lock                    sync.RWMutex
	WChatLoginClient        wchatClient.MinLoginClient
)

// GetGlobalDBByDBName 通过名称获取db list中的db
func GetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	return GVA_DBList[dbname]
}

// MustGetGlobalDBByDBName 通过名称获取db 如果不存在则panic
func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := GVA_DBList[dbname]
	if !ok || db == nil {
		panic("db no init")
	}
	return db
}

func NewWChatMinLogin() wchatClient.MinLoginClient {
	cfg := GVA_CONFIG.WChatLogin
	return wchatClient.NewWChatMinLogin(wchatClient.WChatLogin{
		AppId:   cfg.AppId,
		Secret:  cfg.Secret,
		BaseUrl: cfg.BaseUrl,
	})
}

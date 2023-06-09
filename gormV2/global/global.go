package global

import (
	"gV2/config"
	"gV2/middleware/cache"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GVA_DB     *gorm.DB
	GVA_REDIS  *cache.RCache
	GVA_CONFIG *config.Specification
	GVA_LOG    *zap.Logger
)

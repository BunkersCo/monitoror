//go:generate mockery -name Usecase

package config

import (
	"github.com/monitoror/monitoror/api/config/models"
)

type (
	Usecase interface {
		GetConfigList() *models.ConfigList
		GetConfig(params *models.ConfigParams) *models.ConfigBag
		Verify(config *models.ConfigBag)
		Hydrate(config *models.ConfigBag)
	}
)

package usecase

import (
	"github.com/monitoror/monitoror/api/config/models"
)

// GetConfig and set default value for Config from repository
func (cu *configUsecase) GetConfigList() *models.ConfigList {
	configList := models.ConfigList{}

	for configName := range cu.namedConfigs {
		configList = append(configList, string(configName))
	}

	return &configList
}

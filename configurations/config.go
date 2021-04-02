package configurations

import (
	"admin/types"
	"fmt"
	"sync"

	"github.com/spf13/viper"
	"gopkg.in/ini.v1"
)

type HostelPgConfigs struct {
	initOnce sync.Once
	config   *ini.File
}

var once sync.Once
var configInstance *HostelPgConfigs

func InitConfigs() *HostelPgConfigs {
	once.Do(func() {
		configInstance := &HostelPgConfigs{}
		configInstance.config = nil
	})
	return configInstance
}

func (fc *HostelPgConfigs) getConfigFromFile() (*types.PostgresInstance, error) {
	// define name of the config file
	viper.SetConfigName("config")
	// define the relative path where this file should be looked for
	viper.AddConfigPath("./configurations")
	// enable reading env vars (optional)
	viper.AutomaticEnv()
	// define your config files type
	viper.SetConfigType("json")

	// reading config file and error check
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	pgins := &types.PostgresInstance{}

	// unmarshalling the config in a go struct
	err := viper.Unmarshal(pgins)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}
	// return the config
	return pgins, nil
}

func (fc *HostelPgConfigs) GetPostgresIni(section string) (*types.PostgresInstance, error) {
	resp, err := fc.getConfigFromFile()

	if err != nil {
		return nil, err
	}
	// return the config to dblayer
	return resp, err
}

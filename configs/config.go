package configs

import (
	"barbershop/creativo/internal/validation"
	"barbershop/creativo/pkg/types"
	"reflect"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type (
	AppConfig struct {
		Auth  AuthConfig
		Email EmailConfig
	}

	AuthConfig struct {
		JWTSecret types.JWTSecret `mapstructure:"JWT_SECRET"`
	}

	EmailConfig struct {
		Host        string             `mapstructure:"EMAIL_HOST"`
		Port        int                `mapstructure:"EMAIL_PORT"`
		FromName    string             `mapstructure:"EMAIL_FROM_NAME"`
		FromAddress types.EmailAddress `mapstructure:"EMAIL_FROM_ADDRESS"`
		Password    string             `mapstructure:"EMAIL_PASSWORD"`
	}
)

func InitConfig(configFile string) *AppConfig {
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetConfigFile(configFile)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	var decoderOpts viper.DecoderConfigOption = func(dc *mapstructure.DecoderConfig) {
		dc.DecodeHook = func(from, to reflect.Type, i interface{}) (interface{}, error) {
			if from.Kind() == reflect.String && to == reflect.TypeOf(types.EmailAddress("")) {
				emailAddress, err := validation.NewEmailAddress(i.(string))
				if err != nil {
					return nil, err
				}

				return emailAddress, nil
			}
			return i, nil
		}
	}

	var config AppConfig
	if err := viper.Unmarshal(&config.Email, decoderOpts); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&config.Auth); err != nil {
		panic(err)
	}

	return &config
}

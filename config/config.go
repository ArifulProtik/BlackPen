package config

import "github.com/spf13/viper"

type Config struct {
	ServerConfig `mapstructure:",squash"`
	Postegres    `mapstructure:",squash"`
}
type ServerConfig struct {
	App     string `mapstructure:"APP"`
	Version string `mapstructure:"VERSION"`
	Status  string `mapstructure:"STATUS"`
	Port    string `mapstructure:"PORT"`
	Jwt_Key string `mapstructure:"JWT_KEY"`
}
type Postegres struct {
	DB_HOST string `mapstructure:"DB_HOST"`
	DB_PORT string `mapstructure:"DB_PORT"`
	DB_NAME string `mapstructure:"DB_NAME"`
	DB_USER string `mapstructure:"DB_USER"`
	DB_PASS string `mapstructure:"DB_PASS"`
}

// New maps all enviornment variable to Config Struct Using gomapstructure https://github.com/mitchellh/mapstructure
func New(path string, filename string, configtype string) (*Config, error) {
	var c Config
	v := viper.New()
	v.AddConfigPath(path)
	v.SetConfigName(filename)
	v.SetConfigType(configtype)
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = v.Unmarshal(&c)
	if err != nil {
		return nil, err
	}
	return &c, nil

}

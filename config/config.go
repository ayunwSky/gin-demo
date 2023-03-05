package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Oss OssConfig
}

type OssConfig struct {
	Endpoint        string `yaml:"endpoint"`
	AccessKeyID     string `yaml:"accessKeyID"`
	AccessKeySecret string `yaml:"accessKeySecret"`
	BucketName      string `yaml:"bucketName"`
}

func LoadConfig(configPath string) (*Config, error) {
	// 使用 viper 解析 配置文件
	viper.SetConfigType("yaml")
	viper.SetConfigFile(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshall config: %v", err)
	}

	return &config, nil
}

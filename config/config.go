package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	MysqlUri     int    `env:"MYSQL_URI"`
	S3BucketName string `env:"S3_BUCKET_NAME"`
	S3Region     string `env:"S3_REGION"`
	S3ApiKey     string `env:"S3_API_KEY"`
	S3SecretKey  string `env:"S3_SECRET_KEY"`
	S3Domain     string `env:"S3_DOMAIN"`
	SystemSecret string `env:"SYSTEM_SECRET"`
}

var (
	once sync.Once
	cfg  *Config
)

func NewConfig() (*Config, error) {
	once.Do(func() {
		cfg = &Config{}
		if err := cleanenv.ReadEnv(cfg); err != nil {
			help, _ := cleanenv.GetDescription(cfg, nil)
			log.Fatalf("Failed to read configuration: %v\n%v", err, help)
		}
	})
	return cfg, nil
}

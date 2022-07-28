package models

type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DbName   string `yaml:"dbname"`
		Port     string `yaml:"port"`
	}
}

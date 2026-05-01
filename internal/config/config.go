package config

var Conf Config

type Config struct {
	Application Application `yaml:"application"`
}

type Application struct {
	Datasource Datasource `yaml:"datasource"`
}

type Datasource struct {
	DbName string `yaml:"db-name"`
	Url    string `yaml:"db-url"`
}

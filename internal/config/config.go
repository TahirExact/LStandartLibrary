package config

type Config struct {
	Application Application `yaml:"application"`
}

type Application struct {
	Datasource Datasource `yaml:"datasource"`
}

type Datasource struct {
	UserName     string `yaml:"user-name"`
	UserPassword string `yaml:"user-pass"`
	HostName     string `yaml:"db-host"`
	DbNameName   string `yaml:"db-name"`
}

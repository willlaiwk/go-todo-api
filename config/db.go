package config

type mysqlConfig struct {
	Name     string
	User     string
	Password string
	Host     string
	Port     string
}

var DB = mysqlConfig{
	Name:     "todo",
	User:     "todo",
	Password: "todo123",
	Host:     "127.0.0.1",
	Port:     "3306",
}

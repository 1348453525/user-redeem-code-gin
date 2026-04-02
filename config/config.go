package config

type Config struct {
	Server Server `mapstructure:"server" json:"server"`
	MySQL  MySQL  `mapstructure:"mysql" json:"mysql"`
	Redis  Redis  `mapstructure:"redis" json:"redis"`
}

type Server struct {
	Name    string `mapstructure:"name" json:"name"`
	Version string `mapstructure:"version" json:"version"`
	Addr    string `mapstructure:"addr" json:"addr"`
	Mode    string `mapstructure:"mode" json:"mode"`
}

type MySQL struct {
	Host         string `mapstructure:"host" json:"host"`
	Port         string `mapstructure:"port" json:"port"`
	User         string `mapstructure:"user" json:"user"`
	Password     string `mapstructure:"password" json:"password"`
	DB           string `mapstructure:"db" json:"db"`
	Charset      string `mapstructure:"charset" json:"charset"`
	MaxIdleConns int    `mapstructure:"max_idle_conns" json:"max_idle_conns"`
	MaxOpenConns int    `mapstructure:"max_open_conns" json:"max_open_conns"`
}

type Redis struct {
	Addr     string `mapstructure:"addr" json:"addr"`
	Password string `mapstructure:"password" json:"password"`
}

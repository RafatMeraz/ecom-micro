package configs

type Config struct {
	HTTPServer HTTPServerConfig
	DB         DBConfig
}

type HTTPServerConfig struct {
	ListenAddr string
}

type DBConfig struct {
	Address  string
	Port     int
	Password string
	Name     string
	User     string
}

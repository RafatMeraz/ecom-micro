package configs

type Config struct {
	HTTPServer HTTPServerConfig
	DB         DBConfig
	RateLimit  RateLimit
}

type HTTPServerConfig struct {
	ListenAddr string
}

type DBConfig struct {
	Host     string
	Port     int
	Password string
	Name     string
	User     string
}

type RateLimit struct {
	LimitPerMin int
}

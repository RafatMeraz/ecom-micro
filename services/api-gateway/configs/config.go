package configs

type Config struct {
	HTTPServer HTTPServerConfig
	RateLimit  RateLimit
}

type HTTPServerConfig struct {
	ListenAddr string
}

type RateLimit struct {
	LimitPerMin int
}

package config

type Config struct {
	Env string `yaml: "env" env:"ENV" env-default:"local"`
	StoragePath string `yaml: "storage_path" env-required:"true"`	
}

type HTTPServer struct {
	Address string `yaml: "address" env-default:":8080"`
	Timeout time.Duration `yaml: "timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml: "idle_timeout" env-default:"60s"`
}
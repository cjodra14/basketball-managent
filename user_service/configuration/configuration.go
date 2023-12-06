package configuration

type Configuration struct {
	RESTServer RESTServer
	Logs       Logs
}

type Logs struct {
	LogLevel string `envconfig:"LOG_LEVEL" default:"debug"`
}

type RESTServer struct {
	Address string `envconfig:"REST_SERVER_ADDRESS" default:"localhost"`
	Port    string `envconfig:"REST_SERVER_PORT" default:"8080"`
}

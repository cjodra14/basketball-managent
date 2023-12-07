package configuration

type Configuration struct {
	RESTServer RESTServer
	Logs       Logs
	GRPCServer GRPCServer
	Storage    Storage
}

type Logs struct {
	LogLevel string `envconfig:"LOG_LEVEL" default:"debug"`
}

type RESTServer struct {
	Address string `envconfig:"REST_SERVER_ADDRESS" default:"localhost"`
	Port    int64  `envconfig:"REST_SERVER_PORT" default:"8080"`
}

type GRPCServer struct {
	Address string `envconfig:"GRPC_SERVER_ADDRESS" default:"localhost"`
	Port    int64  `envconfig:"GRPC_SERVER_PORT" default:"3000"`
}

type Storage struct {
	Type                  string `envconfig:"STORAGE_TYPE" default:"memory"`
	PostgresConfiguration PostgresConfiguration
}

type PostgresConfiguration struct {
	URI string `envconfig:"STORAGE_URI" default:"postgredb://localhost:5432"`
}

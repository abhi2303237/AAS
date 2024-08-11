package config

type Config struct {
	Port            string `env:"PORT"`
	ApplicationName string `env:"APPLICATION_NAME"`
	EnableMetricsUi bool   `env:"ENABLE_METRICS_UI" envDefault:"false"`
	LogLevel        string `env:"LOG_LEVEL" envDefault:"info"`
	ContextPath     string `env:"CONTEXT_PATH" envDefault:"/app"`

	Backend string `env:"BACKEND" envDefault:"memmory"`

	MongoConnectionString string `env:"MONGO_CON_STRING"`

	DBServer   string `env:"MSSQL_SERVER"`
	DBPort     int    `env:"MSSQL_PORT"`
	DBUser     string `env:"MSSQL_USER"`
	DBPassword string `env:"MSSQL_PASS"`
	DBDatabase string `env:"MSSQL_DATABASE"`
}

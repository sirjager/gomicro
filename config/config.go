package config

import (
	"fmt"
	"time"

	"github.com/sirjager/gomicro/pkg/db"
	"github.com/sirjager/gomicro/pkg/email"
	"github.com/spf13/viper"
)

// Config represents the application configuration.
type Config struct {
	StartTime   time.Time // StartTime is the timestamp when the application started.
	ServiceName string    // ServiceName is the name of the service.

	GrpcPort    string `mapstructure:"GRPC_PORT"`    // GrpcPort is the port number for the gRPC server.
	GatewayPort string `mapstructure:"GATEWAY_PORT"` // RestPort is the port number for the REST server.

	Mail     email.Config // Mail holds the configuration for mail-related settings.
	Database db.Config    // DBConfig holds the configuration for the database.
}

// LoadConfigs loads the configuration from the specified YAML file.
func LoadConfigs(path string, name string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return
	}
	if err = viper.Unmarshal(&config); err != nil {
		return
	}

	if err = viper.Unmarshal(&config.Mail); err != nil {
		return
	}

	if err = viper.Unmarshal(&config.Database); err != nil {
		return
	}

	// Construct the DBUrl using the DBConfig values.
	config.Database.Url = fmt.Sprintf("%s://%s:%s@%s:%s/%s%s", config.Database.Driver, config.Database.User, config.Database.Pass, config.Database.Host, config.Database.Port, config.Database.Name, config.Database.Args)

	config.Database.Migrate = "file://" + config.Database.Migrate
	return
}

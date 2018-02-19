package cmd

import (
	"log"
	"os"

	"github.com/spf13/viper"
	"github.com/uniplaces/logfairy/infrastructure/api"
	"github.com/uniplaces/logfairy/infrastructure/api/dashboard"
	srv "github.com/uniplaces/logfairy/infrastructure/api/service"
	"github.com/uniplaces/logfairy/infrastructure/api/stream"
	"github.com/uniplaces/logfairy/infrastructure/api/widget"
)

// Dependencies of the application
type Dependencies struct {
	DashboardClient dashboard.Client
	StreamClient    stream.Client
	WidgetClient    widget.Client
}

// setup reads in config file and ENV variables if set.
func Setup() Dependencies {
	viper.SetConfigFile("./logfairy.yaml")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("no config file found")
	}

	service := srv.New(viper.GetString("client.base_url"), viper.GetInt("client.timeout"))
	graylog := api.New(service, os.Getenv("GRAYLOG_USERNAME"), os.Getenv("GRAYLOG_PASSWORD"))

	return Dependencies{
		DashboardClient: dashboard.New(graylog),
		StreamClient:    stream.New(graylog),
		WidgetClient:    widget.New(graylog),
	}
}

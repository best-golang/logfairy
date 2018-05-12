package cmd

import (
	"log"
	"os"

	"github.com/spf13/viper"
	"github.com/uniplaces/logfairy/infrastructure/api"
	dapi "github.com/uniplaces/logfairy/infrastructure/api/dashboard"
	srv "github.com/uniplaces/logfairy/infrastructure/api/service"
	sapi "github.com/uniplaces/logfairy/infrastructure/api/stream"
	wapi "github.com/uniplaces/logfairy/infrastructure/api/widget"
)

// Dependencies of the application
type Dependencies struct {
	DashboardClient dapi.Dashboard
	StreamClient    sapi.Stream
	WidgetClient    wapi.Widget
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
		DashboardClient: dapi.New(service, &graylog),
		StreamClient:    sapi.New(service, &graylog),
		WidgetClient:    wapi.New(service, &graylog),
	}
}

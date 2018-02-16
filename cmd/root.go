package cmd

import (
	"encoding/json"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/uniplaces/logfairy/infrastructure/api"
	srv "github.com/uniplaces/logfairy/infrastructure/api/service"
)

var (
	appConfigFile string
	definitions   string
)

// infrastructure
var (
	service srv.Service
	graylog api.Graylog
)

// clients
var (
	dashboardClient api.DashboardClient
	streamClient    api.StreamClient
	widgetClient    api.WidgetClient
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "logfairy",
	Short: "logfairy to standardize interaction with graylog",
	Long:  `logfairy goal is to standardize the way streams, dashboard and widget are created`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}

func init() {
	initConfig()
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetConfigFile("./graylog.yaml")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		return
	}

	service = srv.New(viper.GetString("client.base_url"), viper.GetInt("client.timeout"))
	graylog = api.New(service, os.Getenv("GRAYLOG_USERNAME"), os.Getenv("GRAYLOG_PASSWORD"))
}

func prettyPrint(object interface{}) {
	prettyObject, err := json.MarshalIndent(object, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}

	println(string(prettyObject))
}

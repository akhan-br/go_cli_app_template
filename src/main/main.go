package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{})
	// Output to stderr instead of stdout, could also be a file.
	log.SetOutput(os.Stderr)
	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

func main() {

	/*
		If you would like to use config file, uncomment this section and set the location for the config file

		viper.SetConfigName("config") // name of config file (without extension)
		viper.AddConfigPath("$PWD/")  // call multiple times to add many search paths
		viper.AddConfigPath(".")      // optionally look for config in the working directory
		err := viper.ReadInConfig()   // Find and read the config file
		if err != nil {               // Handle errors reading the config file
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	*/

	app := cli.NewApp()
	app.Name = "TODO"
	app.Usage = "TODO"
	app.Version = "1.0"
	app.ArgsUsage = "main -h ..."

	app.Flags = []cli.Flag{
	// cli.StringFlag{
	// 	Name:        "tag, t",
	// 	Usage:       "tag for the build",
	// 	Destination: &tag,
	// },
	}

	app.Action = func(c *cli.Context) {
		// Call your app from here
	}
	app.Run(os.Args)
}

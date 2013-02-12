package gatsby

import "github.com/kylelemons/go-gypsy/yaml"
import "fmt"
import "strconv"

type Application struct {
	Config * yaml.File
	Host string
	Port int64
	PublicDir string
}

type DatabaseHandle struct {

}

func (app * Application) LoadConfig (configFile string) (* yaml.File) {
	config, err := yaml.ReadFile(configFile)
	if err != nil {
		fmt.Println("YAML Error: %s", err)
	}
	return config
}

func (app * Application) CreateDatabaseHandle() {

}

func NewApplication(configFile string) (*Application) {
	app := new(Application)
	app.Config = app.LoadConfig(configFile)

	// read Config data into application stash.
	host, err := app.Config.Get("Host")
	if err != nil {
		fmt.Println("Database user is not defined.")
	}
	app.Host = host

	port, err   := app.Config.Get("Port")
	if err != nil {
		fmt.Println("Database user is not defined.")
	}
	app.Port , err = strconv.ParseInt(port,10,32)

	app.PublicDir, err := app.Config.Get("Public")

	return app
}

func init() {

}


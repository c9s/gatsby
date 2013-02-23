package gatsby

import "github.com/kylelemons/go-gypsy/yaml"
// import "encoding/json"
// import "bitbucket.org/zombiezen/yaml"
import "errors"
import "strconv"

type Application struct {
	Config * yaml.File
	Host string
	Port int64
	PublicDir string
}

type DatabaseHandle struct {

}

type ApplicationConfig struct {
	ApplicationName string
	Port string
	Host string
	Public string
}

func init() {

}

func (app * Application) LoadConfig (configFile string) (error) {
	config, err := yaml.ReadFile(configFile)
	if err != nil {
		return err
	}

	app.Config = config

	// read Config data into application stash.
	host, err := app.Config.Get("Host")
	if err != nil {
		return errors.New("You should provide a host in your config file.")
	}
	app.Host = host
	port, err   := app.Config.Get("Port")
	if err != nil {
		return errors.New("Database user is not defined.")
	}
	app.Port , err = strconv.ParseInt(port,10,32)
	if err != nil {
		return errors.New("Can not read server port config.")
	}
	app.PublicDir, err = app.Config.Get("Public")
	if err != nil {
		return errors.New("Can not read server public directory config.")
	}
	return nil;
}

func (app * Application) CreateDatabaseHandle() {

}

func NewApplication() (*Application) {
	app := new(Application)
	return app
}



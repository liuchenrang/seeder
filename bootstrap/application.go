package bootstrap

import (
	"seeder/config"

	"github.com/liuchenrang/log4go"
)

type Application struct {
	mapObj map[string]interface{}
}

func (app *Application) Get(key string) interface{} {
	return app.mapObj[key]
}
func (app *Application) Set(key string, object interface{}) {
	app.mapObj[key] = object
}
func NewApplication() *Application {
	app := &Application{mapObj: make(map[string]interface{})}
	return app
}
func (app *Application) GetLogger() log4go.Logger {
	return app.Get("globalLogger").(log4go.Logger)
}

func (app *Application) GetConfig() config.SeederConfig {
	return app.Get("globalSeederConfig").(config.SeederConfig)
}

package bootstrap


type Application struct {
	mapObj map[string]interface{}
}
func (app *Application) Get(key string ) interface{} {
	return app.mapObj[key]
}
func (app *Application) Set(key string , object interface{})  {
	app.mapObj[key] = object
}
func NewApplication() *Application{
	app := &Application{mapObj:make(map[string]interface{})}
	return app
}
package app
 
import (
    "fmt"
    "log"
    "net/http"
 
    "goapi/src/app/handler"
    "goapi/src/app/model"
    "goapi/src/config"
    "github.com/gorilla/mux"
    "github.com/jinzhu/gorm"
)
 
// App has router and db instances
type App struct {
    Router *mux.Router
    DB     *gorm.DB
}
 
// App initialize with predefined configuration
func (a *App) Initialize(config *config.Config) {
    dbURI := fmt.Sprintf(
        "host=localhost port=5432 user=%s dbname=%s password=%s sslmode=disable",
        config.DB.Username,
        config.DB.Name,
        config.DB.Password)

    db, err := gorm.Open(config.DB.Dialect, dbURI)
    if err != nil {
        log.Fatal(err)
        log.Fatal("Could not connect database")
    }
    // db.LogMode(true)
    db.DB().SetMaxOpenConns(90)
 
    a.DB = model.DBMigrate(db)
    a.Router = mux.NewRouter()
    a.setRouters()
}
 
// Set all required routers
func (a *App) setRouters() {
    // Routing for handling the projects
    a.Get("/employees/", a.GetAllEmployees)
    a.Post("/employees/", a.CreateEmployee)
    a.Get("/employees/{name}/", a.GetEmployee)
    a.Patch("/employees/{name}/", a.UpdateEmployee)
    a.Delete("/employees/{name}/", a.DeleteEmployee)
    a.Patch("/employees/{name}/disable/", a.DisableEmployee)
    a.Patch("/employees/{name}/enable/", a.EnableEmployee)
}
 
// Wrap the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
    a.Router.HandleFunc(path, f).Methods("GET")
}
 
// Wrap the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
    a.Router.HandleFunc(path, f).Methods("POST")
}
 
// Wrap the router for PATCH method
func (a *App) Patch(path string, f func(w http.ResponseWriter, r *http.Request)) {
    a.Router.HandleFunc(path, f).Methods("PATCH")
}
 
// Wrap the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
    a.Router.HandleFunc(path, f).Methods("DELETE")
}
 
// Handlers to manage Employee Data
func (a *App) GetAllEmployees(w http.ResponseWriter, r *http.Request) {
    handler.GetAllEmployees(a.DB, w, r)
}
 
func (a *App) CreateEmployee(w http.ResponseWriter, r *http.Request) {
    handler.CreateEmployee(a.DB, w, r)
}
 
func (a *App) GetEmployee(w http.ResponseWriter, r *http.Request) {
    handler.GetEmployee(a.DB, w, r)
}
 
func (a *App) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
    handler.UpdateEmployee(a.DB, w, r)
}
 
func (a *App) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
    handler.DeleteEmployee(a.DB, w, r)
}
 
func (a *App) DisableEmployee(w http.ResponseWriter, r *http.Request) {
    handler.DisableEmployee(a.DB, w, r)
}
 
func (a *App) EnableEmployee(w http.ResponseWriter, r *http.Request) {
    handler.EnableEmployee(a.DB, w, r)
}
 
// Run the app on it's router
func (a *App) Run(host string) {
    log.Fatal(http.ListenAndServe(host, a.Router))
}

package main
 
import (
    "goapi/src/app"
    "goapi/src/config"
)
 
func main() {
    config := config.GetConfig()
 
    app := &app.App{}
    app.Initialize(config)
    app.Run(":3000")
}

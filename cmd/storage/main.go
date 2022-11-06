package storage

import (
	"fmt"
	"os"

	viper "github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigName("goodmorning")
	viper.SetConfigType("toml")
	viper.AddConfigPath("$HOME/.config/")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			viper.SetDefault("lang", "en")
			viper.SafeWriteConfig()
		} else {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func SaveApp(app App) {
	viper.Set("apps."+app.slug()+".label", app.label)
	viper.Set("apps."+app.slug()+".command", app.command)
	viper.Set("apps."+app.slug()+".active", app.active)
	viper.SafeWriteConfig()
}

func ParseApps() []App {
	apps := viper.Get("apps").(map[string]interface{})
	result := []App{}
	for _, value := range apps {
		app := value.(map[string]interface{})
		result = append(result, App{
			label:   app["label"].(string),
			command: app["command"].(string),
			active:  app["active"].(bool),
		})
	}
	return result
}

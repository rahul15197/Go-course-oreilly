package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var username, password string

func useCredentials() {
	fmt.Printf("Your username is %s\n", viper.GetString("credentials.username"))
	fmt.Printf("Your password is %s\n", viper.GetString("credentials.password"))
}
func main() {
	flag.StringVar(&username, "user", "", "username")
	flag.StringVar(&password, "pass", "", "password")
	flag.Parse()
	viper.SetDefault("webserver.endpoint", "https://httpbin.org/get")
	viper.SetDefault("webserver.name", "Testing Endpoint")
	viper.AddConfigPath(".")
	viper.SetConfigName("creds")
	viper.ReadInConfig()
	if username == "" {
		viper.Set("credentials.username", username)
	}
	if password == "" {
		viper.Set("credentials.password", password)
	}
	if !viper.IsSet("credentials.username") || !viper.IsSet("credentials.password") {
		log.Fatalln("Credentials not set")
	}

	useCredentials()
}

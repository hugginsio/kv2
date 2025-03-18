package main

import (
	"log"
	"os"
)

type Configuration struct {
	DevMode    bool
	PrivateKey string
	PublicKey  string
	TsAuthKey  string
}

// Retrieves the configuration from the environment.
func RetrieveConfiguration() Configuration {
	configuration := Configuration{
		DevMode:    os.Getenv("KV2_DEV_MODE") == "true",
		PrivateKey: os.Getenv("KV2_PRIVATE_KEY"),
		PublicKey:  os.Getenv("KV2_PUBLIC_KEY"),
		TsAuthKey:  os.Getenv("KV2_TS_AUTHKEY"),
	}

	// Go ahead and run this since we don't fetch the config anywhere else.
	preflight(configuration)

	return configuration
}

// Check for misconfigurations, print warnings, etc.
func preflight(configuration Configuration) {
	if configuration.DevMode {
		log.Println("")
		log.Println("<!> RUNNING IN DEVELOPMENT MODE         <!>")
		log.Println("<!> An in-memory database will be used. <!>")
		log.Println("<!> Tailscale will not be used.         <!>")
		log.Println("")
	}

	if !configuration.DevMode && configuration.TsAuthKey == "" {
		log.Fatalln("KV2_TS_AUTHKEY is required outside of development mode.")
	}

	if !configuration.DevMode && configuration.PrivateKey == "" {
		log.Fatalln("KV2_PRIVATE_KEY is required.")
	}

	if !configuration.DevMode && configuration.PublicKey == "" {
		log.Fatalln("KV2_PUBLIC_KEY is required.")
	}
}

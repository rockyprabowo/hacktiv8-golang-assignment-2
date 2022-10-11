package main

import (
	_ "embed"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"io"
	"log"
	"os"
	"rocky.my.id/git/h8-assignment-2/helpers/debugging"
	"rocky.my.id/git/h8-assignment-2/helpers/options"
	"rocky.my.id/git/h8-assignment-2/helpers/slices"
	"strings"
)

//go:embed ascii.txt
var logo string

type _AppState struct {
	Debug       bool
	Environment string
}

var state _AppState

func init() {
	fmt.Println(logo)
	// Loads the environment variables from .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println(
			"WARN: Couldn't load .env file. Load configurations from system environment variables.",
		)
	}

	// Read and setup environment and debug application state
	env, debug := setupEnvDebugMode()
	state = _AppState{
		Environment: env,
		Debug:       debug,
	}
	setupGeneralLogging()

	// Check mandatory environment variables used for configurations
	if err := checkEnvVars(); err != nil {
		debugging.Print(err.Error(), "checkEnvVars")
		log.Fatal("Couldn't load mandatory configuration environment variables!")
	}
	debugging.Print("Mandatory configuration environment variables are available.", "init")
}

// setupServerAddr builds the address for Gin to listen and serve the application to.
func setupServerAddr() string {
	host := options.Default(os.Getenv("APP_HOST"), "localhost")
	port := options.Default(os.Getenv("APP_PORT"), "5001")
	return fmt.Sprintf("%s:%s", host, port)
}

func parseEnvironment() string {
	return options.Default(os.Getenv("APP_ENV"), "production")
}

// isDebugActive returns true if the APP_DEBUG contains a "true-ish" value. Returns false otherwise.
func isDebugActive() bool {
	debugVar, exist := os.LookupEnv("APP_DEBUG")
	return exist && slices.Contains([]string{"true", "1", "on", "active"}, debugVar)
}

// setupEnvDebugMode reads and return environment and debug from environment variables
func setupEnvDebugMode() (env string, debug bool) {
	env = parseEnvironment()
	debug = isDebugActive()
	return
}

// setupGeneralLogging set logging output accordingly
func setupGeneralLogging() {
	if !state.Debug {
		log.SetOutput(io.Discard)
		log.SetFlags(0)
	}
	return
}

// setupGinLogging set Gin related logging accordingly
func setupGinLogging() {
	switch state.Environment {
	case "local":
		gin.SetMode(gin.DebugMode)
	case "production":
		fallthrough
	default:
		gin.SetMode(gin.ReleaseMode)
	}
}

// requiredEnvVars builds a list of mandatory environment variables
func requiredEnvVars() []string {
	var vars = []string{
		"APP_HOST",
		"APP_PORT",
		"APP_STATIC_URL",
	}

	if _, present := os.LookupEnv("DB_URL"); present {
		debugging.Print("config: using DB_URL configurations", "requiredEnvVars()")
		vars = append(vars, "DB_URL")
		return vars
	}

	debugging.Print("config: using DB_* configurations", "requiredEnvVars()")
	vars = append(
		vars,
		[]string{"DB_HOST", "DB_PORT", "DB_USER", "DB_PASSWORD", "DB_DATABASE"}...,
	)

	return vars
}

// checkEnvVars checks every mandatory environment variables used for configurations
func checkEnvVars() error {
	var undefinedEnvVars []string
	for _, v := range requiredEnvVars() {
		if _, present := os.LookupEnv(v); !present {
			undefinedEnvVars = append(undefinedEnvVars, v)
		}
	}

	if len(undefinedEnvVars) > 0 {
		return errors.New(
			"Environment variable " +
				strings.Join(undefinedEnvVars, ", ") +
				" is not set!",
		)
	}
	return nil
}

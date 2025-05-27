package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Database *gorm.DB
	Jira     JiraConfig
	App      Appconfig
}

type DatabaseConfig struct {
	Name string
	Host string
	User string
	Pass string
	Port string
}

type JiraConfig struct {
	Url           string
	User          string
	Pass          string
	SyncInterval  time.Duration // minutes
	Projects      []string
	SyncStartDate string
}

type Appconfig struct {
	Name    string
	Version string
}

func Get() *Config {

	_, err := os.Stat("./.env")
	if err == nil {
		err := godotenv.Load()
		if err != nil {
			panic(err)
		}
	}

	return &Config{
		App: Appconfig{
			Name:    getRequiredEnv("APP_NAME"),
			Version: getRequiredEnv("APP_VERSION"),
		},
		Database: createDbDns(),
		Jira: JiraConfig{
			Url:           getRequiredEnv("JIRA_URL"),
			User:          getRequiredEnv("JIRA_USER"),
			Pass:          getRequiredEnv("JIRA_PASS"),
			SyncInterval:  getJiraSyncInterval() * time.Minute,
			Projects:      getJiraProjects(),
			SyncStartDate: getRequiredEnv("JIRA_SYNC_START_DATE"),
		},
	}

}

func getRequiredEnv(env string) string {
	value, exists := os.LookupEnv(env)
	if !exists || value == "" {
		panic(fmt.Sprintf("Required environment variable %s is not set or is empty.", env))
	}
	return value
}

func createDbDns() *gorm.DB {

	time.Sleep(15 * time.Second) // wait few seconds to database to be ready!

	dbconfig := DatabaseConfig{
		Name: getRequiredEnv("DATABASE_NAME"),
		Host: getRequiredEnv("DATABASE_HOST"),
		User: getRequiredEnv("DATABASE_USER"),
		Pass: getRequiredEnv("DATABASE_PASS"),
		Port: getRequiredEnv("DATABASE_PORT"),
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbconfig.Host, dbconfig.User, dbconfig.Pass, dbconfig.Name, dbconfig.Port)

	db, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn}))
	if err != nil {
		panic(err.Error())
	}

	return db

}

func getJiraProjects() []string {
	file, err := os.ReadFile("config/jira_projects.txt")
	if err != nil {
		panic(err.Error())
	}

	return strings.Split(string(file), "\n")
}

func getJiraSyncInterval() time.Duration {
	timeInt, err := strconv.Atoi(getRequiredEnv("JIRA_SYNC_INTERVAL"))
	jiraSyncInterval := time.Duration(timeInt)
	if err != nil {
		panic("erro converting JiraSyncInterval")
	}
	return jiraSyncInterval
}

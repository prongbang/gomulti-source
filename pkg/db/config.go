package db

// Config is a model for config mariadb
type Config struct {
	User         string
	Pass         string
	Host         string
	DatabaseName string
}

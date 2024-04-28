package database

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const (
	DBHost     = "DB_HOST"
	DBPort     = "DB_PORT"
	DBUser     = "DB_USERNAME"
	DBPassword = "DB_PASSWORD"
	DBName     = "DB_NAME"
)

type PGInstance struct {
	DB *gorm.DB
}

func NewPGInstance() (*PGInstance, error) {
	config, err := getConfigDetails()
	if err != nil {
		return nil, fmt.Errorf("failed to get config details: %v", err)
	}
	db, err := bootDatabase(config)
	if err != nil {
		return nil, fmt.Errorf("failed to boot database: %v", err)
	}
	return &PGInstance{
		DB: db,
	}, nil
}

type databaseconfig struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

func generateDbConnectionString(config databaseconfig) string {
	return fmt.Sprintf(
		"host=%v port=%v user=%v password=%v dbname=%v",
		config.host,
		config.port,
		config.user,
		config.password,
		config.dbname,
	)
}

func getConfigDetails() (databaseconfig, error) {
	host := os.Getenv(DBHost)
	port := os.Getenv(DBPort)
	user := os.Getenv(DBUser)
	password := os.Getenv(DBPassword)
	dbname := os.Getenv(DBName)

	config := databaseconfig{
		host:     host,
		port:     port,
		user:     user,
		password: password,
		dbname:   dbname,
	}

	if host == "" || port == "" || user == "" || password == "" || dbname == "" {
		return config, fmt.Errorf("missing required environment variables")
	}

	return config, nil
}

func bootDatabase(config databaseconfig) (*gorm.DB, error) {
	connectionString := generateDbConnectionString(config)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  connectionString,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		PrepareStmt: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}

	dbSQL, err := db.DB()
	if err != nil {
		return nil, err
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	dbSQL.SetMaxIdleConns(50)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	dbSQL.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	dbSQL.SetConnMaxLifetime(time.Minute * 10)

	if err := dbSQL.Ping(); err != nil {
		return nil, fmt.Errorf("database ping: %v", err)
	}

	err = db.AutoMigrate()
	if err != nil {
		return nil, err
	}

	return db, nil
}

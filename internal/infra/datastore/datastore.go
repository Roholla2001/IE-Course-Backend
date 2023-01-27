package datastore

import (
	"fmt"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type dbConfig struct {
	Host     string
	Port     uint16
	Database string
	User     string
	Password string
	LogLevel uint16
}

var GormLog bool
var dbSaved *gorm.DB

func NewDBConn() (db *gorm.DB, err error) {
	fmt.Println("Connecting to Database")

	//Get database config from environmental variables
	cfg := &dbConfig{}
	if err := cfg.buildConfigFromEnv(); err != nil {
		return nil, fmt.Errorf("could not load config, %w", err)
	}

	//Choose Database Logging policy
	gormLogEnv := os.Getenv("LogGormEnv")
	if gormLogEnv != "" {
		if GormLog, err = strconv.ParseBool(gormLogEnv); err != nil {
			return nil, fmt.Errorf("couldn't parse %v env value '%s': %w",
				"LogGormEnv", gormLogEnv, err)
		}
	} else {
		GormLog = false
	}

	//Create a connection string from config
	connStr := cfg.toConnStr()

	//Get a new Connection Pool from postgres driver
	db, err = gorm.Open(postgres.Open(connStr), &gorm.Config{
		Logger: logger.Default.LogMode(logger.LogLevel(cfg.LogLevel)),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
		},
	})

	if err != nil {
		return nil, err
	}

	dbSaved = db
	return db, nil
}

//fill db config fields from envs
func (cfg *dbConfig) buildConfigFromEnv() error {
	cfg.Host = os.Getenv("LMSDbHost")
	if cfg.Host == "" {
		return fmt.Errorf("%v env must be specified: db host", "LMSDbHost")
	}

	cfg.Database = os.Getenv("LMSDbDatabase")
	if cfg.Database == "" {
		return fmt.Errorf("%v env must be specified: database", "LMSDbDatabase")
	}

	portStr := os.Getenv("LMSDbPort")
	if portStr == "" {
		return fmt.Errorf("%v env must be specified: db port", "LMSDbPort")
	}

	port, err := strconv.ParseUint(portStr, 10, 16)
	if err != nil {
		return fmt.Errorf("couldn't parse %v env value '%s': %w", "LMSDbPort", portStr, err)
	}

	cfg.Port = uint16(port)

	cfg.User = os.Getenv("LMSDbUser")
	if cfg.User == "" {
		return fmt.Errorf("%v env must be specified: db user", "LMSDbUser")
	}

	cfg.Password = os.Getenv("LMSDbPass")
	if cfg.Password == "" {
		return fmt.Errorf("%v env must be specified: db password", "LMSDbPass")
	}

	logStr := os.Getenv("LMSGormLog")
	if portStr == "" {
		return fmt.Errorf("%v env must be specified: gorm log mode", "LMSGormLog")
	}

	logMode, err := strconv.ParseUint(logStr, 10, 16)
	if err != nil {
		return fmt.Errorf("couldn't parse %v env value '%s': %w", "LMSGormLog", logStr, err)
	}

	cfg.LogLevel = uint16(logMode)

	return nil
}

//make a connection string form db config
func (cfg dbConfig) toConnStr() string {
	connString := fmt.Sprintf(
		"host=%s port=%v dbname=%s user=%s password=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.Database, cfg.User, cfg.Password)

	return connString
}

func GetDBConn() *gorm.DB {
	return dbSaved
}

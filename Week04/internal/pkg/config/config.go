package config

import "github.com/mainflux/mainflux"

const (
	defLogLevel       = "error"
	defDBHost         = "localhost"
	defDBPort         = "5432"
	defDBUser         = "test"
	defDBPass         = "test123"
	defDB             = "test_db"
	defDBDebugLevel   = "Silent"
	defDBSSLMode      = "disable"
	defDBSSLCert      = ""
	defDBSSLKey       = ""
	defDBSSLRootCert  = ""
	defPort           = "8180"


	envLogLevel       = "IOT_DP_LOG_LEVEL"
	envDBHost         = "IOT_DP_DB_HOST"
	envDBPort         = "IOT_DP_DB_PORT"
	envDBUser         = "IOT_DP_DB_USER"
	envDBPass         = "IOT_DP_DB_PASS"
	envDB             = "IOT_DP_DB"
	envDBDebugLevel   = "IOT_DP_DB_LOG_LEVEL"
	envDBSSLMode      = "IOT_DP_DB_SSL_MODE"
	envDBSSLCert      = "IOT_DP_DB_SSL_CERT"
	envDBSSLKey       = "IOT_DP_DB_SSL_KEY"
	envDBSSLRootCert  = "IOT_DP_DB_SSL_ROOT_CERT"
	envPort           = "IOT_DP_PORT"

)

type Postgres struct {
	Host        string
	Port        string
	User        string
	Pass        string
	Name        string
	SSLMode     string
	SSLCert     string
	SSLKey      string
	SSLRootCert string
	LogLevel    string
}

type Config struct {
	LogLevel string
	HttpPort string
	DbConfig Postgres
}

func NewConfig() *Config {
	dbConfig := Postgres{
		Host:        mainflux.Env(envDBHost, defDBHost),
		Port:        mainflux.Env(envDBPort, defDBPort),
		User:        mainflux.Env(envDBUser, defDBUser),
		Pass:        mainflux.Env(envDBPass, defDBPass),
		Name:        mainflux.Env(envDB, defDB),
		SSLMode:     mainflux.Env(envDBSSLMode, defDBSSLMode),
		SSLCert:     mainflux.Env(envDBSSLCert, defDBSSLCert),
		SSLKey:      mainflux.Env(envDBSSLKey, defDBSSLKey),
		SSLRootCert: mainflux.Env(envDBSSLRootCert, defDBSSLRootCert),
		LogLevel:    mainflux.Env(envDBDebugLevel, defDBDebugLevel),
	}
	return &Config{
		LogLevel: mainflux.Env(envLogLevel, defLogLevel),
		HttpPort: mainflux.Env(envPort, defPort),
		DbConfig: dbConfig,
	}
}

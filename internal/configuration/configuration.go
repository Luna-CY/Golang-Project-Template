package configuration

var Configuration struct {
	Debug  bool `mapstructure:"debug"` // Debug mode
	Logger struct {
		Level      string   `mapstructure:"level"`       // Log level, allow debug, info, warn, error, panic
		Outputs    []string `mapstructure:"outputs"`     // Log outputs, allow stdout, stderr, or file paths
		MaxSize    int      `mapstructure:"max_size"`    // max log file size, in MB
		MaxAge     int      `mapstructure:"max_age"`     // max log file age, in days
		MaxBackups int      `mapstructure:"max_backups"` // max log file backups
	} `mapstructure:"logger"` // Logger configuration
	Database struct {
		Mysql struct {
			Dsn      string `mapstructure:"dsn"` // DSN connection string
			ConnPool struct {
				Enable          bool `mapstructure:"enable"`             // enable connection pool
				MaxIdleConn     int  `mapstructure:"max_idle_conn"`      // max idle connections
				MaxOpenConn     int  `mapstructure:"max_open_conn"`      // max connections
				MaxIdleLifeTime int  `mapstructure:"max_idle_life_time"` // max idle connection life time, in minutes
			} `mapstructure:"conn_pool"` // connection pool configuration
		} `mapstructure:"mysql"`
	} `mapstructure:"database"` // database configuration
	Cache struct {
		Prefix string `mapstructure:"prefix"` // cache key prefix
	} `mapstructure:"cache"` // cache configuration
	Server struct {
		Http struct {
			Web struct {
				Listen            string   `mapstructure:"listen"`            // listen address. ip:port, default: ":8000"
				GinTrustedProxies []string `mapstructure:"trusted_proxies"`   // trusted proxies
				UnderMaintenance  bool     `mapstructure:"under_maintenance"` // enable under maintenance mode
			} `mapstructure:"web"` // web server configuration
		} `mapstructure:"http"` // http server configuration
	} `mapstructure:"server"` // server configuration
	Sentry struct {
		Enable bool   `mapstructure:"enable"` // if true, enable sentry middleware
		Dsn    string `mapstructure:"dsn"`    // sentry server dsn string
	} `mapstructure:"sentry"` // sentry configuration
}

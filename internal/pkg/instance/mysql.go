package instance

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MysqlConfig struct {
	Host                  string        `mapstructure:"host"`
	User                  string        `mapstructure:"user"`
	Password              string        `mapstructure:"password"`
	Database              string        `mapstructure:"database"`
	MaxIdleConnections    int           `mapstructure:"max-idle-connections"`
	MaxOpenConnections    int           `mapstructure:"max-open-connections"`
	MaxConnectionLifeTime time.Duration `mapstructure:"max-connection-life-time"`
	LogLevel              int           `mapstructure:"log-level"`
}

// NewMySQL create a new gorm store instance with the given options.
func NewMySQL(opts *MysqlConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf(`%s:%s@tcp(%s)/?charset=utf8&parseTime=%t&loc=%s`,
		opts.User,
		opts.Password,
		opts.Host,
		//opts.Database,
		true,
		"Local",
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 打开翻译方言以进行错误处理
		TranslateError: true,
		Logger:         logger.Default.LogMode(logger.LogLevel(opts.LogLevel)),
	})

	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(opts.MaxOpenConnections)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(opts.MaxConnectionLifeTime)

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(opts.MaxIdleConnections)

	return db, nil
}

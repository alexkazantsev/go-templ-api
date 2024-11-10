package config

import (
	"fmt"
	"strconv"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type DatabaseConfig struct {
	Host           string        `env:"DB_HOST"`
	Port           string        `env:"DB_PORT"`
	User           string        `env:"DB_USER"`
	Password       string        `env:"DB_PASSWORD"`
	Name           string        `env:"DB_NAME"`
	DBSSLDisable   string        `env:"DB_SSL_DISABLE" envDefault:"true"`
	DBMaxIdleConns string        `env:"DB_MAX_IDLE_CONNS" envDefault:"25"`
	DBMaxOpenConns string        `env:"DB_MAX_OPEN_CONNS" envDefault:"25"`
	DBConnMaxLife  time.Duration `env:"DB_CONN_MAX_LIFE" envDefault:"1h"`
}

func (c DatabaseConfig) GetDSN() string {
	var sslDisabledStr string
	if disable, _ := strconv.ParseBool(c.DBSSLDisable); disable {
		sslDisabledStr = "sslmode=disable"
	}

	return fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v %s",
		c.Host, c.User, c.Password, c.Name, c.Port, sslDisabledStr,
	)
}

func (c DatabaseConfig) Validate() error {
	err := validation.ValidateStruct(&c,
		validation.Field(&c.Host, validation.Required),
		validation.Field(&c.Port, validation.Required, is.Port),
		validation.Field(&c.User, validation.Required),
		validation.Field(&c.Password, validation.Required),
		validation.Field(&c.Name, validation.Required),
	)

	return err
}

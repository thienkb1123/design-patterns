/*
	#SRP
	Context: You need to create a Config struct to store all configuration information of the application (database, redis, logger, etc).
	Requirement:
		- Allow only a single instance to be created
		- Don't allow new an instance from outside the package
		- Must be thread-safe
*/

package singleton

import "sync"

var (
	instance *Config
	once     sync.Once
)

type Config struct {
	Postgres
	Redis
	Logger
}

type Postgres struct {
	URI     string
	NumPool int
}

type Redis struct {
	Host string
	Port int
}

type Logger struct {
	Level string
}

func NewConfig() *Config {
	once.Do(func() {
		instance = &Config{
			Postgres: Postgres{
				URI:     "postgresql://xxx:xxx@localhost:5432/database",
				NumPool: 1,
			},
			Redis: Redis{
				Host: "localhost",
				Port: 6379,
			},
			Logger: Logger{
				Level: "debug",
			},
		}
	})
	return instance
}

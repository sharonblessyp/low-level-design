## Logger Framework:
 simple structure -> complex system underneath

- frameworks doesnt have
controller -> service -> repository

- LoggerConfig stores configuration, while Logger stores behavior and runtime state.

## LoggerConfig = Settings
- How should the logger behave?
Examples:

Log level = INFO
Output = console
Output = file
Format = JSON

These are just settings.

## Logger = The actual service/object (contains runtime state)

Who performs the logging?
It contains:
- Configuration (config)
- Synchronization (mu)
- Methods like Info, Error, SetLevel


------

More examples
1. Database connection pool
type DBConfig struct {
	MaxConnections int
	Host           string
	Port           int
}

Configuration:

Host = localhost
Port = 5432
MaxConnections = 100

Runtime state:

type DB struct {
	config     DBConfig
	openConns  int
	idleConns  []*Connection
	mu         sync.Mutex
}

During execution:

openConns = 57
idleConns = 12

These values keep changing.
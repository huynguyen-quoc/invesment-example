package config


type Mode string

// Execution modes
const (
	Production  Mode = "production"
	Staging     Mode = "staging"
	Development Mode = "development"
	Test        Mode = "test"
)
// StreamConfig struct  to store the config for stream
type StreamConfig struct {
}

// BaseConfig struct to keep config for common store i.e. Redis, Postgres
type BaseConfig struct {
}

// AppConfig is your application's config
type AppConfig struct {
	Mode      Mode   `json:"mode"`      // to set the current place this instance is running in
	ServiceID string `json:"serviceID"` // used internally while creating UCM client for setting up the configHandler

	// data config
	Data BaseConfig `json:"data"` // config for all common store i.e. Redis, Mysql and Kinesis.

	//stream config
	StreamConfig *StreamConfig `json:"streamConfig"`
}

package config

type DatabaseConfig struct {
	BlocksDatabasePort int
	BlocksDatabaseUri string
	BlocksDatabaseType string
	BlocksDatabaseFile string

	OutputsDatabasePort int
	OutputsDatabaseUri string
	OutputsDatabaseType string
	OutputsDatabaseFile string
}

var DbConf = new(DatabaseConfig)

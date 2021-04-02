package types

type PostgresInstance struct {
	PgConfig pgc `mapstructure:"pg"`
}

type pgc struct {
	Host     string `mapstructure:"host"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

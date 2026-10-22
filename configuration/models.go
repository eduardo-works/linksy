package configuration

type Database struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

type HTTP struct {
	Port int
}

type Config struct {
	Database Database
	HTTP     HTTP
}

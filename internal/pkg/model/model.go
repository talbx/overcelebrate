package model

type BirthdayEntry struct {
	Name string `yaml:"name"`
	Date string `yaml:"date"`
}

type AppConfig struct {
	Pushover struct {
		Apitoken  string
		Usertoken string
	}
}

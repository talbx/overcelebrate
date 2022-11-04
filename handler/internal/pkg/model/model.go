package model

type BirthdayEntry struct {
	Name string `yaml:"name"`
	Date string `yaml:"date"`
}
type PushoverConfig struct {
	Apitoken  string `yaml:"apitoken"`
	Usertoken string `yaml:"usertoken"`
}
type AppConfig struct {
	Birthdays      []BirthdayEntry `yaml:"bdays"`
	PushoverConfig PushoverConfig  `yaml:"pushover"`
}

package config

const LogFile = "LogFile"
const DBConnString = "DBConnString"
const Port = "Port"
const SigningKey = "SigningKey"

var defaults = map[string]string{
	LogFile:      "logs.txt",
	DBConnString: "",
	Port:         "8000",
	SigningKey: "this-is-a-fake-signing-key",
}

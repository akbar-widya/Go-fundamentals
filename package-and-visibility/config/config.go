package config

var Version = "1.0.0"

var secretKey = "super-secret-key"

func GetSystemInfo() string {
	return "System v" + Version + " using key: " + loadSecret()
}

func loadSecret() string {
	return secretKey
}

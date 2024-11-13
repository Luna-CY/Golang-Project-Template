package runtime

import "os"

func GetEnvironment() string {
	env := os.Getenv("ENV")

	switch env {
	case "prod", "production":
		return "production"
	case "sandbox":
		return "sandbox"
	case "test":
		return "test"
	case "dev", "develop":
		return "develop"
	default:
		return "develop"
	}
}

func IsDevelopment() bool {
	return GetEnvironment() == "develop"
}

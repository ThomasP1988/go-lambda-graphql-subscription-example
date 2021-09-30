package config

import "os"

type Stage string

var (
	DEV     Stage = "dev"
	STAGING Stage = "staging"
	PROD    Stage = "prod"
)

type Config struct {
	LocalProfile            string
	Region                  string
	Tables                  TableRegistry
	DataScienceArtEyeAPI    string
	DataScienceArtEyeAPIKey string
}

func SetEnv(env *Stage) *Stage {
	if env == nil {
		envS := Stage(os.Getenv("stage"))
		if envS != "" {
			return &envS

		} else {
			return &DEV

		}
	}
	return env
}

var Conf *Config

func GetConfig(env *Stage) *Config {

	if Conf != nil {
		return Conf
	}
	env = SetEnv(env)

	Conf = &Config{
		// LocalProfile: "yourProfile",
		Region: "eu-west-3",
		Tables: GetTables(env),
		// DataScienceArtEyeAPIAPI:    "https://arteye-api-dev.ds.jha.com", // real dev
		// DataScienceArtEyeAPIAPIKey: "4be82df5-542f-48c2-9f35-a72daa77ada5", // real dev
		DataScienceArtEyeAPI:    "https://arteye-api.ds.jha.com",        // prod that we use on dev, stage, prod
		DataScienceArtEyeAPIKey: "899e5bce-813c-4d20-b334-6a13200e1e0c", // prod that we use on dev, stage, prod
	}

	if *env == STAGING {
		updateConfigForStage(Conf)
	} else if *env == PROD {
		updateConfigForProd(Conf)
	}

	return Conf
}

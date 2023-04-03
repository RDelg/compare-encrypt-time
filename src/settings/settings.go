package setting

import (
	"log"

	env "github.com/Netflix/go-env"
)

type App struct {
	ServerSetting struct {
		RunMode  string `env:"APP_MODE,default=release"`
		HttpPort int    `env:"PORT,default=8000"`
	}
	CypherSetting struct {
		CypherKey string `env:"CYPHER_KEY,required=true"`
		CypherIv  string `env:"CYPHER_IV,required=true"`
	}
}

var AppSetting = &App{}

// Setup initialize the configuration instance
func Setup() {
	_, err := env.UnmarshalFromEnviron(AppSetting)
	if err != nil {
		log.Fatalf("[error] env.UnmarshalFromEnviron: %v", err)
	}

}

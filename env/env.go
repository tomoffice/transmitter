package env

import (
	"log"
	"os"
	"strconv"
	PkgSql "transmitter/sql"
)

type envConfig struct {
	config PkgSql.Config
}

// 如果在環境變數沒有找到將會使用envMap取代
func (env *envConfig) PullEnv(envMap map[string]interface{}) {
	for index := range envMap {
		le, ok := os.LookupEnv(index)
		if ok {
			switch envMap[index].(type) {
			case string:
				envMap[index] = le
			case int:
				v, _ := strconv.Atoi(le)
				envMap[index] = v
			case bool:
				v, _ := strconv.ParseBool(le)
				envMap[index] = v
			}
		} else {
			log.Println("using default env:", index)
		}
	}

	env.config = PkgSql.Config{
		Username:  envMap["dbusername"].(string),
		Password:  envMap["dbpassword"].(string),
		Protocol:  envMap["protocol"].(string),
		Host:      envMap["host"].(string),
		Port:      envMap["port"].(int),
		Database:  envMap["database"].(string),
		Chartset:  envMap["charset"].(string),
		Location:  envMap["location"].(string),
		Parsetime: envMap["parsetime"].(bool),
	}

}

// 原始設計給Docker
func NewEnvConfig(envMap map[string]interface{}) PkgSql.Config {
	env := &envConfig{}
	env.PullEnv(envMap)
	return env.config
}

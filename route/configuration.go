package route

import (
	"github.com/Unknwon/goconfig"
	"path/filepath"
	"github.com/godaner/go-util/fileutil"
	"log"
	"go-util/errorutil"
)


type Configuration struct {
	Port string
}

const(
	CONFIGURATION_FILE_NAME = "app.conf"
)
var (
	conf Configuration
)
func Default()Configuration{
	return Configuration{
		Port:"80",
	}
}

func MakeConfig(cfg *goconfig.ConfigFile) Configuration{
	port, err := cfg.GetValue(goconfig.DEFAULT_SECTION, "port")
	errorutil.CheckErr(err)

	return Configuration{
		Port:port,
	}
}
func LoadConfiguration(){
	workPath := fileutil.Getwd()
	log.Println("LoadConfiguration workPath is : ",workPath)
	appPath := fileutil.GetAppPath()
	log.Println("LoadConfiguration appPath is : ",appPath)
	configPath := filepath.Join(workPath, "conf", CONFIGURATION_FILE_NAME)
	if !fileutil.FileExists(configPath) {
		configPath = filepath.Join(appPath, "conf", CONFIGURATION_FILE_NAME)
		if !fileutil.FileExists(configPath) {
			conf = Default()
			log.Println("LoadConfiguration conf is default , conf is : ",conf)
			return
		}
	}

	cfg, err := goconfig.LoadConfigFile(configPath)
	errorutil.CheckErr(err)

	conf=MakeConfig(cfg)
	log.Println("LoadConfiguration conf is custom , conf is : ",conf)

}



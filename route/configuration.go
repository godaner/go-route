package route

import (
	"github.com/Unknwon/goconfig"
	"os/exec"
	"os"
	"strings"
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

func LoadConfiguration(){
	path := getCurrentPath()
	cfg, err := goconfig.LoadConfigFile(path+"/"+CONFIGURATION_FILE_NAME)
	checkErr(err)

	port, err := cfg.GetValue(goconfig.DEFAULT_SECTION, "port")
	checkErr(err)

	conf=Configuration{
		Port:port,
	}
}



func getCurrentPath() string {
	s, err := exec.LookPath(os.Args[0])
	checkErr(err)
	i := strings.LastIndex(s, "\\")
	path := string(s[0 : i+1])
	return path
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
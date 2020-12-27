package config
import(
	"log"

	"github.com/spf13/viper"
)

type Server struct{
	Port int
}



func ServerConfig() *Server{
	configFile := "config.yaml"
	viper.SetConfigFile(configFile)
	//Read the content of the config file
	if err := viper.ReadInConfig(); err != nil{
		//Report error
		log.Fatalf("Error while reading the %s file: %s", configFile, err)
	}
	port, ok := viper.Get("server.port").(int);
	if	!ok{
		log.Fatal("Error reading the port value")
	}

	return &Server{
		//Set port
		Port:port,
	}
}

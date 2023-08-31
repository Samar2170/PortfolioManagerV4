package internal

import (
	"log"

	"github.com/samar2170/portfolio-manager-v4/client/cognitio/cauth"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var passwordDecryptionKey string
var cognitioClient = new(cauth.AuthServiceClient)

const TradeDateFormat = "2006-01-02"

func init() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	passwordDecryptionKey = viper.GetString("PASSWORD_ENCRYPTION_KEY")
	grpcHost, grpcPort := viper.GetString("GRPC_HOST"), viper.GetString("GRPC_PORT")

	uri := grpcHost + ":" + grpcPort
	log.Println(uri)
	conn, err := grpc.Dial(uri, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(err.Error())
		log.Fatalf("did not connect: %s", err)
	}

	*cognitioClient = cauth.NewAuthServiceClient(conn)
}

package internal

import (
	"log"

	"github.com/samar2170/portfolio-manager-v4/client/cognitio/cauth"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

var passwordDecryptionKey string
var cognitioClient = new(cauth.AuthServiceClient)

func init() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	passwordDecryptionKey = viper.GetString("PASSWORD_ENCRYPTION_KEY")
	grpcHost, grpcPort := viper.GetString("GRPC_HOST"), viper.GetString("GRPC_PORT")

	conn, err := grpc.Dial(grpcHost+":"+grpcPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	*cognitioClient = cauth.NewAuthServiceClient(conn)

}

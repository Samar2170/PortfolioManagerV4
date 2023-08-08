package internal

import (
	"log"

	"github.com/samar2170/portfolio-manager-v4/client/cognitio/cauth"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var passwordDecryptionKey string
var cognitioClient *cauth.AuthServiceClient

func init() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	passwordDecryptionKey = viper.GetString("passwordDecryptionKey")

	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	// cognitioClient = &cauth.NewAuthServiceClient(conn)

}

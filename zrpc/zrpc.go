package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	__ "go-zero-mysql-aula/common/pb"
	"go-zero-mysql-aula/zrpc/internal/config"
	"go-zero-mysql-aula/zrpc/internal/server"
	"go-zero-mysql-aula/zrpc/internal/svc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/zrpc.yaml", "the config file")

type Secrets struct {
	DataSourceName string `json:"UrlBancoMysql"`
	RedisHost      string `json:"RedisHost"`
	RedisPass      string `json:"EtcdHost"`
}

func getSecret() (*Secrets, error) {
	secretName := "go-zero-mysql"
	region := "us-east-1"

	awsconfig2, err := awsconfig.LoadDefaultConfig(context.TODO(), awsconfig.WithRegion(region))
	if err != nil {
		return nil, err
	}

	// Create Secrets Manager client
	svc := secretsmanager.NewFromConfig(awsconfig2)

	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"), // VersionStage defaults to AWSCURRENT if unspecified
	}

	result, err := svc.GetSecretValue(context.TODO(), input)
	if err != nil {
		// For a list of exceptions thrown, see
		// https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_GetSecretValue.html
		return nil, err
	}

	// Decrypts secret using the associated KMS key.
	var secrets Secrets
	err = json.Unmarshal([]byte(*result.SecretString), &secrets)
	if err != nil {
		return nil, err
	}

	return &secrets, nil
}

func main() {
	secret, err := getSecret()
	if err != nil {
		log.Fatalf("Unable to load secrets: %v", err)
	}

	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	// Update config with secrets
	c.DataSourceName = secret.DataSourceName
	if len(c.Cache) > 0 {
		c.Cache[0].Host = secret.RedisHost
		c.Cache[0].Pass = secret.RedisPass
	}

	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		__.RegisterZrpcServiceServer(grpcServer, server.NewZrpcServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

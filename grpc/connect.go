package grpc

import (
	"fmt"
	"github.com/guapo-organizations/go-micro-secret/consul"
	"github.com/guapo-organizations/go-micro-secret/tls"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

//链接服务
func GetGrpcConnet(servier_name string, tag string, config *api.Config, ) (*grpc.ClientConn, error) {
	agent_service, err := consul.FindService(config, servier_name, tag)

	if err != nil {
		return nil, err
	}

	//tls 和 ssl加密
	creds, err := credentials.NewClientTLSFromFile(tls.Path("ca.pem"), "zldz.com")
	if err != nil {
		return nil, err
	}

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", agent_service.Address, agent_service.Port), grpc.WithTransportCredentials(creds))

	if err != nil {
		return nil, err
	}

	return conn, nil
}

package grpc

import (
	"fmt"
	"github.com/guapo-organizations/go-micro-secret/consul"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

//连接服务
func GetGrpcConnet(servier_name string, tag string, config *api.Config, creds credentials.TransportCredentials) (*grpc.ClientConn, error) {
	agent_service, err := consul.FindService(config, servier_name, tag)

	if err != nil {
		return nil, err
	}

	var conn *grpc.ClientConn

	if creds != nil {
		//tls 和 ssl加密
		conn, err = grpc.Dial(fmt.Sprintf("%s:%d", agent_service.Address, agent_service.Port), grpc.WithTransportCredentials(creds))
	} else {
		conn, err = grpc.Dial(fmt.Sprintf("%s:%d", agent_service.Address, agent_service.Port))
	}

	if err != nil {
		return nil, err
	}

	return conn, nil
}

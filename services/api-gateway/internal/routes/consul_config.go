package routes

import (
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
)

var consulClient *consulapi.Client

const (
	AuthServiceId = "ecom-auth"
)

func NewConsulClient() *consulapi.Client {
	config := consulapi.DefaultConfig()
	client, err := consulapi.NewClient(config)
	if err != nil {
		fmt.Println(err)
	}

	consulClient = client

	return client
}

func GetServiceUrl(serviceId string) string {
	services, err := consulClient.Agent().Services()
	if err != nil {
		fmt.Println(err)
	}
	service := services[serviceId]
	return fmt.Sprintf("http://%s:%d/", service.Address, service.Port)
}

package main

import (
	"fmt"
	osb "github.com/pmorie/go-open-service-broker-client/v2"
)

func GetBrokerCatalog(URL string) (*osb.CatalogResponse, error) {
	config := osb.DefaultClientConfiguration()
	config.URL = URL

	client, err := osb.NewClient(config)
	if err != nil {
		return nil, err
	}

	return client.GetCatalog()
}

func main() {
	url := ""
	catalog, err := GetBrokerCatalog(url)
	if err != nil {
		panic(err)
	}

	for _, service := range catalog.Services {
		fmt.Println(service)
	}

}

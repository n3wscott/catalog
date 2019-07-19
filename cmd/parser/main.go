package main

import (
	"encoding/json"
	"fmt"
	osb "github.com/pmorie/go-open-service-broker-client/v2"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/openresourcebroker/catalog/pkg/apis/catalog"
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
		fmt.Printf("%s - %s\n", service.Name, service.ID)
		fmt.Printf("%s\n", service.Description)
		if service.PlanUpdatable != nil {
			fmt.Printf("plan updatable - %t\n", *service.PlanUpdatable)
		}
		for _, plan := range service.Plans {
			fmt.Printf("\t%s - %s\n", plan.Name, plan.ID)
			if plan.Schemas != nil && plan.Schemas.ServiceInstance != nil && plan.Schemas.ServiceInstance.Create != nil {
				fmt.Printf("Schema - %+v\n", plan.Schemas.ServiceInstance.Create)
			}
			fmt.Printf("------------\n")
		}
		fmt.Printf("============\n")
	}

	c := crd()

	b, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}

func crd() *apiextensions.CustomResourceDefinition {
	// kind := "AName"
	name := "a-name" + ".catalog.orb.com"

	kind := schema.GroupVersionKind{
		Kind:    "AName",
		Version: "v1alpha1",
		Group:   catalog.GroupName,
	}
	plural, singular := meta.UnsafeGuessKindToResource(kind)

	return &apiextensions.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
			Labels: map[string]string{
				"catalog.orb.dev/release": "devel",
			},
		},
		Spec: apiextensions.CustomResourceDefinitionSpec{
			Group:   kind.Group,
			Version: kind.Version,
			Names: apiextensions.CustomResourceDefinitionNames{
				Kind:     kind.Kind,
				Plural:   plural.Resource,
				Singular: singular.Resource,
				Categories: []string{
					"orb",
				},
			},
			Subresources: &apiextensions.CustomResourceSubresources{
				Status: &apiextensions.CustomResourceSubresourceStatus{},
			},
		},
	}
}

/*


apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: {name}.catalog.orb.com
  labels:
    catalog.orb.com/release: devel
spec:
  group: networking.internal.knative.dev
  version: v1alpha1
  names:
    kind: ServerlessService
    plural: serverlessservices
    singular: serverlessservice
    categories:
    - orb
  scope: Namespaced
  subresources:
    status: {}


*/

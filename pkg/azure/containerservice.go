package azure

import (
	"context"
	"os"

	"github.com/Azure/azure-sdk-for-go/services/containerservice/mgmt/2019-06-01/containerservice"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2019-05-01/resources"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

var (
	authorizer, _ = auth.NewAuthorizerFromEnvironment()
	ctx = context.Background()
)

func CreateCluster() {
	subscriptionId := os.Getenv("AZURE_SUBSCRIPTION_ID")
	location := "westus2"

	groupsClient := resources.NewGroupsClient(subscriptionId)
	groupsClient.Authorizer = authorizer
	groupsClient.CreateOrUpdate(
		ctx,
		"aks-rg",
		resources.Group{
			Location: &location,
		},
	)

	containerServiceClient := containerservice.NewContainerServicesClient(subscriptionId)
	containerServiceClient.Authorizer = authorizer
	containerServiceClient.CreateOrUpdate(
		ctx, 
		"aks-rg", 
		"aks-pilot",
		containerservice.ContainerService{
			Properties: &containerservice.Properties{

			},
		}, 
	)
}

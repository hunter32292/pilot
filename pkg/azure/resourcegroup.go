package azure

import (
	"log"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2019-05-01/resources"
)

// CreateResourceGroup creates an Azure resource group
func CreateResourceGroup(location string) {
	groupsClient := resources.NewGroupsClient(subscriptionID)
	groupsClient.Authorizer = authorizer
	group, err := groupsClient.CreateOrUpdate(
		ctx,
		"aks-rg",
		resources.Group{
			Location: &location,
		},
	)
	if err != nil {
		log.Fatalf("Unable to create resource group: %v", err)
	}
	log.Printf("Created resource group: %v", group)
}

package azure

import (
	"log"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2019-05-01/resources"
)

var groupsClient resources.GroupsClient

func init() {
	groupsClient = resources.NewGroupsClient(subscriptionID)
	groupsClient.Authorizer = authorizer
}

// CreateResourceGroup creates an Azure resource group
func CreateResourceGroup(location, name string) {
	group, err := groupsClient.CreateOrUpdate(
		ctx,
		name,
		resources.Group{
			Location: &location,
		},
	)
	if err != nil {
		log.Fatalf("Unable to create resource group: %v", err)
	}
	log.Printf("Created resource group: %v", group)
}

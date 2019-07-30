package azure

import (
	"context"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/services/containerservice/mgmt/2019-06-01/containerservice"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

var (
	ctx               = context.Background()
	subscriptionID    = os.Getenv("AZURE_SUBSCRIPTION_ID")
	clientID          = os.Getenv("AZURE_CLIENT_ID")
	clientSecret      = os.Getenv("AZURE_CLIENT_SECRET")
	authorizer, _     = auth.NewAuthorizerFromEnvironment()
	osDiskSize        = int32(0)
)

// CreateCluster creates a managed cluster
func CreateCluster(name string, location string, vmSize string, agentCount int, kubernetesVersion string, dnsPrefix string, agentPoolName string) {
	log.Print("Creating Azure cluster")

	agentCount32 := int32(agentCount)
	var agentPoolProfile = containerservice.ManagedClusterAgentPoolProfile{
		Name:         &agentPoolName,
		Count:        &agentCount32,
		VMSize:       containerservice.VMSizeTypesStandardB2s,
		OsDiskSizeGB: &osDiskSize,
	}
	agentProfileSlice := []containerservice.ManagedClusterAgentPoolProfile{agentPoolProfile}

	containerServiceClient := containerservice.NewManagedClustersClient(subscriptionID)
	containerServiceClient.Authorizer = authorizer

	cluster, err := containerServiceClient.CreateOrUpdate(ctx, "pilot", name, containerservice.ManagedCluster{
		Location: &location,
		ManagedClusterProperties: &containerservice.ManagedClusterProperties{
			KubernetesVersion: &kubernetesVersion,
			DNSPrefix:         &dnsPrefix,
			AgentPoolProfiles: &agentProfileSlice,
			ServicePrincipalProfile: &containerservice.ManagedClusterServicePrincipalProfile{
				ClientID: &clientID,
				Secret:   &clientSecret,
			},
		},
	})
	if err != nil {
		log.Fatalf("Error creating Azure cluster: %v", err)
	}
	log.Printf("Created Azure cluster: %v", cluster)
}

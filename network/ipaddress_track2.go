package network

import (
	"context"
	"log"
	"time"

	"github.com/Azure-Samples/azure-sdk-for-go-samples/internal/config"
	"github.com/Azure/azure-sdk-for-go/sdk/arm/network/2020-07-01/armnetwork"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/to"
)

func getIPClientTrack2() armnetwork.PublicIPAddressesClient {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	client := armnetwork.NewPublicIPAddressesClient(armcore.NewDefaultConnection(cred, nil), config.SubscriptionID())
	return *client
}

func CreatePublicIPTrack2(ctx context.Context, ipName string) {
	ipClient := getIPClientTrack2()
	poller, err := ipClient.BeginCreateOrUpdate(
		ctx,
		config.BaseGroupName(),
		ipName,
		armnetwork.PublicIPAddress{
			Resource: armnetwork.Resource{
				Name:     to.StringPtr(ipName),
				Location: to.StringPtr(config.DefaultLocation()),
			},
			Properties: &armnetwork.PublicIPAddressPropertiesFormat{
				PublicIPAddressVersion:   armnetwork.IPVersionIPv4.ToPtr(),
				PublicIPAllocationMethod: armnetwork.IPAllocationMethodStatic.ToPtr(),
			},
		},
		nil,
	)
	if err != nil {
		log.Fatalf("failed to obtain a response: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to create resource: %v", err)
	}
}

func GetPublicIPTrack2(ctx context.Context, ipName string) {
	ipClient := getIPClientTrack2()
	resp, err := ipClient.Get(ctx, config.BaseGroupName(), ipName, nil)
	if err != nil {
		log.Fatalf("failed to get resource: %v", err)
	}
	log.Printf("public IP address ID: %v", *resp.PublicIPAddress.ID)
}

func DeletePublicIPTrack2(ctx context.Context, ipName string) {
	ipClient := getIPClientTrack2()
	resp, err := ipClient.BeginDelete(ctx, config.BaseGroupName(), ipName, nil)
	if err != nil {
		log.Fatalf("failed to obtain a response: %v", err)
	}
	_, err = resp.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to delete resource: %v", err)
	}
}

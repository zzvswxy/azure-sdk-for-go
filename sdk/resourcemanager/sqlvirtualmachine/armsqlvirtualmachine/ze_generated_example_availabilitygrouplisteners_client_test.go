//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsqlvirtualmachine_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sqlvirtualmachine/armsqlvirtualmachine"
)

// x-ms-original-file: specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2017-03-01-preview/examples/GetAvailabilityGroupListener.json
func ExampleAvailabilityGroupListenersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsqlvirtualmachine.NewAvailabilityGroupListenersClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<sql-virtual-machine-group-name>",
		"<availability-group-listener-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AvailabilityGroupListenersClientGetResult)
}

// x-ms-original-file: specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2017-03-01-preview/examples/CreateOrUpdateAvailabilityGroupListener.json
func ExampleAvailabilityGroupListenersClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsqlvirtualmachine.NewAvailabilityGroupListenersClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<sql-virtual-machine-group-name>",
		"<availability-group-listener-name>",
		armsqlvirtualmachine.AvailabilityGroupListener{
			Properties: &armsqlvirtualmachine.AvailabilityGroupListenerProperties{
				AvailabilityGroupName: to.StringPtr("<availability-group-name>"),
				LoadBalancerConfigurations: []*armsqlvirtualmachine.LoadBalancerConfiguration{
					{
						LoadBalancerResourceID: to.StringPtr("<load-balancer-resource-id>"),
						PrivateIPAddress: &armsqlvirtualmachine.PrivateIPAddress{
							IPAddress:        to.StringPtr("<ipaddress>"),
							SubnetResourceID: to.StringPtr("<subnet-resource-id>"),
						},
						ProbePort: to.Int32Ptr(59983),
						SQLVirtualMachineInstances: []*string{
							to.StringPtr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/testvm2"),
							to.StringPtr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/testvm3")},
					}},
				Port: to.Int32Ptr(1433),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AvailabilityGroupListenersClientCreateOrUpdateResult)
}

// x-ms-original-file: specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2017-03-01-preview/examples/DeleteAvailabilityGroupListener.json
func ExampleAvailabilityGroupListenersClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsqlvirtualmachine.NewAvailabilityGroupListenersClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<sql-virtual-machine-group-name>",
		"<availability-group-listener-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2017-03-01-preview/examples/ListByGroupAvailabilityGroupListener.json
func ExampleAvailabilityGroupListenersClient_ListByGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsqlvirtualmachine.NewAvailabilityGroupListenersClient("<subscription-id>", cred, nil)
	pager := client.ListByGroup("<resource-group-name>",
		"<sql-virtual-machine-group-name>",
		nil)
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}
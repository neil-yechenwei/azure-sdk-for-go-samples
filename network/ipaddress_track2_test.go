package network

import (
	"context"
	"flag"
	"log"
	"os"
	"testing"
	"time"

	"github.com/Azure-Samples/azure-sdk-for-go-samples/internal/config"
	"github.com/Azure-Samples/azure-sdk-for-go-samples/resources"
)

var (
	pipName = "sample-pip-test01"
)

func TestMain(m *testing.M) {
	err := setupEnvironment()
	if err != nil {
		log.Fatalf("could not set up environment: %v\n", err)
	}

	os.Exit(m.Run())
}

func setupEnvironment() error {
	err1 := config.ParseEnvironment()
	err2 := config.AddFlags()

	for _, err := range []error{err1, err2} {
		if err != nil {
			return err
		}
	}

	flag.Parse()
	return nil
}

func TestNetwork(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()
	defer resources.DeleteGroup(ctx, config.BaseGroupName())

	_, err := resources.CreateGroup(ctx, config.BaseGroupName())
	if err != nil {
		t.Fatalf("failed to create group: %+v", err)
	}
	t.Logf("created group %s\n", config.BaseGroupName())

	CreatePublicIPTrack2(ctx, pipName)
	t.Logf("created public ip")

	GetPublicIPTrack2(ctx, pipName)
	t.Logf("Get public ip")

	DeletePublicIPTrack2(ctx, pipName)
	t.Logf("Delete public ip")
}

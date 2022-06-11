package lens_test

import (
	"context"
	"github.com/naturalselectionlabs/pregod/common/lens"
	"testing"
)

var (
	client *lens.Client

	address = "0x3a5bd1e37b099ae3386d13947b6a90d97675e5e3"
	profile = "0x0d"
)

func init() {
	client = lens.NewClient()
}

func TestGetProfiles(t *testing.T) {

	result, err := client.GetProfiles(context.Background(), address)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(result)
}

func TestGetPublications(t *testing.T) {
	result, err := client.GetPublications(context.Background(), profile)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(result)
}

func TestGetPublicationsByAddress(t *testing.T) {
	result, err := client.GetPublicationsByAddress(context.Background(), address)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(result)
}

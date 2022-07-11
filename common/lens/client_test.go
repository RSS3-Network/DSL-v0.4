package lens_test

import (
	"context"
	"github.com/hasura/go-graphql-client"
	"github.com/naturalselectionlabs/pregod/common/lens"
	"testing"
)

var (
	client *lens.Client

	profile = "0x05"
	address = "0x3a5bd1e37b099ae3386d13947b6a90d97675e5e3"

	options = lens.Options{
		Profile: graphql.String(profile),
		Address: graphql.String(address),
		Cursor:  "{}",
	}
)

func init() {
	client = lens.NewClient()
}

func TestGetProfiles(t *testing.T) {

	result, err := client.GetProfiles(context.Background(), &options)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(result)
}

func TestGetPublications(t *testing.T) {

	result, err := client.GetPublications(context.Background(), &options)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(result)
}

func TestGetPublicationPageInfo(t *testing.T) {
	err := client.GetPublicationPageInfo(context.Background(), &options)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(options)
}

func TestGetAllPublicationsByAddress(t *testing.T) {
	result, err := client.GetAllPublicationsByAddress(context.Background(), &options)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(result)
}
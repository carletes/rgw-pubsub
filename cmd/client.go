package cmd

import (
	rgwpubsub "github.com/ceph/rgw-pubsub-api/go/pkg"
	"github.com/golang/glog"
)

func Client() *rgwpubsub.RGWClient {
	client, err := rgwpubsub.NewRGWClient(Username, AccessKey, SecretKey, PubSubEndpoint, Zonegroup)
	if err != nil {
		glog.Fatalf("failed to create RGW PubSub client: %v", err)
	}

	return client
}

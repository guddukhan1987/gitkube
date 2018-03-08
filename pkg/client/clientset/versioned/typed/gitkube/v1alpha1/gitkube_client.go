/*
Copyright 2018 Hasura.io

*/

package v1alpha1

import (
	v1alpha1 "github.com/hasura/gitkube/pkg/apis/gitkube.sh/v1alpha1"
	"github.com/hasura/gitkube/pkg/client/clientset/versioned/scheme"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type GitkubeV1alpha1Interface interface {
	RESTClient() rest.Interface
	RemotesGetter
}

// GitkubeV1alpha1Client is used to interact with features provided by the gitkube.sh group.
type GitkubeV1alpha1Client struct {
	restClient rest.Interface
}

func (c *GitkubeV1alpha1Client) Remotes(namespace string) RemoteInterface {
	return newRemotes(c, namespace)
}

// NewForConfig creates a new GitkubeV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*GitkubeV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &GitkubeV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new GitkubeV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *GitkubeV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new GitkubeV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *GitkubeV1alpha1Client {
	return &GitkubeV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *GitkubeV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}

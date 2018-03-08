/*
Copyright 2018 Hasura.io

*/

package v1alpha1

import (
	v1alpha1 "github.com/hasura/gitkube/pkg/apis/gitkube.sh/v1alpha1"
	scheme "github.com/hasura/gitkube/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// RemotesGetter has a method to return a RemoteInterface.
// A group's client should implement this interface.
type RemotesGetter interface {
	Remotes(namespace string) RemoteInterface
}

// RemoteInterface has methods to work with Remote resources.
type RemoteInterface interface {
	Create(*v1alpha1.Remote) (*v1alpha1.Remote, error)
	Update(*v1alpha1.Remote) (*v1alpha1.Remote, error)
	UpdateStatus(*v1alpha1.Remote) (*v1alpha1.Remote, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Remote, error)
	List(opts v1.ListOptions) (*v1alpha1.RemoteList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Remote, err error)
	RemoteExpansion
}

// remotes implements RemoteInterface
type remotes struct {
	client rest.Interface
	ns     string
}

// newRemotes returns a Remotes
func newRemotes(c *GitkubeV1alpha1Client, namespace string) *remotes {
	return &remotes{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the remote, and returns the corresponding remote object, and an error if there is any.
func (c *remotes) Get(name string, options v1.GetOptions) (result *v1alpha1.Remote, err error) {
	result = &v1alpha1.Remote{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("remotes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Remotes that match those selectors.
func (c *remotes) List(opts v1.ListOptions) (result *v1alpha1.RemoteList, err error) {
	result = &v1alpha1.RemoteList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("remotes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested remotes.
func (c *remotes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("remotes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a remote and creates it.  Returns the server's representation of the remote, and an error, if there is any.
func (c *remotes) Create(remote *v1alpha1.Remote) (result *v1alpha1.Remote, err error) {
	result = &v1alpha1.Remote{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("remotes").
		Body(remote).
		Do().
		Into(result)
	return
}

// Update takes the representation of a remote and updates it. Returns the server's representation of the remote, and an error, if there is any.
func (c *remotes) Update(remote *v1alpha1.Remote) (result *v1alpha1.Remote, err error) {
	result = &v1alpha1.Remote{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("remotes").
		Name(remote.Name).
		Body(remote).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *remotes) UpdateStatus(remote *v1alpha1.Remote) (result *v1alpha1.Remote, err error) {
	result = &v1alpha1.Remote{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("remotes").
		Name(remote.Name).
		SubResource("status").
		Body(remote).
		Do().
		Into(result)
	return
}

// Delete takes name of the remote and deletes it. Returns an error if one occurs.
func (c *remotes) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("remotes").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *remotes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("remotes").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched remote.
func (c *remotes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Remote, err error) {
	result = &v1alpha1.Remote{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("remotes").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}

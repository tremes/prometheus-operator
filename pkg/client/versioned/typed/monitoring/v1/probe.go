// Copyright The prometheus-operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/tremes/prometheus-operator/pkg/apis/monitoring/v1"
	scheme "github.com/tremes/prometheus-operator/pkg/client/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ProbesGetter has a method to return a ProbeInterface.
// A group's client should implement this interface.
type ProbesGetter interface {
	Probes(namespace string) ProbeInterface
}

// ProbeInterface has methods to work with Probe resources.
type ProbeInterface interface {
	Create(ctx context.Context, probe *v1.Probe, opts metav1.CreateOptions) (*v1.Probe, error)
	Update(ctx context.Context, probe *v1.Probe, opts metav1.UpdateOptions) (*v1.Probe, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Probe, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ProbeList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Probe, err error)
	ProbeExpansion
}

// probes implements ProbeInterface
type probes struct {
	client rest.Interface
	ns     string
}

// newProbes returns a Probes
func newProbes(c *ObservabilityV1Client, namespace string) *probes {
	return &probes{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the probe, and returns the corresponding probe object, and an error if there is any.
func (c *probes) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Probe, err error) {
	result = &v1.Probe{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("probes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Probes that match those selectors.
func (c *probes) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ProbeList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ProbeList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("probes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested probes.
func (c *probes) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("probes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a probe and creates it.  Returns the server's representation of the probe, and an error, if there is any.
func (c *probes) Create(ctx context.Context, probe *v1.Probe, opts metav1.CreateOptions) (result *v1.Probe, err error) {
	result = &v1.Probe{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("probes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(probe).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a probe and updates it. Returns the server's representation of the probe, and an error, if there is any.
func (c *probes) Update(ctx context.Context, probe *v1.Probe, opts metav1.UpdateOptions) (result *v1.Probe, err error) {
	result = &v1.Probe{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("probes").
		Name(probe.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(probe).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the probe and deletes it. Returns an error if one occurs.
func (c *probes) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("probes").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *probes) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("probes").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched probe.
func (c *probes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Probe, err error) {
	result = &v1.Probe{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("probes").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

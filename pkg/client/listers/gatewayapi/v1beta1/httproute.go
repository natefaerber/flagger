/*
Copyright 2020 The Flux authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/fluxcd/flagger/pkg/apis/gatewayapi/v1beta1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// HTTPRouteLister helps list HTTPRoutes.
// All objects returned here must be treated as read-only.
type HTTPRouteLister interface {
	// List lists all HTTPRoutes in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.HTTPRoute, err error)
	// HTTPRoutes returns an object that can list and get HTTPRoutes.
	HTTPRoutes(namespace string) HTTPRouteNamespaceLister
	HTTPRouteListerExpansion
}

// hTTPRouteLister implements the HTTPRouteLister interface.
type hTTPRouteLister struct {
	listers.ResourceIndexer[*v1beta1.HTTPRoute]
}

// NewHTTPRouteLister returns a new HTTPRouteLister.
func NewHTTPRouteLister(indexer cache.Indexer) HTTPRouteLister {
	return &hTTPRouteLister{listers.New[*v1beta1.HTTPRoute](indexer, v1beta1.Resource("httproute"))}
}

// HTTPRoutes returns an object that can list and get HTTPRoutes.
func (s *hTTPRouteLister) HTTPRoutes(namespace string) HTTPRouteNamespaceLister {
	return hTTPRouteNamespaceLister{listers.NewNamespaced[*v1beta1.HTTPRoute](s.ResourceIndexer, namespace)}
}

// HTTPRouteNamespaceLister helps list and get HTTPRoutes.
// All objects returned here must be treated as read-only.
type HTTPRouteNamespaceLister interface {
	// List lists all HTTPRoutes in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.HTTPRoute, err error)
	// Get retrieves the HTTPRoute from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.HTTPRoute, error)
	HTTPRouteNamespaceListerExpansion
}

// hTTPRouteNamespaceLister implements the HTTPRouteNamespaceLister
// interface.
type hTTPRouteNamespaceLister struct {
	listers.ResourceIndexer[*v1beta1.HTTPRoute]
}

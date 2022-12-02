/*
Copyright 2021 The Knative Authors

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

package v1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1 "knative.dev/eventing/pkg/apis/sources/v1"
)

// PingSourceLister helps list PingSources.
// All objects returned here must be treated as read-only.
type PingSourceLister interface {
	// List lists all PingSources in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.PingSource, err error)
	// PingSources returns an object that can list and get PingSources.
	PingSources(namespace string) PingSourceNamespaceLister
	PingSourceListerExpansion
}

// pingSourceLister implements the PingSourceLister interface.
type pingSourceLister struct {
	indexer cache.Indexer
}

// NewPingSourceLister returns a new PingSourceLister.
func NewPingSourceLister(indexer cache.Indexer) PingSourceLister {
	return &pingSourceLister{indexer: indexer}
}

// List lists all PingSources in the indexer.
func (s *pingSourceLister) List(selector labels.Selector) (ret []*v1.PingSource, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.PingSource))
	})
	return ret, err
}

// PingSources returns an object that can list and get PingSources.
func (s *pingSourceLister) PingSources(namespace string) PingSourceNamespaceLister {
	return pingSourceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PingSourceNamespaceLister helps list and get PingSources.
// All objects returned here must be treated as read-only.
type PingSourceNamespaceLister interface {
	// List lists all PingSources in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.PingSource, err error)
	// Get retrieves the PingSource from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.PingSource, error)
	PingSourceNamespaceListerExpansion
}

// pingSourceNamespaceLister implements the PingSourceNamespaceLister
// interface.
type pingSourceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PingSources in the indexer for a given namespace.
func (s pingSourceNamespaceLister) List(selector labels.Selector) (ret []*v1.PingSource, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.PingSource))
	})
	return ret, err
}

// Get retrieves the PingSource from the indexer for a given namespace and name.
func (s pingSourceNamespaceLister) Get(name string) (*v1.PingSource, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("pingsource"), name)
	}
	return obj.(*v1.PingSource), nil
}
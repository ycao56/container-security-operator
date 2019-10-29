/*
Copyright The Kubernetes Authors.

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

package internalversion

import (
	v1alpha1 "github.com/coreos-inc/security-labeller/apis/secscan/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ImageManifestVulnLister helps list ImageManifestVulns.
type ImageManifestVulnLister interface {
	// List lists all ImageManifestVulns in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ImageManifestVuln, err error)
	// ImageManifestVulns returns an object that can list and get ImageManifestVulns.
	ImageManifestVulns(namespace string) ImageManifestVulnNamespaceLister
	ImageManifestVulnListerExpansion
}

// imageManifestVulnLister implements the ImageManifestVulnLister interface.
type imageManifestVulnLister struct {
	indexer cache.Indexer
}

// NewImageManifestVulnLister returns a new ImageManifestVulnLister.
func NewImageManifestVulnLister(indexer cache.Indexer) ImageManifestVulnLister {
	return &imageManifestVulnLister{indexer: indexer}
}

// List lists all ImageManifestVulns in the indexer.
func (s *imageManifestVulnLister) List(selector labels.Selector) (ret []*v1alpha1.ImageManifestVuln, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ImageManifestVuln))
	})
	return ret, err
}

// ImageManifestVulns returns an object that can list and get ImageManifestVulns.
func (s *imageManifestVulnLister) ImageManifestVulns(namespace string) ImageManifestVulnNamespaceLister {
	return imageManifestVulnNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ImageManifestVulnNamespaceLister helps list and get ImageManifestVulns.
type ImageManifestVulnNamespaceLister interface {
	// List lists all ImageManifestVulns in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ImageManifestVuln, err error)
	// Get retrieves the ImageManifestVuln from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ImageManifestVuln, error)
	ImageManifestVulnNamespaceListerExpansion
}

// imageManifestVulnNamespaceLister implements the ImageManifestVulnNamespaceLister
// interface.
type imageManifestVulnNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ImageManifestVulns in the indexer for a given namespace.
func (s imageManifestVulnNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ImageManifestVuln, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ImageManifestVuln))
	})
	return ret, err
}

// Get retrieves the ImageManifestVuln from the indexer for a given namespace and name.
func (s imageManifestVulnNamespaceLister) Get(name string) (*v1alpha1.ImageManifestVuln, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("imagemanifestvuln"), name)
	}
	return obj.(*v1alpha1.ImageManifestVuln), nil
}
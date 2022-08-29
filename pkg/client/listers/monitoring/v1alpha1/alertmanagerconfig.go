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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/tremes/prometheus-operator/pkg/apis/monitoring/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AlertmanagerConfigLister helps list AlertmanagerConfigs.
// All objects returned here must be treated as read-only.
type AlertmanagerConfigLister interface {
	// List lists all AlertmanagerConfigs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AlertmanagerConfig, err error)
	// AlertmanagerConfigs returns an object that can list and get AlertmanagerConfigs.
	AlertmanagerConfigs(namespace string) AlertmanagerConfigNamespaceLister
	AlertmanagerConfigListerExpansion
}

// alertmanagerConfigLister implements the AlertmanagerConfigLister interface.
type alertmanagerConfigLister struct {
	indexer cache.Indexer
}

// NewAlertmanagerConfigLister returns a new AlertmanagerConfigLister.
func NewAlertmanagerConfigLister(indexer cache.Indexer) AlertmanagerConfigLister {
	return &alertmanagerConfigLister{indexer: indexer}
}

// List lists all AlertmanagerConfigs in the indexer.
func (s *alertmanagerConfigLister) List(selector labels.Selector) (ret []*v1alpha1.AlertmanagerConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AlertmanagerConfig))
	})
	return ret, err
}

// AlertmanagerConfigs returns an object that can list and get AlertmanagerConfigs.
func (s *alertmanagerConfigLister) AlertmanagerConfigs(namespace string) AlertmanagerConfigNamespaceLister {
	return alertmanagerConfigNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AlertmanagerConfigNamespaceLister helps list and get AlertmanagerConfigs.
// All objects returned here must be treated as read-only.
type AlertmanagerConfigNamespaceLister interface {
	// List lists all AlertmanagerConfigs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AlertmanagerConfig, err error)
	// Get retrieves the AlertmanagerConfig from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.AlertmanagerConfig, error)
	AlertmanagerConfigNamespaceListerExpansion
}

// alertmanagerConfigNamespaceLister implements the AlertmanagerConfigNamespaceLister
// interface.
type alertmanagerConfigNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AlertmanagerConfigs in the indexer for a given namespace.
func (s alertmanagerConfigNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AlertmanagerConfig, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AlertmanagerConfig))
	})
	return ret, err
}

// Get retrieves the AlertmanagerConfig from the indexer for a given namespace and name.
func (s alertmanagerConfigNamespaceLister) Get(name string) (*v1alpha1.AlertmanagerConfig, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("alertmanagerconfig"), name)
	}
	return obj.(*v1alpha1.AlertmanagerConfig), nil
}

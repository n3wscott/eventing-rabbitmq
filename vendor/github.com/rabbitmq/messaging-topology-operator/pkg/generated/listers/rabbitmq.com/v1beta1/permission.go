/*
RabbitMQ Messaging Topology Kubernetes Operator
Copyright 2021 VMware, Inc.

This product is licensed to you under the Mozilla Public License 2.0 license (the "License").  You may not use this product except in compliance with the Mozilla 2.0 License.

This product may include a number of subcomponents with separate copyright notices and license terms. Your use of these subcomponents is subject to the terms and conditions of the subcomponent's license, as noted in the LICENSE file.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/rabbitmq/messaging-topology-operator/api/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PermissionLister helps list Permissions.
// All objects returned here must be treated as read-only.
type PermissionLister interface {
	// List lists all Permissions in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.Permission, err error)
	// Permissions returns an object that can list and get Permissions.
	Permissions(namespace string) PermissionNamespaceLister
	PermissionListerExpansion
}

// permissionLister implements the PermissionLister interface.
type permissionLister struct {
	indexer cache.Indexer
}

// NewPermissionLister returns a new PermissionLister.
func NewPermissionLister(indexer cache.Indexer) PermissionLister {
	return &permissionLister{indexer: indexer}
}

// List lists all Permissions in the indexer.
func (s *permissionLister) List(selector labels.Selector) (ret []*v1beta1.Permission, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.Permission))
	})
	return ret, err
}

// Permissions returns an object that can list and get Permissions.
func (s *permissionLister) Permissions(namespace string) PermissionNamespaceLister {
	return permissionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PermissionNamespaceLister helps list and get Permissions.
// All objects returned here must be treated as read-only.
type PermissionNamespaceLister interface {
	// List lists all Permissions in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.Permission, err error)
	// Get retrieves the Permission from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.Permission, error)
	PermissionNamespaceListerExpansion
}

// permissionNamespaceLister implements the PermissionNamespaceLister
// interface.
type permissionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Permissions in the indexer for a given namespace.
func (s permissionNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.Permission, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.Permission))
	})
	return ret, err
}

// Get retrieves the Permission from the indexer for a given namespace and name.
func (s permissionNamespaceLister) Get(name string) (*v1beta1.Permission, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("permission"), name)
	}
	return obj.(*v1beta1.Permission), nil
}

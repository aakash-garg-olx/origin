// This file was automatically generated by lister-gen

package internalversion

import (
	authorization "github.com/openshift/origin/pkg/authorization/apis/authorization"
	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClusterPolicyLister helps list ClusterPolicies.
type ClusterPolicyLister interface {
	// List lists all ClusterPolicies in the indexer.
	List(selector labels.Selector) (ret []*authorization.ClusterPolicy, err error)
	// Get retrieves the ClusterPolicy from the index for a given name.
	Get(name string) (*authorization.ClusterPolicy, error)
	ClusterPolicyListerExpansion
}

// clusterPolicyLister implements the ClusterPolicyLister interface.
type clusterPolicyLister struct {
	indexer cache.Indexer
}

// NewClusterPolicyLister returns a new ClusterPolicyLister.
func NewClusterPolicyLister(indexer cache.Indexer) ClusterPolicyLister {
	return &clusterPolicyLister{indexer: indexer}
}

// List lists all ClusterPolicies in the indexer.
func (s *clusterPolicyLister) List(selector labels.Selector) (ret []*authorization.ClusterPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*authorization.ClusterPolicy))
	})
	return ret, err
}

// Get retrieves the ClusterPolicy from the index for a given name.
func (s *clusterPolicyLister) Get(name string) (*authorization.ClusterPolicy, error) {
	key := &authorization.ClusterPolicy{ObjectMeta: v1.ObjectMeta{Name: name}}
	obj, exists, err := s.indexer.Get(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(authorization.Resource("clusterpolicy"), name)
	}
	return obj.(*authorization.ClusterPolicy), nil
}

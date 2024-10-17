// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/api/operator/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// KubeAPIServerLister helps list KubeAPIServers.
// All objects returned here must be treated as read-only.
type KubeAPIServerLister interface {
	// List lists all KubeAPIServers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.KubeAPIServer, err error)
	// Get retrieves the KubeAPIServer from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.KubeAPIServer, error)
	KubeAPIServerListerExpansion
}

// kubeAPIServerLister implements the KubeAPIServerLister interface.
type kubeAPIServerLister struct {
	listers.ResourceIndexer[*v1.KubeAPIServer]
}

// NewKubeAPIServerLister returns a new KubeAPIServerLister.
func NewKubeAPIServerLister(indexer cache.Indexer) KubeAPIServerLister {
	return &kubeAPIServerLister{listers.New[*v1.KubeAPIServer](indexer, v1.Resource("kubeapiserver"))}
}
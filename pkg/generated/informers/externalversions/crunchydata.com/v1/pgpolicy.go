/*
 Copyright 2020 Crunchy Data Solutions, Inc.
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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	crunchydatacomv1 "github.com/crunchydata/postgres-operator/pkg/apis/crunchydata.com/v1"
	versioned "github.com/crunchydata/postgres-operator/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/crunchydata/postgres-operator/pkg/generated/informers/externalversions/internalinterfaces"
	v1 "github.com/crunchydata/postgres-operator/pkg/generated/listers/crunchydata.com/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// PgpolicyInformer provides access to a shared informer and lister for
// Pgpolicies.
type PgpolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.PgpolicyLister
}

type pgpolicyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewPgpolicyInformer constructs a new informer for Pgpolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPgpolicyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPgpolicyInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredPgpolicyInformer constructs a new informer for Pgpolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPgpolicyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CrunchydataV1().Pgpolicies(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CrunchydataV1().Pgpolicies(namespace).Watch(context.TODO(), options)
			},
		},
		&crunchydatacomv1.Pgpolicy{},
		resyncPeriod,
		indexers,
	)
}

func (f *pgpolicyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPgpolicyInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *pgpolicyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&crunchydatacomv1.Pgpolicy{}, f.defaultInformer)
}

func (f *pgpolicyInformer) Lister() v1.PgpolicyLister {
	return v1.NewPgpolicyLister(f.Informer().GetIndexer())
}

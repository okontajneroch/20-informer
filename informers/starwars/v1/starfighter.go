package v1

import (
	"context"
	"fmt"
	starwarsv1 "github.com/okontajneroch/starwars/api/starwars/v1"
	starwarsclientset "github.com/okontajneroch/starwars/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
	"time"
)

type StarfighterInformer struct {
	clientset  starwarsclientset.Interface
	namespace  string
	store      cache.Store
	controller cache.Controller
}

func NewStarfighterInformer(client starwarsclientset.Interface, namespace string) *StarfighterInformer {
	sfi := &StarfighterInformer{
		clientset: client,
		namespace: namespace,
	}

	// Create a ListWatch for Starfighter resources.
	lw := &cache.ListWatch{
		ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
			return client.StarwarsV1().Starfighters(namespace).List(context.TODO(), options)
		},
		WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
			return client.StarwarsV1().Starfighters(namespace).Watch(context.TODO(), options)
		},
	}

	// Create the informer
	sfi.store, sfi.controller = cache.NewInformerWithOptions(cache.InformerOptions{
		ListerWatcher: lw,
		ResyncPeriod:  30 * time.Second,
		ObjectType:    &starwarsv1.Starfighter{},
		Handler: cache.ResourceEventHandlerFuncs{
			AddFunc:    onStarfighterAdded,
			UpdateFunc: onStarfighterUpdated,
			DeleteFunc: onStarfighterDeleted,
		},
	})

	return sfi
}

func (sfi *StarfighterInformer) Start(stopCh <-chan struct{}) {
	go sfi.controller.Run(stopCh)
	cache.WaitForCacheSync(stopCh, sfi.controller.HasSynced)
}

func onStarfighterAdded(obj any) {
	sf := obj.(*starwarsv1.Starfighter)
	fmt.Printf("Starfighter added: %s\n", sf.Name)
}

func onStarfighterUpdated(oldObj, newObj any) {
	sf := newObj.(*starwarsv1.Starfighter)
	fmt.Printf("Starfighter updated: %s\n", sf.Name)
}

func onStarfighterDeleted(obj any) {
	sf := obj.(*starwarsv1.Starfighter)
	fmt.Printf("Starfighter deleted: %s\n", sf.Name)
}

package main

import (
	"fmt"
	"github.com/okontajneroch/starwars/clientset"
	swinformer "github.com/okontajneroch/starwars/informers/starwars/v1"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// vytvorim si stop kanál, ktorý bude uzavretý unixovským SIGINT alebo SIGTERM signálom
	stopCh := make(chan struct{})
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigCh
		close(stopCh)
	}()

	// vytrovím si klienta na komunikáciu s kube-apiserverom
	config, _ := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(),
		nil,
	).ClientConfig()
	client, _ := clientset.NewForConfig(config)

	// vytvorím si informer pre "default" namespace a spustím ho
	informer := swinformer.NewStarfighterInformer(client, "default")
	informer.Start(stopCh)
	fmt.Println("Starfighter informers is running. Press Ctrl+C to exit.")

	// čakám na ukončenie programu
	<-stopCh
	fmt.Println("Shutting down informers.")
}

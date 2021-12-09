package main // import "github.com/marcsauter/kube-tmux"

import (
	"fmt"
	"os"

	"k8s.io/client-go/tools/clientcmd"
)

const (
	defaultNamespace = "default"
)

func main() {
	if len(os.Args) > 1 {
		format = os.Args[1]
	}
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	configOverrides := &clientcmd.ConfigOverrides{}
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)
	config, err := kubeConfig.RawConfig()
	if err != nil {
		os.Exit(1)
	}
	currentContext := config.CurrentContext
	currentNamespace := config.Contexts[currentContext].Namespace
	if currentNamespace == "" {
		currentNamespace = defaultNamespace
	}

	fmt.Printf("[%s/%s]", currentContext, currentNamespace)
}

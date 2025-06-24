/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List Kubernetes deployments in the namespace",

	Run: func(cmd *cobra.Command, args []string) {
		log.Debug().Msg("k8s-controller list command executed.")
		InitConfig()
		log.Debug().Msgf("Using namespace %s", Namespace)
		log.Debug().Msgf("Using config %s", KubeConfigPath)

		clientset, err := getKubeClient(KubeConfigPath)
		if err != nil {
			log.Error().Err(err).Msg("Failed to create Kubernetes client")
			os.Exit(1)
		}
		deployments, err := clientset.AppsV1().Deployments(Namespace).List(context.Background(), metav1.ListOptions{})
		if err != nil {
			log.Error().Err(err).Msg("Failed to list deployments")
			os.Exit(1)
		}
		if AllNamespaces {
			fmt.Printf("Found %d deployments in ALL namespace:\n", len(deployments.Items))
		} else {
			fmt.Printf("Found %d deployments in '%s' namespace:\n", len(deployments.Items), Namespace)
		}

		for _, d := range deployments.Items {
			fmt.Println("-", d.Name)
		}
	},
}

func getKubeClient(kubeconfigPath string) (*kubernetes.Clientset, error) {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(config)
}

func init() {
	rootCmd.AddCommand(listCmd)
}

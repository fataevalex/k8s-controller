package cmd

import (
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
)

var (
	KubeConfigPath string = ""
	Namespace      string = "default"
	AllNamespaces         = false
)

func InitConfig() {
	KubeConfigPath = setKubeConfigPath()
	Namespace = setNamespace()
}

func setKubeConfigPath() string {
	if env := os.Getenv("KUBECONFIG"); env != "" {
		return env
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(home, ".kube", "config")
}

func setNamespace() string {
	log.Debug().Msgf("All namespace =  %t", AllNamespaces)
	if AllNamespaces {
		Namespace = ""
	}
	return Namespace
}

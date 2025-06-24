//go:build !test
// +build !test

package cmd

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

func applyTestDeployment(t *testing.T) {
	cmd := exec.Command("kubectl", "apply", "-f", "testdata/nginx-deployment.yaml")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("failed to apply test deployment: %s\n%s", err, output)
	}
}

func cleanupTestDeployment(t *testing.T) {
	cmd := exec.Command("kubectl", "delete", "-f", "testdata/nginx-deployment.yaml", "--ignore-not-found")
	_ = cmd.Run()
}

func TestDeploymentIntegration(t *testing.T) {
	defer cleanupTestDeployment(t)

	applyTestDeployment(t)
	//execute list test
	listOut := &bytes.Buffer{}
	rootCmd.SetOut(listOut)
	rootCmd.SetArgs([]string{"list"})
	if err := rootCmd.Execute(); err != nil {
		t.Fatalf("list command failed: %v", err)
	}

	if !strings.Contains(listOut.String(), "test-deployment") {
		t.Errorf("expected 'test-deployment' in list output, got:\n%s", listOut.String())
	}
	//execute delete test
	deleteOut := &bytes.Buffer{}
	rootCmd.SetOut(deleteOut)
	rootCmd.SetArgs([]string{"delete", "test-deployment"})
	if err := rootCmd.Execute(); err != nil {
		t.Fatalf("delete command failed: %v", err)
	}

	if !strings.Contains(deleteOut.String(), "deleted") {
		t.Errorf("expected 'deleted' in delete output, got:\n%s", deleteOut.String())
	}

}

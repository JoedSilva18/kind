/*
Copyright 2018 The Kubernetes Authors.

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

// Package cluster implements the `list` command
package cluster

import (
	"container/list"
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
	"sigs.k8s.io/kind/pkg/cmd"
	"sigs.k8s.io/kind/pkg/internal/cli"
	"sigs.k8s.io/kind/pkg/log"
	"strings"
	"time"
)

type flagpole struct {
	Name       string
	Config     string
	ImageName  string
	Retain     bool
	Wait       time.Duration
	Kubeconfig string
}

// NewCommand returns a new cobra.Command for list docker images
func NewCommand(logger log.Logger, streams cmd.IOStreams) *cobra.Command {
	flags := &flagpole{}
	cmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "load-images",
		Short: "List all installed docker images",
		Long:  "List all installed docker images",
		RunE: func(cmd *cobra.Command, args []string) error {
			cli.OverrideDefaultName(cmd.Flags())
			return runE(logger, streams, flags)
		},
	}

	return cmd
}

func runE(logger log.Logger, streams cmd.IOStreams, flags *flagpole) error {
	cmd, err := exec.Command("bash", "-c", "kubectl get nodes").Output()

	if err != nil {
		logger.Warn(fmt.Sprint(err))
	}

	kubectlOutputLines := strings.Split(string(cmd), "\n")
	nodeNames := list.New()

	for i := 1; i < len(kubectlOutputLines); i++ {
		kubctlNode := strings.Split(kubectlOutputLines[i], " ")
		kubctlImageName := kubctlNode[0]
		nodeNames.PushBack(kubctlImageName)
	}

	for name := nodeNames.Front(); name != nil; name = name.Next() {
		if fmt.Sprint(name.Value) != "" {
			logger.V(0).Infof("Node name: " + fmt.Sprint(name.Value))

			imageName := "docker exec " + fmt.Sprint(name.Value) + " crictl images"
			cmd := exec.Command("bash", "-c", imageName)
			output, err := cmd.CombinedOutput()

			if err != nil {
				logger.Warn(fmt.Sprint(err) + ": " + string(output))
			} else {
				logger.V(0).Infof(string(output))
			}
		}
	}

	return nil
}

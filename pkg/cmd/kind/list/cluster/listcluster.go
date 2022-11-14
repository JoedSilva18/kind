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
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
	"time"

	"sigs.k8s.io/kind/pkg/cmd"
	"sigs.k8s.io/kind/pkg/log"

	"sigs.k8s.io/kind/pkg/internal/cli"
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
	logger.Warn("Teste comando novo")
	cmd, err := exec.Command("bash", "-c", "docker images").Output()

	if err != nil {
		fmt.Println(err)
	}

	logger.Warn(string(cmd))

	return nil
}

/*
 * Copyright (C) 2022 Canonical Ltd
 *
 *  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 *  in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *
 * SPDX-License-Identifier: Apache-2.0'
 */

package main

import (
	"fmt"
	"os/exec"

	"github.com/canonical/edgex-snap-hooks/v3/env"
	"github.com/canonical/edgex-snap-hooks/v3/log"
	"github.com/canonical/edgex-snap-hooks/v3/snapctl"
)

// installConfig copies all config files from $SNAP to $SNAP_DATA.
func installConfig() error {
	path := "/"

	out, err := exec.Command("cp", "--recursive", env.Snap+path, env.SnapData+path).CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s: %s", out, err)
	}

	return nil
}

func install() {
	log.SetComponentName("install")

	// Install default config files only if no config provider is connected
	isConnected, err := snapctl.IsConnected("ekuiper-data").Run()
	if err != nil {
		log.Fatalf("Error checking interface connection: %s", err)
	}
	if !isConnected {
		if err := installConfig(); err != nil {
			log.Fatalf("Error installing config files: %s", err)
		}
	}
}

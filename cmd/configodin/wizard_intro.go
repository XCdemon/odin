// Copyright 2019-2020 The OdinBrowser Core Developers
// This file is part of the OdinBrowser.
//
// The OdinBrowser is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The OdinBrowser is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the OdinBrowser. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

// makeWizard creates and returns a new puppeth wizard.
func makeWizard(network string) *wizard {
	return &wizard{
		network: network,
		conf: config{
			Servers: make(map[string][]byte),
		},
		servers:  make(map[string]*sshClient),
		services: make(map[string][]string),
		in:       bufio.NewReader(os.Stdin),
	}
}

// run displays some useful infos to the user, starting on the journey of
// setting up a new or managing an existing Ethereum private network.
func (w *wizard) run() {
	// Make sure we have a good network name to work with	fmt.Println()
	// Docker accepts hyphens in image names, but doesn't like it for container names
	fmt.Println("Please specify a network id")
	var id uint64
	for {
		id = uint64(w.readDefaultInt(1))
		break
	}
	w.network = "genesis"
	// Load initial configurations and connect to all live servers
	w.conf.path = filepath.Join(os.Getenv("HOME"), ".puppeth", w.network)

	w.makeGenesis(id)
}

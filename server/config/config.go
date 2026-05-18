// Copyright 2024 The etcd Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package config provides configuration structures and validation for the etcd server.
package config

import (
	"fmt"
	"net/url"
	"time"
)

const (
	// DefaultName is the default name for an etcd member.
	DefaultName = "default"

	// DefaultMaxSnapshots is the default maximum number of snapshots to retain.
	DefaultMaxSnapshots = 5

	// DefaultMaxWALs is the default maximum number of WAL files to retain.
	DefaultMaxWALs = 5

	// DefaultTickMs is the default interval for Raft heartbeat ticks in milliseconds.
	DefaultTickMs = 100

	// DefaultElectionMs is the default timeout for Raft election in milliseconds.
	DefaultElectionMs = 1000

	// DefaultListenPeerURLs is the default URL for peer communication.
	DefaultListenPeerURLs = "http://localhost:2380"

	// DefaultListenClientURLs is the default URL for client communication.
	DefaultListenClientURLs = "http://localhost:2379"

	// DefaultMaxRequestBytes is the default maximum request size in bytes (1.5 MiB).
	DefaultMaxRequestBytes = 1.5 * 1024 * 1024
)

// ServerConfig holds the configuration for an etcd server instance.
type ServerConfig struct {
	// Name is the human-readable name for this etcd member.
	Name string

	// DataDir is the path to the data directory.
	DataDir string

	// WALDir is the dedicated path for WAL files. If empty, DataDir is used.
	WALDir string

	// SnapshotDir is the dedicated path for snapshot files. If empty, DataDir is used.
	SnapshotDir string

	// ListenPeerURLs is the list of URLs to listen on for peer traffic.
	ListenPeerURLs []url.URL

	// ListenClientURLs is the list of URLs to listen on for client traffic.
	ListenClientURLs []url.URL

	// AdvertisePeerURLs is the list of peer URLs to advertise to the cluster.
	AdvertisePeerURLs []url.URL

	// AdvertiseClientURLs is the list of client URLs to advertise to clients.
	AdvertiseClientURLs []url.URL

	// MaxSnapshots is the maximum number of snapshots to retain.
	MaxSnapshots uint

	// MaxWALs is the maximum number of WAL files to retain.
	MaxWALs uint

	// TickMs is the interval for Raft heartbeat ticks in milliseconds.
	TickMs uint

	// ElectionMs is the timeout for Raft elections in milliseconds.
	ElectionMs uint

	// MaxRequestBytes is the maximum size of a client request.
	MaxRequestBytes uint

	// InitialClusterToken is the token for the initial cluster state.
	InitialClusterToken string

	// InitialCluster is the initial cluster configuration for bootstrapping.
	InitialCluster string

	// ClusterState indicates whether this is a new or existing cluster.
	// Valid values: "new" or "existing".
	ClusterState string
}

// NewDefaultServerConfig returns a ServerConfig populated with sensible defaults.
func NewDefaultServerConfig() *ServerConfig {
	peerURL, _ := url.Parse(DefaultListenPeerURLs)
	clientURL, _ := url.Parse(DefaultListenClientURLs)

	return &ServerConfig{
		Name:                DefaultName,
		MaxSnapshots:        DefaultMaxSnapshots,
		MaxWALs:             DefaultMaxWALs,
		TickMs:              DefaultTickMs,
		ElectionMs:          DefaultElectionMs,
		MaxRequestBytes:     DefaultMaxRequestBytes,
		ListenPeerURLs:      []url.URL{*peerURL},
		ListenClientURLs:    []url.URL{*clientURL},
		AdvertisePeerURLs:   []url.URL{*peerURL},
		AdvertiseClientURLs: []url.URL{*clientURL},
		ClusterState:        "new",
	}
}

// Validate checks that the ServerConfig has valid values and returns an error if not.
func (c *ServerConfig) Validate() error {
	if c.Name == "" {
		return fmt.Errorf("member name cannot be empty")
	}
	if c.DataDir == "" {
		return fmt.Errorf("data directory cannot be empty")
	}
	if c.TickMs == 0 {
		return fmt.Errorf("tick interval must be greater than 0")
	}
	if c.ElectionMs < 5*c.TickMs {
		return fmt.Errorf("election timeout (%dms) must be at least 5x tick interval (%dms)",
			c.ElectionMs, c.TickMs)
	}
	if c.ClusterState != "new" && c.ClusterState != "existing" {
		return fmt.Errorf("cluster state must be \"new\" or \"existing\", got: %q", c.ClusterState)
	}
	if len(c.ListenPeerURLs) == 0 {
		return fmt.Errorf("at least one listen peer URL must be specified")
	}
	if len(c.ListenClientURLs) == 0 {
		return fmt.Errorf("at least one listen client URL must be specified")
	}
	return nil
}

// ElectionTimeout returns the election timeout as a time.Duration.
func (c *ServerConfig) ElectionTimeout() time.Duration {
	return time.Duration(c.ElectionMs) * time.Millisecond
}

// TickInterval returns the tick interval as a time.Duration.
func (c *ServerConfig) TickInterval() time.Duration {
	return time.Duration(c.TickMs) * time.Millisecond
}

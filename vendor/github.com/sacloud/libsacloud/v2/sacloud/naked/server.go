// Copyright 2016-2019 The Libsacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package naked

import (
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Server サーバ
type Server struct {
	ID                types.ID               `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Name              string                 `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	Description       string                 `yaml:"description"`
	Tags              types.Tags             `yaml:"tags"`
	Icon              *Icon                  `json:",omitempty" yaml:"icon,omitempty" structs:",omitempty"`
	CreatedAt         *time.Time             `json:",omitempty" yaml:"created_at,omitempty" structs:",omitempty"`
	ModifiedAt        *time.Time             `json:",omitempty" yaml:"modified_at,omitempty" structs:",omitempty"`
	Availability      types.EAvailability    `json:",omitempty" yaml:"availability,omitempty" structs:",omitempty"`
	HostName          string                 `json:",omitempty" yaml:"host_name,omitempty" structs:",omitempty"`
	ServiceClass      string                 `json:",omitempty" yaml:"service_class,omitempty" structs:",omitempty"`
	InterfaceDriver   types.EInterfaceDriver `json:",omitempty" yaml:"interface_driver,omitempty" structs:",omitempty"`
	ServerPlan        *ServerPlan            `json:",omitempty" yaml:"server_plan,omitempty" structs:",omitempty"`
	Zone              *Zone                  `json:",omitempty" yaml:"zone,omitempty" structs:",omitempty"`
	Instance          *Instance              `json:",omitempty" yaml:"instance,omitempty" structs:",omitempty"`
	Disks             []*Disk                `json:",omitempty" yaml:"disks,omitempty" structs:",omitempty"`
	Interfaces        []*Interface           `json:",omitempty" yaml:"interfaces,omitempty" structs:",omitempty"`
	PrivateHost       *PrivateHost           `json:",omitempty" yaml:"private_host,omitempty" structs:",omitempty"`
	WaitDiskMigration bool                   `json:",omitempty" yaml:"wait_disk_migration,omitempty" structs:",omitempty"`
	ConnectedSwitches []*ConnectedSwitch     `json:",omitempty" yaml:"connected_switches,omitempty" structs:",omitempty"`
}

// ConnectedSwitch サーバ作成時に指定する接続先スイッチ
type ConnectedSwitch struct {
	ID    types.ID     `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Scope types.EScope `json:",omitempty" yaml:"scope,omitempty" structs:",omitempty"`
}

// MouseRequestButtons マウスボタン
type MouseRequestButtons struct {
	L bool // 左ボタン
	R bool // 右ボタン
	M bool // 中ボタン
}

// DeleteServerWithDiskParameter サーバ削除時に接続されているディスクを削除するためのパラメータ
type DeleteServerWithDiskParameter struct {
	WithDisk []types.ID
}

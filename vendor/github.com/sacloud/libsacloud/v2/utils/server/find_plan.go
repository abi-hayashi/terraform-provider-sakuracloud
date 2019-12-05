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

package server

import (
	"context"
	"errors"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/search"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// PlanFinder .
type PlanFinder interface {
	Find(ctx context.Context, zone string, conditions *sacloud.FindCondition) (*sacloud.ServerPlanFindResult, error)
}

// FindPlanRequest サーバプラン検索パラメータ
type FindPlanRequest struct {
	CPU        int
	MemoryGB   int
	Commitment types.ECommitment
	Generation types.EPlanGeneration
}

func (f *FindPlanRequest) findCondition() *sacloud.FindCondition {
	cond := &sacloud.FindCondition{
		Sort: search.SortKeys{
			{Key: "Generation", Order: search.SortDesc},
		},
		Filter: search.Filter{
			search.Key("Commitment"): types.Commitments.Standard,
		},
		Count: 1000,
	}
	if f.CPU > 0 {
		cond.Filter[search.Key("CPU")] = f.CPU
	}
	if f.MemoryGB > 0 {
		cond.Filter[search.Key("MemoryMB")] = f.MemoryGB * 1024
	}
	if f.Generation != types.PlanGenerations.Default {
		cond.Filter[search.Key("Generation")] = f.Generation
	}
	if f.Commitment != types.Commitments.Unknown && f.Commitment != types.Commitments.Standard {
		cond.Filter[search.Key("Commitment")] = f.Commitment
	}
	return cond
}

// FindPlan サーバプラン検索
func FindPlan(ctx context.Context, finder PlanFinder, zone string, param *FindPlanRequest) (*sacloud.ServerPlan, error) {
	var cond *sacloud.FindCondition
	if param != nil {
		cond = param.findCondition()
	}

	searched, err := finder.Find(ctx, zone, cond)
	if err != nil {
		return nil, err
	}
	if searched.Count == 0 || len(searched.ServerPlans) == 0 {
		return nil, errors.New("server plan not found")
	}
	return searched.ServerPlans[0], nil
}

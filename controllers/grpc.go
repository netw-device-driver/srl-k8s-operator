/*
	Copyright 2021 Wim Henderickx.

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

package controllers

import (
	"context"

	"github.com/netw-device-driver/netwdevpb"
	"github.com/srl-wim/srl-k8s-operator/pkg/grpcc"
)

func getCachStatus(ctx context.Context, target, resource *string, level int32) (*netwdevpb.CacheStatusReply, error) {
	c := &grpcc.Client{
		Insecure:   true,
		SkipVerify: true,
		Target:     *target,
	}
	req := &netwdevpb.CacheStatusRequest{
		Resource: *resource,
		Level:    level,
	}
	return c.GetCacheStatus(ctx, req)
}

func updateCache(ctx context.Context, target *string, req *netwdevpb.CacheUpdateRequest) (*netwdevpb.CacheUpdateReply, error) {
	c := &grpcc.Client{
		Insecure:   true,
		SkipVerify: true,
		Target:     *target,
	}

	return c.UpdateCache(ctx, req)
}

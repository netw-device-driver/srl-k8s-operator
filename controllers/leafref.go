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
	"encoding/json"

	srlinuxv1alpha1 "github.com/srl-wim/srl-k8s-operator/api/v1alpha1"
)

// getRemoteleafRefResource
func getRemoteleafRefResource(targetName *string, remoteleafRef *srlinuxv1alpha1.RemoteLeafRefInfo) (*string, error) {
	resource := new(string)
	p := new(string)
	for res, dps := range SrlSharedInfo[*targetName] {
		var x1 interface{}
		json.Unmarshal([]byte(dps), &x1)
		f, dp := matchDeletePath(x1, remoteleafRef.RemoteLeafRef)
		if f {
			if len(*dp) > len(*p) {
				resource = stringPtr(res)
				*p = *dp
			}
		}
	}
	return resource, nil
}

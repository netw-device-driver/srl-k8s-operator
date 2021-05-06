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

import "time"

const (
	targetNotFoundRetryDelay  = time.Second * 60
	validationErrorRetyrDelay = time.Second * 60
)

// StringInList returns a boolean indicating whether strToSearch is a
// member of the string slice passed as the first argument.
func StringInList(list []string, strToSearch string) bool {
	for _, item := range list {
		if item == strToSearch {
			return true
		}
	}
	return false
}

// FilterStringFromList produces a new string slice that does not
// include the strToFilter argument.
func FilterStringFromList(list []string, strToFilter string) (newList []string) {
	for _, item := range list {
		if item != strToFilter {
			newList = append(newList, item)
		}
	}
	return
}

func removeString(slice []string, s string) (result []string) {
	for _, v := range slice {
		if v != s {
			result = append(result, s)
		}
	}
	return result
}

func stringPtr(s string) *string { return &s }
func intPtr(i int) *int          { return &i }

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
	"reflect"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	targetNotFoundRetryDelay = time.Second * 60
	//validationErrorRetyrDelay = time.Second * 60
	internalLeafRefvalidationErrorretryDelay = time.Second * 20
	externalLeafRefvalidationErrorretryDelay = time.Second * 10
	parentDependencyRetyrDelay               = time.Second * 5
	deleteDependencyRetryDelay               = time.Second * 10
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

func stringPtr(s string) *string              { return &s }
func intPtr(i int) *int                       { return &i }
func stringSlicePtr(s []string) *[]string     { return &s }
func interfacePtr(i interface{}) *interface{} { return &i }

func getHierarchicalElements(p string) (ekv []ElementKeyValue) {
	skipElement := false

	s1 := strings.Split(p, "/")
	for i := range s1 {
		if i > 0 && !skipElement {
			if strings.Contains(s1[i], "[") {
				s2 := strings.Split(s1[i], "[")
				s3 := strings.Split(s2[1], "=")
				var v string
				if strings.Contains(s3[1], "]") {
					v = strings.Trim(s3[1], "]")
				} else {
					v = s3[1] + "/" + strings.Trim(s1[i+1], "]")
					skipElement = true
				}
				e := ElementKeyValue{
					Element:  s2[0],
					KeyName:  s3[0],
					KeyValue: v,
				}
				ekv = append(ekv, e)
			} else {
				e := ElementKeyValue{
					Element:  s1[i],
					KeyName:  "",
					KeyValue: "",
				}
				ekv = append(ekv, e)
			}
		} else {
			skipElement = false
		}
	}
	return ekv
}

func matchDeletePath(x1 interface{}, remoteLeafRefPath *string) (bool, *string) {
	log.Debugf("matchDeletePath: X1 %v", x1)
	switch x := x1.(type) {
	case map[string]interface{}:
		for k, x2 := range x {
			log.Debugf("matchDeletePath map[string]interface{} Key %s", k)
			if k == "deletePaths" {
				return matchDeletePath(x2, remoteLeafRefPath)
			}
		}
	case []interface{}:
		for _, v := range x {
			switch x2 := v.(type) {
			case string:
				log.Debugf("matchDeletePath []interface{} X2 %v", string(x2))
				if strings.Contains(*remoteLeafRefPath, string(x2)) {
					return true, stringPtr(x2)
				}
			}
		}
	default:
		log.Debugf("matchDeletePath Default type %v", reflect.TypeOf(x))
	}
	return false, nil
}

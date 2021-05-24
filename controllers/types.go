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

type ElementWithLeafRef struct {
	REkvl                          []ElementKeyValue
	RemoteLeafRefs                 []string
	RelativePath2ObjectWithLeafRef string
	AbsolutePath2LeafRef           string
	RelativePath2LeafRef           string
	ElementName                    string
	KeyName                        string
	LocalLeafRefValues             []string
	Exists                         bool
	DependencyCheckSuccess         bool
	LeafRefValues                  []string
}

// ElementKeyValue struct
type ElementKeyValue struct {
	Element  string
	KeyName  string
	KeyValue interface{}
}

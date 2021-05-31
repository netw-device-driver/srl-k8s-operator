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

	srlinuxv1alpha1 "github.com/srl-wim/srl-k8s-operator/api/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type ElementWithLeafRef struct {
	REkvl []ElementKeyValue

	LocalResolvedLeafRefInfo map[string]*srlinuxv1alpha1.RemoteLeafRefInfo

	//RelativePath2ObjectWithLeafRef string
	//AbsolutePath2LeafRef           string
	//RelativePath2LeafRef           string
	//ElementName                    string
	//KeyName                        string
	//DependencyCheckSuccess bool
	//LeafRefValues                  []string
	//Exists                         bool
}

//type RemoteLeafRefInfo struct {
//	RemoteLeafRef          string
//	DependencyCheckSuccess bool
//}

// ElementKeyValue struct
type ElementKeyValue struct {
	Element  string      `json:"element"`
	KeyName  string      `json:"keyName"`
	KeyValue interface{} `json:"keyValue"`
}

/*
type DeletePaths struct {
	ObjectName  *string            `json:"objectName"`
	DeletePaths *[]ElementKeyValue `json:"deletePaths"`
}
*/

type DeletePaths struct {
	DeletePaths *[]string `json:"deletePaths"`
}

type LeafRefResource struct {
	ctx                       context.Context
	client                    client.Client
	nameSpace                 string
	resourceName              string
	resourceObjectName        string
	leafRefResourceName       string
	leafRefResourceObjectName string
	target                    string
}

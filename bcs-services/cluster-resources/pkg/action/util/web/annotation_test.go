/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2022 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * 	http://opensource.org/licenses/MIT
 *
 * Unless required by applicable law or agreed to in writing, software distributed under,
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package web

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWebAnnotations(t *testing.T) {
	exceptAnnos := Annotations{
		Perm: Perm{
			Page: map[ObjName]PermDetail{
				"createBtn": {false, "no perm", ""},
			},
			Items: map[ResUID]ObjPerm{
				"a8ec4e03": {
					"editBtn": {false, "can't edit", ""},
				},
				"ed8250cc": {
					"editBtn": {true, "", ""},
				},
			},
		},
		FeatureFlag: map[FeatureFlagKey]bool{"pvc": false, "hpa": true},
	}
	actualAnnos := NewAnnos(
		NewPagePerm("createBtn", PermDetail{false, "no perm", ""}),
		NewItemPerm("a8ec4e03", "editBtn", PermDetail{false, "can't edit", ""}),
		NewItemPerm("ed8250cc", "editBtn", PermDetail{true, "", ""}),
		NewFeatureFlag("pvc", false),
		NewFeatureFlag("hpa", true),
	)
	assert.Equal(t, exceptAnnos, actualAnnos)

	// 测试转换 pb.Struct 类型
	_, err := actualAnnos.ToPbStruct()
	assert.Nil(t, err)
}

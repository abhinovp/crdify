// Copyright 2025 The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package property

import (
	"testing"

	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	internaltesting "sigs.k8s.io/crdify/pkg/validations/internal/testing"
)

func TestEnum(t *testing.T) {
	testcases := []internaltesting.Testcase[apiextensionsv1.JSONSchemaProps]{
		{
			Name: "no diff, not flagged",
			Old: &apiextensionsv1.JSONSchemaProps{
				Enum: []apiextensionsv1.JSON{
					{
						Raw: []byte("foo"),
					},
				},
			},
			New: &apiextensionsv1.JSONSchemaProps{
				Enum: []apiextensionsv1.JSON{
					{
						Raw: []byte("foo"),
					},
				},
			},
			Flagged:              false,
			ComparableValidation: &Enum{},
		},
		{
			Name: "new enum constraint, flagged",
			Old: &apiextensionsv1.JSONSchemaProps{
				Enum: []apiextensionsv1.JSON{},
			},
			New: &apiextensionsv1.JSONSchemaProps{
				Enum: []apiextensionsv1.JSON{
					{
						Raw: []byte("foo"),
					},
				},
			},
			Flagged:              true,
			ComparableValidation: &Enum{},
		},
		{
			Name: "new allowed enum value added, addition policy not set, flagged",
			Old: &apiextensionsv1.JSONSchemaProps{
				Enum: []apiextensionsv1.JSON{
					{
						Raw: []byte("foo"),
					},
				},
			},
			New: &apiextensionsv1.JSONSchemaProps{
				Enum: []apiextensionsv1.JSON{
					{
						Raw: []byte("foo"),
					},
					{
						Raw: []byte("bar"),
					},
				},
			},
			Flagged:              true,
			ComparableValidation: &Enum{},
		},
		{
			Name: "new allowed enum value added, addition policy set to Disallow, flagged",
			Old: &apiextensionsv1.JSONSchemaProps{
				Enum: []apiextensionsv1.JSON{
					{
						Raw: []byte("foo"),
					},
				},
			},
			New: &apiextensionsv1.JSONSchemaProps{
				Enum: []apiextensionsv1.JSON{
					{
						Raw: []byte("foo"),
					},
					{
						Raw: []byte("bar"),
					},
				},
			},
			Flagged: true,
			ComparableValidation: &Enum{
				EnumConfig: EnumConfig{
					AdditionPolicy: AdditionPolicyDisallow,
				},
			},
		},
		{
			Name: "new allowed enum value added, addition policy set to Allow, not flagged",
			Old: &apiextensionsv1.JSONSchemaProps{
				Enum: []apiextensionsv1.JSON{
					{
						Raw: []byte("foo"),
					},
				},
			},
			New: &apiextensionsv1.JSONSchemaProps{
				Enum: []apiextensionsv1.JSON{
					{
						Raw: []byte("foo"),
					},
					{
						Raw: []byte("bar"),
					},
				},
			},
			Flagged: false,
			ComparableValidation: &Enum{
				EnumConfig: EnumConfig{
					AdditionPolicy: AdditionPolicyAllow,
				},
			},
		},
		{
			Name: "removed enum value, flagged",
			Old: &apiextensionsv1.JSONSchemaProps{
				Enum: []apiextensionsv1.JSON{
					{
						Raw: []byte("foo"),
					},
					{
						Raw: []byte("bar"),
					},
				},
			},
			New: &apiextensionsv1.JSONSchemaProps{
				Enum: []apiextensionsv1.JSON{
					{
						Raw: []byte("bar"),
					},
				},
			},
			Flagged:              true,
			ComparableValidation: &Enum{},
		},
		{
			Name: "different field changed, not flagged",
			Old: &apiextensionsv1.JSONSchemaProps{
				ID: "foo",
			},
			New: &apiextensionsv1.JSONSchemaProps{
				ID: "bar",
			},
			Flagged:              false,
			ComparableValidation: &Enum{},
		},
		{
			Name: "different field changed with enum, not flagged",
			Old: &apiextensionsv1.JSONSchemaProps{
				ID: "foo",
				Enum: []apiextensionsv1.JSON{
					{
						Raw: []byte("foo"),
					},
				},
			},
			New: &apiextensionsv1.JSONSchemaProps{
				ID: "bar",
				Enum: []apiextensionsv1.JSON{
					{
						Raw: []byte("foo"),
					},
				},
			},
			Flagged:              false,
			ComparableValidation: &Enum{},
		},
	}

	internaltesting.RunTestcases(t, testcases...)
}

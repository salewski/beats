// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package kubernetes

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "kubernetes", asset.ModuleFieldsPri, AssetKubernetes); err != nil {
		panic(err)
	}
}

// AssetKubernetes returns asset data.
// This is the base64 encoded gzipped contents of ../metricbeat/module/kubernetes.
func AssetKubernetes() string {
	return "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"
}

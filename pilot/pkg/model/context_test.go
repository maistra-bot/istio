// Copyright 2017 Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"istio.io/istio/pilot/pkg/model"
	"istio.io/istio/pilot/pkg/serviceregistry/memory"
)

func TestServiceNode(t *testing.T) {
	nodes := []struct {
		in  *model.Proxy
		out string
	}{
		{
			in:  &memory.HelloProxyV0,
			out: "sidecar~10.1.1.0~v0.default~default.svc.cluster.local",
		},
		{
			in: &model.Proxy{
				Type:        model.Router,
				ID:          "random",
				IPAddresses: []string{"10.3.3.3"},
				DNSDomain:   "local",
			},
			out: "router~10.3.3.3~random~local",
		},
		{
			in: &model.Proxy{
				Type:        model.SidecarProxy,
				ID:          "random",
				IPAddresses: []string{"10.3.3.3", "10.4.4.4", "10.5.5.5", "10.6.6.6"},
				DNSDomain:   "local",
				Metadata: map[string]string{
					"INSTANCE_IPS": "10.3.3.3,10.4.4.4,10.5.5.5,10.6.6.6",
				},
			},
			out: "sidecar~10.3.3.3~random~local",
		},
	}

	for _, node := range nodes {
		out := node.in.ServiceNode()
		if out != node.out {
			t.Errorf("%#v.ServiceNode() => Got %s, want %s", node.in, out, node.out)
		}
		in, err := model.ParseServiceNodeWithMetadata(node.out, node.in.Metadata)

		if err != nil {
			t.Errorf("ParseServiceNode(%q) => Got error %v", node.out, err)
		}
		if !reflect.DeepEqual(in, node.in) {
			t.Errorf("ParseServiceNode(%q) => Got %#v, want %#v", node.out, in, node.in)
		}
	}
}

func TestParsePort(t *testing.T) {
	if port := model.ParsePort("localhost:3000"); port != 3000 {
		t.Errorf("ParsePort(localhost:3000) => Got %d, want 3000", port)
	}
	if port := model.ParsePort("localhost"); port != 0 {
		t.Errorf("ParsePort(localhost) => Got %d, want 0", port)
	}
}

func TestGetOrDefaultFromMap(t *testing.T) {
	meta := map[string]string{"key1": "key1ValueFromMap"}
	assert.Equal(t, "key1ValueFromMap", model.GetOrDefaultFromMap(meta, "key1", "unexpected"))
	assert.Equal(t, "expectedDefaultKey2Value", model.GetOrDefaultFromMap(meta, "key2", "expectedDefaultKey2Value"))
	assert.Equal(t, "expectedDefaultFromNilMap", model.GetOrDefaultFromMap(nil, "key", "expectedDefaultFromNilMap"))
}

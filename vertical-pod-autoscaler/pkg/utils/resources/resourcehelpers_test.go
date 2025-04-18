/*
Copyright 2025 The Kubernetes Authors.

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

package resourcehelpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"

	"k8s.io/autoscaler/vertical-pod-autoscaler/pkg/utils/test"
)

func TestContainerRequestsAndLimits(t *testing.T) {
	testCases := []struct {
		desc          string
		containerName string
		pod           *apiv1.Pod
		wantRequests  apiv1.ResourceList
		wantLimits    apiv1.ResourceList
	}{
		{
			desc:          "Prefer resource requests from container status",
			containerName: "container",
			pod: test.Pod().AddContainer(
				test.Container().WithName("container").
					WithCPURequest(resource.MustParse("1")).
					WithMemRequest(resource.MustParse("10Mi")).
					WithCPULimit(resource.MustParse("2")).
					WithMemLimit(resource.MustParse("20Mi")).Get()).
				AddContainerStatus(
					test.ContainerStatus().WithName("container").
						WithCPURequest(resource.MustParse("3")).
						WithMemRequest(resource.MustParse("30Mi")).
						WithCPULimit(resource.MustParse("4")).
						WithMemLimit(resource.MustParse("40Mi")).Get()).Get(),
			wantRequests: apiv1.ResourceList{
				apiv1.ResourceCPU:    resource.MustParse("3"),
				apiv1.ResourceMemory: resource.MustParse("30Mi"),
			},
			wantLimits: apiv1.ResourceList{
				apiv1.ResourceCPU:    resource.MustParse("4"),
				apiv1.ResourceMemory: resource.MustParse("40Mi"),
			},
		},
		{
			desc:          "No container status, get resources from pod spec",
			containerName: "container",
			pod: test.Pod().AddContainer(
				test.Container().WithName("container").
					WithCPURequest(resource.MustParse("1")).
					WithMemRequest(resource.MustParse("10Mi")).
					WithCPULimit(resource.MustParse("2")).
					WithMemLimit(resource.MustParse("20Mi")).Get()).Get(),
			wantRequests: apiv1.ResourceList{
				apiv1.ResourceCPU:    resource.MustParse("1"),
				apiv1.ResourceMemory: resource.MustParse("10Mi"),
			},
			wantLimits: apiv1.ResourceList{
				apiv1.ResourceCPU:    resource.MustParse("2"),
				apiv1.ResourceMemory: resource.MustParse("20Mi"),
			},
		},
		{
			desc:          "Only containerStatus, get resources from containerStatus",
			containerName: "container",
			pod: test.Pod().AddContainerStatus(
				test.ContainerStatus().WithName("container").
					WithCPURequest(resource.MustParse("0")).
					WithMemRequest(resource.MustParse("30Mi")).
					WithCPULimit(resource.MustParse("4")).
					WithMemLimit(resource.MustParse("40Mi")).Get()).Get(),
			wantRequests: apiv1.ResourceList{
				apiv1.ResourceCPU:    resource.MustParse("0"),
				apiv1.ResourceMemory: resource.MustParse("30Mi"),
			},
			wantLimits: apiv1.ResourceList{
				apiv1.ResourceCPU:    resource.MustParse("4"),
				apiv1.ResourceMemory: resource.MustParse("40Mi"),
			},
		},
		{
			desc:          "Inexistent container",
			containerName: "inexistent-container",
			pod: test.Pod().AddContainer(
				test.Container().WithName("container").
					WithCPURequest(resource.MustParse("1")).
					WithMemRequest(resource.MustParse("10Mi")).
					WithCPULimit(resource.MustParse("2")).
					WithMemLimit(resource.MustParse("20Mi")).Get()).Get(),
			wantRequests: nil,
			wantLimits:   nil,
		},
		{
			desc:          "Container with no requests or limits",
			containerName: "inexistent-container",
			pod:           test.Pod().AddContainer(test.Container().WithName("container").Get()).Get(),
			wantRequests:  nil,
			wantLimits:    nil,
		},
		{
			desc:          "2 containers",
			containerName: "container-1",
			pod: test.Pod().AddContainer(
				test.Container().WithName("container-1").
					WithCPURequest(resource.MustParse("1")).
					WithMemRequest(resource.MustParse("10Mi")).
					WithCPULimit(resource.MustParse("2")).
					WithMemLimit(resource.MustParse("20Mi")).Get()).
				AddContainerStatus(
					test.ContainerStatus().WithName("container-1").
						WithCPURequest(resource.MustParse("3")).
						WithMemRequest(resource.MustParse("30Mi")).
						WithCPULimit(resource.MustParse("4")).
						WithMemLimit(resource.MustParse("40Mi")).Get()).
				AddContainer(
					test.Container().WithName("container-2").
						WithCPURequest(resource.MustParse("5")).
						WithMemRequest(resource.MustParse("5Mi")).
						WithCPULimit(resource.MustParse("5")).
						WithMemLimit(resource.MustParse("5Mi")).Get()).
				AddContainerStatus(
					test.ContainerStatus().WithName("container-2").
						WithCPURequest(resource.MustParse("5")).
						WithMemRequest(resource.MustParse("5Mi")).
						WithCPULimit(resource.MustParse("5")).
						WithMemLimit(resource.MustParse("5Mi")).Get()).
				Get(),
			wantRequests: apiv1.ResourceList{
				apiv1.ResourceCPU:    resource.MustParse("3"),
				apiv1.ResourceMemory: resource.MustParse("30Mi"),
			},
			wantLimits: apiv1.ResourceList{
				apiv1.ResourceCPU:    resource.MustParse("4"),
				apiv1.ResourceMemory: resource.MustParse("40Mi"),
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			gotRequests, gotLimits := ContainerRequestsAndLimits(tc.containerName, tc.pod)
			assert.Equal(t, tc.wantRequests, gotRequests, "requests don't match")
			assert.Equal(t, tc.wantLimits, gotLimits, "limits don't match")
		})
	}
}

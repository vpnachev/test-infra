// Copyright 2019 Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package framework

import (
	"context"
	"k8s.io/apimachinery/pkg/runtime"
	"time"

	argov1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
	"github.com/gardener/gardener/pkg/client/kubernetes"
	tmv1beta1 "github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1"
	"github.com/gardener/test-infra/pkg/testmachinery"
	"github.com/gardener/test-infra/test/utils"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
)

// Client returns the kubernetes client of the current test cluster
func (o *Operation) Client() kubernetes.Interface {
	return o.tmClient
}

// TestMachineryNamespace returns the current namespace where the testmachinery components are running.
func (o *Operation) TestMachineryNamespace() string {
	return o.config.TmNamespace
}

// TestNamespace returns the name of the current test namespace
func (o *Operation) TestNamespace() string {
	return o.config.Namespace
}

// Commit returns the current commit sha of the test-infra repo
func (o *Operation) Commit() string {
	return o.config.CommitSha
}

func (o *Operation) S3Endpoint() string {
	return o.config.S3Endpoint
}

// Log returns the test logger
func (o *Operation) Log() logr.Logger {
	return o.log
}

// IsLocal indicates if the test is running against a local testmachinery controller
func (o *Operation) IsLocal() bool {
	return o.config.Local
}

// WaitForClusterReadiness waits until all Test Machinery components are ready
func (o *Operation) WaitForClusterReadiness(maxWaitTime time.Duration) error {
	if o.IsLocal() {
		return nil
	}
	return utils.WaitForClusterReadiness(o.Log(), o.tmClient, o.config.TmNamespace, maxWaitTime)
}

// WaitForMinioServiceReadiness waits until the minio service in the testcluster is ready
func (o *Operation) WaitForMinioServiceReadiness(maxWaitTime time.Duration) (*testmachinery.S3Config, error) {
	if o.IsLocal() {
		return nil, nil
	}
	if len(o.config.S3Endpoint) == 0 {
		return nil, errors.New("no s3 endpoint is defined")
	}
	return utils.WaitForMinioService(o.Client(), o.config.S3Endpoint, o.config.TmNamespace, maxWaitTime)
}

func (o *Operation) RunTestrunUntilCompleted(ctx context.Context, tr *tmv1beta1.Testrun, phase argov1.NodePhase, timeout time.Duration) (*tmv1beta1.Testrun, *argov1.Workflow, error) {
	return utils.RunTestrunUntilCompleted(ctx, o.Log().WithValues("namespace", o.TestNamespace()), o.Client(), tr, phase, timeout)
}

func (o *Operation) RunTestrun(ctx context.Context, tr *tmv1beta1.Testrun, phase argov1.NodePhase, timeout time.Duration, watchFunc utils.WatchFunc) (*tmv1beta1.Testrun, *argov1.Workflow, error) {
	return utils.RunTestrun(ctx, o.Log().WithValues("namespace", o.TestNamespace()), o.Client(), tr, phase, timeout, watchFunc)
}

// AppendObject adds a kubernetes objects to the start of the state's objects.
// These objects are meant to be cleaned up after the test has run.
func (s *OperationState) AppendObject(obj runtime.Object) {
	if s.Objects == nil {
		s.Objects = make([]runtime.Object, 0)
	}
	s.Objects = append([]runtime.Object{obj}, s.Objects...)
}

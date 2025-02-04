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

package common

// Annotations
const (
	PurposeTestrunAnnotation = "testmachinery.sapcloud.io/purpose"
	ResumeTestrunAnnotation  = "testmachinery.sapcloud.io/resume"

	AnnotationSystemStep = "testmachinery.sapcloud.io/system-step"

	// images
	DockerImageGardenerApiServer = "eu.gcr.io/gardener-project/gardener/apiserver"

	// Repositories
	TestInfraRepo          = "https://github.com/gardener/test-infra.git"
	GardenSetupRepo        = "https://github.com/gardener/garden-setup.git"
	GardenerRepo           = "https://github.com/gardener/gardener.git"
	GardenerExtensionsRepo = "https://github.com/gardener/gardener-extensions.git"

	PatternLatest = "latest"
)

var (
	// Default timeout of 4 hours to wait before resuming the testrun
	DefaultPauseTimeout = 14400
)

/**
 * Copyright (C) 2015 Red Hat, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *         http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package main

import (
	"fmt"
	"github.com/fabric8io/kubernetes-client/generator/pkg/schemagen"
	"github.com/kudobuilder/kudo/pkg/apis/kudo/v1beta1"
	"github.com/kudobuilder/kudo/pkg/kudoctl/packages"
	"k8s.io/apimachinery/pkg/runtime"
	"reflect"
)

func main() {

	// the CRD List types for which the model should be generated
	// no other types need to be defined as they are auto discovered
	crdLists := map[reflect.Type]schemagen.CrdScope{
		// kudo
		reflect.TypeOf(v1beta1.AggregatedStatus{}): schemagen.Namespaced,
		reflect.TypeOf(v1beta1.DummyTaskSpec{}):    schemagen.Namespaced,
		// reflect.TypeOf(v1beta1.ExecutionStatus):       schemagen.Namespaced,
		reflect.TypeOf(v1beta1.Instance{}):              schemagen.Namespaced,
		reflect.TypeOf(v1beta1.InstanceError{}):         schemagen.Namespaced,
		reflect.TypeOf(v1beta1.InstanceList{}):          schemagen.Namespaced,
		reflect.TypeOf(v1beta1.InstanceSpec{}):          schemagen.Namespaced,
		reflect.TypeOf(v1beta1.Maintainer{}):            schemagen.Namespaced,
		reflect.TypeOf(v1beta1.InstanceStatus{}):        schemagen.Namespaced,
		reflect.TypeOf(v1beta1.Operator{}):              schemagen.Namespaced,
		reflect.TypeOf(v1beta1.OperatorDependency{}):    schemagen.Namespaced,
		reflect.TypeOf(v1beta1.OperatorList{}):          schemagen.Namespaced,
		reflect.TypeOf(v1beta1.OperatorSpec{}):          schemagen.Namespaced,
		reflect.TypeOf(v1beta1.OperatorStatus{}):        schemagen.Namespaced,
		reflect.TypeOf(v1beta1.OperatorVersion{}):       schemagen.Namespaced,
		reflect.TypeOf(v1beta1.OperatorVersionList{}):   schemagen.Namespaced,
		reflect.TypeOf(v1beta1.OperatorVersionSpec{}):   schemagen.Namespaced,
		reflect.TypeOf(v1beta1.OperatorVersionStatus{}): schemagen.Namespaced,
		// reflect.TypeOf(v1beta1.Ordering):       schemagen.Namespaced,
		reflect.TypeOf(v1beta1.Parameter{}): schemagen.Namespaced,
		// reflect.TypeOf(v1beta1.ParameterType):       schemagen.Namespaced,
		reflect.TypeOf(v1beta1.Phase{}):            schemagen.Namespaced,
		reflect.TypeOf(v1beta1.PhaseStatus{}):      schemagen.Namespaced,
		reflect.TypeOf(v1beta1.PipeSpec{}):         schemagen.Namespaced,
		reflect.TypeOf(v1beta1.PipeTaskSpec{}):     schemagen.Namespaced,
		reflect.TypeOf(v1beta1.Plan{}):             schemagen.Namespaced,
		reflect.TypeOf(v1beta1.PlanExecution{}):    schemagen.Namespaced,
		reflect.TypeOf(v1beta1.PlanStatus{}):       schemagen.Namespaced,
		reflect.TypeOf(v1beta1.ResourceTaskSpec{}): schemagen.Namespaced,
		reflect.TypeOf(v1beta1.Step{}):             schemagen.Namespaced,
		reflect.TypeOf(v1beta1.StepStatus{}):       schemagen.Namespaced,
		reflect.TypeOf(v1beta1.Task{}):             schemagen.Namespaced,
		reflect.TypeOf(v1beta1.TaskSpec{}):         schemagen.Namespaced,
		reflect.TypeOf(v1beta1.ToggleTaskSpec{}):   schemagen.Namespaced,

		// kudo package
		reflect.TypeOf(packages.Files{}):        schemagen.Namespaced,
		reflect.TypeOf(packages.Package{}):      schemagen.Namespaced,
		reflect.TypeOf(packages.Resources{}):    schemagen.Namespaced,
		reflect.TypeOf(packages.OperatorFile{}): schemagen.Namespaced,
		reflect.TypeOf(packages.ParamsFile{}):   schemagen.Namespaced,
		reflect.TypeOf(packages.Parameters{}):   schemagen.Namespaced,
		reflect.TypeOf(packages.Parameter{}):    schemagen.Namespaced,
		reflect.TypeOf(packages.Templates{}):    schemagen.Namespaced,
	}

	// constraints and patterns for fields
	constraints := map[reflect.Type]map[string]*schemagen.Constraint{}

	// types that are manually defined in the model
	providedTypes := []schemagen.ProvidedType{}

	// go packages that are provided and where no generation is required and their corresponding java package
	providedPackages := map[string]string{
		// external
		"k8s.io/api/core/v1":                   "io.fabric8.kubernetes.api.model",

		"k8s.io/apimachinery/pkg/api/resource": "io.fabric8.kubernetes.api.model",
		"k8s.io/apimachinery/pkg/apis/meta/v1": "io.fabric8.kubernetes.api.model",
		"k8s.io/apimachinery/pkg/util": "io.fabric8.kubernetes.api.model",
	}

	// mapping of go packages of this module to the resulting java package
	// optional ApiGroup and ApiVersion for the go package (which is added to the generated java class)
	packageMapping := map[string]schemagen.PackageInformation{
		"github.com/kudobuilder/kudo/pkg/apis/kudo/v1beta1": {JavaPackage: "io.fabric8.kudo.v1beta1", ApiGroup: "kudo.dev", ApiVersion: "v1beta1"},
		"github.com/kudobuilder/kudo/pkg/kudoctl/packages":  {JavaPackage: "io.fabric8.kudo.packages", ApiGroup: "kudo.dev", ApiVersion: "v1beta1"},
	}

	// converts all packages starting with <key> to a java package using an automated scheme:
	//  - replace <key> with <value> aka "package prefix"
	//  - replace '/' with '.' for a valid java package name
	mappingSchema := map[string]string{
		"kudo.dev": "io.fabric8.kudo.v1beta1",
	}

	// overwriting some times
	manualTypeMap := map[reflect.Type]string{
		// reflect.TypeOf(apis.URL{}):             "java.lang.String",
		// reflect.TypeOf(apis.VolatileTime{}):    "java.lang.String",
		reflect.TypeOf(runtime.RawExtension{}): "io.fabric8.kubernetes.api.model.HasMetadata",
	}

	json := schemagen.GenerateSchema("http://fabric8.io/kudo/KUDOSchema#", crdLists, providedPackages, manualTypeMap, packageMapping, mappingSchema, providedTypes, constraints)

	fmt.Println(json)
}

/*
Copyright 2017 The Kubernetes Authors.

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

// Note: the example only works with the code within the same release/branch.
package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	machinelearningv1alpha2 "github.com/seldonio/seldon-core/operator/api/v1alpha2"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
)

func main() {
	var fpath *string
	fpath = flag.String("filepath", "", "absolute path to the deployment file")
	flag.Parse()

	// create scheme
	scheme := runtime.NewScheme()
	_ = machinelearningv1alpha2.AddToScheme(scheme)
	k8sManager, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:         scheme,
		LeaderElection: false,
	})

	client := k8sManager.GetClient()

	dat, err := ioutil.ReadFile(*fpath)
	if err != nil {
		panic(err)
	}
	seldonDeployment := &machinelearningv1alpha2.SeldonDeployment{}
	err = json.NewDecoder(bytes.NewBuffer(dat)).Decode(seldonDeployment)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v \n", seldonDeployment)

	// Create Deployment
	fmt.Println("Creating deployment...")
	//result, err := deploymentsClient.Create(deployment)
	seldonDeployment.Default()
	err = client.Create(context.Background(), seldonDeployment)
	if err != nil {
		panic(err)
	}
	fmt.Print("I created a seldon deployment \n")
	prompt()

}

func prompt() {
	fmt.Printf("-> Press Return key to continue.")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		break
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println()
}

func int32Ptr(i int32) *int32 { return &i }

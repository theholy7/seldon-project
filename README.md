# Challenge

To demonstrate your knowledge of Cloud Native Technologies and general programming we would like you to complete the following task.

Task:
  * Create a standalone program in Go which takes in a Seldon Core Custom Resource and creates it over the Kubernetes API
  * Watch the created resource to wait for it to become available.
  * When it is available delete the Custom Resource.

For this task you will need:
  * A running Kubernetes cluster (e.g. Minikube)
  * Seldon Core running on it
  * A Seldon Custom Resource
    * An example SeldonDeploment can be found at [here](https://github.com/SeldonIO/seldon-core/).

Please return your submission by providing the link to a public Github repository when your solution is complete. You will be evaluated on functionality achieved, clarity of code and ease of use for us to test your solution.
Please don't hesitate to come back to us if you have any questions.

# Solution

## Setup

Make sure you have Helm and Minikube installed and running. Make sure that `kubectl` is correctly configured to apply changes to Minikube.

To deploy Seldon Core on Minikube use the script in `./scripts/start_seldon_core.sh`.

To remove Seldon Core you can run `./scripts/remove_seldon_core.sh`.

Install ambassador with `helm install ambassador stable/ambassador`.

Run `go run main.go --filepath seldon_deployment.json`. This will deploy an example mock model to Minikube.

Run on a new terminal `minikube tunnel` to get an external IP for ambassador. You can check the IP by doing `kubectl get svc --all-namespaces`

Run `curl -XPOST -k http://<Ambassador-IP>/seldon/default/seldon-model/api/v0.1/predictions -v`. The response will be an error, but we sucessfully reached the mock model deployed.

## Useful commands

Use `kubectl get sdep` to check if a SeldonCustomDeployment exists.



## Version used

```
helm version
version.BuildInfo{Version:"v3.0.0", GitCommit:"e29ce2a54e96cd02ccfce88bee4f58bb6e2a28b6", GitTreeState:"clean", GoVersion:"go1.13.4"}

minikube version: v1.5.2
commit: 792dbf92a1de583fcee76f8791cff12e0c9440ad

hyperkit -v
hyperkit: 0.20190802
```

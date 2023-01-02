# kubectl-categories
Kubectl plugin to list resources categories

This command displays the names of categories you can use with the following command, and the list of resources belonging to each category:

```
$ kubectl get <category>
```

## Example

```console
$ go build

$ ./kubectl-categories 
all:
  pods (v1)
  replicationcontrollers (v1)
  services (v1)
  daemonsets (apps/v1)
  deployments (apps/v1)
  replicasets (apps/v1)
  statefulsets (apps/v1)
  horizontalpodautoscalers (autoscaling/v2)
  cronjobs (batch/v1)
  jobs (batch/v1)
api-extensions:
  mutatingwebhookconfigurations (admissionregistration.k8s.io/v1)
  validatingwebhookconfigurations (admissionregistration.k8s.io/v1)
  customresourcedefinitions (apiextensions.k8s.io/v1)
  apiservices (apiregistration.k8s.io/v1)

$ kubectl get api-extensions
[...]
```

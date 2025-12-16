# Bifrost Operator — Zero-Effort Kubernetes Log Aggregation

**_An open source project from [Peek8.io](https://peek8.io/)._** 

**Bifrost** is a Kubernetes Operator that automatically deploys a full observability stack for logs — **Grafana Alloy → Loki → Grafana** — and wires everything up to collect logs from selected namespaces with no manual setup.

Define which namespaces you want to monitor using the `LogSpace` CRD and Bifrost provisions **all required components, dashboards, datasources, storage, RBAC and configurations** in minutes.


## 🌈 Why the name “Bifrost”?

The name `Bifrost` sounds familiear to you ? Probably, you have heard it in `Thor Movie Series` :).   In [**Norse Mythology**](https://en.wikipedia.org/wiki/Norse_mythology)(from where Loki came from :) ), *Bifrost* is the magical rainbow bridge that connects **Midgard (the Earth)** to **Asgard (the world of the gods)**.

Similarly, the `Bifrost Operator` forms a **bridge between your Kubernetes applications and their observability**, connecting logs across namespaces to a central dashboard you can rely on.

## Acknowledgement
The `bifrost operator` uses statemachine. The tool [vectorSigma](https://github.com/mhersson/vectorsigma) has been used to generate the state machine in GO from UML. To know more about how to use it in operator, you can have look at the [vectorsigma k8s operator guide](https://github.com/mhersson/vectorsigma/blob/main/docs/k8s-operator-guide.md).


## Motivation
Setting up a complete, production-grade log aggregation stack in Kubernetes is not straightforward.

Manually deploying and configuring **Grafana Alloy, Loki, persistent storage, Grafana, datasources, dashboards, label filters, namespace scoping, service accounts, and RBAC permissions** can easily take days of trial and error — even for experienced DevOps engineers.

It involves:

- Understanding and writing Alloy pipelines
- Scaling and configuring Loki stateful storage
- Deploying Grafana and patching datasource JSON
- Building dashboards manually and maintaining them as services change
- Managing service accounts, RBAC, ports, selectors and labels across namespaces

**Bifrost reduces that to a single YAML file and a few minutes.**

With just(see more details at `Quick Start` section):

```yaml
apiVersion: bifrost.peek8.io/v1alpha1
kind: LogSpace
metadata:
  name: sample
spec:
  targetNamespaces:
    - dev
    - staging
```

the operator will:

- Deploy and configure Grafana Alloy (DaemonSet)
- Deploy and configure Loki (StatefulSet)
- Deploy and configure Grafana (Deployment)
- Configure datasources and provision a ready-to-use dashboard
- Wire everything automatically without user intervention

No deep observability expertise required. No manual YAML surgery. No multi-day setup.


## Features
| Category         | Capability                                                  |
| ---------------- | ----------------------------------------------------------- |
| Log Collection   | Collect logs from multiple namespaces using Grafana Alloy   |
| Storage          | Fully provisioned Loki cluster (StatefulSet + PVCs)         |
| Visualization    | Auto-generated Grafana dashboards and log explorer          |
| Config           | Datasource + dashboard provisioning via the operator        |
| Zero Manual Work | No Helm charts, no dashboard JSON wiring, no RBAC debugging |

Once applied, Bifrost sets up everything inside the bifrost namespace, and logs from all pods in the target namespaces become visible in Grafana instantly.

## Quick Start

### Installation
- **Apply the installation manifest**

```bash
$ kubectl apply -f https://raw.githubusercontent.com/peek8/bifrost/refs/heads/main/installations/install.yaml
```

This will install the operator (as deployment) along with the RBAC manifests needed for the operator. This will create a namespace named `bifrost` and the operator will be running there.

**Check the operator pods, the pod should show ready in a while:**

```bash
$ kubectl get pods -n bifrost
```

**Apply the following sample LogSpace Resource (you can also directly apply from [github sample](https://raw.githubusercontent.com/peek8/bifrost/refs/heads/main/config/samples/bifrost_v1alpha1_logspace.yaml)):**

```yaml
apiVersion: bifrost.peek8.io/v1alpha1
kind: LogSpace
metadata:
  labels:
    app.kubernetes.io/name: bifrost
    app.kubernetes.io/managed-by: kustomize
  name: sample
spec:
  targetNamespaces:
    # - Add your namespaces here 
  pvc:
    # storageClass: add your storage class here eg it will be "local-path" For rancher k3s. 
  collector:
    storage: 
      size: "5Gi"
  loki:
    storage: 
      size: "5Gi"
  grafana:
    storage: 
      size: "5Gi"    
```

Add your namespaces at `targetNamespaces` and add proper storageClass for your cluster at `pvc.storageClass`. if you are not sure about it, you can get it by the following command:

```bash
$ kubectl get storageclass
```

if you save it as `logspace.yaml`, apply it as :

```bash
$ kubectl apply -f logsapce.yaml -n bifrost
```

And that's it, in a minute or so you will see all the alloy, loki and grafana pods are ready:

```
$ watch kubectl get pods -n bifrost
```

### Accessing Grafana

After all the pods are ready, now you can access the grafana dashboard. Access the grafana from localhost using `kubectl port-forward`.

```
$ kubectl port-forward svc/grafana 3000:3000 -n bifrost
```

Then navigate to: 
```
http://localhost:3000
```
The default grafana user/pass is `admin/admin`

If you browse the dashboards you can see a dashboard created (name has prefix `bifrost`)be visible and filtered using namespace and app labels.
And that's it, Cheers, Happy Logging and Monitoring !! 🎉 🎉

## Give it a Try
You can try with some sample applications from [this repository](./config/samples/sample-logger-apps/). I have used [mingrammer/flog](https://github.com/mingrammer/flog) image to generate some fake logs. The deployments.yaml file has three application: `api`, `frontend` and `backend`. Apply it in two namespace ie `asgard` and `midgard`:

```
$ kubectl create namespace asgard
$ kubectl apply -f https://raw.githubusercontent.com/peek8/bifrost/refs/heads/main/config/samples/sample-logger-apps/deployments.yaml -n asgard

$ kubectl create namespace midgard
$ kubectl apply -f https://raw.githubusercontent.com/peek8/bifrost/refs/heads/main/config/samples/sample-logger-apps/deployments.yaml -n asgard

$ kubectl apply -f https://raw.githubusercontent.com/peek8/bifrost/refs/heads/main/config/samples/bifrost_v1alpha1_logspace.yaml -n bifrost
```

This will create grafana dashboard at bifrost namespace which will scrape logs from asgard and midgard namespaces.


## Architecture
[target namespaces pods] → Grafana Alloy → Loki → Grafana Dashboard

| Component     | Deployment Model                |
| ------------- | ------------------------------- |
| Grafana Alloy | DaemonSet on all nodes          |
| Loki          | StatefulSet with PVCs           |
| Grafana       | Deployment with provisioning    |
| Dashboard     | Managed dynamically by operator |
| Datasource    | Managed dynamically by operator |


## Roadmap

- Now Loki only Supports monolithic mode, support other mode ie `simplScalable`, `microservice`
- Support s3 and other storages for Loki.
- Support for syslog, audit logs, and custom pipelines
- More status with conditions at logspace resource
- Add Routing for grafana from operator.
- Add more options to select pods from namespace eg using pod labels.
- Alerting & ruler provisioning.

## Contributing
Issues and PRs are most welcome! Whether it's docs, CRD spec improvements, code, or examples — contributions help the community.

## License
Apache 2.0, see more details at [LICENSE File](./LICENSE).

## Community
Bifrost is a [Peek8](https://peek8.io/) open source project.
Learn about our open source work and portfolio [here](https://peek8.io/#products).
If you want to collaborate with us or Invest at [Peek8](https://peek8.io/), please [contact us here](https://peek8.io/#contact).

Issues and PRs are most welcome! Whether it's docs, code improvement, or examples — contributions help the community.

Last but not the least, If Bifrost saves you hours of setup, consider giving the repository a ⭐.
Your support helps build a better developer-first observability ecosystem.
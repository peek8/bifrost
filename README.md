# Bifrost Operator — Zero-Effort Kubernetes Log Aggregation

Bifrost is a Kubernetes Operator that automatically deploys a full observability stack for logs — **Grafana Alloy → Loki → Grafana** — and wires everything up to collect logs from selected namespaces with no manual setup.

Define which namespaces you want to monitor using the `LogSpace` CRD and Bifrost provisions **all required components, dashboards, datasources, storage, RBAC and configurations** in minutes.

---

## 🌈 Why the name “Bifrost”?

The name bifrost sounds familiear to you ? Probably, you have heard it in `Thor Movie Series` :).   In [**Norse Mythology**](https://en.wikipedia.org/wiki/Norse_mythology)(from where Loki came from :) ), *Bifrost* is the magical rainbow bridge that connects **Midgard (the Earth)** to **Asgard (the world of the gods)**.

Similarly, the `Bifrost Operator` forms a **bridge between your Kubernetes applications and their observability**, connecting logs across namespaces to a central dashboard you can rely on.

---
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

With just:

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

kubectl apply -f https://github.com/peek8/

Confirm CRD:
kubectl get crds | grep logspaces


### Accessing Grafana

After the operator reconciles the LogSpace resource:

kubectl port-forward svc/grafana 3000:3000 -n bifrost


Then navigate to:

http://localhost:3000


Default dashboards will already be visible and filtered using namespace and app labels.

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

- Multi-cluster aggregation mode
- Multi-tenant LogSpace isolation
- Support for syslog, audit logs, and custom pipelines
- Alerting & ruler provisioning

## Contributing
PRs are welcome! Whether it's docs, CRD spec improvements, code, or examples — contributions help the community.

## License
Apache License

## Support the project
If Bifrost saves you hours of setup, consider giving the repository a ⭐.
Your support helps build a better developer-first observability ecosystem.
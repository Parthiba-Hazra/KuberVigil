# Kubervigil

Kubervigil is a command-line tool for analyzing the health of Kubernetes resources and API endpoints in a given namespace. It provides detailed health reports for different Kubernetes resources and checks the status of API endpoints.

## Installation

To install Kubervigil, follow these steps:

1. Clone the repository:

```bash
git clone https://github.com/Parthiba-Hazra/kubervigil.git
```
Change to the project directory:
```bash
cd kubervigil
```
Build the project:
```bash
go build
```
Move the binary to a directory in your system's PATH:
```bash
go install
```
## Usage
To use Kubervigil, you need to have a valid kubeconfig file with access to your Kubernetes cluster.

```bash
kubervigil check [command]
```
### Commands

#### APIversion command
Check preferred api-versions for all resources.

```bash
kubervigil check apiversion
```
NOTE: Providing kubeconfig path or namespace is optional if you dont provide config path it will try to get path from ~/.kube/config, and if you dont provide any namespace it will autometically serlect the default namespace.

#### API Command
Check the preferred Kubernetes API versions for resources in a given namespace.

```bash
kubervigil check api --config=<path/to/kubeconfig> --ns=<namespace>
```
### Pods Command
Analyze the health of Pods in a given namespace.

```bash
kubervigil check pods --config=<path/to/kubeconfig> --ns=<namespace>
```
#### Services Command
Analyze the health of Services in a given namespace.

```bash
kubervigil check services --config=<path/to/kubeconfig> --ns=<namespace>
````
#### Deployments Command
Analyze the health of Deployments in a given namespace.

```bash
kubervigil check deployments --config=<path/to/kubeconfig> --ns=<namespace>
```
#### Daemonsets Command
Analyze the health of DaemonSets in a given namespace.

```bash
kubervigil check daemonsets --config=<path/to/kubeconfig> --ns=<namespace>
```
#### Statefulsets command
Analyze the health of Statefulsets in a given namespace.

```bash
kubervigil check statefulsets --config=<path/to/kubeconfig> --ns=<namespace>
```

#### Configmaps Command
Analyze the health of ConfigMaps in a given namespace.

```bash
kubervigil check configmaps --config=<path/to/kubeconfig> --ns=<namespace>
```
#### Ingress Command
Analyze the health of Ingress resources in a given namespace.

```bash
kubervigil check ingress --config=<path/to/kubeconfig> --ns=<namespace>
```

We will be add more helath analysis features.
Feel free to contribute.. 
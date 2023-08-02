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
kubervigil [command]
```
### Commands
#### API Command
Check the preferred Kubernetes API versions for resources in a given namespace.

```bash
kubervigil api --config <path/to/kubeconfig> --ns=<namespace>
```
### Pods Command
Analyze the health of Pods in a given namespace.

```bash
kubervigil pods --config <path/to/kubeconfig> --ns=<namespace>
```
#### Services Command
Analyze the health of Services in a given namespace.

```bash
kubervigil services --config <path/to/kubeconfig> --ns=<namespace>
````
#### Deployments Command
Analyze the health of Deployments in a given namespace.

```bash
kubervigil deployments --config <path/to/kubeconfig> --ns=<namespace>
```
#### Daemonsets Command
Analyze the health of DaemonSets in a given namespace.

```bash
kubervigil daemonsets --config <path/to/kubeconfig> --ns=<namespace>
```
#### Configmaps Command
Analyze the health of ConfigMaps in a given namespace.

```bash
kubervigil configmaps --config <path/to/kubeconfig> --ns=<namespace>
```
#### Ingress Command
Analyze the health of Ingress resources in a given namespace.

```bash
kubervigil ingress --config <path/to/kubeconfig> --ns=<namespace>
```

We will be add more helath analysis features.
Feel free to contribute.. 
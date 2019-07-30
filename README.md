# Pilot
Pilot is a CLI tool used for managing Kubernetes clusters through a cloud provider's managed service. Currently support is being built for:

* Azure Kubernetes Service
* Google Kubernetes Engine
* Amazon Elastic Kubernetes Service

Pilot takes a declarative approach by looking at your configuration file, assessing the state of the resources from the cloud provider, and reconciliing the differences.

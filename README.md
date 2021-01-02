# GO REST-API App on Kubernetes
This repository basically helps you to deploy an App written in Go to recieve and serve http requests on kubernetes. We use the minukube tool to deploy and test the application. Minikube is a tool that helps us to run the kubernetes cluster on the local machine. It's a single node kubernetes cluster that helps the local development of applications and k8 objects. 

# About the App
The App which is written in Go will create a web server that listens to a specified port and serve the requests coming to it depends on the routing added to the application. The App here is a containerized one and it would listen to the port 8080 in the container, when the path http://hostname:8080/hello is requested, it will fetch some data from the database associated with the application and will send it as the response. Note that the URL changes depends how we expose the app to the outside world or within kubernetes itself.

# prerequisites 
In order to run and test the App on your local machine we need the below tools installed on your local machine depends on your Operating System flavor. 

1. Install VirtualBox on your local machine which is necessary for running minikube, feel free to skip the steps if you are already having minikube instaleld and running. Refer the link below for VirtualBox downloads: 
https://www.virtualbox.org/wiki/Downloads

2.Download and install docker desktop as we are using docker environment set up to run k8s and test the application, follow the link below for the same.
https://docs.docker.com/desktop/

3. Download and install the kubectl tool which helps to communicate with the cluster, follow the link below to get more information.
https://kubernetes.io/docs/tasks/tools/

4. Download and install helm which is a package manager for k8s and helps to make the deployment of apps on kubernetes easy. Install helm from the below link.
https://helm.sh/docs/intro/install/


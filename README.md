# GO REST-API App on Kubernetes
This repository basically helps you to deploy an App written in Go to recieve and serve http requests on kubernetes. We use the minukube tool to deploy and test the application. Minikube is a tool that helps us to run the kubernetes cluster on the local machine. It's a single node kubernetes cluster that helps the local development of applications and k8 objects.

# Overview
The App which is written in Go will create a web server that listens to a specified port and serve the requests coming to it depends on the routing added to the application. The App here is a containerized one and it would listen to the port 8080 in the container, when the path http://hostname:8080/hello is requested, it will fetch some data from the database associated with the application and will send it as the response. Note that the URL changes depends how we expose the app to the outside world or within kubernetes itself.

# Prerequisites
In order to run and test the App on your local machine we need the below tools installed on your local machine depends on your Operating System flavor.

1. Install VirtualBox on your local machine which is necessary for running minikube, feel free to skip the steps if you are already having minikube instaleld and running. Refer the link below for VirtualBox downloads: https://www.virtualbox.org/wiki/Downloads

2. Download and install docker desktop as we are using docker environment set up to run k8s and test the application, follow the link below for the same. https://docs.docker.com/desktop/

3. Download and install the kubectl tool which helps to communicate with the cluster, follow the link below to get more information. https://kubernetes.io/docs/tasks/tools/

4. Download and install helm which is a package manager for k8s and helps to make the deployment of apps on kubernetes easy. Install helm from the below link. 
https://helm.sh/docs/intro/install/

# Running the App
The repository contains all the necessary app files, kubernetes manisfests, helmcharts and Dockerfile to make the deployment easy. The Dockerfile will build and package the application will give us the required docker image to run on kubernetes. It's a multi stage Dockerfile which actually reduces the complexity and size of the final image. Follow the steps below to build the docker image and run the app on kubernetes environment powered by minikube. 

1. Open the terminal and run the below command to clone the repository to the local machine. 
```
git clone https://github.com/jishnu-codes/go-on-k8s
```
2. Start the minikube service on the local machine using the Virtualbox as driver. The command may change according to the flavor of your operating system. (This commands were tested using a Mac machine at the time of development)
```
minikube start — driver=virtualbox
```
3. Run the below command so that the minikube will be able to use the Docker daemon inside the minikube instance. This is required because we are using all development operations locally and this enables to load the docker image from the local machine itself. The `imagePullPolicy: Never` in the deployment manifest of App triggers the same. 
```
eval $(minikube docker-env)
```
4. Move to the directory where the Dockerfile resides and run the below command to build the docker image and tag it as you wish.
```
docker build -t mygoapp:latest .
```
5. Now that you have the docker image to run on k8s, the application uses MySQL to serve the requests coming to it as the data being fetched from the database. The kubernetes manifests includes all the required YAML files to deploy the app. It also has the helmchart to make the deployment easy and automated. 

# Creating the secret for MySQL: 
When deploy the MySQL container, the root password must be passed to it in order to initialize the datbase. This can be done in two ways. 
a) Create the secret manually and pass it as environment variable to the container of MySQL. Use the below kubectl command to achieve the same:
```
kubectl create secret generic mysql-pass --from-literal=password=<desired-password>
```
b) Keep the secret yaml file inside MySQL helmchart that contains the encoded MySQL root password and automatically deploy it when helm install command runs. 

Choose either of the above methods, here the chart creates the secret and pass it to the MySQL container as env variable. The same will be used by the app container to connect to the database while initializing the app container.

6. Move to /helmcharts/mysql directory and run the command below to install MySQL, it will be having a persitent volume taken from the host machine also a claim bounded to it. 
```
helm install mysql . 
```
7. Install the app using the below command from /helmcharts/goapp
```
helm install goapp .
```
8. The goapp is exposed as a NodePort service so that we can access it from the local machine, using command line or using web browser. The following command will give you the URL that the service is exposed to: 
```
minikube service goapp --url
```
9. According to the app, the route `hostname/hello` will fetch something from the datasbase. Use the command below to test the application using a curl command, assuming that the output from the above command is `http://192.168.99.106:30374`:
```
curl http://192.168.99.106:30374/hello
```
10. The desired output is shown below, provided the data inserted into the database is as per mysql-init-config.yaml under MySQL helmchart templates:
```
 jishnuks@My-MacBook-Air  ~  curl http://192.168.99.106:30374/hello                                                                                        ✔  5307  10:41:48
Hello New Year, 2021
 jishnuks@My-MacBook-Air  ~ 
 ```

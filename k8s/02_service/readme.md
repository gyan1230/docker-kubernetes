1. Deployment:
    1. Rolling upgrade from v1 to v2 one by one
	2. Roll back v2 to v1
	3. Pause the env , make the changes(resources allocation, web server), then resume
All of these done by deployment.
Not much different between deployment and RS.

2. Namespaces:
	1. Creating pods, svc, namespaces is having default namespace, created by k8s when cluster setup
	2. K8s create setup of pods, svc required for networking solutions dns, to isolate from user and prevent accidentally delete, - kube-system
	3. Kube-public also created for user

Q. How to connect different namespace service with each other?

Web-pod of default name space want to connect db-service of dev name space?
Ans:
	service.namespace.svc.cluster.local

Note:
	create namespace by cmdline or move the namespace to metadata section

To switch the dev namespace permanently , so that no need to specify namespace option anymore.
	1. Kube Ctl config command in the current context.
Then kubectl get pod will get dev namespace details
For others (default, prod) need to specify -namespace=prod/default

kubectl get pods --all-namespaces to get all pods with all namespace 

Limit resources in namespaces create resourceQuota

3. Services:
	1. Services enable loose coupling between Microservices  in our application.

To access the web app outside we need service like NodePort svc, for accessible internal port on the node.

2.  types of services
	1.	NodePort
	2.	Cluster IP : it creates a virtual IP inside the cluster to enable communication between diff svcs such as frontend server, backend server
	3.	Load-balancer : for load balance our app in supported cloud provided 

1. NodePort:
	3 ports are involved : 
	a.	One web server is running in the pod referred as TargetPort (where the svc forward the request to)
	b.	Second on the service itself referred as Port itself - svc is like virtual server inside the node , so inside cluster it has its own ipaddr, that IP called as cluster IP of the service
	c.	Third one is the NodePort

How to create service?

But we didn’t mention the target port of which pod? ,there could be 100 no of pods with web svc running on port 80

So we use labels and selectors to link together.

Pull the labels from pod-definition file under selector section.
 Create the service by kubectl create -f service.-definition.yml 


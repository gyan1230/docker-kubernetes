Labels are property attached to each items.
Selectors filter the items based on labels.

Different types of object created in k8s, like pods, services, replicates, deployments etc.
Over time we may end of 100 or 1000 different objects in the cluster.
Then we need a way to filter objects by different categories, like by types, by applications, or by their functionalities.
Each object attach labels by as per our need, like app, function etc.
￼
   
How to specify labels?
While creating pods, specify labels under metadata, labels as a key value format. Add as many labels as we like. 
￼

Once the pod created , use kubectl get commands with —selector key=value to get the filtered pods with the labels.
Kubectl get pods —selector app=app1

K8s objects use labels and selectors internally to connect different objects together. 
Fo eg. To create a replicaset consisting of three different pods, we first label the pod definition yml file and use selector in replicaset to group the pods.
U can see label define in two places
￼

The labels under the template section are configured for the pods.
The labels on the top is for the RS itself.(not concern for nw)
On creation if the label match, it create RS successfully.
Same for while creating services.

Annotation.
	Annotation used to record other details for  informative purpose.
Like tool details name,build, version etc.
￼

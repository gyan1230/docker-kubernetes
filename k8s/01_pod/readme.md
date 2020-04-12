1. Pods with yaml

Only name and labels allow under metadata.
Under labels , any key-value pair is ok

2. Replication controller :(RC)

Copy the pod metadata part to template part of spec.
kubectl create -f rc-definition.yml

3. ReplicaSet: (RS)
Selector is required as it is taken the pods into consideration which created before in the label as in selector.
Pod definition file already in the template
Because rs also manage pods that were not created part of rs creation.
Pods created b4 rs , match labels specify in the selector , rs take those pods into consideration, when creating rs.

Selector is not  a required field in case of rc.
But it is still available, when we skip it , it assume label as the pod’s definition file.

In RS user input is required for this property and it has to be written in the form of matchLabels as below.
Matchlabels selectors simply match the labels specify under it to the labels on the pod.
Eg : type - front-end match labels to the labels in the pod-template.

Other than type matchLabels are available in RS not available in RC.

Under the selector section , use matchLabels filter, and provide the same label while creating the pods.

How to scale replicas?
	1.	Change the definition file to 6 and recreate it by using kubectl replace -f rs-def.yml
	2.	kubectl scale  —replicas=6 -f rs-def.yml

To delete rs or rc,pod,
kubectl delete rs/rc/po <name of replicaset>/<name of replicationController>/<name of pod>

Exercise:
Cmd to get yml file 
Kubectl get rs new-replica-set -o yaml > my-def.yml

Then delete wrong replica-set.
And edit the created one and create new rs.
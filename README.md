# ssh-bastion

bastion-controller folder contains a kubebuilder generated operator. The operator watches the Tenant CRs and updates a configmap inserting the new public keys.

## Usage

```shell
# Install the CRDs
kubectl apply -f bastion-operator/config/crd/bases/ 

# Deploy the bastion and create a placeholder configmap to get the pod started
kubectl apply -f bastion-deployment.yml

# Open another terminal tab and start the operator
cd bastion-operator
make run

# Go back to the first terminal tab, insert one of your pubkeys to a tenant CR (i.e. cr_mariorossi.yml)
nano cr_mariorossi.yml

# Create the new resource
kubectl apply -f cr_mariorossi.yml

# Check that the configmap is updated
kubectl describe configmap auth-keys-bastion

# Start LoadBalancer service and then connect
kubectl apply -f service.yml
ssh bastion@127.0.0.1 -p 2222
```
### Note:
It could take ~1 minute to update the authorized_keys in the pods (See [kubernetes docs](https://kubernetes.io/docs/tasks/configure-pod-container/configure-pod-configmap/#mounted-configmaps-are-updated-automatically)).


Desired output:
```
Welcome to Alpine!

The Alpine Wiki contains a large amount of how-to guides and general
information about administrating Alpine systems.
See <http://wiki.alpinelinux.org/>.

You can setup the system with the command: setup-alpine

You may change this message by editing /etc/motd.

This account is not available

```

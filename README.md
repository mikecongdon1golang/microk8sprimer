This is my testing / planning on how to use Kubernetes microk8s on my ubuntu dev env.


references:
# https://microk8s.io/docs/
# https://kukulinski.com/10-most-common-reasons-kubernetes-deployments-fail-part-1/
# https://microk8s.io/docs/working

I have docker, docker-compose, kubernetes, microk8s installed already.
I have microk8s.kubectl aliased as mkubectl


Step 1 (this was annoying)
*   you have to export the image out of docker - local and put it in a registry, or you have to somehow get it somewhere that kubernetes can see it. This is how I did it:

>docker build . -t hello:local

then to make sure it built properly:
>docker image ls 
export to tar:
>docker save hello:local > hello.tar

import it to k8s namespace (this part didn't seem to quite work as expected - details below)
>microk8s.ctr -n k8s.io image import hello.tar

>microk8s.ctr -n k8s.io images ls
it shows up as *docker.io/library/hello:local* and I would have expected it to show up as *k8s.io/hello:local*

To create a deployment:
>mkubectl create deployment hellodep --image=docker.io/library/hello:local

Then scale it:
>mkubectl scale deployment hellodep --replicas=2

Then, expose it:
>mkubectl expose deployment hellodep --type=NodePort --port=3000 --name=hello-exp

Then, check what your local port is:
>mkubectl get all --all-namespaces

*NAMESPACE     NAME                           TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)*

*default       service/hello-exp            NodePort    10.152.183.253   <none>        3000:31035/TCP*
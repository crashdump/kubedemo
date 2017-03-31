Kube Demo
=========

Description
-----------

This is a simple demonstration of Kubernetes composed of two Go webapp (frontend
and backend).

How To
------

You will find a Makefile at the root of the repo which include most of the
command one would need to create a Namespace, deploy, scale containers, etc..

*Examples:*

Build all the containers (if you are using minikube, you  might want to expose
its Docker environemt first)
```
eval $(minikube docker-env)
make all
```

Deploy the applications in Kubernetes
```
make kube-deploy
```

Destroy all the resources in Kubernetes
```
make kube-destroy
```

Notes
-----

I realise building the simple Go apps in Docker is _totally overkill_ but this
project is meant to explain what a full featured project would look like.

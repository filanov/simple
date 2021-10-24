# Simplesrv
simple golang rest service

## Local deployment (temp)

Create k3d cluster with local registry and port configuraiton
```
k3d create --enable-registry -p 8081:80@server
```

Build an push can be done with a single command, makefile configuration match k3d config from above
```
make container-push
```

Export kubeconfig
```
export KUBECONFIG="$(k3d get-kubeconfig --name='k3s-default')"
```
Deploy the application
```
kubectl apply -f deploy/app.yaml
```
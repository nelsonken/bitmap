# bitmap

> redist bitmap get and set, for play with docker and drone

## plays

- dockerfile
- docker-compose
- drone
- kubernets

## play step

### drone and compose

- docker-compose up

### kubernets

- docker build .
- docker pull redis
- create ./data/ put config.ini into ./data/
- kubectl create -f kubernets-dashboard.yaml
- kubectl proxy
- [view http://127.0.0.1:8001/api/v1/namespaces/kube-system/services/https:kubernetes-dashboard:/proxy/#!/login](http://127.0.0.1:8001/api/v1/namespaces/kube-system/services/https:kubernetes-dashboard:/proxy/#!/login)
- play it

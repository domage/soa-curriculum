htpasswd -c auth loki
kubectl create secret generic loki-auth --from-file=auth

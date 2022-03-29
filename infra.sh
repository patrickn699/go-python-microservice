gcloud config set compute/zone us-central1-a
gcloud container clusters create pn-cluster --num-nodes 3 
gcloud container clusters get-credentials pn-cluster

docker pull patrick699/frontend:1.0.1
docker pull patrick699/conv:1.0.0

docker tag patrick699/frontend:1.0.1 gcr.io/patrick-project/frontend:1.0.1
docker tag patrick699/conv:1.0.0 gcr.io/patrick-project/conv:1.0.0

docker push gcr.io/patrick-project/frontend:1.0.1
docker push gcr.io/patrick-project/conv:1.0.0

kubectl create service clusterip conv --tcp 9000
kubectl create deployment conv --image=gcr.io/patrick-project/conv:1.0.0
kubectl create deployment front --image=gcr.io/patrick-project/frontend:1.0.1

kubectl expose deployment front --type=LoadBalancer --port=8000 --target-port=5000

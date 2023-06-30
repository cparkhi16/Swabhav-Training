docker build -t cparkhi16/multi-client-k8s:latest -t cparkhi16/multi-client-k8s:$SHA -f ./client/Dockerfile ./client
docker build -t cparkhi16/multi-server-k8s:latest -t cparkhi16/multi-server-k8s:$SHA -f ./server/Dockerfile ./server
docker build -t cparkhi16/multi-worker-k8s:latest -t cparkhi16/multi-worker-k8s:$SHA -f ./worker/Dockerfile ./worker

docker push cparkhi16/multi-client-k8s:latest
docker push cparkhi16/multi-server-k8s:latest
docker push cparkhi16/multi-worker-k8s:latest

docker push cparkhi16/multi-client-k8s:$SHA
docker push cparkhi16/multi-server-k8s:$SHA
docker push cparkhi16/multi-worker-k8s:$SHA

kubectl apply -f k8s
kubectl set image deployments/server-deployment server=cparkhi16/multi-server-k8s:$SHA
kubectl set image deployments/client-deployment client=cparkhi16/multi-client-k8s:$SHA
kubectl set image deployments/worker-deployment worker=cparkhi16/multi-worker-k8s:$SHA
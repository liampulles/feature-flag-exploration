GOBIN := $(shell go env GOBIN)

run:
	go run .

test:
	while sleep 0.5; do curl http://localhost:9080/; done

apply-flags:
	-kubectl delete configmap flags
	kubectl create configmap flags --from-file=flags/flags.yaml

# k3d cluster create --registry-use k3d-registry:33065 -p "8081:80@loadbalancer" --agents 1
push: apply-flags
	docker build -t k3d-registry:33065/basic:latest .
	docker push k3d-registry:33065/basic:latest
	kubectl apply -f k8s.yaml
	kubectl rollout restart deployment/basic

test-cluster:
	while sleep 0.5; do curl http://localhost:8081/; done
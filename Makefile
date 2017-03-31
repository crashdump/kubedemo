.PHONY: frontend backend

all: frontend backend

clean:
	rm -f dist/*
	$(shell docker images|grep kubedemo|awk '{ print $3}'|xargs docker rmi -f)

frontend:
	docker run --rm -v "$(PWD)":/usr/src/myapp -w /usr/src/myapp golang:1.8 go build -v -o dist/frontend frontend/main.go
	docker build -f frontend/Dockerfile -t kubedemo/kubedemo-frontend:$(shell cat frontend/VERSION) .

backend:
	docker run --rm -v "$(PWD)":/usr/src/myapp -w /usr/src/myapp golang:1.8 go build -v -o dist/backend backend/main.go
	docker build -f backend/Dockerfile -t kubedemo/kubedemo-backend:$(shell cat backend/VERSION) .

kube-deploy:
	-kubectl create namespace kubedemo
	kubectl -n kubedemo apply -f kubernetes/

kube-destroy:
	kubectl delete namespace kubedemo

kube-scale-frontend:
	kubectl -n kubedemo scale deployment frontend --replicas 2
	kubectl -n kubedemo scale deployment backend --replicas 6

kube-get-info:
	kubectl -n kubedemo get services
	kubectl -n kubedemo get deployments


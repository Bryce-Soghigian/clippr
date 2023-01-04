deploy:
	kubectl apply -k $PWD/kustomize/ -n $(NAMESPACE) 

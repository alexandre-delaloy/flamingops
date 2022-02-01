# the template validation
val=false
# the chart name
chart=flamingops
# the app version
ver=1.0.0

.PHONY: help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(firstword $(MAKEFILE_LIST)) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

encode-config:
	cat envs/kubeconfig | base64 > envs/kubeconfig.base64

login: ## connect to k8s cluster
	export KUBECONFIG=envs/kubeconfig

check: ## check the chart output
	helm template . \
		--kubeconfig envs/kubeconfig \
		--values envs/values.yaml \
		--validate=$(val) \
		--generate-name \
		--debug > envs/output.yaml  

upgrade: ## upgrade the actual version of chart
	helm upgrade $(chart) \
		--install \
		--kubeconfig envs/kubeconfig \
		--values envs/values.yaml \
		--cleanup-on-fail \
		--debug \
		--version=$(ver) .

install: ## install all charts
	make install-ingress-nginx
	make install-cert-manager
	
install-ingress-nginx: ## install ingress nginx controller chart
	make login
	echo -e "\e[34m[INSTALLING] ingress-nginx/ingress-nginx\e[0m"
	sleep 1
	helm repo add ingress-nginx 'https://kubernetes.github.io/ingress-nginx'
	helm repo update
	helm install --kubeconfig envs/kubeconfig ingress-nginx ingress-nginx/ingress-nginx
	echo -e "\e[32m[INSTALLED] ingress-nginx/ingress-nginx\e[0m"
	sleep 1

install-cert-manager: ## install cert-manager chart
	make login
	echo -e "\e[34m[INSTALLING] jetstack/cert-manager\e[0m"
	sleep 1
	helm repo add jetstack 'https://charts.jetstack.io'
	helm repo update
	helm install --kubeconfig envs/kubeconfig \
	cert-manager jetstack/cert-manager \
	--namespace cert-manager \
	--create-namespace \
	--version v1.5.3 \
	--set installCRDs=true
	echo -e "\e[32m[INSTALLING] jetstack/cert-manager\e[0m"
	sleep 1

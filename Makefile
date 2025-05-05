bind:
	(kubectl port-forward service/sample-go-app-service 8081:8081 &) && \
	(kubectl port-forward service/argocd-server -n argocd 8080:443 &) && \
	wait
.PHONY: generate-deepcopies
generate-deepcopies:
	deepcopy-gen ./api/starwars/v1 --output-file zz_generated.deepcopy.go
	controller-gen crd paths=./api/... output:crd:dir=./k8s

.PHONY: generate-crd
generate-crd:
	controller-gen crd paths=./api/... output:crd:dir=./k8s

.PHONY: generate-clientset
generate-clientset:
	client-gen --input-base="github.com/okontajneroch/starwars/api" \
	   --input="starwars/v1" \
	   --clientset-name="clientset" \
	   --output-dir=. \
	   --output-pkg="github.com/okontajneroch/starwars"

.PHONY: generate
generate: generate-deepcopies generate-crd generate-clientset


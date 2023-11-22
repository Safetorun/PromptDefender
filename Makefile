MODULES := $(shell (find .  -type f -name '*.go' -maxdepth 2 | sed -r 's|/[^/]+$$||' |cut -c 3-|sort |uniq))
AWS_MODULES := $(shell cd deployments/aws && find . -type f -name '*.go' -maxdepth 2 | sed -r 's|^\./|deployments/aws/|' | grep "lambda_" | sed -r 's|/[^/]+$$||' | sort | uniq)
PROJECT_DIR := $(shell pwd)

setup-workspace:
	export TF_VAR_branch_name=$$(git rev-parse --abbrev-ref HEAD);\
	if [ "$$TF_VAR_branch_name" = "main" ] && [ "$$INTEGRATION_TEST" != "true" ]; then \
		echo "On 'main' branch. Using the 'default' workspace..."; \
		cd terraform && terraform workspace select -or-create default; \
		terraform workspace show; \
		cd ..; \
	else \
		echo "Workspace $$TF_VAR_branch_name exists. Selecting it..."; \
		cd terraform && terraform workspace select -or-create  $$TF_VAR_branch_name; \
		terraform workspace show; \
		cd ..; \
	fi

test: build
	cd aiprompt && go test -v ./... -cover
	cd pii && go test -v ./... -cover
	cd pii_aws && go test -v ./... -cover
	cd canary && go test -v ./... -cover
	cd prompt && go test -v ./... -cover
	cd moat  && go test -v ./... -cover
	cd keep  && go test -v ./... -cover
	cd wall && go test -v ./... -cover

build: generate
	for aws_module in $(AWS_MODULES) ; do \
	   cd $$aws_module && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main || exit 1; cd $(PROJECT_DIR) ; \
	done

deploy: setup-workspace build
	export TF_VAR_commit_version=`git rev-parse --short HEAD` &&\
	cd terraform && terraform init && terraform apply -auto-approve &&\
	terraform output -json > terraform_output.json

install:
	for number in  $(MODULES) ; do \
       cd $$number && go get ./... || exit 1; cd .. ; \
    done
	for aws_module in $(AWS_MODULES) ; do \
	   cd $$aws_module && go get ./... || exit 1; cd $(PROJECT_DIR) ; \
	done

tidy:
	for number in $(MODULES); do \
		cd $$number && go mod tidy || exit 1; cd .. ; \
	done
	for aws_module in $(AWS_MODULES) ; do \
	   cd $$aws_module && go mod tidy || exit 1; cd $(PROJECT_DIR) ; \
	done

upgrade:
	for number in  $(MODULES) ; do \
	   cd $$number && go get -u all  || exit 1; cd .. ; \
	done
	for aws_module in $(AWS_MODULES) ; do \
	   cd $$aws_module && go get -u all || exit 1; cd $(PROJECT_DIR) ; \
	done

clean:
	for number in  $(MODULES) ; do \
	   cd $$number && go clean || exit 1; cd .. ; \
	done
	for aws_module in $(AWS_MODULES) ; do \
	   cd $$aws_module && go clean || exit 1; cd $(PROJECT_DIR) ; \
	done

generate:
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
	for aws_module in $(AWS_MODULES) ; do \
	   cd $$aws_module && oapi-codegen -package main -generate types $(PROJECT_DIR)/openapi.yml > api.gen.go || exit 1; cd $(PROJECT_DIR); \
	done
	oapi-codegen -package integration_test_harness -generate types,client $(PROJECT_DIR)/openapi.yml > integration_test_harness/api.gen.go

generate_jailbreak:
	cd builder\
	 && pip install -r requirements.txt && python3 clean_jailbreaks_into_json.py\
  	 && python3 jailbreak_embeddings.py && go build -o main && ./main

integration_test:
	go install github.com/tomwright/dasel/cmd/dasel@latest
	export URL=`dasel select -f terraform/terraform_output.json '.api_url.value' | tr -d '"'` &&\
	export DEFENDER_API_KEY=`dasel select -f terraform/terraform_output.json '.api_key_value.value' | tr -d '"'` &&\
	cd integration_test_harness && go test ./...

destroy: setup-workspace
	export TF_VAR_commit_version=`git rev-parse --short HEAD` &&\
	cd terraform && terraform init && terraform destroy -auto-approve

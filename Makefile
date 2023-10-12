MODULES := $(shell (find .  -type f -name '*.go' -maxdepth 2 | sed -r 's|/[^/]+$$||' |cut -c 3-|sort |uniq))

test:
	cd aiprompt && go test -v ./... -cover
	cd app && go test -v ./... -cover
	cd pii && go test -v ./... -cover
	cd pii_aws && go test -v ./... -cover
	cd canary && go test -v ./... -cover
	cd prompt && go test -v ./... -cover

build:
	cd lambda_prompt_shield && go build -v ./...
	cd lambda_prompt_shield && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ../main lambda.go
	cd lambda_prompt_builder && go build -v ./...
	cd lambda_prompt_builder && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main lambda.go


deploy:
	cd lambda_prompt_shield && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ../main lambda.go
	cd lambda_prompt_builder && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main lambda.go
	cd terraform && terraform init && terraform apply -auto-approve

install:
	for number in  $(MODULES) ; do \
       cd $$number && go get ./... || exit 1; cd .. ; \
    done

tidy:
	for number in  $(MODULES) ; do \
	   cd $$number && go mod tidy  || exit 1; cd .. ; \
	done

upgrade:
	for number in  $(MODULES) ; do \
	   cd $$number && go get -u all  || exit 1; cd .. ; \
	done

clean:
	for number in  $(MODULES) ; do \
	   cd $$number && go clean || exit 1; cd .. ; \
	done
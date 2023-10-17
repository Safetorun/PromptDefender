MODULES := $(shell (find .  -type f -name '*.go' -maxdepth 2 | sed -r 's|/[^/]+$$||' |cut -c 3-|sort |uniq))
AWS_MODULES := $(shell cd deployments/aws && find . -type f -name '*.go' -maxdepth 2 | sed -r 's|/[^/]+$$||' | cut -c 3- | sort | uniq)

test:
	cd aiprompt && go test -v ./... -cover
	cd pii && go test -v ./... -cover
	cd pii_aws && go test -v ./... -cover
	cd canary && go test -v ./... -cover
	cd prompt && go test -v ./... -cover
	cd moat  && go test -v ./... -cover
	cd keep  && go test -v ./... -cover

build:
	cd deployments/aws/lambda_moat && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main lambda.go
	cd deployments/aws/lambda_keep && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main lambda.go


deploy:
	cd deployments/aws/lambda_moat && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main lambda.go
	cd deployments/aws/lambda_keep && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main lambda.go
	cd terraform && terraform init && terraform apply -auto-approve

install:
	for number in  $(MODULES) ; do \
       cd $$number && go get ./... || exit 1; cd .. ; \
    done

tidy:
	for number in $(MODULES); do \
		cd $$number && go mod tidy || exit 1; cd .. ; \
	done
	cd deployments/aws/lambda_moat && go mod tidy
	cd deployments/aws/lambda_keep && go mod tidy

upgrade:
	for number in  $(MODULES) ; do \
	   cd $$number && go get -u all  || exit 1; cd .. ; \
	done
	cd deployments/aws/lambda_moat && go get -u all
	cd deployments/aws/lambda_keep && go get -u all

clean:
	for number in  $(MODULES) ; do \
	   cd $$number && go clean || exit 1; cd .. ; \
	done
	cd deployments/aws/lambda_moat && go clean
	cd deployments/aws/lambda_keep && go clean

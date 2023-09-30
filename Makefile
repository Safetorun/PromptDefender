test:
	cd aiprompt && go test -v ./... -cover
	cd app && go test -v ./... -cover
	cd pii && go test -v ./... -cover
	cd pii_aws && go test -v ./... -cover
	cd canary && go test -v ./... -cover
	cd prompt && go test -v ./... -cover

build:
	cd lambda && go build -v ./...
	cd lambda && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ../main lambda.go

deploy:
	cd lambda && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ../main lambda.go
	cd terraform && terraform init && terraform apply -auto-approve
update:
	cd aiprompt && go get -u ./...
	cd app && go get -u ./...
	cd pii && go get -u ./...
	cd pii_aws && go get -u ./...
	cd canary && go get -u ./...
	cd prompt && go get -u ./...

tidy:
	cd lambda && go mod tidy
	cd app && go mod tidy
	cd aiprompt && go mod tidy
	cd pii && go mod tidy 
	cd pii_aws && go mod tidy 
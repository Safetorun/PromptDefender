test:
	cd aiprompt && go test -v ./...
	cd app && go test -v ./...
	cd pii && go test -v ./...
	cd pii_aws && go test -v ./...

build:
	cd lambda && go build -v ./...
	cd lambda && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ../main lambda.go

deploy:
	cd lambda && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ../main lambda.go
	cd terraform && terraform init && terraform apply -auto-approve

tidy:
	cd lambda && go mod tidy
	cd app && go mod tidy
	cd aiprompt && go mod tidy
	cd pii && go mod tidy 
	cd pii_aws && go mod tidy 
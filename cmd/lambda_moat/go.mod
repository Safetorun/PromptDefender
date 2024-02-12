module github.com/safetorun/PromptDefender/deployments/aws/lambda_moat

go 1.20

require (
	github.com/aws/aws-lambda-go v1.46.0
	github.com/safetorun/PromptDefender/badwords v0.0.0-20231210112259-15b98f65cf67
	github.com/safetorun/PromptDefender/badwords_embeddings v0.0.0-20231210112259-15b98f65cf67
	github.com/safetorun/PromptDefender/embeddings v0.0.0-20231210112259-15b98f65cf67
	github.com/safetorun/PromptDefender/internal/base_aws v0.0.0-20240130072532-9a4bc83dc7d2
	github.com/safetorun/PromptDefender/moat v0.0.0-20231210112259-15b98f65cf67
	github.com/safetorun/PromptDefender/pii_aws v0.0.0-20231210112259-15b98f65cf67
	go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-lambda-go/otellambda v0.48.0
	go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-lambda-go/otellambda/xrayconfig v0.48.0
	go.opentelemetry.io/contrib/propagators/aws v1.23.0
	go.opentelemetry.io/otel v1.23.0
)

require (
	github.com/aws/aws-sdk-go v1.50.7 // indirect
	github.com/cenkalti/backoff/v4 v4.2.1 // indirect
	github.com/drewlanenga/govector v0.0.0-20220726163947-b958ac08bc93 // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.19.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/rogpeppe/go-internal v1.11.0 // indirect
	github.com/safetorun/PromptDefender/pii v0.0.0-20231210081706-40db07878111 // indirect
	github.com/sashabaranov/go-openai v1.19.2 // indirect
	go.opentelemetry.io/contrib/detectors/aws/lambda v0.48.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.23.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.23.0 // indirect
	go.opentelemetry.io/otel/metric v1.23.0 // indirect
	go.opentelemetry.io/otel/sdk v1.23.0 // indirect
	go.opentelemetry.io/otel/trace v1.23.0 // indirect
	go.opentelemetry.io/proto/otlp v1.1.0 // indirect
	golang.org/x/net v0.20.0 // indirect
	golang.org/x/sys v0.16.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240102182953-50ed04b92917 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240102182953-50ed04b92917 // indirect
	google.golang.org/grpc v1.61.0 // indirect
	google.golang.org/protobuf v1.32.0 // indirect
)

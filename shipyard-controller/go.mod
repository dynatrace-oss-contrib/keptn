module github.com/keptn/keptn/shipyard-controller

go 1.16

require (
	github.com/Masterminds/semver/v3 v3.1.1
	github.com/benbjohnson/clock v1.1.0
	github.com/cloudevents/sdk-go/v2 v2.5.0
	github.com/gin-gonic/gin v1.7.4
	github.com/go-openapi/swag v0.19.14 // indirect
	github.com/go-test/deep v1.0.5
	github.com/google/uuid v1.3.0
	github.com/jeremywohl/flatten v1.0.1
	github.com/keptn/go-utils v0.9.0
	github.com/ory/dockertest/v3 v3.6.5
	github.com/sirupsen/logrus v1.8.1
	github.com/stretchr/testify v1.7.0
	github.com/swaggo/swag v1.7.0
	go.mongodb.org/mongo-driver v1.7.2
	go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin v0.22.0
	go.opentelemetry.io/otel v1.0.0-RC2
	go.opentelemetry.io/otel/exporters/jaeger v1.0.0-RC2
	go.opentelemetry.io/otel/sdk v1.0.0-RC2
	go.opentelemetry.io/otel/trace v1.0.0-RC2
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
	k8s.io/api v0.22.1
	k8s.io/apimachinery v0.22.1
	k8s.io/client-go v0.22.1
)

replace github.com/keptn/go-utils => github.com/dynatrace-oss-contrib/go-utils v0.8.5-0.20210902083938-166ac7f851b4


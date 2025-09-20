module warhammer

go 1.23.0

toolchain go1.23.8

require (
	github.com/aws/aws-lambda-go v1.47.0
	github.com/awslabs/aws-lambda-go-api-proxy v0.16.2
	github.com/go-chi/cors v1.2.1
	github.com/joho/godotenv v1.5.1
	github.com/lib/pq v1.10.9
)

require (
	github.com/go-chi/chi/v5 v5.2.2
	github.com/golang-jwt/jwt/v4 v4.5.2
	github.com/google/uuid v1.6.0
	golang.org/x/crypto v0.37.0
)

replace github.com/go-chi/chi/v5 => github.com/go-chi/chi/v5 v5.2.2

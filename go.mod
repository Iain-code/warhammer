module warhammer

go 1.23.0

toolchain go1.23.8

require (
	github.com/go-chi/cors v1.2.1
	github.com/joho/godotenv v1.5.1
	github.com/lib/pq v1.10.9
)

require (
	github.com/go-chi/chi/v5 v5.2.2
	github.com/golang-jwt/jwt/v4 v4.5.2
	github.com/google/uuid v1.6.0
	golang.org/x/crypto v0.36.0
)

require github.com/stripe/stripe-go/v82 v82.5.0 // indirect

replace github.com/go-chi/chi/v5 => github.com/go-chi/chi/v5 v5.2.2

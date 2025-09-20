# Build
FROM golang:1.23-bookworm AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# build your app binary; name it "main"
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o main ./

# Runtime: Go Lambda base (includes the runtime so you don't need /var/runtime/bootstrap)
FROM public.ecr.aws/lambda/go:1
# copy your handler into $LAMBDA_TASK_ROOT
COPY --from=build /src/main ${LAMBDA_TASK_ROOT}
# tell Lambda which executable to run
CMD ["main"]

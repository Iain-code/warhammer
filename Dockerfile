# Build
FROM golang:1.23-bookworm AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bootstrap ./

# Lambda runtime
FROM public.ecr.aws/lambda/provided:al2023
COPY --from=build /src/bootstrap /var/task/bootstrap
CMD ["bootstrap"]

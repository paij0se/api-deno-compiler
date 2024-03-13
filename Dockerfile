FROM golang:1.21
WORKDIR /app
ENV GITHUB=
ENV DATABASE_URL=
COPY . .
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-api-d-c
CMD ["/docker-api-d-c"]
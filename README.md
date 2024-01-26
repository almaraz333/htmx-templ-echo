# A Go Echo server using Templ components in a Dockerized container

## To Run Locally:
### Dev:
  - Run `go mod download`
  - Run `air`
### Prod
- Run `docker build -t htmx .`
- Run `docker run -p 8080:3000 htmx`

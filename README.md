# A Go Echo server using Templ components in a Dockerized container

## To Run Locally:
### Dev:
- With Air (live reloading): 
  - Install Air with `go install github.com/cosmtrek/air@latest`
  - `cd` into the project
  - Run `air`
- Without Air:
  - Run `go mod download`
  - Run `go build -o main`
  - Run `./main`
### Prod
- Run `docker build -t htmx .`
- Run `docker run -p 8080:3000 htmx`

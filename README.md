SWStarterBackend

# Run the app
## Running during development
- Check if you have go installed by running `go version`. If not install it by following [this guide](https://go.dev/doc/install)
- Run `go run .` at project's root folder!

## Running using docker

- Make sure you have Docker in your machine
- Create a `.env` file in project's root folder (use `.env.sample` as base)
- Run `docker-compose up -d` and you're good to go!

# Examples
Assuming your PORT=8080 and you're running on localhost, an example of an endpoint would be:
```curl
curl --request GET \
  --url 'http://localhost:8080/searchMovies?query=Return'
```
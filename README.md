# README Generator

## Description
A Demo Golang project, topics :

- Simple Http Server
- Get csv from endpoint
- Convert csv to struct
- sort slice of struct
- filter slice of struct



## Installation
To use this application, please clone and get packages from root folder:
```shell
go mod tidy
```

## Usage
After cloning the repo and installing packages you can run the server:
```shell
go run cmd/ranking/main.go
```

To test using curl, postman or in my case httpie
```shell
http :8080/top top==10 language==javascript
```
These returns the top 10 Javascript repos



## License
This application is licensed under the MIT license.

## Contributing
There are no guidelines for contributing at this time.

## Tests
To run tests on the application, install
```
There is no test information for this application at this time.
```
and run `npm run test` from the command line.

    
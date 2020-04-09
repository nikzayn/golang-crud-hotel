# Hotel Golang-CRUD

Creating a web api using golang to setup crud functionalities.

## Installation

```bash
go get -v github.com/nikzayn/golang-crud-hotel
cd golang-crud-hotel && cp sample.env .env
```

## Usage

Navigate into the project directory and run the following command

```
sudo docker-compose build
sudo docker-compose up
```

### Environment Variables

| Enivironment Variable           | Description                                                                                     |
| ------------------------------- | ----------------------------------------------------------------------------------------------- |
| **PORT**                        | Port on to which application server listens to. Default value is 8080                           |
| **RESPONSE_TIMEOUT**            | Timeout for the server to write response. Default value is 100ms                                |
| **REQUEST_BODY_READ_TIMEOUT**   | Timeout for reading the request body send to the server. Default value is 20ms                  |
| **RESPONSE_BODY_WRITE_TIMEOUT** | Timeout for writing the response body. Default value is 20ms                                    |
| **PRODUCTION**                  | Flag to denote whether the server is running in production. Default value is `false`            |
| **MAX_REQUESTS**                | Maximum no. of concurrent requests supported by the server. Default value is 1000               |
| **REQUEST_CLEAN_UP_CHECK**      | Time interval after which error request app context cleanup has to be done. Default value is 2m |
| **Database**                    | All information related to database are stored in sample.env file                               |

### Endpoints
- Getting all the available hotels - ```/hotel```
- Creating the hotels information  - ```/hotel/create```
- Updating the hotels information  - ```/hotel/update```
- Deleting the hotels information  - ```/hotel/delete```

## Author

Nikhil Vaidyar<nikhilvaidyar1997@gmail.com>

# Car rental

Test task for Car rental service
It's running via docker and supporting Hot reload.
## Installation
Update .env according to your envirement, on Mac-OS need to replace variable DATABASE_HOST= via your virtual machine host IP. Same api you can use to run it in browser IP:3000
For starting run
```go
docker-compose up
```

## Environment variables
```
DATABASE_USERNAME=
DATABASE_PASSWORD=
DATABASE_NANE=
DATABASE_HOST=
DATABASE_PORT=
DEBUG_MODE=true
SERVER_PORT=3000
```

## Supported endpoint
```
GET - https://{host}/ - home endpoint
POST - https://{host}/car - create car
GET - https://{host}/car - get list of all cars
GET - https://{host}/car/color - get car collors
POST - https://{host}/car/color - create new car collor
GET - https://{host}/car/available?start_date=2022-02-01T12:00:00Z&end_date=2022-02-13T12:00:00Z - get list of available cars on selected dates
GET - https://{host}/car/{carID} - get specific car
DELETE - https://{host}/car/{carID} - delete a specific car
GET - https://{host}/car/{carID}/booking - get all booking for a specific car
POST - https://{host}/car/{carID}/booking - create booking for a specific car
GET - https://{host}/car/{carID}/booking/{bookingID} - get info about a specific booking
DELETE - https://{host}/car/{carID}/booking/{bookingID} - delete a specific booking
```

## Good to be done if it were real API:
```
- Pagination for GET endpoints
- Cors
- User auth by token
- Test
```
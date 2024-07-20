# News App Project BE 23

## Open Api 
This is link open api news app:
```
https://app.swaggerhub.com/apis/ilham.ilham1997/News-App/1.0.0
```

## Requirements
For running this project please install go. Go:
```
go version
go1.22.3 darwin/arm64
```

## The software you need
Recommended software to install. Example:
```
Visual Studio Code is editor code
Supabase is an open source Firebase alternative.
Insomina is an API platform for building and using APIs.
```

## Getting Started
Init folder name:
```
go mod init news-app-be23
```

## Setup supabase
```
Login supabase
Insert project
Select menu table editor
Create new schema
Rename "newsapp"
```

## Setup database:
```
Add .env
setup env value according to your database settings, for example:
poshost= is your host name
posuser= is your username
pospw= is your password
posport= is port your database
dbname= is your db name

JWT_SECRET= is your pass jwt
```

## Install library:
install several libraries needed for the project, for example:
```
go get gorm.io/gorm
go get gorm.io/driver/postgres
go get github.com/joho/godotenv
go get github.com/golang-jwt/jwt/v5
go get github.com/labstack/echo-jwt/v4
go get github.com/labstack/echo/v4
go get github.com/labstack/echo/v4/middleware
```

## Run the app:
```
open terminal
go run main.go
```

## Testing api:
```
open insomnia
add htpp request
select method POST 
testing api, for example:
http://localhost:5000/login
click send
```
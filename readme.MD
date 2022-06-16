# NiQurl - Simple URL Shortener App

**NiQurl** is simple app implementing functionalities of **Golang**, **Redis** and **Docker**. Go to section **Features** to read more about app.

## Features

NiQurl is Command Line Interface app which lets user to shorten any URL and convert it into shortURL based on NiQurl server address (currently hard-coded to `localhost:8081/[shorturl]`. After visiting address above, user is automatically redirected to original URL. If user visits address like `localhost:8081/![shorturl]`, instead of redirecting, server's API shows details about shortened URL. Length of shortened URL is configurable. 

App also works on presumption that its users are randomly generated using randomuser.me API. This presumption was made on project specification which assumes that if app user shortens URL, it should be assigned to random user from app database. When this random user was assigned to any other URL before, the app veryfies if this user can shorten URL in specified span of time. That span of time is configurable. Default value is set to 30 seconds, so if span of time does not exceed this duration, app rejects shortening URL and prompts that this user has to wait x time to be able to shorten URL again.

Because of this functionality, app cannot work properly without users and these two configurable values. That's why app checks these things before executing functionalities. During initial run, app automatically generates 5 random users, sets `SHORT_URL_LEN` value to 4 characters and `USER_WAIT_TIME` to 30 seconds.

App stores its data and variables in Redis database. 

Userdata consists of keys:
- username
- firstname
- lastname
- email
- regdate

URL data consists of keys:
- longurl
- shorturl
- createdby

Apart from these data, database also stores temporary keys named by username to verify `USER_WAIT_TIME` constraint.

## Commands

**NiQurl** CLI offers following commands:

#### make

```
make [url]
```

**make** command provides user shortening URL functionality. After invoking command, app veryfies basic url correctness and returns short URL which currently looks like `localhost:8081/[shorturl]` or warning related to `USER_WAIT_TIME` variable. If url was shortened before, make command returns short url from database.

#### setlen

```
setlen [number]
```
_number has to be between 1 and 20_

**setlen** command changes value of `SHORT_URL_LEN` if user wants to change default value set to 4

#### settime

```
settime [seconds]
```

_command accepts values between 1 and 1048576_

**settime** command changes value of `USER_WAIT_TIME` if user wants to change default value set to 30

#### settings

```
settings
```

**setttings** command changes prints currently set values of `USER_WAIT_TIME` and `SHORT_URL_LEN` variables.

#### clear

```
clear
```

**clear** command clears previous operations from terminal

#### help

```
help
```

**help** command shows all of the available commands and their description

#### exit

```
exit
```

**exit** command quits from NiQurl CLI to system shell 

## Flags

**NiQurl** CLI offers following flags:

#### generate-fake-users

```
./niqurl --generate-fake-users [number]
```
or

```
./niqurl -g [number]
```

_number has to be between 1 and 1000_

**generate-fake-users** flag used while running CLI, lets user to generate and add to database specified number of random users from randomuser.me API

## How to start

#### Requirements

The only requirement of NiQurl app is having Docker and being able to run `docker compose` command

#### Instalation

1. Save docker-compose.yml file into any location in your computer
   Link: https://github.com/nicolebroyak/niqurl/raw/dev/deployments/remote/docker-compose.yml 

   If you are using linux you can use following command:
   
```
wget https://github.com/nicolebroyak/niqurl/raw/dev/deployments/remote/docker-compose.yml
```

2. Use docker compose to build app. Run following command in the location where you have downloaded docker-compose.yml file.

```
docker compose build
```

3. Use docker compose to run app.

```
docker compose up -d
```

After that you should see app running in your Docker Containers list. App consists of three containers.

#### Run CLI

To run use command:

```
docker exec -it niqurl-cli sh -c "./niqurl"
```

Now you should be see app shell like: `NiQurl> `. If so, you can start using app and its commands.

#### Use API

After you will have some of urls shortened, you can confirm that app is working by curling short url address, for example: `localhost:8081/F48n` or `localhost:8081/!F48n`

You should see appropriate response from the server.

#### Query database

If you would like to query data from database, use command:

```
docker exec -it niqurl-redis sh -c "redis-cli"
```

You should see shell of redis client like `127.0.0.1:6379>`

## What next

If you encounter any problems using NiQurl, you can issue them in this repository. 


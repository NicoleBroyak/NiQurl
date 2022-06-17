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
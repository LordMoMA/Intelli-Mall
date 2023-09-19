# Intelli-Mall

## High-level view of the components
<img width="712" alt="Screenshot 2023-09-09 at 11 25 09 PM" src="https://github.com/LordMoMA/Intelli-Mall/assets/67067729/097cfc9d-8d45-48fc-afda-a052d03eb96a">

## How to start the application

Starting the monolith:
```bash
docker compose --profile monolith up
```
Starting the microservices

```bash
docker compose --profile microservices up  
```

> Note: my local machine is Mac M2 ARM64, be sure to locate the docker image with the tag version compatible with your machine architecture.

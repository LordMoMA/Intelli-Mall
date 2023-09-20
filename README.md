# Intelli-Mall

## High-level view of the components

![Intelli-Mall Architecture](https://github.com/LordMoMA/Intelli-Mall/assets/67067729/097cfc9d-8d45-48fc-afda-a052d03eb96a.png)

## Intelli-Mall AWS Architecture

![Intelli-Mall AWS Architecture](https://github.com/LordMoMA/Intelli-Mall/assets/67067729/176f52d6-8bf4-4bb4-bdb6-15ea9ef5a836.png)

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

## Docker Compose with either a monolith or microservices

![Screenshot of Intelli-Mall](https://github.com/LordMoMA/Intelli-Mall/assets/67067729/dec1b2ff-57a5-4966-80b8-7a1e74ad748f.png)

## Swagger UI

![Screenshot of Intelli-Mall](https://github.com/LordMoMA/Intelli-Mall/assets/67067729/1e4a4341-4d80-4161-b8a0-cd08b2c7712d.png)

## The monitoring services

![Screenshot of Intelli-Mall](https://github.com/LordMoMA/Intelli-Mall/assets/67067729/5f0d72d1-eb6a-4ce4-b842-2f6d4bc50744.png)

## Use `/cmd/busywork` to simulate several users making requests to perform several different activities:

```bash
cd cmd/busywork
go run .
```

### Busywork Output

```bash
07:54:29.751693 [Client 4] is considering setting up a new store
07:54:29.751713 [Client 5] is considering adding new inventory
07:54:29.751706 [Client 1] is considering adding new inventory
07:54:29.751739 [Client 3] is considering buying some things
07:54:29.751721 [Client 2] is considering updating product branding
07:54:29.768769 [Client 5] has no store to work with
07:54:29.779070 [Client 1] has no store to work with
07:54:30.441748 [Client 4] is getting "Phoenix Moore" ready
07:54:30.994412 [Client 4] is adding "Rustic Wooden Chicken" for $5.49
07:54:31.689962 [Client 4] is adding "Luxurious Rubber Chips" for $8.64
07:54:32.492293 [Client 3] but not enough store are available
07:54:32.568757 [Client 2] has rebranded the product "Luxurious Rubber Chips" to "Generic Wooden Chicken"
07:54:32.746445 [Client 4] is adding "Tasty Bronze Computer" for $8.35
07:54:32.757783 [Client 4] has finished setting up "Phoenix Moore"
07:54:38.288137 [Client 1] is considering adding new inventory
07:54:40.098567 [Client 1] is adding "Fantastic Wooden Mouse" for $6.39
07:54:40.104955 [Client 1] is done adding new inventory
07:54:40.544104 [Client 5] is considering updating product branding
07:54:41.231602 [Client 3] is considering adding new inventory
07:54:42.087757 [Client 2] is considering setting up a new store
07:54:42.313003 [Client 3] is adding "Rustic Wooden Shirt" for $9.26
07:54:43.249574 [Client 2] is getting "William Connelly" ready
07:54:43.677610 [Client 4] is considering adding new inventory
07:54:43.772456 [Client 5] has rebranded the product "Generic Wooden Chicken" to "Recycled Fresh Sausages"
07:54:44.054260 [Client 3] is adding "Licensed Steel Chips" for $8.13
07:54:44.062578 [Client 3] is done adding new inventory
07:54:44.535895 [Client 2] is adding "Handcrafted Concrete Table" for $10.09
07:54:45.595938 [Client 4] is adding "Elegant Granite Tuna" for $7.78
07:54:46.461319 [Client 2] is adding "Fantastic Bronze Pants" for $9.73
07:54:47.416518 [Client 4] is adding "Fantastic Plastic Chair" for $5.55
07:54:47.749307 [Client 2] is adding "Rustic Wooden Pizza" for $11.96
07:54:48.612844 [Client 1] is considering buying some things
07:54:48.950366 [Client 2] is adding "Elegant Steel Soap" for $7.67
07:54:48.962393 [Client 2] has finished setting up "William Connelly"
07:54:49.010368 [Client 4] is adding "Gorgeous Rubber Chair" for $9.40
07:54:49.017702 [Client 4] is done adding new inventory
07:54:52.326077 [Client 1] is browsing the items from "William Connelly"
07:54:52.802885 [Client 3] is considering adding new inventory
07:54:53.276149 [Client 1] might buy 1 "Rustic Wooden Pizza" for $11.96 each
07:54:53.926178 [Client 3] is adding "Handcrafted Soft Car" for $5.01
07:54:54.201275 [Client 1] is OK with $11.96
07:54:54.548199 [Client 5] is considering browsing for new things
07:54:55.889883 [Client 3] is adding "Refined Plastic Pants" for $7.80
07:54:55.896221 [Client 3] is done adding new inventory
07:54:56.755084 [Client 5] is browsing the items from "Phoenix Moore"
07:54:58.220493 [Client 5] might buy 2 "Fantastic Wooden Mouse" for $6.39 each
07:54:58.481385 [Client 2] is considering registering a new account
07:54:59.303871 [Client 5] thinks $12.78 is too much
07:54:59.937482 [Client 4] is considering buying some things
07:55:01.991065 [Client 4] is browsing the items from "William Connelly"
07:55:02.797352 [Client 4] might buy 3 "Handcrafted Soft Car" for $5.01 each
07:55:03.964219 [Client 1] is considering browsing for new things
07:55:04.630438 [Client 4] is OK with $15.03
07:55:04.635503 [Client 3] is considering registering a new account
07:55:06.397583 [Client 1] but not enough stores are available
07:55:09.474708 [Client 2] is considering browsing for new things
07:55:10.081959 [Client 5] is considering registering a new account
07:55:12.793284 [Client 2] but not enough stores are available
07:55:14.904765 [Client 1] is considering browsing for new things
07:55:14.924097 [Client 3] is considering rebranding a store
07:55:16.608062 [Client 4] is considering adding new inventory
07:55:16.807551 [Client 3] has rebranded the store "Phoenix Moore" to "Connor Williamson"
07:55:17.474107 [Client 4] is adding "Ergonomic Concrete Car" for $5.83
07:55:17.670906 [Client 1] but not enough stores are available
07:55:18.251951 [Client 4] is adding "Rustic Concrete Pants" for $7.37
07:55:18.257743 [Client 4] is done adding new inventory
07:55:21.749334 [Client 5] is considering adding new inventory
07:55:22.311291 [Client 2] is considering browsing for new things
07:55:22.528879 [Client 5] is adding "Small Soft Towels" for $9.46
07:55:23.540356 [Client 5] is adding "Electronic Fresh Soap" for $7.25
07:55:23.547253 [Client 5] is done adding new inventory
07:55:24.125859 [Client 2] is browsing the items from "William Connelly"
07:55:24.817106 [Client 2] might buy 4 "Ergonomic Concrete Car" for $5.83 each
07:55:25.547038 [Client 3] is considering browsing for new things
07:55:26.179075 [Client 1] is considering adding new inventory
07:55:26.774182 [Client 2] is browsing the items from "Connor Williamson"
07:55:27.706740 [Client 1] is adding "Awesome Wooden Keyboard" for $9.15
07:55:27.714322 [Client 1] is done adding new inventory
07:55:27.947808 [Client 3] but not enough stores are available
07:55:28.284000 [Client 2] might buy 1 "Rustic Wooden Shirt" for $9.26 each
07:55:29.177746 [Client 4] is considering browsing for new things
07:55:29.593234 [Client 2] thinks $32.58 is too much
07:55:31.675344 [Client 4] but not enough stores are available
07:55:34.322599 [Client 5] is considering registering a new account
07:55:36.221473 [Client 1] is considering adding new inventory
07:55:36.687106 [Client 3] is considering registering a new account
07:55:37.281486 [Client 1] is adding "Refined Wooden Computer" for $6.76
07:55:38.797600 [Client 1] is adding "Oriental Granite Keyboard" for $8.81
07:55:39.115718 [Client 2] is considering registering a new account
07:55:40.790283 [Client 1] is adding "Unbranded Steel Chair" for $8.65
07:55:40.797666 [Client 1] is done adding new inventory
07:55:42.595664 [Client 4] is considering adding new inventory
07:55:43.460873 [Client 4] is adding "Rustic Rubber Fish" for $9.26
07:55:44.069827 [Client 4] is adding "Licensed Frozen Pants" for $11.21
07:55:45.709748 [Client 5] is considering browsing for new things
07:55:45.721676 [Client 4] is adding "Practical Metal Towels" for $6.27
07:55:45.729938 [Client 4] is done adding new inventory
07:55:46.598130 [Client 3] is considering adding new inventory
^C07:55:46.856464 [Client 4] Quitting time
07:55:46.856470 [Client 1] Quitting time
07:55:46.856466 [Client 2] Quitting time
07:55:47.884613 [Client 5] is browsing the items from "William Connelly"
07:55:48.285565 [Client 3] is adding "Incredible Granite Chips" for $10.04
07:55:49.448966 [Client 3] is adding "Handmade Bronze Chicken" for $6.83
07:55:49.651385 [Client 5] might buy 3 "Rustic Concrete Pants" for $7.37 each
07:55:50.290852 [Client 5] thinks $22.11 is too much
07:55:50.297213 [Client 5] Quitting time
07:55:50.394300 [Client 3] is adding "Intelligent Rubber Shirt" for $10.36
07:55:50.400688 [Client 3] is done adding new inventory
07:55:50.400713 [Client 3] Quitting time
07:55:50 busywork shutdown
```

You can increase the number of clients by passing in the -clients=n flag, with an upper limit of 25.

## The Jaeger UI for tracing

Open http:// localhost:8081 in your browser to open Jaeger.

![Screenshot of Intelli-Mall](https://github.com/LordMoMA/Intelli-Mall/assets/67067729/9d1c3d63-da20-46ae-ad69-396fbbb7c350 "Screenshot of Intelli-Mall")

## Traces that involved the baskets service

![Screenshot of the Jaeger UI for tracing](https://github.com/LordMoMA/Intelli-Mall/assets/67067729/8bed563c-b362-425a-8532-b324a8a1ad8b "Screenshot of the Jaeger UI for tracing")

## Viewing the monitoring data

![Screenshot of the Jaeger UI for tracing](https://github.com/LordMoMA/Intelli-Mall/assets/67067729/fe008c04-aca7-4189-a204-3408484e6d02 "Screenshot of the Jaeger UI for tracing")

Clicking on one of the rows in the graph will provide you with additional details. 

## The Prometheus UI

We also have the metrics to check out in Prometheus at `http://localhost:9090`

![Screenshot of received messages counts for the cosec service](https://github.com/LordMoMA/Intelli-Mall/assets/67067729/231d151b-cd4c-4978-b2eb-519452832f44 "Screenshot of received messages counts for the cosec service")

## Searching for the received messages counts for the cosec service

![Screenshot of received messages counts for the cosec service](https://github.com/LordMoMA/Intelli-Mall/assets/67067729/790cdfd8-405a-4d66-a79d-ee5670724be0 "Screenshot of received messages counts for the cosec service")
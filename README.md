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

## Grafana UI for more compelling intepretation - Intelli Mall App Dashboard

Opening `localhost:3000/` and then browsing for dashboards will show the two dashboards that are installed under the intellimall folder.

![Screenshot of OpenTelemetry Collector dashboard](https://github.com/LordMoMA/Intelli-Mall/assets/67067729/03e64fbc-6eff-40ee-a89b-9e3d8a13771f "Screenshot of OpenTelemetry Collector dashboard")

How much activity you see in the dashboard will depend on how many clients you have running in the busywork application and the random interactions that the clients are performing.

## OpenTelemetry Collector dashboard

![Screenshot of OpenTelemetry Collector dashboard](https://github.com/LordMoMA/Intelli-Mall/assets/67067729/7393a06d-de86-4575-b0d6-2359a5deb957 "Screenshot of OpenTelemetry Collector dashboard")

Details about how much work the collector is doing.

## How does Terminal reflect the events:

```bash
grafana        | logger=context traceID=00000000000000000000000000000000 userId=0 orgId=1 uname= t=2023-09-20T12:06:26.480212513Z level=info msg="Request Completed" method=POST path=/api/ds/query status=400 remote_addr=172.18.0.1 time_ms=21 duration=21.768709ms size=99 referer="http://localhost:3000/d/Pc9ixd4Vk/application?orgId=1&refresh=30s" traceID=00000000000000000000000000000000
collector      | 2023-09-20T12:13:41.244Z       info    TracesExporter  {"kind": "exporter", "data_type": "traces", "name": "logging", "#spans": 2}
grafana        | logger=live t=2023-09-20T12:18:33.088289544Z level=info msg="Initialized channel handler" channel=grafana/dashboard/uid/BKf2sowmj address=grafana/dashboard/uid/BKf2sowmj
prometheus     | ts=2023-09-20T13:04:51.975Z caller=compact.go:519 level=info component=tsdb msg="write block" mint=1695204291419 maxt=1695211200000 ulid=01HASB306VTMA1K6NRP5ZCCEQ3 duration=44.237792ms
```

## Business logic flow

### read after write

![read after write](docs/Diagrams/read_after_write.png)

### transactions

![transactions](docs/Diagrams/transactions.png)

### notification ordering

![notification ordering](docs/Diagrams/notification_ordering.png)

### adding items

![adding items](docs/Diagrams/adding_items.png)

### async pay invoice

![async pay invoice](docs/Diagrams/async_pay_invoice.png)


### create order with domain events

![create order with domain events](docs/Diagrams/create_order_with_domain_events.png)

<!-- ### deduplication flow

![deduplication flow](docs/Diagrams/deduplication_flow.png) -->

## Structure of the project

```bash
├── LICENSE
├── Makefile
├── README.md
├── baskets
│   ├── basketsclient
│   │   ├── basket
│   │   │   ├── basket_client.go
│   │   │   ├── cancel_basket_parameters.go
│   │   │   ├── cancel_basket_responses.go
│   │   │   ├── checkout_basket_parameters.go
│   │   │   ├── checkout_basket_responses.go
│   │   │   ├── get_basket_parameters.go
│   │   │   ├── get_basket_responses.go
│   │   │   ├── start_basket_parameters.go
│   │   │   └── start_basket_responses.go
│   │   ├── item
│   │   │   ├── add_item_parameters.go
│   │   │   ├── add_item_responses.go
│   │   │   ├── item_client.go
│   │   │   ├── remove_item_parameters.go
│   │   │   └── remove_item_responses.go
│   │   ├── models
│   │   │   ├── add_item_params_body.go
│   │   │   ├── basketspb_add_item_response.go
│   │   │   ├── basketspb_basket.go
│   │   │   ├── basketspb_cancel_basket_response.go
│   │   │   ├── basketspb_checkout_basket_response.go
│   │   │   ├── basketspb_get_basket_response.go
│   │   │   ├── basketspb_item.go
│   │   │   ├── basketspb_remove_item_response.go
│   │   │   ├── basketspb_start_basket_request.go
│   │   │   ├── basketspb_start_basket_response.go
│   │   │   ├── checkout_basket_params_body.go
│   │   │   ├── protobuf_any.go
│   │   │   ├── remove_item_params_body.go
│   │   │   └── rpc_status.go
│   │   └── shopping_baskets_client.go
│   ├── basketspb
│   │   ├── api.pb.go
│   │   ├── api.pb.gw.go
│   │   ├── api.proto
│   │   ├── api_grpc.pb.go
│   │   ├── events.go
│   │   ├── events.pb.go
│   │   ├── events.proto
│   │   ├── mock_basket_service_client.go
│   │   ├── mock_basket_service_server.go
│   │   └── mock_unsafe_basket_service_server.go
│   ├── buf.gen.yaml
│   ├── buf.yaml
│   ├── cmd
│   │   └── service
│   │       └── main.go
│   ├── generate.go
│   ├── internal
│   │   ├── application
│   │   │   ├── application.go
│   │   │   ├── application_test.go
│   │   │   ├── instrumented_app.go
│   │   │   └── mock_app.go
│   │   ├── constants
│   │   │   └── constants.go
│   │   ├── domain
│   │   │   ├── basket.go
│   │   │   ├── basket_events.go
│   │   │   ├── basket_repository.go
│   │   │   ├── basket_snapshots.go
│   │   │   ├── basket_status.go
│   │   │   ├── basket_test.go
│   │   │   ├── fake_basket_repository.go
│   │   │   ├── fake_product_cache_repository.go
│   │   │   ├── fake_store_cache_repository.go
│   │   │   ├── item.go
│   │   │   ├── mock_basket_repository.go
│   │   │   ├── mock_product_cache_repository.go
│   │   │   ├── mock_product_repository.go
│   │   │   ├── mock_store_cache_repository.go
│   │   │   ├── mock_store_repository.go
│   │   │   ├── product.go
│   │   │   ├── product_cache_repository.go
│   │   │   ├── product_repository.go
│   │   │   ├── registrations.go
│   │   │   ├── store.go
│   │   │   ├── store_cache_repository.go
│   │   │   └── store_repository.go
│   │   ├── grpc
│   │   │   ├── product_repository.go
│   │   │   ├── server.go
│   │   │   ├── server_integration_test.go
│   │   │   ├── server_transaction.go
│   │   │   └── store_repository.go
│   │   ├── handlers
│   │   │   ├── domain_events.go
│   │   │   ├── domain_events_transaction.go
│   │   │   ├── integration_events.go
│   │   │   ├── integration_events_contract_test.go
│   │   │   ├── integration_events_integration_test.go
│   │   │   ├── integration_events_transaction.go
│   │   │   └── pacts
│   │   │       └── baskets-sub-stores-pub.json
│   │   ├── postgres
│   │   │   ├── product_cache_repository.go
│   │   │   ├── product_cache_repository_integration_test.go
│   │   │   └── store_cache_repository.go
│   │   └── rest
│   │       ├── api.annotations.yaml
│   │       ├── api.openapi.yaml
│   │       ├── api.swagger.json
│   │       ├── gateway.go
│   │       ├── gateway_contract_test.go
│   │       ├── index.html
│   │       └── swagger.go
│   ├── migrations
│   │   ├── 001_create_tables.sql
│   │   └── migrations.go
│   ├── module.go
│   └── ui
│       ├── client.js
│       ├── client.spec.js
│       ├── package-lock.json
│       ├── package.json
│       └── pacts
│           └── baskets-ui-baskets-api.json
├── cmd
│   ├── busywork
│   │   ├── baskets
│   │   │   └── client.go
│   │   ├── client.go
│   │   ├── customers
│   │   │   └── client.go
│   │   ├── main.go
│   │   ├── payments
│   │   │   └── client.go
│   │   └── stores
│   │       └── client.go
│   └── intellimall
│       └── main.go
├── cosec
│   ├── cmd
│   │   └── service
│   │       └── main.go
│   ├── internal
│   │   ├── constants
│   │   │   └── constants.go
│   │   ├── handlers
│   │   │   ├── integration_events.go
│   │   │   ├── integration_events_transaction.go
│   │   │   ├── replies.go
│   │   │   └── replies_transaction.go
│   │   ├── models
│   │   │   └── data.go
│   │   └── saga.go
│   ├── migrations
│   │   ├── 001_create_tables.sql
│   │   └── migrations.go
│   └── module.go
├── create_order.http
├── customers
│   ├── buf.gen.yaml
│   ├── buf.yaml
│   ├── cmd
│   │   └── service
│   │       └── main.go
│   ├── customersclient
│   │   ├── customer
│   │   │   ├── change_sms_number_parameters.go
│   │   │   ├── change_sms_number_responses.go
│   │   │   ├── customer_client.go
│   │   │   ├── disable_customer_parameters.go
│   │   │   ├── disable_customer_responses.go
│   │   │   ├── enable_customer_parameters.go
│   │   │   ├── enable_customer_responses.go
│   │   │   ├── get_customer_parameters.go
│   │   │   ├── get_customer_responses.go
│   │   │   ├── register_customer_parameters.go
│   │   │   └── register_customer_responses.go
│   │   ├── customers_client.go
│   │   └── models
│   │       ├── change_sms_number_params_body.go
│   │       ├── customerspb_change_sms_number_response.go
│   │       ├── customerspb_customer.go
│   │       ├── customerspb_disable_customer_response.go
│   │       ├── customerspb_enable_customer_response.go
│   │       ├── customerspb_get_customer_response.go
│   │       ├── customerspb_register_customer_request.go
│   │       ├── customerspb_register_customer_response.go
│   │       ├── protobuf_any.go
│   │       └── rpc_status.go
│   ├── customerspb
│   │   ├── api.pb.go
│   │   ├── api.pb.gw.go
│   │   ├── api.proto
│   │   ├── api_grpc.pb.go
│   │   ├── messages.go
│   │   ├── messages.pb.go
│   │   ├── messages.proto
│   │   ├── mock_customers_service_client.go
│   │   ├── mock_customers_service_server.go
│   │   └── mock_unsafe_customers_service_server.go
│   ├── generate.go
│   ├── internal
│   │   ├── application
│   │   │   ├── application.go
│   │   │   ├── instrumented_app.go
│   │   │   └── mock_app.go
│   │   ├── constants
│   │   │   └── constants.go
│   │   ├── domain
│   │   │   ├── customer.go
│   │   │   ├── customer_events.go
│   │   │   ├── customer_repository.go
│   │   │   └── mock_customer_repository.go
│   │   ├── grpc
│   │   │   ├── server.go
│   │   │   └── server_transaction.go
│   │   ├── handlers
│   │   │   ├── commands.go
│   │   │   ├── commands_transaction.go
│   │   │   ├── domain_events.go
│   │   │   └── domain_events_transaction.go
│   │   ├── postgres
│   │   │   └── customer_repository.go
│   │   └── rest
│   │       ├── api.annotations.yaml
│   │       ├── api.openapi.yaml
│   │       ├── api.swagger.json
│   │       ├── gateway.go
│   │       ├── index.html
│   │       └── swagger.go
│   ├── migrations
│   │   ├── 001_create_tables.sql
│   │   └── migrations.go
│   └── module.go
├── deployment
│   ├── application
│   │   ├── Makefile
│   │   ├── database.tf
│   │   ├── kubernetes.tf
│   │   ├── main.tf
│   │   ├── nats.tf
│   │   ├── providers.tf
│   │   ├── sql
│   │   │   ├── init_db.psql
│   │   │   └── init_service_db.psql
│   │   ├── svc_baskets.tf
│   │   ├── svc_cosec.tf
│   │   ├── svc_customers.tf
│   │   ├── svc_depot.tf
│   │   ├── svc_notifications.tf
│   │   ├── svc_ordering.tf
│   │   ├── svc_payments.tf
│   │   ├── svc_search.tf
│   │   └── svc_stores.tf
│   ├── infrastructure
│   │   ├── Makefile
│   │   ├── alb.tf
│   │   ├── ecr.tf
│   │   ├── eks.tf
│   │   ├── main.tf
│   │   ├── providers.tf
│   │   ├── rds.tf
│   │   ├── security_groups.tf
│   │   └── vpc.tf
│   └── setup-tools
│       ├── Dockerfile
│       ├── set-tool-alias.sh
│       └── win-set-tool-alias.ps1
├── depot
│   ├── buf.gen.yaml
│   ├── buf.yaml
│   ├── cmd
│   │   └── service
│   │       └── main.go
│   ├── depotclient
│   │   ├── depot_operations_client.go
│   │   ├── models
│   │   │   ├── assign_shopping_list_params_body.go
│   │   │   ├── depotpb_assign_shopping_list_response.go
│   │   │   ├── depotpb_cancel_shopping_list_response.go
│   │   │   ├── depotpb_complete_shopping_list_response.go
│   │   │   ├── depotpb_create_shopping_list_request.go
│   │   │   ├── depotpb_create_shopping_list_response.go
│   │   │   ├── depotpb_order_item.go
│   │   │   ├── protobuf_any.go
│   │   │   └── rpc_status.go
│   │   └── shopping_list
│   │       ├── assign_shopping_list_parameters.go
│   │       ├── assign_shopping_list_responses.go
│   │       ├── cancel_shopping_list_parameters.go
│   │       ├── cancel_shopping_list_responses.go
│   │       ├── ccomplete_shopping_list_parameters.go
│   │       ├── ccomplete_shopping_list_responses.go
│   │       ├── create_shopping_list_parameters.go
│   │       ├── create_shopping_list_responses.go
│   │       └── shopping_list_client.go
│   ├── depotpb
│   │   ├── api.pb.go
│   │   ├── api.pb.gw.go
│   │   ├── api.proto
│   │   ├── api_grpc.pb.go
│   │   ├── messages.go
│   │   ├── messages.pb.go
│   │   ├── messages.proto
│   │   ├── mock_depot_service_client.go
│   │   ├── mock_depot_service_server.go
│   │   └── mock_unsafe_depot_service_server.go
│   ├── generate.go
│   ├── internal
│   │   ├── application
│   │   │   ├── application.go
│   │   │   ├── commands
│   │   │   │   ├── assign_shopping_list.go
│   │   │   │   ├── cancel_shopping_list.go
│   │   │   │   ├── complete_shopping_list.go
│   │   │   │   ├── create_shopping_list.go
│   │   │   │   ├── initiate_shopping.go
│   │   │   │   └── order_item.go
│   │   │   ├── mock_app.go
│   │   │   ├── mock_commands.go
│   │   │   ├── mock_queries.go
│   │   │   └── queries
│   │   │       └── get_shopping_list.go
│   │   ├── constants
│   │   │   └── constants.go
│   │   ├── domain
│   │   │   ├── bot.go
│   │   │   ├── bot_status.go
│   │   │   ├── fake_product_cache_repository.go
│   │   │   ├── fake_store_cache_repository.go
│   │   │   ├── item.go
│   │   │   ├── mock_order_repository.go
│   │   │   ├── mock_product_cache_repository.go
│   │   │   ├── mock_product_repository.go
│   │   │   ├── mock_shopping_list_repository.go
│   │   │   ├── mock_store_cache_repository.go
│   │   │   ├── mock_store_repository.go
│   │   │   ├── order_repository.go
│   │   │   ├── product.go
│   │   │   ├── product_cache_repository.go
│   │   │   ├── product_repository.go
│   │   │   ├── shopping_list.go
│   │   │   ├── shopping_list_events.go
│   │   │   ├── shopping_list_repository.go
│   │   │   ├── shopping_list_status.go
│   │   │   ├── stop.go
│   │   │   ├── store.go
│   │   │   ├── store_cache_repository.go
│   │   │   └── store_repository.go
│   │   ├── grpc
│   │   │   ├── product_repository.go
│   │   │   ├── server.go
│   │   │   ├── server_transaction.go
│   │   │   └── store_repository.go
│   │   ├── handlers
│   │   │   ├── commands.go
│   │   │   ├── commands_transaction.go
│   │   │   ├── domain_events.go
│   │   │   ├── domain_events_transaction.go
│   │   │   ├── integration_events.go
│   │   │   ├── integration_events_contract_test.go
│   │   │   ├── integration_events_transaction.go
│   │   │   └── pacts
│   │   │       └── depot-sub-stores-pub.json
│   │   ├── postgres
│   │   │   ├── product_cache_repository.go
│   │   │   ├── shopping_list_repository.go
│   │   │   └── store_cache_repository.go
│   │   └── rest
│   │       ├── api.annotations.yaml
│   │       ├── api.openapi.yaml
│   │       ├── api.swagger.json
│   │       ├── gateway.go
│   │       ├── index.html
│   │       └── swagger.go
│   ├── migrations
│   │   ├── 001_create_tables.sql
│   │   └── migrations.go
│   └── module.go
├── docker
│   ├── Dockerfile
│   ├── Dockerfile.microservices
│   ├── database
│   │   ├── 001_create_monolith_db.sh
│   │   ├── 002_create_pact_db.sh
│   │   └── 003_create_service_dbs.sh
│   ├── grafana
│   │   ├── grafana.ini
│   │   └── provisioning
│   │       ├── dashboards
│   │       │   ├── default.yaml
│   │       │   ├── intellimall.json
│   │       │   └── opentelemetry-collector.json
│   │       └── datasources
│   │           ├── default.yaml
│   │           └── jaeger.yaml
│   ├── nginx.conf
│   ├── otel
│   │   └── otel-config.yml
│   ├── prometheus
│   │   └── prometheus-config.yml
│   └── wait-for
├── docker-compose.yml
├── docs
│   ├── ADL
│   │   ├── 0001-keep-an-architecture-decision-log.md
│   │   └── 0002-use-a-modular-monolith-architecture.md
│   ├── Diagrams
│   │   ├── 2pc_transaction.png
│   │   ├── adding_items.plantuml
│   │   ├── adding_items.png
│   │   ├── async_pay_invoice.png
│   │   ├── create_order_with_domain_events.plantuml
│   │   ├── create_order_with_domain_events.png
│   │   ├── deduplication_flow.png
│   │   ├── deduplication_flow.puml
│   │   ├── local_transaction.png
│   │   ├── notification_ordering.png
│   │   ├── notification_ordering.puml
│   │   ├── pay_invoice.png
│   │   ├── read_after_write.png
│   │   ├── read_after_write.puml
│   │   ├── transactions.png
│   │   └── transactions.puml
│   └── EventStorming
│       └── BigPicture
│           ├── 2_chaotic_exploration.png
│           ├── 3_enforcing_the_timeline.png
│           ├── 4_people_and_systems.png
│           ├── 5_explicit_walkthrough.png
│           └── bounded_contexts.png
├── go.mod
├── go.sum
├── internal
│   ├── am
│   │   ├── buf.gen.yaml
│   │   ├── buf.yaml
│   │   ├── command_messages.go
│   │   ├── event_messages.go
│   │   ├── fake_event_publisher.go
│   │   ├── generate.go
│   │   ├── message.go
│   │   ├── message_types.pb.go
│   │   ├── message_types.proto
│   │   ├── middleware.go
│   │   ├── mock_command_publisher.go
│   │   ├── mock_event_publisher.go
│   │   ├── mock_reply_publisher.go
│   │   ├── reply_messages.go
│   │   ├── subscriber_config.go
│   │   └── subscription.go
│   ├── amotel
│   │   ├── extractor.go
│   │   ├── injector.go
│   │   ├── metadata_carrier.go
│   │   └── trace.go
│   ├── amprom
│   │   ├── received.go
│   │   └── sent.go
│   ├── config
│   │   └── config.go
│   ├── ddd
│   │   ├── aggregate.go
│   │   ├── aggregate_build_options.go
│   │   ├── command.go
│   │   ├── entity.go
│   │   ├── entity_build_options.go
│   │   ├── event.go
│   │   ├── event_dispatcher.go
│   │   ├── generate.go
│   │   ├── metadata.go
│   │   ├── mock_aggregate.go
│   │   ├── mock_command_handler.go
│   │   ├── mock_entity.go
│   │   ├── mock_event_handler.go
│   │   ├── mock_event_publisher.go
│   │   ├── mock_event_subscriber.go
│   │   ├── mock_reply_handler.go
│   │   └── reply.go
│   ├── di
│   │   ├── api.go
│   │   ├── container.go
│   │   └── tracked.go
│   ├── errorsotel
│   │   └── errors.go
│   ├── es
│   │   ├── aggregate.go
│   │   ├── aggregate_build_options.go
│   │   ├── aggregate_repository.go
│   │   ├── aggregate_store.go
│   │   ├── event.go
│   │   ├── event_publisher.go
│   │   ├── fake_aggregate_repository.go
│   │   ├── generate.go
│   │   ├── mock_aggregate.go
│   │   ├── mock_aggregate_repository.go
│   │   ├── mock_aggregate_store.go
│   │   ├── mock_event_sourced_aggregate.go
│   │   └── snapshot.go
│   ├── jetstream
│   │   ├── buf.gen.yaml
│   │   ├── buf.yaml
│   │   ├── generate.go
│   │   ├── raw_message.go
│   │   ├── stream.go
│   │   ├── stream_message.pb.go
│   │   ├── stream_message.proto
│   │   └── subscription.go
│   ├── logger
│   │   ├── log
│   │   │   └── silent_logger.go
│   │   └── logger.go
│   ├── postgres
│   │   ├── event_store.go
│   │   ├── inbox_store.go
│   │   ├── outbox_store.go
│   │   ├── saga_store.go
│   │   ├── snapshot_store.go
│   │   └── types.go
│   ├── postgresotel
│   │   └── trace.go
│   ├── registry
│   │   ├── build_options.go
│   │   ├── errors.go
│   │   ├── registration.go
│   │   ├── registry.go
│   │   ├── serde.go
│   │   └── serdes
│   │       ├── doc.go
│   │       ├── json.go
│   │       └── proto.go
│   ├── rpc
│   │   ├── config.go
│   │   └── conn.go
│   ├── sec
│   │   ├── orchestrator.go
│   │   ├── saga.go
│   │   ├── saga_repository.go
│   │   └── saga_step.go
│   ├── system
│   │   ├── system.go
│   │   └── types.go
│   ├── tm
│   │   ├── inbox_middleware.go
│   │   ├── outbox_middleware.go
│   │   └── outbox_processor.go
│   ├── tools.go
│   ├── waiter
│   │   ├── options.go
│   │   └── waiter.go
│   └── web
│       ├── config.go
│       ├── embed.go
│       ├── index.html
│       └── swagger-ui
│           ├── favicon-16x16.png
│           ├── favicon-32x32.png
│           ├── index.html
│           ├── oauth2-redirect.html
│           ├── swagger-ui-bundle.js
│           ├── swagger-ui-es-bundle-core.js
│           ├── swagger-ui-es-bundle.js
│           ├── swagger-ui-standalone-preset.js
│           ├── swagger-ui.css
│           └── swagger-ui.js
├── migrations
│   ├── 001_create_basket_schema.sql
│   ├── 002_create_cosec_schema.sql
│   ├── 003_create_customers_schema.sql
│   ├── 004_create_depot_schema.sql
│   ├── 005_create_notifications_schema.sql
│   ├── 006_create_ordering_schema.sql
│   ├── 007_create_payments_schema.sql
│   ├── 008_create_search_schema.sql
│   ├── 009_create_stores_schema.sql
│   └── migrations.go
├── notifications
│   ├── buf.gen.yaml
│   ├── buf.yaml
│   ├── cmd
│   │   └── service
│   │       └── main.go
│   ├── generate.go
│   ├── internal
│   │   ├── application
│   │   │   ├── application.go
│   │   │   ├── customer_cache_repository.go
│   │   │   └── customer_repository.go
│   │   ├── constants
│   │   │   └── constants.go
│   │   ├── grpc
│   │   │   ├── customer_repository.go
│   │   │   └── server.go
│   │   ├── handlers
│   │   │   └── integration_events.go
│   │   ├── models
│   │   │   └── customer.go
│   │   └── postgres
│   │       └── customer_cache_repository.go
│   ├── migrations
│   │   ├── 001_create_tables.sql
│   │   └── migrations.go
│   ├── module.go
│   └── notificationspb
│       ├── api.pb.go
│       ├── api.proto
│       └── api_grpc.pb.go
├── ordering
│   ├── buf.gen.yaml
│   ├── buf.yaml
│   ├── cmd
│   │   └── service
│   │       └── main.go
│   ├── generate.go
│   ├── internal
│   │   ├── application
│   │   │   ├── application.go
│   │   │   ├── commands
│   │   │   │   ├── approve_order.go
│   │   │   │   ├── cancel_order.go
│   │   │   │   ├── complete_order.go
│   │   │   │   ├── create_order.go
│   │   │   │   ├── ready_order.go
│   │   │   │   └── reject_order.go
│   │   │   ├── mock_app.go
│   │   │   ├── mock_commands.go
│   │   │   ├── mock_queries.go
│   │   │   └── queries
│   │   │       └── get_order.go
│   │   ├── constants
│   │   │   └── constants.go
│   │   ├── domain
│   │   │   ├── customer_repository.go
│   │   │   ├── item.go
│   │   │   ├── mock_customer_repository.go
│   │   │   ├── mock_order_repository.go
│   │   │   ├── mock_payment_repository.go
│   │   │   ├── mock_shopping_repository.go
│   │   │   ├── order.go
│   │   │   ├── order_events.go
│   │   │   ├── order_repository.go
│   │   │   ├── order_snapshots.go
│   │   │   ├── order_status.go
│   │   │   ├── payment_repository.go
│   │   │   └── shopping_repository.go
│   │   ├── grpc
│   │   │   ├── server.go
│   │   │   └── server_transaction.go
│   │   ├── handlers
│   │   │   ├── commands.go
│   │   │   ├── commands_transaction.go
│   │   │   ├── domain_events.go
│   │   │   ├── domain_events_transaction.go
│   │   │   ├── integration_events.go
│   │   │   └── integration_events_transaction.go
│   │   └── rest
│   │       ├── api.annotations.yaml
│   │       ├── api.openapi.yaml
│   │       ├── api.swagger.json
│   │       ├── gateway.go
│   │       ├── index.html
│   │       └── swagger.go
│   ├── migrations
│   │   ├── 001_create_tables.sql
│   │   └── migrations.go
│   ├── module.go
│   ├── orderingclient
│   │   ├── models
│   │   │   ├── orderingpb_cancel_order_response.go
│   │   │   ├── orderingpb_create_order_request.go
│   │   │   ├── orderingpb_create_order_response.go
│   │   │   ├── orderingpb_get_order_response.go
│   │   │   ├── orderingpb_item.go
│   │   │   ├── orderingpb_order.go
│   │   │   ├── protobuf_any.go
│   │   │   └── rpc_status.go
│   │   ├── order
│   │   │   ├── cancel_order_parameters.go
│   │   │   ├── cancel_order_responses.go
│   │   │   ├── create_order_parameters.go
│   │   │   ├── create_order_responses.go
│   │   │   ├── get_order_parameters.go
│   │   │   ├── get_order_responses.go
│   │   │   └── order_client.go
│   │   └── order_processing_client.go
│   └── orderingpb
│       ├── api.pb.go
│       ├── api.pb.gw.go
│       ├── api.proto
│       ├── api_grpc.pb.go
│       ├── messages.go
│       ├── messages.pb.go
│       ├── messages.proto
│       ├── mock_ordering_service_client.go
│       ├── mock_ordering_service_server.go
│       └── mock_unsafe_ordering_service_server.go
├── payments
│   ├── buf.gen.yaml
│   ├── buf.yaml
│   ├── cmd
│   │   └── service
│   │       └── main.go
│   ├── generate.go
│   ├── internal
│   │   ├── application
│   │   │   ├── application.go
│   │   │   ├── invoice_repository.go
│   │   │   ├── mock_app.go
│   │   │   ├── mock_invoice_repository.go
│   │   │   ├── mock_payment_repository.go
│   │   │   └── payment_repository.go
│   │   ├── constants
│   │   │   └── constants.go
│   │   ├── grpc
│   │   │   ├── server.go
│   │   │   └── server_transaction.go
│   │   ├── handlers
│   │   │   ├── commands.go
│   │   │   ├── commands_transaction.go
│   │   │   ├── domain_events.go
│   │   │   ├── domain_events_transaction.go
│   │   │   ├── integration_events.go
│   │   │   └── integration_events_transaction.go
│   │   ├── models
│   │   │   ├── invoice.go
│   │   │   ├── invoice_events.go
│   │   │   └── payment.go
│   │   ├── postgres
│   │   │   ├── invoice_repository.go
│   │   │   └── payment_repository.go
│   │   └── rest
│   │       ├── api.annotations.yaml
│   │       ├── api.openapi.yaml
│   │       ├── api.swagger.json
│   │       ├── gateway.go
│   │       ├── index.html
│   │       └── swagger.go
│   ├── migrations
│   │   ├── 001_create_tables.sql
│   │   └── migrations.go
│   ├── module.go
│   ├── paymentsclient
│   │   ├── invoice
│   │   │   ├── invoice_client.go
│   │   │   ├── pay_invoice_parameters.go
│   │   │   └── pay_invoice_responses.go
│   │   ├── models
│   │   │   ├── paymentspb_authorize_payment_request.go
│   │   │   ├── paymentspb_authorize_payment_response.go
│   │   │   ├── paymentspb_pay_invoice_response.go
│   │   │   ├── protobuf_any.go
│   │   │   └── rpc_status.go
│   │   ├── payment
│   │   │   ├── authorize_payment_parameters.go
│   │   │   ├── authorize_payment_responses.go
│   │   │   └── payment_client.go
│   │   └── payments_client.go
│   └── paymentspb
│       ├── api.pb.go
│       ├── api.pb.gw.go
│       ├── api.proto
│       ├── api_grpc.pb.go
│       ├── messages.go
│       ├── messages.pb.go
│       ├── messages.proto
│       ├── mock_payments_service_client.go
│       ├── mock_payments_service_server.go
│       └── mock_unsafe_payments_service_server.go
├── search
│   ├── buf.gen.yaml
│   ├── buf.yaml
│   ├── cmd
│   │   └── service
│   │       └── main.go
│   ├── generate.go
│   ├── internal
│   │   ├── application
│   │   │   ├── application.go
│   │   │   ├── customer_repository.go
│   │   │   ├── order_repository.go
│   │   │   ├── product_repository.go
│   │   │   └── store_repository.go
│   │   ├── constants
│   │   │   └── constants.go
│   │   ├── grpc
│   │   │   ├── customer_repository.go
│   │   │   ├── product_repository.go
│   │   │   ├── server.go
│   │   │   ├── server_transaction.go
│   │   │   └── store_repository.go
│   │   ├── handlers
│   │   │   ├── integration_events.go
│   │   │   └── integration_events_transaction.go
│   │   ├── models
│   │   │   ├── customer.go
│   │   │   ├── order.go
│   │   │   ├── product.go
│   │   │   └── store.go
│   │   ├── postgres
│   │   │   ├── customer_cache_repository.go
│   │   │   ├── order_repository.go
│   │   │   ├── product_cache_repository.go
│   │   │   └── store_cache_repository.go
│   │   └── rest
│   │       ├── api.annotations.yaml
│   │       ├── api.openapi.yaml
│   │       ├── api.swagger.json
│   │       ├── gateway.go
│   │       ├── index.html
│   │       └── swagger.go
│   ├── migrations
│   │   ├── 001_create_tables.sql
│   │   └── migrations.go
│   ├── module.go
│   └── searchpb
│       ├── api.pb.go
│       ├── api.pb.gw.go
│       ├── api.proto
│       └── api_grpc.pb.go
├── stores
│   ├── buf.gen.yaml
│   ├── buf.yaml
│   ├── cmd
│   │   └── service
│   │       └── main.go
│   ├── generate.go
│   ├── internal
│   │   ├── application
│   │   │   ├── application.go
│   │   │   ├── commands
│   │   │   │   ├── add_product.go
│   │   │   │   ├── create_store.go
│   │   │   │   ├── decrease_product_price.go
│   │   │   │   ├── disable_participation.go
│   │   │   │   ├── enable_participation.go
│   │   │   │   ├── increase_product_price.go
│   │   │   │   ├── rebrand_product.go
│   │   │   │   ├── rebrand_store.go
│   │   │   │   └── remove_product.go
│   │   │   ├── mock_app.go
│   │   │   ├── mock_commands.go
│   │   │   ├── mock_queries.go
│   │   │   └── queries
│   │   │       ├── get_catalog.go
│   │   │       ├── get_participating_stores.go
│   │   │       ├── get_product.go
│   │   │       ├── get_store.go
│   │   │       └── get_stores.go
│   │   ├── constants
│   │   │   └── constants.go
│   │   ├── domain
│   │   │   ├── catalog_repository.go
│   │   │   ├── fake_catalog_repository.go
│   │   │   ├── fake_mall_repository.go
│   │   │   ├── fake_product_repository.go
│   │   │   ├── fake_store_repository.go
│   │   │   ├── mall_repository.go
│   │   │   ├── mock_catalog_repository.go
│   │   │   ├── mock_mall_repository.go
│   │   │   ├── mock_product_repository.go
│   │   │   ├── mock_store_repository.go
│   │   │   ├── product.go
│   │   │   ├── product_events.go
│   │   │   ├── product_repository.go
│   │   │   ├── product_snapshots.go
│   │   │   ├── store.go
│   │   │   ├── store_events.go
│   │   │   ├── store_repository.go
│   │   │   └── store_snapshots.go
│   │   ├── grpc
│   │   │   ├── server.go
│   │   │   └── server_transaction.go
│   │   ├── handlers
│   │   │   ├── catalog.go
│   │   │   ├── domain_events.go
│   │   │   ├── domain_events_contract_test.go
│   │   │   ├── domain_events_transaction.go
│   │   │   └── mall.go
│   │   ├── postgres
│   │   │   ├── catalog_repository.go
│   │   │   └── mall_repository.go
│   │   └── rest
│   │       ├── api.annotations.yaml
│   │       ├── api.openapi.yaml
│   │       ├── api.swagger.json
│   │       ├── gateway.go
│   │       ├── index.html
│   │       └── swagger.go
│   ├── migrations
│   │   ├── 001_create_tables.sql
│   │   └── migrations.go
│   ├── module.go
│   ├── storesclient
│   │   ├── models
│   │   │   ├── add_product_params_body.go
│   │   │   ├── decrease_product_price_params_body.go
│   │   │   ├── increase_product_price_params_body.go
│   │   │   ├── protobuf_any.go
│   │   │   ├── rebrand_product_params_body.go
│   │   │   ├── rebrand_store_params_body.go
│   │   │   ├── rpc_status.go
│   │   │   ├── storespb_add_product_response.go
│   │   │   ├── storespb_create_store_request.go
│   │   │   ├── storespb_create_store_response.go
│   │   │   ├── storespb_decrease_product_price_response.go
│   │   │   ├── storespb_disable_participation_response.go
│   │   │   ├── storespb_enable_participation_response.go
│   │   │   ├── storespb_get_catalog_response.go
│   │   │   ├── storespb_get_participating_stores_response.go
│   │   │   ├── storespb_get_product_response.go
│   │   │   ├── storespb_get_store_response.go
│   │   │   ├── storespb_get_stores_response.go
│   │   │   ├── storespb_increase_product_price_response.go
│   │   │   ├── storespb_product.go
│   │   │   ├── storespb_rebrand_product_response.go
│   │   │   ├── storespb_rebrand_store_response.go
│   │   │   ├── storespb_remove_product_response.go
│   │   │   └── storespb_store.go
│   │   ├── participation
│   │   │   ├── disable_participation_parameters.go
│   │   │   ├── disable_participation_responses.go
│   │   │   ├── enable_participation_parameters.go
│   │   │   ├── enable_participation_responses.go
│   │   │   ├── get_participating_stores_parameters.go
│   │   │   ├── get_participating_stores_responses.go
│   │   │   └── participation_client.go
│   │   ├── product
│   │   │   ├── add_product_parameters.go
│   │   │   ├── add_product_responses.go
│   │   │   ├── decrease_product_price_parameters.go
│   │   │   ├── decrease_product_price_responses.go
│   │   │   ├── get_product_parameters.go
│   │   │   ├── get_product_responses.go
│   │   │   ├── get_store_products_parameters.go
│   │   │   ├── get_store_products_responses.go
│   │   │   ├── increase_product_price_parameters.go
│   │   │   ├── increase_product_price_responses.go
│   │   │   ├── product_client.go
│   │   │   ├── rebrand_product_parameters.go
│   │   │   ├── rebrand_product_responses.go
│   │   │   ├── remove_product_parameters.go
│   │   │   └── remove_product_responses.go
│   │   ├── store
│   │   │   ├── create_store_parameters.go
│   │   │   ├── create_store_responses.go
│   │   │   ├── get_store_parameters.go
│   │   │   ├── get_store_responses.go
│   │   │   ├── get_stores_parameters.go
│   │   │   ├── get_stores_responses.go
│   │   │   ├── rebrand_store_parameters.go
│   │   │   ├── rebrand_store_responses.go
│   │   │   └── store_client.go
│   │   └── store_management_client.go
│   └── storespb
│       ├── api.pb.go
│       ├── api.pb.gw.go
│       ├── api.proto
│       ├── api_grpc.pb.go
│       ├── asyncapi.go
│       ├── asyncapi.yaml
│       ├── css
│       │   ├── asyncapi.min.css
│       │   └── global.min.css
│       ├── events.go
│       ├── events.pb.go
│       ├── events.proto
│       ├── index.html
│       ├── js
│       │   └── asyncapi-ui.min.js
│       ├── mock_stores_service_client.go
│       ├── mock_stores_service_server.go
│       └── mock_unsafe_stores_service_server.go
└── testing
    └── e2e
        ├── baskets_feature_test.go
        ├── customers_feature_test.go
        ├── e2e_test.go
        ├── features
        │   ├── baskets
        │   │   ├── add_item.feature
        │   │   ├── checkout_basket.feature
        │   │   └── start_basket.feature
        │   ├── customers
        │   │   └── register_customer.feature
        │   ├── kiosk
        │   │   └── shopping.feature
        │   ├── orders
        │   │   └── processing.feature
        │   └── stores
        │       ├── create_product.feature
        │       └── create_store.feature
        └── stores_feature_test.go

202 directories, 790 files
```
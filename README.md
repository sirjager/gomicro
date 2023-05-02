# GoMicro Starter Template

This repository serves as a starter template for building microservices in Go. It provides a basic structure and essential components to kickstart the development of a microservice-based application.

## Directory Structure

Here is an overview of the directory structure:

```sh
Permissions Size User Date Modified Name
drwxrwxrwx     - root  1 May 23:01   ./
drwxrwxrwx     - root  1 May 22:36  ├──  cmd/
drwxrwxrwx     - root  1 May 22:36  │  ├──  gateway/
.rwxrwxrwx  1.8k root  1 May 22:36  │  │  ├──  gateway.go*
.rwxrwxrwx   729 root  1 May 22:36  │  │  ├──  headers.go*
.rwxrwxrwx  1.0k root  1 May 22:35  │  │  └──  logger.go*
drwxrwxrwx     - root  1 May 22:35  │  ├──  grpc/
.rwxrwxrwx  1.2k root  1 May 22:35  │  │  ├──  grpc.go*
.rwxrwxrwx  1.7k root  1 May 22:35  │  │  └──  logger.go*
.rwxrwxrwx  2.0k root  1 May 22:57  │  └──  main.go*
drwxrwxrwx     - root  1 May 15:19  ├──  config/
.rwxrwxrwx  1.7k root  1 May 22:56  │  └──  config.go*
.rwxrwxrwx   953 root  1 May 22:50  ├──  docker-compose.yml*
.rwxrwxrwx   558 root  1 May 22:41  ├──  Dockerfile*
drwxrwxrwx     - root  1 May 15:14  ├──  docs/
drwxrwxrwx     - root  1 May 15:14  │  ├──  statik/
.rwxrwxrwx  7.6M root  1 May 15:14  │  │  └──  statik.go*
drwxrwxrwx     - root  1 May 15:14  │  └──  swagger/
.rwxrwxrwx   665 root  1 May 13:42  │     ├──  favicon-16x16.png*
.rwxrwxrwx   628 root  1 May 13:42  │     ├──  favicon-32x32.png*
.rwxrwxrwx  3.8k root  1 May 15:14  │     ├──  gomicro.swagger.json*
.rwxrwxrwx   202 root  1 May 13:42  │     ├──  index.css*
.rwxrwxrwx   734 root  1 May 13:42  │     ├──  index.html*
.rwxrwxrwx  2.7k root  1 May 13:42  │     ├──  oauth2-redirect.html*
.rwxrwxrwx   488 root  1 May 13:42  │     ├──  swagger-initializer.js*
.rwxrwxrwx  1.0M root  1 May 13:42  │     ├──  swagger-ui-bundle.js*
.rwxrwxrwx  1.5M root  1 May 13:42  │     ├──  swagger-ui-bundle.js.map*
.rwxrwxrwx  369k root  1 May 13:42  │     ├──  swagger-ui-es-bundle-core.js*
.rwxrwxrwx  1.3M root  1 May 13:42  │     ├──  swagger-ui-es-bundle-core.js.map*
.rwxrwxrwx  1.0M root  1 May 13:42  │     ├──  swagger-ui-es-bundle.js*
.rwxrwxrwx  1.5M root  1 May 13:42  │     ├──  swagger-ui-es-bundle.js.map*
.rwxrwxrwx  323k root  1 May 13:42  │     ├──  swagger-ui-standalone-preset.js*
.rwxrwxrwx  518k root  1 May 13:42  │     ├──  swagger-ui-standalone-preset.js.map*
.rwxrwxrwx  145k root  1 May 13:42  │     ├──  swagger-ui.css*
.rwxrwxrwx  251k root  1 May 13:42  │     ├──  swagger-ui.css.map*
.rwxrwxrwx  257k root  1 May 13:42  │     ├──  swagger-ui.js*
.rwxrwxrwx  305k root  1 May 13:42  │     └──  swagger-ui.js.map*
.rwxrwxrwx   308 root  1 May 23:05  ├──  example.env*
.rwxrwxrwx  2.6k root  1 May 19:32  ├──  go.mod*
.rwxrwxrwx  194k root  1 May 19:32  ├──  go.sum*
drwxrwxrwx     - root  1 May 12:35  ├──  internal/
drwxrwxrwx     - root  1 May 13:55  │  ├──  db/
.rwxrwxrwx  1.8k root  1 May 13:55  │  │  ├──  db.dbml*
drwxrwxrwx     - root  1 May 13:50  │  │  ├──  migration/
.rwxrwxrwx    38 root  1 May 12:35  │  │  │  ├──  000001_init_database.down.sql*
.rwxrwxrwx   192 root  1 May 12:35  │  │  │  ├──  000001_init_database.up.sql*
.rwxrwxrwx    33 root  1 May 12:35  │  │  │  ├──  000002_users_table.down.sql*
.rwxrwxrwx  2.6k root  1 May 12:35  │  │  │  └──  000002_users_table.up.sql*
drwxrwxrwx     - root  1 May 13:51  │  │  ├──  queries/
.rwxrwxrwx  1.7k root  1 May 12:35  │  │  │  └──  users.sql*
.rwxrwxrwx  2.6k root  1 May 17:14  │  │  ├──  schema.sql*
drwxrwxrwx     - root  1 May 13:51  │  │  └──  sqlc/
.rwxrwxrwx   597 root  1 May 15:14  │  │     ├──  db.go*
.rwxrwxrwx  1.6k root  1 May 15:14  │  │     ├──  models.go*
.rwxrwxrwx  1.3k root  1 May 15:14  │  │     ├──  querier.go*
.rwxrwxrwx   516 root  1 May 13:50  │  │     ├──  querier_extended.go*
.rwxrwxrwx   641 root  1 May 14:50  │  │     ├──  store.go*
.rwxrwxrwx   12k root  1 May 15:14  │  │     ├──  users.sql_gen.go*
.rwxrwxrwx  2.5k root  1 May 13:49  │  │     └──  users_extended.go*
drwxrwxrwx     - root  1 May 13:08  │  ├──  proto/
drwxrwxrwx     - root  1 May 12:29  │  │  ├──  google/
drwxrwxrwx     - root  1 May 12:29  │  │  │  └──  api/
.rwxrwxrwx  1.0k root  1 May 12:29  │  │  │     ├──  annotations.proto*
.rwxrwxrwx  3.6k root  1 May 12:29  │  │  │     ├──  field_behavior.proto*
.rwxrwxrwx   15k root  1 May 12:29  │  │  │     ├──  http.proto*
.rwxrwxrwx  2.7k root  1 May 12:29  │  │  │     └──  httpbody.proto*
drwxrwxrwx     - root  1 May 12:29  │  │  ├──  protoc-gen-openapiv2/
drwxrwxrwx     - root  1 May 12:29  │  │  │  └──  options/
.rwxrwxrwx  1.8k root  1 May 12:29  │  │  │     ├──  annotations.proto*
.rwxrwxrwx   29k root  1 May 12:29  │  │  │     └──  openapiv2.proto*
.rwxrwxrwx   900 root  1 May 13:17  │  │  ├──  rpc-health.proto*
.rwxrwxrwx   415 root  1 May 13:18  │  │  ├──  rpc-welcome.proto*
.rwxrwxrwx  2.4k root  1 May 13:16  │  │  └──  service.proto*
drwxrwxrwx     - root  1 May 13:59  │  └──  service/
.rwxrwxrwx   538 root  1 May 15:08  │     ├──  rpc-health.go*
.rwxrwxrwx   499 root  1 May 15:08  │     ├──  rpc-welcome.go*
.rwxrwxrwx   842 root  1 May 15:07  │     └──  service.go*
.rwxrwxrwx  2.1k root  1 May 22:39  ├──  Makefile*
drwxrwxrwx     - root  1 May 14:41  ├──  pkg/
drwxrwxrwx     - root  1 May 12:18  │  ├──  crypt/
.rwxrwxrwx  1.5k root  1 May 13:19  │  │  └──  hash.go*
drwxrwxrwx     - root  1 May 16:18  │  ├──  db/
.rwxrwxrwx  1.8k root  1 May 22:55  │  │  ├──  connection.go*
.rwxrwxrwx  1.7k root  1 May 16:18  │  │  └──  migrate.go*
drwxrwxrwx     - root  1 May 14:41  │  ├──  email/
.rwxrwxrwx   683 root  1 May 22:53  │  │  ├──  email.go*
.rwxrwxrwx  2.1k root  1 May 22:53  │  │  └──  gmail.go*
drwxrwxrwx     - root  1 May 12:18  │  ├──  tokens/
.rwxrwxrwx   340 root  1 May 12:18  │  │  ├──  builder.go*
.rwxrwxrwx   277 root  1 May 12:18  │  │  ├──  main_test.go*
.rwxrwxrwx  1.2k root  1 May 12:18  │  │  ├──  paseto.go*
.rwxrwxrwx  1.8k root  1 May 13:01  │  │  ├──  paseto_test.go*
.rwxrwxrwx  1.2k root  1 May 12:18  │  │  └──  payload.go*
drwxrwxrwx     - root  1 May 12:18  │  ├──  utils/
.rwxrwxrwx   319 root  1 May 12:53  │  │  ├──  numbers.go*
.rwxrwxrwx  1.9k root  1 May 19:02  │  │  └──  randoms.go*
drwxrwxrwx     - root  1 May 12:17  │  └──  validator/
.rwxrwxrwx  1.9k root  1 May 23:04  │     └──  validator.go*
drwxrwxrwx     - root  1 May 20:41  ├──  public/
.rwxrwxrwx   53k root  1 May 19:52  │  ├──  console-log.png*
.rwxrwxrwx   21k root  1 May 19:52  │  ├──  evans-grpc.png*
.rwxrwxrwx   30k root  1 May 19:52  │  ├──  make-run.png*
.rwxrwxrwx  196k root  1 May 19:52  │  ├──  multiple-logs.png*
.rwxrwxrwx   48k root  1 May 19:52  │  ├──  postman.png*
.rwxrwxrwx   33k root  1 May 19:52  │  ├──  rpc-health.png*
.rwxrwxrwx   24k root  1 May 19:52  │  ├──  rpc-welcome.png*
.rwxrwxrwx  236k root  1 May 19:52  │  ├──  swagger-playground.png*
.rwxrwxrwx  229k root  1 May 19:52  │  └──  swagger.png*
.rwxrwxrwx   34k root  1 May 23:24  ├──  README.md*
.rwxrwxrwx   383 root  1 May 13:35  └──  sqlc.yaml*
```

## Directory Structure

### 1. `cmd/`

- ```sh
  Size Name
     -  ./
     - ├──  gateway/
  1.8k │  ├──  gateway.go*
   729 │  ├──  headers.go*
  1.0k │  └──  mw-logger.go*
     - ├──  grpc/
  1.2k │  ├──  grpc.go*
  1.7k │  └──  mw-logger.go*
  2.0k └──  main.go*
  ```

This directory contains the main executable files for the microservice. It includes the `gateway`and `grpc`subdirectories.

- The `gateway` directory houses the implementation for the gateway, which handles HTTP requests and forwards them to the appropriate services.

  - `gateway/headers.go` file contains utility functions and middleware to handle headers in the microservice. It provides functions to get, and set allow custom headers for incoming and outgoing requests.

  - `gateway/mw-logger.go` file contains the implementation of a logger middleware specifically designed for the `gateway`. It intercepts HTTP requests and logs relevant information such as the request method, URL, and response status.

  - Making new middlewares for gateway server is very easy. You have to follow a pattern prefix filename with `mw-` and then `name`. For an example `gateway/mw-name.go`

- The `grpc` directory contains the implementation for the gRPC server, which handles communication between microservices using the gRPC protocol.

  - `grpc/mw-logger.go` file contains the implementation of a logger middleware for the `gRPC server`. Similar to the gateway logger, it intercepts incoming gRPC requests and logs relevant information about the request and response.

  - Just like gateway middlewares you can create custom middlewares for grpc server. For example `grpc/mw-name.go`

### 2. `config/`

This directory contains configuration files for the microservice. Configuration is an essential part of any application, and here you can define various settings and parameters for your microservice, such as database connections.

- The `config.go` file contains the `Config` struct and the `LoadConfigs` function.

- The `Config` struct represents the application configuration and includes fields for `StartTime`, `ServiceName`, `GrpcPort`, `GatewayPort`, `Email`, and `Database`.

- The `LoadConfigs` function is responsible for loading the application configuration from a YAML file. It takes the file path and name as parameters and returns the loaded configuration.

### 3. `docs/`

The `docs` folder contains documentation-related files for the microservice. Here is an explanation of the files within the `docs` folder:

- `statik` folder: This folder contains a generated file named `statik.go`. The `statik.go` file is generated using the Statik library and serves as an embedded file system for serving static assets in the microservice.

- `swagger` folder: This folder contains the Swagger documentation files for the microservice's API. The Swagger documentation provides a detailed description of the API endpoints, request/response schemas, and other relevant information. The files within this folder include HTML, CSS, JavaScript, and image files used to render the Swagger UI.

### 4. `internal/`

The `internal/db` folder in the repository contains files related to the database operations of the microservice. Here is an explanation of the files within the `internal/db` folder:

- `db.dbml` file: This file represents the database schema using the DBML (Database Markup Language) syntax. It defines the tables, columns, relationships, and other database entities in a human-readable format.

- #### `internal/db/db.dbml`

  - The `db.dbml` file is a database schema file that represents the structure of the database using the **DBML (Database Markup Language)** syntax. DBML is a human-readable language for defining database schemas in a simplified and intuitive way.
  - In `db.dbml` file you can find the **definition of tables, columns, relationships**, and other entities that make up the database schema for the microservice. This file provides a high-level overview of the database structure and serves as a reference for developers and database administrators.

  - **You can visualize and interact with the `db.dbml` file using tools like [dbdiagram.io](https://dbdiagram.io/home)**, which is a popular online platform for designing and documenting database schemas. dbdiagram.io allows you to import the `db.dbml` file and generate interactive diagrams, shareable documentation, and even perform reverse engineering to generate SQL scripts or other database-related artifacts.

  - With the help of `db.dbml` we can generate `schema.sql` for `postgres`, `mysql` and `mssql` database

- #### `internal/db/schema.sql`

  - The `schema.sql` file represents the generated SQL script that defines the database schema. It is created using the **dbml2sql** command-line interface (CLI) tool, which takes the db.dbml file as input and generates the corresponding SQL statements.
  - By generating `schema.sql` it becomes easier to write migration files.

- #### `internal/db/migrations/`

- ```sh
  Size Name
     -  ./
    38 ├──  000001_init_database.down.sql*
   192 ├──  000001_init_database.up.sql*
    33 ├──  000002_users_table.down.sql*
  2.6k └──  000002_users_table.up.sql*
  ```

  - The migration files are written in SQL and consist of two parts: the "up" migration and the "down" migration. The "up" migration contains the SQL statements to apply the changes and update the schema to a new version, while the "down" migration contains the SQL statements to revert the changes and rollback the schema to a previous version.

- Writing migrations are much more easier when you already have `schema.sql` generated.

- #### `internal/db/queries/`
- ```sh
  Size Name
     -  ./
  1.7k └──  users.sql*
  ```

  - The `queries` directory contains raw SQL query files that allow developers to write custom and optimized queries tailored to their application's specific needs.

  - To simplify database interaction and enhance productivity, the [**`sqlc`**](https://sqlc.dev/) tool is utilized. It automatically generates Go code based on the SQL queries and database schema, producing type-safe store interfaces. These interfaces provide convenient methods for executing operations on database tables and executing the custom queries defined in the `queries` directory.

  - The generated store interfaces ensure compile-time safety and enable developers to write robust code for interacting with the database. The `sqlc` tool supports various databases and generates database-specific code based on the provided queries and schema.

- #### `internal/db/sqlc/` [sqlc](https://github.com/kyleconroy/sqlc)

  - ##### `sqlc.yaml` in root directory
  - ```yaml
      version: '2'

      sql:
      - schema: ./internal/db/migration
          queries: ./internal/db/queries
          engine: postgresql
          gen:
          go:
              package: sqlc
              out: ./internal/db/sqlc
              emit_json_tags: true
              emit_interface: true
              emit_empty_slices: true
              json_tags_case_style: snake
              emit_prepared_queries: false
              output_files_suffix: _gen
    ```

  - ##### `/internal/db/sqlc`
  - ```sh
       Size Name
           -  ./
       597 ├──  db.go*
      1.6k ├──  models.go*
      1.3k ├──  querier.go*
       516 ├──  querier_extended.go*
       641 ├──  store.go*
       12k ├──  users.sql_gen.go*
      2.5k └──  users_extended.go*
    ```

  - The `sqlc` directory contains the generated code produced by the [SQLC](https://sqlc.dev/) . SQLC is a code generator that automatically creates Go code based on the SQL queries and the database schema. It utilizes the references from the `migrations` and `queries` directories to generate the required code for interacting with the database.

  - The generated code includes database models, query executors, and store interfaces, providing type-safe access to the database schema and the custom queries defined in the `queries` directory. This generated code eliminates the need for writing repetitive and error-prone code for database operations.

  - By leveraging the `sqlc` directory, developers can perform database operations efficiently and maintain consistency in their codebase. The generated code serves as a reliable bridge between the database schema, custom queries, and the application code, enabling the development of robust and maintainable database-driven applications.

- #### `internal/proto/` [a bit of everything ](https://github.com/grpc-ecosystem/grpc-gateway/blob/main/examples/internal/proto/examplepb/a_bit_of_everything.proto)

  - ```sh
      Size Name
          -  ./
          - ├──  google/
          - │  └──  api/
      1.0k │     ├──  annotations.proto*
      3.6k │     ├──  field_behavior.proto*
      15k │     ├──  http.proto*
      2.7k │     └──  httpbody.proto*
          - ├──  protoc-gen-openapiv2/
          - │  └──  options/
      1.8k │     ├──  annotations.proto*
      29k │     └──  openapiv2.proto*
      900 ├──  rpc-health.proto*
      415 ├──  rpc-welcome.proto*
      2.4k └──  service.proto*
    ```

  - The `google` directory contains protobuf files used for generating the gateway server.
  - The `protoc-gen-openapiv2` directory contains protobuf files with options for generating OpenAPI v2 documentation.
  - To create a new API route and an RPC function, you can create a new file in the proto directory with the naming convention rpc-openration-name.proto. This file will define the new RPC operation and its associated request and response message types.

  - Here is how a `rpc-welcome.proto` api looks like
  - ```ts
        syntax = "proto3";

        package gomicro;

        // Specify the Go package name for the generated code.
        option go_package = "github.com/sirjager/rpcs/gomicro";

        // WelcomeRequest represents a request for the welcome message.
        message WelcomeRequest {}

        // WelcomeResponse represents the response containing the welcome message.
        message WelcomeResponse {
            // Message represents the welcome message.
            string message = 1;
        }
    ```

  - `service.proto`

  - ```ts
    syntax = "proto3";

    package gomicro;

    import "google/api/annotations.proto";
    import "protoc-gen-openapiv2/options/annotations.proto";

    import "rpc-welcome.proto";
    import "rpc-health.proto";

    // Specify the Go package name for the generated gRPC stubs.
    // The package name determines the import path used in other Go files to reference the generated code.
    // In this case, the github.com/sirjager/rpcs/gomicro package is used to organize and import the generated gRPC stubs
    // for the Go Microservice Template.
    option go_package = "github.com/sirjager/rpcs/gomicro";


    // Define OpenAPI v2 options for generating Swagger documentation.
    option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_swagger) = {
        info: {
            title: "Go Microservice Template";
            version: "0.1.0";
            contact: {
                name: "SirJager";
                url: "https://github.com/sirjager/gomicro";
            };
        };
    };

    // Micro service
    service GoMicro {
        // Define OpenAPI v2 options for generating tags and external documentation.
        option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_tag) = {
            description: "GoMicro is a comprehensive starter template designed to accelerate the development of microservices using the Go programming language. It provides a solid foundation and essential components to expedite the creation of scalable and modular microservice architectures";
            external_docs: {
                url: "https://github.com/sirjager/gomicro";
                description: "Find out more about GoMicro";
            };
        };
        // Welcome endpoint returns a welcome message.
        rpc Welcome(WelcomeRequest) returns (WelcomeResponse) {
            option (google.api.http) = {
                get: "/"
            };
            option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_operation) = {
                description: "Returns a welcome message";
                summary: "Welcome Message";
                tags: "System";
            };
        }
        // Health endpoint returns the health status of the API.
        rpc Health(HealthRequest) returns (HealthResponse) {
            option (google.api.http) = {
                get: "/v1/health"
            };
            option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_operation) = {
                description: "Returns the health status of the API";
                summary: "API Health";
                tags: "System";
            };
        }
    }
    ```

- #### `internal/service/`

  - ```sh
        -  ./
    538 ├──  rpc-health.go*
    499 ├──  rpc-welcome.go*
    842 └──  service.go*
    ```

  - The `service.go` file in the `internal/service` directory contains the implementation of the core service of your application. This file defines a struct named `CoreService`, which represents the main service instance. The `CoreService` struct contains various fields that hold important components for your service.
    `service.go`
  - ```go
    package service

    import (
    "github.com/rs/zerolog"
    "github.com/sirjager/gomicro/config"
    "github.com/sirjager/gomicro/internal/db/sqlc"
    "github.com/sirjager/gomicro/pkg/email"

        rpc "github.com/sirjager/rpcs/gomicro/go"

    )

    // CoreService represents the core service of your application.
    type CoreService struct {
    rpc.UnimplementedGoMicroServer
    Logr zerolog.Logger // Logger is used for logging.
    Config config.Config // Config holds the configuration settings.
    mailer email.MailSender // Mailer is used for sending emails.
    store sqlc.Store
    }

    // NewService creates a new instance of the CoreService.
    func NewService(logr zerolog.Logger, config config.Config, mailer email.MailSender, store sqlc.Store) (\*CoreService, error) {

        return &CoreService{
            Logr:   logr,
            Config: config,
            mailer: mailer,
            store:  store,
        }, nil

    }
    ```

Here's a breakdown of the important parts of the `service.go` file:

- `CoreService` struct: This struct represents the core service of your application. It implements the `rpc.UnimplementedGoMicroServer` interface, which is generated from the protobuf service definition. The struct has the following fields:

- `Logr`: An instance of the `zerolog.Logger` type, which is used for logging purposes.
- `Config`: An instance of the `config.Config` type, which holds the configuration settings for your service.
- `mailer`: An instance of the `email.MailSender` type, which is used for sending emails.
- `store`: An instance of the `sqlc.Store` type, which represents the database store used by your service.

- `NewService` function: This function is a constructor for creating a new instance of the `CoreService`. It takes in various parameters, such as the logger, configuration, mailer, and store, and returns a pointer to a `CoreService` instance. Inside the function, the fields of the `CoreService` struct are initialized with the provided values.

The `service.go` file is where you can define and implement the core functionalities of your service. You can customize this file to add additional methods and logic to handle the RPC operations defined in the protobuf service.

- `rpc-welcome.go` Here is how we write business logic

- ```go
  package service

  import (
      "context"
      "fmt"

      rpc "github.com/sirjager/rpcs/gomicro/go"
  )

  // welcomeMessage returns a formatted welcome message string.
  func welcomeMessage(name string) string {
      return fmt.Sprintf("Welcome to %s Api", name)
  }

  // Welcome is a service method that returns a welcome message.
  func (s *CoreService) Welcome(ctx context.Context, req *rpc.WelcomeRequest) (*rpc.WelcomeResponse, error) {
      return &rpc.WelcomeResponse{Message: welcomeMessage(s.Config.ServiceName)}, nil
  }

  ```

### 5. `pkg/`

- The `pkg` directory contains various packages that provide reusable functionalities for your application. Here's a brief overview of the packages found in the `pkg` directory:

```sh
Size Name
   -  ./
   - ├──  crypt/
1.5k │  └──  hash.go*
   - ├──  db/
1.8k │  ├──  connection.go*
1.7k │  └──  migrate.go*
   - ├──  email/
 582 │  ├──  email.go*
2.3k │  └──  gmail.go*
   - ├──  tokens/
 340 │  ├──  builder.go*
 277 │  ├──  main_test.go*
1.2k │  ├──  paseto.go*
1.8k │  ├──  paseto_test.go*
1.2k │  └──  payload.go*
   - ├──  utils/
 319 │  ├──  numbers.go*
2.1k │  └──  randoms.go*
   - └──  validator/
1.8k    └──  validator.go*
```

- `crypt`: This package contains the `hash.go` file, which provides functions for hashing and verifying hashes.

- `db`: This package contains the `connection.go` and `migrate.go` files. `connection.go` is responsible for establishing connections to the database, while `migrate.go` handles database migrations, ensuring that the database schema is up to date.

- `email`: This package includes the `email.go` and `gmail.go` files. It provides functionalities for sending emails. This pkg can be easily extended to use other services like `Mailgun`, `SendGrid` etc.

- `tokens`: The `tokens` package contains files such as `builder.go`, `paseto.go`, `paseto_test.go`, and `payload.go`. This package provides utilities for working with authentication and authorization tokens, such as building and parsing tokens using the Paseto token format.

  > _I choose to use **`paseto`(Platform-Agnostic Security Tokens) instead of `jwt` (JSON Web Token) tokens**_, [Here is why.](https://paragonie.com/blog/2017/03/jwt-json-web-tokens-is-bad-standard-that-everyone-should-avoid)
  > You can always implement jwt tokens or any other tokens by satisfying builder interface

- `utils`: The `utils` package contains files like `numbers.go` and `randoms.go`. This package provides various utility functions related to numbers and randomization like `random email`. `random username`, `random password` etc.

- `validator`: This package includes the `validator.go` file, which provides validation functionalities. It contains functions for validating data, such as validating `email`, `username`, `uuid` etc.

- These packages in the `pkg` directory are designed to promote code reusability and organization. You can leverage these packages to simplify your application development by utilizing well-defined and tested functionalities in different parts of your codebase.

### 6. `Makefile`

The `Makefile` is a build automation tool used to define and organize tasks in a project. It provides a convenient way to execute commands for building, testing, and managing various aspects of the project. Let's go through the sections of the provided `Makefile`:

- `proto`: This target generates protobuf and gRPC code, Swagger documentation, and static assets. It uses the `protoc` command to compile the `.proto` files in the specified directory (`PROTO_DIR`) and generates Go code for gRPC, gRPC-gateway, and OpenAPI Swagger. The resulting files are placed in the `$(GO_RPC_DIR)` directory.

- `tidy`: This target cleans up the `go.sum` file and tidies the project dependencies using `go mod tidy`. It also fetches the latest version of the `github.com/sirjager/rpcs` package.

- `test`: This target runs tests for the project using the `go test` command. It includes the `-v` flag for verbose output, `-cover` flag for code coverage analysis, and `-short` flag to skip long-running tests.

- `build`: This target builds the project by running the `go build` command. It creates an executable named `main` in the `./dist` directory by compiling the `cmd/main.go` file.

- `run`: This target runs the project using the `go run` command. It executes the `cmd/main.go` file.

- `dbdocs`: This target generates database documentation using the `dbdocs` tool. It builds the documentation based on the `db.dbml` file located in `./internal/db` directory.

- `dbschema`: This target generates the database schema by running the `dbml2sql` command. It converts the `db.dbml` file into SQL code and outputs it to the `./internal/db/schema.sql` file.

- `migratenew`: This target creates a new database migration file using the `migrate create` command. It creates a migration file with the given name in the `$(DB_MIGRATIONS)` directory.

- `migrateup`: This target runs the database migrations (up) using the `migrate` command. It applies any pending migrations to the database specified by the `$(DB_URL)`.

- `migratedown`: This target rolls back the database migrations (down) using the `migrate` command. It reverses all applied migrations in the database specified by the `$(DB_URL)`.

- `sqlc`: This target generates SQL code using SQLC. It runs the `sqlc generate` command to generate code based on the SQL queries defined in the project.

These targets can be executed by running `make <target>` in the command line, where `<target>` is the desired task from the Makefile.

```markdown
SERVICE_NAME=gomicro
RPCS_DIR=../rpcs
PROTO_DIR=./internal/proto

GO_RPC_DIR=$(RPCS_DIR)/$(SERVICE_NAME)/go

# open api swagger documentations

STATIK_OUT=./docs
SWAGGER_OUT=./docs/swagger

# TEST: database configs

DB_MIGRATIONS=./internal/db/migration
DB_URL=postgres://postgres:password@localhost:5432/databasename?sslmode=disable
```

### 7. `example.env`

```go
  config, err := config.LoadConfigs(".", "example")
  if err != nil {
    logr.Fatal().Err(err).Msg("failed to load configurations")
  }
```

The provided configuration consists of environment variables that can be used to configure a gRPC server, a gateway server, a database connection, and mail SMTP settings. Here's an explanation of each variable:

`GRPC_PORT`: The port number on which the gRPC server will listen.
`GATEWAY_PORT`: The port number on which the gateway server (e.g., REST server) will listen.
`DB_HOST`: The host address of the database server.
`DB_PORT`: The port number of the database server.
`DB_NAME`: The name of the database to connect to.
`DB_USER`: The username to use for authenticating the database connection.
`DB_PASS`: The password to use for authenticating the database connection.
`DB_DRIVER`: The driver or database type to use for the connection.
`DB_ARGS`: Additional arguments or options for the database connection string, such as SSL mode.
`DB_MIGRATE`: The path to the database migration files.

For the mail SMTP configuration:

`GMAIL_NAME`: The name or display name associated with the Gmail account.
`GMAIL_USER`: The email address of the Gmail account.
`GMAIL_PASS`: The application-specific password generated for the Gmail account to be used for authentication. You need to get app password for gmail stmp from https://myaccount.google.com/apppasswords

These environment variables can be used to provide the necessary configuration values for the corresponding components of your application, such as the gRPC server, gateway server, database connection, and mail sending functionality.

### 8. `.goreleaser.yaml` config for [GoReleaser](https://goreleaser.com/)

- command `make release` or `goreleaser --snapshot --clean`

The build configuration for a project named "gomicro" using a build tool [GoReleaser](https://github.com/goreleaser/goreleaser). It specifies different builds for multiple operating systems and architectures.

Here's an explanation of the configuration:

- `project_name: gomicro`: Specifies the name of the project as "gomicro". **This is the only thing you will need to change for minimum release configs**.

Builds:

- `id`: An identifier for the build.
- `binary`: Specifies the name of the binary output file. The `{{.ProjectName}}`, `{{.Os}}`, and `{{.Arch}}` are placeholders that will be replaced with actual values during the build process.
- `main`: Specifies the directory path where the main package or entry point of the application is located.
- `goos`: Specifies the target operating system(s) for the build.
- `goarch`: Specifies the target architecture(s) for the build.
- `no_unique_dist_dir`: Indicates whether to use a unique distribution directory for each build. In this case, it is set to `true`, meaning each build will have its own directory.

Archives:

- `format`: Specifies the format for archiving the build. In this case, it is set to "binary", indicating that the build output will be archived as a binary file.
- `name_template`: Specifies the template for naming the archived file. The `{{.ProjectName}}`, `{{.Os}}`, and `{{.Arch}}` are placeholders that will be replaced with actual values during the build process.

Checksum:

- `name_template`: Specifies the template for naming the checksum file. The `{{.ProjectName}}` is a placeholder that will be replaced with the actual project name during the build process.

This configuration allows you to build the "gomicro" project for different operating systems and architectures, generating binaries and archives with specific naming conventions. It also includes checksum generation for the built files.

### 9. `Dockerfile`

- This Dockerfile defines a multi-stage build process to create a container image for a Go application.
- This needs no configuration to work unless you want to modif for your own custom needs.

### 10. `docker-compose.yml`

This is a Docker Compose file that defines two services: a `PostgreSQL database` service and a `gomicro` service.

1. **Database Service:**

   - `container_name: db`: Sets the container name for the database service as "db".
   - `image: postgres:14-alpine`: Specifies the image to use for the database service, which is `postgres:14-alpine` (PostgreSQL 14 with Alpine Linux).
   - `volumes: - db:/var/lib/postgresql/data`: Creates a named volume called "db" and mounts it to the `/var/lib/postgresql/data` directory inside the database container. This volume is used to persist the database data.
   - `environment`: Sets environment variables for the database service, including the username, password, and database name.

2. **gomicro Service:**
   - `image: ghcr.io/sirjager/gomicro:latest`: Specifies the image to use for the `gomicro` service, which is `ghcr.io/sirjager/gomicro:latest`.
   - `container_name: gomicro`: Sets the container name for the `gomicro` service as "gomicro".
   - `depends_on: - db`: Specifies that the `gomicro` service depends on the `db` service. This ensures that the database service starts before the `gomicro` service.
   - `ports: - 4420 - 4424`: Maps the container's ports 4420 and 4424 to the same ports on the host machine, allowing external access to the `gomicro` service.
   - `environment`: Sets environment variables for the `gomicro` service, including the database connection details, email-related variables, and other configuration variables.
   - `entrypoint: ['/app/wait-for.sh', 'db:5432', '--', '/app/start.sh']`: Sets the entrypoint command for the `gomicro` container. It first waits for the database service to be available before executing the `/app/start.sh` script.
   - `command: ['/app/main']`: Specifies the command to run when the `gomicro` container starts, which is `/app/main`.

This Docker Compose file allows you to run both the database service and the `gomicro` service together, with the `gomicro` service depending on the database service. The services can communicate with each other, and the `gomicro` service is accessible externally through ports 4420 and 4424.

- This Dockerfile defines a multi-stage build process to create a container image for a Go application.
- This needs no configuration to work unless you want to modif for your own custom needs.

### 11. `.github/workflows/release.yml`

This is a GitHub Actions workflow that automates the release and deployment process for a Go application. Let's break down the different sections and steps of the workflow:

1. **Workflow Configuration**: This section defines the name of the workflow and specifies that it should only run when a tag is pushed to the repository. The tags are filtered to match the pattern "v\*", which means only tags starting with "v" will trigger the workflow.

2. **Environment Variables**: This section defines two environment variables - `REGISTRY` and `REPOSITORY`. `REGISTRY` is set to `ghcr.io`, which is the GitHub Container Registry. `REPOSITORY` is set to the name of the GitHub repository where the workflow is running.

3. **Jobs**: The workflow has a single job called "release" that runs on the latest version of Ubuntu.

4. **Permissions**: The permissions section specifies that the job needs write access to the repository contents and packages.

5. **Steps**:
   a. **Checkout Repository**: This step checks out the repository code using the `actions/checkout` action.
   b. **Setup Golang**: This step sets up the Go environment using the `actions/setup-go` action. It specifies the minimum required Go version and enables caching to speed up subsequent builds.
   c. **Create Release**: This step uses the `goreleaser/goreleaser-action` action to create a Go release. It invokes the `goreleaser` command with the `release --clean` arguments, which generates a release artifact and publishes it in the repository's release section on GitHub. The `GITHUB_TOKEN` environment variable is passed to authenticate the action.
   d. **Log in to the Container registry**: This step uses the `docker/login-action` action to authenticate with the GitHub Container Registry using the provided `GITPAT` secret. It uses the `ghcr.io` registry and sets the username to the actor (user) who triggered the workflow.
   e. **Extract metadata (tags, labels) for Docker**: This step uses the `docker/metadata-action` action to extract metadata (such as tags and labels) for the Docker image. It specifies the image to extract metadata from, which is in the format `ghcr.io/repository`.
   f. **Build and push Docker image**: This step uses the `docker/build-push-action` action to build and push a Docker image. It specifies the build context (the current directory), enables pushing the image, and sets the tags and labels for the image based on the metadata extracted in the previous step.

Overall, this workflow listens for tagged commits, creates a Go release, and deploys the built image to the GitHub Container Registry.

### 12. `public/` for storing assets.

#### run server : `make run`

![command: make run, showing some initial logs screenshot](/public/make-run.png)

#### OpenAPI swagger documentation

![swagger documentation screenshot](/public/swagger.png)

#### Testing GRPC server using [`Evans`](https://github.com/ktr0731/evans)

![testing  grpc server using evans client screenshot](/public/evans-grpc.png)

#### GRPC Welcome Call

![grpc server welcome call screenshot](/public/rpc-welcome.png)

#### GRPC Health Call

![grpc server health call screenshot](/public/rpc-health.png)

#### Console Logs

![console logs of incoming requests screenshot](/public/console-log.png)

#### Testing `/v1/health` in OpenAPI swagger playground

![testing /v1/health route in swagger playground  screenshot](/public/swagger-playground.png)

#### Testing `/v1/health` in Postman

![testing /v1/health route in postman screenshot](/public/postman.png)

#### Console logs

![testing /v1/health route in postman screenshot](/public/multiple-logs.png)

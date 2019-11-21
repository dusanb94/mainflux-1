# Auth service

Auth service provides an API for managing API keys. Through this API clients
are able to do the following actions on API keys:

- create
- obtain
- verify
- revoke

There are *three types of API keys*:

- session key - keys issued to the user upon Login request
- user key - keys issued upon the user request
- reset key - password reset key

Session keys are issued when user logs in. Each user request (other than `registration` and `login`) contains session key that is used to authenticate the user. User keys are similar to the session keys. The main difference is that user keys have configurable expiration time. If no time is set, the key will never expire. For that reason, user keys are _the only key type that can be revoked_. Reset key is the password reset key. It's short-lived token used for password recovery process.

For in-depth explanation of the aforementioned scenarios, as well as thorough
understanding of Mainflux, please check out the [official documentation][doc].

## Configuration

The service is configured using the environment variables presented in the
following table. Note that any unset variables will be replaced with their
default values.

| Variable                 | Description                                                             | Default        |
|--------------------------|-------------------------------------------------------------------------|----------------|
| MF_AUTH_LOG_LEVEL        | Service level (debug, info, warn, error)                                | error          |
| MF_AUTH_DB_HOST          | Database host address                                                   | localhost      |
| MF_AUTH_DB_PORT          | Database host port                                                      | 5432           |
| MF_AUTH_DB_USER          | Database user                                                           | mainflux       |
| MF_AUTH_DB_PASSWORD      | Database password                                                       | mainflux       |
| MF_AUTH_DB               | Name of the database used by the service                                | auth           |
| MF_AUTH_DB_SSL_MODE      | Database connection SSL mode (disable, require, verify-ca, verify-full) | disable        |
| MF_AUTH_DB_SSL_CERT      | Path to the PEM encoded certificate file                                |                |
| MF_AUTH_DB_SSL_KEY       | Path to the PEM encoded key file                                        |                |
| MF_AUTH_DB_SSL_ROOT_CERT | Path to the PEM encoded root certificate file                           |                |
| MF_AUTH_HTTP_PORT        | Auth service HTTP port                                                  | 8180           |
| MF_AUTH_GRPC_PORT        | Auth service gRPC port                                                  | 8181           |
| MF_AUTH_SERVER_CERT      | Path to server certificate in pem format                                |                |
| MF_AUTH_SERVER_KEY       | Path to server key in pem format                                        |                |
| MF_AUTH_SECRET           | String used for signing tokens                                          | auth          |
| MF_JAEGER_URL            | Jaeger server URL                                                       | localhost:6831 |

## Deployment

The service itself is distributed as Docker container. The following snippet
provides a compose file template that can be used to deploy the service container
locally:

```yaml
version: "2"
services:
  auth:
    image: mainflux/auth:[version]
    container_name: [instance name]
    ports:
      - [host machine port]:[configured HTTP port]
    environment:
      MF_AUTH_LOG_LEVEL: [Service log level]
      MF_AUTH_DB_HOST: [Database host address]
      MF_AUTH_DB_PORT: [Database host port]
      MF_AUTH_DB_USER: [Database user]
      MF_AUTH_DB_PASS: [Database password]
      MF_AUTH_DB: [Name of the database used by the service]
      MF_AUTH_DB_SSL_MODE: [SSL mode to connect to the database with]
      MF_AUTH_DB_SSL_CERT: [Path to the PEM encoded certificate file]
      MF_AUTH_DB_SSL_KEY: [Path to the PEM encoded key file]
      MF_AUTH_DB_SSL_ROOT_CERT: [Path to the PEM encoded root certificate file]
      MF_AUTH_HTTP_PORT: [Service HTTP port]
      MF_AUTH_GRPC_PORT: [Service gRPC port]
      MF_AUTH_SECRET: [String used for signing tokens]
      MF_AUTH_SERVER_CERT: [String path to server certificate in pem format]
      MF_AUTH_SERVER_KEY: [String path to server key in pem format]
      MF_JAEGER_URL: [Jaeger server URL]
```

To start the service outside of the container, execute the following shell script:

```bash
# download the latest version of the service
go get github.com/mainflux/mainflux

cd $GOPATH/src/github.com/mainflux/mainflux

# compile the service
make auth

# copy binary to bin
make install

# set the environment variables and run the service
MF_AUTH_LOG_LEVEL=[Service log level] MF_AUTH_DB_HOST=[Database host address] MF_AUTH_DB_PORT=[Database host port] MF_AUTH_DB_USER=[Database user] MF_AUTH_DB_PASS=[Database password] MF_AUTH_DB=[Name of the database used by the service] MF_AUTH_DB_SSL_MODE=[SSL mode to connect to the database with] MF_AUTH_DB_SSL_CERT=[Path to the PEM encoded certificate file] MF_AUTH_DB_SSL_KEY=[Path to the PEM encoded key file] MF_AUTH_DB_SSL_ROOT_CERT=[Path to the PEM encoded root certificate file] MF_AUTH_HTTP_PORT=[Service HTTP port] MF_AUTH_GRPC_PORT=[Service gRPC port] MF_AUTH_SECRET=[String used for signing tokens] MF_AUTH_SERVER_CERT=[Path to server certificate] MF_AUTH_SERVER_KEY=[Path to server key] MF_JAEGER_URL=[Jaeger server URL] $GOBIN/mainflux-auth
```

If `MF_EMAIL_TEMPLATE` doesn't point to any file service will function but password reset functionality will not work.

## Usage

For more information about service capabilities and its usage, please check out
the [API documentation](swagger.yaml).

[doc]: http://mainflux.readthedocs.io
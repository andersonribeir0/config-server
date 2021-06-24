# Config-server

Config server is a golang proof of concept where environment settings are constantly updated.
By executing this job the environment configuration should be always up to date without performing a new deployment.

# Prerequisites
Docker and docker compose.
# How does this POC keep everything up to date?
Config server uses Hashicorp Consul as service discovery and environment key-value pairs provider.

When the server starts it performs a polling job into consul getting services addresses and application service key-value pairs.
In case of consul hit failure, it should try to get local environment variable with the same key.

# How to run locally
1. Build an image by performing `docker build -t config-server .`


2. Inside docker-compose.yaml file you should be able to find commands into `command` key into config-server service.
There is some params that may be passed to run the http server:
```
-n, --app-name string             Application name. (default "config-server")
--auto-refresh                    Refresh key-value pairs from consult integration. (default true)
--auto-refresh-duration int       Interval between key-value pairs refreshes in seconds. (default 5)
--consul-port string              Port to be used on consul. (default ":8500")
--consul-prefix string            Consul key value prefix. (default "config")
--consul-url string               Consul URL. (default "consul")
-h, --help                        help for http
-p, --port string                 Port to be used on http server. (default ":4040")
```

3. After filling up with the desired configuration, you may start the local environment by executing `docker-compose up`

4. You should be able to see logging config variables by config-server or getting them by performing a cURL to `<local_config_server_address>:<local_config_server_port>`.
   If you leave the default values, the config-server address should be `localhost:4040` and the consul ui address `localhost:8500`.

5. Insert some key-value pairs into consul.
If you leave the default values, you should:
   - Access `localhost:8500`
   - Click on `Key/Value` nav bar option
   - Click on `Create` button
   - Insert `config/config-server/<any_key>` on `Key or Folder` input
   - Insert `<any_value>` into `Value` input
   - Click on `Save` button

6. Now you can see the config variable changing dynamically

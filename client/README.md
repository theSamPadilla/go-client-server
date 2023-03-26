# Client implementation
This is a standalone package to instantiate client that writes a message to a RabbitMQ queue.

Clients can ask the server to perform the following operations:
- Add item
- Remove item
- Get item
- Get item by index
- Get all items

This diretory contains bash scripts of example operations to use.

## Usage
To spin up a new client and send a new message:
```
go run client.go <command> [flags]
```

The following commands and required flags are:
- **Command: `add`**
    - Adds a new value to the ordered map. Resets the value if it exists.
    - Required flags: `-key` and `-value`
- **Command: `remove`**
    - Removes a key and its respective value from the ordered map.
    - Required flag: `-key`
- **Command: `get`**
    - Gets corresponding value in the orded map.
    - Required flag: `-key`
- **Command: `geti`**
    - Gets value corresponding to the index passed.
    - Required flag: `-index`
        - `index` value must be an unsigned int.
- **Command: `getall`**
    - Gets all the elements of the ordered map in order of insertion.
    - No required flags.

# Http server
This go script is not needed to spin up an http client. Simply make a `curl` requests to the appropriate endpoint.
Server listens to calls on `localhost:6969`.

### REST endpoints, methods, and formats:
- **`GET /`**
    - Gets all the elements of the ordered map in order of insertion.

- **`GET /key/<key>`**
    - Gets value corresponding to the key passed.

- **`GET /index/<index>`**
    - Gets value corresponding to the index passed.
    - Index must be an unsigned int.

- **`POST /add`**
    - Adds a key value pair to the ordered map.
    - Requires a request body of `Content-type: application/json` with the following values:
        - `key: <key>`
        - `value: <value>`

- **`POST /remove`**
    - Removes a key and its respecitve value from the ordered map.
    - Requires a request body of `Content-type: application/json` with the following values:
        - `key: <key>`
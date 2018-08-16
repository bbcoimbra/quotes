Quotes
======

Pet-project to learn some of Go Lang and its ecossystem.

Build
-----

Get source by running `go get` as follow:

```bash
go get github.com/bbcoimbra/quotes
```

and build as follow:

```bash
go build
```

Running
-------

To run server execute (inside build directory):

```bash
./quotes server
```

Server will running on all interfaces on port TCP/8080.

Client
------

The HTTP API is very simple, but there is a client to make interaction even easier.

### Add Quote

```bash
./quotes client add --text 'A vingança nunca é plena. Mata a alma e a envenena.' --author 'Seu Madruga'
```

### List Quotes

```bash
./quotes client list
```

### Get Quote

```bash
./quotes client get --quote-id 1
```

### Delete Quote


```bash
./quotes client get --quote-id 1
```

Configuration
-------------

Some configuration can be done into `$HOME/.quotes.yaml` file.

### Server

* listen: set pair IP/port that server will try to bind and answer requests

### Client

* api\_endpoint: Address that client will send requests

### Example

```yaml
server:
  listen: localhost:8080
client:
  api_endpoint: http://localhost:8080
```

TODO
----

* [x] add persistent configuration to interfaces to bind
* [x] add persistent configuration to ports to bind
* [x] add flags to configure interfaces to bind
* [x] add flags to configure ports to bind
* [ ] persist Quotes into a database (actually Quotes is stored only in memory)
* [ ] add Votes to Quotes (:+1: and :-1:)
* [ ] add authentication
* [ ] add authorization

License
=======

See LICENCE file.

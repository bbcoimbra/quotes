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

TODO
----

* [ ] add persistent configuration to interfaces to bind
* [ ] add persistent configuration to ports to bind
* [ ] add flags to configure interfaces to bind
* [ ] add flags to configure ports to bind
* [ ] add authentication
* [ ] add authorization
* [ ] persist Quotes into a database (actually Quotes is stored only in memory)
* [ ] add Votes to Quotes (:+1: and :-1:)

License
=======

See LICENCE file.
Example bot implementation for [checkers-bot](https://github.com/tleyden/checkers-bot) which makes random moves.

# Big Picture

![architecture png](http://cl.ly/image/3S0G0h2U0R2b/Screen%20Shot%202013-10-08%20at%2010.43.00%20PM.png)

# Install 

First you will need to [install Go 1.1](http://golang.org/doc/install) or later.

```
$ go get github.com/tleyden/checkers-bot-random
$ go get github.com/couchbaselabs/go.assert
```

It is also possible to download the latest build of the [OSX binary](http://cbfs.hq.couchbase.com:8484/projects/checkers-bot/checkers-bot-random.mac.gz) directly.

# Validate installation - run tests

```
$ cd $GOPATH/github.com/tleyden/checkers-bot-random
$ go test -v
```

# Install other system components

To get a fully working system, you'll need the following:

* [Couchbase Server](http://www.couchbase.com/download)

* [Sync Gateway](https://github.com/couchbase/sync_gateway)

* [Checkers Overlord](https://github.com/apage43/checkers-overlord)

* [Checkers-iOS](https://github.com/couchbaselabs/Checkers-iOS)

Checkers-iOS is not strictly required, but very useful in order to view the game.

The other way to install [Checkers-iOS](https://github.com/couchbaselabs/Checkers-iOS) from github.

It can be installed for the [iTunes Store](https://itunes.apple.com/us/app/id698034787), however that version is only able to connect to the non-public production server.



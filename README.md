# tachyon cli plugin

*there was once a great and powerful command for the cf cli, it went away, now it's back!*

## about

Ever want to create a space and then switch to it immediately? Yeah, I thought
so. This thing does that thing.

## installation

```
$ go get github.com/xoebus/tachyon
$ cf install-plugin $GOPATH/bin/tachyon
```

## usage

```
$ cf take-space awesome-new-space
```

### Thanks (embeds)

- https://gist.github.com/albertbori/1798d88a93175b9da00b

#### Function

    - When Ethernet connected, WiFi will be closed automatically.
    - When Ethernet disconnected, WiFi will be open automatically.

### Install

```shell
go install github.com/honwen/mac-auto-xlan-switch@latest
```

### Use

```shell
sudo mac-auto-xlan-switch start
sudo mac-auto-xlan-switch stop
```

### Help

```shell
$ mac-auto-xlan-switch -h
NNAME:
   switcher - mac-auto-xlan-switch

USAGE:
   mac-auto-xlan-switch [global options] command [command options] [arguments...]

VERSION:
   Git:[MISSING BUILD VERSION [GIT HASH]] (go1.20.4)

COMMANDS:
   start    Start Auto xLan Switcher
   stop     Stop Auto xLan Switcher
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

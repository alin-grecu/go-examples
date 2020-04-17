## Simple Network CLI tool


```bash
NAME:
   netcli - A new cli application

USAGE:
   netcli.exe [global options] command [command options] [arguments...]

COMMANDS:
   ns       Looks up the nameservers for a particular host
   ip       Looks up the IP addresses for a particular host
   cname    Looks up the CNAME for a particular host
   mx       Looks up the MX records for a particular host
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)
PS C:\Users\alins\Documents\GitHub\go-examples\cli-tools\network-cli> .\netcli.exe
NAME:
   netcli.exe - A new cli application

USAGE:
   netcli.exe [global options] command [command options] [arguments...]

COMMANDS:
   ns       Looks up the nameservers for a particular host
   ip       Looks up the IP addresses for a particular host
   cname    Looks up the CNAME for a particular host
   mx       Looks up the MX records for a particular host
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)
```

## Source
https://tutorialedge.net/golang/building-a-cli-in-go/

## Changes
The tutorial uses [ufave/cli/v1](https://github.com/urfave/cli/blob/master/docs/v1/manual.md) which is not on master anymore.
This example is adapted for v2.

## Run application
### Prerequisites
```bash
go get github.com/urface/cli
```

### Build and run
```bash
go build cmd/network-cli/netcli.go
./netcli
```
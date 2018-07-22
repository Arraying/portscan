# portscan

Port range scanner implementation in Go

## Compiling

1) Install Go & set up *GOPATH*/*GOROOT*
2) `git clone https://github.com/Arraying/portscan.git`
3) `cd portscan`
4) Optional - set *GOOS* and *GOARCH* environment variables if intending to cross compile
5) `go build`

## Usage

Run the binary with the desired flags.

### Required flags

* `host` The host address (IPv4) of the target.
* `min` The minimum port number (inclusive).
* `max` The maximum port number (inclusive).

### Optional flags

* `suppress` Suppresses output for invidivual ports, only results will show, disabled by default.
* `timeout` The TCP timeout (ms), 1000 by default. May require alteration depending on latency to target.

## Example

`./portscan -host="127.0.0.1" -min=22 -max=443 -suppress -timeout=500`

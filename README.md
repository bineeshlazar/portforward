# Simple TCP port forwarder

Inspired from [tcp-port-forward](https://github.com/fidian/tcp-port-forward). This tool opens a port on the local machine and forwards the connection to another machine.

## Usage

### Compiling

To download/compile the code run,

    go get github.com/bnsh12/portforward

This will build the executable in your $GOPATH/bin directory.

### Running

Run the command, (Assuming that $GOPATH/bin is added to your PATH variable)

    portforward server-address:port local-address:port

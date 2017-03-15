# cxmate

You can set the address and port that symbiont listens on with environment variables:

```
export LISTENING_ADDRESS = "0.0.0.0"
export LISTENING_PORT = "80"
```

You can set the address and port of the proxied biological service in the same way:

```
export SERVICE_ADDRESS = "127.0.0.1"
export SERVICE_PORT = "8080"
```

You can set the aspects cxmate will forward and accept from the service like this:

```
export RECEIVES_ASPECTS="edges,nodes"
export SENDS_ASPECTS="nodeAttributes"
```

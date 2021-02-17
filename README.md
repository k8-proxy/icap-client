# icap-client

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)
[![Project status](https://img.shields.io/badge/version-0.1.0-green.svg)](https://github.com/egirna/icap-client/releases)
[![GoDoc](https://godoc.org/github.com/egirna/icap-client?status.svg)](https://godoc.org/github.com/egirna/icap-client)


Talk to the ICAP servers using probably the first ICAP client package in GO!

### Installing
```console
go get -u github.com/haitham911/icap-client

```

### Usage

**Import The Package**

```go
import ic "github.com/haitham911/icap-client"

```

**Making a simple RESPMOD call**

```go

  req, err := ic.NewRequest(ic.MethodRESPMOD, "icap://<host>:<port>/<path>", httpReq, httpResp)

  if err != nil {
    log.Fatal(err)
  }

  client := &ic.Client{
		Timeout: 5 * time.Second,
	}

  resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

```

**Note**: ``httpReq`` & ``httpResp`` here are ``*http.Response`` & ``*http.Request`` respectively

**Making Tls InsecureSkipVerify call**

```go

  req, err := ic.NewRequestTLS(ic.MethodRESPMOD, "icap://<host>:<port>/<path>", httpReq, httpResp,"tls")

  if err != nil {
    log.Fatal(err)
  }

  client := &ic.Client{
		Timeout: 5 * time.Second,
	}

  resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

```


**Setting preview obtained from OPTIONS call**

```go

  optReq, err := ic.NewRequest(ic.MethodOPTIONS, "icap://<host>:<port>/<path>", nil, nil)

  if err != nil {
    log.Fatal(err)
    return
  }

  client := &ic.Client{
    Timeout: 5 * time.Second,
  }

  req.SetPreview(optReq.PreviewBytes)

  // do something with req(ICAP *Request)

```

**DEBUG Mode**

Turn on debug mode to inspect detailed & verbose logs to debug your code during development

```go
  ic.SetDebugMode(true)

```

By default the icap-client will dump the debugging logs to the standard output(stdout), but you can always add your custom writer

```go
  f, _ := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
  ic.SetDebugOutput(f)
```

For more details, see the [docs](https://godoc.org/github.com/egirna/icap-client) and [examples](examples/).


### Contributing

This package is still WIP, so totally open to suggestions. See the contributions guide [here](CONTRIBUTING.md).

### License

**icap-client** is licensed under the [Apache License](LICENSE).

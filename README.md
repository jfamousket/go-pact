# go-pact

Module to ease interaction with Pact's development server &amp; ScalableBFT

## Install

```bash
   go install github.com/jfamousket/go-pact@latest
```

## Functions

### Hashing

Create a `blake2b` hash

```go
import (
    "github.com/jfamousket/go-pact"
)

pact.CreateBlake2Hash(<cmd_string>) string
```

### Pact API functions

This functions are same as described in [pact-api-documentation](https://pact-language.readthedocs.io/en/stable/pact-reference.html?highlight=%2Fsend#rest-api)

```go
import (
    "github.com/jfamousket/go-pact"
)

pact.Send(<valid_cmd_object>, <api_host>) pact.SendResponse
pact.Local(<valid_cmd_object>, <api_host>) pact.LocalResponse
pact.Listen(<request_key>, <api_host>) pact.ListenResponse
pact.Poll(<array_of_request_keys>, <api_host>) pact.PollResponse
```

## Usage

See the `example` directory in this project, for an example TODO app.

## TODOS

- [ ] Improve docs

# Contributors

Built with :heart: by [jfamousket](https://jfamousket@gmail.com)

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
    "github.com/jfamousket/go-kadena/helpers"
)

CreateBlake2Hash(<cmd_string>) string
```

### Pact API functions

This functions are same as described in [pact-api-documentation](https://pact-language.readthedocs.io/en/stable/pact-reference.html?highlight=%2Fsend#rest-api)

```go
import (
    "github.com/jfamousket/go-kadena/fetch"
)

Send(<valid_cmd_object>, <api_host>) SendResponse
Local(<valid_cmd_object>, <api_host>) LocalResponse
Listen(<request_key>, <api_host>) ListenResponse
Poll(<array_of_request_keys>, <api_host>) PollResponse
```

## TODOS

- [ ] Key generation and manipulation functions
- [x] Create wallet Kadena blockchain
- [ ] Other expected features for blockchain library

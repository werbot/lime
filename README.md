[![Go Report Card](https://goreportcard.com/badge/github.com/werbot/lime)](https://goreportcard.com/report/github.com/werbot/lime) ![Docker](https://github.com/werbot/lime/workflows/Docker/badge.svg)

<img src="https://werbot.com/img/projects/lime.png" height="70" />


## Installation 
```
$ git clone https://github.com/werbot/lime.git
```


## Setup
1. Modify config for DB in `config/config.go`
2. Update parameters for privateKey, publicKey in file `license/license.go`

### Run server
```
$ go run main.go server 
```


### Available Commands:
- `healthcheck` : Check healthcheck
- `help` : Help about any command
- `server` : Start license server


## API list
* `GET /ping ` : Health server
* `POST /key` : Generate new license
* `GET /key/:customer_id ` : Get active license
* `PATCH  /key/:customer_id` : Update license
* `POST   /verify` : Check status license


## To-do
- [x] Generating license
- [x] Verification license
- [ ] Auto-create and install license on the client
- [ ] Command-line utility for generating key pair 
- [ ] Integration with Stripe
- [ ] Example client
- [ ] Admin console
- [ ] Support IP address check

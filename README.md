[![Go Report Card](https://goreportcard.com/badge/github.com/werbot/lime)](https://goreportcard.com/report/github.com/werbot/lime) [![CodeFactor](https://www.codefactor.io/repository/github/werbot/lime/badge)](https://www.codefactor.io/repository/github/werbot/lime) ![Docker](https://github.com/werbot/lime/workflows/Docker/badge.svg) 

<img src="https://werbot.com/assets/github/lime.png" height="70" />


## Installation 
```
$ git clone https://github.com/werbot/lime.git
```


## Setup
1. Modify config for DB in `config/config.go`
2. Update parameters for privateKey, publicKey in file `license/license.go` 
To generate new key pair use command ```go run main.go pkey```

## Run server
```
$ go run main.go server 
```

## Available Commands:
- `healthcheck` : Check healthcheck
- `help` : Help about any command
- `server` : Start license server
- `pkey` : Generating key pair


## Admin console
Link for admin console http://localhost:8080/admin/
default login - admin, password - admin
<img src="https://werbot.com/assets/github/lime/login.png" />
<img src="https://werbot.com/assets/github/lime/customers.png" />
<img src="https://werbot.com/assets/github/lime/subscriptions.png" />


## API list
* `GET      /api/ping ` : Health server
* `POST     /api/key` : Generate new license
* `GET      /api/key/:customer_id ` : Get active license
* `PATCH    /api/key/:customer_id` : Update license
* `POST     /api/verify` : Check status license


## TODO
- [x] Generating license
- [x] Verification license
- [ ] Auto-create and install license on the client
- [x] Command-line utility for generating key pair 
- [ ] Integration with Stripe
- [ ] Example client
- [x] Admin console
- [ ] Support IP address check

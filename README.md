# License server - Werbot Lime
Wery light license server


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
- [ ] Integration with Stripe
- [ ] Example client
- [ ] Admin-panel
- [ ] Support ip check
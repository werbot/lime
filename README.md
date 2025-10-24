<p align="center">
    <a href="https://werbot.github.io/lime/" target="_blank" rel="noopener">
        <img src="https://raw.githubusercontent.com/werbot/lime/v2/.github/assets/promo.png" alt="light license-key server in 1 file" />
    </a>
</p>

<p align="center">
    <a href="https://github.com/werbot/lime/releases"><img src="https://img.shields.io/github/v/release/werbot/lime?sort=semver&label=Release&color=651FFF"></a>
    &nbsp;
    <a href="/LICENSE"><img src="https://img.shields.io/badge/MIT-green.svg"></a>
    &nbsp;
    <a href="https://goreportcard.com/report/github.com/werbot/lime"><img src="https://goreportcard.com/badge/github.com/werbot/lime"></a>
    &nbsp;
    <a href="https://www.codefactor.io/repository/github/werbot/lime"><img src="https://www.codefactor.io/repository/github/werbot/lime/badge" alt="CodeFactor" /></a>
    &nbsp;
    <a href="https://github.com/werbot/lime"><img src="https://img.shields.io/badge/backend-go-orange.svg"></a>
    &nbsp;
    <a href="   https://github.com/werbot/lime/blob/v2/go.mod"><img src="https://img.shields.io/github/go-mod/go-version/werbot/lime?color=7fd5ea"></a>
    &nbsp;
    <a href="https://twitter.com/werbot_"><img src="https://img.shields.io/twitter/follow/werbot_?style=social"></a>
</p>

## 🍋‍🟩&nbsp;&nbsp;What is lime?

Light license-key server in 1 file

> [!WARNING]
> Current major version is zero (`v0.x.x`) to accommodate rapid development and fast iteration while getting early feedback from users. Please keep in mind that lime is still under active development and therefore full backward compatibility is not guaranteed before reaching v1.0.0.

## 🏆&nbsp;&nbsp;Features

- **License Key Management** - Generate, validate, and manage license keys with Ed25519 cryptographic signatures
- **Multi-tier Licensing** - Support for different license types (Trial, Standard, Enterprise, etc.) with customizable limits
- **Customer Portal** - Self-service portal for customers to manage their licenses
- **Admin Dashboard** - Comprehensive admin panel built with Vue3 and TailwindCSS
- **Payment Integration** - Built-in support for Stripe payment processing
- **Database Support** - SQLite (default) and PostgreSQL support
- **Geo-location Tracking** - Track license usage by country (MaxMind, IPinfo, GeoOpen)
- **JWT Authentication** - Secure API access with JWT tokens
- **Email Notifications** - Automated email notifications for license events
- **Multi-currency Support** - Handle payments in multiple currencies (USD, EUR, GBP, etc.)
- **Audit Logging** - Track all administrative actions
- **RESTful API** - Complete API for integration with your applications
- **Single Binary** - All features packed into one executable file
- **Docker Support** - Easy deployment with Docker and Docker Compose

## ⬇️&nbsp;&nbsp;Installation

### Binary Installation

Download the latest release for your platform from [GitHub Releases](https://github.com/werbot/lime/releases):

```bash
# Linux
wget https://github.com/werbot/lime/releases/latest/download/lime-linux-amd64
chmod +x lime-linux-amd64
mv lime-linux-amd64 /usr/local/bin/lime

# macOS
wget https://github.com/werbot/lime/releases/latest/download/lime-darwin-amd64
chmod +x lime-darwin-amd64
mv lime-darwin-amd64 /usr/local/bin/lime
```

### Docker Installation

```bash
docker pull ghcr.io/werbot/lime:latest

# Run with Docker
docker run -d \
  -p 8088:8088 \
  -v $(pwd)/data:/data \
  --name lime \
  ghcr.io/werbot/lime:latest
```

### Docker Compose

```bash
# Clone the repository
git clone https://github.com/werbot/lime.git
cd lime

# Start with Docker Compose
docker compose -f docker/docker-compose.yaml up -d
```

### Build from Source

```bash
# Clone the repository
git clone https://github.com/werbot/lime.git
cd lime

# Build frontend
yarn --cwd ./web run build

# Build and run
cd ./cmd
go build -o lime main.go
./lime serve
```

## ⬇️&nbsp;&nbsp;Updating

### Binary Update

1. Download the latest release
2. Stop the running instance
3. Replace the old binary with the new one
4. Restart the service

```bash
# Backup your data
cp -r ./lime_base ./lime_base.backup
cp -r ./lime_keys ./lime_keys.backup

# Download and replace binary
wget https://github.com/werbot/lime/releases/latest/download/lime-linux-amd64
chmod +x lime-linux-amd64
mv lime-linux-amd64 /usr/local/bin/lime

# Restart service
systemctl restart lime
```

### Docker Update

```bash
# Pull the latest image
docker pull ghcr.io/werbot/lime:latest

# Stop and remove old container
docker stop lime
docker rm lime

# Start with new image
docker run -d \
  -p 8088:8088 \
  -v $(pwd)/data:/data \
  --name lime \
  ghcr.io/werbot/lime:latest
```

## 🚀&nbsp;&nbsp;Getting started

### 1. Generate Configuration File

```bash
lime gen --config
```

This creates `lime.toml` with default settings. Edit the file to configure:
- HTTP address and port (default: `0.0.0.0:8088`)
- Admin credentials (default: `admin@mail.com` / `Pass123`)
- Database settings (SQLite or PostgreSQL)
- Geo-location provider
- Email settings

### 2. Initialize Keys

Keys are generated automatically on first run, or manually:

```bash
# Generate JWT keys
lime gen --jwt

# Generate license master keys
lime gen --license
```

### 3. Start the Server

```bash
lime serve
```

The server will start on `http://0.0.0.0:8088`

### 4. Access Admin Panel

Navigate to `http://localhost:8088/_/` and login with:
- **Email:** admin@mail.com
- **Password:** Pass123

> [!IMPORTANT]
> Change the default admin credentials immediately after first login!

### 5. Create License Patterns

1. Go to Admin Panel → Patterns
2. Create license patterns (e.g., Trial, Standard, Enterprise)
3. Define limits, pricing, and duration for each pattern

### 6. Generate Licenses

1. Add customers via Admin Panel → Customers
2. Create payments for customers
3. Generate license keys from payments

## 📚&nbsp;&nbsp;Commands

### `lime serve`

Start the license server.

```bash
lime serve
```

Server will start on the address defined in `lime.toml` (default: `0.0.0.0:8088`)

### `lime gen`

Generate configuration files and cryptographic keys.

```bash
# Generate configuration file
lime gen --config

# Generate JWT key pair
lime gen --jwt

# Generate license master key pair
lime gen --license
```

### Helper Scripts

Development helper scripts are available in the `./scripts` folder:

```bash
# Install or update Go version
./scripts/golang

# Database migrations
./scripts/migrate dev up      # Apply migrations with fixtures
./scripts/migrate dev down    # Rollback migrations
./scripts/migrate up          # Apply migrations only
./scripts/migrate down        # Rollback migrations

# Optimize SQLite database
./scripts/sqlite

# Generate keys (alternative to 'lime gen')
./scripts/gen

# Clean up hung processes
./scripts/clear
```

## 🏦&nbsp;&nbsp;Adding payment systems

Currently, Lime supports **Stripe** for payment processing. Additional payment providers can be added by extending the payment handler.

### Stripe Integration

1. Get your Stripe API keys from [Stripe Dashboard](https://dashboard.stripe.com/apikeys)

2. Configure Stripe in your application by adding webhook URLs:
   - Webhook endpoint: `https://yourdomain.com/_/api/webhook/stripe`

3. Handle payment events in the admin panel:
   - Go to Admin Panel → Payments
   - View and manage payment transactions
   - Monitor payment status (Paid, Unpaid, Processed, Canceled, Failed)

### Adding Custom Payment Providers

To add support for additional payment providers:

1. Update `internal/models/payments.go`:
   ```go
   const (
       _ PaymentProvider = iota
       NONE
       STRIPE
       YOUR_PROVIDER  // Add your provider
   )
   ```

2. Implement payment webhook handler in `internal/handlers/admin/payment.go`

3. Update frontend in `web/src/utils/index.ts`:
   ```typescript
   export const paymentProvidersObj = [
       { name: "None", color: "gray" },
       { name: "Stripe", color: "purple" },
       { name: "YourProvider", color: "blue" },  // Add your provider
   ];
   ```



## 🧩&nbsp;&nbsp;For developers

The backend is developed in Go language. The frontend (admin site and base site) operates on the Vue3 and TailwindCSS.  

There are a number of scripts (in the ./scripts folder) that simplify development:  
`./scripts/golang` - Installs or updates a previously installed version of go (if needed).  
`./scripts/migration` - Helps to work with migrations. For instance, the `./scripts/migration dev up` command will apply new migrations from folder ./migrations, then implement the migrations from folder ./fixtures.  
`./scripts/sqlite` - Optimizes the existing database.  
`./scripts/gen` - Generate JWT or master-License keys.  
`./scripts/clear` - Removing hung golang or vite processes.  

First run:  
1. `yarn --cwd ./web run build` - This is necessary in order to be able to compile and run a go app.  
2. `cd ./cmd/ && go run main.go gen --config` - (if need) To save the configuration file with default parameters for further modification  
3. `cd ./cmd/ && go run main.go serve` - Launch the license server  
4. `yarn --cwd ./web run dev` - (if need) If you need to change the server for if you're going to modify the UI admin or manager panel, it will launch a dev environment with Vite. 

> [!NOTE] 
> I recommend running the `./scripts/migration dev up` command. It will add test data to the database, which makes it easier to work with. For example, it will create products, transfer test images and create a test user for access to the admin panel on address http://0.0.0.0:8088/_/:  
> login - admin@mail.com  
> password - Pass123

## 👍&nbsp;&nbsp;Contribute

If you want to say **thank you** and/or support the active development of `lime`:

1. Add a [GitHub Star](https://github.com/werbot/lime/stargazers) to the project.
2. Tweet about the project [on your Twitter](https://twitter.com/intent/tweet?text=%F0%9F%8D%8B%E2%80%8D%F0%9F%9F%A9%20light%20license-key%20server%20in%201%20file%20on%20%23Go%20https%3A%2F%2Fgithub.com%2Fwerbot%2Flime)


3. Write a review or tutorial on [Medium](https://medium.com/), [Dev.to](https://dev.to/) or personal blog.
4. Support the project by donating a [cup of coffee](https://github.com/sponsors/werbot).

You can learn more about how you can contribute to this project in the [contribution guide](https://github.com/werbot/lime/blob/v2/.github/CONTRIBUTING.md).
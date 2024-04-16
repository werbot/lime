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
> Current major version is zero (`v0.x.x`) to accommodate rapid development and fast iteration while getting early feedback from users. Please keep in mind that litecart is still under active development and therefore full backward compatibility is not guaranteed before reaching v1.0.0.

## 🏆&nbsp;&nbsp;Features
... coming soon ...  

## ⬇️&nbsp;&nbsp;Installation
... coming soon ...  

## ⬇️&nbsp;&nbsp;Updating
... coming soon ...  

## 🚀&nbsp;&nbsp;Getting started
... coming soon ...  

## 📚&nbsp;&nbsp;Commands
... coming soon ...  

## 🏦&nbsp;&nbsp;Adding payment systems
... coming soon ...  

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
# Habit Tracker

![GitHub go.mod Go version (branch & subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/infamous55/habit-tracker/master?style=flat-square)
![](https://tokei.rs/b1/github/infamous55/habit-tracker?style=flat-square)

A GraphQL API for tracking your habits.

## About

The GraphQL API is written in Golang and relies on the [gqlgen](https://gqlgen.com/) library and the [echo](https://echo.labstack.com/) framework. It interacts with MongoDB using the [official drivers](https://github.com/mongodb/mongo-go-driver). The data layer is unit tested.

The GraphQL schema, encompassing all available types, queries, and mutations, can be located within the `internal/graphql/schema.graphql` file. Additionally, the API is secured with JWT-based authentication.

The project includes a CI/CD worflow built with [GitHub Actions](https://github.com/features/actions). Infrastructured is created in [DigitalOcean](https://www.digitalocean.com/) using [Pulumi](https://www.pulumi.com/), and the API is deployed as a [Docker](https://www.docker.com/) image.

## To Do

- [ ] Improve the error messages.
- [ ] Implement [validator](https://pkg.go.dev/github.com/go-playground/validator/v10#section-readme).

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

Distributed under the GNU General Public License v3.0. See `LICENSE` for more information.

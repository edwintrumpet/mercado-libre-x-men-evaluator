<h1 align="center">X-men evaluator</h1>

<p align="center">
  <img src="https://github.com/edwintrumpet/experiment_golang_ci_server/workflows/Tests/badge.svg" alt="Tests badge">
</p>

<p align="center">
  <img src="https://simpleicons.org/icons/githubactions.svg" alt="github actions" width="40" height="40"/>
  <img src="https://raw.githubusercontent.com/devicons/devicon/2809b567852a4648062a2d3e7c1c531367458c0b/icons/go/go-original.svg" alt="go" width="40" height="40"/>
</p>

This program have some DNA secuences and uses the IsMutant function to evaluate if DNA secuence belongs to a mutant.

If DNA secuence belongs to a mutant IsMutant function will return a true, if DNA secuence does not belong a mutant so it will return a false.

## Scripts

- **`./run.sh`** => Execute API in develop mode
- **`./run.sh build`** => Build API
- **`./run.sh test`** => Run tests
- **`./run.sh coverage`** => Run tests coverage

## Develop

To run API in develop mode you must to have Go installed and execute the following command into the project folder

```shell
./run.sh
```

API will be listening in http://localhost:8080

## Deploy

To build API you must have Go installed and execute the following command into the project folder

```shell
./run.sh build
```

## Author

Edwin García  
spark.com.co@gmail.com

## License

[MIT](./LICENSE)

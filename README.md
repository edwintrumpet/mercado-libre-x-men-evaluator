<h1 align="center">X-men evaluator</h1>

<p align="center">
  <img src="https://github.com/edwintrumpet/experiment_golang_ci_server/workflows/Tests/badge.svg" alt="Tests badge">
</p>

<p align="center">
  <img src="https://raw.githubusercontent.com/devicons/devicon/2809b567852a4648062a2d3e7c1c531367458c0b/icons/docker/docker-original.svg" alt="docker" width="40" height="40"/>
  <img src="https://simpleicons.org/icons/githubactions.svg" alt="github actions" width="40" height="40"/>
  <img src="https://raw.githubusercontent.com/devicons/devicon/2809b567852a4648062a2d3e7c1c531367458c0b/icons/go/go-original.svg" alt="go" width="40" height="40"/>
</p>

This program have some DNA secuences and uses the IsMutant
function to evaluate if DNA secuence belongs to a mutant.

If DNA secuence belongs to a mutant IsMutant function will
return a true, if DNA secuence does not belong a mutant so
it will return a false.

## Scripts

- **`./run.sh`** => Execute API in develop mode
- **`./run.sh build`** => Build API
- **`./run.sh test`** => Run tests
- **`./run.sh coverage`** => Run tests coverage

## Develop

### Using Docker

To run API in develop mode you must to have Docker installed

If you run this project first time execute the following command
to build the image

```shell
./run.sh buildDev
```

To run API in develop mode

```shell
./run.sh
```

API will be listening on http://localhost

To remove docker container

```shell
./run.sh down
```

### Without Docker

To run API in develop mode you must to have Go installed

To run API in develop mode

```shell
./run.sh run
```

API will be listening on http://localhost

## Deploy

This API uses amazon ECS to deploy it, you need to have a
ECR repository, and create an ECS cluster with a task to deploy
the image.

To make push in repository you can set the environment variables
in `.env` taking `.env.example` as example and set
`~/.aws/credentials` with your IAM access key

After that you need to autenticate to ECR

```shell
./run.sh ecrAuth
```

And make push

```shell
./run.sh push
```

After that you can run task in cluster to deploy the image and API will
be available in public IP.

If you want just obtain the docker image to production

```shell
./run.sh build
```

This image is running on port 80

## Author

Edwin García  
spark.com.co@gmail.com

## License

[MIT](./LICENSE)

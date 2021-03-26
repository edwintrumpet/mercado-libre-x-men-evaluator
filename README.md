<h1 align="center">X-men evaluator</h1>

<p align="center">
  <img src="https://github.com/edwintrumpet/experiment_golang_ci_server/workflows/Tests/badge.svg" alt="Tests badge">
</p>

<p align="center">
  <img src="https://raw.githubusercontent.com/devicons/devicon/2809b567852a4648062a2d3e7c1c531367458c0b/icons/amazonwebservices/amazonwebservices-original.svg" alt="aws" width="40" height="40"/>
  <img src="https://raw.githubusercontent.com/devicons/devicon/2809b567852a4648062a2d3e7c1c531367458c0b/icons/docker/docker-original.svg" alt="docker" width="40" height="40"/>
  <img src="https://simpleicons.org/icons/githubactions.svg" alt="github actions" width="40" height="40"/>
  <img src="https://raw.githubusercontent.com/devicons/devicon/2809b567852a4648062a2d3e7c1c531367458c0b/icons/go/go-original.svg" alt="go" width="40" height="40"/>
</p>

API that takes some DNA secuences and evaluates if DNA secuence
belongs to a mutant.

## Routes

API available on http://3.238.192.139/

Request instructions

| método | endpoint  | parámetros |            valores            |
| :----: | :-------: | :--------: | :---------------------------: |
|  POST  | `/mutant` |    body    | See the following json format |
|  GET   | `/stats`  |            |                               |

```json
{
  "dna": ["ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"]
}
```

Possible answers

- `/mutant`  
  Returns a status code 200 with a succes message if DNA belongs to
  a mutant

  ```json
  {
    "message": "OK"
  }
  ```

  Returns a status code 403 with a Forbidden message if DNA does not
  belongs to a mutant

  ```json
  {
    "message": "Forbidden"
  }
  ```

- `/stats`  
  Returns a status code 200 with a json that contains info about requests
  amount
  ```json
  {
    "count_mutant_dna": 9,
    "count_human_dna": 10,
    "ratio": 0.9
  }
  ```

## Scripts

- **`./run.sh`** => Execute API and database in develop mode
- **`./run.sh buildDev`** => Build Docker image to develop
- **`./run.sh down`** => Down develop container
- **`./run.sh build`** => Build Docker image to production
- **`./run.sh ecrAuth`** => Authenticate in ECR to make push
- **`./run.sh push`** => Push docker image to ECR
- **`./run.sh runDB`** => Enter the database to make SQL instructions with psql
- **`./run.sh test`** => Run tests
- **`./run.sh coverage`** => Run tests coverage

## Develop

### Using Docker

To run API in develop mode you must to have Docker installed

First you must to create `.env` file taking `.env.example` as example

If you run this project first time execute the following command
to build the image

```shell
./run.sh buildDev
```

To run API in develop mode

```shell
./run.sh
```

Database will be executed automatically and API will be listening on http://localhost

In another terminal execute

```shell
./run.sh runDB
```

Create an user with the same name and password that you have in `.env`

```sql
CREATE USER <username> WITH PASSWORD '<password>';
```

Grant all privileges to user

```sql
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO <username>;
```

Create the mutants database

```sql
CREATE DATABASE mutants OWNER <username>
```

Use mutants database

```psql
\c mutants
```

Create **stats** table

```sql
CREATE TABLE IF NOT EXISTS stats (
  id SERIAL PRIMARY KEY,
  name VARCHAR(20) NOT NULL,
  value INT DEFAULT 0
);
```

```sql
ALTER TABLE stats
  ADD CONSTRAINT uq_name_stat
  UNIQUE (name);
```

Populate database

```sql
INSERT INTO stats (name) VALUES
  ('count_mutant_dna'),
  ('count_human_dna');
```

You're ready to develop!

To remove docker container

```shell
./run.sh down
```

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

Database is hosted in AWS RDS

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

# Kargo

Simple wrapper around [Docker](https://www.docker.com/docker-engine), [Docker Compose](https://www.docker.com/docker-compose) and [Docker Machine](https://www.docker.com/docker-machine) to do basic actions more easily.

## Build

```bash
make
```

## Usage

```
NAME:
   kargo - using docker should be easy

USAGE:
   kargo [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
   destroy	destroy the docker containers
   machine	start the docker machine if required
   open, o	open the project in your browser
   shell, s	open a shell into a container
   start	start the docker containers
   help, h	Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --machine, -m "default"	name of the docker machine to use
   --project-name, -p 		name of the project
   --help, -h			show help
   --version, -v		print the version
```

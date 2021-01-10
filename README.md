# dev-cli

[![dev-cli](https://github.com/Kaiser925/dev-cli/workflows/dev-cli/badge.svg)](https://github.com/Kaiser925/dev-cli/actions)

dev-cli is a commandline interface tool for building development environment.

It encapsulates the commands of docker and docker-compose, enabling developers to quickly build a local development environment

## Usage

> If you want to use dev-cli, you should install **docker** and **docker-compose** on your system. 

### install

```bash
go get https://github.com/Kaiser925/dev-cli
```

### Commands

Show help messages

```bash
dev-cli --help
```

#### create

Command "create" is used to create new resource.

examples:

Create a new mongo replica set.

```bash
dev-cli create mongors
```

#### delete

Command "delete" is used to delete resource.

## Next step

- [ ] Implement Docker client.

- [ ] Implement Docker-compose.

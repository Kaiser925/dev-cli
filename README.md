# devctl

[![devctl](https://github.com/Kaiser925/devctl/workflows/devctl/badge.svg)](https://github.com/Kaiser925/devctl/actions)

devctl is a tool for building development environment.

It encapsulates the commands of docker and docker-compose, enabling developers to quickly build a local development environment

## Usage

> If you want to use devctl, you should install **docker** and **docker-compose** on your system. 

### install

```bash
go get https://github.com/Kaiser925/devctl
```

### Commands

Show help messages

```bash
devctl --help
```

#### create

Command "create" is used to create new resource.

examples:

1. Create a new mongo replica set.

```bash
devctl create mongors
```

2. Create a new mongo user for the database.

```bash
devctl create mongousr <db> <user> <password>
```

#### delete

Command "delete" is used to delete resource.

## Next step

- [ ] Implement Docker client.

- [ ] Implement Docker-compose.

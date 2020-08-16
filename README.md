# `hikctl`

[![PkgGoDev](https://pkg.go.dev/badge/github.com/loozhengyuan/hikctl)](https://pkg.go.dev/github.com/loozhengyuan/hikctl)
[![Go Report Card](https://goreportcard.com/badge/github.com/loozhengyuan/hikctl)](https://goreportcard.com/report/github.com/loozhengyuan/hikctl)
![ci](https://github.com/loozhengyuan/hikctl/workflows/ci/badge.svg)

> **Important**: This tool is still under heavy development so only a basic set of configuration variables are available. This will be gradually expanded over time.

Configuration management tool for Hikvision devices.

The `hikctl` command line tool, inspired by the [`kubectl`](https://kubernetes.io/docs/reference/kubectl/overview/) tool for Kubernetes, aims to provide a declarative way to manage configuration on Hikvision devices.

## Install

### Compile from source

To compile `hikctl` from source, clone the git repository.

```shell
git clone https://github.com/loozhengyuan/hikctl.git && cd hikctl
```

Next, run `go install`, which will compile the binary and install it to the `$GOPATH/bin` directory.

```shell
go install
```

Check that the binary is correctly installed by running `hikctl version`

```shell
hikctl version
```

## Usage

### Applying a configuration file

First, start by creating a set of configuration variables in a `yaml` file.

```yaml
version: v0
```

To apply the configuration file, run the `hikctl apply` command as shown below.

```shell
hikctl apply hikvision.yaml
```

## License

[GPL-3.0](https://choosealicense.com/licenses/gpl-3.0/)

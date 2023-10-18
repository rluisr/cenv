cenv
====

[![release](https://github.com/rluisr/cenv/actions/workflows/deploy.yml/badge.svg)](https://github.com/rluisr/cenv/actions/workflows/deploy.yml)

cenv is a version manager for [AWS Copilot CLI](https://github.com/aws/copilot-cli).

```shell
$ cat .copilot-version
1.30.1

$ sudo cenv #or cenv --user
copilot version: v1.30.1
```

Installation
------------

| Platform           | Command to install                                                                                                                             |
|--------------------|------------------------------------------------------------------------------------------------------------------------------------------------
| macOS M1           | `curl -Lo cenv https://github.com/rluisr/cenv/releases/latest/download/cenv-darwin-arm64 && chmod +x cenv && sudo mv cenv /usr/local/bin/cenv` |
| macOS Intel        | `curl -Lo cenv https://github.com/rluisr/cenv/releases/latest/download/cenv-darwin-amd64 && chmod +x cenv && sudo mv cenv /usr/local/bin/cenv` |
| Linux x86 (64-bit) | `curl -Lo cenv https://github.com/rluisr/cenv/releases/latest/download/cenv-linux-amd64 && chmod +x cenv && sudo mv cenv /usr/local/bin/cenv`  |
| Linux (ARM)        | `curl -Lo cenv https://github.com/rluisr/cenv/releases/latest/download/cenv-linux-arm64 && chmod +x cenv && sudo mv cenv /usr/local/bin/cenv`  |

Usage
-----

`--user` option install copilot-cli to current directory instead of `/usr/local/bin`.

Limitations
-----------

cenv does not have the bash affinity of nvm or tfenv.

This tool simply downloads and installs the binary version written in the current .copilot-version.
# Introduction

When developing softwars, I often have to create some random stuffs and confiugraion over and over, so I make this
small CLI tool to make my life easier.

This tool call quickgen, it supports the following command, please check command help for more detail.

* Generate UUID
* Generate MD5
* Encode/Decode base64
* Format JSON
* Generate Dockerfile for Golang app
* Generate Makefile for Golang app

## Installation

Download the prebuild binary for your OS [Relase](https://github.com/phucngodev/quickgen/releases/)
and put it in your PATH if you want for convenience


## Usage

```bash
CLI tool generate random stuffs

Usage:
  quickgen [command]

Available Commands:
  base64      base64 sub command
  completion  Generate the autocompletion script for the specified shell
  docker      Docker subcommands
  help        Help about any command
  json        JSON subcommands
  makefile    Generate Makefile for some languages
  md5         Generate md5 for given string
  uuid        Generate uuid code
  version     Show quickgen version info

Flags:
  -h, --help     help for quickgen
  -t, --toggle   Help message for toggle

Use "quickgen [command] --help" for more information about a command.

```

# License MIT

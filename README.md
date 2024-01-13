<h1 align="center">gh-q</h1>

<p align="center">
<a href="https://github.com/koki-develop/gh-q/releases/latest"><img src="https://img.shields.io/github/v/release/koki-develop/gh-q" alt="GitHub release (latest by date)"></a>
<a href="https://goreportcard.com/report/github.com/koki-develop/gh-q"><img src="https://goreportcard.com/badge/github.com/koki-develop/gh-q" alt="Go Report Card"></a>
<a href="./LICENSE"><img src="https://img.shields.io/github/license/koki-develop/gh-q" alt="LICENSE"></a>
</p>

<p align="center">
gh extension to manage GitHub repositories like <a href="https://github.com/x-motemen/ghq">ghq</a>
</p>

## Contents

- [Installation](#installation)
- [Configuration](#configuration)
  - [`ghq.root`](#ghqroot)
- [Usage](#usage)
  - [`create` - Create a new repository on local](#create---create-a-new-repository-on-local)
  - [`get` - Clone repository](#get---clone-repository)
  - [`list` - List managed repositories](#list---list-managed-repositories)
  - [`remove` - Remove repository from local](#remove---remove-repository-from-local)
- [LICENSE](#license)

## Installation

```console
$ gh extension install koki-develop/gh-q
```

## Configuration

### `ghq.root`

`ghq.root` is the root directory of the repository managed by.  
If not set, it will be `~/ghq` by default.

```console
$ git config --global ghq.root ~/your/ghq
```

`GHQ_ROOT` environment variable can also be used to set it.

```console
$ export GHQ_ROOT=~/your/ghq
```

## Usage

```console
$ gh q --help
gh extension to manage GitHub repositories like `ghq`

Usage:
  gh q [command]

Available Commands:
  create      Create a new repository on local
  get         Clone repository
  help        Help about any command
  list        List managed repositories
  remove      Remove repository from local

Flags:
  -h, --help   help for q
```

### `create` - Create a new repository on local

```console
$ gh q create --help
Create a new repository on local.

Usage:
  gh q create OWNER/REPO|REPO... [flags]

Aliases:
  create, c

Flags:
  -h, --help   help for create
```

### `get` - Clone repository

```console
$ gh q get --help
Clone repository.

Usage:
  gh q get OWNER/REPO|REPO... [flags]

Aliases:
  get, g

Flags:
  -h, --help   help for get
```

### `list` - List managed repositories

```console
$ gh q list --help
List managed repositories.

Usage:
  gh q list [flags]

Aliases:
  list, ls

Flags:
  -f, --filter      filter by fuzzy search
  -p, --full-path   print full path
  -h, --help        help for list
  -m, --multiple    allow multiple selection (only available with --filter)
```

### `remove` - Remove repository from local

```console
$ gh q remove --help
Remove repository from local.

Usage:
  gh q remove OWNER/REPO|REPO... [flags]

Aliases:
  remove, rm

Flags:
  -f, --force   Remove without confirmation
  -h, --help    help for remove
```

## LICENSE

[MIT](./LICENSE)

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
- [Usage](#usage)
  - [`clone` - Clone repository](#clone---clone-repository)
  - [`list` - List managed repositories](#list---list-managed-repositories)
- [LICENSE](#license)

## Installation

```console
$ gh extension install koki-develop/gh-q
```

## Usage

```console
$ gh q --help
gh extension to manage GitHub repositories like `ghq`

Usage:
  gh q [command]

Available Commands:
  clone       Clone repository
  help        Help about any command
  list        List managed repositories

Flags:
  -h, --help   help for q
```

### `clone` - Clone repository

```console
$ gh q clone --help
Clone repository.

Usage:
  gh q clone OWNER/REPO [flags]

Aliases:
  clone, c

Flags:
  -h, --help   help for clone
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

## LICENSE

[MIT](./LICENSE)

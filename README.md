# mgit

mgit (multi git) lets you manage multiple git repositories.

![View outdated git tags](docs/view-outdated-tags.png)

## Installation

```shell
go get -u github.com/akyoto/mgit
```

## Usage

### View tags

```shell
mgit -tags
```

Lets you view outdated or untagged git repositories by recursively searching everything in the working directory.

### Run a command in each repository

```shell
mgit -run "git pull"
```

```shell
mgit -run "go get -u"
```

```shell
mgit -run "go mod tidy"
```

The `-run` flag lets you specify a command to run in every git repository.

The command will be executed in parallel (one async routine per repository).

## FAQ

### Why is this needed?

When you manage multiple repositories, it is quite difficult to keep in mind which git repositories were already tagged. This tool helps you remember if your last commit is already tagged or not.

### What does "not tagged" mean?

It means that the repository doesn't have any tags.

### What does "outdated" mean?

It means that your last commit hasn't been tagged yet and users of your repository might still be on an outdated version.

### Can mgit automatically tag the outdated ones?

Not yet, although this feature is work in progress. Feel free to send a PR if you need this functionality.

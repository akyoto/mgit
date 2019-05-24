![mgit](docs/mgit.png)

[![Godoc][godoc-image]][godoc-url]
[![Report][report-image]][report-url]
[![Tests][tests-image]][tests-url]
[![Coverage][coverage-image]][coverage-url]
[![Sponsor][sponsor-image]][sponsor-url]

mgit (multi git) lets you manage multiple git repositories.

## Installation

```shell
go get -u github.com/akyoto/mgit/...
```

## Usage

### Increment all semver tags

![Increment semver minor version in all tags](docs/multi-tag-semver.png)

```shell
mgit -tag +0.0.1
```

```shell
mgit -tag +0.0.1 -dry
```

Adds the given version increment to each tag. Specify `-dry` to stop the actual tagging and only display the diff to confirm the changes. This will never increment repositories where the last commit is already tagged / up to date.

### View tags

![View git tags](docs/view-tags.png)

```shell
mgit -tags
```

Lets you view outdated or untagged git repositories by recursively searching everything in the working directory.

### Run a command

![Run command](docs/run-command.png)

```shell
mgit -run "go get -u"
```

```shell
mgit -run "npm update"
```

The `-run` flag lets you specify a command to run in every git repository. The command will be executed in parallel (one async routine per repository).

### Setting working directory

```shell
mgit -root ~/ -tags
```

Use `-root` to use a different directory than the current working directory.

## FAQ

### Why is this needed?

* Increase semver tags of multiple repositories
* Update dependencies of multiple repositories
* Get information about multiple repositories
* See if your last commits are already tagged or not
* Runs every command in parallel which makes it pretty fast

### What does "not tagged" mean?

It means that the repository doesn't have any tags.

### What does "outdated" mean?

It means that your last commit hasn't been tagged yet and users of your repository might still be on an outdated version.

## Style

Please take a look at the [style guidelines](https://github.com/akyoto/quality/blob/master/STYLE.md) if you'd like to make a pull request.

## Sponsors

| [![Cedric Fung](https://avatars3.githubusercontent.com/u/2269238?s=70&v=4)](https://github.com/cedricfung) | [![Scott Rayapoullé](https://avatars3.githubusercontent.com/u/11772084?s=70&v=4)](https://github.com/soulcramer) | [![Eduard Urbach](https://avatars3.githubusercontent.com/u/438936?s=70&v=4)](https://twitter.com/eduardurbach) |
| --- | --- | --- |
| [Cedric Fung](https://github.com/cedricfung) | [Scott Rayapoullé](https://github.com/soulcramer) | [Eduard Urbach](https://eduardurbach.com) |

Want to see [your own name here?](https://github.com/users/akyoto/sponsorship)

[godoc-image]: https://godoc.org/github.com/akyoto/mgit?status.svg
[godoc-url]: https://godoc.org/github.com/akyoto/mgit
[report-image]: https://goreportcard.com/badge/github.com/akyoto/mgit
[report-url]: https://goreportcard.com/report/github.com/akyoto/mgit
[tests-image]: https://cloud.drone.io/api/badges/akyoto/mgit/status.svg
[tests-url]: https://cloud.drone.io/akyoto/mgit
[coverage-image]: https://codecov.io/gh/akyoto/mgit/graph/badge.svg
[coverage-url]: https://codecov.io/gh/akyoto/mgit
[sponsor-image]: https://img.shields.io/badge/github-donate-green.svg
[sponsor-url]: https://github.com/users/akyoto/sponsorship

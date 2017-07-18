# Dave

Dave cuts VoIPGRID releases.

## Status

Maintained

## Usage

### Requirements

You'll need a Github access token, which you can generate
[here](https://github.com/settings/tokens). Dave only needs `repo` scope to
work.

### Installation

Quite simple: get [Go](https://golang.org/dl), set a `$GOPATH`, and run

    go get github.com/VoIPGRID/dave

### Running

Make sure `$GOPATH/bin` is in your `$PATH`.

    $ dave -h
    Usage of dave:
      -owner string
        	github owner to find repo under (default "VoIPGRID")
      -prefix string
        	branch name prefix (default "release-")
      -repo string
        	github repository to bump version of (default "voipgrid")
      -token string
        	github access token

All flags have an environment variable that sets the default, so setting
`DAVE_REPO` also changes the repo Dave looks at. Flags override environment
variables.

## Contributing

See the [CONTRIBUTING.md](CONTRIBUTING.md) file on how to contribute to this project.

## Contributors

See the [CONTRIBUTORS.md](CONTRIBUTORS.md) file for a list of contributors to the project.

## Roadmap

### Changelog

The changelog can be found in the [CHANGELOG.md](CHANGELOG.md) file.

### In progress

 * Maintaining

### Future

 * Nothing planned

## Get in touch with a developer

If you want to report an issue see the [CONTRIBUTING.md](CONTRIBUTING.md) file for more info.

We will be happy to answer your other questions at opensource@wearespindle.com

## License

Dave is made available under the MIT license. See the [LICENSE file](LICENSE) for more info.

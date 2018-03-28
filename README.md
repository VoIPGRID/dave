# dave
Dave cuts VoIPGRID releases.

## Installation

Quite simple: get [Go](https://golang.org/dl), set a `$GOPATH`, and run

    go get github.com/VoIPGRID/dave

## Use

Make sure `$GOPATH/bin` is in your `$PATH`.

    $ dave -h
    Usage of dave:
      -branch string
        	branch to base release branch on (default "develop")
      -dryrun
        	don't actually create the branch, just print
      -owner string
        	github owner to find repo under (default "VoIPGRID")
      -prefix string
        	branch name prefix (default "release-")
      -repo string
        	github repository to bump version of (default "voipgrid")
      -token string
        	github access token

You'll need a Github access token, which you can generate
[here](https://github.com/settings/tokens). Dave only needs `repo` scope to
work.

All flags have an environment variable that sets the default, so setting
`DAVE_REPO` also changes the repo Dave looks at. Flags override environment
variables.

go.travis [![Build Status](https://secure.travis-ci.org/daaku/go.travis.png)](http://travis-ci.org/daaku/go.travis)
=========

A script to enable [Travis CI](http://travis-ci.org/) integration for
[Go](http://golang.org/) projects. This is augments the stock Travis Go setup
to automatically pull test dependencies and also allow for proper use of
absolutely referenced dependencies that are in the same repository. Once your
project has some tests, setup a `.travis.yml` config as such:

```yaml
language: go
go:
  - 1.1
script:
 - curl https://raw.github.com/daaku/go.travis/master/install | sh
```

And you're all set.

Example config:
https://github.com/daaku/go.travis/blob/master/.travis.yml

Example run:
https://travis-ci.org/daaku/go.travis

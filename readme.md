go.travis [![Build Status](https://secure.travis-ci.org/nshah/go.travis.png)](http://travis-ci.org/nshah/go.travis)
=========

A script to enable [Travis CI](http://travis-ci.org/) integration for
[Go](http://golang.org/) projects. Once your project has some tests,
setup a `.travis.yml` config as such:

```yaml
language: node_js
script:
 - curl https://raw.github.com/nshah/go.travis/master/install | sh
```

And you're all set.

Example config:
https://github.com/nshah/go.travis/blob/master/.travis.yml

Example run:
http://travis-ci.org/#!/nshah/go.travis

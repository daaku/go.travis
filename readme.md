go.travis
=========

A script to enable [Travis CI](http://travis-ci.org/) integration for
[Go](http://golang.org/) projects. Once you project has some tests,
setup a `.travis.yml` config as such:

```yaml
language: erlang
script:
 - curl https://raw.github.com/nshah/go.travis/master/install | sh
```

And you're all set.

Example config:
https://github.com/nshah/go.freeport/blob/master/.travis.yml

Example run:
http://travis-ci.org/#!/nshah/go.freeport/builds/1093516

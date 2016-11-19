dns [![Build Status](https://travis-ci.org/voutasaurus/dns.svg?branch=master)](https://travis-ci.org/voutasaurus/dns)
=======

dns is a tool to perform dns lookup.

You should probably just `alias dns = dig +short` instead. I still use this to find out what Go's dns is going to give me in more complicated programs because it works slightly differently to dig.

Usage
=====

```
$ dns google.com
IP addrs: [2607:f8b0:4005:804::200e 216.58.192.14]
CNAME: google.com.
```

Installation
============

```
$ go get github.com/voutasaurus/dns
```

Contributions
=============

Issues welcome. PRs welcome.

sslciphersuitescheck is a simple program to test which clients suppport or doesn't support your ssl ciphersuites, by use [https://api.ssllabs.com/api/v3/getClients](https://api.ssllabs.com/api/v3/getClients).

Installing from source
----------------------

To install, run

    $ go get github.com/yangxikun/sslciphersuitescheck

Build

    $ go install github.com/yangxikun/sslciphersuitescheck 

You will now find a `sslciphersuitescheck` binary in your `$GOPATH/bin` directory.

Usage
-----

Pipe your ssl ciphersuites to sslciphersuitescheck

    $ openssl ciphers -V 'ECDHE+ECDSA ECDHE AESGCM AES HIGH MEDIUM !kDH !kECDH !aNULL !eNULL !LOW !MD5 !EXP !DSS !PSK !SRP !CAMELLIA !IDEA !SEED !RC4 !3DES' | sslciphersuitescheck

Run `sslciphersuitescheck -help` for more information.

package assets
var Init_doc_security_notes = `                    IPFS Alpha Security Notes

We try hard to ensure our system is safe and robust, but all software
has bugs, especially new software. This distribution is meant to be an
alpha preview, don't use it for anything mission critical.

Please note the following:

- This is alpha software and has not been audited. It is our goal
  to conduct a proper security audit once we close in on a 1.0 release.

- ipfs is a networked program, and may have serious undiscovered
  vulnerabilities. It is written in Go, and we do not execute any
  user provided data. But please point any problems out to us in a
  github issue, or email security@ipfs.io privately.

- ipfs uses encryption for all communication, but it's NOT PROVEN SECURE
  YET!  It may be totally broken. For now, the code is included to make
  sure we benchmark our operations with encryption in mind. In the future,
  there will be an "unsafe" mode for high performance intranet apps.
  If this is a blocking feature for you, please contact us.
`

# yb_log_uploader

This utility will upload files directly to YugabyteDB Support's SendSafely dropzone. YugabyteDB customers can use this tool to upload files directly to Zendesk tickets.


## Openpgp in Golang

SendSafely dropzone uploads require the use of PGP encryption. Google has deprecated the openpgp library for Go as of March 2021:

https://github.com/golang/go/issues/44226


`yb_log_uploader` uses a fork of Go's openpgp library:

https://github.com/ProtonMail/gopenpgp

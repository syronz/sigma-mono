# SIGMAMONO
Accounting application, inspired by OMEGA and SIGMA

[![BuildStatus](https://api.travis-ci.org/syronz/sigma.svg?branch=master)](http://travis-ci.org/syronz/sigma) 
[![Go Report Card](https://goreportcard.com/badge/github.com/syronz/sigma)](https://goreportcard.com/report/github.com/syronz/sigma)
[![Go Coverage](https://github.com/syronz/sigma/blob/master/coverage_badge.png)](https://gocover.io/github.com/syronz/sigma)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/e927e927d84447a3967de50c0c155eba)](https://www.codacy.com/manual/syronz/sigma?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=syronz/sigma&amp;utm_campaign=Badge_Grade)
[![codebeat badge](https://codebeat.co/badges/5bfc77f0-d7d0-450c-bbb1-b6f8521e1630)](https://codebeat.co/projects/github-com-syronz-sigma-master)
[![Maintainability](https://api.codeclimate.com/v1/badges/1402fbdb45356d914cb1/maintainability)](https://codeclimate.com/github/syronz/sigma/maintainability)
[![GolangCI](https://golangci.com/badges/github.com/gojek/darkroom.svg)](https://golangci.com/r/github.com/syronz/sigma)

[![GoDoc](https://godoc.org/github.com/syronz/go-log?status.svg)](https://godoc.org/github.com/syronz/sigma)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fsyronz%2Fsigma.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fsyronz%2Fsigma?ref=badge_shield)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://github.com/syronz/sigma/blob/master/LICENSE)

## Coverage

For calculating coverage run below command
```bash
gopherbadger -md="README.md"
```

## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fsyronz%2Fsigma.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fsyronz%2Fsigma?ref=badge_large)

### Disable gorm log in tests
```bash
noLogSQL=true go test -v
```

### Testing Coverage
```bash
go test -atomic -coverprofile=coverage.out
go tool cover -html coverage.out
go tool cover -func coverage.out
```

## format of ID
```
ID: 1001  100  100 000 001
ID: 9999, 999, 999,999,999
    Comp  Nod  ---- ID ---

1844 6744 ,073,709,551,615
  a    b         c
a company_id
b node_id
c regular id
```

## Features
Features should be encrypted and save in the company table, controlled by license

1.  Delete List
2.  Undelete
3.  Inventory
4.  Currency number
5.  Network accessibility (192.* or public)
6.  Auto backup
7.  Number of user
8.  Income Statemnt
9.  Cash Flow Statement
10. Balance Sheet
11. Activity
12. Log API

## TODO
implement https://github.com/syronz/machineid


## Find test files for updating travis
```shell
find . -regex ".*test.go"
```

## Save coverage for each part
```shell
noLogSQL=true go test -run Test -v -cover -coverpkg "sigma/domain/sync/service" "sigma/domain/sync/test" -coverprofile coverage.out
go tool cover -html coverage.out
```


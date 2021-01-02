## antideb

```go
             _         _      _     
            | |       | |    | |    
  __ _ _ __ | |_ _  __| | ___| |__  
 / _` | '_ \| __| |/ _` |/ _ \ '_ \ 
| (_| | | | | |_| | (_| |  __/ |_) |
 \__,_|_| |_|\__|_|\__,_|\___|_.__/ 
                                    
```

Package antideb - basic anti-debugging and anti-reverse engineering protection for your application. Performs basic detection functions such as ptrace, int3, time slots, vdso and others (don't foget to obfuscate your code).
<br/><br/>
[![GoDoc](https://godoc.org/github.com/biter777/antideb?status.svg)](https://godoc.org/github.com/antideb/countries)
[![GoDev](https://img.shields.io/badge/godev-reference-5b77b3)](https://pkg.go.dev/github.com/biter777/antideb?tab=doc)
[![Go Walker](https://img.shields.io/badge/gowalker-reference-5b77b3)](https://gowalker.org/github.com/biter777/antideb)
[![GolangCI](https://golangci.com/badges/github.com/biter777/antideb.svg?style=flat)](https://golangci.com/r/github.com/biter777/antideb)
[![GoReport](https://goreportcard.com/badge/github.com/biter777/antideb)](https://goreportcard.com/report/github.com/biter777/antideb)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/08eb1d2ff62e465091b3a288ae078a96)](https://www.codacy.com/manual/biter777/antideb?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=biter777/antideb&amp;utm_campaign=Badge_Grade)
[![codecov](https://codecov.io/gh/biter777/antideb/branch/master/graph/badge.svg)](https://codecov.io/gh/biter777/antideb)
[![Coverage Status](https://coveralls.io/repos/github/biter777/antideb/badge.svg?branch=master)](https://coveralls.io/github/biter777/antideb?branch=master)
[![Coverage](https://img.shields.io/badge/coverage-gocover.io-brightgreen)](https://gocover.io/github.com/biter777/antideb)
[![License](https://img.shields.io/badge/License-BSD%202--Clause-brightgreen.svg)](https://opensource.org/licenses/BSD-2-Clause)
[![Build Status](https://travis-ci.org/biter777/antideb.svg?branch=master)](https://travis-ci.org/biter777/antideb)
[![Build status](https://ci.appveyor.com/api/projects/status/t9lpor9o8tpacpmr/branch/master?svg=true)](https://ci.appveyor.com/project/biter777/antideb/branch/master)
[![Circle CI](https://circleci.com/gh/biter777/antideb/tree/master.svg?style=shield)](https://circleci.com/gh/biter777/antideb/tree/master)
[![Semaphore Status](https://biter777.semaphoreci.com/badges/antideb.svg?style=shields)](https://biter777.semaphoreci.com/projects/antideb)
[![Build Status](https://github.com/go-vgo/robotgo/workflows/Go/badge.svg)](https://github.com/go-vgo/robotgo/commits/master)
[![Codeship Status](https://codeship.com/projects/00db4400-1803-0138-1132-7ab932dd1523/status?branch=master)](https://app.codeship.com/projects/381056) 
[![Gluten Free](https://img.shields.io/badge/gluten-free-brightgreen)](https://www.scsglobalservices.com/services/gluten-free-certification)
[![DepShield Badge](https://depshield.sonatype.org/badges/biter777/antideb/depshield.svg)](https://depshield.github.io)
<br/>

### installation

    go get github.com/biter777/antideb

### usage

```go
func main() {
	debug := false // set to false for production
	// ... do litle work
	if !debug {
		go antideb.Detect(true)
	}
	// ... do main work
}
```

### options

```go
import "github.com/biter777/antideb"
```

For more complex options, consult the [documentation](http://godoc.org/github.com/biter777/antideb).

### contributing

1) <b>Welcome pull requests, bug fixes and issue reports.</b><br/>
Before proposing a change, please discuss it first by raising an issue.<br/>

2) <b>Donate</b>. A donation isn't necessary, but it's welcome.<br/>
<noscript><a href="https://liberapay.com/biter777/donate"><img alt="Donate using Liberapay" src="https://liberapay.com/assets/widgets/donate.svg"></a></noscript> 
[![ko-fi](https://www.ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/I2I61D1XZ) <a href="https://pay.cloudtips.ru/p/94fc4268" target="_blank"><img height="30" src="https://usa.visa.com/dam/VCOM/regional/lac/ENG/Default/Partner%20With%20Us/Payment%20Technology/visapos/full-color-800x450.jpg"></a> <a href="https://pay.cloudtips.ru/p/94fc4268" target="_blank"><img height="30" src="https://brand.mastercard.com/content/dam/mccom/brandcenter/thumbnails/mastercard_debit_sym_decal_web_105px.png"></a> <a href="https://pay.cloudtips.ru/p/94fc4268" target="_blank"><img height="30" src="https://developer.apple.com/assets/elements/icons/apple-pay/apple-pay.svg"></a> <a href="https://pay.cloudtips.ru/p/94fc4268" target="_blank"><img height="30" src="https://developers.google.com/pay/api/images/brand-guidelines/google-pay-mark.png"></a> <a href="https://money.yandex.ru/to/4100164702007" target="_blank"><img width="125" height="25" src="https://yastatic.net/q/logoaas/v1/Yandex%20Money.svg"></a><br/>

3) <b>Star us</b>. Give us a star, please, if it's not against your religion :)

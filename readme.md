# BAD IP 

`badip` is an IP GeolocationInfo tool that can help security researchers in recon. 
```bash
./badip 1.1.1.1
```
``` json
IP :  1.1.1.1
City :  Los Angeles
Region :  California
Country :  US
Location :  34.0522,-118.2437
Company :  AS13335 Cloudflare, Inc.
Postal Code :  90076
TimeZone :  America/Los_Angeles
```
Requirements
------------

BadIP requires the following to run:

  * [GO][go] 1.15+
  * [ipinfo.io Token][ipinfo] 


[go]: https://go.dev/doc/install
[ipinfo]: https://ipinfo.io/
## Installation

Clone the repository [BadIp](https://github.com/iamtakdir/badip) in your local machine. Grab your free API token from here [IP Info](https://ipinfo.io/pricing), choose free package and copy the token from dashboard. 
Then replace the `token` in `config.go`file with your token. 

`config/config.go`

#### Example: 
```go
package config

func GetToken() string {
	var token = "Put Your Token Here"
	return token
}
```

### Usage `./badip 1.1.1.1`
Running using Golang
```bash
go run main.go
```
Runing as an 
executable  file 
```bash
go build
```
Installing in PATH
```bash
go install 
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## Tech Stack 
<img src="{https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white
}"/>


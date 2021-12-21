# gonetrc

This is the library parse the system netrc file, support linux/macos/windows.

We have two Function,

```
// for http Request, add credentials from netrc
func AddCredentials(req *http.Request) (added bool) 
// for host name, get matched username/password from netrc
func GetCredentials(host string) (string, string)
```
package gonetrc

import "net/http"

// AddCredentials fills in the user's credentials for req, if any.
// The return value reports whether any matching credentials were found.
func AddCredentials(req *http.Request) (added bool) {
	host := req.URL.Hostname()

	// TODO(golang.org/issue/26232): Support arbitrary user-provided credentials.
	netrcOnce.Do(readNetrc)
	for _, l := range netrc {
		if l.machine == host {
			req.SetBasicAuth(l.login, l.password)
			return true
		}
	}

	return false
}

// GetCredentials if matched the host name
func GetCredentials(host string) (string, string) {

	// TODO(golang.org/issue/26232): Support arbitrary user-provided credentials.
	netrcOnce.Do(readNetrc)
	for _, l := range netrc {
		if l.machine == host {
			return l.login, l.password
		}
	}

	return "", ""
}

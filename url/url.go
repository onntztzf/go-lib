package url

import (
	"fmt"
	"net/url"
)

//AppendParams Append parameters to the url
func AppendParams(URL string, newParams map[string]interface{}) string {
	u, err := url.Parse(URL)
	if err != nil {
		return URL
	}
	q := u.Query()
	for k, v := range newParams {
		q.Set(k, fmt.Sprintf("%v", v))
	}
	u.RawQuery = q.Encode()
	return u.String()
}

/*
Hand-written wrapper to encapsulate the details of the generated library
*/

package lotw

import (
	"context"
)

/*
Query LotW and return an ADIF string with the results.
*/
func Query(user string, pw string, opts *QueryOpts) (string, error) {
	config := NewConfiguration()
	config.UserAgent = "k0swe/lotw-qsl 0.0.1"
	client := NewAPIClient(config)
	queryResp, _, err := client.DefaultApi.Query(context.TODO(), user, pw, 1, opts)
	if err != nil {
		return "", err
	}
	return queryResp, nil
}

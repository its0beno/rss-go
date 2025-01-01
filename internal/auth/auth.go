package auth

import (
	"errors"
	"net/http"
	"strings"
)

// extract api key from the headers
// Example
// Authorization : APIKey {insert an api_key}
func GetAPiKey(headers http.Header) (string, error){
 val := headers.Get("Authorization")
 if val == ""{
	return "", errors.New("no authentication info found")
 }
 vals:= strings.Split(val, " ")
 if len(vals)!=2{
	return "", errors.New("malformed header")
 }
 if vals[0]!= "APiKey"{
	return "", errors.New("malformed first part of auth header")
 }
 return vals[1],nil
}
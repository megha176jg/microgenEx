// Code generated by microgen 1.0.4. DO NOT EDIT.

package transport

import "context"

func (set EndpointsSet) Validate(arg0 context.Context, arg1 int, arg2 string) (res0 error) {
	request := ValidateRequest{
		Token:  arg2,
		UserId: arg1,
	}
	_, res0 = set.ValidateEndpoint(arg0, &request)
	if res0 != nil {
		return
	}
	return res0
}

func (set EndpointsSet) Delete(arg0 context.Context, arg1 int, arg2 string) (res0 error) {
	request := DeleteRequest{
		Token:  arg2,
		UserId: arg1,
	}
	_, res0 = set.DeleteEndpoint(arg0, &request)
	if res0 != nil {
		return
	}
	return res0
}

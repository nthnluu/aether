package hours

type uppercaseRequest struct {
	S string `json:"s"`
}

type uppercaseResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}

//func makeUppercaseEndpoint(svc StringService) endpoint.Endpoint {
//	return func(_ context.Context, request interface{}) (interface{}, error) {
//		req := request.(uppercaseRequest)
//		v, err := svc.Uppercase(req.S)
//		if err != nil {
//			return uppercaseResponse{v, err.Error()}, nil
//		}
//		return uppercaseResponse{v, ""}, nil
//	}
//}

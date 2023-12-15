package models

type Response struct {
	Data interface{} `json:"data"`
}

type ResponseBody struct {
	ID         string      `json:"id"`
	Type       string      `json:"type"`
	Attributes interface{} `json:"attributes"`
}

func NewResponse(i Item) *Response {
	r := &Response{}
	rb := &ResponseBody{}
	r.Data = rb
	rb.Attributes = i
	rb.ID = i.GetID()
	rb.Type = i.GetType()
	return r
}

func NewListResponse(res []Item) *Response {
	r := &Response{}
	bodies := make([]*ResponseBody, len(res))
	for index, i := range res {
		rb := &ResponseBody{}
		r.Data = rb
		rb.Attributes = i
		rb.ID = i.GetID()
		rb.Type = i.GetType()
		bodies[index] = rb
	}
	r.Data = bodies
	return r
}

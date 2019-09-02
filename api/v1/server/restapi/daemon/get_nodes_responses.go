// Code generated by go-swagger; DO NOT EDIT.

package daemon

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/cilium/cilium/api/v1/models"
)

// GetNodesOKCode is the HTTP code returned for type GetNodesOK
const GetNodesOKCode int = 200

/*GetNodesOK Success

swagger:response getNodesOK
*/
type GetNodesOK struct {

	/*
	  In: Body
	*/
	Payload *models.StatusResponse `json:"body,omitempty"`
}

// NewGetNodesOK creates GetNodesOK with default headers values
func NewGetNodesOK() *GetNodesOK {

	return &GetNodesOK{}
}

// WithPayload adds the payload to the get nodes o k response
func (o *GetNodesOK) WithPayload(payload *models.StatusResponse) *GetNodesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get nodes o k response
func (o *GetNodesOK) SetPayload(payload *models.StatusResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNodesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

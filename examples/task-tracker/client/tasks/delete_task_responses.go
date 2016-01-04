package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"
)

type DeleteTaskReader struct {
	formats strfmt.Registry
}

func (o *DeleteTaskReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteTaskNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteTaskDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewDeleteTaskNoContent creates a DeleteTaskNoContent with default headers values
func NewDeleteTaskNoContent() *DeleteTaskNoContent {
	return &DeleteTaskNoContent{}
}

/*DeleteTaskNoContent

Task deleted
*/
type DeleteTaskNoContent struct {
}

func (o *DeleteTaskNoContent) Error() string {
	return fmt.Sprintf("[DELETE /tasks/{id}][%d] deleteTaskNoContent ", 204)
}

func (o *DeleteTaskNoContent) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTaskDefault creates a DeleteTaskDefault with default headers values
func NewDeleteTaskDefault(code int) *DeleteTaskDefault {
	return &DeleteTaskDefault{
		_statusCode: code,
	}
}

/*DeleteTaskDefault

DeleteTaskDefault delete task default
*/
type DeleteTaskDefault struct {
	_statusCode int
}

// Code gets the status code for the delete task default response
func (o *DeleteTaskDefault) Code() int {
	return o._statusCode
}

func (o *DeleteTaskDefault) Error() string {
	return fmt.Sprintf("[DELETE /tasks/{id}][%d] deleteTask default ", o._statusCode)
}

func (o *DeleteTaskDefault) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"
)

/*Order order

swagger:model Order
*/
type Order struct {

	/* Complete complete
	 */
	Complete *bool `json:"complete,omitempty"`

	/* ID id
	 */
	ID int64 `json:"id,omitempty"`

	/* PetID pet id
	 */
	PetID int64 `json:"petId,omitempty"`

	/* Quantity quantity
	 */
	Quantity int32 `json:"quantity,omitempty"`

	/* ShipDate ship date
	 */
	ShipDate strfmt.DateTime `json:"shipDate,omitempty"`

	/* Order Status
	 */
	Status *string `json:"status,omitempty"`
}

// Validate validates this order
func (m *Order) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var orderStatusEnum []interface{}

func (m *Order) validateStatusEnum(path, location string, value string) error {
	if orderStatusEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["placed","approved","delivered"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			orderStatusEnum = append(orderStatusEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, orderStatusEnum); err != nil {
		return err
	}
	return nil
}

func (m *Order) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

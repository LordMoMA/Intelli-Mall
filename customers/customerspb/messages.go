package customerspb

import (
	"github.com/LordMoMA/Intelli-Mall/internal/registry"
	"github.com/LordMoMA/Intelli-Mall/internal/registry/serdes"
)

const (
	CustomerAggregateChannel = "intellimall.customers.events.Customer"

	CustomerRegisteredEvent = "customersapi.CustomerRegistered"
	CustomerSmsChangedEvent = "customersapi.CustomerSmsChanged"
	CustomerEnabledEvent    = "customersapi.CustomerEnabled"
	CustomerDisabledEvent   = "customersapi.CustomerDisabled"

	CommandChannel = "intellimall.customers.commands"

	AuthorizeCustomerCommand = "customersapi.AuthorizeCustomer"
)

func Registrations(reg registry.Registry) error {
	serde := serdes.NewProtoSerde(reg)

	// Customer events
	if err := serde.Register(&CustomerRegistered{}); err != nil {
		return err
	}
	if err := serde.Register(&CustomerSmsChanged{}); err != nil {
		return err
	}
	if err := serde.Register(&CustomerEnabled{}); err != nil {
		return err
	}
	if err := serde.Register(&CustomerDisabled{}); err != nil {
		return err
	}

	// commands
	if err := serde.Register(&AuthorizeCustomer{}); err != nil {
		return err
	}
	return nil
}

func (*CustomerRegistered) Key() string { return CustomerRegisteredEvent }
func (*CustomerSmsChanged) Key() string { return CustomerSmsChangedEvent }
func (*CustomerEnabled) Key() string    { return CustomerEnabledEvent }
func (*CustomerDisabled) Key() string   { return CustomerDisabledEvent }

func (*AuthorizeCustomer) Key() string { return AuthorizeCustomerCommand }

package internal

import (
	"context"

	"github.com/LordMoMA/Intelli-Mall/cosec/internal/models"
	"github.com/LordMoMA/Intelli-Mall/customers/customerspb"
	"github.com/LordMoMA/Intelli-Mall/depot/depotpb"
	"github.com/LordMoMA/Intelli-Mall/internal/ddd"
	"github.com/LordMoMA/Intelli-Mall/internal/sec"
	"github.com/LordMoMA/Intelli-Mall/ordering/orderingpb"
	"github.com/LordMoMA/Intelli-Mall/payments/paymentspb"
)

const CreateOrderSagaName = "cosec.CreateOrder"
const CreateOrderReplyChannel = "intellimall.cosec.replies.CreateOrder"

type createOrderSaga struct {
	sec.Saga[*models.CreateOrderData]
}

func NewCreateOrderSaga() sec.Saga[*models.CreateOrderData] {
	saga := createOrderSaga{
		Saga: sec.NewSaga[*models.CreateOrderData](CreateOrderSagaName, CreateOrderReplyChannel),
	}

	// 0. -RejectOrder
	saga.AddStep().
		Compensation(saga.rejectOrder)

	// 1. AuthorizeCustomer
	saga.AddStep().
		Action(saga.authorizeCustomer)

	// 2. CreateShoppingList, -CancelShoppingList
	saga.AddStep().
		Action(saga.createShoppingList).
		OnActionReply(depotpb.CreatedShoppingListReply, saga.onCreatedShoppingListReply).
		Compensation(saga.cancelShoppingList)

	// 3. ConfirmPayment
	saga.AddStep().
		Action(saga.confirmPayment)

	// 4. InitiateShopping
	saga.AddStep().
		Action(saga.initiateShopping)

	// 5. ApproveOrder
	saga.AddStep().
		Action(saga.approveOrder)

	return saga
}

func (s createOrderSaga) rejectOrder(ctx context.Context, data *models.CreateOrderData) (string, ddd.Command, error) {
	return orderingpb.CommandChannel, ddd.NewCommand(orderingpb.RejectOrderCommand, &orderingpb.RejectOrder{Id: data.OrderID}), nil
}

func (s createOrderSaga) authorizeCustomer(ctx context.Context, data *models.CreateOrderData) (string, ddd.Command, error) {
	return customerspb.CommandChannel, ddd.NewCommand(customerspb.AuthorizeCustomerCommand, &customerspb.AuthorizeCustomer{Id: data.CustomerID}), nil
}

func (s createOrderSaga) createShoppingList(ctx context.Context, data *models.CreateOrderData) (string, ddd.Command, error) {
	items := make([]*depotpb.CreateShoppingList_Item, len(data.Items))
	for i, item := range data.Items {
		items[i] = &depotpb.CreateShoppingList_Item{
			ProductId: item.ProductID,
			StoreId:   item.StoreID,
			Quantity:  int32(item.Quantity),
		}
	}

	return depotpb.CommandChannel, ddd.NewCommand(depotpb.CreateShoppingListCommand, &depotpb.CreateShoppingList{
		OrderId: data.OrderID,
		Items:   items,
	}), nil
}

func (s createOrderSaga) onCreatedShoppingListReply(ctx context.Context, data *models.CreateOrderData, reply ddd.Reply) error {
	payload := reply.Payload().(*depotpb.CreatedShoppingList)

	data.ShoppingID = payload.GetId()

	return nil
}

func (s createOrderSaga) cancelShoppingList(ctx context.Context, data *models.CreateOrderData) (string, ddd.Command, error) {
	return depotpb.CommandChannel, ddd.NewCommand(depotpb.CancelShoppingListCommand, &depotpb.CancelShoppingList{Id: data.ShoppingID}), nil
}

func (s createOrderSaga) confirmPayment(ctx context.Context, data *models.CreateOrderData) (string, ddd.Command, error) {
	return paymentspb.CommandChannel, ddd.NewCommand(paymentspb.ConfirmPaymentCommand, &paymentspb.ConfirmPayment{
		Id:     data.PaymentID,
		Amount: data.Total,
	}), nil
}

func (s createOrderSaga) initiateShopping(ctx context.Context, data *models.CreateOrderData) (string, ddd.Command, error) {
	return depotpb.CommandChannel, ddd.NewCommand(depotpb.InitiateShoppingCommand, &depotpb.InitiateShopping{Id: data.ShoppingID}), nil
}

func (s createOrderSaga) approveOrder(ctx context.Context, data *models.CreateOrderData) (string, ddd.Command, error) {
	return orderingpb.CommandChannel, ddd.NewCommand(orderingpb.ApproveOrderCommand, &orderingpb.ApproveOrder{
		Id:         data.OrderID,
		ShoppingId: data.ShoppingID,
	}), nil
}

package domain

import (
	"context"
	"log"

	"github.com/AlfredoPastor/ddd-go/shared/eventbus"
	"github.com/AlfredoPastor/ddd-go/shared/vo"
)

func BookProducts(ctx context.Context, repo OrderRepository, bus eventbus.Bus, orderLines []OrderLine) error {
	idsProductsBooked := []vo.ID{}
	for _, orderLine := range orderLines {
		idBooked, err := repo.BookProductFromInventory(ctx, orderLine.ProductID, orderLine.Quantity)
		if err != nil {
			events := []eventbus.Event{}
			for _, idBookedFail := range idsProductsBooked {
				events = append(events, NewProductBookedFailEvent(idBookedFail))
			}
			errP := bus.Publish(ctx, events)
			if errP != nil {
				log.Println(errP.Error())
			}
			return err
		}
		idsProductsBooked = append(idsProductsBooked, idBooked)
	}
	idsProductsBooked = idsProductsBooked[:0]
	
	return nil
}

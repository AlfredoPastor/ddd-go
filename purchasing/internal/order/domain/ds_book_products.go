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
			for _, idBookedFail := range idsProductsBooked {
				errP := bus.Publish(ctx, []eventbus.Event{NewProductBookedFailEvent(idBookedFail)})
				if errP != nil {
					log.Println(errP.Error())
				}
			}
			return err
		}
		idsProductsBooked = append(idsProductsBooked, idBooked)
	}

	return nil
}

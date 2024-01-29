package pickup

import (
	"encoding/json"
	dto "for_learning/src/app/dto/pickup"
	natsPublisher "for_learning/src/infra/broker/nats/publisher"
	Const "for_learning/src/infra/constants"
	"log"
)

type PickUpUCInterface interface {
	Create(req *dto.ReqPickupDTO) error
}

type pickUpUseCase struct {
	Publisher natsPublisher.PublisherInterface
}

func NewPickUpUseCase(publiser natsPublisher.PublisherInterface) *pickUpUseCase {
	return &pickUpUseCase{
		Publisher: publiser,
	}
}

func (uc *pickUpUseCase) Create(req *dto.ReqPickupDTO) error {
	newData, _ := json.Marshal(req)
	err := uc.Publisher.Nats(newData, Const.BOOKS)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

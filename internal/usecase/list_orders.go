package usecase

import "github.com/luizffdemoraes/20-CleanArch/internal/entity"

type ListOrdersOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(orderRepository entity.OrderRepositoryInterface) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: orderRepository,
	}
}

func (u *ListOrdersUseCase) Execute() ([]ListOrdersOutputDTO, error) {
	orders, err := u.OrderRepository.FindAll()
	if err != nil {
		return nil, err
	}

	output := make([]ListOrdersOutputDTO, 0, len(orders))
	for _, order := range orders {
		output = append(output, ListOrdersOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}

	return output, nil
}

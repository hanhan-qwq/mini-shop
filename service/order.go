package service

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"mini_shop/global"
	"mini_shop/model"
	"mini_shop/repository"
	"mini_shop/web/request"
)

type OrderService struct {
	OrderDAO   *repository.OrderDAO
	ProductDAO *repository.ProductDAO
	DB         *gorm.DB
}

func NewOrderService() *OrderService {
	return &OrderService{
		OrderDAO:   repository.NewOrderDAO(),
		ProductDAO: repository.NewProductDAO(),
		DB:         global.GetDB(),
	}
}

type GetOrderDetailResponse struct {
	Order model.Order       `json:"order"`
	Items []model.OrderItem `json:"items"`
}

func (s *OrderService) CreateOrder(userID uint, remark string, itemReqs []request.CreateOrderItemRequest) (*model.Order, error) {
	if len(itemReqs) == 0 {
		return nil, errors.New("订单项不能为空")
	}

	var order *model.Order

	err := s.DB.Transaction(func(tx *gorm.DB) error {
		var items []model.OrderItem
		var total float64

		for _, req := range itemReqs {
			// 1 校验商品是否存在
			product, err := s.ProductDAO.GetProductByIDWithLock(tx, req.ProductID)
			if err != nil {
				return fmt.Errorf("商品不存在: %d", req.ProductID)
			}

			//  2 校验库存是否足够
			if product.Stock < req.Quantity {
				return fmt.Errorf("库存不足: %s (剩余 %d)", product.Name, product.Stock)
			}

			// 3 扣减库存
			if err := s.ProductDAO.DecreaseStockInTx(tx, product.ID, req.Quantity); err != nil {
				return errors.New("扣减库存失败")
			}

			// 4 创建订单项
			item := model.OrderItem{
				ProductID:   product.ID,
				ProductName: product.Name,
				Price:       product.Price,
				Quantity:    req.Quantity,
				TotalPrice:  product.Price * float64(req.Quantity),
			}

			items = append(items, item)
			total += item.TotalPrice
		}

		//   5 创建订单
		orderNo := fmt.Sprintf("OD%s%d", uuid.New().String()[:8], userID)
		order = &model.Order{
			OrderNo:    orderNo,
			UserID:     userID,
			TotalPrice: total,
			Status:     model.OrderStatusPendingPay,
			Remark:     remark,
		}

		if err := s.OrderDAO.CreateOrderInTx(tx, order, items); err != nil {
			return err
		}

		return nil
	})

	return order, err
}

func (s *OrderService) GetOrderDetail(userID, orderID uint) (*GetOrderDetailResponse, error) {
	order, err := s.OrderDAO.GetOrderByID(userID, orderID)
	if err != nil {
		return nil, errors.New("订单不存在或无权访问")
	}

	items, err := s.OrderDAO.GetOrderItems(order.ID)
	if err != nil {
		return nil, err
	}

	return &GetOrderDetailResponse{
		Order: *order,
		Items: items,
	}, nil
}

func (s *OrderService) ListUserOrders(userID uint, page, pageSize, status int) ([]model.Order, int64, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	// 调用 DAO 查询
	orders, count, err := s.OrderDAO.FindUserOrders(userID, page, pageSize, status)
	if err != nil {
		return nil, 0, err
	}

	return orders, count, nil
}

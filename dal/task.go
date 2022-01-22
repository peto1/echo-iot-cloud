package dal

import (
	"context"
	"time"

	"github.com/airren/echo-iot-backend/model"
	"github.com/airren/echo-iot-backend/model/req"
	"github.com/airren/echo-iot-backend/utils"
)

func CreateOrder(ctx context.Context, task *model.Order) (*model.Order, error) {
	org := utils.GetOrgFromCtx(ctx)
	task.Org = org
	err := db.Create(task).Error
	return task, err
}

func UpdateOrder(ctx context.Context, task *model.Order) (*model.Order, error) {
	org := utils.GetOrgFromCtx(ctx)
	task.Org = org
	err := db.Save(task).Error
	return task, err
}

func QueryOrders(ctx context.Context, req *req.OrderReq) ([]*model.Order, error) {
	org := utils.GetOrgFromCtx(ctx)
	var tasks []*model.Order
	query := db.Where(map[string]interface{}{"period_id": req.PeriodId, "org": org})
	err := query.Find(&tasks).Error
	return tasks, err
}

func GetOrderById(ctx context.Context, id int64) (*model.Order, error) {
	org := utils.GetOrgFromCtx(ctx)
	task := &model.Order{}
	query := db.Where("id=? and org=?", id, org)
	err := query.Find(task).Error
	return task, err

}

func QueryOrderByPriority(ctx context.Context, priority string) ([]*model.Order, error) {
	org := utils.GetOrgFromCtx(ctx)
	tasks := make([]*model.Order, 0)
	query := db.Where("priority = ?and org=?", priority, org)
	err := query.Find(&tasks).Error
	return tasks, err
}

func QueryOrderByTime(ctx context.Context, startAt time.Time, endAt time.Time) ([]*model.Order, error) {
	org := utils.GetOrgFromCtx(ctx)
	tasks := make([]*model.Order, 0)
	query := db.Where("created_at >? and created_at <=? and org=?", startAt, endAt, org)
	err := query.Find(&tasks).Error
	return tasks, err
}

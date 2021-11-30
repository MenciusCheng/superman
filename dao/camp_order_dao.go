package dao

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/MenciusCheng/superman/conf"
	"github.com/MenciusCheng/superman/util/dragons/proxy"
	"github.com/MenciusCheng/superman/util/gendbinfo"
)

type CampOrderDao struct {
	c  *conf.Config
	db *proxy.SQL
}

func NewCampOrderDao(c *conf.Config) *CampOrderDao {
	return &CampOrderDao{
		c:  c,
		db: proxy.InitSQL("camp_order"),
	}
}

// Ping check db resource status
func (d *CampOrderDao) Ping(ctx context.Context) error {
	return nil
}

// Close release resource
func (d *CampOrderDao) Close() error {
	return nil
}

func (d *CampOrderDao) GetOrder(ctx context.Context) error {
	var OrderFormData struct {
		Orderid     string `json:"orderid" gorm:"orderid"`           //订单id
		CommodityId int64  `json:"commodity_id" gorm:"commodity_id"` //商品id
	}
	err := d.db.Master(ctx).Table("order_form").Take(&OrderFormData).Error
	fmt.Printf("OrderFormData: %+v", OrderFormData)
	return err
}

func (d *CampOrderDao) GenModel(ctx context.Context) error {
	gendbinfo.Gen(d)
	return nil
}

func (d *CampOrderDao) RawRows(sql string, values ...interface{}) (*sql.Rows, error) {
	return d.db.Master(context.Background()).Raw(sql, values).Rows()
}
func (d *CampOrderDao) RawScan(dest interface{}, sql string, values ...interface{}) error {
	db := d.db.Master(context.Background()).Raw(sql, values).Scan(dest)
	return db.Error
}

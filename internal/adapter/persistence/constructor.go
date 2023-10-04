/*
Package persistence

DB로의 접근을 담당하는 persistence 객체입니다.
*/
package persistence

import (
	"github.com/Goboolean/shared/pkg/rdbms"
)

type Adapter struct {
	rdbms *rdbms.Queries
}

func New(rdbms *rdbms.Queries) *Adapter {
	return &Adapter{
		rdbms: rdbms,
	}
}

//func (a *Adapter) IsExistStockById(id string) (bool, error) {
//	return rdbms.CheckStockExist
//}

/*
func (a *Adapter) StoreStock(tx port.Transactioner, stockId string, agg *entity.StockAggregate) error {
	dto := &mongo.StockAggregate{
		EventType: agg.EventType,
		Avg:       agg.Average,
		Min:       agg.Min,
		Max:       agg.Max,
		Start:     agg.Start,
		End:       agg.End,
		StartTime: agg.StartTime,
		EndTime:   agg.EndTime,
	}

	if err := a.mongo.InsertStockBatch(tx.(*mongo.Transaction), stockId, []*mongo.StockAggregate{dto}); err != nil {
		return err
	}

	a.prom.StoreCounter()().Inc()
	return nil
}

func (a *Adapter) StoreStockBatch(tx port.Transactioner, stockId string, aggs []*entity.StockAggregate) error {
	dtos := make([]*mongo.StockAggregate, 0, len(aggs))

	for _, agg := range aggs {
		dtos = append(dtos, &mongo.StockAggregate{
			EventType: agg.EventType,
			Avg:       agg.Average,
			Min:       agg.Min,
			Max:       agg.Max,
			Start:     agg.Start,
			End:       agg.End,
			StartTime: agg.StartTime,
			EndTime:   agg.EndTime,
		})
	}

	if err := a.mongo.InsertStockBatch(tx.(*mongo.Transaction), stockId, dtos); err != nil {
		return err
	}

	a.prom.StoreCounter()().Add(float64(len(aggs)))
	return nil
}

func (a *Adapter) CreateStoringStartedLog(ctx context.Context, stockId string) error {

	return a.rdbms.CreateAccessInfo(ctx, rdbms.CreateAccessInfoParams{
		ProductID: stockId,
		Status:    "started",
	})
}

func (a *Adapter) CreateStoringStoppedLog(ctx context.Context, stockId string) error {

	return a.rdbms.CreateAccessInfo(ctx, rdbms.CreateAccessInfoParams{
		ProductID: stockId,
		Status:    "stopped",
	})
}

func (a *Adapter) CreateStoringFailedLog(ctx context.Context, stockId string) error {

	return a.rdbms.CreateAccessInfo(ctx, rdbms.CreateAccessInfoParams{
		ProductID: stockId,
		Status:    "failed",
	})
}
*/

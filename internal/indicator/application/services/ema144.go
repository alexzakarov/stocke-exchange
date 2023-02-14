package services

import (
	"context"
	ent "main/internal/indicator/domain/entities"
	"main/pkg/market_data/binance"
	fx "main/pkg/utils/formulas"

	ti "github.com/cinar/indicator"
)

// IndicatorEma144 Indicator Calculate Func
func (u *service) IndicatorEma144(ctx context.Context, chartData []binance.ChartData) (responser ent.IndicatorCalcResponse, err error) {
	if u.cfg.Server.APP_DEBUG == true {
		println("IndicatorEma144 begin to work")
	}

	var calculatedData []float64
	var closeLine []float64
	var signal int
	var result []float64
	period := 144

	for _, values := range chartData {
		closeLine = append(closeLine, values.ClosePrice)
	}

	calculatedData = ti.Ema(period, closeLine)
	for i := 0; i < len(calculatedData); i++ {
		if closeLine[i] > calculatedData[i] {
			result = append(result, 1)
		} else if closeLine[i] < calculatedData[i] {
			result = append(result, -1)
		} else {
			result = append(result, 0)
		}
	}

	signal = int(fx.GetByIndexN(result, 2))

	responser = ent.IndicatorCalcResponse{
		Signal:       signal,
		Result:       nil,
		CalculatedAt: chartData[len(chartData)-1].CloseTime,
	}

	return
}

package dummy_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kodefluence/altair/mock"
	"github.com/kodefluence/altair/plugin/metric/module/dummy"
	"github.com/kodefluence/altair/plugin/metric/module/dummy/controller/metric"
)

func TestProvider(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	appBearer := mock.NewMockAppBearer(mockController)
	appBearer.EXPECT().SetMetricProvider(metric.NewDummy())
	dummy.Load(appBearer)
}

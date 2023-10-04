package simulation_test

import (
	"context"
	"github.com/Goboolean/core-system.command/cmd/inject"
	"github.com/Goboolean/core-system.command/internal/adapter/broker"
	"github.com/Goboolean/core-system.command/internal/adapter/persistence"
	"github.com/Goboolean/core-system.command/internal/domain/entity"
	"github.com/Goboolean/core-system.command/internal/domain/service/simulation"
	"github.com/Goboolean/core-system.command/internal/domain/service/task"
	"github.com/Goboolean/core-system.command/internal/util/hash"
	"github.com/Goboolean/core-system.command/internal/util/log"
	"github.com/Goboolean/core-system.command/internal/util/msg"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

var (
	tm    *task.Manager
	kafka *broker.MockAdapter
	db    *persistence.MockAdapter
	sm    *simulation.Manager
)

func SetUp() {
	kafka = broker.NewMockAdapter(false)
	tm, _ = inject.InitializeMockTaskManager(nil, kafka)
	db = persistence.NewMockAdapter()
	sm, _ = inject.InitializeMockSimulationManager(nil, tm, db)
	log.Init()
}

func TearDown() {
	// nothing
}

func TestMain(m *testing.M) {
	SetUp()
	code := m.Run()
	TearDown()
	os.Exit(code)
}

func testCleanUp() {
	kafka.SetError(false)
	tm.RemoveTaskAll()
	db.ClearModel()
	db.ClearStock()
	db.ClearSimulation()
}

func Test_Simulation_AllocateTask(t *testing.T) {

	type args struct {
		modelId string
		stockId string
		client  *entity.Client
	}

	testModelId := "test-model-id"
	testStockId := "test-stock-id"
	testUUID := hash.Sha256(testModelId + testStockId)

	// task			X
	// simulation	X
	// model		O
	// stock		O
	t.Run("Allocate task (case: model, stock exists)", func(t *testing.T) {
		defer testCleanUp()

		// given
		db.AddModel(testModelId)
		db.AddStock(testStockId)

		ch := make(chan *entity.TaskAllocResponse)
		defer close(ch)
		c := &entity.Config{
			Channel:      true,
			Notification: false,
			AutoTrading:  false,
		}
		arg := args{
			modelId: testModelId,
			stockId: testStockId,
			client:  entity.NewClient(ch, c),
		}

		assertions := assert.New(t)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// when
		_, err := sm.Allocate(ctx, arg.modelId, arg.stockId, arg.client)

		// then
		assertions.NoError(err, "task allocate error.")
		assertions.Equal(true, tm.IsExistTask(testUUID), "task is not exist.")
	})

	// task			X
	// simulation	O
	// model		O
	// stock		O
	t.Run("Allocate task (case: simulation, model, stock exist)", func(t *testing.T) {
		defer testCleanUp()

		// given
		db.AddSimulation(testUUID)
		db.AddModel(testModelId)
		db.AddStock(testStockId)

		ch := make(chan *entity.TaskAllocResponse)
		defer close(ch)
		c := &entity.Config{
			Channel:      true,
			Notification: false,
			AutoTrading:  false,
		}
		arg := args{
			modelId: testModelId,
			stockId: testStockId,
			client:  entity.NewClient(ch, c),
		}

		assertions := assert.New(t)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// when
		_, err := sm.Allocate(ctx, arg.modelId, arg.stockId, arg.client)

		// then
		assertions.NoError(err, "task allocate error.")

		// task는 존재해서는 안 됨.
		assertions.Equal(false, tm.IsExistTask(testUUID), "task exist.")

		// simulation이 이미 존재하기에 channel에 메시지가 전송되어야 함.
		select {
		case m := <-ch:
			assertions.Equal(msg.Success(testUUID), m, "message is not equal.")
		case <-time.After(2 * time.Second):
			t.Error("timed out waiting for data on the channel.")
		}
	})

	// task			O
	// simulation	X
	// model		O
	// stock		O
	t.Run("Allocate task (case: task, model, stock exist)", func(t *testing.T) {
		defer testCleanUp()

		// given
		db.AddModel(testModelId)
		db.AddStock(testStockId)

		ch := make(chan *entity.TaskAllocResponse)
		defer close(ch)
		c := &entity.Config{
			Channel:      true,
			Notification: false,
			AutoTrading:  false,
		}
		arg := args{
			modelId: testModelId,
			stockId: testStockId,
			client:  entity.NewClient(ch, c),
		}

		assertions := assert.New(t)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		_, err := sm.Allocate(ctx, arg.modelId, arg.stockId, arg.client)
		assertions.NoError(err, "task allocate error.")
		assertions.Equal(true, tm.IsExistTask(testUUID), "task is not exist.")

		// new request
		ch2 := make(chan *entity.TaskAllocResponse)
		defer close(ch2)
		c2 := &entity.Config{
			Channel:      true,
			Notification: false,
			AutoTrading:  false,
		}
		arg2 := args{
			modelId: testModelId,
			stockId: testStockId,
			client:  entity.NewClient(ch2, c2),
		}

		// when
		_, err2 := sm.Allocate(ctx, arg2.modelId, arg2.stockId, arg2.client)

		// then
		assertions.NoError(err2, "task allocate error.")

		// 이미 존재하는 task에 구독되어야 함.
		expectedData := &entity.TaskAllocResponse{
			Status:      123,
			AccessToken: "test-access-token",
		}

		tm.Notify(testUUID, expectedData.Status, expectedData.AccessToken, true)

		select {
		case m := <-ch:
			assertions.Equal(expectedData, m, "message is not equal.")
		case <-time.After(2 * time.Second):
			t.Error("timed out waiting for data on the channel.")
		}

		select {
		case m := <-ch2:
			assertions.Equal(expectedData, m, "message is not equal.")
		case <-time.After(2 * time.Second):
			t.Error("timed out waiting for data on the channel.")
		}
	})

	// task			X
	// simulation	X
	// model		X
	// stock		O
	t.Run("Allocate task (case: stock exist)", func(t *testing.T) {
		defer testCleanUp()

		// given
		db.AddStock(testStockId)

		ch := make(chan *entity.TaskAllocResponse)
		defer close(ch)
		c := &entity.Config{
			Channel:      true,
			Notification: false,
			AutoTrading:  false,
		}
		arg := args{
			modelId: testModelId,
			stockId: testStockId,
			client:  entity.NewClient(ch, c),
		}

		assertions := assert.New(t)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// when
		_, err := sm.Allocate(ctx, arg.modelId, arg.stockId, arg.client)

		// then
		assertions.NoError(err, "task allocate error.")
		assertions.Equal(false, tm.IsExistTask(testUUID), "task exist.")

		select {
		case m := <-ch:
			assertions.Equal(msg.NotFound(), m, "message is not equal.")
		case <-time.After(2 * time.Second):
			t.Error("timed out waiting for data on the channel.")
		}
	})

	// task			X
	// simulation	X
	// model		O
	// stock		X
	t.Run("Allocate task (case: model exist)", func(t *testing.T) {
		defer testCleanUp()

		// given
		db.AddModel(testModelId)

		ch := make(chan *entity.TaskAllocResponse)
		defer close(ch)
		c := &entity.Config{
			Channel:      true,
			Notification: false,
			AutoTrading:  false,
		}
		arg := args{
			modelId: testModelId,
			stockId: testStockId,
			client:  entity.NewClient(ch, c),
		}

		assertions := assert.New(t)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// when
		_, err := sm.Allocate(ctx, arg.modelId, arg.stockId, arg.client)

		// then
		assertions.NoError(err, "task allocate error.")
		assertions.Equal(false, tm.IsExistTask(testUUID), "task exist.")

		select {
		case m := <-ch:
			assertions.Equal(msg.NotFound(), m, "message is not equal.")
		case <-time.After(2 * time.Second):
			t.Error("timed out waiting for data on the channel.")
		}
	})

	// task			X
	// simulation	X
	// model		X
	// stock		X
	t.Run("Allocate task (case: nothing exist)", func(t *testing.T) {
		defer testCleanUp()

		// given
		ch := make(chan *entity.TaskAllocResponse)
		defer close(ch)
		c := &entity.Config{
			Channel:      true,
			Notification: false,
			AutoTrading:  false,
		}
		arg := args{
			modelId: testModelId,
			stockId: testStockId,
			client:  entity.NewClient(ch, c),
		}

		assertions := assert.New(t)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// when
		_, err := sm.Allocate(ctx, arg.modelId, arg.stockId, arg.client)

		// then
		assertions.NoError(err, "task allocate error.")
		assertions.Equal(false, tm.IsExistTask(testUUID), "task exist.")

		select {
		case m := <-ch:
			assertions.Equal(msg.NotFound(), m, "message is not equal.")
		case <-time.After(2 * time.Second):
			t.Error("timed out waiting for data on the channel.")
		}
	})
}

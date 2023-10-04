package task_test

import (
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/Goboolean/core-system.command/cmd/inject"
	"github.com/Goboolean/core-system.command/internal/adapter/broker"
	"github.com/Goboolean/core-system.command/internal/domain/entity"
	"github.com/Goboolean/core-system.command/internal/domain/service/task"
	"github.com/Goboolean/core-system.command/internal/util/log"
	"github.com/stretchr/testify/assert"
	"os"
	"sync"
	"testing"
	"time"
)

var (
	manager *task.Manager
	kafka   *broker.MockAdapter
)

func SetUp() {
	kafka = broker.NewMockAdapter(false)
	manager, _ = inject.InitializeMockTaskManager(nil, kafka)
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

func Test_Task_AllocateTask(t *testing.T) {

	type args struct {
		modelId string
		stockId string
		client  *entity.Client
	}

	// 최초 테스크 추가 + 메시지 테스트
	t.Run("Allocate task (case: send message)", func(t *testing.T) {
		// given
		ch := make(chan *entity.TaskAllocResponse)
		defer close(ch)
		c := &entity.Config{
			Channel:      true,
			Notification: false,
			AutoTrading:  false,
		}
		arg := args{
			modelId: "test-model-id",
			stockId: "test-stock-id",
			client:  entity.NewClient(ch, c),
		}
		uuid := fmt.Sprintf("%x", sha256.Sum256([]byte(arg.modelId+arg.stockId)))

		// when
		assertions := assert.New(t)
		assertions.Equal(false, manager.IsExistTask(uuid), "task is exist.")

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// then
		clientKey, err := manager.Allocate(ctx, arg.modelId, arg.stockId, arg.client)
		success := assertions.NoError(err, "task allocate error.")
		if success {
			defer manager.RemoveTask(uuid)
			defer manager.RemoveSubscriber(uuid, clientKey)
		}
		assertions.Equal(true, manager.IsExistTask(uuid), "task is not exist.")

		// client가 해당 task에 잘 추가되었는지 확인 -> 메시지 전달
		// given
		expectedData := &entity.TaskAllocResponse{
			Status:      123,
			AccessToken: "test-access-token",
		}

		// when
		go manager.Notify(uuid, expectedData.Status, expectedData.AccessToken, true)

		// then
		select {
		case receivedData := <-ch:
			assertions.Equal(expectedData, receivedData, "received data is not equal to expected data.")
		case <-time.After(2 * time.Second):
			t.Error("timed out waiting for data on the channel.")
		}
	})
	assert.Equal(t, 0, manager.GetSizeOfTask())

	// 최초 테스크 추가 실패
	t.Run("Allocate task (case: kafka infra error)", func(t *testing.T) {

		// given
		kafka.SetError(true)
		defer func() {
			kafka.SetError(false)
		}()

		ch := make(chan *entity.TaskAllocResponse)
		defer close(ch)
		c := &entity.Config{
			Channel:      true,
			Notification: false,
			AutoTrading:  false,
		}
		arg := args{
			modelId: "test-model-id-2",
			stockId: "test-stock-id-2",
			client:  entity.NewClient(ch, c),
		}
		uuid := fmt.Sprintf("%x", sha256.Sum256([]byte(arg.modelId+arg.stockId)))

		// when
		assertions := assert.New(t)
		assertions.Equal(false, manager.IsExistTask(uuid), "task is exist.")

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		clientKey, err := manager.Allocate(ctx, arg.modelId, arg.stockId, arg.client)

		// then
		success := assertions.Error(err, "kafka에 문제가 있지만, task 등록에 성공하는 문제가 발생했습니다.")
		if !success {
			defer manager.RemoveTask(uuid)
			defer manager.RemoveSubscriber(uuid, clientKey)
		}

		select {
		case msg := <-ch:
			assertions.Equal(&entity.TaskAllocResponse{
				Status:      500,
				AccessToken: "",
			}, msg, "received data is not equal to expected data.")
		case <-time.After(2 * time.Second):
			t.Error("timed out waiting for data on the channel.")
		}
	})
	assert.Equal(t, 0, manager.GetSizeOfTask())

	// 이미 존재하는 테스크에 추가
	t.Run("Allocate task (case: already exist task)", func(t *testing.T) {

		// given
		ch := make(chan *entity.TaskAllocResponse)
		defer close(ch)
		c := &entity.Config{
			Channel:      true,
			Notification: false,
			AutoTrading:  false,
		}
		arg := args{
			modelId: "test-model-id-3",
			stockId: "test-stock-id-3",
			client:  entity.NewClient(ch, c),
		}
		uuid := fmt.Sprintf("%x", sha256.Sum256([]byte(arg.modelId+arg.stockId)))

		// when
		assertions := assert.New(t)
		assertions.Equal(false, manager.IsExistTask(uuid), "task is exist.")

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		clientKey, err := manager.Allocate(ctx, arg.modelId, arg.stockId, arg.client)

		// then
		isPass := assertions.NoError(err, "task allocate error.")
		if isPass {
			defer manager.RemoveTask(uuid)
			defer manager.RemoveSubscriber(uuid, clientKey)
		}
		assertions.Equal(true, manager.IsExistTask(uuid), "task is not exist.")

		// 이미 존재하는 테스크에 추가
		clientKey, err = manager.Allocate(ctx, arg.modelId, arg.stockId, arg.client)
		isPass = assertions.Error(err, "이미 존재하는 테스크에 추가하는데 성공했습니다.")
		if !isPass {
			defer manager.RemoveTask(uuid)
			defer manager.RemoveSubscriber(uuid, clientKey)
		}
	})
	assert.Equal(t, 0, manager.GetSizeOfTask())

	// 이미 존재하는 테스크에 클라이언트만 추가한 후 메시지 테스트
	t.Run("Allocate task (case: already exist task, only add client with sending message)", func(t *testing.T) {

		// given
		ch := make(chan *entity.TaskAllocResponse)
		defer close(ch)
		c := &entity.Config{
			Channel:      true,
			Notification: false,
			AutoTrading:  false,
		}
		arg := args{
			modelId: "test-model-id",
			stockId: "test-stock-id",
			client:  entity.NewClient(ch, c),
		}
		uuid := fmt.Sprintf("%x", sha256.Sum256([]byte(arg.modelId+arg.stockId)))

		// when
		assertions := assert.New(t)
		assertions.Equal(false, manager.IsExistTask(uuid), "task is exist.")

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// then
		clientKey, err := manager.Allocate(ctx, arg.modelId, arg.stockId, arg.client)
		success := assertions.NoError(err, "task allocate error.")
		if success {
			defer manager.RemoveTask(uuid)
			defer manager.RemoveSubscriber(uuid, clientKey)
		}
		assertions.Equal(true, manager.IsExistTask(uuid), "task is not exist.")

		// client만 새로 추가
		// given
		ch2 := make(chan *entity.TaskAllocResponse)
		defer close(ch2)
		c2 := &entity.Config{
			Channel:      true,
			Notification: false,
			AutoTrading:  false,
		}
		arg2 := args{
			modelId: "test-model-id",
			stockId: "test-stock-id",
			client:  entity.NewClient(ch2, c2),
		}
		uuid2 := fmt.Sprintf("%x", sha256.Sum256([]byte(arg.modelId+arg.stockId)))
		assertions.Equal(true, manager.IsExistTask(uuid2), "task is not exist.")

		// when
		clientKey2 := manager.AddSubscriber(uuid2, arg2.client)
		defer manager.RemoveSubscriber(uuid2, clientKey2)

		// then
		// 메시지 전송 후 두 client에서 receive되는지 확인
		expectedData := &entity.TaskAllocResponse{
			Status:      123,
			AccessToken: "test-access-token",
		}

		// when
		go manager.Notify(uuid, expectedData.Status, expectedData.AccessToken, true)

		// then
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			select {
			case receivedData := <-ch:
				assertions.Equal(expectedData, receivedData, "received data is not equal to expected data.")
			case <-time.After(2 * time.Second):
				t.Error("timed out waiting for data on the channel.")
			}
			wg.Done()
		}()

		wg.Add(1)
		go func() {
			select {
			case receivedData := <-ch2:
				assertions.Equal(expectedData, receivedData, "received data is not equal to expected data.")
			case <-time.After(2 * time.Second):
				t.Error("timed out waiting for data on the channel.")
			}
			wg.Done()
		}()

		wg.Wait()
	})
	assert.Equal(t, 0, manager.GetSizeOfTask())

	// 이미 존재하는 테스크에서 클라이언트, 테스크 삭제
	t.Run("Allocate task (case: already exist task, remove client, task)", func(t *testing.T) {

		// given
		ch := make(chan *entity.TaskAllocResponse)
		defer close(ch)
		c := &entity.Config{
			Channel:      true,
			Notification: false,
			AutoTrading:  false,
		}
		arg := args{
			modelId: "test-model-id",
			stockId: "test-stock-id",
			client:  entity.NewClient(ch, c),
		}
		uuid := fmt.Sprintf("%x", sha256.Sum256([]byte(arg.modelId+arg.stockId)))

		// when
		assertions := assert.New(t)
		assertions.Equal(false, manager.IsExistTask(uuid), "task is exist.")

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// then
		clientKey, err := manager.Allocate(ctx, arg.modelId, arg.stockId, arg.client)
		assertions.NoError(err, "task allocate error.")
		assertions.Equal(true, manager.IsExistTask(uuid), "task is not exist.")

		// client 삭제
		assertions.True(manager.RemoveSubscriber(uuid, clientKey), "client remove error.")
		// task 삭제
		assertions.True(manager.RemoveTask(uuid), "task remove error.")
	})
	assert.Equal(t, 0, manager.GetSizeOfTask())
}

package sw

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gitlab.bianjie.ai/adb/adb-modules/proto/superwallet/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const SuccessCode = 0

var Cli *SWClient

type SWClient struct {
	cli     pb.UserServiceClient
	timeout time.Duration
}

func Init(endPoint string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	conn, err := grpc.DialContext(ctx, endPoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Fatalf("sw did not connect: %v", err)
	}
	Cli = &SWClient{
		cli:     pb.NewUserServiceClient(conn),
		timeout: 30,
	}
}

func (c *SWClient) DelConfigCache(code string) error {
	method := "DelConfigCache"
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout*time.Second)
	defer cancel()
	req := &pb.DelConfigCacheReq{
		Base: &pb.BaseReq{RequestId: c.genRequestId()},
		Code: code,
	}
	c.logParams(method, req.Base.RequestId, req)
	resp, err := c.cli.DelConfigCache(ctx, req)
	if err != nil {
		c.logError(method, req.Base.RequestId, req, err)
		return err
	}
	if resp.Code != SuccessCode {
		c.logBadResp(method, req.Base.RequestId, req, resp, resp.Code, resp.Message)
		return fmt.Errorf("response err,code:%d,message:%s", resp.Code, resp.Message)
	}
	return nil
}

func (c *SWClient) logError(method, requestId string, params interface{}, err error) {
	logrus.WithField("request_id", requestId).
		WithField("params", params).
		//WithField("resp", res).
		Errorf("call sw spi error, method: %s, err: %v", method, err)
}

func (c *SWClient) logBadResp(method, requestId string, params, resp interface{}, respCode int32, respMsg string) {
	logrus.WithField("request_id", requestId).
		WithField("params", params).
		WithField("resp", resp).
		Errorf("sw spi bad response, method: %s, code: %d(%s)", method, respCode, respMsg)
}

func (c *SWClient) logParams(method, requestId string, params interface{}) {
	logrus.WithField("request_id", requestId).
		WithField("params", params).
		Debugf("------------>[sw] method: %s", method)
}

func (c *SWClient) logResponse(method, requestId string, resp interface{}) {
	logrus.WithField("request_id", requestId).
		WithField("resp", resp).
		Debugf("<------------[sw] method: %s", method)
}

func (c *SWClient) genRequestId() string {
	return uuid.NewString()
}

// Code generated by protoc-gen-ybrpc. DO NOT EDIT.

package rpc

// service: yb.rpc_test.CalculatorService
// service: CalculatorService

// methods
type CalculatorServiceImpl struct{}

func (s *CalculatorServiceImpl) Add(ts string, pb *AddRequestPB) (*AddResponsePB, error) {
	return &AddResponsePB{}, nil
}

func (s *CalculatorServiceImpl) Sleep(ts string, pb *SleepRequestPB) (*SleepResponsePB, error) {
	return &SleepResponsePB{}, nil
}

func (s *CalculatorServiceImpl) Echo(ts string, pb *EchoRequestPB) (*EchoResponsePB, error) {
	return &EchoResponsePB{}, nil
}

func (s *CalculatorServiceImpl) WhoAmI(ts string, pb *WhoAmIRequestPB) (*WhoAmIResponsePB, error) {
	return &WhoAmIResponsePB{}, nil
}

func (s *CalculatorServiceImpl) TestArgumentsInDiffPackage(ts string, pb *ReqDiffPackagePB) (*RespDiffPackagePB, error) {
	return &RespDiffPackagePB{}, nil
}

func (s *CalculatorServiceImpl) Panic(ts string, pb *PanicRequestPB) (*PanicResponsePB, error) {
	return &PanicResponsePB{}, nil
}

func (s *CalculatorServiceImpl) Ping(ts string, pb *PingRequestPB) (*PingResponsePB, error) {
	return &PingResponsePB{}, nil
}

func (s *CalculatorServiceImpl) Disconnect(ts string, pb *DisconnectRequestPB) (*DisconnectResponsePB, error) {
	return &DisconnectResponsePB{}, nil
}

func (s *CalculatorServiceImpl) Forward(ts string, pb *ForwardRequestPB) (*ForwardResponsePB, error) {
	return &ForwardResponsePB{}, nil
}
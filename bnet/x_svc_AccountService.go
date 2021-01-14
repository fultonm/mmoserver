// generated by protoc-gen-gcraft : DO NOT EDIT
package bnet

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	protocol "github.com/superp00t/gophercraft/bnet/bgs/protocol"
	v1 "github.com/superp00t/gophercraft/bnet/bgs/protocol/account/v1"
	math "math"
)

// shut the compiler up
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = protocol.E_MethodOptions

const AccountServiceHash = 0x62DA0891

type AccountService interface {
	ResolveAccount(*Conn, uint32, *v1.ResolveAccountRequest)
	Subscribe(*Conn, uint32, *v1.SubscriptionUpdateRequest)
	Unsubscribe(*Conn, uint32, *v1.SubscriptionUpdateRequest)
	GetAccountState(*Conn, uint32, *v1.GetAccountStateRequest)
	GetGameAccountState(*Conn, uint32, *v1.GetGameAccountStateRequest)
	GetLicenses(*Conn, uint32, *v1.GetLicensesRequest)
	GetGameTimeRemainingInfo(*Conn, uint32, *v1.GetGameTimeRemainingInfoRequest)
	GetGameSessionInfo(*Conn, uint32, *v1.GetGameSessionInfoRequest)
	GetCAISInfo(*Conn, uint32, *v1.GetCAISInfoRequest)
	GetAuthorizedData(*Conn, uint32, *v1.GetAuthorizedDataRequest)
	GetSignedAccountState(*Conn, uint32, *v1.GetSignedAccountStateRequest)
}

func DispatchAccountService(conn *Conn, svc AccountService, token uint32, method uint32, data []byte) error {
	switch method {
	case 13:
		var args v1.ResolveAccountRequest
		err := proto.Unmarshal(data, &args)
		if err != nil {
			return err
		}
		svc.ResolveAccount(conn, token, &args)
	case 25:
		var args v1.SubscriptionUpdateRequest
		err := proto.Unmarshal(data, &args)
		if err != nil {
			return err
		}
		svc.Subscribe(conn, token, &args)
	case 26:
		var args v1.SubscriptionUpdateRequest
		err := proto.Unmarshal(data, &args)
		if err != nil {
			return err
		}
		svc.Unsubscribe(conn, token, &args)
	case 30:
		var args v1.GetAccountStateRequest
		err := proto.Unmarshal(data, &args)
		if err != nil {
			return err
		}
		svc.GetAccountState(conn, token, &args)
	case 31:
		var args v1.GetGameAccountStateRequest
		err := proto.Unmarshal(data, &args)
		if err != nil {
			return err
		}
		svc.GetGameAccountState(conn, token, &args)
	case 32:
		var args v1.GetLicensesRequest
		err := proto.Unmarshal(data, &args)
		if err != nil {
			return err
		}
		svc.GetLicenses(conn, token, &args)
	case 33:
		var args v1.GetGameTimeRemainingInfoRequest
		err := proto.Unmarshal(data, &args)
		if err != nil {
			return err
		}
		svc.GetGameTimeRemainingInfo(conn, token, &args)
	case 34:
		var args v1.GetGameSessionInfoRequest
		err := proto.Unmarshal(data, &args)
		if err != nil {
			return err
		}
		svc.GetGameSessionInfo(conn, token, &args)
	case 35:
		var args v1.GetCAISInfoRequest
		err := proto.Unmarshal(data, &args)
		if err != nil {
			return err
		}
		svc.GetCAISInfo(conn, token, &args)
	case 37:
		var args v1.GetAuthorizedDataRequest
		err := proto.Unmarshal(data, &args)
		if err != nil {
			return err
		}
		svc.GetAuthorizedData(conn, token, &args)
	case 44:
		var args v1.GetSignedAccountStateRequest
		err := proto.Unmarshal(data, &args)
		if err != nil {
			return err
		}
		svc.GetSignedAccountState(conn, token, &args)
	}
	return nil
}

type EmptyAccountService struct{}

func (e EmptyAccountService) ResolveAccount(conn *Conn, token uint32, args *v1.ResolveAccountRequest) {
	conn.SendResponseCode(token, ERROR_RPC_NOT_IMPLEMENTED)
}
func (e EmptyAccountService) Subscribe(conn *Conn, token uint32, args *v1.SubscriptionUpdateRequest) {
	conn.SendResponseCode(token, ERROR_RPC_NOT_IMPLEMENTED)
}
func (e EmptyAccountService) Unsubscribe(conn *Conn, token uint32, args *v1.SubscriptionUpdateRequest) {
	conn.SendResponseCode(token, ERROR_RPC_NOT_IMPLEMENTED)
}
func (e EmptyAccountService) GetAccountState(conn *Conn, token uint32, args *v1.GetAccountStateRequest) {
	conn.SendResponseCode(token, ERROR_RPC_NOT_IMPLEMENTED)
}
func (e EmptyAccountService) GetGameAccountState(conn *Conn, token uint32, args *v1.GetGameAccountStateRequest) {
	conn.SendResponseCode(token, ERROR_RPC_NOT_IMPLEMENTED)
}
func (e EmptyAccountService) GetLicenses(conn *Conn, token uint32, args *v1.GetLicensesRequest) {
	conn.SendResponseCode(token, ERROR_RPC_NOT_IMPLEMENTED)
}
func (e EmptyAccountService) GetGameTimeRemainingInfo(conn *Conn, token uint32, args *v1.GetGameTimeRemainingInfoRequest) {
	conn.SendResponseCode(token, ERROR_RPC_NOT_IMPLEMENTED)
}
func (e EmptyAccountService) GetGameSessionInfo(conn *Conn, token uint32, args *v1.GetGameSessionInfoRequest) {
	conn.SendResponseCode(token, ERROR_RPC_NOT_IMPLEMENTED)
}
func (e EmptyAccountService) GetCAISInfo(conn *Conn, token uint32, args *v1.GetCAISInfoRequest) {
	conn.SendResponseCode(token, ERROR_RPC_NOT_IMPLEMENTED)
}
func (e EmptyAccountService) GetAuthorizedData(conn *Conn, token uint32, args *v1.GetAuthorizedDataRequest) {
	conn.SendResponseCode(token, ERROR_RPC_NOT_IMPLEMENTED)
}
func (e EmptyAccountService) GetSignedAccountState(conn *Conn, token uint32, args *v1.GetSignedAccountStateRequest) {
	conn.SendResponseCode(token, ERROR_RPC_NOT_IMPLEMENTED)
}

func (c *Conn) AccountService_Request_ResolveAccount(args *v1.ResolveAccountRequest) (*v1.ResolveAccountResponse, error) {
	header, bytes, err := c.Request(AccountServiceHash, 13, args)
	if err != nil {
		return nil, err
	}
	if Status(header.GetStatus()) != ERROR_OK {
		return nil, fmt.Errorf("bnet: non-ok status 0x%08X", header.GetStatus())
	}
	var out v1.ResolveAccountResponse
	err = proto.Unmarshal(bytes, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Conn) AccountService_Request_Subscribe(args *v1.SubscriptionUpdateRequest) (*v1.SubscriptionUpdateResponse, error) {
	header, bytes, err := c.Request(AccountServiceHash, 25, args)
	if err != nil {
		return nil, err
	}
	if Status(header.GetStatus()) != ERROR_OK {
		return nil, fmt.Errorf("bnet: non-ok status 0x%08X", header.GetStatus())
	}
	var out v1.SubscriptionUpdateResponse
	err = proto.Unmarshal(bytes, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Conn) AccountService_Request_Unsubscribe(args *v1.SubscriptionUpdateRequest) error {
	header, _, err := c.Request(AccountServiceHash, 26, args)
	if err != nil {
		return err
	}
	if Status(header.GetStatus()) != ERROR_OK {
		return fmt.Errorf("bnet: non-ok status 0x%08X", header.GetStatus())
	}
	return nil
}

func (c *Conn) AccountService_Request_GetAccountState(args *v1.GetAccountStateRequest) (*v1.GetAccountStateResponse, error) {
	header, bytes, err := c.Request(AccountServiceHash, 30, args)
	if err != nil {
		return nil, err
	}
	if Status(header.GetStatus()) != ERROR_OK {
		return nil, fmt.Errorf("bnet: non-ok status 0x%08X", header.GetStatus())
	}
	var out v1.GetAccountStateResponse
	err = proto.Unmarshal(bytes, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Conn) AccountService_Request_GetGameAccountState(args *v1.GetGameAccountStateRequest) (*v1.GetGameAccountStateResponse, error) {
	header, bytes, err := c.Request(AccountServiceHash, 31, args)
	if err != nil {
		return nil, err
	}
	if Status(header.GetStatus()) != ERROR_OK {
		return nil, fmt.Errorf("bnet: non-ok status 0x%08X", header.GetStatus())
	}
	var out v1.GetGameAccountStateResponse
	err = proto.Unmarshal(bytes, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Conn) AccountService_Request_GetLicenses(args *v1.GetLicensesRequest) (*v1.GetLicensesResponse, error) {
	header, bytes, err := c.Request(AccountServiceHash, 32, args)
	if err != nil {
		return nil, err
	}
	if Status(header.GetStatus()) != ERROR_OK {
		return nil, fmt.Errorf("bnet: non-ok status 0x%08X", header.GetStatus())
	}
	var out v1.GetLicensesResponse
	err = proto.Unmarshal(bytes, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Conn) AccountService_Request_GetGameTimeRemainingInfo(args *v1.GetGameTimeRemainingInfoRequest) (*v1.GetGameTimeRemainingInfoResponse, error) {
	header, bytes, err := c.Request(AccountServiceHash, 33, args)
	if err != nil {
		return nil, err
	}
	if Status(header.GetStatus()) != ERROR_OK {
		return nil, fmt.Errorf("bnet: non-ok status 0x%08X", header.GetStatus())
	}
	var out v1.GetGameTimeRemainingInfoResponse
	err = proto.Unmarshal(bytes, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Conn) AccountService_Request_GetGameSessionInfo(args *v1.GetGameSessionInfoRequest) (*v1.GetGameSessionInfoResponse, error) {
	header, bytes, err := c.Request(AccountServiceHash, 34, args)
	if err != nil {
		return nil, err
	}
	if Status(header.GetStatus()) != ERROR_OK {
		return nil, fmt.Errorf("bnet: non-ok status 0x%08X", header.GetStatus())
	}
	var out v1.GetGameSessionInfoResponse
	err = proto.Unmarshal(bytes, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Conn) AccountService_Request_GetCAISInfo(args *v1.GetCAISInfoRequest) (*v1.GetCAISInfoResponse, error) {
	header, bytes, err := c.Request(AccountServiceHash, 35, args)
	if err != nil {
		return nil, err
	}
	if Status(header.GetStatus()) != ERROR_OK {
		return nil, fmt.Errorf("bnet: non-ok status 0x%08X", header.GetStatus())
	}
	var out v1.GetCAISInfoResponse
	err = proto.Unmarshal(bytes, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Conn) AccountService_Request_GetAuthorizedData(args *v1.GetAuthorizedDataRequest) (*v1.GetAuthorizedDataResponse, error) {
	header, bytes, err := c.Request(AccountServiceHash, 37, args)
	if err != nil {
		return nil, err
	}
	if Status(header.GetStatus()) != ERROR_OK {
		return nil, fmt.Errorf("bnet: non-ok status 0x%08X", header.GetStatus())
	}
	var out v1.GetAuthorizedDataResponse
	err = proto.Unmarshal(bytes, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Conn) AccountService_Request_GetSignedAccountState(args *v1.GetSignedAccountStateRequest) (*v1.GetSignedAccountStateResponse, error) {
	header, bytes, err := c.Request(AccountServiceHash, 44, args)
	if err != nil {
		return nil, err
	}
	if Status(header.GetStatus()) != ERROR_OK {
		return nil, fmt.Errorf("bnet: non-ok status 0x%08X", header.GetStatus())
	}
	var out v1.GetSignedAccountStateResponse
	err = proto.Unmarshal(bytes, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

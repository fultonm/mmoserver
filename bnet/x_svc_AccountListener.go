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

const AccountListenerHash = 0x54DFDA17

type AccountListener interface {
	OnAccountStateUpdated(*Conn, uint32, *v1.AccountStateNotification)
	OnGameAccountStateUpdated(*Conn, uint32, *v1.GameAccountStateNotification)
	OnGameAccountsUpdated(*Conn, uint32, *v1.GameAccountNotification)
	OnGameSessionUpdated(*Conn, uint32, *v1.GameAccountSessionNotification)
}

func DispatchAccountListener(conn *Conn, svc AccountListener, token uint32, method uint32, data []byte) error {
	switch method {
	case 1:
		var args v1.AccountStateNotification
		err := proto.Unmarshal(data, &args)
		if err != nil {
			return err
		}
		svc.OnAccountStateUpdated(conn, token, &args)
	case 2:
		var args v1.GameAccountStateNotification
		err := proto.Unmarshal(data, &args)
		if err != nil {
			return err
		}
		svc.OnGameAccountStateUpdated(conn, token, &args)
	case 3:
		var args v1.GameAccountNotification
		err := proto.Unmarshal(data, &args)
		if err != nil {
			return err
		}
		svc.OnGameAccountsUpdated(conn, token, &args)
	case 4:
		var args v1.GameAccountSessionNotification
		err := proto.Unmarshal(data, &args)
		if err != nil {
			return err
		}
		svc.OnGameSessionUpdated(conn, token, &args)
	}
	return nil
}

type EmptyAccountListener struct{}

func (e EmptyAccountListener) OnAccountStateUpdated(conn *Conn, token uint32, args *v1.AccountStateNotification) {
	conn.SendResponseCode(token, ERROR_RPC_NOT_IMPLEMENTED)
}
func (e EmptyAccountListener) OnGameAccountStateUpdated(conn *Conn, token uint32, args *v1.GameAccountStateNotification) {
	conn.SendResponseCode(token, ERROR_RPC_NOT_IMPLEMENTED)
}
func (e EmptyAccountListener) OnGameAccountsUpdated(conn *Conn, token uint32, args *v1.GameAccountNotification) {
	conn.SendResponseCode(token, ERROR_RPC_NOT_IMPLEMENTED)
}
func (e EmptyAccountListener) OnGameSessionUpdated(conn *Conn, token uint32, args *v1.GameAccountSessionNotification) {
	conn.SendResponseCode(token, ERROR_RPC_NOT_IMPLEMENTED)
}

func (c *Conn) AccountListener_Request_OnAccountStateUpdated(args *v1.AccountStateNotification) error {
	header, _, err := c.Request(AccountListenerHash, 1, args)
	if err != nil {
		return err
	}
	if Status(header.GetStatus()) != ERROR_OK {
		return fmt.Errorf("bnet: non-ok status 0x%08X", header.GetStatus())
	}
	return nil
}

func (c *Conn) AccountListener_Request_OnGameAccountStateUpdated(args *v1.GameAccountStateNotification) error {
	header, _, err := c.Request(AccountListenerHash, 2, args)
	if err != nil {
		return err
	}
	if Status(header.GetStatus()) != ERROR_OK {
		return fmt.Errorf("bnet: non-ok status 0x%08X", header.GetStatus())
	}
	return nil
}

func (c *Conn) AccountListener_Request_OnGameAccountsUpdated(args *v1.GameAccountNotification) error {
	header, _, err := c.Request(AccountListenerHash, 3, args)
	if err != nil {
		return err
	}
	if Status(header.GetStatus()) != ERROR_OK {
		return fmt.Errorf("bnet: non-ok status 0x%08X", header.GetStatus())
	}
	return nil
}

func (c *Conn) AccountListener_Request_OnGameSessionUpdated(args *v1.GameAccountSessionNotification) error {
	header, _, err := c.Request(AccountListenerHash, 4, args)
	if err != nil {
		return err
	}
	if Status(header.GetStatus()) != ERROR_OK {
		return fmt.Errorf("bnet: non-ok status 0x%08X", header.GetStatus())
	}
	return nil
}

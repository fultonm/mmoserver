// generated by protoc-gen-gcraft : DO NOT EDIT
package bnet

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	protocol "github.com/superp00t/gophercraft/bnet/bgs/protocol"
	v1 "github.com/superp00t/gophercraft/bnet/bgs/protocol/friends/v1"
	math "math"
)

// shut the compiler up
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = protocol.E_MethodOptions

const FriendsListenerHash = 0x6F259A13

type FriendsListener interface {
	OnFriendAdded(*Conn, uint32, *v1.FriendNotification)
	OnFriendRemoved(*Conn, uint32, *v1.FriendNotification)
	OnReceivedInvitationAdded(*Conn, uint32, *v1.InvitationNotification)
	OnReceivedInvitationRemoved(*Conn, uint32, *v1.InvitationNotification)
	OnSentInvitationAdded(*Conn, uint32, *v1.SentInvitationAddedNotification)
	OnSentInvitationRemoved(*Conn, uint32, *v1.SentInvitationRemovedNotification)
	OnUpdateFriendState(*Conn, uint32, *v1.UpdateFriendStateNotification)
}

func DispatchFriendsListener(conn *Conn, svc FriendsListener, token uint32, method uint32, data []byte) error {
	switch method {
	case 1:
		var args v1.FriendNotification
		err := proto.Unmarshal(data, &args)
		if err != nil {
			return err
		}
		svc.OnFriendAdded(conn, token, &args)
	case 2:
		var args v1.FriendNotification
		err := proto.Unmarshal(data, &args)
		if err != nil {
			return err
		}
		svc.OnFriendRemoved(conn, token, &args)
	case 3:
		var args v1.InvitationNotification
		err := proto.Unmarshal(data, &args)
		if err != nil {
			return err
		}
		svc.OnReceivedInvitationAdded(conn, token, &args)
	case 4:
		var args v1.InvitationNotification
		err := proto.Unmarshal(data, &args)
		if err != nil {
			return err
		}
		svc.OnReceivedInvitationRemoved(conn, token, &args)
	case 5:
		var args v1.SentInvitationAddedNotification
		err := proto.Unmarshal(data, &args)
		if err != nil {
			return err
		}
		svc.OnSentInvitationAdded(conn, token, &args)
	case 6:
		var args v1.SentInvitationRemovedNotification
		err := proto.Unmarshal(data, &args)
		if err != nil {
			return err
		}
		svc.OnSentInvitationRemoved(conn, token, &args)
	case 7:
		var args v1.UpdateFriendStateNotification
		err := proto.Unmarshal(data, &args)
		if err != nil {
			return err
		}
		svc.OnUpdateFriendState(conn, token, &args)
	}
	return nil
}

type EmptyFriendsListener struct{}

func (e EmptyFriendsListener) OnFriendAdded(conn *Conn, token uint32, args *v1.FriendNotification) {
	conn.SendResponseCode(token, ERROR_RPC_NOT_IMPLEMENTED)
}
func (e EmptyFriendsListener) OnFriendRemoved(conn *Conn, token uint32, args *v1.FriendNotification) {
	conn.SendResponseCode(token, ERROR_RPC_NOT_IMPLEMENTED)
}
func (e EmptyFriendsListener) OnReceivedInvitationAdded(conn *Conn, token uint32, args *v1.InvitationNotification) {
	conn.SendResponseCode(token, ERROR_RPC_NOT_IMPLEMENTED)
}
func (e EmptyFriendsListener) OnReceivedInvitationRemoved(conn *Conn, token uint32, args *v1.InvitationNotification) {
	conn.SendResponseCode(token, ERROR_RPC_NOT_IMPLEMENTED)
}
func (e EmptyFriendsListener) OnSentInvitationAdded(conn *Conn, token uint32, args *v1.SentInvitationAddedNotification) {
	conn.SendResponseCode(token, ERROR_RPC_NOT_IMPLEMENTED)
}
func (e EmptyFriendsListener) OnSentInvitationRemoved(conn *Conn, token uint32, args *v1.SentInvitationRemovedNotification) {
	conn.SendResponseCode(token, ERROR_RPC_NOT_IMPLEMENTED)
}
func (e EmptyFriendsListener) OnUpdateFriendState(conn *Conn, token uint32, args *v1.UpdateFriendStateNotification) {
	conn.SendResponseCode(token, ERROR_RPC_NOT_IMPLEMENTED)
}

func (c *Conn) FriendsListener_Request_OnFriendAdded(args *v1.FriendNotification) error {
	header, _, err := c.Request(FriendsListenerHash, 1, args)
	if err != nil {
		return err
	}
	if Status(header.GetStatus()) != ERROR_OK {
		return fmt.Errorf("bnet: non-ok status 0x%08X", header.GetStatus())
	}
	return nil
}

func (c *Conn) FriendsListener_Request_OnFriendRemoved(args *v1.FriendNotification) error {
	header, _, err := c.Request(FriendsListenerHash, 2, args)
	if err != nil {
		return err
	}
	if Status(header.GetStatus()) != ERROR_OK {
		return fmt.Errorf("bnet: non-ok status 0x%08X", header.GetStatus())
	}
	return nil
}

func (c *Conn) FriendsListener_Request_OnReceivedInvitationAdded(args *v1.InvitationNotification) error {
	header, _, err := c.Request(FriendsListenerHash, 3, args)
	if err != nil {
		return err
	}
	if Status(header.GetStatus()) != ERROR_OK {
		return fmt.Errorf("bnet: non-ok status 0x%08X", header.GetStatus())
	}
	return nil
}

func (c *Conn) FriendsListener_Request_OnReceivedInvitationRemoved(args *v1.InvitationNotification) error {
	header, _, err := c.Request(FriendsListenerHash, 4, args)
	if err != nil {
		return err
	}
	if Status(header.GetStatus()) != ERROR_OK {
		return fmt.Errorf("bnet: non-ok status 0x%08X", header.GetStatus())
	}
	return nil
}

func (c *Conn) FriendsListener_Request_OnSentInvitationAdded(args *v1.SentInvitationAddedNotification) error {
	header, _, err := c.Request(FriendsListenerHash, 5, args)
	if err != nil {
		return err
	}
	if Status(header.GetStatus()) != ERROR_OK {
		return fmt.Errorf("bnet: non-ok status 0x%08X", header.GetStatus())
	}
	return nil
}

func (c *Conn) FriendsListener_Request_OnSentInvitationRemoved(args *v1.SentInvitationRemovedNotification) error {
	header, _, err := c.Request(FriendsListenerHash, 6, args)
	if err != nil {
		return err
	}
	if Status(header.GetStatus()) != ERROR_OK {
		return fmt.Errorf("bnet: non-ok status 0x%08X", header.GetStatus())
	}
	return nil
}

func (c *Conn) FriendsListener_Request_OnUpdateFriendState(args *v1.UpdateFriendStateNotification) error {
	header, _, err := c.Request(FriendsListenerHash, 7, args)
	if err != nil {
		return err
	}
	if Status(header.GetStatus()) != ERROR_OK {
		return fmt.Errorf("bnet: non-ok status 0x%08X", header.GetStatus())
	}
	return nil
}

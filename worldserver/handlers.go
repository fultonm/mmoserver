package worldserver

import (
	p "github.com/superp00t/gophercraft/packet"
	v "github.com/superp00t/gophercraft/vsn"
)

type Handlers struct {
	Map map[p.WorldType]*WorldClientHandler
}

func (ws *WorldServer) initHandlers() {
	h := &Handlers{make(map[p.WorldType]*WorldClientHandler)}

	// Classic protocol selector
	c := v.Selector("-12340")

	h.On(p.CMSG_PET_NAME_QUERY, c, 0, (*Session).SendPet)
	h.On(p.CMSG_WARDEN_DATA, c, 0, (*Session).WardenResponse)
	h.On(p.CMSG_PING, c, 0, (*Session).HandlePong)
	h.On(p.CMSG_UPDATE_ACCOUNT_DATA, c, 0, (*Session).HandleAccountDataUpdate)

	// Character selection menu
	h.On(p.CMSG_CHAR_ENUM, c, 0, (*Session).CharacterList)
	h.On(p.CMSG_CHAR_DELETE, c, 0, (*Session).DeleteCharacter)
	h.On(p.CMSG_CHAR_CREATE, c, 0, (*Session).CreateCharacter)
	h.On(p.CMSG_PLAYER_LOGIN, c, 0, (*Session).HandleJoin)
	h.On(p.CMSG_LOGOUT_REQUEST, c, 1, (*Session).HandleLogoutRequest)

	h.On(p.CMSG_NAME_QUERY, c, 1, (*Session).HandleNameQuery)
	h.On(p.CMSG_MESSAGECHAT, c, 1, (*Session).HandleChat)
	h.On(p.CMSG_WHO, c, 1, (*Session).HandleWho)
	h.On(p.CMSG_SET_SELECTION, c, 1, (*Session).HandleTarget)

	// Movement
	h.On(p.MSG_MOVE_HEARTBEAT, c, 1, (*Session).HandleMoves)
	h.On(p.MSG_MOVE_START_FORWARD, c, 1, (*Session).HandleMoves)
	h.On(p.MSG_MOVE_START_BACKWARD, c, 1, (*Session).HandleMoves)
	h.On(p.MSG_MOVE_STOP, c, 1, (*Session).HandleMoves)
	h.On(p.MSG_MOVE_START_STRAFE_LEFT, c, 1, (*Session).HandleMoves)
	h.On(p.MSG_MOVE_START_STRAFE_RIGHT, c, 1, (*Session).HandleMoves)
	h.On(p.MSG_MOVE_STOP_STRAFE, c, 1, (*Session).HandleMoves)
	h.On(p.MSG_MOVE_JUMP, c, 1, (*Session).HandleMoves)
	h.On(p.MSG_MOVE_START_TURN_LEFT, c, 1, (*Session).HandleMoves)
	h.On(p.MSG_MOVE_START_TURN_RIGHT, c, 1, (*Session).HandleMoves)
	h.On(p.MSG_MOVE_STOP_TURN, c, 1, (*Session).HandleMoves)
	h.On(p.MSG_MOVE_START_PITCH_UP, c, 1, (*Session).HandleMoves)
	h.On(p.MSG_MOVE_START_PITCH_DOWN, c, 1, (*Session).HandleMoves)
	h.On(p.MSG_MOVE_SET_PITCH, c, 1, (*Session).HandleMoves)
	h.On(p.MSG_MOVE_STOP_PITCH, c, 1, (*Session).HandleMoves)
	h.On(p.MSG_MOVE_SET_RUN_MODE, c, 1, (*Session).HandleMoves)
	h.On(p.MSG_MOVE_SET_FACING, c, 1, (*Session).HandleMoves)
	h.On(p.MSG_MOVE_FALL_LAND, c, 1, (*Session).HandleMoves)
	h.On(p.MSG_MOVE_START_SWIM, c, 1, (*Session).HandleMoves)
	h.On(p.MSG_MOVE_STOP_SWIM, c, 1, (*Session).HandleMoves)
	h.On(p.MSG_MOVE_WORLDPORT_ACK, c, 1, (*Session).HandleWorldportAck)
	h.On(p.CMSG_SUMMON_RESPONSE, c, 1, (*Session).HandleSummonResponse)

	// Location
	h.On(p.CMSG_ZONEUPDATE, c, 1, (*Session).HandleZoneUpdate)
	h.On(p.CMSG_AREATRIGGER, c, 1, (*Session).HandleAreaTrigger)

	// Animation
	h.On(p.CMSG_STANDSTATECHANGE, c, 1, (*Session).HandleStandStateChange)
	h.On(p.CMSG_TEXT_EMOTE, c, 1, (*Session).HandleTextEmote)

	// Gameobjects
	h.On(p.CMSG_GAMEOBJECT_QUERY, c, 1, (*Session).HandleGameObjectQuery)
	h.On(p.CMSG_GAMEOBJ_USE, c, 1, (*Session).HandleGameObjectUse)

	// Inventory
	h.On(p.CMSG_ITEM_QUERY_SINGLE, c, 1, (*Session).HandleItemQuery)
	h.On(p.CMSG_SWAP_INV_ITEM, c, 1, (*Session).HandleSwapInventoryItem)
	h.On(p.CMSG_SWAP_ITEM, c, 1, (*Session).HandleSwapItem)
	h.On(p.CMSG_AUTOEQUIP_ITEM, c, 1, (*Session).HandleAutoEquipItem)
	h.On(p.CMSG_DESTROYITEM, c, 1, (*Session).HandleDestroyItem)
	h.On(p.CMSG_SPLIT_ITEM, c, 1, (*Session).HandleSplitItem)

	ws.handlers = h
}

type HandlerOptionType uint8

const (
	AsynchronousHandler HandlerOptionType = iota
)

type HandlerOption struct {
	Type HandlerOptionType
}

var Async = &HandlerOption{Type: AsynchronousHandler}

func (h *Handlers) On(pt p.WorldType, selector v.Selector, requiredState SessionState, function interface{}, option ...*HandlerOption) {
	h.Map[pt] = &WorldClientHandler{pt, selector, requiredState, function}
}

type WorldClientHandler struct {
	Op            p.WorldType
	Selector      v.Selector
	RequiredState SessionState
	Fn            interface{}
}

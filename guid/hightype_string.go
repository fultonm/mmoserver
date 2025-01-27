// Code generated by "gcraft_stringer -type=HighType"; DO NOT EDIT.

package guid

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Null-0]
	_ = x[Uniq-1]
	_ = x[Player-2]
	_ = x[Item-3]
	_ = x[WorldTransaction-4]
	_ = x[StaticDoor-5]
	_ = x[Transport-6]
	_ = x[Conversation-7]
	_ = x[Creature-8]
	_ = x[Vehicle-9]
	_ = x[Pet-10]
	_ = x[GameObject-11]
	_ = x[DynamicObject-12]
	_ = x[AreaTrigger-13]
	_ = x[Corpse-14]
	_ = x[LootObject-15]
	_ = x[SceneObject-16]
	_ = x[Scenario-17]
	_ = x[AIGroup-18]
	_ = x[DynamicDoor-19]
	_ = x[ClientActor-20]
	_ = x[Vignette-21]
	_ = x[CallForHelp-22]
	_ = x[AIResource-23]
	_ = x[AILock-24]
	_ = x[AILockTicket-25]
	_ = x[ChatChannel-26]
	_ = x[Party-27]
	_ = x[Guild-28]
	_ = x[WowAccount-29]
	_ = x[BNetAccount-30]
	_ = x[GMTask-31]
	_ = x[MobileSession-32]
	_ = x[RaidGroup-33]
	_ = x[Spell-34]
	_ = x[Mail-35]
	_ = x[WebObj-36]
	_ = x[LFGObject-37]
	_ = x[LFGList-38]
	_ = x[UserRouter-39]
	_ = x[PVPQueueGroup-40]
	_ = x[UserClient-41]
	_ = x[PetBattle-42]
	_ = x[UniqUserClient-43]
	_ = x[BattlePet-44]
	_ = x[CommerceObj-45]
	_ = x[ClientSession-46]
	_ = x[Cast-47]
	_ = x[ClientConnection-48]
	_ = x[Mo_Transport-60]
	_ = x[Instance-61]
	_ = x[Group-62]
}

const (
	_HighType_name_0 = "NullUniqPlayerItemWorldTransactionStaticDoorTransportConversationCreatureVehiclePetGameObjectDynamicObjectAreaTriggerCorpseLootObjectSceneObjectScenarioAIGroupDynamicDoorClientActorVignetteCallForHelpAIResourceAILockAILockTicketChatChannelPartyGuildWowAccountBNetAccountGMTaskMobileSessionRaidGroupSpellMailWebObjLFGObjectLFGListUserRouterPVPQueueGroupUserClientPetBattleUniqUserClientBattlePetCommerceObjClientSessionCastClientConnection"
	_HighType_name_1 = "Mo_TransportInstanceGroup"
)

var (
	_HighType_index_0 = [...]uint16{0, 4, 8, 14, 18, 34, 44, 53, 65, 73, 80, 83, 93, 106, 117, 123, 133, 144, 152, 159, 170, 181, 189, 200, 210, 216, 228, 239, 244, 249, 259, 270, 276, 289, 298, 303, 307, 313, 322, 329, 339, 352, 362, 371, 385, 394, 405, 418, 422, 438}
	_HighType_index_1 = [...]uint8{0, 12, 20, 25}
)

func (i HighType) String() string {
	switch {
	case 0 <= i && i <= 48:
		return _HighType_name_0[_HighType_index_0[i]:_HighType_index_0[i+1]]
	case 60 <= i && i <= 62:
		i -= 60
		return _HighType_name_1[_HighType_index_1[i]:_HighType_index_1[i+1]]
	default:
		return "HighType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}

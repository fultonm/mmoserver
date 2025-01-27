package update

import (
	"encoding/binary"
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/superp00t/gophercraft/vsn"
)

//go:generate gcraft_stringer -type=ItemFlag -method=ToString -fromString
type ItemFlag uint64

const (
	ItemFlagNoPickup ItemFlag = 1 << iota
	ItemFlagConjured
	ItemFlagHasLoot
	ItemFlagHeroicTooltip
	ItemFlagDeprecated
	ItemFlagNoUserDestroy
	ItemFlagPlayerCast
	ItemFlagNoEquipCooldown
	ItemFlagMultiLootQuest
	ItemFlagIsWrapper
	ItemFlagUsesResources
	ItemFlagMultiDrop
	ItemFlagItemPurchaseRecord
	ItemFlagPetition
	ItemFlagHasText
	ItemFlagNoDisenchant
	ItemFlagRealDuration
	ItemFlagNoCreator
	ItemFlagIsProspectable
	ItemFlagUniqueEquippable
	ItemFlagIgnoreForAuras
	ItemFlagIgnoreDefaultArenaRestrictions
	ItemFlagNoDurabilityLoss
	ItemFlagUseWhenShapeshifted
	ItemFlagHasQuestGlow
	ItemFlagHideUnusableRecipe
	ItemFlagNotUseableInArena
	ItemFlagBindsToAccount
	ItemFlagNoReagentCost
	ItemFlagIsMillable
	ItemFlagReportToGuildChat
	ItemFlagNoProgressiveLoot
)

type ItemFlagDescriptor map[ItemFlag]uint64

var ItemFlagDescriptors = map[vsn.Build]ItemFlagDescriptor{
	vsn.Alpha: {
		// Guessed cuz lazy
		ItemFlagNoPickup:                       0x00000001, // not used
		ItemFlagConjured:                       0x00000002,
		ItemFlagHasLoot:                        0x00000004, // affect only non container items that can be "open" for loot. It or lockid set enable for client sh$
		ItemFlagHeroicTooltip:                  0x00000008, // heroic item version
		ItemFlagDeprecated:                     0x00000010, // can't repeat old note: appears red icon (like when item durability==0)
		ItemFlagNoUserDestroy:                  0x00000020, // used for totem. Item can not be destroyed, except by using spell (item can be reagent for spell an$
		ItemFlagPlayerCast:                     0x00000040, // ? old note: usable
		ItemFlagNoEquipCooldown:                0x00000080,
		ItemFlagMultiLootQuest:                 0x00000100, // saw this on item 47115, 49295...
		ItemFlagIsWrapper:                      0x00000200, // used or not used wrapper
		ItemFlagUsesResources:                  0x00000400, // ignore bag space at new item creation?
		ItemFlagMultiDrop:                      0x00000800, // determines if item is party loot or not
		ItemFlagItemPurchaseRecord:             0x00001000, // item cost can be refunded within 2 hours after purchase
		ItemFlagPetition:                       0x00002000, // arena/guild charter
		ItemFlagHasText:                        0x00004000,
		ItemFlagNoDisenchant:                   0x00008000, // a lot of items have this
		ItemFlagRealDuration:                   0x00010000, // a lot of items have this
		ItemFlagNoCreator:                      0x00020000,
		ItemFlagIsProspectable:                 0x00040000, // item can have prospecting loot (in fact some items expected have empty loot)
		ItemFlagUniqueEquippable:               0x00080000,
		ItemFlagIgnoreForAuras:                 0x00100000,
		ItemFlagIgnoreDefaultArenaRestrictions: 0x00200000, // last used flag in 1.12.1
	},

	vsn.V1_12_1: {
		ItemFlagNoPickup:                       0x00000001, // not used
		ItemFlagConjured:                       0x00000002,
		ItemFlagHasLoot:                        0x00000004, // affect only non container items that can be "open" for loot. It or lockid set enable for client sh$
		ItemFlagHeroicTooltip:                  0x00000008, // heroic item version
		ItemFlagDeprecated:                     0x00000010, // can't repeat old note: appears red icon (like when item durability==0)
		ItemFlagNoUserDestroy:                  0x00000020, // used for totem. Item can not be destroyed, except by using spell (item can be reagent for spell an$
		ItemFlagPlayerCast:                     0x00000040, // ? old note: usable
		ItemFlagNoEquipCooldown:                0x00000080,
		ItemFlagMultiLootQuest:                 0x00000100, // saw this on item 47115, 49295...
		ItemFlagIsWrapper:                      0x00000200, // used or not used wrapper
		ItemFlagUsesResources:                  0x00000400, // ignore bag space at new item creation?
		ItemFlagMultiDrop:                      0x00000800, // determines if item is party loot or not
		ItemFlagItemPurchaseRecord:             0x00001000, // item cost can be refunded within 2 hours after purchase
		ItemFlagPetition:                       0x00002000, // arena/guild charter
		ItemFlagHasText:                        0x00004000,
		ItemFlagNoDisenchant:                   0x00008000, // a lot of items have this
		ItemFlagRealDuration:                   0x00010000, // a lot of items have this
		ItemFlagNoCreator:                      0x00020000,
		ItemFlagIsProspectable:                 0x00040000, // item can have prospecting loot (in fact some items expected have empty loot)
		ItemFlagUniqueEquippable:               0x00080000,
		ItemFlagIgnoreForAuras:                 0x00100000,
		ItemFlagIgnoreDefaultArenaRestrictions: 0x00200000, // last used flag in 1.12.1
	},

	vsn.V2_4_3: {
		ItemFlagNoPickup:                       0x00000001, // not used
		ItemFlagConjured:                       0x00000002,
		ItemFlagHasLoot:                        0x00000004, // affect only non container items that can be "open" for loot. It or lockid set enable for client sh$
		ItemFlagHeroicTooltip:                  0x00000008, // heroic item version
		ItemFlagDeprecated:                     0x00000010, // can't repeat old note: appears red icon (like when item durability==0)
		ItemFlagNoUserDestroy:                  0x00000020, // used for totem. Item can not be destroyed, except by using spell (item can be reagent for spell an$
		ItemFlagPlayerCast:                     0x00000040, // ? old note: usable
		ItemFlagNoEquipCooldown:                0x00000080,
		ItemFlagMultiLootQuest:                 0x00000100, // saw this on item 47115, 49295...
		ItemFlagIsWrapper:                      0x00000200, // used or not used wrapper
		ItemFlagUsesResources:                  0x00000400, // ignore bag space at new item creation?
		ItemFlagMultiDrop:                      0x00000800, // determines if item is party loot or not
		ItemFlagItemPurchaseRecord:             0x00001000, // item cost can be refunded within 2 hours after purchase
		ItemFlagPetition:                       0x00002000, // arena/guild charter
		ItemFlagHasText:                        0x00004000,
		ItemFlagNoDisenchant:                   0x00008000, // a lot of items have this
		ItemFlagRealDuration:                   0x00010000, // a lot of items have this
		ItemFlagNoCreator:                      0x00020000,
		ItemFlagIsProspectable:                 0x00040000, // item can have prospecting loot (in fact some items expected have empty loot)
		ItemFlagUniqueEquippable:               0x00080000,
		ItemFlagIgnoreForAuras:                 0x00100000,
		ItemFlagIgnoreDefaultArenaRestrictions: 0x00200000, // last used flag in 1.12.1
		ItemFlagNoDurabilityLoss:               0x00400000,
		ItemFlagUseWhenShapeshifted:            0x00800000,
	},

	vsn.V3_3_5a: {
		ItemFlagNoPickup:                       0x00000001,
		ItemFlagConjured:                       0x00000002, // Conjured item
		ItemFlagHasLoot:                        0x00000004, // Item can be right clicked to open for loot
		ItemFlagHeroicTooltip:                  0x00000008, // Makes green "Heroic" text appear on item
		ItemFlagDeprecated:                     0x00000010, // Cannot equip or use
		ItemFlagNoUserDestroy:                  0x00000020, // Item can not be destroyed, except by using spell (item can be reagent for spell)
		ItemFlagPlayerCast:                     0x00000040, // Item's spells are castable by players
		ItemFlagNoEquipCooldown:                0x00000080, // No default 30 seconds cooldown when equipped
		ItemFlagMultiLootQuest:                 0x00000100,
		ItemFlagIsWrapper:                      0x00000200, // Item can wrap other items
		ItemFlagUsesResources:                  0x00000400,
		ItemFlagMultiDrop:                      0x00000800, // Looting this item does not remove it from available loot
		ItemFlagItemPurchaseRecord:             0x00001000, // Item can be returned to vendor for its original cost (extended cost)
		ItemFlagPetition:                       0x00002000, // Item is guild or arena charter
		ItemFlagHasText:                        0x00004000, // Only readable items have this (but not all)
		ItemFlagNoDisenchant:                   0x00008000,
		ItemFlagRealDuration:                   0x00010000,
		ItemFlagNoCreator:                      0x00020000,
		ItemFlagIsProspectable:                 0x00040000, // Item can be prospected
		ItemFlagUniqueEquippable:               0x00080000, // You can only equip one of these
		ItemFlagIgnoreForAuras:                 0x00100000,
		ItemFlagIgnoreDefaultArenaRestrictions: 0x00200000, // Item can be used during arena match
		ItemFlagNoDurabilityLoss:               0x00400000, // Some Thrown weapons have it (and only Thrown) but not all
		ItemFlagUseWhenShapeshifted:            0x00800000, // Item can be used in shapeshift forms
		ItemFlagHasQuestGlow:                   0x01000000,
		ItemFlagHideUnusableRecipe:             0x02000000, // Profession recipes: can only be looted if you meet requirements and don't already know it
		ItemFlagNotUseableInArena:              0x04000000, // Item cannot be used in arena
		ItemFlagBindsToAccount:                 0x08000000, // Item binds to account and can be sent only to your own characters
		ItemFlagNoReagentCost:                  0x10000000, // Spell is cast ignoring reagents
		ItemFlagIsMillable:                     0x20000000, // Item can be milled
		ItemFlagReportToGuildChat:              0x40000000,
		ItemFlagNoProgressiveLoot:              0x80000000,
	},
}

func (x ItemFlag) EncodeWord() (string, error) {
	return x.String(), nil
}

func (x ItemFlag) DecodeWord(out reflect.Value, str string) error {
	data, err := ParseItemFlag(str)
	if err != nil {
		return err
	}
	out.Set(reflect.ValueOf(data))
	return nil
}

func (x ItemFlag) String() string {
	var s []string
	if x&ItemFlagNoPickup != 0 {
		s = append(s, "NoPickup")
	}
	if x&ItemFlagConjured != 0 {
		s = append(s, "Conjured")
	}
	if x&ItemFlagHasLoot != 0 {
		s = append(s, "HasLoot")
	}
	if x&ItemFlagHeroicTooltip != 0 {
		s = append(s, "HeroicTooltip")
	}
	if x&ItemFlagDeprecated != 0 {
		s = append(s, "Deprecated")
	}
	if x&ItemFlagNoUserDestroy != 0 {
		s = append(s, "NoUserDestroy")
	}
	if x&ItemFlagPlayerCast != 0 {
		s = append(s, "PlayerCast")
	}
	if x&ItemFlagNoEquipCooldown != 0 {
		s = append(s, "NoEquipCooldown")
	}
	if x&ItemFlagMultiLootQuest != 0 {
		s = append(s, "MultiLootQuest")
	}
	if x&ItemFlagIsWrapper != 0 {
		s = append(s, "IsWrapper")
	}
	if x&ItemFlagUsesResources != 0 {
		s = append(s, "UsesResources")
	}
	if x&ItemFlagMultiDrop != 0 {
		s = append(s, "MultiDrop")
	}
	if x&ItemFlagItemPurchaseRecord != 0 {
		s = append(s, "ItemPurchaseRecord")
	}
	if x&ItemFlagPetition != 0 {
		s = append(s, "Petition")
	}
	if x&ItemFlagHasText != 0 {
		s = append(s, "HasText")
	}
	if x&ItemFlagNoDisenchant != 0 {
		s = append(s, "NoDisenchant")
	}
	if x&ItemFlagRealDuration != 0 {
		s = append(s, "RealDuration")
	}
	if x&ItemFlagNoCreator != 0 {
		s = append(s, "NoCreator")
	}
	if x&ItemFlagIsProspectable != 0 {
		s = append(s, "IsProspectable")
	}
	if x&ItemFlagUniqueEquippable != 0 {
		s = append(s, "UniqueEquippable")
	}
	if x&ItemFlagIgnoreForAuras != 0 {
		s = append(s, "IgnoreForAuras")
	}
	if x&ItemFlagIgnoreDefaultArenaRestrictions != 0 {
		s = append(s, "IgnoreDefaultArenaRestrictions")
	}
	if x&ItemFlagNoDurabilityLoss != 0 {
		s = append(s, "NoDurabilityLoss")
	}
	if x&ItemFlagUseWhenShapeshifted != 0 {
		s = append(s, "UseWhenShapeshifted")
	}
	if x&ItemFlagHasQuestGlow != 0 {
		s = append(s, "HasQuestGlow")
	}
	if x&ItemFlagHideUnusableRecipe != 0 {
		s = append(s, "HideUnusableRecipe")
	}
	if x&ItemFlagNotUseableInArena != 0 {
		s = append(s, "NotUseableInArena")
	}
	if x&ItemFlagBindsToAccount != 0 {
		s = append(s, "BindsToAccount")
	}
	if x&ItemFlagNoReagentCost != 0 {
		s = append(s, "NoReagentCost")
	}
	if x&ItemFlagIsMillable != 0 {
		s = append(s, "IsMillable")
	}
	if x&ItemFlagReportToGuildChat != 0 {
		s = append(s, "ReportToGuildChat")
	}
	if x&ItemFlagNoProgressiveLoot != 0 {
		s = append(s, "NoProgressiveLoot")
	}
	return strings.Join(s, "|")
}

func ParseItemFlag(str string) (ItemFlag, error) {
	if str == "" {
		return 0, nil
	}

	s := strings.Split(str, "|")

	flag := ItemFlag(0)

	for _, v := range s {
		iflag, err := ItemFlagFromString("ItemFlag" + v)
		if err != nil {
			return flag, err
		}

		flag |= iflag
	}

	return flag, nil
}

func DecodeItemFlagInteger(version vsn.Build, value uint64) (ItemFlag, error) {
	descriptor, ok := ItemFlagDescriptors[version]
	if !ok {
		return 0, fmt.Errorf("no descriptor found for version %d", version)
	}

	iflag := ItemFlag(0)

	for k, v := range descriptor {
		if value&v != 0 {
			iflag |= k
		}
	}

	return iflag, nil
}

func (iflg ItemFlag) Resolve(version vsn.Build) (uint64, error) {
	desc, ok := ItemFlagDescriptors[version]

	if !ok {
		return 0, fmt.Errorf("no descriptor found for version %d", version)
	}
	out := uint64(0)

	for code, flags := range desc {
		if iflg&code != 0 {
			out |= flags
		}
	}

	return out, nil
}

func (iflg ItemFlag) Encode(wr io.Writer, version vsn.Build) error {
	code, err := iflg.Resolve(version)
	if err != nil {
		panic(err)
	}
	var data [4]byte

	binary.LittleEndian.PutUint32(data[:], uint32(code))
	_, err = wr.Write(data[:])

	return err
}

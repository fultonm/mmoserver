// Code generated by "gcraft_stringer -type=Race"; DO NOT EDIT.

package packet

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[RACE_NONE-0]
	_ = x[RACE_HUMAN-1]
	_ = x[RACE_ORC-2]
	_ = x[RACE_DWARF-3]
	_ = x[RACE_NIGHTELF-4]
	_ = x[RACE_UNDEAD_PLAYER-5]
	_ = x[RACE_TAUREN-6]
	_ = x[RACE_GNOME-7]
	_ = x[RACE_TROLL-8]
	_ = x[RACE_GOBLIN-9]
	_ = x[RACE_BLOODELF-10]
	_ = x[RACE_DRAENEI-11]
	_ = x[RACE_FEL_ORC-12]
	_ = x[RACE_NAGA-13]
	_ = x[RACE_BROKEN-14]
	_ = x[RACE_SKELETON-15]
	_ = x[RACE_VRYKUL-16]
	_ = x[RACE_TUSKARR-17]
	_ = x[RACE_FOREST_TROLL-18]
	_ = x[RACE_TAUNKA-19]
	_ = x[RACE_NORTHREND_SKELETON-20]
	_ = x[RACE_ICE_TROLL-21]
	_ = x[RACE_WORGEN-22]
	_ = x[RACE_PANDAREN_NEUTRAL-24]
	_ = x[RACE_PANDAREN_ALLIANCE-25]
	_ = x[RACE_PANDAREN_HORDE-26]
	_ = x[RACE_NIGHTBORNE-27]
	_ = x[RACE_HIGHMOUNTAIN_TAUREN-28]
	_ = x[RACE_VOID_ELF-29]
	_ = x[RACE_LIGHTFORGED_DRAENEI-30]
	_ = x[RACE_ZANDALARI_TROLL-31]
	_ = x[RACE_KUL_TIRAN-32]
	_ = x[RACE_THIN_HUMAN-33]
	_ = x[RACE_DARK_IRON_DWARF-34]
	_ = x[RACE_MAGHAR_ORC-36]
}

const (
	_Race_name_0 = "RACE_NONERACE_HUMANRACE_ORCRACE_DWARFRACE_NIGHTELFRACE_UNDEAD_PLAYERRACE_TAURENRACE_GNOMERACE_TROLLRACE_GOBLINRACE_BLOODELFRACE_DRAENEIRACE_FEL_ORCRACE_NAGARACE_BROKENRACE_SKELETONRACE_VRYKULRACE_TUSKARRRACE_FOREST_TROLLRACE_TAUNKARACE_NORTHREND_SKELETONRACE_ICE_TROLLRACE_WORGEN"
	_Race_name_1 = "RACE_PANDAREN_NEUTRALRACE_PANDAREN_ALLIANCERACE_PANDAREN_HORDERACE_NIGHTBORNERACE_HIGHMOUNTAIN_TAURENRACE_VOID_ELFRACE_LIGHTFORGED_DRAENEIRACE_ZANDALARI_TROLLRACE_KUL_TIRANRACE_THIN_HUMANRACE_DARK_IRON_DWARF"
	_Race_name_2 = "RACE_MAGHAR_ORC"
)

var (
	_Race_index_0 = [...]uint16{0, 9, 19, 27, 37, 50, 68, 79, 89, 99, 110, 123, 135, 147, 156, 167, 180, 191, 203, 220, 231, 254, 268, 279}
	_Race_index_1 = [...]uint8{0, 21, 43, 62, 77, 101, 114, 138, 158, 172, 187, 207}
)

func (i Race) String() string {
	switch {
	case 0 <= i && i <= 22:
		return _Race_name_0[_Race_index_0[i]:_Race_index_0[i+1]]
	case 24 <= i && i <= 34:
		i -= 24
		return _Race_name_1[_Race_index_1[i]:_Race_index_1[i+1]]
	case i == 36:
		return _Race_name_2
	default:
		return "Race(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}

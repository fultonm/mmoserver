// Code generated by "gcraft_stringer -type=AuthType"; DO NOT EDIT.

package auth

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[AUTH_LOGON_CHALLENGE-0]
	_ = x[AUTH_LOGON_PROOF-1]
	_ = x[AUTH_RECONNECT_CHALLENGE-2]
	_ = x[AUTH_RECONNECT_PROOF-3]
	_ = x[REALM_LIST-16]
	_ = x[XFER_INITIATE-48]
	_ = x[XFER_DATA-49]
	_ = x[XFER_ACCEPT-50]
	_ = x[XFER_RESUME-51]
	_ = x[XFER_CANCEL-52]
}

const (
	_AuthType_name_0 = "AUTH_LOGON_CHALLENGEAUTH_LOGON_PROOFAUTH_RECONNECT_CHALLENGEAUTH_RECONNECT_PROOF"
	_AuthType_name_1 = "REALM_LIST"
	_AuthType_name_2 = "XFER_INITIATEXFER_DATAXFER_ACCEPTXFER_RESUMEXFER_CANCEL"
)

var (
	_AuthType_index_0 = [...]uint8{0, 20, 36, 60, 80}
	_AuthType_index_2 = [...]uint8{0, 13, 22, 33, 44, 55}
)

func (i AuthType) String() string {
	switch {
	case 0 <= i && i <= 3:
		return _AuthType_name_0[_AuthType_index_0[i]:_AuthType_index_0[i+1]]
	case i == 16:
		return _AuthType_name_1
	case 48 <= i && i <= 52:
		i -= 48
		return _AuthType_name_2[_AuthType_index_2[i]:_AuthType_index_2[i+1]]
	default:
		return "AuthType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}

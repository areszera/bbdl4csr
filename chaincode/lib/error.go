package lib

import "errors"

var (
	ErrArgsLength         = errors.New("argument length too shorty")
	ErrArgsEmpty          = errors.New("empty argument contained")
	ErrUserNotFound       = errors.New("user not found")
	ErrUserExists         = errors.New("user exists")
	ErrReviewerNotEnough  = errors.New("reviewer not enough")
	ErrPaperNotFound      = errors.New("paper not found")
	ErrPeerReviewNotFound = errors.New("peer review not found")
)
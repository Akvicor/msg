package status

import "errors"

var (
	SenderErrorNotSupported = errors.New("not supported")
	SenderErrorConnecting   = errors.New("connecting")
	SenderErrorDown         = errors.New("down")
	SenderErrorWrongTarget  = errors.New("wrong target")
	SenderErrorUnknown      = errors.New("unknown")
)

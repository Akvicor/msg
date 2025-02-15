package status

import "sync/atomic"

type Status interface {
	Status() SenderStatus
}

type SenderStatus int32

const (
	_ SenderStatus = iota
	SenderStatusNotSupported
	SenderStatusConnecting
	SenderStatusOK
	SenderStatusDown
)

func (s *SenderStatus) ToString() string {
	switch *s {
	case SenderStatusNotSupported:
		return "not supported"
	case SenderStatusConnecting:
		return "connecting"
	case SenderStatusOK:
		return "ok"
	case SenderStatusDown:
		return "down"
	default:
		return "unknown"
	}
}

func (s *SenderStatus) ToError() error {
	switch *s {
	case SenderStatusNotSupported:
		return SenderErrorNotSupported
	case SenderStatusConnecting:
		return SenderErrorConnecting
	case SenderStatusOK:
		return nil
	case SenderStatusDown:
		return SenderErrorDown
	default:
		return SenderErrorUnknown
	}
}

type AtomicSenderStatus struct {
	status atomic.Int32
}

func NewSenderStatus(stat SenderStatus) *AtomicSenderStatus {
	st := &AtomicSenderStatus{
		status: atomic.Int32{},
	}
	st.Store(stat)
	return st
}

func (s *AtomicSenderStatus) Load() SenderStatus {
	return SenderStatus(s.status.Load())
}

func (s *AtomicSenderStatus) Store(status SenderStatus) {
	s.status.Store(int32(status))
}

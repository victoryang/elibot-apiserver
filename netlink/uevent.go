package netlink

import (
	"bytes"

	Log "elibot-apiserver/log"
)

const (
	ADD      = "add"
	REMOVE   = "remove"

	SUBSYSTEM = "SUBSYSTEM"
)

type UEvent struct {
	Action		string					`json:"action"`
	KObject		string					`json:"kobject"`
	Env			map[string]string 		`json:"env"`
}

func parseUEventHeader(action string) string {
	switch action {
		case ADD, REMOVE:
			return action
		default:
			return ""
	}
}

func parseUEvent(raw []byte) *UEvent {
	fields := bytes.Split(raw, []byte{0x00})
	if len(fields) == 0 {
		Log.Error("wrong uevent format")
		return nil
	}

	headers := bytes.Split(fields[0], []byte("@"))
	if len(headers) != 2 {
		Log.Error("wrong uevent header")
		return nil
	}

	action := parseUEventHeader(string(headers[0]))
	if action == "" {
		Log.Error("do not support uevent action: ", action)
		return nil
	}

	e := &UEvent {
		Action:		action,
		KObject:	string(headers[1]),
		Env:		make(map[string]string),
	}

	for _, v := range fields[1:len(fields)-1] {
		env := bytes.Split(v, []byte("="))
		if len(env) != 2 {
			continue
		}
		e.Env[string(env[0])] = string(env[1])

		// parse stops at SUBSYSTEM
		if string(env[0]) == SUBSYSTEM {
			break
		}
	}

	if _, ok := e.Env[SUBSYSTEM]; !ok {
		return nil
	}

	return e
}
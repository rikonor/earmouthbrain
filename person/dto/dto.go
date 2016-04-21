package dto

type Message string
type MessageHandler func(Message)

func (msg *Message) String() string {
	return string(*msg)
}

func ByteSliceToMessage(bs []byte) Message {
	return Message(string(bs))
}

func StringToMessage(s string) Message {
	return Message(s)
}

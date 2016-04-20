package dto

type Message string
type MessageHandler func(Message)

func ByteSliceToMessage(bs []byte) Message {
	return Message(string(bs))
}

func StringToMessage(s string) Message {
	return Message(s)
}

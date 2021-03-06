package message

import (
	"encoding/json"
	"fmt"
	"time"
)

const Broadcast = ""

// ConvertJSON converts message array to JSON string.
func ConvertToJSON(msgs []*Message) ([]byte, error) {
	return json.Marshal(msgs)
}

// MessageBox is a message box which holds all messages in its drawers.
type MessageBox struct {
	Drawers map[string]*Drawer
}

// NewMessageBox creates a new MessageBox object with no Drawers.
func NewMessageBox() *MessageBox {
	m := new(MessageBox)
	m.Drawers = make(map[string]*Drawer)
	return m
}

func (m *MessageBox) Dump() ([]byte, error) {
	return json.Marshal(m.Drawers)
}

// Pickup take out all message in a Drawer.
func (m *MessageBox) Pickup(name string) []*Message {
	drawer, ok := m.Drawers[name]
	if !ok {
		return nil
	}

	messages := drawer.Messages
	drawer.truncate()
	return messages
}

// Post puts message into correct Drawers.
func (m *MessageBox) Post(msg *Message) error {
	if msg == nil {
		return fmt.Errorf("Post message is nil.")
	}

	if msg.To == Broadcast {
		for name, drawer := range m.Drawers {
			if name != msg.From {
				drawer.appendMessage(msg)
			}
		}
	} else {
		m.addDrawer(msg.To)
		m.Drawers[msg.To].appendMessage(msg)
	}
	return nil
}

func (m *MessageBox) addDrawer(name string) {
	if _, exists := m.Drawers[name]; !exists {
		m.Drawers[name] = NewDrawer()
	}
}

// Drawer is a drawer of message box.
// Every drawer holds messages for someone.
type Drawer struct {
	Messages []*Message `json:"messages"`
}

const initialCapacity = 10

// NewDrawer creates a new Drawer object with no Messages.
func NewDrawer() *Drawer {
	d := new(Drawer)
	d.truncate()
	return d
}

func (d *Drawer) appendMessage(msg *Message) {
	d.Messages = append(d.Messages, msg)
}

func (d *Drawer) truncate() {
	d.Messages = make([]*Message, 0, initialCapacity)
}

// Message is a message sent from someone to someone.
// If value of To is "all", it is a broadcast message.
type Message struct {
	From      string `json:"from"`
	To        string `json:"to"`
	Body      string `json:"body"`
	Timestamp string `json:"timestamp"`
}

// New creates a new Message object.
func New(from, to, body string) *Message {
	m := new(Message)
	m.From = from
	m.To = to
	m.Body = body
	m.Timestamp = time.Now().Format("2006/01/02 15:04:05")
	return m
}

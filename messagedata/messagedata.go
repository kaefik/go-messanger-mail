package messagedata

// структура сообщения
type Message struct {
	id          int      // идентификатор сообщения
	idSender    int      // идентификатор отправитель
	idRecipient int      // идентификатор получателя
	txt         string   // содержание сообщения
	attach      []string // аттач к сообщению
}

func (this *Message) getId() int {
	return this.id
}

func (this *Message) setId(idi int) {
	this.id = idi
}

func (this *Message) getIdSender() int {
	return this.idSender
}

func (this *Message) setIdSender(idi int) {
	this.idSender = idi
}

func (this *Message) getIdRecipient() int {
	return this.idRecipient
}

func (this *Message) setIdRecipient(idi int) {
	this.idRecipient = idi
}

func (this *Message) getTxt() string {
	return this.txt
}

func (this *Message) setTxt(msg string) {
	this.txt = msg
}

func (this *Message) getAttach() []string {
	return this.attach
}

func (this *Message) setAttach(attach []string) {
	this.attach = attach
}

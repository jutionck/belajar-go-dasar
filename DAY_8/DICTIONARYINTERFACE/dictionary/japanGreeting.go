package dictionary

//JapanDict is Model of English Dictionary
type JapanDict struct{}

//GetMorningGreeting is A Method Receiver
func (en JapanDict) GetMorningGreeting() string {
	return "おはようございまず"
}

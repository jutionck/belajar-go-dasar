package dictionary

//EnglishDict is Model of English Dictionary
type EnglishDict struct{}

//GetMorningGreeting is A Method Receiver
func (en EnglishDict) GetMorningGreeting() string {
	return "Hello, good morning"
}

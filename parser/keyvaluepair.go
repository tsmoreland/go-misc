package parser

type KeyValuePair struct {
	key   string
	value string
}

var (
	EmptyKeyValuePair = KeyValuePair{"", ""}
)

func NewKeyValuePair(key, value string) KeyValuePair {
	return KeyValuePair{key, value}
}

func (k *KeyValuePair) Key() string {
	return k.key
}

func (k *KeyValuePair) Value() string {
	return k.value
}

func (k *KeyValuePair) Deconstruct() (string, string) {
	return k.key, k.value
}

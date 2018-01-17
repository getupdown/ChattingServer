package InfoParser

import "encoding/json"

type JsonParser struct {
	Parser
}

func (parser JsonParser) Parse(content []byte, v interface{}) (err error) {
	err = json.Unmarshal(content, v)
	return err
}


func (parser JsonParser) Encode(v interface{}) (content []byte, err error) {
	content, err = json.Marshal(v)
	return content, err
}

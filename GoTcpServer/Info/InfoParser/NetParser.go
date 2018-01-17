package InfoParser



type Parser interface {
	//parse the information into the map[string]string
	Parse(content []byte, v interface{}) (err error)
	Encode(v interface{}) (content []byte, err error)
}



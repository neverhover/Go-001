package logic

import "fmt"

func Handle(infos []byte) ([]byte, error) {
	str := string(infos)
	var (
		act      string
		content string
	)
	_, err := fmt.Sscanf(str, "%s %s", &act, &content)
	if err != nil {
		fmt.Printf("Sscanf error %s\n", err)
		return nil, err
	}
	switch act {
	case "auth":
		return []byte(Auth(content)), nil
	case "read":
		return Read(content)
	case "exec":
		return Exec(content)
	}
	return nil, nil
}

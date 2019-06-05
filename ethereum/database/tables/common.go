package tables

import "strings"

func MysqlFormat(s string) string {
	if !strings.ContainsAny(s, `;'\"&<>`) {
		return s
	}
	s = strings.Replace(s, `\`, `\\`, -1)
	s = strings.Replace(s, `'`, `\'`, -1)
	s = strings.Replace(s, `;`, `\;`, -1)
	s = strings.Replace(s, `"`, `\"`, -1)
	//s = strings.Replace(s, `&`, `\&`, -1)
	//s = strings.Replace(s, `<`, `\<`, -1)
	//s = strings.Replace(s, `>`, `\>`, -1)
	return s
}

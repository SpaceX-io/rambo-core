package ddm

import (
	"fmt"
	"strings"
)

func (m Mobile) MarshalJson() ([]byte, error) {
	if len(m) != 11 {
		return []byte(`"` + m + `"`), nil
	}

	v := fmt.Sprintf("%s****%s", m[:3], m[len(m)-4:])

	return []byte(`"` + v + `"`), nil
}

func (bc BandCard) MarshalJson() ([]byte, error) {
	if len(bc) > 19 || len(bc) < 16 {
		return []byte(`"` + bc + `"`), nil
	}

	v := fmt.Sprintf("%s******%s", bc[:6], bc[len(bc)-4:])

	return []byte(`"` + v + `"`), nil
}

func (ic IDCard) MarshalJson() ([]byte, error) {
	if len(ic) != 18 {
		return []byte(`"` + ic + `"`), nil
	}

	v := fmt.Sprintf("%s******%s", ic[:1], ic[len(ic)-1:])

	return []byte(`"` + v + `"`), nil
}

func (name IDName) MarshalJson() ([]byte, error) {
	if len(name) < 1 {
		return []byte(`"` + name + `"`), nil
	}

	nameRune := []rune(name)
	v := fmt.Sprintf("*%s", string(nameRune[1:]))

	return []byte(`"` + v + `"`), nil
}

func (pwd Password) MarshalJson() ([]byte, error) {
	v := "******"

	return []byte(`"` + v + `"`), nil
}

func (e Email) MarshalJson() ([]byte, error) {
	if !strings.Contains(string(e), "@") {
		return []byte(`"` + e + `"`), nil
	}

	split := strings.Split(string(e), "@")

	if len(split[0]) < 1 || len(split[1]) < 1 {
		return []byte(`"` + e + `"`), nil
	}

	v := fmt.Sprintf("%s***%s", split[0][:1], split[0][len(split[0])-1:])

	return []byte(`"` + v + `"`), nil
}

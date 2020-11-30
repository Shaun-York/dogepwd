package characters

import (
	"strings"
)

const (
	//Setatozl all lower-case letters
	Setatozl = "abcdefghijklmnopqrstuvwxyz"
	//Setatozu all upper-case letters
	Setatozu = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	//Setdigits all digits
	Setdigits = "0123456789"
	//Setspecial all special
	Setspecial = "{}:\"<>?[];\\',./~!@#$%^&*()_+`-=|"
)

//Charset Charset
type Charset struct {
	Chars string
}
//Atozl Atozl
func (c *Charset) Atozl() {
	if strings.Contains(c.Chars, Setatozl) {
		c.Chars = strings.Trim(c.Chars, Setatozl)
	} else {
		s := strings.Builder{}
		s.WriteString(c.Chars)
		s.WriteString(Setatozl)
		c.Chars = s.String()
	}
}
//Atozu Atozu
func (c *Charset) Atozu() {
	if strings.Contains(c.Chars, Setatozu) {
		c.Chars = strings.Trim(c.Chars, Setatozu)
	} else {
		s := strings.Builder{}
		s.WriteString(c.Chars)
		s.WriteString(Setatozu)
		c.Chars = s.String()
	}
}
//Special use special
func (c *Charset) Special() {
	if strings.Contains(c.Chars, Setspecial) {
		c.Chars = strings.Trim(c.Chars, Setspecial)
	} else {
		s := strings.Builder{}
		s.WriteString(c.Chars)
		s.WriteString(Setspecial)
		c.Chars = s.String()
	}
}
//Digits use digits
func (c *Charset) Digits() {
	if strings.Contains(c.Chars, Setdigits) {
		c.Chars = strings.Trim(c.Chars, Setdigits)
	} else {
		s := strings.Builder{}
		s.WriteString(c.Chars)
		s.WriteString(Setdigits)
		c.Chars = s.String()
	}
}
//Clear clear set
func (c *Charset) Clear() {
	c.Chars = ""
}
//All - Use all characters
func (c *Charset) All() {
	s := strings.Builder{}
	s.WriteString(Setatozl)
	s.WriteString(Setatozu)
	s.WriteString(Setspecial)
	s.WriteString(Setdigits)
	c.Chars = s.String()
}
// NotSet is set
func (c *Charset) NotSet() bool {
	return len(c.Chars) == 0
}
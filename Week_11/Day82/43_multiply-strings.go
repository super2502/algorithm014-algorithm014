package Day82

import "strconv"

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	ret := ""
	for i := 0; i < len(num1); i++ {
		ret = merge10(ret, oneProduct(num1[i], num2))
	}
	return ret
}

func merge10(s1 string, s2 string) string {
	if s1 != "" {
		s1 = s1 + "0"
	}
	return sum(s1, s2)
}

func oneProduct(a byte, b string) string {
	ret := ""
	for j := 0; j < len(b); j++ {
		ret = merge10b(ret, a, b[j])
	}
	return ret
}

func merge10b(s string, a, b byte) string {
	if s != "" {
		s = s + "0"
	}
	prod := int(a-'0') * int(b-'0')
	return sum(s, strconv.Itoa(prod))
}

func sum(s1, s2 string) string {
	s := ""
	carry := 0
	i, j := len(s1)-1, len(s2)-1
	for i >= 0 && j >= 0 {
		sum := int(s1[i]-'0') + int(s2[j]-'0') + carry
		carry = sum / 10
		s = string(byte(sum%10+'0')) + s
		i--
		j--
	}
	for i >= 0 {
		sum := int(s1[i]-'0') + carry
		carry = sum / 10
		s = string(byte(sum%10+'0')) + s
		i--
	}
	for j >= 0 {
		sum := int(s2[j]-'0') + carry
		carry = sum / 10
		s = string(byte(sum%10+'0')) + s
		j--
	}
	if carry != 0 {
		s = "1" + s
	}
	return s
}

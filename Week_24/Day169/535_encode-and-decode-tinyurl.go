package Day169

import (
	"encoding/base64"
	"fmt"
	"strings"
)

type Codec struct {
	prefix string
}

func Constructor() Codec {
	return Codec{
		prefix: "http://tinyurl.com/",
	}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	uri := base64.StdEncoding.EncodeToString([]byte(longUrl))
	return fmt.Sprintf("%s%s", this.prefix, uri)
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	uri := strings.TrimPrefix(shortUrl, this.prefix)
	bs, _ := base64.StdEncoding.DecodeString(uri)
	return string(bs)
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */

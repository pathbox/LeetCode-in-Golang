package LeetCode535

type Codec struct {
	cache map[string]string
	count int
}

func Constructor() Codec {
	return Codec{
		cache: make(map[string]string, 0),
		count: 0,
	}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	if _, ok := this.cache[longUrl]; ok {
		return this.cache[longUrl]
	}
	this.count++
	str := this.trans()
	this.cache[longUrl] = str
	return str
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	for k, v := range this.cache {
		if v == shortUrl {
			return k
		}
	}
	return ""
}

func (this *Codec) trans() string {
	chars := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	tmp := this.count
	var res []byte
	for tmp > 0 {
		b := tmp % 62
		res = append(res, chars[b])
		tmp = tmp / 62
	}
	return string(res)
}

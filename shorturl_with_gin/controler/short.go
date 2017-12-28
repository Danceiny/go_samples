package controler

import (
	"crypto/md5"
	"fmt"
	"io"
	"samples/shorturl_with_gin/common/cache"
)
var (
	urlcache cache.Cache
	globalnum int
)


func init() {
	globalnum = 100000000
	urlcache, _ = cache.NewCache("memory", `{"interval":0}`)
}

type ShortResult struct {
	UrlShort string
	UrlLong  string
}
func GetMD5(lurl string) string {
	h := md5.New()
	salt1 := "salt4shorturl"
	io.WriteString(h, lurl+salt1)
	urlmd5 := fmt.Sprintf("%x", h.Sum(nil))
	return urlmd5
}
func ShortenControler(l string)(result ShortResult){
	result.UrlLong = l
	urlmd5 := GetMD5(l)

	println(urlmd5)

	if urlcache.IsExist(urlmd5) {
		result.UrlShort = urlcache.Get(urlmd5).(string)
		println(result.UrlShort)
	} else {
		result.UrlShort = Generate()
		println(result.UrlShort)
		err := urlcache.Put(urlmd5, result.UrlShort, 0)
		if err != nil {
			//beego.Info(err)
		}
		err = urlcache.Put(result.UrlShort, l, 0)
		if err != nil {
			//beego.Info(err)
		}
	}



	return
}

func Generate() (tiny string) {
	globalnum++
	num := globalnum
	fmt.Println(num)
	alpha := merge(getRange(48, 57), getRange(65, 90))
	alpha = merge(alpha, getRange(97, 122))
	if num < 62 {
		tiny = string(alpha[num])
		return tiny
	} else {
		var runes []rune
		runes = append(runes, alpha[num%62])
		num = num / 62
		for num >= 1 {
			if num < 62 {
				runes = append(runes, alpha[num-1])
			} else {
				runes = append(runes, alpha[num%62])
			}
			num = num / 62

		}
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		tiny = string(runes)
		return tiny
	}
	return tiny
}

func getRange(start, end rune) (ran []rune) {
	for i := start; i <= end; i++ {
		ran = append(ran, i)
	}
	return ran
}

func merge(a, b []rune) []rune {
	c := make([]rune, len(a)+len(b))
	copy(c, a)
	copy(c[len(a):], b)
	return c
}

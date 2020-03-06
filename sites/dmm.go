package sites

import "fmt"

func Dmm(id string) Site {
	return Site{
		Url:       url(id),
		UserAgent: MobileUserAgent,
		Title:     ".ttl-grp",
	}
}

func url(id string) string {
	return fmt.Sprintf("https://www.dmm.co.jp/mono/dvd/-/detail/=/cid=%s/", id)
}

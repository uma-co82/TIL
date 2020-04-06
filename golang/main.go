package main

import "fmt"

func main() {
	hoge := []string{"1", "2", "3"}
	var sum int
	if len(hoge) < sum {
		goto END_LABEL
	}
	for i := 0; i <= 2; i++ {
		fmt.Println(hoge[i])
	}
END_LABEL:
	fmt.Println("end")
}

// func main() {
// 	re := regexp.MustCompile(`\d+`)
// 	// ho := "/korea-cosmetic6【韓国コスメ♡】韓国旅行で買うべき６の理由！ | Stayway Media"
// 	// ff := "/osaka【2019年最新版】意外な穴場もあり？関西人が選ぶ、大阪でおすすめの観光スポットベスト30 | Stayway Media"
// 	// foo := "/hawaii-clothes10女子必読！ハワイ旅行の服装で留意すべき１０か条！ | Stayway Media"
// 	// tt := "/osaka-travel22019年【ひとり旅必見】大阪でおすすめの観光スポット10選！ | Stayway Media"
// 	foo := "/singapore-restaurant25/3シンガポールでグルメを楽しもう！おすすめレストラン25選 | ページ 3 | Stayway Media ページ 4"
// 	ree := regexp.MustCompile(`ページ (\d)`)
// 	fmt.Println(ree.FindAllStringSubmatch(foo, 1))
// 	tt := "/hawaii-clothes10女子必読！ハワイ旅行の服装で留意すべき１０か条！ | Stayway Media"
// 	fmt.Println(ree.FindAllStringSubmatch(tt, 1))
// 	// fmt.Println(strings.Split(ho, "/"))
// 	// fmt.Println(strings.Split(ff, "/"))
// 	// fmt.Printf("Run: %d", utf8.RuneCountInString(ho))

// 	var ss string
// 	var sss string
// 	// 1. 日本語を発見した段階でそれ以降全て削除
// 	r := []rune(tt)
// 	for _, v := range r {
// 		b := []byte(string(v))
// 		if len(b) >= 3 {
// 			ss = tt[0:strings.Index(tt, string(v))]
// 			break
// 		}
// 	}

// 	// 2. xxxx年などを削除するために数字を全部抜き出し、4文字以上だったら末4文字を削除する
// 	for _, str := range re.FindAllStringSubmatch(ss, -1) {
// 		for _, st := range str {
// 			if utf8.RuneCountInString(st) >= 4 {
// 				sss = strings.Replace(ss, st[len(st)-4:], "", -1)
// 			}
// 		}
// 	}

// 	fmt.Println(ss)
// 	fmt.Println(sss)
// }

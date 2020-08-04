package cmd

import (
	"fmt"
	"net/http"
	"net/url"
)

// jsonをパースする為の構造体を定義する

type Data struct {
	ID             string `json:id`
	Url            string `json:"url"`
	Title          string `json:"title"`
	LikesCount     int    `json:"likes_count"`
	PageViewsCount int    `json:"page_views_count"`
}

func FetchQiitaData(accessToken string) []Data {

	baseUrl := "https://qiita.com/api/v2/items?query=user:mizuki0414"
	// 様々な検索条件をかけるときはbaseUrlをv2/までにして他を変数で定義してurl.Parseで合体させる
	endpointURL, err := url.Parse(baseUrl)
	_ = err
	var resp = &http.Response{}
	// qiitaのアクセストークンがない場合はAuthorizationを付与しない
	// 2パターン作っておく。
	// accessトークンは環境変数に入れておく。自分の場合は.bash_profileにexport文を書いている。

	if len(accessToken) > 0 {
		// QiitaAPIにリクエストを送ってレスポンスをrespに格納する。
		resp, err = http.DefaultClient.Do(&http.Request{
			URL:    endpointURL,
			Method: "GET",
			Header: http.Header{
				"Content-Type":  {"application/json"},
				"Authorization": {"Bearer " + accessToken},
			},
		})
	} else {
		fmt.Println("***** Access Token 無しでQiitaAPIを叩いています アクセス制限に注意して下さい*****")

		resp, err = http.DefaultClient.Do(&http.Request{
			URL:    endpointURL,
			Method: "GET",
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
		})
	}
	defer resp.Body.Close()

	var datas []Data

	/*********一覧取得では、ページビューがnilになるので個別で取りに行ってデータを得る*****************/
	for i, val := range datas {

		article_id := val.ID
		baseUrl := "https://qiita.com/api/v2/items/"
		endpointURL2, err := url.Parse(baseUrl + article_id)
		_ = err
		resp, err = http.DefaultClient.Do(&http.Request{
			URL:    endpointURL2,
			Method: "GET",
			Header: http.Header{
				"Content-Type":  {"application/json"},
				"Authorization": {"Bearer " + accessToken},
			},
		})

		var m map[string]interface{}

		datas[i].PageViewsCount = int(m["page_views_count"].(float64))
	}
	return datas
}

// データの出力
func OutputQiitaData(datas []Data) {
	fmt.Println("************************自分のQiita投稿一覧******************************")
	for _, val := range datas {
		fmt.Printf("%-15v%v%v\n", "ID", ": ", val.ID)
		fmt.Printf("%-15v%v%v\n", "Title", ": ", val.Title)
		fmt.Printf("%-12v%v%v\n", "いいね", ": ", val.LikesCount)
		fmt.Printf("%-9v%v%v\n", "ページビュー", ": ", val.PageViewsCount)
		fmt.Printf("%-15v%v%v\n", "URL", ": ", val.Url)
		fmt.Println("-------------------------------------------------------------------------")
	}
}

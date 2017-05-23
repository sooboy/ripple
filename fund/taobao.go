package fund

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	db "ripple/database"
)

var (
	// Client 磨人Client
	Client = new(http.Client)
)

var (
	url = "https://detailskip.taobao.com/service/getData/1/p1/item/detail/sib.htm?itemId=%s&sellerId=295065471&modules=soldQuantity,price"
	ids = []string{
		"541820182958",
		"541881604174",
		"541830335367",
		"544410378063",
	}
)

type Data struct {
	Code struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"code"`
	Data struct {
		Price        string `json:"price"`
		SoldQuantity struct {
			ConfirmGoodsCount float64 `json:"confirmGoodsCount"`
			SoldTotalCount    float64 `json:"soldTotalCount"`
		} `json:"soldQuantity"`
	} `json:"data"`
}
type Result struct {
	Time string
	Fund float64
}

func (d *Data) Result() (*Result, error) {
	f, err := strconv.ParseFloat(d.Data.Price, 64)
	if err != nil {
		return nil, err
	}
	return &Result{
		Time: time.Now().Format("2006-01-02 15:04:05"),
		Fund: f * d.Data.SoldQuantity.SoldTotalCount,
	}, nil
}

func TaobaoFundIn() {
	var counter = 0.0
	for _, item := range ids {
		fmt.Println(fmt.Sprintf(url, item))
		var req, err = http.NewRequest("GET", fmt.Sprintf(url, item), nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		var h = http.Header{}
		h.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36")
		h.Add("referer", "https://item.taobao.com/item.htm")
		req.Header = h

		res, err := Client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(body))
		var data Data
		err = json.Unmarshal(body, &data)
		if err != nil {
			fmt.Println(err)
			return
		}
		f, err := data.Result()
		if err != nil {
			fmt.Println(err)
			return
		}
		counter += f.Fund
	}
	fmt.Println(counter)

	err := db.Session.DB("test").C("ripple.fund").Insert(&Result{
		Time: time.Now().Format("2006-01-02 15:04:05"),
		Fund: counter,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("store Ok")
}

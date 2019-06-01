package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type InitImageUrl interface {
	InitClear()
}

type Shop struct {
	Id string		`json:"id"`
	Name string		`json:"name"`
}

type Product struct {
	Name string		`json:"name"`
	Brand string	`json:"brand_id"`
	Price int		`json:"price"`
	Shops []Shop	`json:"shops"`
	Cover string	`json:"cover"`
	Banners []string	`json:"banner_images"`
	Details []string	`json:"details"`
}

type Brand struct {
	Id string		`json:"id"`
	Name string 	`json:"name"`
}

type Brands struct {
	Data []Brand
}

type ListItem struct {
	Url string
	Img string
	Title string
}

func (item *ListItem) InitClear() {
	item.Url = WEBSITE_ROOT + strings.Trim(item.Url, " ")
	// item.Img = UploadImage("https://" + strings.Trim(item.Img, " "))
	item.Img = "https://" + strings.Trim(item.Img, " ")
	item.Title = strings.Trim(item.Title, " ")
}

type Images struct {
	Urls []string
}

func (img *Images) InitClear ()  {
	var newUrls []string
	for _, v := range img.Urls {
		newUrls = append(newUrls, "https://" + strings.Trim(v, " "))
		time.Sleep(time.Microsecond)
	}
	img.Urls = newUrls
}

func (product *Product) InitClear ()  {
	product.Name = func() string {
		name := []rune(product.Name)
		return string(name[0: 200])
	}()
	product.Brand = GetRandBrand().Id
	product.Shops = GetRandShops()
}

var (
	BaseApi = ""
	AuthToken = ""
	WEBSITE_ROOT = "https://www.meilele.com"
)

func main()  {
	var pageIndex = 1
	for pageIndex < 48 {
		AddProducts(pageIndex)
		pageIndex++
	}
}

/**
添加商品
*/
func AddProducts(pageIndex int)  {
	var uri = "/weapi/v1/product"
	body, err := DownloadHtml("https://www.meilele.com/category-keting/list-p" + strconv.Itoa(pageIndex) + "/?from=page")
	if err != nil {
		fmt.Println(err)
		return
	}
	html := string(body)
	html = RegexpReplaceEmpty(html, "\\s{2,}")
	re := regexp.MustCompile(`<li data\-index="[0-9]+" data\-goods\-id="[0-9]+" class="ga\-list g\-item"[^>]*>.*?<\/li>`)
	res := re.FindAllString(html, -1)
	for _, match := range res {
		item, e := Parse(match)
		if e != nil {
			continue
		}
		product, e2 := ParseItemDetail(item)
		if e2 != nil {
			continue
		}
		reader, err := json.Marshal(product)
		if err != nil {
			fmt.Println("json err.   ", reader)
			return
		}
		fmt.Println(item.Title)
		// 提交到服务器
		SendPostRequest(uri, product)
	}
}

/**
	解析单个Item
 */
func Parse(html string) (ListItem ,error) {
	var res []string
	re1 := regexp.MustCompile(`<a href="(\/category\-shafa\/goods\-[0-9]*\.html\?page=[0-9]+&index=[0-9]+)" target="_blank">.*<img class="d\-img" src="//.*"[^>]*data\-src="//(.*?)"[^>]*data\-wide\-src="//.*"[^>]*alt="(.*?)图片"[^>]*><\/a>`)
	findres := re1.FindString(html)
	if findres == ""{
		re2 := regexp.MustCompile(`<a href="(\/category\-shafa\/goods\-[0-9]*\.html\?page=[0-9]+&index=[0-9]+)" target="_blank">.*<img class="d\-img" src="//(.*?)"[^>]*alt="(.*?)图片"[^>]*><\/a>`)
		res = re2.FindStringSubmatch(html)
	} else {
		res = re1.FindStringSubmatch(html)
	}
	if len(res) >= 3 {
		item := ListItem{
			Url: res[1],
			Img: res[2],
			Title: res[3],
		}
		item.InitClear()
		return item, nil
	}

	return ListItem{}, errors.New("not found.")
}


/**
	解析Item详情
 */
func ParseItemDetail(item ListItem) (Product, error) {
	product := Product{
		Name: item.Title,
		Cover: item.Img,
	}

	html, _ := DownloadHtml(item.Url)
	priceRe := regexp.MustCompile(`<strong id="JS_effect_price" class="gnum">([0-9]+)\.00<\/strong>`)
	prices := priceRe.FindStringSubmatch(html)
	price, _ := strconv.Atoi(prices[1])
	product.Price = price

	imagesRe := regexp.MustCompile(`<div class="img" data\-source="handledpictures" style="margin\-bottom:[0-9]+px;">[^<]*<img src="//image\.meilele\.com\/images\/blank\.gif"[^>]*data\-src="//(.+\.(jpg|png|bum|gif|jpeg|ico))"`)
	images := imagesRe.FindAllStringSubmatch(html, -1)
	imgs := Images{}
	for _, image := range images {
		imgs.Urls = append(imgs.Urls, image[1])
	}
	imgs.InitClear()

	if len(imgs.Urls) >= 3 {
		product.Banners = imgs.Urls[:3]
		if len(imgs.Urls) > 20 {
			product.Details = imgs.Urls[3:20]
		} else {
			product.Details = imgs.Urls[3:]
		}
		product.InitClear()

		return product, nil
	}

	return Product{}, errors.New("not found.")
}

/**
	随机获取一个品牌
 */
func GetRandBrand() Brand {
	// 这里有个坑 Golang的相对路径是相对于命令执行的路径 不是可执行文件的路径!!!
	// 读文件时要写绝对路径
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	jsonFile := dir + "/pack28/brands.json"
	// 解析品牌json
	res, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		return Brand{}
	}
	brands := Brands{}
	_ = json.Unmarshal(res, &brands)
	allBrands := make(map[int]Brand)
	for i, v := range brands.Data {
		allBrands[i] = v
	}
	Length := len(allBrands) - 1
	// 这也是个坑 Golang 随机数在程序重启后生成的随机数跟上次运行结果一样  需要重新播随机种子
	// 这里使用Unix时间戳
	rand.Seed(time.Now().UnixNano())

	randIndex := rand.Intn(Length)
	return allBrands[randIndex]
}

/**
	随机获取一个门店列表
 */
func GetRandShops() []Shop {
	AllShops := []Shop{
		{Id: "a9f1f02e", Name: "徐东销品茂店"},
		{Id: "c3ff61a4", Name: "武汉广场店"},
		{Id: "d37e9f1c", Name: "武昌区天天开心店"},
		{Id: "8cabc921", Name: "默认门店"},
	}
	rand.Seed(time.Now().UnixNano())
	Length := len(AllShops)
	doRand := func () Shop {
		randIndex := rand.Intn(Length)
		return AllShops[randIndex]
	}
	var result []Shop
	for i := 0;i < 2 ; i++ {
		result = append(result, doRand())
	}
	return result
}

/**
	发送POST请求
 */
func SendPostRequest(uri string, data interface{})  {
	reader, err := json.Marshal(data)
	if err != nil {
		fmt.Println("json err.   ", reader)
		return
	}
	url := BaseApi+ uri
	request, _ := http.NewRequest("POST", url, bytes.NewReader(reader))
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	request.Header.Set("Authorization", AuthToken)
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	jsons, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsons))
}

/**
	上传图片
 */
func UploadImage(uri string) string  {
	var (
		base = "https://www.superbed.cn/upload?token=3a2be755bfe849acb8ce491de7217414&sync=1&endpoints=alicdn&v=1&src="
	)
	url := base + uri
	request, _ := http.NewRequest("POST", url, bytes.NewReader([]byte("")))
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if response.StatusCode != http.StatusOK {
		fmt.Println("status code => ", response.StatusCode)
		return ""
	}
	jsons := string(body)
	jsre := regexp.MustCompile(`https://.+.alicdn.com/.+.(jpg|png|jpeg|gif|bmp)`)
	return jsre.FindString(jsons)
}

/**
	正则替换空
 */
func RegexpReplaceEmpty(html string, regex string) string {
	re, _ := regexp.Compile(regex)
	return re.ReplaceAllString(html, "")
}

/**
	下载网页
 */
func DownloadHtml(uri string) (string, error)  {
	response, err := http.Get(uri)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return "", errors.New("status code ===> " + strconv.Itoa(response.StatusCode))
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

/**
下载远程文件保存到本地
*/
func DownloadImageSave(url string) string  {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	imgDir := dir + "/pack28/img/"
	response, err2 := http.Get(url)
	if err2 != nil {
		panic(err2)
	}
	defer response.Body.Close()
	body, err3 := ioutil.ReadAll(response.Body)
	if err3 != nil {
		panic(err3)
	}
	reg, _ := regexp.Compile(`(\w|\d|_)*.(jpg|png|bum|gif|jpeg|ico)`)
	name := reg.FindStringSubmatch(url)[0]
	imgPath := imgDir + name
	out, err4 := os.Create(imgPath)
	if err4 != nil {
		panic(err4)
	}
	_, _ = io.Copy(out, bytes.NewReader(body))
	return imgPath
}

/**
添加商品品牌
无需重复调用
*/
func AddBeands()  {
	var (
		uri = "/weapi/v1/brands"
	)
	all, err := DownloadHtml("https://www.meilele.com/category-keting/")
	if err != nil {
		panic(err)
	}
	re := regexp.MustCompile(`<a href="\/category-keting\/list-b[0-9]+\/\?from=b" title="(.*)" class="abs"><span class="txt"[^>]*>[^<]+<\/span><\/a>`)
	res := re.FindAllSubmatch([]byte(all), -1)
	var brands []string
	for _, match := range res {
		brands = append(brands, string(match[1]))
	}
	for _, brand := range brands {
		data := make(map[string]interface{})
		data["name"] = brand
		fmt.Println("Add ===> ", brand)
		SendPostRequest(uri, data)
	}
}

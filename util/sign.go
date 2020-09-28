package util

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
)

const MERICO = "https://cloud.merico.cn"
const CONTENTTYPE = "application/json"


//options选项
const (
	DEVNUM = "developer_num" //开发者数量
	DEVEQU = "dev_equivalent" //开发当量(规模)
   	DEVVAL = "dev_value"  //开发价值
 	DEVEQUEVERYDEVELOPER = "dev_equivalent_every_developer" //平均开发当量
 	DEVEQUGROUPPERID = "dev_equivalent_group_period" //按照指定周期的开发当量
	DEVEQUEVERYGROUP = "dev_equivalent_every_developer_group_period" //按照指定周期的平均开发当量
)

//签名类
type SignUtils struct {
	params map[string]interface{}
	key string
}

//选项类
type Options struct {
	SelectColumn string	`json:"selectColumn"`
	TargetTimezoneName string `json:"targetTimezoneName"`
	SelectProjectId string `json:"selectProjectId,omitempty"`
}

//实例化签名
func NewSign()(sign *SignUtils) {
	sign = new(SignUtils)
	sign.key = MericoSecret
	sign.params = make(map[string]interface{})
	sign.params["appid"] = MericoAppid
	return sign
}


func (s *SignUtils)SetNonceStr(nonceStr string){
	s.params["nonce_str"] = nonceStr
}

/*
对map的key排序，返回有序的slice
 */
func (s *SignUtils)sortMapbyKey()(keys []string){
	for k := range s.params{
		keys = append(keys,k)
	}
	sort.Strings(keys)
	return keys
}

/*
设置k-v键值对
 */
func (s *SignUtils)Set(key,value string){
	s.params[key] = value
}

/*
设置数组和对象k-v
 */
func (s *SignUtils)SetObjectOrArray(key string,value interface{})(err error){
	s.params[key] = value
	return
}

/*
生成sign
 */
func (s *SignUtils)sign()(result string){
	//对key排序
	keys := s.sortMapbyKey()
	//拼接
	for _,val := range keys{
		//对于array和object需要先转义再拼接
		v,err := s.params[val].(string)
		//断言失败ok为false
		if !err{
			//序列化
			if b,err2 := json.Marshal(s.params[val]);err2 ==nil{
				v = string(b)
			}else{
				fmt.Println(err2)
				return
			}
		}
		result += val+ "=" + v + "&"
	}
	result += "key=" + MericoSecret
	//md5加密
	result = fmt.Sprintf("%x", md5.Sum([]byte(result)))
	//转化大写
	result = strings.ToUpper(result)
	return
}

/*
获取要post的数据
 */
func (s *SignUtils)GetPostData()(result *strings.Reader,err error){
	//获取sign值
	s.params["sign"] = s.sign()
	//序列化 两次序列化导致转义
	b,err := json.Marshal(s.params)
	if err != nil{
		return
	}
	result = strings.NewReader(string(b))
	return
}

func test(){
	signTool := NewSign()
	signTool.SetNonceStr("ibuaiVcKdpRxkhJ9A")

	options := Options{
		SelectColumn:       DEVVAL,
		TargetTimezoneName: "UTC",
		SelectProjectId: 	"9b8225e8-6ab3-4209-80d7-670f32e2f331",
	}
	err := signTool.SetObjectOrArray("options",options)
	if err != nil{
		fmt.Println(err)
	}

	res,err  := signTool.GetPostData()
	if err != nil{
		fmt.Println(err)
		return
	}

	url := MERICO+"/openapi/openapi/developer/get-efficiency-metric"
	fmt.Println(url)

	fmt.Println(*res)
	resp,err := http.Post(url,CONTENTTYPE,res)
	if err!= nil{
		fmt.Println(err)
	}
	defer resp.Body.Close()

	b,err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return
	}

	fmt.Println(string(b))
	fmt.Println(resp.Status)
}

func testDeveloper() {
	signTool := NewSign()
	signTool.SetNonceStr("ibuaiVcKdpRxkhJ9A")

	options := Options{
		SelectColumn:       DEVVAL,
		TargetTimezoneName: "UTC",
		SelectProjectId: 	"9b8225e8-6ab3-4209-80d7-670f32e2f331",
	}
	err := signTool.SetObjectOrArray("options",options)
	if err != nil{
		fmt.Println(err)
	}

	var primaryEmails []string
	primaryEmails = append(primaryEmails,"1005035266@qq.com")
	err = signTool.SetObjectOrArray("primaryEmailStrs",primaryEmails)
	if err != nil{
		fmt.Println(err)
	}

	res,err  := signTool.GetPostData()
	if err != nil{
		fmt.Println(err)
		return
	}

	url := MERICO+"/openapi/openapi/developer/get-efficiency-metric"
	fmt.Println(url)

	fmt.Println(*res)
	resp,err := http.Post(url,CONTENTTYPE,res)
	if err!= nil{
		fmt.Println(err)
	}
	defer resp.Body.Close()

	b,err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return
	}

	fmt.Println(string(b))
	fmt.Println(resp.Status)
}

func testProjectAdd(){
	signTool := NewSign()
	signTool.SetNonceStr("ibuaiVcKdpRxkhJ9A")
	signTool.Set("gitUrl","https://github.com/zby1218/zbygit.git")

	res,err  := signTool.GetPostData()
	if err != nil{
		fmt.Println(err)
		return
	}

	url := MERICO+"/openapi/openapi/project/add"
	fmt.Println(url)

	fmt.Println(*res)
	resp,err := http.Post(url,CONTENTTYPE,res)
	if err!= nil{
		fmt.Println(err)
	}
	defer resp.Body.Close()

	b,err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return
	}

	fmt.Println(string(b))
	fmt.Println(resp.Status)
}

func testGroupAdd(){
	signTool := NewSign()
	signTool.SetNonceStr("ibuaiVcKdpRxkhJ9A")
	signTool.Set("name","sky的测试name")
	signTool.Set("description","sky的测试description")
	signTool.Set("logo","https://www.baidu.com/link?url=q8pK1hrK1CCaVb9vBNq7n09Yj2-IY6TGTX6DmYyKSBVm5pf2bwBGlNOYFfOQp_Z9CE3Q7kdrRY8gWcFj0CjP1Y9uDb5Psc9EjQR_9T7wFnQ615wFGGNgdPA2i1vWQhqiyEtWE0POz_DGGBpDQJ5uPBIAtnAwFKNT-gXM7VYil-kwtIeiprB3iqp5_TZ56anAqPUVF0LSTDVQOCA0W2G59GltEI9rEXaZfQlwY4cvp6K-gNOHZIwa6brujqyU4cD3OctYuq1I9PaSxMWnbdAXX70YcftBbBBHQkFTw5bVWUKyrDaW2UZTLLCT0HewpY6o3NnYl4Hb3F7Y0kwYdUSBZfqMjdUyiA1Fxr1yq9-hK40YbhqmsWrs81h6XMoZTzBfVIuZHZlG1k-TjvUW1DT47EeLHTQesw_KQtCpyiiFYrEIZR6VJZM0oAKkWG0f5Vqvbvdcbnz-86wcZb01aqURTWDoRe0md621gCdJTXQeCKZbvZgP-078x8kdry0YbZzoJ9e-b7QQ5SeZYD_jbjZelH8m7RLt2FJxG1P7E7RX7BZpmIk3ZNgzVHglYMRtTFstRswLS2-oDuLeDlyYA0AlFRCTf54t34Yp1vyfV4m0q8q&timg=&click_t=1600950954494&s_info=1351_624&wd=&eqid=d23436200003f5f6000000065f6c92a6")

	res,err  := signTool.GetPostData()
	if err != nil{
		fmt.Println(err)
		return
	}

	url := MERICO+"/openapi/openapi/group/add"
	fmt.Println(url)

	fmt.Println(*res)
	resp,err := http.Post(url,CONTENTTYPE,res)
	if err!= nil{
		fmt.Println(err)
	}
	defer resp.Body.Close()

	b,err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return
	}

	fmt.Println(string(b))
	fmt.Println(resp.Status)
}

func testProjectGet() {
	signTool := NewSign()
	signTool.SetNonceStr("ibuaiVcKdpRxkhJ9A")
	//id或者gitUrl
	signTool.Set("gitUrl","https://github.com/g1050/test.git")

	res,err  := signTool.GetPostData()
	if err != nil{
		fmt.Println(err)
		return
	}

	url := MERICO+"/openapi/openapi/project/get"
	fmt.Println(url)

	fmt.Println(*res)
	resp,err := http.Post(url,CONTENTTYPE,res)
	if err!= nil{
		fmt.Println(err)
	}
	defer resp.Body.Close()

	b,err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return
	}

	fmt.Println(string(b))
	fmt.Println(resp.Status)
}

func testGroupfind() {
	signTool := NewSign()
	signTool.SetNonceStr("ibuaiVcKdpRxkhJ9A")

	res,err  := signTool.GetPostData()
	if err != nil{
		fmt.Println(err)
		return
	}

	url := MERICO+"/openapi/openapi/group/find"
	fmt.Println(url)

	fmt.Println(*res)
	resp,err := http.Post(url,CONTENTTYPE,res)
	if err!= nil{
		fmt.Println(err)
	}
	defer resp.Body.Close()

	b,err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return
	}

	fmt.Println(string(b))
	fmt.Println(resp.Status)
}

func test_main()  {

	// /openapi/developer/get-efficiency-metric
	//testDeveloper()

	/*
	project
	 */
	// /openapi/project/add gitUrl方式
	//testProjectAdd()

	// /openapi/project/get
	//testProjectGet()

	/*
	group
	 */
	// /openapi/group/add
	//testGroupAdd()

	// /openapi/group/find 查询team所有组
	//testGroupfind()
}
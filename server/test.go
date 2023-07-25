package main

import "log"

func Test() {
	log.Println(FindByWechatOpenID("x1xx", "xxx"))            // ""
	log.Println(FindByWechatOpenID("x1xx", "memos_open_api")) // ""
	log.Println(FindByWechatOpenID("xxx", "memos_open_api"))  // ""
	log.Println(FindByWechatOpenID("xxx", "memos"))

	log.Println(FindAOpenIDExist("xxx"))
	log.Println(FindAOpenIDExist("xxx1"))
	log.Println(FindAOpenIDExist("xxx2"))
}

func Test2() {
	memosOpenAPIUrl := FindByWechatOpenID("kkbt", "memos_open_api")
	err := CreateMemo(memosOpenAPIUrl, "This is a memo1", "PRIVATE", []int{})
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Memo created")
	}
	log.Println(memosOpenAPIUrl)
	id, err := CreateResourceByLink(memosOpenAPIUrl, "https://img-s-msn-com.akamaized.net/tenant/amp/entityid/AA1eey6U.img?w=640&h=360&m=6")
	log.Println(id, err)
	err = CreateMemo(memosOpenAPIUrl, "This is a memo2", "PRIVATE", []int{id})
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Memo created")
	}
}

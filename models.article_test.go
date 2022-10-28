package main

import "testing"

func TestGetAllAritcles(t *testing.T) {
	// Testing article data的数据读取，
	alist := getAllArticles()
	if len(alist) != len(articleList) {
		t.Fail()
	}

	for i, v := range alist {
		if v.Content != articleList[i].Content || v.ID != articleList[i].ID || v.Title != articleList[i].Title {
			t.Fail()
			break
		}
	}
}

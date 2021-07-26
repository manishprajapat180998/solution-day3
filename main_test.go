package main

import (
	"bytes"
	"fmt"
	"github.com/exercise2/Config"
	"github.com/exercise2/Controllers"
	"github.com/exercise2/Routes"
	"github.com/jinzhu/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetFunc(t *testing.T) {
	Config.DB,_ = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	defer Config.DB.Close()
	router:=Routes.SetupRouter()
	router.GET("/user-api/user/",Controllers.GetUsers)

	req,_ := http.NewRequest("GET", "/user-api/user/",nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp,req)
	fmt.Println(resp.Body.String(),"reponse body")
	if resp.Code != 200 {
		t.Errorf("returned unexpected body: got %v want %v",resp.Code,200)
	}
}



func TestPostfunc(t *testing.T) {
	Config.DB,_=gorm.Open("mysql",Config.DbURL(Config.BuildDBConfig()))
	defer Config.DB.Close()
	router := Routes.SetupRouter()
	router.POST("/user-api/user/",Controllers.CreateUser)
	req,_ := http.NewRequest("POST", "/user-api/user/", bytes.NewBuffer([]byte(`{"name":"manish","lastname":"prajapat","dob":"12345678","address":"xyz lane","subject":"maths","marks":"100"}`)))
	resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
	fmt.Println(resp.Body.String(),"reponse body")
	if resp.Code != 200 {
		t.Errorf("handler returned unexpected body: got %v want %v",resp.Code,200)
	}

}

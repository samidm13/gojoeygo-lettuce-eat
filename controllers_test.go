package main

import (
  "testing"
  "io/ioutil"
  "net/http"
  "net/http/httptest"
  "strings"

  // "github.com/gin-gonic/gin"
  )

// func getLoginPOSTPayload() string {
// 	params := url.Values{}
// 	params.Add("username", "user1")
// 	params.Add("password", "pass1")
//
// 	return params.Encode()
// }
//
// func getRegistrationPOSTPayload() string {
// 	params := url.Values{}
// 	params.Add("username", "u1")
// 	params.Add("password", "p1")
//
// 	return params.Encode()
// }

// func getRouter(withTemplates bool) *gin.Engine {
// 	r := gin.Default()
// 	if withTemplates {
// 		r.LoadHTMLGlob("templates/*")
// 		r.Use(setUserStatus()) // new line
// 	}
// 	return r
// }

// func TestShowIndexPage(t *testing.T) {
// 	r := setupRouter()
// 	w := httptest.NewRecorder()
//   c.SetCookie("name", 1, 3600, "", "", false, true)
// 	req, _ := http.NewRequest("GET", "/", nil)
//
//   r.ServeHTTP(w, req)
//
// 	if status := w.Code; status != http.StatusOK {
//     t.Errorf("handler returned wrong status code: actual %v expected %v", status, http.StatusOK)
//   }
//
// 		p, err := ioutil.ReadAll(w.Body)
//     if err != nil || strings.Index(string(p), "<title>Home Page</title>") < 0 {
//       	t.Fail()
// 	   }
// }

func TestShowSignLogPage(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/signlog", nil)

  r.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusOK {
    t.Errorf("handler returned wrong status code: actual %v expected %v", status, http.StatusOK)
  }

		p, err := ioutil.ReadAll(w.Body)
    if err != nil || strings.Index(string(p), "<title>Sign Up / Log In</title>") < 0 {
      	t.Fail()
	   }
}

func TestShowRestaurants(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
  http.SetCookie(w, &http.Cookie{Name: "name", Value: "1"})
	req, _ := http.NewRequest("GET", "/restaurants", nil)

  r.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusSeeOther {
    t.Errorf("handler returned wrong status code: actual %v expected %v", status, http.StatusSeeOther)
  }

		p, err := ioutil.ReadAll(w.Body)
    if err != nil || strings.Index(string(p), "<title>Restaurants</title>") < 0 {
      	t.Fail()
	   }
}

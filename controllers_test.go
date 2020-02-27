package main

import (
  "testing"
  "io/ioutil"
  "net/http"
  "net/http/httptest"
  "strings"
  )

func TestShowIndexPage(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
  http.SetCookie(w, &http.Cookie{Name: "name", Value: "1"})
	req, _ := http.NewRequest("GET", "/", nil)

  r.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusSeeOther {
    t.Errorf("handler returned wrong status code: actual %v expected %v", status, http.StatusSeeOther)
  }

		p, err := ioutil.ReadAll(w.Body)
    if err != nil || strings.Index(string(p), "<title>Home Page</title>") < 0 {
      	t.Fail()
	   }
}

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

// func TestOrderSetUp(t *testing.T) {
// 	r := setupRouter()
// 	w := httptest.NewRecorder()
//   http.SetCookie(w, &http.Cookie{Name: "name", Value: "1"})
//
//   orderPayload := getOrderPOSTPayload()
//   req, _ := http.NewRequest("POST", "/orders", strings.NewReader(orderPayload))
// 	req.Header = http.Header{"Cookie": w.HeaderMap["Set-Cookie"]}
// 	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
// 	req.Header.Add("Content-Length", strconv.Itoa(len(orderPayload)))
//
//   r.ServeHTTP(w, req)
//
// 	if status := w.Code; status != http.StatusSeeOther {
//     t.Errorf("handler returned wrong status code: actual %v expected %v", status, http.StatusSeeOther)
//   }
//
// 		p, err := ioutil.ReadAll(w.Body)
//     if err != nil || strings.Index(string(p), "<title>Order</title>") < 0 {
//       	t.Fail()
// 	   }
// }
//
// func getOrderPOSTPayload() string {
// 	params := url.Values{}
// 	params.Add("RestID", "1")
// 	params.Add("token", "4321")
//   params.Add("address", "London")
//   params.Add("token", "4321")
//   params.Add("appt", time.Now())
//
// 	return params.Encode()
// }

func TestShowOrderPage(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
  http.SetCookie(w, &http.Cookie{Name: "name", Value: "1"})
	req, _ := http.NewRequest("GET", "/orders", nil)

  r.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusSeeOther {
    t.Errorf("handler returned wrong status code: actual %v expected %v", status, http.StatusSeeOther)
  }

		p, err := ioutil.ReadAll(w.Body)
    if err != nil || strings.Index(string(p), "<title>Order</title>") < 0 {
      	t.Fail()
	   }
}

func TestCreateBasket(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
  http.SetCookie(w, &http.Cookie{Name: "name", Value: "1"})
	req, _ := http.NewRequest("POST", "/basket", nil)

  r.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusOK {
    t.Errorf("handler returned wrong status code: actual %v expected %v", status, http.StatusOK)
  }

		p, err := ioutil.ReadAll(w.Body)
    if err != nil || strings.Index(string(p), "<title>User Confirmation</title>") < 0 {
      	t.Fail()
	   }
}

func TestShowAdminPage(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
  http.SetCookie(w, &http.Cookie{Name: "name", Value: "1"})
	req, _ := http.NewRequest("POST", "/adminconf", nil)

  r.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusOK {
    t.Errorf("handler returned wrong status code: actual %v expected %v", status, http.StatusOK)
  }

		p, err := ioutil.ReadAll(w.Body)
    if err != nil || strings.Index(string(p), "<title>Admin Confirmation</title>") < 0 {
      	t.Fail()
	   }
}

func TestAdminBasket(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
  http.SetCookie(w, &http.Cookie{Name: "name", Value: "1"})
	req, _ := http.NewRequest("POST", "/menu/adminorder", nil)

  r.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusOK {
    t.Errorf("handler returned wrong status code: actual %v expected %v", status, http.StatusOK)
  }

		p, err := ioutil.ReadAll(w.Body)
    if err != nil || strings.Index(string(p), "<title>Menu</title>") < 0 {
      	t.Fail()
	   }
}

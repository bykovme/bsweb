package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/bykovme/bslib"
)

// LoginPage - registration page structure
type ItemPage struct {
	ErrorText string
	IsError   bool
	Item      bslib.BSItem
}

func viewItemPage(w http.ResponseWriter, r *http.Request) {
	var item ItemPage
	bsInstance := bslib.GetInstance()
	var itemId string
	itemId, r.URL.Path = ShiftPath(r.URL.Path)

	if itemId == "" {
		ErrorPage(w, "Error", "Nothing found")
		return
	}
	iid, err := strconv.ParseInt(itemId, 10, 64)
	if err != nil {
		ErrorPage(w, "Error", err.Error())
		return
	}

	item.Item, err = bsInstance.ReadItemById(iid)

	if err != nil {
		log.Println(err.Error())
		ErrorPage(w, "Error", err.Error())
		return
	}

	drillDownViewItem(w, r, item)

}

func showItemDetails(w http.ResponseWriter, item ItemPage) {
	err := renderHTMLTemplate(w, "item_view", item)
	if err != nil {
		ErrorPage(w, "Error", err.Error())
	}
}

func drillDownViewItem(w http.ResponseWriter, r *http.Request, item ItemPage) {
	var action string
	action, r.URL.Path = ShiftPath(r.URL.Path)

	switch action {
	case "addfield":
		processAddNewField(w, r, item)
	case "deletefield":
		processDeleteField(w, r, item)
	default:
		showItemDetails(w, item)
	}
}

func processDeleteField(w http.ResponseWriter, r *http.Request, item ItemPage) {
	if r.Method == "GET" {
		showItemDetails(w, item)
	} else if r.Method == "POST" {
		deleteField(w, r, item)
	}
}

func processAddNewField(w http.ResponseWriter, r *http.Request, item ItemPage) {
	if r.Method == "GET" {
		err := renderHTMLTemplate(w, "add_field", item)
		if err != nil {
			ErrorPage(w, "Error", err.Error())
		}
	} else if r.Method == "POST" {
		addNewField(w, r, item)
	}
}

func deleteField(w http.ResponseWriter, r *http.Request, item ItemPage) {
	var delField bslib.UpdateFieldForm

	bsInst := bslib.GetInstance()
	fId, err := getValueByName(r, "field_id")
	if err != nil || fId == "" {
		if err != nil {
			item.ErrorText = err.Error()
		} else {
			item.ErrorText = "Field id is empty"
		}
		item.IsError = true
	}
	if item.IsError {
		if err != nil {
			ErrorPage(w, "Error", item.ErrorText)
		}
		return
	}
	delField.ID, err = strconv.ParseInt(fId, 10, 64)
	if err != nil {
		ErrorPage(w, "Error", err.Error())
		return
	}
	resp, delErr := bsInst.DeleteField(delField)

	if delErr != nil {
		ErrorPage(w, "Error", delErr.Error())
		return
	}
	if resp.Status != bslib.ConstSuccessResponse {
		ErrorPage(w, "Error", resp.MsgTxt)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/items/%d", item.Item.ID), 302)
}

func addNewField(w http.ResponseWriter, r *http.Request, item ItemPage) {
	var newField bslib.UpdateFieldForm
	var err error
	bsInst := bslib.GetInstance()

	newField.Name, err = getValueByName(r, "field_name")
	if err != nil || newField.Name == "" {
		if err != nil {
			item.ErrorText = err.Error()
		} else {
			item.ErrorText = "Field name is empty"
		}
		item.IsError = true
	}

	newField.Icon, err = getValueByName(r, "field_icon")
	if err != nil || newField.Icon == "" {
		if err != nil {
			item.ErrorText = err.Error()
		} else {
			item.ErrorText = "Field icon is empty"
		}
		item.IsError = true
	}

	newField.ValueType, err = getValueByName(r, "value_type")
	if err != nil || newField.ValueType == "" {
		if err != nil {
			item.ErrorText = err.Error()
		} else {
			item.ErrorText = "Value type is empty"
		}
		item.IsError = true
	}

	newField.Value, err = getValueByName(r, "field_value")
	if err != nil || newField.Value == "" {
		if err != nil {
			item.ErrorText = err.Error()
		} else {
			item.ErrorText = "Value is empty"
		}
		item.IsError = true
	}

	if item.IsError {
		err := renderHTMLTemplate(w, "add_field", item)
		if err != nil {
			ErrorPage(w, "Error", err.Error())
		}
		return
	}

	newField.ItemID = item.Item.ID
	res, err := bsInst.AddNewField(newField)
	if err != nil {
		item.IsError = true
		item.ErrorText = err.Error()
		err := renderHTMLTemplate(w, "add_field", item)
		if err != nil {
			ErrorPage(w, "Error", err.Error())
		}
		return
	}

	if res.Status != bslib.ConstSuccessResponse {
		item.IsError = true
		item.ErrorText = res.MsgTxt
		err := renderHTMLTemplate(w, "add_field", item)
		if err != nil {
			ErrorPage(w, "Error", err.Error())
		}
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/items/%d", item.Item.ID), 302)

}

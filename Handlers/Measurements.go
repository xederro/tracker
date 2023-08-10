package Handlers

import (
	"fmt"
	"healthtracker/Database/Measurement"
	"healthtracker/Templater"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func Measurements(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		s := strings.Split(req.URL.Path, "/")

		if len(s) > 2 && s[2] == "add" {
			getCreator(w, req)
		} else if len(s) > 3 && s[2] == "edit" {
			editForm(w, req, s[3])
		} else if len(s) > 3 && s[2] == "page" {
			atoi, err := strconv.Atoi(s[3])
			if err != nil {
				atoi = 1
			}

			getAll(w, req, atoi)
		} else if len(s) > 2 && s[2] != "" {

			getOne(w, req, s[2])
		} else {
			getAll(w, req, 1)
		}
	} else if req.Method == http.MethodDelete {
		s := strings.Split(req.URL.Path, "/")

		deleteRow(w, req, s[2])
	} else if req.Method == http.MethodPost {
		create(w, req)
	} else if req.Method == http.MethodPut {
		s := strings.Split(req.URL.Path, "/")

		edit(w, req, s[2])
	}
}

func getCreator(w http.ResponseWriter, req *http.Request) {
	err := Templater.Templater.ExecuteTemplate(w, "creator.html", "")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func getOne(w http.ResponseWriter, req *http.Request, id string) {
	m, err := Measurement.GetOne(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = Templater.Templater.ExecuteTemplate(w, "measurement.html", m)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func getAll(w http.ResponseWriter, req *http.Request, page int) {
	m, err := Measurement.GetAll(page)
	if err != nil {
		fmt.Println(err)
		return
	}

	t := struct {
		New   bool
		Mes   []*Measurement.Measurement
		Pages int
		Cur   int
		Next  int
		Prev  int
	}{Measurement.DoneToday(), m, Measurement.Count(), page, page + 1, page - 1}

	err = Templater.Templater.ExecuteTemplate(w, "measurements.html", t)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func deleteRow(w http.ResponseWriter, req *http.Request, id string) {
	err := Measurement.Delete(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	getAll(w, req, 1)
}

func create(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}

	form, err := bindForm(req.Form)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = Measurement.Create(form)
	if err != nil {
		fmt.Println(err)
		return
	}

	getAll(w, req, 1)
}

func editForm(w http.ResponseWriter, req *http.Request, id string) {
	m, err := Measurement.GetOne(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = Templater.Templater.ExecuteTemplate(w, "editor.html", m)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func edit(w http.ResponseWriter, req *http.Request, id string) {
	err := req.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}

	form, err := bindForm(req.Form)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = Measurement.Edit(form, id)
	if err != nil {
		fmt.Println(err)
		return
	}

	getOne(w, req, id)
}

func bindForm(Form url.Values) (*Measurement.Measurement, error) {
	Steps, err := strconv.Atoi(Form.Get("Steps"))
	if err != nil {
		return nil, err
	}

	SleepDuration, err := strconv.Atoi(Form.Get("SleepDuration"))
	if err != nil {
		return nil, err
	}

	SleepQuality, err := strconv.Atoi(Form.Get("SleepQuality"))
	if err != nil {
		return nil, err
	}

	Mood, err := strconv.Atoi(Form.Get("Mood"))
	if err != nil {
		return nil, err
	}

	Energy, err := strconv.Atoi(Form.Get("Energy"))
	if err != nil {
		return nil, err
	}

	m := Measurement.Measurement{
		Steps:         Steps,
		SleepDuration: SleepDuration,
		SleepQuality:  SleepQuality,
		Mood:          Mood,
		Energy:        Energy,
	}

	return &m, nil
}

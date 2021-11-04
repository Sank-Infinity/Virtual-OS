package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io/ioutil"
	"strconv"

	"fyne.io/fyne/v2"
	// "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"

	// "fyne.io/fyne/v2/container"
	// "fyne.io/fyne/v2/widget"
	"net/http"
)

func showWeatherApp( w fyne.Window) {
	// a := app.New()
	// w := a.NewWindow("Weather App")
	// w.Resize(fyne.NewSize(800, 600))
	
	// api part

	responce, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=mumbai&APPID=24beb57b7231f56a8dc3ebd3ad367e5d")

	if err != nil {
		fmt.Println(err)
	}

	defer responce.Body.Close()

	body , err := ioutil.ReadAll(responce.Body)
	if err != nil {
		fmt.Println(err)
	}

	weather, err := UnmarshalWelcome(body)

	if err != nil {
		fmt.Println(err)
	}

	img := canvas.NewImageFromFile("weather.png")
	img.FillMode = canvas.ImageFillOriginal

	label1 := canvas.NewText("Weather Details", color.White)
	label1.TextStyle = fyne.TextStyle{Bold: true}

	label2 := canvas.NewText(fmt.Sprintf("Country %s", weather.Sys.Country), color.White)

	label3 := canvas.NewText(fmt.Sprintf("Wind Speed %.2f", weather.Wind.Speed), color.White)

	label4 := canvas.NewText(fmt.Sprintf("Temperature %.2f", weather.Main.Temp - 273), color.White)

	hum := strconv.FormatInt(int64( weather.Main.Humidity ), 10)
	label5 := canvas.NewText(fmt.Sprintf("Humidity %v", hum), color.White)

	sr := strconv.FormatInt(int64( weather.Sys.Sunrise  ), 10)
	label6 := canvas.NewText(fmt.Sprintf("Sunrise %v", sr), color.White)

	st := strconv.FormatInt(int64( weather.Sys.Sunset  ), 10)
	label7 := canvas.NewText(fmt.Sprintf("Sunset %v", st), color.White)

	prs := strconv.FormatInt(int64( weather.Main.Pressure  ), 10)
	label8 := canvas.NewText(fmt.Sprintf("Pressure %v", prs), color.White)

	weatherContainer :=  container.NewVBox(
			label1,
			img,
			label2,
			label3,
			label4,
			label5,
			label6,
			label7,
			label8,
		)
	

	w.SetContent(container.NewBorder(panelContent, nil, nil, nil, weatherContainer),)
	w.Show()
}

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    welcome, err := UnmarshalWelcome(bytes)
//    bytes, err = welcome.Marshal()

func UnmarshalWelcome(data []byte) (Welcome, error) {
	var r Welcome
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Welcome) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Welcome struct {
	Coord      Coord     `json:"coord"`     
	Weather    []Weather `json:"weather"`   
	Base       string    `json:"base"`      
	Main       Main      `json:"main"`      
	Visibility int64     `json:"visibility"`
	Wind       Wind      `json:"wind"`      
	Clouds     Clouds    `json:"clouds"`    
	Dt         int64     `json:"dt"`        
	Sys        Sys       `json:"sys"`       
	Timezone   int64     `json:"timezone"`  
	ID         int64     `json:"id"`        
	Name       string    `json:"name"`      
	Cod        int64     `json:"cod"`       
}

type Clouds struct {
	All int64 `json:"all"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Main struct {
	Temp      float64 `json:"temp"`      
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`  
	TempMax   float64 `json:"temp_max"`  
	Pressure  int64   `json:"pressure"`  
	Humidity  int64   `json:"humidity"`  
	SeaLevel  int64   `json:"sea_level"` 
	GrndLevel int64   `json:"grnd_level"`
}

type Sys struct {
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"` 
}

type Weather struct {
	ID          int64  `json:"id"`         
	Main        string `json:"main"`       
	Description string `json:"description"`
	Icon        string `json:"icon"`       
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int64   `json:"deg"`  
	Gust  float64 `json:"gust"` 
}

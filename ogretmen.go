package main

import "fmt"

type ogretmen struct {
	ad    string
	yas   int
	brans string
}

func (ogr *ogretmen) bilgileriAl() string {
	return ("selam ben " + ogr.ad + " ve ben " + ogr.brans + " ogretmeniyim")
}
func (ogr ogretmen) derseGir() {
	fmt.Println(ogr.brans+" dersine giriliyor -", ogr.ad+" hoca")
}

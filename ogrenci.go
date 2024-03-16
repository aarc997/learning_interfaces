package main

import "strconv"

type ogrenci struct {
	ogrencino int
	ad        string
	sinif     int
}

func (ogr *ogrenci) bilgileriAl() string {
	return ("selam ben " + ogr.ad + " okul numaram " + strconv.Itoa(ogr.ogrencino) + " ve ben " + strconv.Itoa(ogr.sinif) + " . sinif ogrencisiyim")
}

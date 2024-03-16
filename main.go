package main

import "fmt"

func main() {
	var ogrenci1 insan = &ogrenci{1, "nazif süleyman", 11}
	var ogretmen1 insan = &ogretmen{"kemal", 35, "Web Tabanlı Uygulamalar"}
	var ogretmen2 ogretmen = ogretmen{"selim", 49, "mobil uygulamalar"}
	fmt.Println(ogrenci1.bilgileriAl())
	fmt.Println(ogretmen1.bilgileriAl())
	ogretmen2.derseGir()
}

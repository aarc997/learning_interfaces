# learning_interfaces
bu repo ile golang de interface implementasyonu ögreniyoruz!

<h1>Hedeflenen Uygulama :</h1>
<p>ogretmen ve ogrenci structına sahibiz ve bu 2 struct arasında bir ortak nokta oluşturmak için insan adında bir interface kullanıyoruz</p>

<h3>insan interface'i özellikleri</h3>
<ul>
  <li>string değer döndüren "bilgileriAl()" methodunu barındırıyor</li>
</ul>

<h3>ogretmen structı özellikleri</h3>
<ul>
  <li>ad, yas, brans değişkenlerine sahip olacak</li>
  <li>derseGir() ve imlement ettigi insan interface'inden gelen bilgileriAl() methoduna sahip</li>
</ul>

<h3>ogrenci structı özellikleri</h3>
<ul>
  <li>ogrencino, ad, sinif değişkenlerine sahip olacak</li>
  <li>imlement ettigi insan interfaceinden gelen bilgileriAl() methoduna sahip</li>
</ul>

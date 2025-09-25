package budaya

import (
	"time"

	"github.com/wkloose/Buddya/initializers"
	"github.com/wkloose/Buddya/models"
)

func SeedBudayaBatch1() {
	cities := []models.City{
		{
			NamaKota: "Jakarta",
			Provinsi: "DKI Jakarta",
			DeskripsiBudaya: "Jakarta, sebagai ibu kota Indonesia, memiliki budaya Betawi yang kental. Budaya ini tercermin dalam kesenian ondel-ondel, lenong, gambang kromong, serta kuliner khas seperti kerak telor dan soto Betawi. Selain itu, Jakarta juga merupakan pusat percampuran berbagai budaya nusantara dan mancanegara karena menjadi magnet urbanisasi. Keberagaman ini menciptakan identitas unik yang membedakan Jakarta dari kota lainnya di Indonesia.",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			NamaKota: "Tangerang",
			Provinsi: "Banten",
			DeskripsiBudaya: "Tangerang memiliki identitas budaya yang kuat melalui komunitas Tionghoa Benteng yang sudah ada sejak ratusan tahun. Tradisi Cap Go Meh, Barongsai, dan rumah-rumah tua bergaya arsitektur Cina menjadi ciri khas. Selain itu, budaya Betawi pinggiran juga masih lestari melalui kesenian gambang kromong dan tradisi pernikahan adat Betawi. Tangerang menjadi salah satu contoh akulturasi budaya lokal, Tionghoa, dan modernitas perkotaan.",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			NamaKota: "Cilegon",
			Provinsi: "Banten",
			DeskripsiBudaya: "Cilegon dikenal sebagai Kota Baja, namun di balik modernitas industrinya, budaya tradisional masih hidup. Seni debus khas Banten yang menampilkan kekuatan fisik luar biasa masih sering dipentaskan. Masyarakat Cilegon juga memegang teguh tradisi keislaman, yang tampak dalam kegiatan keagamaan, pengajian, dan ritual adat. Perpaduan antara budaya pesisir, tradisi Banten, dan pengaruh modern menjadikan Cilegon unik.",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			NamaKota: "Serang",
			Provinsi: "Banten",
			DeskripsiBudaya: "Serang adalah pusat kebudayaan Banten yang terkenal dengan seni bela diri debus. Selain itu, tradisi pencak silat, seni terbang gede (rebana besar), dan upacara adat keagamaan masih lestari. Serang juga memiliki peninggalan sejarah seperti Keraton Surosowan dan Masjid Agung Banten yang menjadi simbol kejayaan Kesultanan Banten. Budaya masyarakat Serang sangat erat kaitannya dengan nilai-nilai Islam dan sejarah kejayaan Banten.",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			NamaKota: "Tangerang Selatan",
			Provinsi: "Banten",
			DeskripsiBudaya: "Tangerang Selatan adalah kota urban yang tumbuh pesat, namun tetap menyimpan jejak budaya Betawi dan Sunda. Tradisi pernikahan adat, seni tari topeng, dan kuliner khas masih bisa ditemukan. Masyarakat Tangsel juga memelihara kegiatan seni modern yang berpadu dengan budaya tradisional. Hal ini menjadikan Tangsel sebagai kota dengan identitas budaya hybrid, menggabungkan tradisi lokal dengan gaya hidup metropolitan.",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			NamaKota: "Bandung",
			Provinsi: "Jawa Barat",
			DeskripsiBudaya: "Bandung adalah pusat kebudayaan Sunda dengan kesenian tradisional seperti angklung, wayang golek, dan tari jaipong. Kota ini juga dikenal dengan fashion, musik, dan kulinernya yang kreatif sehingga dijuluki Paris van Java. Tradisi lokal Sunda masih sangat hidup di tengah masyarakat, seperti upacara adat seren taun dan berbagai ritual pertanian. Bandung adalah kota yang menggabungkan warisan budaya Sunda dengan inovasi modern.",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			NamaKota: "Bogor",
			Provinsi: "Jawa Barat",
			DeskripsiBudaya: "Bogor memiliki identitas budaya Sunda yang kental, dengan tradisi wayang golek, kesenian calung, dan angklung gubrak. Kota ini juga terkenal dengan kuliner khasnya seperti asinan Bogor dan doclang. Selain itu, Bogor memiliki sejarah panjang sebagai tempat peristirahatan bangsawan Belanda, yang meninggalkan warisan arsitektur kolonial. Budaya Bogor mencerminkan perpaduan antara tradisi Sunda, sejarah kolonial, dan kehidupan urban modern.",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			NamaKota: "Bekasi",
			Provinsi: "Jawa Barat",
			DeskripsiBudaya: "Bekasi memiliki akar budaya Betawi dan Sunda. Sejarah panjang Bekasi terkait dengan kerajaan Tarumanegara menjadikan kota ini memiliki nilai sejarah tinggi. Kesenian Betawi seperti tanjidor, ondel-ondel, dan lenong masih bisa ditemui di beberapa wilayah. Selain itu, perkembangan urbanisasi membuat Bekasi menjadi kota dengan budaya yang majemuk dan beragam.",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			NamaKota: "Depok",
			Provinsi: "Jawa Barat",
			DeskripsiBudaya: "Depok dikenal dengan budaya Betawi Depok, yang memiliki ciri khas berbeda dari Betawi Jakarta. Tradisi ini terlihat dalam upacara adat, kesenian, dan bahasa sehari-hari masyarakatnya. Depok juga memiliki pengaruh budaya Sunda karena letaknya yang berdekatan dengan Bogor. Kota ini menjadi contoh percampuran dua budaya besar: Betawi dan Sunda.",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			NamaKota: "Cimahi",
			Provinsi: "Jawa Barat",
			DeskripsiBudaya: "Cimahi memiliki budaya Sunda yang kuat, dengan kesenian seperti pencak silat, tari tradisional, dan musik angklung. Kota ini juga dikenal sebagai kota militer yang memengaruhi perkembangan sosial dan budayanya. Masyarakat Cimahi menjaga tradisi Sunda melalui berbagai festival dan kegiatan budaya. Perpaduan antara budaya tradisional dan pengaruh militer menjadikan Cimahi unik di Jawa Barat.",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	initializers.DB.Create(&cities)
}

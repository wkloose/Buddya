package budaya

import (
	"time"

	"github.com/wkloose/Buddya/initializers"
	"github.com/wkloose/Buddya/models"
)

func SeedBudayaBatch2() {
	cities := []models.City{
		{
			NamaKota: "Sukabumi",
			Provinsi: "Jawa Barat",
			DeskripsiBudaya: "Sukabumi memiliki budaya Sunda yang kuat dengan kehidupan masyarakat agraris dan pesisir. Tradisi kesenian rakyat seperti pencak silat, wayang golek, dan gondang sunda masih sering dipertunjukkan. Selain itu, masyarakat Sukabumi menjunjung tinggi tradisi gotong royong dan nilai kekeluargaan. Kuliner khas seperti nasi uduk dan kue mochi menjadi bagian dari identitas budaya Sukabumi.",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			NamaKota: "Tasikmalaya",
			Provinsi: "Jawa Barat",
			DeskripsiBudaya: "Tasikmalaya terkenal sebagai pusat kerajinan tangan di Jawa Barat, dengan produk unggulan berupa bordir, payung geulis, dan batik tasik. Budaya Sunda di Tasikmalaya juga tercermin dalam kesenian tari jaipong, angklung, dan upacara adat seren taun. Kehidupan masyarakatnya banyak dipengaruhi oleh tradisi agraris yang masih bertahan hingga kini. Tasikmalaya memadukan kreativitas seni dengan kearifan lokal Sunda.",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			NamaKota: "Cirebon",
			Provinsi: "Jawa Barat",
			DeskripsiBudaya: "Cirebon dikenal sebagai kota dengan akulturasi budaya Jawa, Sunda, dan Tionghoa. Hal ini terlihat jelas dalam kesenian tari topeng Cirebon, batik motif megamendung, dan arsitektur keraton yang megah. Sebagai kota pelabuhan, Cirebon memiliki sejarah panjang interaksi dengan pedagang asing yang memperkaya kebudayaannya. Hingga kini, Cirebon tetap menjadi salah satu pusat budaya pesisir Jawa.",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			NamaKota: "Banjar",
			Provinsi: "Jawa Barat",
			DeskripsiBudaya: "Banjar, meskipun kota kecil, memiliki budaya Sunda yang kental dengan pengaruh Priangan. Kesenian rakyat seperti calung, jaipong, dan wayang golek masih dilestarikan masyarakat. Tradisi agraris dan nilai religius menjadi bagian penting dari kehidupan sehari-hari. Kuliner tradisional dan festival budaya daerah juga memperkuat identitas Banjar sebagai bagian dari warisan Sunda.",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			NamaKota: "Semarang",
			Provinsi: "Jawa Tengah",
			DeskripsiBudaya: "Semarang adalah kota pesisir yang mencerminkan budaya Jawa dengan pengaruh kuat dari komunitas Tionghoa dan Belanda. Tradisi Dugderan yang diadakan menjelang Ramadan adalah salah satu budaya khas kota ini. Selain itu, arsitektur kolonial, klenteng Sam Poo Kong, dan masjid bersejarah menjadi simbol keberagaman budaya Semarang. Batik Semarangan dengan motif Tugu Muda dan Lawang Sewu juga memperkaya identitas budaya kota ini.",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			NamaKota: "Surakarta",
			Provinsi: "Jawa Tengah",
			DeskripsiBudaya: "Surakarta atau Solo dikenal sebagai salah satu pusat kebudayaan Jawa. Kota ini terkenal dengan keraton Kasunanan dan Pura Mangkunegaran yang menjaga tradisi Jawa klasik. Kesenian wayang kulit, gamelan, dan batik Solo menjadi warisan budaya yang masih dilestarikan. Tradisi Jawa halus yang mengutamakan tata krama, unggah-ungguh, dan kesenian membuat Solo dijuluki 'Spirit of Java'.",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			NamaKota: "Salatiga",
			Provinsi: "Jawa Tengah",
			DeskripsiBudaya: "Salatiga merupakan kota kecil yang multikultural, dengan perpaduan budaya Jawa, Tionghoa, dan Belanda. Kehidupan masyarakatnya kental dengan tradisi Jawa seperti slametan, wayang, dan gamelan. Salatiga juga dikenal sebagai kota pendidikan yang terbuka terhadap budaya modern, sehingga perpaduan tradisi dan modernitas terlihat jelas di kota ini. Kulinernya, seperti ronde dan enting-enting gepuk, juga menjadi bagian dari identitas budaya Salatiga.",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			NamaKota: "Magelang",
			Provinsi: "Jawa Tengah",
			DeskripsiBudaya: "Magelang memiliki identitas budaya Jawa yang kental, ditandai dengan keberadaan Candi Borobudur sebagai warisan dunia. Kehidupan masyarakatnya dipengaruhi oleh tradisi agraris dan nilai religius yang kuat. Upacara adat Waisak di Borobudur adalah salah satu peristiwa budaya dan spiritual penting di kota ini. Selain itu, Magelang juga memiliki seni tari, musik gamelan, dan kuliner khas yang memperkaya budayanya.",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			NamaKota: "Pekalongan",
			Provinsi: "Jawa Tengah",
			DeskripsiBudaya: "Pekalongan dikenal sebagai Kota Batik Dunia dengan tradisi membatik yang sudah berlangsung ratusan tahun. Motif batik pesisir Pekalongan terkenal berwarna cerah dan dipengaruhi budaya Tionghoa, Arab, dan Belanda. Selain batik, Pekalongan juga memiliki kesenian rakyat seperti tari sintren dan tradisi laut sedekah laut. Kota ini adalah contoh nyata akulturasi budaya pesisir Jawa yang tetap berkembang hingga kini.",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			NamaKota: "Tegal",
			Provinsi: "Jawa Tengah",
			DeskripsiBudaya: "Tegal memiliki budaya pesisir Jawa dengan logat bahasa yang khas dan dikenal dengan sebutan ngapak. Tradisi nelayan, kesenian rakyat seperti wayang golek cepak, dan tari endel merupakan bagian dari identitas budaya Tegal. Kuliner khas seperti sate tegal dan teh poci juga memperkuat daya tarik kota ini. Budaya Tegal mencerminkan keberanian, keterbukaan, dan humor yang khas dari masyarakatnya.",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	initializers.DB.Create(&cities)
}

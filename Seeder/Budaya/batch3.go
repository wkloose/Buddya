package budaya

import (
	"time"

	"github.com/wkloose/Buddya/initializers"
	"github.com/wkloose/Buddya/models"
)

func SeedBudayaBatch3() {
	cities := []models.City{
		{
			NamaKota: "Yogyakarta",
			Provinsi: "DI Yogyakarta",
			DeskripsiBudaya: "Yogyakarta adalah pusat kebudayaan Jawa yang masih hidup hingga kini. Kota ini dikenal dengan Keraton Yogyakarta yang menjaga tradisi kesultanan, seni gamelan, batik, dan wayang kulit. Tradisi akademik dan seni modern juga berkembang pesat, menjadikan Yogyakarta sebagai kota pelajar dan kota budaya. Ritual keagamaan, upacara adat, serta kehidupan seni jalanan membuat Yogyakarta menjadi simbol harmoni antara tradisi dan modernitas.",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			NamaKota: "Surabaya",
			Provinsi: "Jawa Timur",
			DeskripsiBudaya: "Surabaya dikenal dengan julukan Kota Pahlawan, dengan budaya arek Suroboyo yang tegas, lugas, dan penuh semangat. Seni tradisional seperti ludruk, remo, dan karapan sapi di sekitarnya menjadi ciri khas budaya Jawa Timuran. Surabaya juga menjadi kota multikultural dengan keberadaan kampung Arab, Pecinan, dan kawasan kolonial Belanda. Karakter masyarakatnya yang terbuka dan berani mencerminkan jiwa kepahlawanan Surabaya.",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			NamaKota: "Malang",
			Provinsi: "Jawa Timur",
			DeskripsiBudaya: "Malang memiliki budaya Arek Malangan yang khas dengan seni tari topeng Malangan sebagai identitas penting. Kota ini juga terkenal dengan tradisi wayang, musik gamelan, dan kuliner khas seperti rawon dan bakso Malang. Sejarah kerajaan Singhasari dan Majapahit turut memperkaya nilai budaya Malang. Dengan iklim sejuk pegunungan, Malang menjadi pusat pariwisata sekaligus budaya di Jawa Timur.",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			NamaKota: "Batu",
			Provinsi: "Jawa Timur",
			DeskripsiBudaya: "Batu dikenal sebagai kota wisata pegunungan yang juga memiliki budaya Jawa Timuran yang khas. Tradisi pertanian dan kesenian rakyat seperti jaranan dan campursari masih hidup di masyarakat. Kehidupan religius dengan perayaan tradisi desa membuat Batu tetap menjaga akar budayanya meski berkembang sebagai kota wisata modern. Batu mencerminkan perpaduan antara budaya agraris dan industri pariwisata.",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			NamaKota: "Mojokerto",
			Provinsi: "Jawa Timur",
			DeskripsiBudaya: "Mojokerto memiliki nilai sejarah tinggi karena menjadi pusat Kerajaan Majapahit, salah satu kerajaan terbesar di nusantara. Peninggalan berupa candi dan situs sejarah menjadi bagian penting dari identitas budaya kota ini. Tradisi Jawa klasik masih terjaga melalui kesenian gamelan, wayang, dan upacara adat. Mojokerto adalah simbol kejayaan masa lalu yang masih hidup dalam budaya masyarakatnya.",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			NamaKota: "Madiun",
			Provinsi: "Jawa Timur",
			DeskripsiBudaya: "Madiun terkenal dengan budaya pencak silat yang menjadi bagian dari identitas masyarakatnya. Selain itu, kesenian tradisional seperti gendhing gamelan dan tari tradisional masih lestari. Pecel Madiun menjadi ikon kuliner yang memperkuat karakter budaya kota ini. Kehidupan masyarakat Madiun mencerminkan perpaduan antara budaya agraris dan semangat kemandirian khas Jawa Timur.",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			NamaKota: "Kediri",
			Provinsi: "Jawa Timur",
			DeskripsiBudaya: "Kediri memiliki sejarah panjang sebagai pusat Kerajaan Kediri di masa lalu. Budaya Jawa klasik masih sangat terasa melalui seni wayang, gamelan, dan tradisi keagamaan. Kediri juga dikenal dengan kuliner khas seperti tahu takwa dan gethuk pisang. Kehidupan masyarakatnya masih memegang nilai gotong royong dan kearifan lokal yang diwariskan turun-temurun.",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			NamaKota: "Blitar",
			Provinsi: "Jawa Timur",
			DeskripsiBudaya: "Blitar dikenal sebagai kota bersejarah karena menjadi tempat makam Bung Karno, proklamator kemerdekaan Indonesia. Budaya Jawa Timuran dengan kesenian ludruk, wayang, dan jaranan masih kuat di Blitar. Selain itu, masyarakatnya menjunjung tinggi nilai patriotisme dan nasionalisme. Blitar adalah kota yang memadukan budaya tradisional, sejarah, dan semangat kebangsaan.",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			NamaKota: "Probolinggo",
			Provinsi: "Jawa Timur",
			DeskripsiBudaya: "Probolinggo memiliki budaya masyarakat pesisir yang bercampur dengan tradisi Jawa Timuran. Kesenian jaranan, hadrah, dan tarian daerah menjadi bagian penting dari identitas budaya kota ini. Probolinggo juga dikenal dengan festival anggur dan mangga yang mencerminkan kehidupan agraris masyarakatnya. Tradisi dan keanekaragaman kuliner membuat Probolinggo memiliki warna budaya yang khas.",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			NamaKota: "Pasuruan",
			Provinsi: "Jawa Timur",
			DeskripsiBudaya: "Pasuruan adalah kota pesisir yang dikenal dengan budaya nelayan dan tradisi maritimnya. Upacara adat sedekah laut menjadi simbol penting kehidupan masyarakat pesisir Pasuruan. Kesenian rakyat seperti wayang kulit, hadrah, dan tari tradisional masih sering dipentaskan. Budaya agraris dan maritim berpadu di Pasuruan, menciptakan identitas unik bagi kota ini.",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	initializers.DB.Create(&cities)
}

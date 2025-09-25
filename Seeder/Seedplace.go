package Seeder

import (
	"log"

	"github.com/wkloose/Buddya/initializers"
	"github.com/wkloose/Buddya/models"
)

func SeedPlaces() {
	var DB = initializers.DB
	var cities []models.City
	if err := DB.Find(&cities).Error; err != nil {
		log.Fatal("❌ gagal ambil cities:", err)
	}

	cityMap := make(map[string]uint)
	for _, c := range cities {
		cityMap[c.NamaKota] = c.ID
	}
	var placesData = []struct { 
		NamaTempat string
		Kategori string 
		Alamat string 
		Latitude float64 
		Longitude float64 
		Deskripsi string 
		ImageURL string 
		NamaKota string 
		}{

	{"Klenteng Boen Tek Bio", "wisata", "Jl. Bhakti No.14, Pasar Lama, Tangerang", -6.178, 106.632, "Klenteng bersejarah di pusat kota.", "", "Tangerang"},
	{"Museum Benteng Heritage", "wisata", "Jl. Cilame No.18-20, Pasar Lama, Tangerang", -6.177, 106.631, "Museum sejarah China-Indonesia lokal.", "", "Tangerang"},
	{"Masjid Agung Al-Azhom", "wisata", "Jl. Satria Sudirman, Sukaasih, Tangerang", -6.176, 106.63, "Masjid besar dan ikon religi.", "", "Tangerang"},
	{"Laksa Tangerang", "kuliner", "Jl. Kisamaun, Pasar Lama, Tangerang", -6.176, 106.637, "Laksa khas Tangerang.", "", "Tangerang"},
	{"Nasi Sumsum", "kuliner", "Jl. Kisamaun, Pasar Lama, Tangerang", -6.177, 106.635, "Nasi sumsum tradisional.", "", "Tangerang"},
	{"Gecom Rawa Belong", "kuliner", "Jl. Rawa Belong, Karawaci, Tangerang", -6.178, 106.634, "Makanan lokal populer.", "", "Tangerang"},

	{"Gunung Batu Lawang", "wisata", "Jl. Gunung Batu Lawang, Citangkil, Cilegon", -6.012, 106.06, "Formasi batu dan pemandangan alam.", "", "Cilegon"},
	{"Monumen Krakatau Steel", "wisata", "Jl. Raya Krakatau, Kebonsari, Cilegon", -6.014, 106.054, "Monumen industri dan sejarah kota.", "", "Cilegon"},
	{"Masjid Agung Cilegon", "wisata", "Jl. Jenderal Sudirman No.1, Jombang Wetan, Cilegon", -6.016, 106.05, "Masjid utama di kota.", "", "Cilegon"},
	{"Sate Bandeng Cilegon", "kuliner", "Jl. Raya Cilegon, Jombang, Cilegon", -6.017, 106.052, "Sate bandeng khas Cilegon.", "", "Cilegon"},
	{"Nasi Gonjleng", "kuliner", "Jl. Raya Anyer, Cibeber, Cilegon", -6.018, 106.055, "Kuliner tradisional lokal.", "", "Cilegon"},
	{"Rengginang Cilegon", "kuliner", "Jl. Seruni, Cilegon", -6.02, 106.056, "Jajanan rengginang khas.", "", "Cilegon"},

	{"Keraton Surosowan", "wisata", "Kampung Kroya, Kasemen, Serang", -6.044, 106.163, "Sisa benteng kesultanan Banten.", "", "Serang"},
	{"Masjid Agung Banten", "wisata", "Jl. Raya Banten Lama, Kasemen, Serang", -6.043, 106.163, "Masjid bersejarah di Banten.", "", "Serang"},
	{"Benteng Speelwijk", "wisata", "Kampung Pamarican, Kasemen, Serang", -6.041, 106.165, "Benteng kolonial Belanda.", "", "Serang"},
	{"Sate Bandeng Serang", "kuliner", "Jl. Raya Banten Lama, Serang", -6.042, 106.166, "Sate bandeng lokal.", "", "Serang"},
	{"Rabeg Khas Serang", "kuliner", "Jl. Raya Serang-Banten Lama, Serang", -6.041, 106.167, "Masakan khas Banten.", "", "Serang"},
	{"Emping Melinjo", "kuliner", "Jl. Raya Serang, Serang", -6.043, 106.168, "Camilan tradisional.", "", "Serang"},

	{"Masjid Raya Bani Umar", "wisata", "Jl. Graha Raya Bintaro, Pondok Aren, Tangerang Selatan", -6.293, 106.72, "Masjid ikonik daerah.", "", "Tangerang Selatan"},
	{"Situ Gintung", "wisata", "Ciputat Timur, Tangerang Selatan", -6.317, 106.75, "Danau bekas tanggul dengan taman.", "", "Tangerang Selatan"},
	{"Taman Kota BSD", "wisata", "Jl. Letnan Sutopo, BSD City, Serpong, Tangerang Selatan", -6.301, 106.65, "Ruang terbuka hijau dan event.", "", "Tangerang Selatan"},
	{"Nasi Uduk Tangsel", "kuliner", "Jl. Raya Pondok Aren, Tangerang Selatan", -6.29, 106.72, "Nasi uduk lokal populer.", "", "Tangerang Selatan"},
	{"Sate Bandeng Tangsel", "kuliner", "Jl. Raya Serpong, Tangerang Selatan", -6.3, 106.73, "Sate bandeng versi Tangsel.", "", "Tangerang Selatan"},
	{"Ketan Bintaro", "kuliner", "Jl. Bintaro Utara Raya, Pondok Aren, Tangerang Selatan", -6.2905, 106.74, "Ketan dan jajanan pasar.", "", "Tangerang Selatan"},
		{"Gedung Sate", "wisata", "Jl. Diponegoro No.22, Cibeunying Kaler, Bandung", -6.9025, 107.6186, "Landmark pemerintahan bergaya kolonial.", "", "Bandung"},
{"Museum Konferensi Asia Afrika", "wisata", "Jl. Asia Afrika No.65, Braga, Bandung", -6.922, 107.606, "Museum sejarah politik Asia-Afrika.", "", "Bandung"},
{"Saung Angklung Udjo", "wisata", "Jl. Padasuka No.118, Cibeunying Kidul, Bandung", -6.903, 107.654, "Pertunjukan angklung tradisional.", "", "Bandung"},
{"Batagor Kingsley", "kuliner", "Jl. Veteran No.25, Sumur Bandung, Bandung", -6.919, 107.61, "Batagor populer lokal.", "", "Bandung"},
{"Surabi Enhaii", "kuliner", "Jl. Setiabudi No.137-139, Bandung", -6.901, 107.619, "Surabi khas Bandung.", "", "Bandung"},
{"Mie Kocok Bandung", "kuliner", "Jl. KH Ahmad Dahlan No.67, Bandung", -6.915, 107.61, "Mie kocok legendaris.", "", "Bandung"},

{"Kebun Raya Bogor", "wisata", "Jl. Ir. H. Juanda No.13, Bogor Tengah, Bogor", -6.5969, 106.8055, "Kebun botani bersejarah.", "", "Bogor"},
{"Istana Bogor", "wisata", "Jl. Ir. H. Juanda No.1, Bogor Tengah, Bogor", -6.5965, 106.806, "Istana negara dengan taman luas.", "", "Bogor"},
{"Museum Zoologi", "wisata", "Jl. Ir. H. Juanda No.9, Bogor Tengah, Bogor", -6.5958, 106.8065, "Koleksi fauna Indonesia.", "", "Bogor"},
{"Toge Goreng Bogor", "kuliner", "Jl. Surya Kencana, Bogor Tengah, Bogor", -6.595, 106.8166, "Kuliner khas Bogor.", "", "Bogor"},
{"Asinan Bogor", "kuliner", "Jl. Surya Kencana No.303, Bogor Tengah, Bogor", -6.597, 106.816, "Asinan sayur buah legendaris.", "", "Bogor"},
{"Roti Unyil", "kuliner", "Jl. Pajajaran No.1, Bogor Tengah, Bogor", -6.5955, 106.8162, "Roti khas kota.", "", "Bogor"},

{"Bekasi Cyber Park", "wisata", "Jl. KH. Noer Ali, Bekasi Selatan, Bekasi", -6.234, 106.989, "Taman kota modern.", "", "Bekasi"},
{"Museum Bekasi", "wisata", "Jl. Ir. H. Juanda, Bekasi Timur, Bekasi", -6.24, 106.99, "Museum sejarah lokal.", "", "Bekasi"},
{"Rumah Sakit Pusaka Bekasi (kultural spot)", "wisata", "Jl. Mayor Oking, Bekasi Timur, Bekasi", -6.241, 106.991, "Spot komunitas dan budaya.", "", "Bekasi"},
{"Sate Maranggi Bekasi", "kuliner", "Jl. Ir. H. Juanda, Bekasi Timur, Bekasi", -6.238, 106.99, "Sate maranggi lokal.", "", "Bekasi"},
{"Nasi Uduk Khas Bekasi", "kuliner", "Jl. KH. Agus Salim, Bekasi Timur, Bekasi", -6.239, 106.9895, "Nasi uduk populer.", "", "Bekasi"},
{"Kupat Tahu Bekasi", "kuliner", "Jl. M. Hasibuan, Bekasi Timur, Bekasi", -6.2375, 106.9892, "Kupat tahu legendaris.", "", "Bekasi"},

{"Kota Depok Lama", "wisata", "Jl. Pemuda, Pancoran Mas, Depok", -6.396, 106.82, "Kawasan bersejarah Depok.", "", "Depok"},
{"Taman Wiladatika", "wisata", "Jl. Jambore No.1, Cimanggis, Depok", -6.354, 106.902, "Taman rekreasi dan budaya.", "", "Depok"},
{"Museum Pengasingan Depok", "wisata", "Jl. Pemuda, Pancoran Mas, Depok", -6.395, 106.819, "Sejarah lokal.", "", "Depok"},
{"Soto Depok", "kuliner", "Jl. Margonda Raya, Depok", -6.4, 106.794, "Soto khas lokal.", "", "Depok"},
{"Es Dawet Depok", "kuliner", "Jl. Margonda Raya, Depok", -6.401, 106.795, "Minuman tradisional.", "", "Depok"},
{"Nasi Uduk Kampung Depok", "kuliner", "Jl. Sawangan, Pancoran Mas, Depok", -6.402, 106.7945, "Nasi uduk rumahan.", "", "Depok"},

{"Alun-Alun Cimahi", "wisata", "Jl. Raden Demang Hardjakusumah, Cimahi Tengah, Cimahi", -6.9, 107.545, "Pusat kegiatan kota.", "", "Cimahi"},
{"Museum Mandalawangi", "wisata", "Jl. Mandalawangi, Cimahi", -6.902, 107.544, "Museum lokal.", "", "Cimahi"},
{"Masjid Agung Cimahi", "wisata", "Jl. Rd. Demang Hardjakusumah, Cimahi Tengah, Cimahi", -6.901, 107.543, "Sentra religi.", "", "Cimahi"},
{"Sate Maranggi Cimahi", "kuliner", "Jl. Gatot Subroto, Cimahi", -6.891, 107.545, "Sate khas lokal.", "", "Cimahi"},
{"Mie Ayam Cimahi", "kuliner", "Jl. Gatot Subroto, Cimahi", -6.892, 107.546, "Mie ayam populer.", "", "Cimahi"},
{"Kupat Tahu Cimahi", "kuliner", "Jl. Rd. Demang Hardjakusumah, Cimahi Tengah, Cimahi", -6.893, 107.547, "Kupat tahu lokal.", "", "Cimahi"},

{"Pantai Pelabuhan Ratu (kultur area)", "wisata", "Pelabuhan Ratu, Sukabumi", -6.957, 106.971, "Pantai dan budaya pesisir.", "", "Sukabumi"},
{"Museum Prasasti Sukabumi", "wisata", "Jl. Bhayangkara, Sukabumi", -6.926, 106.934, "Sejarah lokal.", "", "Sukabumi"},
{"Cikundul Cultural Park", "wisata", "Jl. Cikundul, Lembursitu, Sukabumi", -6.93, 106.925, "Taman budaya.", "", "Sukabumi"},
{"Nasi Liwet Sukabumi", "kuliner", "Jl. Ahmad Yani, Sukabumi", -6.922, 106.925, "Masakan khas lokal.", "", "Sukabumi"},
{"Sate Maranggi Sukabumi", "kuliner", "Jl. Ahmad Yani, Sukabumi", -6.923, 106.926, "Sate lokal.", "", "Sukabumi"},
{"Keripik Singkong Sukabumi", "kuliner", "Jl. Ahmad Yani, Sukabumi", -6.924, 106.927, "Oleh-oleh khas.", "", "Sukabumi"},

{"Museum Asia Afrika (Tasik variant)", "wisata", "Jl. KHZ Mustofa, Tasikmalaya", -7.325, 108.219, "Spot sejarah regional.", "", "Tasikmalaya"},
{"Alun-Alun Tasikmalaya", "wisata", "Jl. KHZ Mustofa, Cihideung, Tasikmalaya", -7.326, 108.222, "Pusat kota dan acara budaya.", "", "Tasikmalaya"},
{"Cagar Budaya Tasik", "wisata", "Jl. Sutisna Senjaya, Tasikmalaya", -7.3275, 108.2205, "Situs bersejarah.", "", "Tasikmalaya"},
{"Soto tauco Tasik", "kuliner", "Jl. Tentara Pelajar, Tasikmalaya", -7.328, 108.22, "Soto khas.", "", "Tasikmalaya"},
{"Tape Ketan Tasik", "kuliner", "Jl. Tentara Pelajar, Tasikmalaya", -7.329, 108.221, "Tape ketan lokal.", "", "Tasikmalaya"},
{"Kupat Tahu Tasik", "kuliner", "Jl. Tentara Pelajar, Tasikmalaya", -7.33, 108.222, "Kupat tahu populer.", "", "Tasikmalaya"},

{"Keraton Kasepuhan", "wisata", "Jl. Jagasatru No.36, Kesepuhan, Cirebon", -6.729, 108.551, "Keraton bersejarah.", "", "Cirebon"},
{"Gua Sunyaragi", "wisata", "Jl. Sunyaragi, Kesambi, Cirebon", -6.733, 108.554, "Kompleks taman bersejarah.", "", "Cirebon"},
{"Museum Cirebon", "wisata", "Jl. Yos Sudarso, Lemahwungkuk, Cirebon", -6.73, 108.552, "Sejarah lokal dan budaya.", "", "Cirebon"},
{"Empal Gentong", "kuliner", "Jl. Slamet Riyadi, Cirebon", -6.726, 108.552, "Masakan khas Cirebon.", "", "Cirebon"},
{"Nasi Jamblang", "kuliner", "Jl. Cangkring II, Cirebon", -6.727, 108.5515, "Kuliner unik dengan daun jati.", "", "Cirebon"},
{"Dawet Ayu Cirebon", "kuliner", "Jl. Pasuketan, Cirebon", -6.725, 108.553, "Minuman tradisional.", "", "Cirebon"},

{"Alun-Alun Banjar", "wisata", "Jl. Letnan Harun, Banjar", -6.9736, 108.4478, "Pusat kegiatan budaya.", "", "Banjar"},
{"Museum Banjar", "wisata", "Jl. Letnan Harun, Banjar", -6.974, 108.448, "Sejarah setempat.", "", "Banjar"},
{"Masjid Agung Banjar", "wisata", "Jl. Letnan Harun, Banjar", -6.973, 108.447, "Sentra religi lokal.", "", "Banjar"},
{"Soto Banjar", "kuliner", "Jl. Letnan Harun, Banjar", -6.9735, 108.4475, "Soto khas Banjar.", "", "Banjar"},
{"Kue Bagea Banjar", "kuliner", "Jl. Letnan Harun, Banjar", -6.9732, 108.4472, "Kue tradisional.", "", "Banjar"},
{"Lontong Kari Banjar", "kuliner", "Jl. Letnan Harun, Banjar", -6.9738, 108.4482, "Lontong khas.", "", "Banjar"},

{"Lawang Sewu", "wisata", "Jl. Pemuda No.160, Semarang Tengah, Semarang", -6.9806, 110.4146, "Bangunan kolonial ikonik.", "", "Semarang"},
{"Kota Lama Semarang", "wisata", "Jl. Letjen Suprapto, Semarang Utara, Semarang", -6.9735, 110.414, "Area bersejarah bergaya kolonial.", "", "Semarang"},
{"Sam Poo Kong", "wisata", "Jl. Simongan No.129, Semarang Barat, Semarang", -6.9674, 110.4152, "Klenteng bersejarah.", "", "Semarang"},
{"Tahu Gimbal Pak Karmin", "kuliner", "Jl. Pandanaran II No.2, Semarang Tengah, Semarang", -6.973, 110.414, "Tahu gimbal khas Semarang.", "", "Semarang"},
{"Lumpia Semarang", "kuliner", "Jl. Pandanaran No.57, Semarang Tengah, Semarang", -6.9732, 110.4142, "Lumpia khas kota.", "", "Semarang"},
{"Nasi Gandul", "kuliner", "Jl. Anjasmoro Raya, Semarang Barat, Semarang", -6.97, 110.42, "Nasi gulai tradisional.", "", "Semarang"},

		{"Keraton Surakarta", "wisata", "Jl. Sidikoro, Baluwarti, Pasar Kliwon, Surakarta", -7.5678, 110.831, "Istana tradisional Keraton.", "", "Surakarta"},
{"Pasar Klewer", "wisata", "Jl. DR. Radjiman, Kauman, Pasar Kliwon, Surakarta", -7.565, 110.8315, "Pusat batik tradisional.", "", "Surakarta"},
{"Taman Balekambang", "wisata", "Jl. Ahmad Yani, Manahan, Banjarsari, Surakarta", -7.566, 110.828, "Taman kota dengan event seni.", "", "Surakarta"},
{"Nasi Liwet Wongso", "kuliner", "Jl. Teuku Umar, Keprabon, Banjarsari, Surakarta", -7.5665, 110.8312, "Nasi liwet khas Solo.", "", "Surakarta"},
{"Serabi Notosuman", "kuliner", "Jl. Mohammad Yamin No.28, Jayengan, Serengan, Surakarta", -7.5659, 110.8318, "Serabi tradisional.", "", "Surakarta"},
{"Timlo Solo", "kuliner", "Jl. Jend. Urip Sumoharjo No.94, Purwodiningratan, Jebres, Surakarta", -7.5652, 110.8305, "Sup khas Solo.", "", "Surakarta"},

{"Alun-Alun Salatiga", "wisata", "Jl. Sukowati, Kalicacing, Sidomukti, Salatiga", -7.3306, 110.5089, "Pusat kota untuk acara budaya.", "", "Salatiga"},
{"Museum Salatiga", "wisata", "Jl. Diponegoro No.1, Kalicacing, Salatiga", -7.331, 110.509, "Sejarah lokal.", "", "Salatiga"},
{"Gereja Blenduk (dekat)", "wisata", "Jl. Jenderal Sudirman, Salatiga", -7.329, 110.5095, "Arsitektur kolonial sekitar.", "", "Salatiga"},
{"Es Dawet Salatiga", "kuliner", "Jl. Sukowati, Kalicacing, Sidomukti, Salatiga", -7.33, 110.508, "Dawet tradisional.", "", "Salatiga"},
{"Sate Ayam Salatiga", "kuliner", "Jl. Jenderal Sudirman, Salatiga", -7.3315, 110.5092, "Sate lokal.", "", "Salatiga"},
{"Kue Lontong Salatiga", "kuliner", "Jl. Diponegoro, Kalicacing, Sidomukti, Salatiga", -7.332, 110.51, "Kue tradisional.", "", "Salatiga"},

{"Candi Borobudur", "wisata", "Jl. Badrawati, Borobudur, Magelang", -7.6079, 110.2038, "Candi Buddha terbesar di dunia.", "", "Magelang"},
{"Alun-Alun Magelang", "wisata", "Jl. Alun-Alun Utara, Magelang Tengah, Magelang", -7.4682, 110.217, "Pusat kota dan acara.", "", "Magelang"},
{"Museum Kartini", "wisata", "Jl. Ade Irma Suryani, Magelang Tengah, Magelang", -7.467, 110.218, "Koleksi sejarah lokal.", "", "Magelang"},
{"Tahu Aci Magelang", "kuliner", "Jl. Pemuda, Magelang Tengah, Magelang", -7.468, 110.2175, "Tahu khas setempat.", "", "Magelang"},
{"Sate Kelinci Magelang", "kuliner", "Jl. Tidar, Magelang Selatan, Magelang", -7.469, 110.219, "Kuliner tradisional.", "", "Magelang"},
{"Bakpia Magelang", "kuliner", "Jl. Mayor Jenderal Sutoyo, Magelang Tengah, Magelang", -7.466, 110.216, "Bakpia khas area Borobudur.", "", "Magelang"},

{"Museum Batik Pekalongan", "wisata", "Jl. Jetayu No.1, Pekalongan Barat, Pekalongan", -6.883, 109.6665, "Pusat batik dan edukasi.", "", "Pekalongan"},
{"Alun-Alun Pekalongan", "wisata", "Jl. Alun-Alun Barat, Pekalongan", -6.884, 109.667, "Pusat acara budaya.", "", "Pekalongan"},
{"Kampung Batik", "wisata", "Jl. Hayam Wuruk, Pekalongan", -6.885, 109.668, "Sentra pembuatan batik.", "", "Pekalongan"},
{"Lentho Pekalongan", "kuliner", "Jl. Dr. Cipto, Pekalongan Barat, Pekalongan", -6.8835, 109.6675, "Makanan tradisional.", "", "Pekalongan"},
{"Pindang Lemah Wangi", "kuliner", "Jl. Sultan Agung, Pekalongan", -6.8845, 109.6685, "Ikan pindang khas.", "", "Pekalongan"},
{"Sego Megono", "kuliner", "Jl. Sultan Agung, Pekalongan", -6.8855, 109.6695, "Nasi megono khas.", "", "Pekalongan"},

{"Alun-Alun Tegal", "wisata", "Jl. Pancasila, Tegal Barat, Tegal", -6.8698, 109.1362, "Pusat kota.", "", "Tegal"},
{"Keraton Kacirebonan (dekat)", "wisata", "Jl. Sultan Agung, Tegal", -6.8705, 109.137, "Sejarah lokal.", "", "Tegal"},
{"Museum Negeri Slawi (dekat)", "wisata", "Jl. Ahmad Yani, Slawi, Tegal", -6.871, 109.138, "Koleksi budaya lokal.", "", "Tegal"},
{"Gado-Gado Tegal", "kuliner", "Jl. Pancasila, Tegal Barat, Tegal", -6.87, 109.1365, "Gado-gado khas.", "", "Tegal"},
{"Tahu Petis Tegal", "kuliner", "Jl. Mayjen Sutoyo, Tegal", -6.8712, 109.1375, "Tahu petis lokal.", "", "Tegal"},
{"Es Durian Tegal", "kuliner", "Jl. Ahmad Yani, Tegal", -6.872, 109.1385, "Dessert populer.", "", "Tegal"},
{"Keraton Yogyakarta", "wisata", "Jl. Rotowijayan Blok No.1, Kraton, Yogyakarta", -7.805, 110.364, "Istana Sultan Yogyakarta.", "", "Yogyakarta"},
{"Taman Sari", "wisata", "Jl. Tamanan, Kraton, Yogyakarta", -7.807, 110.36, "Benteng air istana.", "", "Yogyakarta"},
{"Museum Sonobudoyo", "wisata", "Jl. Trikora No.6, Gondomanan, Yogyakarta", -7.802, 110.36, "Koleksi seni dan budaya.", "", "Yogyakarta"},
{"Gudeg Yu Djum", "kuliner", "Jl. Wijilan No.167, Kraton, Yogyakarta", -7.801, 110.367, "Gudeg terkenal.", "", "Yogyakarta"},
{"Bakpia Pathok", "kuliner", "Jl. KS Tubun No.75, Ngampilan, Yogyakarta", -7.793, 110.365, "Bakpia khas Yogyakarta.", "", "Yogyakarta"},
{"Sate Klathak Pak Pong", "kuliner", "Jl. Sultan Agung No.18, Jejeran, Bantul, Yogyakarta", -7.839, 110.408, "Sate khas Bantul.", "", "Yogyakarta"},

{"Tugu Pahlawan", "wisata", "Jl. Pahlawan, Alun-alun Contong, Surabaya", -7.245, 112.737, "Monumen pahlawan nasional.", "", "Surabaya"},
{"House of Sampoerna", "wisata", "Jl. Taman Sampoerna No.6, Krembangan Utara, Surabaya", -7.226, 112.739, "Museum rokok dan sejarah.", "", "Surabaya"},
{"Museum 10 November", "wisata", "Jl. Pahlawan, Alun-alun Contong, Surabaya", -7.2455, 112.7375, "Sejarah perjuangan.", "", "Surabaya"},
{"Rawon Setan", "kuliner", "Jl. Embong Malang No.78/I, Surabaya", -7.268, 112.745, "Rawon khas Surabaya.", "", "Surabaya"},
{"Lontong Balap Pak Gendut", "kuliner", "Jl. Kranggan No.60, Sawahan, Surabaya", -7.257, 112.73, "Lontong balap legendaris.", "", "Surabaya"},
{"Tahu Tek Pak Jayen", "kuliner", "Jl. Dharmahusada No.112, Surabaya", -7.266, 112.736, "Tahu tek populer.", "", "Surabaya"},

{"Candi Singosari", "wisata", "Jl. Kertanegara, Singosari, Malang", -7.892, 112.665, "Candi peninggalan kerajaan Singosari.", "", "Malang"},
{"Museum Malang Tempo Doeloe", "wisata", "Jl. Gajahmada No.2, Klojen, Malang", -7.978, 112.633, "Sejarah kota Malang.", "", "Malang"},
{"Alun-Alun Malang", "wisata", "Jl. Merdeka Selatan, Klojen, Malang", -7.982, 112.63, "Pusat kegiatan kota.", "", "Malang"},
{"Bakso President", "kuliner", "Jl. Batanghari No.5, Klojen, Malang", -7.967, 112.632, "Bakso legendaris Malang.", "", "Malang"},
{"Rawon Nguling", "kuliner", "Jl. Zainul Arifin No.62, Klojen, Malang", -7.983, 112.632, "Rawon khas.", "", "Malang"},
{"Pos Ketan Legenda Batu", "kuliner", "Jl. Kartini No.6, Batu", -7.872, 112.524, "Ketan khas Batu.", "", "Batu"},

{"Jatim Park 2", "wisata", "Jl. Oro-Oro Ombo No.9, Batu", -7.882, 112.524, "Taman wisata dan edukasi.", "", "Batu"},
{"Museum Angkut", "wisata", "Jl. Terusan Sultan Agung No.2, Batu", -7.878, 112.531, "Museum transportasi.", "", "Batu"},
{"Alun-Alun Batu", "wisata", "Jl. Diponegoro, Batu", -7.873, 112.523, "Pusat kota.", "", "Batu"},
{"Pos Ketan Legenda", "kuliner", "Jl. Kartini No.6, Batu", -7.872, 112.524, "Ketan legendaris.", "", "Batu"},
{"Bakso President (Batu branch)", "kuliner", "Jl. Diponegoro No.10, Batu", -7.873, 112.525, "Bakso populer.", "", "Batu"},
{"Apple Market Batu", "kuliner", "Jl. Brantas, Batu", -7.874, 112.526, "Apple and local snacks.", "", "Batu"},

{"Candi Bajang Ratu", "wisata", "Jl. Candi Bajang Ratu, Trowulan, Mojokerto", -7.488, 112.434, "Gapura Majapahit bersejarah.", "", "Mojokerto"},
{"Trowulan (kawasan cagar budaya)", "wisata", "Jl. Raya Trowulan, Mojokerto", -7.557, 112.51, "Situs purbakala Majapahit.", "", "Mojokerto"},
{"Museum Majapahit", "wisata", "Jl. Raya Trowulan, Mojokerto", -7.475, 112.435, "Museum sejarah Majapahit.", "", "Mojokerto"},
{"Nasi Pecel Mojokerto", "kuliner", "Jl. Majapahit, Mojokerto", -7.474, 112.436, "Nasi pecel lokal.", "", "Mojokerto"},
{"Sate Klothok", "kuliner", "Jl. Benteng Pancasila, Mojokerto", -7.4755, 112.437, "Sate tradisional.", "", "Mojokerto"},

		{"Getuk Lindri Mojokerto", "kuliner", "Jl. Majapahit, Mojokerto", -7.476, 112.438, "Kudapan tradisional.", "", "Mojokerto"},

{"Alun-Alun Madiun", "wisata", "Jl. Pahlawan, Madiun", -7.6309, 111.5239, "Pusat kegiatan kota.", "", "Madiun"},
{"Stasiun Madiun (heritage spot)", "wisata", "Jl. Kompol Sunaryo No.6, Madiun", -7.634, 111.516, "Stasiun bersejarah.", "", "Madiun"},
{"Museum Madiun", "wisata", "Jl. Pahlawan, Madiun", -7.631, 111.524, "Sejarah lokal.", "", "Madiun"},
{"Pecel Madiun", "kuliner", "Jl. Cokroaminoto, Madiun", -7.63, 111.523, "Pecel khas Madiun.", "", "Madiun"},
{"Rujak Soto Madiun", "kuliner", "Jl. HOS Cokroaminoto, Madiun", -7.631, 111.525, "Kombinasi khas lokal.", "", "Madiun"},
{"Oleh-Oleh Madiun", "kuliner", "Jl. Panglima Sudirman, Madiun", -7.632, 111.526, "Snacks and local sweets.", "", "Madiun"},

{"Candi Tegowangi (dekat)", "wisata", "Jl. Candi Tegowangi, Plemahan, Kediri", -7.869, 111.997, "Situs candi lokal.", "", "Kediri"},
{"Alun-Alun Kediri", "wisata", "Jl. Panglima Sudirman, Kediri", -7.848, 112.008, "Pusat kota.", "", "Kediri"},
{"Museum Airlangga (dekat)", "wisata", "Jl. Airlangga, Kediri", -7.85, 112.01, "Sejarah lokal.", "", "Kediri"},
{"Pecel Kediri", "kuliner", "Jl. Doho, Kediri", -7.845, 112.001, "Pecel khas.", "", "Kediri"},
{"Rujak Cingur Kediri", "kuliner", "Jl. Joyoboyo, Kediri", -7.846, 112.002, "Rujak cingur khas Jawa Timur.", "", "Kediri"},
{"Lontong Balap Kediri", "kuliner", "Jl. Pattimura, Kediri", -7.847, 112.003, "Lontong balap lokal.", "", "Kediri"},

{"Makam Bung Karno", "wisata", "Jl. Ir. Soekarno, Bendogerit, Blitar", -8.0984, 112.1554, "Makam Presiden Soekarno.", "", "Blitar"},
{"Candi Penataran (dekat)", "wisata", "Jl. Penataran, Nglegok, Blitar", -8.0815, 112.148, "Kompleks candi besar.", "", "Blitar"},
{"Museum Bung Karno", "wisata", "Jl. Kalasan, Blitar", -8.098, 112.155, "Sejarah presiden pertama Indonesia.", "", "Blitar"},
{"Pecel Blitar", "kuliner", "Jl. Merdeka, Blitar", -8.072, 112.147, "Pecel khas Blitar.", "", "Blitar"},
{"Nasi Tiwul", "kuliner", "Jl. Kenari, Blitar", -8.073, 112.148, "Makanan tradisional.", "", "Blitar"},
{"Oleh-Oleh Blitar", "kuliner", "Jl. Kartini, Blitar", -8.074, 112.149, "Snacks and souvenirs.", "", "Blitar"},

{"Pelabuhan Probolinggo", "wisata", "Jl. Pelabuhan, Probolinggo", -7.7657, 113.217, "Pelabuhan dan area sejarah.", "", "Probolinggo"},
{"Museum Probolinggo", "wisata", "Jl. Suroyo, Probolinggo", -7.766, 113.218, "Sejarah lokal.", "", "Probolinggo"},
{"Alun-Alun Probolinggo", "wisata", "Jl. Suroyo, Probolinggo", -7.767, 113.219, "Pusat kota.", "", "Probolinggo"},
{"Rujak Soto Probolinggo", "kuliner", "Jl. Suroyo, Probolinggo", -7.768, 113.22, "Kombinasi kuliner khas.", "", "Probolinggo"},
{"Ikan Bakar Pantai", "kuliner", "Jl. Pantai Permata, Probolinggo", -7.769, 113.221, "Seafood lokal.", "", "Probolinggo"},
{"Sate Madura (lokal)", "kuliner", "Jl. Suroyo, Probolinggo", -7.77, 113.222, "Sate khas Jawa Timur.", "", "Probolinggo"},

{"Alun-Alun Pasuruan", "wisata", "Jl. Alun-Alun Utara, Pasuruan", -7.6426, 112.9136, "Pusat kegiatan kota.", "", "Pasuruan"},
{"Museum Pasuruan (lokal)", "wisata", "Jl. Sultan Agung, Pasuruan", -7.643, 112.914, "Sejarah lokal.", "", "Pasuruan"},
{"Kawasan Panggung (kultur)", "wisata", "Jl. Panggungrejo, Pasuruan", -7.644, 112.915, "Panggung seni dan budaya.", "", "Pasuruan"},
{"Lontong Balap Pasuruan", "kuliner", "Jl. Gajah Mada, Pasuruan", -7.645, 112.916, "Lontong balap khas.", "", "Pasuruan"},
{"Nasi Campur Pasuruan", "kuliner", "Jl. Hayam Wuruk, Pasuruan", -7.646, 112.917, "Nasi campur lokal.", "", "Pasuruan"},
{"Pesmol Ikan", "kuliner", "Jl. Hasanuddin, Pasuruan", -7.647, 112.918, "Masakan tradisional.", "", "Pasuruan"},

{"Alun-Alun Salatiga", "wisata", "Jl. Jenderal Sudirman, Sidorejo, Salatiga", -7.3306, 110.5089, "Pusat kota untuk acara budaya.", "", "Salatiga"},
{"Museum Salatiga", "wisata", "Jl. Diponegoro No.23, Sidorejo, Salatiga", -7.331, 110.509, "Sejarah lokal.", "", "Salatiga"},
{"Gereja Blenduk (Salatiga)", "wisata", "Jl. Jend. Sudirman, Sidorejo, Salatiga", -7.329, 110.5095, "Arsitektur kolonial sekitar.", "", "Salatiga"},
{"Es Dawet Salatiga", "kuliner", "Jl. Jend. Sudirman, Salatiga", -7.33, 110.508, "Dawet tradisional.", "", "Salatiga"},
{"Sate Ayam Salatiga", "kuliner", "Jl. Diponegoro, Sidorejo, Salatiga", -7.3315, 110.5092, "Sate lokal.", "", "Salatiga"},
{"Kue Lontong Salatiga", "kuliner", "Jl. Jend. Sudirman, Salatiga", -7.332, 110.51, "Kue tradisional.", "", "Salatiga"},

{"Candi Borobudur", "wisata", "Jl. Badrawati, Borobudur, Magelang", -7.6079, 110.2038, "Candi Buddha terbesar di dunia.", "", "Magelang"},
{"Alun-Alun Magelang", "wisata", "Jl. Alun-Alun Utara, Magelang", -7.4682, 110.217, "Pusat kota dan acara.", "", "Magelang"},
{"Museum Kartini Magelang", "wisata", "Jl. Jalanan No.12, Magelang", -7.467, 110.218, "Koleksi sejarah lokal.", "", "Magelang"},
{"Tahu Aci Magelang", "kuliner", "Jl. Pemuda, Magelang", -7.468, 110.2175, "Tahu khas setempat.", "", "Magelang"},
{"Sate Kelinci Magelang", "kuliner", "Jl. Diponegoro, Magelang", -7.469, 110.219, "Kuliner tradisional.", "", "Magelang"},
{"Bakpia Magelang", "kuliner", "Jl. Borobudur, Magelang", -7.466, 110.216, "Bakpia khas area Borobudur.", "", "Magelang"},

{"Museum Batik Pekalongan", "wisata", "Jl. Jetayu No.1, Pekalongan", -6.883, 109.6665, "Pusat batik dan edukasi.", "", "Pekalongan"},
{"Alun-Alun Pekalongan", "wisata", "Jl. Alun-Alun Utara, Pekalongan", -6.884, 109.667, "Pusat acara budaya.", "", "Pekalongan"},
{"Kampung Batik Kauman", "wisata", "Jl. Hayam Wuruk, Pekalongan", -6.885, 109.668, "Sentra pembuatan batik.", "", "Pekalongan"},
{"Lentho Pekalongan", "kuliner", "Jl. Jetayu, Pekalongan", -6.8835, 109.6675, "Makanan tradisional.", "", "Pekalongan"},
{"Pindang Lemah Wangi", "kuliner", "Jl. Dr. Cipto, Pekalongan", -6.8845, 109.6685, "Ikan pindang khas.", "", "Pekalongan"},
{"Sego Megono", "kuliner", "Jl. Sultan Agung, Pekalongan", -6.8855, 109.6695, "Nasi megono khas.", "", "Pekalongan"},

{"Alun-Alun Tegal", "wisata", "Jl. Veteran, Tegal", -6.8698, 109.1362, "Pusat kota.", "", "Tegal"},
{"Keraton Kacirebonan (Tegal)", "wisata", "Jl. Keraton, Tegal", -6.8705, 109.137, "Sejarah lokal.", "", "Tegal"},
{"Museum Negeri Slawi", "wisata", "Jl. Ahmad Yani, Slawi, Tegal", -6.871, 109.138, "Koleksi budaya lokal.", "", "Tegal"},
{"Gado-Gado Tegal", "kuliner", "Jl. Ahmad Yani, Tegal", -6.87, 109.1365, "Gado-gado khas.", "", "Tegal"},
{"Tahu Petis Tegal", "kuliner", "Jl. Veteran, Tegal", -6.8712, 109.1375, "Tahu petis lokal.", "", "Tegal"},
{"Es Durian Tegal", "kuliner", "Jl. Diponegoro, Tegal", -6.872, 109.1385, "Dessert populer.", "", "Tegal"},

{"Tugu Pahlawan", "wisata", "Jl. Pahlawan, Surabaya", -7.245, 112.737, "Monumen pahlawan nasional.", "", "Surabaya"},
{"House of Sampoerna", "wisata", "Jl. Taman Sampoerna No.6, Surabaya", -7.226, 112.739, "Museum rokok dan sejarah.", "", "Surabaya"},
{"Museum 10 November", "wisata", "Jl. Pahlawan No.1, Surabaya", -7.2455, 112.7375, "Sejarah perjuangan.", "", "Surabaya"},
{"Rawon Setan", "kuliner", "Jl. Embong Malang No.78, Surabaya", -7.268, 112.745, "Rawon khas Surabaya.", "", "Surabaya"},
{"Lontong Balap Pak Gendut", "kuliner", "Jl. Kranggan No.60, Surabaya", -7.257, 112.73, "Lontong balap legendaris.", "", "Surabaya"},
{"Tahu Tek Pak Jayen", "kuliner", "Jl. Dharmawangsa No.2, Surabaya", -7.266, 112.736, "Tahu tek populer.", "", "Surabaya"},

	
}

	for _, p := range placesData {
		// Cari CityID dari map
		cityID, ok := cityMap[p.NamaKota]
		if !ok {
			log.Printf("⚠️ Kota %s tidak ditemukan di DB, skip %s", p.NamaKota, p.NamaTempat)
			continue
		}

		place := models.Place{
			NamaTempat: p.NamaTempat,
			Kategori:   p.Kategori,
			Alamat:     p.Alamat,
			Latitude:   p.Latitude,
			Longitude:  p.Longitude,
			Deskripsi:  p.Deskripsi,
			ImageURL:   p.ImageURL,
			CityID:     cityID,
		}

		// Cek dulu apakah sudah ada
		var existing models.Place
		if err := DB.Where("nama_tempat = ?", place.NamaTempat).First(&existing).Error; err == nil {
			log.Printf("ℹ️ Tempat %s sudah ada, skip", place.NamaTempat)
			continue
		}

		if err := DB.Create(&place).Error; err != nil {
			log.Printf("❌ Gagal seed %s: %v", place.NamaTempat, err)
		} else {
			log.Printf("✅ Berhasil seed %s", place.NamaTempat)
		}
	}
}

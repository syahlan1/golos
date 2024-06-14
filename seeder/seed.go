package seeder

import (
	"os"

	"github.com/syahlan1/golos/connection"
	"github.com/syahlan1/golos/models"
	"golang.org/x/crypto/bcrypt"
)

func UserSeed() {
	var adminPassword = os.Getenv("ADMIN_PASSWORD")

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(adminPassword), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	var users = []models.Users{
		{
			Username: os.Getenv("ADMIN_USERNAME"),
			Email:    os.Getenv("ADMIN_EMAIL"),
			Password: hashedPassword,
			IsLogin:  0,
			RoleId:   1,
		},
	}

	var roles = []models.Roles{
		{
			Name:        "Superadmin",
			Description: "High Tier Admin",
		},
	}

	var permission = []models.Permission{
		{
			Name: "create",
		},
		{
			Name: "update",
		},
		{
			Name: "delete",
		},
		{
			Name: "approve",
		},
		{
			Name: "create_role",
		},
		{
			Name: "delete_role",
		},
	}

	var rolePermissions = []models.RolePermission{
		{
			RolesId:      1,
			PermissionId: 1,
		},
		{
			RolesId:      1,
			PermissionId: 2,
		},
		{
			RolesId:      1,
			PermissionId: 3,
		},
		{
			RolesId:      1,
			PermissionId: 4,
		},
		{
			RolesId:      1,
			PermissionId: 5,
		},
		{
			RolesId:      1,
			PermissionId: 6,
		},
	}

	connection.DB.Save(&permission)
	connection.DB.Save(&roles)
	connection.DB.Save(&users)
	connection.DB.Save(&rolePermissions)
}

func GeneralInformationSeed() {

	var cabang = []models.Cabang{

		{
			BranchCode:      "01BJI",
			Name:            "KC Jakarta Sudirman",
			CBCCode:         "00BJI",
			Address:         "Gedung Sahid Sudirman Center Lt. GF, Jl. Sudirman No. 86",
			Cabang:          true,
			CabangPencairan: true,
		},
		{
			BranchCode:      "02BJI",
			Name:            "KCP Jakarta RS Fatmawati",
			CBCCode:         "00BJI",
			Address:         "Jl. R.S. Fatmawati No. 22 B-D, Cipete Selatan, Cilandak",
			Cabang:          true,
			CabangPencairan: true,
		},
		{
			BranchCode:      "03BJI",
			Name:            "KCP Jakarta Pondok Indah",
			CBCCode:         "00BJI",
			Address:         "Jl. Metro Pondok Indah Blok UA No. 71",
			Cabang:          true,
			CabangPencairan: true,
		},
		{
			BranchCode:      "04BJI",
			Name:            "KCP Jakarta Kemang",
			CBCCode:         "00BJI",
			Address:         "Jl. Kemang Selatan Raya No. 111 H, Mampang Prapatan",
			Cabang:          true,
			CabangPencairan: true,
		},
		{
			BranchCode:      "05BJI",
			Name:            "KCP Jakarta Tanah Abang",
			CBCCode:         "00BJI",
			Address:         "Komplek Pertokoan Tanah Abang, Bukit Blok F No. 16-17",
			Cabang:          true,
			CabangPencairan: true,
		},
		{
			BranchCode:      "06BJI",
			Name:            "KCP Jakarta Gajahmada",
			CBCCode:         "00BJI",
			Address:         "Jalan Gajah Mada No. 11 A-B Jakarta Pusat ",
			Cabang:          true,
			CabangPencairan: true,
		},
		{
			BranchCode:      "07BJI",
			Name:            "KCP Jakarta Orion Dusit",
			CBCCode:         "00BJI",
			Address:         "Jalan Mangga Dua Raya No. 1.06",
			Cabang:          true,
			CabangPencairan: true,
		},
		{
			BranchCode:      "08BJI",
			Name:            "KCP Jakarta Klender",
			CBCCode:         "00BJI",
			Address:         "Buaran Plaza Lantai Dasar No. 8-10, Jl. Raden Inten No. 1, Buaran, Klender",
			Cabang:          true,
			CabangPencairan: true,
		},
		{
			BranchCode:      "09BJI",
			Name:            "KCP Jakarta  AEON Mall ",
			CBCCode:         "00BJI",
			Address:         "AEON Mall Jakarta Garden City Lt. GF Unit G - 68. Jl Boulevard Garden City, Cakung",
			Cabang:          true,
			CabangPencairan: true,
		},
		{
			BranchCode:      "10BJI",
			Name:            "KCP Jakarta Muara Karang",
			CBCCode:         "00BJI",
			Address:         "Jl. Muara Karang Raya Blok A 8 Utara No. 21",
			Cabang:          true,
			CabangPencairan: true,
		},
		{
			BranchCode:      "11BJI",
			Name:            "KCP Jakarta Kelapa Gading Boulevard",
			CBCCode:         "00BJI",
			Address:         "Jl. Boulevard Barat, Blok LC 6 Kavling No. 55, Kelapa Gading Permai",
			Cabang:          true,
			CabangPencairan: true,
		},
		{
			BranchCode:      "12BJI",
			Name:            "KCP Jakarta Puri Indah",
			CBCCode:         "00BJI",
			Address:         "Pasar Puri Indah Blok I No. 37, Jl. Puri Indah Raya",
			Cabang:          true,
			CabangPencairan: true,
		},
		{
			BranchCode:      "13BJI",
			Name:            "KCP Jakarta Pintu-Kecil",
			CBCCode:         "00BJI",
			Address:         "Jl. Pasar Pagi No. 101 A (dh No. 99), Roa Malaka, Tambora",
			Cabang:          true,
			CabangPencairan: true,
		},
		{
			BranchCode:      "14BJI",
			Name:            "KCP Jakarta Tomang",
			CBCCode:         "00BJI",
			Address:         "Gedung Graha Sukandamulia, Lt. 1 (Dasar), Jl. Tomang Raya Terusan Kav.71-72",
			Cabang:          true,
			CabangPencairan: true,
		},
		{
			BranchCode:      "15BJI",
			Name:            "KC Medan Putri Hijau",
			CBCCode:         "00BJI",
			Address:         "Jl. Putri Hijau No. 4 BC",
			Cabang:          true,
			CabangPencairan: true,
		},
		{
			BranchCode:      "16BJI",
			Name:            "KCP Medan Asia",
			CBCCode:         "00BJI",
			Address:         "Jl. Asia No. 172 C",
			Cabang:          true,
			CabangPencairan: true,
		},
		{
			BranchCode:      "17BJI",
			Name:            "KC Batam",
			CBCCode:         "00BJI",
			Address:         "Komplek Pertokoan Costa Rica Blok B1 No. 3 – 3A",
			Cabang:          true,
			CabangPencairan: true,
		},
		{
			BranchCode:      "18BJI",
			Name:            "KC Pekanbaru",
			CBCCode:         "00BJI",
			Address:         "Jl. Jendral Sudirman No. 150 A-B",
			Cabang:          true,
			CabangPencairan: true,
		},
		{
			BranchCode:      "19BJI",
			Name:            "KC Pangkal Pinang Sudirman",
			CBCCode:         "00BJI",
			Address:         "Jl. Jenderal Sudirman No. 30-32",
			Cabang:          true,
			CabangPencairan: true,
		},
		{
			BranchCode:      "20BJI",
			Name:            "KCP Sungai Liat Sudirman",
			CBCCode:         "00BJI",
			Address:         "Komplek Ruko Permata Indah, Blok A No. 1A-B, Jl. Jend. Sudirman",
			Cabang:          true,
			CabangPencairan: true,
		},
		{
			BranchCode:      "21BJI",
			Name:            "KC Jambi",
			CBCCode:         "00BJI",
			Address:         "Jl. Gatot Subroto No. 75, Jambi",
			Cabang:          true,
			CabangPencairan: true,
		},
		{
			BranchCode:      "22BJI",
			Name:            "KC Palembang Kebumen",
			CBCCode:         "00BJI",
			Address:         "Jl. Kebumen Darat No. 834",
			Cabang:          true,
			CabangPencairan: true,
		},
		{
			BranchCode:      "23BJI",
			Name:            "KCP Palembang Rajawali",
			CBCCode:         "00BJI",
			Address:         "Jl. Rajawali No. 1087 - 1088 Lantai 2",
			Cabang:          true,
			CabangPencairan: true,
		},
		{
			BranchCode:      "24BJI",
			Name:            "KC Bandar Lampung Sudirman",
			CBCCode:         "00BJI",
			Address:         "Jl.Jendral Sudirman No. 23 E – F Bandar Lampung",
			Cabang:          true,
			CabangPencairan: true,
		},
		{
			BranchCode:      "25BJI",
			Name:            "KC Karawang",
			CBCCode:         "00BJI",
			Address:         "Sentra KIIC LT. 1, Jl. Permata Raya Lot CA-1, KIIC, Karawang",
			Cabang:          true,
			CabangPencairan: true,
		},
		{
			BranchCode:      "26BJI",
			Name:            "KC Bandung",
			CBCCode:         "00BJI",
			Address:         "Jl. Ir. H. Juanda No. 28, Bandung",
			Cabang:          true,
			CabangPencairan: true,
		},
		{
			BranchCode:      "27BJI",
			Name:            "KC Bogor",
			CBCCode:         "00BJI",
			Address:         "Jl. Suryakencana No. 294-296",
			Cabang:          true,
			CabangPencairan: true,
		},
		{
			BranchCode:      "28BJI",
			Name:            "KC Cirebon Yos Sudarso",
			CBCCode:         "00BJI",
			Address:         "Jalan Yos Sudarso No. 15 D-F",
			Cabang:          true,
			CabangPencairan: true,
		},
		{
			BranchCode:      "29BJI",
			Name:            "KCP Bekasi",
			CBCCode:         "00BJI",
			Address:         "Grand Mall Bekasi Blok B No. 8, Jl. Jenderal Sudirman",
			Cabang:          true,
			CabangPencairan: true,
		},
		{
			BranchCode:      "30BJI",
			Name:            "KCP Cikarang",
			CBCCode:         "00BJI",
			Address:         "Hotel Holiday Inn Cikarang Jababeka Lt. 1, Jl. Jababeka Raya Kav. A-2 Jababeka 1, Cikarang",
			Cabang:          true,
			CabangPencairan: true,
		},
		{
			BranchCode:      "31BJI",
			Name:            "KCP Depok Margonda",
			CBCCode:         "00BJI",
			Address:         "Jl. Margonda Raya No. 252 D, Kel. Kemiri Muka, Kec. Beji",
			Cabang:          true,
			CabangPencairan: true,
		},
		{
			BranchCode:      "32BJI",
			Name:            "KCP Serpong",
			CBCCode:         "00BJI",
			Address:         "Jalur Sutera 29 D Nomor 39, Perumahan Alam Sutera, Serpong Utara",
			Cabang:          true,
			CabangPencairan: true,
		},
		{
			BranchCode:      "33BJI",
			Name:            "KC Semarang",
			CBCCode:         "00BJI",
			Address:         "Ruko pemuda mas blok A4, Jln. Pemuda No. 150",
			Cabang:          true,
			CabangPencairan: true,
		},
		{
			BranchCode:      "34BJI",
			Name:            "KC Solo",
			CBCCode:         "00BJI",
			Address:         "Jl. Slamet Riyadi No. 295, Laweyan",
			Cabang:          true,
			CabangPencairan: true,
		},
		{
			BranchCode:      "35BJI",
			Name:            "KC Yogyakarta",
			CBCCode:         "00BJI",
			Address:         "Jl. P. Diponegoro No. 9",
			Cabang:          true,
			CabangPencairan: true,
		},
		{
			BranchCode:      "36BJI",
			Name:            "KC Malang",
			CBCCode:         "00BJI",
			Address:         "Jl. Letjen. Sutoyo No. 124",
			Cabang:          true,
			CabangPencairan: true,
		},
		{
			BranchCode:      "37BJI",
			Name:            "KC Surabaya Darmo",
			CBCCode:         "00BJI",
			Address:         "Jl. Raya Darmo No. 105 - 107, Kota Surabaya",
			Cabang:          true,
			CabangPencairan: true,
		},
		{
			BranchCode:      "38BJI",
			Name:            "KC Bali - Denpasar",
			CBCCode:         "00BJI",
			Address:         "Komp. Pertokoan dan Perkantoran Teuku Umar Investama, Jl. Teuku Umar No. 121 Blok D1 dan D2",
			Cabang:          true,
			CabangPencairan: true,
		},
		{
			BranchCode:      "39BJI",
			Name:            "KC Balikpapan",
			CBCCode:         "00BJI",
			Address:         "Jl. Jenderal Sudirman No. 11",
			Cabang:          true,
			CabangPencairan: true,
		},
		{
			BranchCode:      "40BJI",
			Name:            "KC Samarinda",
			CBCCode:         "00BJI",
			Address:         "Jl. Jenderal Sudirman No. 4C Samarinda",
			Cabang:          true,
			CabangPencairan: true,
		},
		{
			BranchCode:      "41BJI",
			Name:            "KC Makassar Botolempangan",
			CBCCode:         "00BJI",
			Address:         "Jl. Botolempangan No. 18",
			Cabang:          true,
			CabangPencairan: true,
		},
		{
			BranchCode:      "42BJI",
			Name:            "KC Pontianak - Juanda",
			CBCCode:         "00BJI",
			Address:         "Jalan Ir. H. Juanda No. 55 -56 RT 003 RW 005 Kelurahan Darat Sekip Kecamatan Pontianak Kota",
			Cabang:          true,
			CabangPencairan: true,
		},
	}

	var program = []models.Program{
		{
			Code: "020",
			Name: "periodik Rating",
		},
		{
			Code: "021",
			Name: "Kredit Commercial Multifinance",
		},
		{
			Code: "022",
			Name: "Periodic Rating Multifinance Commercial",
		},
	}

	var segment = []models.Segment{
		{
			Code: "MB100",
			Name: "BJI Consumer Loan",
		},
		{
			Code: "SM100",
			Name: "BJI SME & Micro",
		},
	}

	connection.DB.Save(&cabang)
	connection.DB.Save(&program)
	connection.DB.Save(&segment)
}

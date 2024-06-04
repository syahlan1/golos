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

func CreditSeed() {
	var credit = []models.CreditType{
		{
			Code: "0213",
			Name: "MIKRO INVESTASI",
		},
		{
			Code: "0216",
			Name: "MIKRO MODAL KERJA",
		},
		{
			Code: "0521",
			Name: "PEMILIKAN KENDARAAN BERMOTOR",
		},
		{
			Code: "0527",
			Name: "MULTIGUNA UMUM",
		},
		{
			Code: "0529",
			Name: "KREDIT TANPA AGUNAN",
		},
	}

	var purpose = []models.Purpose{
		{
			Code: "06",
			Name: "Kredit Rekening Koran (KRK)",
		  },
		  {
			Code: "10",
			Name: "Kredit Atas Permintaan (KAP)",
		  },
		  {
			Code: "18",
			Name: "Kredit Angsuran Berjangka (KAB)",
		  },
		  {
			Code: "25",
			Name: "KMK Khusus Kepada Jasa Pengembang",
		  },
		  {
			Code: "26",
			Name: "Kredit Investasi (KIN)",
		  },
		  {
			Code: "31",
			Name: "Kredit Kepemilikan Rumah (KPR)",
		  },
		  {
			Code: "33",
			Name: "Kredit Kepemilikan Kendaraan Bermotor (KKB)",
		  },
		  {
			Code: "34",
			Name: "Kredit Serba Guna (KSG)",
		  },
		  {
			Code: "35",
			Name: "Kredit Tanpa Agunan (KTA)",
		  },
		  {
			Code: "39",
			Name: "Kredit Dengan Agunan Deposito",
		  },
		  {
			Code: "45",
			Name: "Kredit Excecuting Multifinance",
		  },
		  {
			Code: "46",
			Name: "Joint Financing",
		  },
		  {
			Code: "47",
			Name: "Channeling Loan",
		  },
		  {
			Code: "48",
			Name: "Sindikasi",
		  },
		  {
			Code: "49",
			Name: "Agriculture Financing",
		  },
		  {
			Code: "50",
			Name: "Bank Garansi",
		  },
		  {
			Code: "51",
			Name: "Standby Letter of Credit (L/C)",
		  },
		  {
			Code: "52",
			Name: "Pre Export Financing",
		  },
	}

	var currency = []models.Currency{
		{
			Code: "IDR",
			Name: "Rupiah",
		},
		{
			Code: "USD",
			Name: "Dollar",
		},
		{
			Code: "EUR",
			Name: "Euro",
		},
	}

	var assessment = []models.Assessment{
		{
			Name: "Menurut penilaian bank",
		},
		{
			Name: "Menurut penilaian appraisal",
		},
	}

	var collateral = []models.CollateralType{
		{

			Code:         "RE",
			Name:         "Tanah Dan Bangunan",
			LinkTable:    "COLLATERAL_RE",
			CodeIbs:      "001",
			ReqAppraisal: "1",
			RatingCode:   "A",
		},
		{

			Code:         "VEH",
			Name:         "Mesin-Mesin",
			LinkTable:    "COLLATERAL_VEH",
			CodeIbs:      "061",
			ReqAppraisal: "1",
			RatingCode:   "4",
		},
		{

			Code:         "DEP",
			Name:         "Deposito",
			LinkTable:    "COLLATERAL_DEP",
			CodeIbs:      "101",
			ReqAppraisal: "1",
			RatingCode:   "7",
		},
		{

			Code:         "BOND",
			Name:         "Garansi Bank",
			LinkTable:    "COLLATERAL_BOND",
			CodeIbs:      "202",
			ReqAppraisal: "1",
			RatingCode:   "1",
		},
		{

			Code:         "PG",
			Name:         "Personal Guarantee",
			LinkTable:    "COLLATERAL_PG",
			CodeIbs:      "281",
			ReqAppraisal: "1",
			RatingCode:   "1",
		},
		{

			Code:         "MISC",
			Name:         "Emas dan Perhiasan",
			LinkTable:    "COLLATERAL_MISC",
			CodeIbs:      "331",
			ReqAppraisal: "1",
			RatingCode:   "7",
		},
		{

			Code:         "STOCK",
			Name:         "Saham",
			LinkTable:    "COLLATERAL_STOCK",
			CodeIbs:      "351",
			ReqAppraisal: "1",
			RatingCode:   "7",
		},
		{

			Code:         "AR",
			Name:         "Piutang/Tagihan",
			LinkTable:    "COLLATERAL_AR",
			CodeIbs:      "381",
			ReqAppraisal: "1",
			RatingCode:   "2",
		},
		{

			Code:         "LC",
			Name:         "Standby LC",
			LinkTable:    "COLLATERAL_LC",
			CodeIbs:      "401",
			ReqAppraisal: "1",
			RatingCode:   "1",
		},
		{

			Code:         "PNCHQ",
			Name:         "Promisary Note/Cheques",
			LinkTable:    "COLLATERAL_PNCHQ",
			CodeIbs:      "431",
			ReqAppraisal: "1",
			RatingCode:   "1",
		},
		{

			Code:         "TRCON",
			Name:         "Transfer Contract",
			LinkTable:    "COLLATERAL_TRCON",
			CodeIbs:      "451",
			ReqAppraisal: "1",
			RatingCode:   "1",
		},
		{

			Code:         "SPK",
			Name:         "Contract dan SPK",
			LinkTable:    "COLLATERAL_SPK",
			CodeIbs:      "481",
			ReqAppraisal: "1",
			RatingCode:   "1",
		},
		{

			Code:         "LSAGR",
			Name:         "Lease Agreement",
			LinkTable:    "COLLATERAL_LSAGR",
			CodeIbs:      "501",
			ReqAppraisal: "1",
			RatingCode:   "1",
		},
		{

			Code:         "RE",
			Name:         "Tanah Darat",
			LinkTable:    "COLLATERAL_RE",
			CodeIbs:      "004",
			ReqAppraisal: "1",
			RatingCode:   "3",
		},
		{

			Code:         "RE",
			Name:         "Tanah dan Pabrik/Proyek",
			LinkTable:    "COLLATERAL_RE",
			CodeIbs:      "005",
			ReqAppraisal: "1",
			RatingCode:   "A",
		},
		{

			Code:         "RE",
			Name:         "Bangunan",
			LinkTable:    "COLLATERAL_RE",
			CodeIbs:      "002",
			ReqAppraisal: "1",
			RatingCode:   "A",
		},
		{

			Code:         "RE",
			Name:         "Pabrik/Proyek",
			LinkTable:    "COLLATERAL_RE",
			CodeIbs:      "003",
			ReqAppraisal: "1",
			RatingCode:   "A",
		},
		{

			Code:         "RE",
			Name:         "Tanah Kebun",
			LinkTable:    "COLLATERAL_RE",
			CodeIbs:      "006",
			ReqAppraisal: "1",
			RatingCode:   "3",
		},
		{

			Code:         "VEH",
			Name:         "Kapal dan Sejenisnya",
			LinkTable:    "COLLATERAL_VEH",
			CodeIbs:      "062",
			ReqAppraisal: "1",
			RatingCode:   "9",
		},
		{

			Code:         "VEH",
			Name:         "Pesawat terbang dan sejenisnya",
			LinkTable:    "COLLATERAL_VEH",
			CodeIbs:      "063",
			ReqAppraisal: "1",
			RatingCode:   "9",
		},
		{

			Code:         "VEH",
			Name:         "Kenderaan Bermotor",
			LinkTable:    "COLLATERAL_VEH",
			CodeIbs:      "064",
			ReqAppraisal: "1",
			RatingCode:   "5",
		},
		{

			Code:         "VEH",
			Name:         "Alat Berat",
			LinkTable:    "COLLATERAL_VEH",
			CodeIbs:      "065",
			ReqAppraisal: "1",
			RatingCode:   "4",
		},
		{

			Code:         "BOND",
			Name:         "Obligasi",
			LinkTable:    "COLLATERAL_BOND",
			CodeIbs:      "203",
			ReqAppraisal: "1",
			RatingCode:   "6",
		},
		{

			Code:         "BOND",
			Name:         "Retention Bond",
			LinkTable:    "COLLATERAL_BOND",
			CodeIbs:      "204",
			ReqAppraisal: "1",
			RatingCode:   "6",
		},
		{

			Code:         "BOND",
			Name:         "Performance Bond",
			LinkTable:    "COLLATERAL_BOND",
			CodeIbs:      "205",
			ReqAppraisal: "1",
			RatingCode:   "6",
		},
		{

			Code:         "BOND",
			Name:         "Advance Payment Bond",
			LinkTable:    "COLLATERAL_BOND",
			CodeIbs:      "206",
			ReqAppraisal: "1",
			RatingCode:   "6",
		},
		{

			Code:         "BOND",
			Name:         "Surety Bond",
			LinkTable:    "COLLATERAL_BOND",
			CodeIbs:      "207",
			ReqAppraisal: "1",
			RatingCode:   "6",
		},
		{

			Code:         "BOND",
			Name:         "Sertifikat Bank Indonesia",
			LinkTable:    "COLLATERAL_BOND",
			CodeIbs:      "208",
			ReqAppraisal: "1",
			RatingCode:   "6",
		},
		{

			Code:         "PG",
			Name:         "Pemerintahan",
			LinkTable:    "COLLATERAL_PG",
			CodeIbs:      "282",
			ReqAppraisal: "1",
			RatingCode:   "1",
		},
		{

			Code:         "PG",
			Name:         "Askrindo",
			LinkTable:    "COLLATERAL_PG",
			CodeIbs:      "283",
			ReqAppraisal: "1",
			RatingCode:   "1",
		},
		{

			Code:         "PG",
			Name:         "Perum PKK/PT SPU",
			LinkTable:    "COLLATERAL_PG",
			CodeIbs:      "284",
			ReqAppraisal: "1",
			RatingCode:   "1",
		},
		{

			Code:         "PG",
			Name:         "Perusahaan (Corp GTE)",
			LinkTable:    "COLLATERAL_PG",
			CodeIbs:      "285",
			ReqAppraisal: "1",
			RatingCode:   "1",
		},
		{

			Code:         "PG",
			Name:         "Avalist Guarantee",
			LinkTable:    "COLLATERAL_PG",
			CodeIbs:      "286",
			ReqAppraisal: "1",
			RatingCode:   "1",
		},
		{

			Code:         "PG",
			Name:         "PT. ASEI",
			LinkTable:    "COLLATERAL_PG",
			CodeIbs:      "287",
			ReqAppraisal: "1",
			RatingCode:   "1",
		},
		{

			Code:         "MISC",
			Name:         "Lainnya",
			LinkTable:    "COLLATERAL_MISC",
			CodeIbs:      "332",
			ReqAppraisal: "1",
			RatingCode:   "1",
		},
		{

			Code:         "LC",
			Name:         "SKBDN",
			LinkTable:    "COLLATERAL_LC",
			CodeIbs:      "402",
			ReqAppraisal: "1",
			RatingCode:   "1",
		},
		{

			Code:         "INV",
			Name:         "Barang Inventory",
			LinkTable:    "COLLATERAL_INV",
			CodeIbs:      "161",
			ReqAppraisal: "1",
			RatingCode:   "2",
		},
		{

			Code:         "INV",
			Name:         "Persediaan Barang / Stock",
			LinkTable:    "COLLATERAL_INV",
			CodeIbs:      "162",
			ReqAppraisal: "1",
			RatingCode:   "2",
		},
		{

			Code:         "VEH",
			Name:         "Mesin dan Tanah",
			LinkTable:    "COLLATERAL_VEH",
			CodeIbs:      "066",
			ReqAppraisal: "1",
			RatingCode:   "4",
		},
		{

			Code:         "VEH",
			Name:         "Mesin dan Alat Berat",
			LinkTable:    "COLLATERAL_VEH",
			CodeIbs:      "067",
			ReqAppraisal: "1",
			RatingCode:   "4",
		},
	}

	adrstr := func(s string) *string { return &s }
	adrflt := func(s float64) *float64 { return &s }
	var proofOfOwnership = []models.ProofOfOwnership{

		{
			Code:       "A-FIDUCIA",
			Name:       "FEO-AKTA FIDUCIA",
			Flag:       nil,
			RatingCode: adrstr("A"),
		},
		{
			Code:       "BOT",
			Name:       "BANGUN GUNA SERAH (BUILD OPERATE TRANSFER)",
			Flag:       nil,
			RatingCode: adrstr(""),
		},
		{
			Code:       "BPKB",
			Name:       "BUKTI PEMILIKAN KENDARAAN BERMOTOR",
			Flag:       nil,
			RatingCode: nil,
		},
		{
			Code:       "CLERANCE",
			Name:       "CLERANCE SERT./KETERANGAN DARI BPN",
			Flag:       adrstr("RE"),
			RatingCode: adrstr("10"),
		},
		{
			Code:       "COVERNOTE",
			Name:       "SURAT KETERANGAN NOTARIS",
			Flag:       adrstr("RE"),
			RatingCode: adrstr("10"),
		},
		{
			Code:       "GIRIK",
			Name:       "GIRIK/PETOK D/LETTER C",
			Flag:       adrstr("RE"),
			RatingCode: adrstr("10"),
		},
		{
			Code:       "KAD",
			Name:       "milik sendiri",
			Flag:       adrstr("DEP"),
			RatingCode: nil,
		},
		{
			Code:       "SHGB",
			Name:       "SERT. HAK GUNA BANGUNAN",
			Flag:       adrstr("RE"),
			RatingCode: adrstr("9"),
		},
		{
			Code:       "SHGBHM",
			Name:       "SERT. HAK GUNA BANGUNAN ATAS HAK MILIK",
			Flag:       adrstr("RE"),
			RatingCode: adrstr("9"),
		},
		{
			Code:       "SHGU",
			Name:       "SERT. HAK GUNA USAHA",
			Flag:       adrstr("RE"),
			RatingCode: adrstr("9"),
		},
		{
			Code:       "SHM",
			Name:       "SERT. HAK MILIK",
			Flag:       adrstr("RE"),
			RatingCode: adrstr("8"),
		},
		{
			Code:       "SHMSRS",
			Name:       "SERT. HAK MILIK ATAS SATUAN RUSUN/APT.",
			Flag:       adrstr("RE"),
			RatingCode: adrstr("8"),
		},
		{
			Code:       "SHP",
			Name:       "SERT. HAK PAKAI",
			Flag:       adrstr("RE"),
			RatingCode: adrstr("10"),
		},
		{
			Code:       "SHPM",
			Name:       "SERT. HAK PAKAI DI ATAS TANAH HAK MILIK",
			Flag:       adrstr("RE"),
			RatingCode: adrstr("10"),
		},
		{
			Code:       "SKPT",
			Name:       "SURAT KETERANGAN PENDAFTARAN TANAH",
			Flag:       adrstr("RE"),
			RatingCode: adrstr("10"),
		},
	}

	var formOfBindings = []models.FormOfBinding{
		{
			Name: "Hak Tanggungan",
		},
		{
			Name: "APHT",
		},
		{
			Name: "Fiducial Bawah Tangan",
		},
		{
			Name: "Akta Fiducia",
		},
		{
			Name: "Kuasa Jual Bawah Tangan",
		},
		{
			Name: "Kuasa Jual Notarial",
		},
		{
			Name: "Tanpa Pengikatan",
		},
		{
			Name: "Gadai",
		},
		{
			Name: "Fidusia",
		},
		{
			Name: "SKMHT",
		},
		{
			Name: "Cessie",
		},
		{
			Name: "Belum Diikat atau Tidak Diikat",
		},
		{
			Name: "Hipotik",
		},
		{
			Name: "Surat Kuasa Jual (SKJ)",
		},
		{
			Name: "SKHMT",
		},
		{
			Name: "Lainnya",
		},
	}

	var collateralClassification = []models.CollateralClassification{
		{
			Name: "Jaminan Utama",
		},
		{
			Name: "Jaminan Tambahan",
		},
	}

	var submissionType = []models.SubmissionType{
		{
			Code: "01",
			Name: "Periodik Rating",
			SibsCode: adrstr("1"),
			Scoring: adrstr("1"),
			StopTrack: nil,
			GoToTrack: nil,
			SibsLimit: adrflt(9000000000000.0),
			Channeling: adrstr("0"),
		  },
		  {
			Code: "02",
			Name: "Perubahan Jaminan",
			SibsCode: adrstr("3"),
			Scoring: adrstr("0"),
			StopTrack: nil,
			GoToTrack: nil,
			SibsLimit: nil,
			Channeling: adrstr("0"),
		  },
		  {
			Code: "03",
			Name: "Perubahan Limit",
			SibsCode: adrstr("2"),
			Scoring: adrstr("1"),
			StopTrack: nil,
			GoToTrack: nil,
			SibsLimit: nil,
			Channeling: adrstr("0"),
		  },
		  {
			Code: "04",
			Name: "Renewal",
			SibsCode: adrstr("4"),
			Scoring: adrstr("1"),
			StopTrack: nil,
			GoToTrack: nil,
			SibsLimit: nil,
			Channeling: adrstr("1"),
		  },
		  {
			Code: "05",
			Name: "Perubahan Syarat Kredit",
			SibsCode: adrstr("5"),
			Scoring: adrstr("0"),
			StopTrack: nil,
			GoToTrack: nil,
			SibsLimit: nil,
			Channeling: adrstr("1"),
		  },
		  {
			Code: "06",
			Name: "Withdrawal",
			SibsCode: adrstr("6"),
			Scoring: adrstr("1"),
			StopTrack: adrstr("3.6"),
			GoToTrack: adrstr("5.1"),
			SibsLimit: nil,
			Channeling: adrstr("1"),
		  },
		  {
			Code: "07",
			Name: "Analisa Restrukturisasi",
			SibsCode: adrstr("7"),
			Scoring: adrstr("0"),
			StopTrack: nil,
			GoToTrack: nil,
			SibsLimit: nil,
			Channeling: adrstr("0"),
		  },
		  {
			Code: "08",
			Name: "Penyelesaian Kredit",
			SibsCode: adrstr("8"),
			Scoring: adrstr("0"),
			StopTrack: nil,
			GoToTrack: nil,
			SibsLimit: nil,
			Channeling: adrstr("0"),
		  },
		  {
			Code: "09",
			Name: "Withdrawal Past Due NCL",
			SibsCode: adrstr("9"),
			Scoring: adrstr("0"),
			StopTrack: nil,
			GoToTrack: nil,
			SibsLimit: nil,
			Channeling: adrstr("0"),
		  },
		  {
			Code: "10",
			Name: "Periodik Rating Old",
			SibsCode: adrstr("10"),
			Scoring: adrstr("1"),
			StopTrack: nil,
			GoToTrack: nil,
			SibsLimit: nil,
			Channeling: adrstr("0"),
		  },
		  {
			Code: "11",
			Name: "Automatic Renewal",
			SibsCode: adrstr("4"),
			Scoring: adrstr("1"),
			StopTrack: nil,
			GoToTrack: nil,
			SibsLimit: nil,
			Channeling: adrstr("1"),
		  },
		  {
			Code: "T01",
			Name: "ITTP - Rating Customer",
			SibsCode: nil,
			Scoring: nil,
			StopTrack: nil,
			GoToTrack: nil,
			SibsLimit: nil,
			Channeling: nil,
		  },
		  {
			Code: "T02",
			Name: "ITTP - Penerbitan",
			SibsCode: nil,
			Scoring: nil,
			StopTrack: nil,
			GoToTrack: nil,
			SibsLimit: nil,
			Channeling: nil,
		  },
		  {
			Code: "T03",
			Name: "ITTP - Perpanjangan",
			SibsCode: nil,
			Scoring: nil,
			StopTrack: nil,
			GoToTrack: nil,
			SibsLimit: nil,
			Channeling: nil,
		  },
		  {
			Code: "T04",
			Name: "ITTP - Transaksi",
			SibsCode: nil,
			Scoring: nil,
			StopTrack: nil,
			GoToTrack: nil,
			SibsLimit: nil,
			Channeling: nil,
		  },
	}

	connection.DB.Save(&credit)
	connection.DB.Save(&purpose)
	connection.DB.Save(&currency)
	connection.DB.Save(&assessment)
	connection.DB.Save(&collateral)
	connection.DB.Save(&proofOfOwnership)
	connection.DB.Save(&formOfBindings)
	connection.DB.Save(&collateralClassification)
	connection.DB.Save(&submissionType)
}

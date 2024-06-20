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
		panic(err,)
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

func ApplicantSeed() {
	var homeStatus = []models.HomeStatus{
		{
			Code: "1",
			Name : "Credit.",
			Sibs : "KANTR",
		},
		{
			Code: "2",
			Name : "Institution",
			Sibs : "DINAS",
		},
		{
			Code: "3",
			Name : "Others",
			Sibs : "LAIN",
		},
		{
			Code: "4",
			Name : "Own",
			Sibs : "MILIK",
		},
		{
			Code: "5",
			Name : "Parents",
			Sibs : "KLRGA",
		},
		{
			Code: "6",
			Name : "Rented",
			Sibs : "SEWA",
		},
	}

	var maritalStatus = []models.MaritalStatus{
		{
			Code : "A",
			Name : "Kawin",
		},
		{
			Code : "B",
			Name : "Belum Kawin",
		},
		{
			Code : "C",
			Name : "Janda",
		},
		{
			Code : "D",
			Name : "Duda",
		},
		{
			Code : "E",
			Name : "Tidak Diketahui",
		},
	}


	var nationality = []models.Nationality{
		{
			Code: "WNI",
			Name: "Indonesia",
		},
		{
			Code: "WNA",
			Name: "Asing",
		},
	}

	var education = []models.Education{
		{
			Code : "A",
			Name : "Tanpa Gelar",
		},
		{
			Code : "B",
			Name : "Diploma 1",
		},
		{
			Code : "C",
			Name : "Diploma 2",
		},
		{
			Code : "D",
			Name : "Diploma 3",
		},
		{
			Code : "E",
			Name : "S - 1",
		},
		{
			Code : "F",
			Name : "S - 2",
		},
		{
			Code : "G",
			Name : "S - 3",
		},
		{
			Code : "Z",
			Name : "Lainnya",
		},
	}

	var jobPosition = []models.JobPosition{
		{
			Code : "001",
			Name : "Accounting/finance officer",
		},
		{
			Code : "002",
			Name : "Customer service",
		},
		{
			Code : "003",
			Name : "Engineering",
		},
		{
			Code : "004",
			Name : "Eksekutif",
		},
		{
			Code : "005",
			Name : "Administrasi umum",
		},
		{
			Code : "006",
			Name : "Teknologi informasi",
		},
		{
			Code : "007",
			Name : "Konsultan/Analis",
		},
		{
			Code : "008",
			Name : "Marketing",
		},
		{
			Code : "009",
			Name : "Pengajar (Guru, Dosen)",
		},
		{
			Code : "01",
			Name : "PEMILIK-Direktur Utama",
		},
		{
			Code : "010",
			Name : "Militer",
		},
		{
			Code : "011",
			Name : "Pensiunan",
		},
		{
			Code : "012",
			Name : "Pelajar/Mahasiswa",
		},
		{
			Code : "013",
			Name : "Wiraswasta",
		},
		{
			Code : "014",
			Name : "Polisi",
		},
		{
			Code : "015",
			Name : "Petani",
		},
		{
			Code : "016",
			Name : "Nelayan",
		},
		{
			Code : "017",
			Name : "Peternak",
		},
		{
			Code : "018",
			Name : "Dokter",
		},
		{
			Code : "019",
			Name : "Tenaga Medis",
		},
		{
			Code : "02",
			Name : "PEMILIK-Direktur",
		},
		{
			Code : "020",
			Name : "Pengacara/Notaris",
		},
		{
			Code : "021",
			Name : "Pekerja Hotel/Restoran",
		},
		{
			Code : "022",
			Name : "Peneliti",
		},
		{
			Code : "023",
			Name : "Desainer",
		},
		{
			Code : "024",
			Name : "Arsitek",
		},
		{
			Code : "025",
			Name : "Pekerja Seni",
		},
		{
			Code : "026",
			Name : "Pengamanan",
		},
		{
			Code : "027",
			Name : "Pialang/Broker",
		},
		{
			Code : "028",
			Name : "Distributor",
		},
		{
			Code : "029",
			Name : "Pilot/Awak Pesawat",
		},
		{
			Code : "03",
			Name : "PEMILIK-Komusaris Utama",
		},
		{
			Code : "030",
			Name : "Nahkoda/Awak Kapal",
		},
		{
			Code : "031",
			Name : "Masinis/Sopir/Kondektur",
		},
		{
			Code : "032",
			Name : "Buruh Pabrik/Bangunan/Tani",
		},
		{
			Code : "033",
			Name : "Tukang / Pengrajin",
		},
		{
			Code : "034",
			Name : "Ibu Rumah Tangga",
		},
		{
			Code : "035",
			Name : "Pekerja Informal",
		},
		{
			Code : "036",
			Name : "Pejabat Negara / Pemerintahan",
		},
		{
			Code : "037",
			Name : "Pegawai Negara/Pemerintahan",
		},
		{
			Code : "04",
			Name : "PEMILIK-Komusaris",
		},
		{
			Code : "05",
			Name : "PEMILIK-Kuasa Direksi",
		},
		{
			Code : "06",
			Name : "PEMILIK-Bukan Pengurus",
		},
		{
			Code : "07",
			Name : "PEMILIK-Masyarakat",
		},
		{
			Code : "08",
			Name : "PEMILIK-Ketua Umum",
		},
		{
			Code : "09",
			Name : "PEMILIK-Ketua",
		},
		{
			Code : "099",
			Name : "Lain-lain",
		},
		{
			Code : "10",
			Name : "PEMILIK-Sekretaris",
		},
		{
			Code : "11",
			Name : "PEMILIK-Bendahara",
		},
		{
			Code : "12",
			Name : "PEMILIK-Lainnya",
		},
		{
			Code : "13",
			Name : "BUKAN PEMILIK-Direktur Utama",
		},
		{
			Code : "14",
			Name : "BUKAN PEMILIK-Direktur",
		},
		{
			Code : "15",
			Name : "BUKAN PEMILIK-Komisaris Utama",
		},
		{
			Code : "16",
			Name : "BUKAN PEMILIK-Komisaris",
		},
		{
			Code : "17",
			Name : "BUKAN PEMILIK-Kuasa Direksi",
		},
		{
			Code : "18",
			Name : "BUKAN PEMILIK-Ketua Umum",
		},
		{
			Code : "19",
			Name : "BUKAN PEMILIK-Ketua",
		},
		{
			Code : "20",
			Name : "BUKAN PEMILIK-Sekretaris",
		},
		{
			Code : "21",
			Name : "BUKAN PEMILIK-Bendahara",
		},
		{
			Code : "22",
			Name : "BUKAN PEMILIK-Lainnya",
		},
		{
			Code : "23",
			Name : "Notaris/Pengacara",
		},
		{
			Code : "24",
			Name : "Guru/Dosen",
		},
		{
			Code : "25",
			Name : "Mahasiswa/Pelajar",
		},
		{
			Code : "26",
			Name : "Akuntan",
		},
		{
			Code : "27",
			Name : "GM atau Kepala Biro",
		},
		{
			Code : "28",
			Name : "Staff",
		},
		{
			Code : "29",
			Name : "Bintara",
		},
		{
			Code : "30",
			Name : "Tamtama",
		},
		{
			Code : "31",
			Name : "Pengurus Koperasi/Yayasan/LSM",
		},
		{
			Code : "32",
			Name : "Anggota Koperasi/Yayasan/LSM",
		},
		{
			Code : "33",
			Name : "Kepala Pemerintahan",
		},
		{
			Code : "34",
			Name : "Anggota Lembaga Tinggi Negara",
		},
		{
			Code : "35",
			Name : "Menteri",
		},
		{
			Code : "36",
			Name : "Sekretaris",
		},
		{
			Code : "37",
			Name : "Risk Manager",
		},
		{
			Code : "38",
			Name : "Operational Manager",
		},
	}

	var businessSector = []models.BusinessSector{
		{
			Code : "01",
			Name : "Agriculture / Plantation",
		},
		{
			Code : "02",
			Name : "Fishery / Farm",
		},
		{
			Code : "03",
			Name : "Mining",
		},
		{
			Code : "04",
			Name : "Processing industry",
		},
		{
			Code : "05",
			Name : "Manufacture",
		},
		{
			Code : "06",
			Name : "Energy & Water",
		},
		{
			Code : "06001",
			Name : "Construction",
		},
		{
			Code : "06002",
			Name : "Automotive",
		},
		{
			Code : "06003",
			Name : "Retail/Wholesale trading",
		},
		{
			Code : "06004",
			Name : "Import/export/domestic trading",
		},
		{
			Code : "06005",
			Name : "Transportation",
		},
		{
			Code : "06006",
			Name : "Telecommunication",
		},
		{
			Code : "06007",
			Name : "Financial services",
		},
		{
			Code : "06008",
			Name : "Education Services",
		},
		{
			Code : "06009",
			Name : "Health Services",
		},
		{
			Code : "07",
			Name : "Non-financial services",
		},
		{
			Code : "07001",
			Name : "Hotels & Accommodation",
		},
		{
			Code : "08",
			Name : "Restaurant",
		},
		{
			Code : "08001",
			Name : "Real Estate",
		},
		{
			Code : "08002",
			Name : "Rental",
		},
		{
			Code : "09",
			Name : "Research and development",
		},
		{
			Code : "09001",
			Name : "Publish,Print&Advertising",
		},
		{
			Code : "10",
			Name : "Entertainment Activities",
		},
		{
			Code : "6001",
			Name : "Household",
		},
	}


	var NegaraDomisili = []models.Negara{
		{
			Code : "AE",
			Name : "UNITED ARAB EMIRATES",

		},
		{
			Code : "AF",
			Name : "AFGHANISTAN",

		},
		{
			Code : "AG",
			Name : "ANTIGUA AND BARBUDA",

		},
		{
			Code : "AI",
			Name : "ANGUILLA",

		},
		{
			Code : "AL",
			Name : "ALBANIA",

		},
		{
			Code : "AM",
			Name : "ARMENIA",

		},
		{
			Code : "AN",
			Name : "NETH. ANTILLES",

		},
		{
			Code : "AO",
			Name : "ANGOLA",

		},
		{
			Code : "AQ",
			Name : "ANTARCTICA",

		},
		{
			Code : "AR",
			Name : "ARGENTINA",

		},
		{
			Code : "AS",
			Name : "AMERICAN SAMOA",

		},
		{
			Code : "AT",
			Name : "AUSTRIA",

		},
		{
			Code : "AU",
			Name : "AUSTRALIA",

		},
		{
			Code : "AW",
			Name : "ARUBA",

		},
		{
			Code : "AZ",
			Name : "AZERBAIJAN",

		},
		{
			Code : "BA",
			Name : "BOSNIA AND HERZEGOVINA",

		},
		{
			Code : "BB",
			Name : "BARBADOS",

		},
		{
			Code : "BD",
			Name : "BANGLADESH",

		},
		{
			Code : "BE",
			Name : "BELGIUM",

		},
		{
			Code : "BF",
			Name : "BURKINA FASO",

		},
		{
			Code : "BG",
			Name : "BULGARIA",

		},
		{
			Code : "BH",
			Name : "BAHRAIN",

		},
		{
			Code : "BI",
			Name : "BURUNDI",

		},
		{
			Code : "BJ",
			Name : "BENIN",

		},
		{
			Code : "BM",
			Name : "BERMUDA",

		},
		{
			Code : "BN",
			Name : "BRUNEI DARUSSALAM",

		},
		{
			Code : "BO",
			Name : "BOLIVIA",

		},
		{
			Code : "BR",
			Name : "BRAZIL",

		},
		{
			Code : "BS",
			Name : "BAHAMAS",

		},
		{
			Code : "BT",
			Name : "BHUTAN",

		},
		{
			Code : "BUD",
			Name : "Budha",

		},
		{
			Code : "BV",
			Name : "BOUVET ISLAND",

		},
		{
			Code : "BW",
			Name : "BOTSWANA",

		},
		{
			Code : "BY",
			Name : "BELARUS",

		},
		{
			Code : "BZ",
			Name : "BELIZE",

		},
		{
			Code : "CA",
			Name : "CANADA",

		},
		{
			Code : "CC",
			Name : "COCOS (KEELING) ISLANDS",

		},
		{
			Code : "CD",
			Name : "CONGO, THE DEMOCRATIC REPUBLIC",

		},
		{
			Code : "CF",
			Name : "CENTRAL AFRICAN REPUBLIC",

		},
		{
			Code : "CG",
			Name : "CONGO",

		},
		{
			Code : "CH",
			Name : "SWITZERLAND",

		},
		{
			Code : "CI",
			Name : "COTE DIVOIRE",

		},
		{
			Code : "CK",
			Name : "COOK ISLANDS",

		},
		{
			Code : "CL",
			Name : "CHILE",

		},
		{
			Code : "CM",
			Name : "CAMEROON",

		},
		{
			Code : "CN",
			Name : "CHINA",

		},
		{
			Code : "CO",
			Name : "COLOMBIA",

		},
		{
			Code : "CR",
			Name : "COSTA RICA",

		},
		{
			Code : "CU",
			Name : "CUBA",

		},
		{
			Code : "CV",
			Name : "CAPE VERDE",

		},
		{
			Code : "CX",
			Name : "CHRISTMAS ISLAND",

		},
		{
			Code : "CY",
			Name : "CYPRUS",

		},
		{
			Code : "CZ",
			Name : "CZECH REPUBLIC",

		},
		{
			Code : "DE",
			Name : "GERMANY",

		},
		{
			Code : "DJ",
			Name : "DJIBOUTI",

		},
		{
			Code : "DK",
			Name : "DENMARK",

		},
		{
			Code : "DM",
			Name : "DOMINICA",

		},
		{
			Code : "DO",
			Name : "DOMINICAN REPUBLIC",

		},
		{
			Code : "DZ",
			Name : "ALGERIA",

		},
		{
			Code : "EC",
			Name : "ECUADOR",

		},
		{
			Code : "EE",
			Name : "ESTONIA",

		},
		{
			Code : "EG",
			Name : "EGYPT",

		},
		{
			Code : "EH",
			Name : "WESTERN SAHARA",

		},
		{
			Code : "ER",
			Name : "ERITREA",

		},
		{
			Code : "ES",
			Name : "SPAIN",

		},
		{
			Code : "ET",
			Name : "ETHIOPIA",

		},
		{
			Code : "FI",
			Name : "FINLAND",

		},
		{
			Code : "FJ",
			Name : "FIJI",

		},
		{
			Code : "FK",
			Name : "FALKLAND  ISLANDS (MALVINAS)",

		},
		{
			Code : "FM",
			Name : "MICRONESIA (FEDERATED STATES O",

		},
		{
			Code : "FO",
			Name : "FAEROE ISLANDS",

		},
		{
			Code : "FR",
			Name : "FRANCE",

		},
		{
			Code : "GA",
			Name : "GABON",

		},
		{
			Code : "GB",
			Name : "UNITED KINGDOM",

		},
		{
			Code : "GD",
			Name : "GRENADA",

		},
		{
			Code : "GE",
			Name : "GEORGIA",

		},
		{
			Code : "GF",
			Name : "FRENCH GUIANA",

		},
		{
			Code : "GG",
			Name : "GUERNSEY, C.I.",

		},
		{
			Code : "GH",
			Name : "GHANA",

		},
		{
			Code : "GI",
			Name : "GIBRALTAR",

		},
		{
			Code : "GL",
			Name : "GREENLAND",

		},
		{
			Code : "GM",
			Name : "GAMBIA",

		},
		{
			Code : "GN",
			Name : "GUINEA",

		},
		{
			Code : "GP",
			Name : "GUADELOUPE",

		},
		{
			Code : "GQ",
			Name : "EQUATORIAL GUINEA",

		},
		{
			Code : "GR",
			Name : "GREECE",

		},
		{
			Code : "GT",
			Name : "GUATEMALA",

		},
		{
			Code : "GU",
			Name : "GUAM",

		},
		{
			Code : "GW",
			Name : "GUINEA-BISSAU",

		},
		{
			Code : "GY",
			Name : "GUYANA",

		},
		{
			Code : "HIN",
			Name : "Hindu",

		},
		{
			Code : "HK",
			Name : "HONG KONG",

		},
		{
			Code : "HM",
			Name : "HEARD AND MCDONALD ISLANDS",

		},
		{
			Code : "HN",
			Name : "HONDURAS",

		},
		{
			Code : "HR",
			Name : "CROATIA",

		},
		{
			Code : "HT",
			Name : "HAITI",

		},
		{
			Code : "HU",
			Name : "HUNGARY",

		},
		{
			Code : "ID",
			Name : "INDONESIA",

		},
		{
			Code : "IE",
			Name : "IRELAND",

		},
		{
			Code : "IL",
			Name : "ISRAEL",

		},
		{
			Code : "IM",
			Name : "ISLE OF MAN",

		},
		{
			Code : "IN",
			Name : "INDIA",

		},
		{
			Code : "IO",
			Name : "BRITISH INDIAN OCEAN TERRITORY",

		},
		{
			Code : "IQ",
			Name : "IRAQ",

		},
		{
			Code : "IR",
			Name : "IRAN (ISLAMIC REPUBLIC OF)",

		},
		{
			Code : "IS",
			Name : "ICELAND",

		},
		{
			Code : "ISL",
			Name : "Islam",

		},
		{
			Code : "IT",
			Name : "ITALY",

		},
		{
			Code : "JE",
			Name : "JERSEY, C.I.",

		},
		{
			Code : "JM",
			Name : "JAMAICA",

		},
		{
			Code : "JO",
			Name : "JORDAN",

		},
		{
			Code : "JP",
			Name : "JAPAN",

		},
		{
			Code : "KAT",
			Name : "Kristen Katolik",

		},
		{
			Code : "KE",
			Name : "KENYA",

		},
		{
			Code : "KG",
			Name : "KYRGYZSTAN",

		},
		{
			Code : "KH",
			Name : "CAMBODIA",

		},
		{
			Code : "KHC",
			Name : "Kong Hu Chu",

		},
		{
			Code : "KI",
			Name : "KIRIBATI",

		},
		{
			Code : "KM",
			Name : "COMOROS",

		},
		{
			Code : "KN",
			Name : "SAINT KITTS AND NEVIS",

		},
		{
			Code : "KP",
			Name : "KOREA, DEMOCRATIC PEOPLES REP.",

		},
		{
			Code : "KR",
			Name : "KOREA, REPUBLIC OF",

		},
		{
			Code : "KW",
			Name : "KUWAIT",

		},
		{
			Code : "KY",
			Name : "CAYMAN ISLANDS",

		},
		{
			Code : "KZ",
			Name : "KAZAKHSTAN",

		},
		{
			Code : "LA",
			Name : "LAO PEOPLES DEMOCRATIC REPUBLI",

		},
		{
			Code : "LB",
			Name : "LEBANON",

		},
		{
			Code : "LC",
			Name : "SAINT LUCIA",

		},
		{
			Code : "LI",
			Name : "LIECHTENSTEIN",

		},
		{
			Code : "LK",
			Name : "SRI LANKA",

		},
		{
			Code : "LLN",
			Name : "Lainnya",

		},
		{
			Code : "LR",
			Name : "LIBERIA",

		},
		{
			Code : "LS",
			Name : "LESOTHO",

		},
		{
			Code : "LT",
			Name : "LITHUANIA",

		},
		{
			Code : "LU",
			Name : "LUXEMBOURG",

		},
		{
			Code : "LV",
			Name : "LATVIA",

		},
		{
			Code : "LY",
			Name : "LIBYAN ARAB JAMAHIRIYA",

		},
		{
			Code : "MA",
			Name : "MOROCCO",

		},
		{
			Code : "MC",
			Name : "MONACO",

		},
		{
			Code : "MD",
			Name : "MOLDOVA, REPUBLIC OF",

		},
		{
			Code : "MG",
			Name : "MADAGASCAR",

		},
		{
			Code : "MH",
			Name : "MARSHALL ISLANDS",

		},
		{
			Code : "MK",
			Name : "MACEDONIA,THE FORMER YUGOSLAV ",

		},
		{
			Code : "ML",
			Name : "MALI",

		},
		{
			Code : "MM",
			Name : "MYANMAR",

		},
		{
			Code : "MN",
			Name : "MONGOLIA",

		},
		{
			Code : "MO",
			Name : "MACAU",

		},
		{
			Code : "MP",
			Name : "NORTHERN MARIANA ISLANDS",

		},
		{
			Code : "MQ",
			Name : "MARTINIQUE",

		},
		{
			Code : "MR",
			Name : "MAURITANIA",

		},
		{
			Code : "MS",
			Name : "MONTSERRAT",

		},
		{
			Code : "MT",
			Name : "MALTA",

		},
		{
			Code : "MU",
			Name : "MAURITIUS",

		},
		{
			Code : "MV",
			Name : "MALDIVES",

		},
		{
			Code : "MW",
			Name : "MALAWI",

		},
		{
			Code : "MX",
			Name : "MEXICO",

		},
		{
			Code : "MY",
			Name : "MALAYSIA",

		},
		{
			Code : "MZ",
			Name : "MOZAMBIQUE",

		},
		{
			Code : "NA",
			Name : "NAMIBIA",

		},
		{
			Code : "NC",
			Name : "NEW CALEDONIA",

		},
		{
			Code : "NE",
			Name : "NIGER",

		},
		{
			Code : "NF",
			Name : "NORFOLK ISLAND",

		},
		{
			Code : "NG",
			Name : "NIGERIA",

		},
		{
			Code : "NI",
			Name : "NICARAGUA",

		},
		{
			Code : "NL",
			Name : "NETHERLANDS",

		},
		{
			Code : "NO",
			Name : "NORWAY",

		},
		{
			Code : "NP",
			Name : "NEPAL",

		},
		{
			Code : "NR",
			Name : "NAURU",

		},
		{
			Code : "NU",
			Name : "NIUE",

		},
		{
			Code : "NZ",
			Name : "NEW ZEALAND",

		},
		{
			Code : "OM",
			Name : "OMAN",

		},
		{
			Code : "PA",
			Name : "PANAMA",

		},
		{
			Code : "PE",
			Name : "PERU",

		},
		{
			Code : "PF",
			Name : "FRENCH POLYNESIA",

		},
		{
			Code : "PG",
			Name : "PAPUA NEW GUINEA",

		},
		{
			Code : "PH",
			Name : "PHILIPPINES",

		},
		{
			Code : "PK",
			Name : "PAKISTAN",

		},
		{
			Code : "PL",
			Name : "POLAND",

		},
		{
			Code : "PM",
			Name : "SAINT PIERRE AND MIQUELON",

		},
		{
			Code : "PN",
			Name : "PITCAIRN",

		},
		{
			Code : "PR",
			Name : "PUERTO RICO",

		},
		{
			Code : "PRO",
			Name : "Kristen Protestan",

		},
		{
			Code : "PS",
			Name : "PALESTINIAN TERRITORY, OCCUPIE",

		},
		{
			Code : "PT",
			Name : "PORTUGAL",

		},
		{
			Code : "PW",
			Name : "PALAU",

		},
		{
			Code : "PY",
			Name : "PARAGUAY",

		},
		{
			Code : "PZ",
			Name : "PANAMA CANAL ZONE",

		},
		{
			Code : "QA",
			Name : "QATAR",

		},
		{
			Code : "RE",
			Name : "REUNION",

		},
		{
			Code : "RO",
			Name : "ROMANIA",

		},
		{
			Code : "RU",
			Name : "RUSSIAN FEDERATION",

		},
		{
			Code : "RW",
			Name : "RWANDA",

		},
		{
			Code : "SA",
			Name : "SAUDI ARABIA",

		},
		{
			Code : "SB",
			Name : "SOLOMON ISLANDS",

		},
		{
			Code : "SC",
			Name : "SEYCHELLES",

		},
		{
			Code : "SD",
			Name : "SUDAN",

		},
		{
			Code : "SE",
			Name : "SWEDEN",

		},
		{
			Code : "SG",
			Name : "SINGAPORE",

		},
		{
			Code : "SH",
			Name : "SAINT HELENA",

		},
		{
			Code : "SI",
			Name : "SLOVENIA",

		},
		{
			Code : "SJ",
			Name : "SVALBARD AND JAN MAYEN ISLANDS",

		},
		{
			Code : "SK",
			Name : "SLOVAKIA",

		},
		{
			Code : "SL",
			Name : "SIERRA LEONE",

		},
		{
			Code : "SM",
			Name : "SAN MARINO",

		},
		{
			Code : "SN",
			Name : "SENEGAL",

		},
		{
			Code : "SO",
			Name : "SOMALIA",

		},
		{
			Code : "SR",
			Name : "SURINAME",

		},
		{
			Code : "ST",
			Name : "SAO TOME AND PRINCIPE",

		},
		{
			Code : "SV",
			Name : "EL SALVADOR",

		},
		{
			Code : "SY",
			Name : "SYRIAN ARAB REPUBLIC",

		},
		{
			Code : "SZ",
			Name : "SWAZILAND",

		},
		{
			Code : "TC",
			Name : "TURKS AND CAICOS ISLANDS",

		},
		{
			Code : "TD",
			Name : "CHAD",

		},
		{
			Code : "TF",
			Name : "FRENCH SOUTHERN TERRITORIES",

		},
		{
			Code : "TG",
			Name : "TOGO",

		},
		{
			Code : "TH",
			Name : "THAILAND",

		},
		{
			Code : "TJ",
			Name : "TAJIKISTAN",

		},
		{
			Code : "TK",
			Name : "TOKELAU",

		},
		{
			Code : "TM",
			Name : "TURKMENISTAN",

		},
		{
			Code : "TN",
			Name : "TUNISIA",

		},
		{
			Code : "TO",
			Name : "TONGA",

		},
		{
			Code : "TP",
			Name : "EAST TIMOR",

		},
		{
			Code : "TR",
			Name : "TURKEY",

		},
		{
			Code : "TT",
			Name : "TRINIDAD AND TOBAGO",

		},
		{
			Code : "TV",
			Name : "TUVALU",

		},
		{
			Code : "TW",
			Name : "TAIWAN",

		},
		{
			Code : "TZ",
			Name : "TANZANIA, UNITED REPUBLIC OF",

		},
		{
			Code : "UA",
			Name : "UKRAINE",

		},
		{
			Code : "UG",
			Name : "UGANDA",

		},
		{
			Code : "UM",
			Name : "UNITED STATES MINOR OUTLAYING ",

		},
		{
			Code : "US",
			Name : "UNITED STATES",

		},
		{
			Code : "UY",
			Name : "URUGUAY",

		},
		{
			Code : "UZ",
			Name : "UZBEKISTAN",

		},
		{
			Code : "VA",
			Name : "HOLY SEE (VATICAN CITY STATE)",

		},
		{
			Code : "VC",
			Name : "SAINT VINCENT AND THE GRENADIN",

		},
		{
			Code : "VE",
			Name : "VENEZUELA",

		},
		{
			Code : "VG",
			Name : "VIRGIN ISLANDS, BRITISH",

		},
		{
			Code : "VI",
			Name : "VIRGIN ISLANDS, U.S.",

		},
		{
			Code : "VN",
			Name : "VIET NAM",

		},
		{
			Code : "VU",
			Name : "VANUATU",

		},
		{
			Code : "WF",
			Name : "WALLIS AND FUTUNA ISLANDS",

		},
		{
			Code : "WS",
			Name : "SAMOA",

		},
		{
			Code : "YE",
			Name : "YEMEN",

		},
		{
			Code : "YT",
			Name : "MAYOTTE",

		},
		{
			Code : "YU",
			Name : "YUGOSLAVIA",

		},
		{
			Code : "ZA",
			Name : "SOUTH AFRICA",

		},
		{
			Code : "ZM",
			Name : "ZAMBIA",

		},
		{
			Code : "ZW",
			Name : "ZIMBABWE",

		},
	}

	var gender = []models.Gender{
		{
			Name: "Laki-laki",
		},
		{
			Name: "Perempuan",
		},
	}

	var addressType = []models.AddressType{
		{
			Code : "DINAS",
			Name : "Rumah Dinas/Instansi",
		},
		{
			Code : "KANTR",
			Name : "Kantor",
		},
		{
			Code : "KLRGA",
			Name : "Rumah Milik Keluarga",
		},
		{
			Code : "KOST",
			Name : "Rumah Kost",
		},
		{
			Code : "LAIN",
			Name : "Lainnya",
		},
		{
			Code : "MILIK",
			Name : "Rumah Milik Sendiri",
		},
		{
			Code : "PROJ",
			Name : "Proyek",
		},
		{
			Code : "SEWA",
			Name : "Rumah Sewa/Kontrakan",
		},
		{
			Code : "PBRK",
			Name : "Pabrik",
		},
		{
			Code : "INDS",
			Name : "Lokasi di Indonesia",
		},
		{
			Code : "KTRP",
			Name : "Kantor Pusat",
		},
	}

	connection.DB.Save(&homeStatus)
	connection.DB.Save(&maritalStatus)
	connection.DB.Save(&nationality)
	connection.DB.Save(&education)
	connection.DB.Save(&jobPosition)
	connection.DB.Save(&businessSector)
	connection.DB.Save(&NegaraDomisili)
	connection.DB.Save(&gender)
	connection.DB.Save(&addressType)
}

func BusinessSeed() {

	var companyFirstName = []models.CompanyFirstName{
		{
			Code : "01",
			Name : "BUMDES",
		},
		{
			Code : "02",
			Name : "CV",
		},
		{
			Code : "03",
			Name : "Debitur Kelompok",
		},
		{
			Code : "04",
			Name : "EMKL",
		},
		{
			Code : "05",
			Name : "Firma",
		},
		{
			Code : "06",
			Name : "Gabungan Koperasi",
		},
		{
			Code : "07",
			Name : "Induk Koperasi",
		},
		{
			Code : "08",
			Name : "Koperasi",
		},
		{
			Code : "09",
			Name : "KUD",
		},
		{
			Code : "10",
			Name : "Limited",
		},
		{
			Code : "11",
			Name : "MAI",
		},
		{
			Code : "12",
			Name : "NV",
		},
		{
			Code : "13",
			Name : "PD",
		},
		{
			Code : "14",
			Name : "Persero",
		},
		{
			Code : "15",
			Name : "Persekutuan Perdata",
		},
		{
			Code : "16",
			Name : "PU",
		},
		{
			Code : "17",
			Name : "Primer Koperasi",
		},
		{
			Code : "18",
			Name : "PT",
		},
		{
			Code : "19",
			Name : "Pusat Koperasi",
		},
		{
			Code : "20",
			Name : "Pusat KUD",
		},
		{
			Code : "21",
			Name : "UD",
		},
		{
			Code : "22",
			Name : "UDKP",
		},
		{
			Code : "23",
			Name : "Yayasan",
		},
		{
			Code : "24",
			Name : "PP Daerah",
		},
		{
			Code : "25",
			Name : "PU Daerah",
		},
		{
			Code : "99",
			Name : "Lainnya",
		},
	}

	var companyType = []models.CompanyType{
		{
			Code : "A",
			Name : "Perorangan",
		},
		{
			Code : "B",
			Name : "Badan Usaha Sws",
		},
		{
			Code : "C",
			Name : "Bank",
		},
		{
			Code : "D",
			Name : "Lbg.Kg.Non Bank",
		},
		{
			Code : "E",
			Name : "Badan Sosial",
		},
		{
			Code : "F",
			Name : "Lbg.Pemrth.",
		},
		{
			Code : "G",
			Name : "Lbg.Intrnsl.",
		},
		{
			Code : "H",
			Name : "Pwkl.Neg.Asing",
		},
		{
			Code : "I",
			Name : "Koperasi",
		},
		{
			Code : "J",
			Name : "BUMN",
		},
		{
			Code : "K",
			Name : "BUMD",
		},
		{
			Code : "L",
			Name : "Pen.Modal Asing",
		},
		{
			Code : "M",
			Name : "BPR",
		},
		{
			Code : "N",
			Name : "Lembaga PEMDA",
		},
		{
			Code : "O",
			Name : "Yayasan",
		},
		{
			Code : "P",
			Name : "Prsh Securitas",
		},
		{
			Code : "Q",
			Name : "PMDN",
		},
		{
			Code : "R",
			Name : "Lembaga Nirlaba",
		},
	}

	var externalRatingCompany = []models.ExternalRatingCompany{
		{
			Code : "CARD",
			Name : "Credit Card Center",
		},
		{
			Code : "FITCH",
			Name : "Fitch Rating",
		},
		{
			Code : "FITID",
			Name : "Fitch Indonesia",
		},
		{
			Code : "LAIN",
			Name : "Lainnya",
		},
		{
			Code : "LPINT",
			Name : "Lembaga Pemeringkatan Internasional",
		},
		{
			Code : "LPNAS",
			Name : "Lembaga Pemeringkatan Nasional",
		},
		{
			Code : "MARKT",
			Name : "Marketing / Business Unit",
		},
		{
			Code : "MDS",
			Name : "Moody's",
		},
		{
			Code : "MDSID",
			Name : "Moodys Indonesia",
		},
		{
			Code : "MEDIA",
			Name : "Media Cetak dan Elektronik",
		},
		{
			Code : "MODDY",
			Name : "Quality",
		},
		{
			Code : "PEFIN",
			Name : "Pefindo",
		},
		{
			Code : "S&P",
			Name : "Standard & Poor's",
		},
		{
			Code : "SCPIN",
			Name : "Sucofindo",
		},
		{
			Code : "SPI",
			Name : "Satuan Pengawas Intern (SPI)",
		},
		{
			Code : "SYR",
			Name : "Surveyor",
		},
		{
			Code : "TDK",
			Name : "Tidak Ada",
		},
	}

	var ratingClass = []models.RatingClass{
		{
			ExternalRatingId : 1,
			Code : "AA",
			Name : "Excellent",
			Sibs : "AA",
		},
		{
			ExternalRatingId : 1,
			Code : "BB",
			Name : "Good",
			Sibs : "BB",
		},
		{
			ExternalRatingId : 1,
			Code : "CC",
			Name : "Fair",
			Sibs : "CC",
		},
		{
			ExternalRatingId : 1,
			Code : "DD",
			Name : "Bad Card Holder",
			Sibs : "DD",
		},
		{
			ExternalRatingId : 1,
			Code : "EE",
			Name : "Black List",
			Sibs : "EE",
		},
		{
			ExternalRatingId : 2,
			Code : "01",
			Name : "AAA/ Highest Credit Quality",
			Sibs : "01",
		},
		{
			ExternalRatingId : 2,
			Code : "02",
			Name : "AA+",
			Sibs : "02",
		},
		{
			ExternalRatingId : 2,
			Code : "03",
			Name : "AA/ Very High Credit Quality",
			Sibs : "03",
		},
		{
			ExternalRatingId : 2,
			Code : "04",
			Name : "AA-",
			Sibs : "04",
		},
		{
			ExternalRatingId : 2,
			Code : "05",
			Name : "A+",
			Sibs : "05",
		},
		{
			ExternalRatingId : 2,
			Code : "06",
			Name : "A/ High Credit Quality",
			Sibs : "06",
		},
		{
			ExternalRatingId : 2,
			Code : "07",
			Name : "A-",
			Sibs : "07",
		},
		{
			ExternalRatingId : 2,
			Code : "08",
			Name : "BBB+",
			Sibs : "08",
		},
		{
			ExternalRatingId : 2,
			Code : "09",
			Name : "BBB/ Good Credit Quality",
			Sibs : "09",
		},
		{
			ExternalRatingId : 2,
			Code : "10",
			Name : "BBB-",
			Sibs : "10",
		},
		{
			ExternalRatingId : 2,
			Code : "11",
			Name : "BB+",
			Sibs : "11",
		},
		{
			ExternalRatingId : 2,
			Code : "12",
			Name : "BB/ Speculative",
			Sibs : "12",
		},
		{
			ExternalRatingId : 2,
			Code : "13",
			Name : "BB-",
			Sibs : "13",
		},
		{
			ExternalRatingId : 2,
			Code : "14",
			Name : "B+",
			Sibs : "14",
		},
		{
			ExternalRatingId : 2,
			Code : "15",
			Name : "B/ Highly Speculative",
			Sibs : "15",
		},
		{
			ExternalRatingId : 2,
			Code : "16",
			Name : "B-",
			Sibs : "16",
		},
		{
			ExternalRatingId : 2,
			Code : "17",
			Name : "CCC+",
			Sibs : "17",
		},
		{
			ExternalRatingId : 2,
			Code : "18",
			Name : "CCC/ High Default Risk",
			Sibs : "18",
		},
		{
			ExternalRatingId : 2,
			Code : "19",
			Name : "CCC-",
			Sibs : "19",
		},
		{
			ExternalRatingId : 2,
			Code : "20",
			Name : "CC/ High Default Risk",
			Sibs : "20",
		},
		{
			ExternalRatingId : 2,
			Code : "21",
			Name : "C/ High Default Risk",
			Sibs : "21",
		},
		{
			ExternalRatingId : 2,
			Code : "22",
			Name : "D",
			Sibs : "22",
		},
		{
			ExternalRatingId : 2,
			Code : "61",
			Name : "F1+",
			Sibs : "61",
		},
		{
			ExternalRatingId : 2,
			Code : "A",
			Name : "High Credit Quality",
			Sibs : "A",
		},
		{
			ExternalRatingId : 2,
			Code : "AA",
			Name : "Very High Credit Quality",
			Sibs : "AA",
		},
		{
			ExternalRatingId : 2,
			Code : "AAA",
			Name : "Highest Credit Quality",
			Sibs : "AAA",
		},
		{
			ExternalRatingId : 2,
			Code : "B",
			Name : "Highly Speculative",
			Sibs : "B",
		},
		{
			ExternalRatingId : 2,
			Code : "BB",
			Name : "Speculative",
			Sibs : "BB",
		},
		{
			ExternalRatingId : 2,
			Code : "BBB",
			Name : "Good Credit Quality",
			Sibs : "BBB",
		},
		{
			ExternalRatingId : 2,
			Code : "C",
			Name : "High Default Risk",
			Sibs : "C",
		},
		{
			ExternalRatingId : 2,
			Code : "CC",
			Name : "High Default Risk",
			Sibs : "CC",
		},
		{
			ExternalRatingId : 2,
			Code : "CCC",
			Name : "High Default Risk",
			Sibs : "CCC",
		},
		{
			ExternalRatingId : 2,
			Code : "D",
			Name : "Default and potential recovery < 50%",
			Sibs : "D",
		},
		{
			ExternalRatingId : 2,
			Code : "DD",
			Name : "Default and potential recovery 50%-90%",
			Sibs : "DD",
		},
		{
			ExternalRatingId : 2,
			Code : "DDD",
			Name : "Default",
			Sibs : "DDD",
		},
		{
			ExternalRatingId : 4,
			Code : "ISO",
			Name : "Memenuhi Standard ISO",
			Sibs : "ISO",
		},
		{
			ExternalRatingId : 8,
			Code : "A",
			Name : "Upper-Medium-Grade",
			Sibs : "A",
		},
		{
			ExternalRatingId : 8,
			Code : "AA",
			Name : "High Quality",
			Sibs : "AA",
		},
		{
			ExternalRatingId : 8,
			Code : "AAA",
			Name : "Best Quality",
			Sibs : "AAA",
		},
		{
			ExternalRatingId : 8,
			Code : "B",
			Name : "Lack Characteristics",
			Sibs : "B",
		},
		{
			ExternalRatingId : 8,
			Code : "BA",
			Name : "Hot Well Assured",
			Sibs : "BA",
		},
		{
			ExternalRatingId : 8,
			Code : "BAA",
			Name : "Medium Grade",
			Sibs : "BAA",
		},
		{
			ExternalRatingId : 8,
			Code : "C",
			Name : "Extreamly Poor",
			Sibs : "C",
		},
		{
			ExternalRatingId : 8,
			Code : "CA",
			Name : "Highly Speculative",
			Sibs : "CA",
		},
		{
			ExternalRatingId : 8,
			Code : "CAA",
			Name : "Poor",
			Sibs : "CAA",
		},
		{
			ExternalRatingId : 10,
			Code : "AAA",
			Name : "Excellent",
			Sibs : "AAA",
		},
		{
			ExternalRatingId : 10,
			Code : "BBB",
			Name : "Good",
			Sibs : "BBB",
		},
		{
			ExternalRatingId : 10,
			Code : "CCC",
			Name : "Average",
			Sibs : "CCC",
		},
		{
			ExternalRatingId : 10,
			Code : "DDD",
			Name : "Poor",
			Sibs : "DDD",
		},
		{
			ExternalRatingId : 11,
			Code : "A",
			Name : "High Credit Quality",
			Sibs : "A",
		},
		{
			ExternalRatingId : 11,
			Code : "AA",
			Name : "Very High Credit Quality",
			Sibs : "AA",
		},
		{
			ExternalRatingId : 11,
			Code : "B",
			Name : "Higly Speculative/Poor Credit Quality",
			Sibs : "B",
		},
		{
			ExternalRatingId : 11,
			Code : "BA",
			Name : "Speculation/Questionable Credit Quality",
			Sibs : "BA",
		},
		{
			ExternalRatingId : 11,
			Code : "BAA",
			Name : "Good Credit Quality",
			Sibs : "BAA",
		},
		{
			ExternalRatingId : 11,
			Code : "C",
			Name : "Imminent Default with recovery low",
			Sibs : "C",
		},
		{
			ExternalRatingId : 11,
			Code : "CA",
			Name : "Usually in Default",
			Sibs : "CA",
		},
		{
			ExternalRatingId : 11,
			Code : "CAA",
			Name : "Extremely Poor Credit Quality",
			Sibs : "CAA",
		},
		{
			ExternalRatingId : 13,
			Code : "A",
			Name : "High Credit Quality",
			Sibs : "A",
		},
		{
			ExternalRatingId : 13,
			Code : "AA",
			Name : "Very High Credit Quality",
			Sibs : "AA",
		},
		{
			ExternalRatingId : 13,
			Code : "AAA",
			Name : "Highest Credit Quality",
			Sibs : "AAA",
		},
		{
			ExternalRatingId : 13,
			Code : "B",
			Name : "Highly Default Risk",
			Sibs : "B",
		},
		{
			ExternalRatingId : 13,
			Code : "BB",
			Name : "Speculative",
			Sibs : "BB",
		},
		{
			ExternalRatingId : 13,
			Code : "BBB",
			Name : "Good Credit Quality",
			Sibs : "BBB",
		},
		{
			ExternalRatingId : 13,
			Code : "C",
			Name : "Highly Vulnerable",
			Sibs : "C",
		},
		{
			ExternalRatingId : 13,
			Code : "CC",
			Name : "Highly Vulnerable",
			Sibs : "CC",
		},
		{
			ExternalRatingId : 13,
			Code : "CCC",
			Name : "Vulnerable",
			Sibs : "CCC",
		},
		{
			ExternalRatingId : 13,
			Code : "D",
			Name : "Default",
			Sibs : "D",
		},
		{
			ExternalRatingId : 15,
			Code : "AA",
			Name : "Excellent",
			Sibs : "AA",
		},
		{
			ExternalRatingId : 15,
			Code : "BB",
			Name : "Good",
			Sibs : "BB",
		},
		{
			ExternalRatingId : 15,
			Code : "CC",
			Name : "Fair",
			Sibs : "CC",
		},
		{
			ExternalRatingId : 15,
			Code : "DD",
			Name : "Bad",
			Sibs : "DD",
		},
		{
			ExternalRatingId : 15,
			Code : "EE",
			Name : "Black List",
			Sibs : "EE",
		},
		{
			ExternalRatingId : 17,
			Code : "TDK",
			Name : "Tidak Ada",
			Sibs : "TDK",
		},
	}

	var kodeBursa = []models.KodeBursa{
		{
			Code : "01",
			Name : "001",			
		},
		{
			Code : "02",
			Name : "002",
		},
	}

	var businessType = []models.BusinessType{
		{
			Code : "01",
			Name : "Agriculture / Plantation",
		},
		{
			Code : "02",
			Name : "Fishery / Farm",
		},
		{
			Code : "03",
			Name : "Mining",
		},
		{
			Code : "04",
			Name : "Processing industry",
		},
		{
			Code : "05",
			Name : "Manufacture",
		},
		{
			Code : "06",
			Name : "Energy & Water",
		},
		{
			Code : "06001",
			Name : "Construction",
		},
		{
			Code : "06002",
			Name : "Automotive",
		},
		{
			Code : "06003",
			Name : "Retail/Wholesale trading",
		},
		{
			Code : "06004",
			Name : "Import/export/domestic trading",
		},
		{
			Code : "06005",
			Name : "Transportation",
		},
		{
			Code : "06006",
			Name : "Telecommunication",
		},
		{
			Code : "06007",
			Name : "Financial services",
		},
		{
			Code : "06008",
			Name : "Education Services",
		},
		{
			Code : "06009",
			Name : "Health Services",
		},
		{
			Code : "07",
			Name : "Non-financial services",
		},
		{
			Code : "07001",
			Name : "Hotels & Accommodation",
		},
		{
			Code : "08",
			Name : "Restaurant",
		},
		{
			Code : "08001",
			Name : "Real Estate",
		},
		{
			Code : "08002",
			Name : "Rental",
		},
		{
			Code : "09",
			Name : "Research and development",
		},
		{
			Code : "09001",
			Name : "Publish,Print&Advertising",
		},
		{
			Code : "10",
			Name : "Entertainment Activities",
		},
		{
			Code : "6001",
			Name : "Household",
		},
	}

	connection.DB.Save(&companyFirstName)
	connection.DB.Save(&companyType)
	connection.DB.Save(&externalRatingCompany)
	connection.DB.Save(&ratingClass)
	connection.DB.Save(&kodeBursa)
	connection.DB.Save(&businessType)
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

func SectorEconomySeed() {

	var sectorEconomy1 = []models.SectorEconomy1{

		{
			Code: "A00000",
			Name: "PERTANIAN, KEHUTANAN DAN PERIKANAN",
		},
		{
			Code: "AA0000",
			Name: "RUMAH TANGGA",
		},
		{
			Code: "B00000",
			Name: "PERTAMBANGAN DAN PENGGALIAN",
		},
		{
			Code: "BB0000",
			Name: "BUKAN LAPANGAN USAHA LAINNYA",
		},
		{
			Code: "C00000",
			Name: "INDUSTRI PENGOLAHAN",
		},
		{
			Code: "D00000",
			Name: "PENGADAAN LISTRIK, GAS, UAP/AIR PANAS DAN UDARA DINGIN",
		},
		{
			Code: "E00000",
			Name: "PENGELOLAAN AIR, PENGELOLAAN AIR LIMBAH, PENGELOLAAN DAN DAUR ULANG SAMPAH, DAN AKTIVITAS REMEDIASI",
		},
		{
			Code: "F00000",
			Name: "KONSTRUKSI",
		},
		{
			Code: "G00000",
			Name: "PERDAGANGAN BESAR DAN ECERAN; REPARASI DAN PERAWATAN MOBIL DAN SEPEDA MOTOR",
		},
		{
			Code: "H00000",
			Name: "PENGANGKUTAN DAN PERGUDANGAN",
		},
		{
			Code: "I00000",
			Name: "PENYEDIAAN AKOMODASI DAN PENYEDIAAN MAKAN MINUM",
		},
		{
			Code: "J00000",
			Name: "INFORMASI DAN KOMUNIKASI",
		},
		{
			Code: "K00000",
			Name: "AKTIVITAS KEUANGAN DAN ASURANSI",
		},
		{
			Code: "L00000",
			Name: "REAL ESTAT",
		},
		{
			Code: "M00000",
			Name: "AKTIVITAS PROFESIONAL, ILMIAH DAN TEKNIS",
		},
		{
			Code: "N00000",
			Name: "AKTIVITAS PENYEWAAN DAN SEWA GUNA USAHA TANPA HAK OPSI, KETENAGAKERJAAN, AGEN PERJALANAN DAN PENUNJANG USAHA LAINNYA",
		},
		{
			Code: "O00000",
			Name: "ADMINISTRASI PEMERINTAHAN, PERTAHANAN DAN JAMINAN SOSIAL WAJIB",
		},
		{
			Code: "P00000",
			Name: "PENDIDIKAN",
		},
		{
			Code: "Q00000",
			Name: "AKTIVITAS KESEHATAN MANUSIA DAN AKTIVITAS SOSIAL",
		},
		{
			Code: "R00000",
			Name: "KESENIAN, HIBURAN DAN REKREASI",
		},
		{
			Code: "S00000",
			Name: "AKTIVITAS JASA LAINNYA",
		},
		{
			Code: "T00000",
			Name: "AKTIVITAS RUMAH TANGGA SEBAGAI PEMBERI KERJA; AKTIVITAS YANG MENGHASILKAN BARANG DAN JASA OLEH RUMAH TANGGA YANG DIGUNAKAN UNTUK MEMENUHI KEBUTUHAN SENDIRI",
		},
		{
			Code: "U00000",
			Name: "AKTIVITAS BADAN INTERNASIONAL DAN BADAN EKSTRA INTERNASIONAL LAINNYA",
		},
	}

	var sectorEconomy2 = []models.SectorEconomy2{
		{
			Code : "009000",
			Name : "BUKAN LAPANGAN USAHA LAINNYA",
			SectorEconomy1Id : 4,
		},
		{
			Code : "351001",
			Name : "KETENAGALISTRIKAN PEDESAAN",
			SectorEconomy1Id : 6,
		},
		{
			Code : "351002",
			Name : "KETENAGALISTRIKAN LAINNYA",
			SectorEconomy1Id : 6,
		},
		{
			Code : "352000",
			Name : "PENGADAAN DAN DISTRIBUSI GAS ALAM DAN BUATAN",
			SectorEconomy1Id : 6,
		},
		{
			Code : "353000",
			Name : "PENGADAAN UAP/AIR PANAS, UDARA DINGIN DAN PRODUKSI ES",
			SectorEconomy1Id : 6,
		},
		{
			Code : "360000",
			Name : "PENGELOLAAN AIR",
			SectorEconomy1Id : 7,
		},
		{
			Code : "370000",
			Name : "PENGELOLAAN AIR LIMBAH",
			SectorEconomy1Id : 7,
		},
		{
			Code : "380000",
			Name : "PENGELOLAAN DAN DAUR ULANG SAMPAH",
			SectorEconomy1Id : 7,
		},
		{
			Code : "390000",
			Name : "AKTIVITAS REMEDIASI DAN PENGELOLAAN SAMPAH LAINNYA",
			SectorEconomy1Id : 7,
		},
		{
			Code : "551100",
			Name : "HOTEL BINTANG",
			SectorEconomy1Id : 11,
		},
		{
			Code : "551200",
			Name : "HOTEL MELATI",
			SectorEconomy1Id : 11,
		},
		{
			Code : "559000",
			Name : "PENYEDIAAN AKOMODASI LAINNYA",
			SectorEconomy1Id : 11,
		},
		{
			Code : "561001",
			Name : "RESTORAN DAN RUMAH MAKAN",
			SectorEconomy1Id : 11,
		},
		{
			Code : "561009",
			Name : "PENYEDIAAN MAKANAN DAN MINUMAN LAINNYA",
			SectorEconomy1Id : 11,
		},
		{
			Code : "580000",
			Name : "AKTIVITAS PENERBITAN",
			SectorEconomy1Id : 12,
		},
		{
			Code : "591000",
			Name : "AKTIVITAS PRODUKSI GAMBAR BERGERAK, VIDEO DAN PROGRAM TELEVISI",
			SectorEconomy1Id : 12,
		},
		{
			Code : "592000",
			Name : "AKTIVITAS PEREKAMAN SUARA DAN PENERBITAN MUSIK",
			SectorEconomy1Id : 12,
		},
		{
			Code : "600000",
			Name : "AKTIVITAS PENYIARAN DAN PEMROGRAMAN",
			SectorEconomy1Id : 12,
		},
		{
			Code : "610001",
			Name : "AKTIVITAS TELEKOMUNIKASI DENGAN KABEL, TANPA KABEL DAN SATELIT",
			SectorEconomy1Id : 12,
		},
		{
			Code : "610002",
			Name : "JASA NILAI TAMBAH TELEPONI DAN JASA MULTIMEDIA",
			SectorEconomy1Id : 12,
		},
		{
			Code : "610009",
			Name : "AKTIVITAS TELEKOMUNIKASI LAINNYA YTDL",
			SectorEconomy1Id : 12,
		},
		{
			Code : "620100",
			Name : "AKTIVITAS PEMROGRAMAN KOMPUTER",
			SectorEconomy1Id : 12,
		},
		{
			Code : "620200",
			Name : "AKTIVITAS KONSULTASI KOMPUTER DAN MANAJEMEN FASILITAS KOMPUTER",
			SectorEconomy1Id : 12,
		},
		{
			Code : "631110",
			Name : "AKTIVITAS PENGOLAHAN DATA",
			SectorEconomy1Id : 12,
		},
		{
			Code : "631120",
			Name : "AKTIVITAS HOSTING DAN YBDI",
			SectorEconomy1Id : 12,
		},
		{
			Code : "631210",
			Name : "PORTAL WEB DAN/ATAU PLATFORM DIGITAL TANPA TUJUAN KOMERSIAL",
			SectorEconomy1Id : 12,
		},
		{
			Code : "631220",
			Name : "PORTAL WEB DAN/ATAU PLATFORM DIGITAL DENGAN TUJUAN KOMERSIAL",
			SectorEconomy1Id : 12,
		},
		{
			Code : "639100",
			Name : "AKTIVITAS KANTOR BERITA",
			SectorEconomy1Id : 12,
		},
		{
			Code : "639900",
			Name : "AKTIVITAS JASA INFORMASI LAINNYA YTDL",
			SectorEconomy1Id : 12,
		},
		{
			Code : "641000",
			Name : "PERANTARA MONETER",
			SectorEconomy1Id : 13,
		},
		{
			Code : "649100",
			Name : "SEWA GUNA USAHA DENGAN HAK OPSI",
			SectorEconomy1Id : 13,
		},
		{
			Code : "649900",
			Name : "AKTIVITAS JASA KEUANGAN LAINNYA YTDL, BUKAN ASURANSI DAN DANA PENSIUN",
			SectorEconomy1Id : 13,
		},
		{
			Code : "650000",
			Name : "ASURANSI, REASURANSI DAN DANA PENSIUN, BUKAN JAMINAN SOSIAL WAJIB",
			SectorEconomy1Id : 13,
		},
		{
			Code : "661001",
			Name : "KEGIATAN PENUKARAN VALUTA ASING (MONEY CHANGER)",
			SectorEconomy1Id : 13,
		},
		{
			Code : "661009",
			Name : "AKTIVITAS PENUNJANG JASA KEUANGAN LAINNYA",
			SectorEconomy1Id : 13,
		},
		{
			Code : "662000",
			Name : "AKTIVITAS PENUNJANG ASURANSI DAN DANA PENSIUN",
			SectorEconomy1Id : 13,
		},
		{
			Code : "681101",
			Name : "REAL ESTATE PERUMAHAN SEDERHANA PERUMNAS",
			SectorEconomy1Id : 14,
		},
		{
			Code : "681102",
			Name : "REAL ESTATE PERUMAHAN SEDERHANA PERUMNAS TIPE 21",
			SectorEconomy1Id : 14,
		},
		{
			Code : "681103",
			Name : "REAL ESTATE PERUMAHAN SEDERHANA PERUMNAS TIPE 22 S.D. 70",
			SectorEconomy1Id : 14,
		},
		{
			Code : "681104",
			Name : "REAL ESTATE PERUMAHAN MENENGAH, BESAR ATAU MEWAH (TIPE DIATAS 70)",
			SectorEconomy1Id : 14,
		},
		{
			Code : "681105",
			Name : "REAL ESTATE PERUMAHAN FLAT / APARTEMEN",
			SectorEconomy1Id : 14,
		},
		{
			Code : "681106",
			Name : "REAL ESTATE GEDUNG PERBELANJAAN (MAL, PLAZA)",
			SectorEconomy1Id : 14,
		},
		{
			Code : "681107",
			Name : "REAL ESTATE GEDUNG PERKANTORAN",
			SectorEconomy1Id : 14,
		},
		{
			Code : "681108",
			Name : "REAL ESTATE GEDUNG RUMAH TOKO (RUKO) ATAU RUMAH KANTOR (RUKAN)",
			SectorEconomy1Id : 14,
		},
		{
			Code : "681109",
			Name : "REAL ESTATE LAINNYA",
			SectorEconomy1Id : 14,
		},
		{
			Code : "681200",
			Name : "KAWASAN PARIWISATA",
			SectorEconomy1Id : 14,
		},
		{
			Code : "681300",
			Name : "KAWASAN INDUSTRI",
			SectorEconomy1Id : 14,
		},
		{
			Code : "682000",
			Name : "REAL ESTAT ATAS DASAR BALAS JASA (FEE) ATAU KONTRAK",
			SectorEconomy1Id : 14,
		},
		{
			Code : "690000",
			Name : "AKTIVITAS HUKUM DAN AKUNTANSI",
			SectorEconomy1Id : 15,
		},
		{
			Code : "702010",
			Name : "AKTIVITAS KONSULTASI PARIWISATA",
			SectorEconomy1Id : 15,
		},
		{
			Code : "702090",
			Name : "AKTIVITAS KANTOR PUSAT DAN KONSULTASI MANAJEMEN LAINNYA",
			SectorEconomy1Id : 15,
		},
		{
			Code : "710000",
			Name : "AKTIVITAS ARSITEKTUR DAN KEINSINYURAN; ANALISIS DAN UJI TEKNIS",
			SectorEconomy1Id : 15,
		},
		{
			Code : "721000",
			Name : "PENELITIAN DAN PENGEMBANGAN ILMU PENGETAHUAN ALAM DAN ILMU TEKNOLOGI DAN REKAYASA",
			SectorEconomy1Id : 15,
		},
		{
			Code : "722000",
			Name : "PENELITIAN DAN PENGEMBANGAN ILMU PENGETAHUAN SOSIAL DAN HUMANIORA",
			SectorEconomy1Id : 15,
		},
		{
			Code : "730000",
			Name : "PERIKLANAN DAN PENELITIAN PASAR",
			SectorEconomy1Id : 15,
		},
		{
			Code : "740000",
			Name : "AKTIVITAS PROFESIONAL, ILMIAH DAN TEKNIS LAINNYA",
			SectorEconomy1Id : 15,
		},
		{
			Code : "750000",
			Name : "AKTIVITAS KESEHATAN HEWAN",
			SectorEconomy1Id : 15,
		},
		{
			Code : "841000",
			Name : "ADMINISTRASI PEMERINTAHAN DAN KEBIJAKAN EKONOMI DAN SOSIAL",
			SectorEconomy1Id : 17,
		},
		{
			Code : "842000",
			Name : "PENYEDIAAN LAYANAN UNTUK MASYARAKAT DALAM BIDANG HUBUNGAN LUAR NEGERI, PERTAHANAN, KEAMANAN DAN KETERTIBAN",
			SectorEconomy1Id : 17,
		},
		{
			Code : "843000",
			Name : "JAMINAN SOSIAL WAJIB",
			SectorEconomy1Id : 17,
		},
		{
			Code : "851000",
			Name : "PENDIDIKAN DASAR DAN PENDIDIKAN ANAK USIA DINI",
			SectorEconomy1Id : 18,
		},
		{
			Code : "852000",
			Name : "PENDIDIKAN MENENGAH",
			SectorEconomy1Id : 18,
		},
		{
			Code : "853000",
			Name : "PENDIDIKAN TINGGI",
			SectorEconomy1Id : 18,
		},
		{
			Code : "854000",
			Name : "PENDIDIKAN LAINNYA",
			SectorEconomy1Id : 18,
		},
		{
			Code : "855000",
			Name : "KEGIATAN PENUNJANG PENDIDIKAN",
			SectorEconomy1Id : 18,
		},
		{
			Code : "861000",
			Name : "AKTIVITAS RUMAH SAKIT",
			SectorEconomy1Id : 19,
		},
		{
			Code : "862000",
			Name : "AKTIVITAS PRAKTIK DOKTER DAN DOKTER GIGI",
			SectorEconomy1Id : 19,
		},
		{
			Code : "869000",
			Name : "AKTIVITAS PELAYANAN KESEHATAN MANUSIA LAINNYA",
			SectorEconomy1Id : 19,
		},
		{
			Code : "870000",
			Name : "AKTIVITAS SOSIAL",
			SectorEconomy1Id : 19,
		},
		{
			Code : "900001",
			Name : "JASA IMPRESARIAT BIDANG SENI",
			SectorEconomy1Id : 20,
		},
		{
			Code : "900009",
			Name : "AKTIVITAS HIBURAN, SENI DAN KREATIVITAS LAINNYA",
			SectorEconomy1Id : 20,
		},
		{
			Code : "910100",
			Name : "PERPUSTAKAAN DAN ARSIP",
			SectorEconomy1Id : 20,
		},
		{
			Code : "910200",
			Name : "MUSEUM DAN OPERASIONAL BANGUNAN DAN SITUS BERSEJARAH",
			SectorEconomy1Id : 20,
		},
		{
			Code : "930000",
			Name : "AKTIVITAS OLAHRAGA DAN REKREASI LAINNYA",
			SectorEconomy1Id : 20,
		},
		{
			Code : "941000",
			Name : "AKTIVITAS ORGANISASI BISNIS, PENGUSAHA DAN PROFESI",
			SectorEconomy1Id : 21,
		},
		{
			Code : "942000",
			Name : "AKTIVITAS ORGANISASI BURUH",
			SectorEconomy1Id : 21,
		},
		{
			Code : "949000",
			Name : "AKTIVITAS ORGANISASI KEANGGOTAAN LAINNYA YTDL",
			SectorEconomy1Id : 21,
		},
		{
			Code : "950000",
			Name : "REPARASI KOMPUTER DAN BARANG KEPERLUAN PRIBADI DAN PERLENGKAPAN RUMAH TANGGA",
			SectorEconomy1Id : 21,
		},
		{
			Code : "960001",
			Name : "AKTIVITAS PANTI PIJAT DAN SPA",
			SectorEconomy1Id : 21,
		},
		{
			Code : "960009",
			Name : "AKTIVITAS JASA PERORANGAN LAINNYA",
			SectorEconomy1Id : 21,
		},
		{
			Code : "970000",
			Name : "AKTIVITAS RUMAH TANGGA SEBAGAI PEMBERI KERJA DARI PERSONIL DOMESTIK",
			SectorEconomy1Id : 22,
		},
		{
			Code : "990000",
			Name : "AKTIVITAS BADAN INTERNASIONAL DAN BADAN EKSTRA INTERNASIONAL LAINNYA",
			SectorEconomy1Id : 23,
		},
		{
			Code : "A00001",
			Name : "PERTANIAN",
			SectorEconomy1Id : 1,
		},
		{
			Code : "A00002",
			Name : "PERKEBUNAN",
			SectorEconomy1Id : 1,
		},
		{
			Code : "A00003",
			Name : "BUDIDAYA / PEMBIBITAN",
			SectorEconomy1Id : 1,
		},
		{
			Code : "A00004",
			Name : "PETERNAKAN / PERBURUAN / PENANGKAPAN",
			SectorEconomy1Id : 1,
		},
		{
			Code : "AA0001",
			Name : "PEMILIKAN TEMPAT TINGGAL",
			SectorEconomy1Id : 2,
		},
		{
			Code : "AA0002",
			Name : "PEMILIKAN KENDARAAN",
			SectorEconomy1Id : 2,
		},
		{
			Code : "AA0003",
			Name : "PEMILIKAN PERABOTAN / PERALATAN",
			SectorEconomy1Id : 2,
		},
		{
			Code : "AA0004",
			Name : "MULTIGUNA",
			SectorEconomy1Id : 2,
		},
		{
			Code : "B00001",
			Name : "PERTAMBANGAN BIJIH",
			SectorEconomy1Id : 3,
		},
		{
			Code : "B00002",
			Name : "PERTAMBANGAN NON BIJIH",
			SectorEconomy1Id : 3,
		},
		{
			Code : "C00001",
			Name : "BARANG TAMBANG / LOGAM / MESIN",
			SectorEconomy1Id : 5,
		},
		{
			Code : "C00002",
			Name : "MAKANAN / MINUMAN",
			SectorEconomy1Id : 5,
		},
		{
			Code : "C00003",
			Name : "BERBAHAN DASAR HEWANI",
			SectorEconomy1Id : 5,
		},
		{
			Code : "C00004",
			Name : "BERBAHAN DASAR NABATI",
			SectorEconomy1Id : 5,
		},
		{
			Code : "C00005",
			Name : "BAHAN SANDANG",
			SectorEconomy1Id : 5,
		},
		{
			Code : "C00006",
			Name : "KARET / PLASTIK / KACA / KAYU",
			SectorEconomy1Id : 5,
		},
		{
			Code : "C00007",
			Name : "BAHAN KIMIA",
			SectorEconomy1Id : 5,
		},
		{
			Code : "C00008",
			Name : "BARANG PENGANGKUTAN / TRANSPORTASI",
			SectorEconomy1Id : 5,
		},
		{
			Code : "C00009",
			Name : "BARANG ELEKTRONIK",
			SectorEconomy1Id : 5,
		},
		{
			Code : "C00010",
			Name : "BARANG LAINNYA",
			SectorEconomy1Id : 5,
		},
		{
			Code : "F00001",
			Name : "BANGUNAN",
			SectorEconomy1Id : 8,
		},
		{
			Code : "F00002",
			Name : "INFRASTUKTUR",
			SectorEconomy1Id : 8,
		},
		{
			Code : "G00001",
			Name : "PERDAGANGAN BESAR: MAKANAN/MINUMAN",
			SectorEconomy1Id : 9,
		},
		{
			Code : "G00002",
			Name : "PERDAGANGAN BESAR: HASIL HUTAN/LAUT/SUNGAI/ALAM",
			SectorEconomy1Id : 9,
		},
		{
			Code : "G00003",
			Name : "PERDAGANGAN BESAR: SANDANG",
			SectorEconomy1Id : 9,
		},
		{
			Code : "G00004",
			Name : "PERDAGANGAN BESAR: LAINNYA",
			SectorEconomy1Id : 9,
		},
		{
			Code : "G00005",
			Name : "PERDAGANGAN ECERAN: KHUSUS",
			SectorEconomy1Id : 9,
		},
		{
			Code : "G00006",
			Name : "PERDAGANGAN ECERAN: KAKI LIMA/LOS/KIOS",
			SectorEconomy1Id : 9,
		},
		{
			Code : "G00007",
			Name : "PERDAGANGAN ECERAN: LAINNYA",
			SectorEconomy1Id : 9,
		},
		{
			Code : "G00008",
			Name : "PERDAGANGAN LAINNYA",
			SectorEconomy1Id : 9,
		},
		{
			Code : "H00001",
			Name : "PENGANGKUTAN DARAT",
			SectorEconomy1Id : 10,
		},
		{
			Code : "H00002",
			Name : "PENGANGKUTAN LAUT / SUNGAI",
			SectorEconomy1Id : 10,
		},
		{
			Code : "H00003",
			Name : "PENGANGKUTAN UDARA",
			SectorEconomy1Id : 10,
		},
		{
			Code : "H00004",
			Name : "PENGANGKUTAN DAN PERGUDANGAN",
			SectorEconomy1Id : 10,
		},
		{
			Code : "N00001",
			Name : "AKTIVITAS PENYEWAAN DAN SEWA GUNA USAHA TANPA HAK OPSI",
			SectorEconomy1Id : 16,
		},
		{
			Code : "N00002",
			Name : "KETENAGAKERJAAN, AGEN PERJALANAN DAN PENUNJANG USAHA LAINNYA",
			SectorEconomy1Id : 16,
		},
	}

	var sectorEconomy3 =  []models.SectorEconomy3{
		{
			Code : "001110",
			Name : "RUMAH TANGGA UNTUK PEMILIKAN RUMAH TINGGAL S.D. TIPE 21",
			SectorEconomy2Id : 87,
			Seq : "001110",
			
			
		},
		{
			Code : "001120",
			Name : "RUMAH TANGGA UNTUK PEMILIKAN RUMAH TINGGAL TIPE DIATAS 21 S.D. 70",
			SectorEconomy2Id : 87,
			Seq : "001120",
			
			
		},
		{
			Code : "001130",
			Name : "RUMAH TANGGA UNTUK PEMILIKAN RUMAH TINGGAL TIPE DIATAS 70",
			SectorEconomy2Id : 87,
			Seq : "001130",
			
			
		},
		{
			Code : "001210",
			Name : "RUMAH TANGGA UNTUK PEMILIKAN FLAT ATAU APARTEMEN S.D. TIPE 21",
			SectorEconomy2Id : 87,
			Seq : "001210",
			
			
		},
		{
			Code : "001220",
			Name : "RUMAH TANGGA UNTUK PEMILIKAN FLAT ATAU APARTEMEN TIPE DIATAS 21 S.D. 70",
			SectorEconomy2Id : 87,
			Seq : "001220",
			
			
		},
		{
			Code : "001230",
			Name : "RUMAH TANGGA UNTUK PEMILIKAN FLAT ATAU APARTEMEN TIPE DIATAS 70",
			SectorEconomy2Id : 87,
			Seq : "001230",
			
			
		},
		{
			Code : "001300",
			Name : "RUMAH TANGGA UNTUK PEMILIKAN RUMAH TOKO (RUKO) ATAU RUMAH KANTOR (RUKAN)",
			SectorEconomy2Id : 87,
			Seq : "001300",
			
			
		},
		{
			Code : "002100",
			Name : "RUMAH TANGGA UNTUK PEMILIKAN MOBIL RODA EMPAT",
			SectorEconomy2Id : 88,
			Seq : "002100",
			
			
		},
		{
			Code : "002200",
			Name : "RUMAH TANGGA UNTUK PEMILIKAN SEPEDA BERMOTOR",
			SectorEconomy2Id : 88,
			Seq : "002200",
			
			
		},
		{
			Code : "002300",
			Name : "RUMAH TANGGA UNTUK PEMILIKAN TRUK DAN KENDARAAN BERMOTOR RODA ENAM ATAU LEBIH",
			SectorEconomy2Id : 88,
			Seq : "002300",
			
			
		},
		{
			Code : "002900",
			Name : "RUMAH TANGGA UNTUK PEMILIKAN KENDARAAN BERMOTOR LAINNYA",
			SectorEconomy2Id : 88,
			Seq : "002900",
			
			
		},
		{
			Code : "003100",
			Name : "RUMAH TANGGA UNTUK PEMILIKAN FURNITUR DAN PERALATAN RUMAH TANGGA",
			SectorEconomy2Id : 89,
			Seq : "003100",
			
			
		},
		{
			Code : "003200",
			Name : "RUMAH TANGGA UNTUK PEMILIKAN TELEVISI, RADIO, DAN ALAT ELEKTRONIK",
			SectorEconomy2Id : 89,
			Seq : "003200",
			
			
		},
		{
			Code : "003300",
			Name : "RUMAH TANGGA UNTUK PEMILIKAN KOMPUTER DAN ALAT KOMUNIKASI",
			SectorEconomy2Id : 89,
			Seq : "003300",
			
			
		},
		{
			Code : "003900",
			Name : "RUMAH TANGGA UNTUK PEMILIKAN PERALATAN LAINNYA",
			SectorEconomy2Id : 89,
			Seq : "003900",
			
			
		},
		{
			Code : "004120",
			Name : "RUMAH TANGGA UNTUK KEPERLUAN MULTIGUNA BERAGUNAN RUMAH TINGGAL S.D TIPE 21",
			SectorEconomy2Id : 90,
			Seq : "004120",
			
			
		},
		{
			Code : "004130",
			Name : "RUMAH TANGGA UNTUK KEPERLUAN MULTIGUNA BERAGUNAN RUMAH TINGGAL TIPE DIATAS 21 S.D 70",
			SectorEconomy2Id : 90,
			Seq : "004130",
			
			
		},
		{
			Code : "004140",
			Name : "RUMAH TANGGA UNTUK KEPERLUAN MULTIGUNA BERAGUNAN RUMAH TINGGAL TIPE DIATAS 70",
			SectorEconomy2Id : 90,
			Seq : "004140",
			
			
		},
		{
			Code : "004150",
			Name : "RUMAH TANGGA UNTUK KEPERLUAN MULTIGUNA BERAGUNAN APARTEMEN S.D TIPE 21",
			SectorEconomy2Id : 90,
			Seq : "004150",
			
			
		},
		{
			Code : "004160",
			Name : "RUMAH TANGGA UNTUK KEPERLUAN MULTIGUNA BERAGUNAN APARTEMEN TIPE 22 S.D 70",
			SectorEconomy2Id : 90,
			Seq : "004160",
			
			
		},
		{
			Code : "004170",
			Name : "RUMAH TANGGA UNTUK KEPERLUAN MULTIGUNA BERAGUNAN APARTEMEN TIPE DIATAS 70",
			SectorEconomy2Id : 90,
			Seq : "004170",
			
			
		},
		{
			Code : "004180",
			Name : "RUMAH TANGGA UNTUK KEPERLUAN MULTIGUNA BERAGUNAN RUKO/RUKAN",
			SectorEconomy2Id : 90,
			Seq : "004180",
			
			
		},
		{
			Code : "004190",
			Name : "RUMAH TANGGA UNTUK KEPERLUAN MULTIGUNA LAINNYA",
			SectorEconomy2Id : 90,
			Seq : "004190",
			
			
		},
		{
			Code : "004900",
			Name : "RUMAH TANGGA UNTUK KEPERLUAN YANG TIDAK DIKLASIFIKASIKAN DI TEMPAT LAIN",
			SectorEconomy2Id : 89,
			Seq : "004900",
			
			
		},
		{
			Code : "009000",
			Name : "BUKAN LAPANGAN USAHA LAINNYA",
			SectorEconomy2Id : 1,
			Seq : "009000",
			
			
		},
		{
			Code : "011110",
			Name : "PERTANIAN JAGUNG",
			SectorEconomy2Id : 83,
			Seq : "011110",
			
			
		},
		{
			Code : "011130",
			Name : "PERTANIAN KEDELAI",
			SectorEconomy2Id : 83,
			Seq : "011130",
			
			
		},
		{
			Code : "011140",
			Name : "PERTANIAN KACANG TANAH",
			SectorEconomy2Id : 83,
			Seq : "011140",
			
			
		},
		{
			Code : "011190",
			Name : "PERTANIAN SEREALIA LAINNYA, ANEKA KACANG DAN BIJI-BIJIAN PENGHASIL MINYAK LAINNYA",
			SectorEconomy2Id : 83,
			Seq : "011190",
			
			
		},
		{
			Code : "011200",
			Name : "PERTANIAN PADI",
			SectorEconomy2Id : 83,
			Seq : "011200",
			
			
		},
		{
			Code : "011301",
			Name : "PERTANIAN HORTIKULTURA BAWANG MERAH",
			SectorEconomy2Id : 83,
			Seq : "011301",
			
			
		},
		{
			Code : "011302",
			Name : "PERTANIAN ANEKA UMBI PALAWIJA",
			SectorEconomy2Id : 83,
			Seq : "011302",
			
			
		},
		{
			Code : "011303",
			Name : "PERTANIAN BIT GULA DAN TANAMAN PEMANIS BUKAN TEBU",
			SectorEconomy2Id : 83,
			Seq : "011303",
			
			
		},
		{
			Code : "011309",
			Name : "PERTANIAN SAYURAN, BUAH DAN ANEKA UMBI LAINNYA",
			SectorEconomy2Id : 83,
			Seq : "011309",
			
			
		},
		{
			Code : "011400",
			Name : "PERKEBUNAN TEBU",
			SectorEconomy2Id : 84,
			Seq : "011400",
			
			
		},
		{
			Code : "011500",
			Name : "PERKEBUNAN TEMBAKAU",
			SectorEconomy2Id : 84,
			Seq : "011500",
			
			
		},
		{
			Code : "011600",
			Name : "PERTANIAN TANAMAN BERSERAT",
			SectorEconomy2Id : 83,
			Seq : "011600",
			
			
		},
		{
			Code : "011909",
			Name : "PERTANIAN TANAMAN SEMUSIM LAINNYA YTDL",
			SectorEconomy2Id : 83,
			Seq : "011909",
			
			
		},
		{
			Code : "011930",
			Name : "PERTANIAN TANAMAN BUNGA",
			SectorEconomy2Id : 83,
			Seq : "011930",
			
			
		},
		{
			Code : "011940",
			Name : "PERTANIAN PEMBIBITAN TANAMAN BUNGA",
			SectorEconomy2Id : 83,
			Seq : "011940",
			
			
		},
		{
			Code : "012201",
			Name : "PERTANIAN BUAH PISANG",
			SectorEconomy2Id : 83,
			Seq : "012201",
			
			
		},
		{
			Code : "012209",
			Name : "PERTANIAN BUAH-BUAHAN TROPIS DAN SUBTROPIS LAINNYA",
			SectorEconomy2Id : 83,
			Seq : "012209",
			
			
		},
		{
			Code : "012300",
			Name : "PERTANIAN BUAH JERUK",
			SectorEconomy2Id : 83,
			Seq : "012300",
			
			
		},
		{
			Code : "012400",
			Name : "PERTANIAN BUAH APEL DAN BUAH BATU (POME AND STONE FRUITS)",
			SectorEconomy2Id : 83,
			Seq : "012400",
			
			
		},
		{
			Code : "012500",
			Name : "PERTANIAN SAYURAN DAN BUAH SEMAK DAN BUAH BIJI KACANG-KACANGAN LAINNYA",
			SectorEconomy2Id : 83,
			Seq : "012500",
			
			
		},
		{
			Code : "012610",
			Name : "PERKEBUNAN BUAH KELAPA",
			SectorEconomy2Id : 84,
			Seq : "012610",
			
			
		},
		{
			Code : "012620",
			Name : "PERKEBUNAN BUAH KELAPA SAWIT",
			SectorEconomy2Id : 84,
			Seq : "012620",
			
			
		},
		{
			Code : "012690",
			Name : "PERKEBUNAN BUAH OLEAGINOUS LAINNYA",
			SectorEconomy2Id : 84,
			Seq : "012690",
			
			
		},
		{
			Code : "012701",
			Name : "PERKEBUNAN TANAMAN KOPI",
			SectorEconomy2Id : 84,
			Seq : "012701",
			
			
		},
		{
			Code : "012702",
			Name : "PERKEBUNAN TANAMAN TEH",
			SectorEconomy2Id : 84,
			Seq : "012702",
			
			
		},
		{
			Code : "012703",
			Name : "PERKEBUNAN TANAMAN COKLAT (KAKAO)",
			SectorEconomy2Id : 84,
			Seq : "012703",
			
			
		},
		{
			Code : "012709",
			Name : "PERTANIAN TANAMAN UNTUK BAHAN MINUMAN LAINNYA",
			SectorEconomy2Id : 83,
			Seq : "012709",
			
			
		},
		{
			Code : "012810",
			Name : "PERKEBUNAN LADA",
			SectorEconomy2Id : 84,
			Seq : "012810",
			
			
		},
		{
			Code : "012820",
			Name : "PERKEBUNAN CENGKEH",
			SectorEconomy2Id : 84,
			Seq : "012820",
			
			
		},
		{
			Code : "012830",
			Name : "PERTANIAN CABAI",
			SectorEconomy2Id : 83,
			Seq : "012830",
			
			
		},
		{
			Code : "012840",
			Name : "PERKEBUNAN TANAMAN AROMATIK/PENYEGAR",
			SectorEconomy2Id : 84,
			Seq : "012840",
			
			
		},
		{
			Code : "012850",
			Name : "PERKEBUNAN TANAMAN OBAT / BAHAN FARMASI",
			SectorEconomy2Id : 84,
			Seq : "012850",
			
			
		},
		{
			Code : "012891",
			Name : "PERKEBUNAN TANAMAN REMPAH PANILI",
			SectorEconomy2Id : 84,
			Seq : "012891",
			
			
		},
		{
			Code : "012892",
			Name : "PERKEBUNAN TANAMAN REMPAH PALA",
			SectorEconomy2Id : 84,
			Seq : "012892",
			
			
		},
		{
			Code : "012899",
			Name : "PERKEBUNAN TANAMAN REMPAH YANG TIDAK DIKLASIFIKASIKAN DI TEMPAT LAIN",
			SectorEconomy2Id : 84,
			Seq : "012899",
			
			
		},
		{
			Code : "012910",
			Name : "PERKEBUNAN KARET DAN TANAMAN PENGHASIL GETAH LAINNYA",
			SectorEconomy2Id : 84,
			Seq : "012910",
			
			
		},
		{
			Code : "012990",
			Name : "PERTANIAN CEMARA DAN TANAMAN TAHUNAN LAINNYA",
			SectorEconomy2Id : 83,
			Seq : "012990",
			
			
		},
		{
			Code : "013010",
			Name : "PERTANIAN TANAMAN HIAS",
			SectorEconomy2Id : 83,
			Seq : "013010",
			
			
		},
		{
			Code : "013020",
			Name : "PERTANIAN PENGEMBANGBIAKAN TANAMAN",
			SectorEconomy2Id : 83,
			Seq : "013020",
			
			
		},
		{
			Code : "014110",
			Name : "PEMBIBITAN DAN BUDIDAYA SAPI POTONG",
			SectorEconomy2Id : 85,
			Seq : "014110",
			
			
		},
		{
			Code : "014120",
			Name : "PEMBIBITAN DAN BUDIDAYA SAPI PERAH",
			SectorEconomy2Id : 85,
			Seq : "014120",
			
			
		},
		{
			Code : "014130",
			Name : "PEMBIBITAN DAN BUDIDAYA KERBAU POTONG",
			SectorEconomy2Id : 85,
			Seq : "014130",
			
			
		},
		{
			Code : "014140",
			Name : "PEMBIBITAN DAN BUDIDAYA KERBAU PERAH",
			SectorEconomy2Id : 85,
			Seq : "014140",
			
			
		},
		{
			Code : "014400",
			Name : "PETERNAKAN DOMBA DAN KAMBING",
			SectorEconomy2Id : 86,
			Seq : "014400",
			
			
		},
		{
			Code : "014500",
			Name : "PETERNAKAN BABI",
			SectorEconomy2Id : 86,
			Seq : "014500",
			
			
		},
		{
			Code : "014600",
			Name : "PETERNAKAN UNGGAS",
			SectorEconomy2Id : 86,
			Seq : "014600",
			
			
		},
		{
			Code : "014900",
			Name : "PETERNAKAN LAINNYA",
			SectorEconomy2Id : 86,
			Seq : "014900",
			
			
		},
		{
			Code : "016000",
			Name : "JASA PENUNJANG PERTANIAN DAN PASCA PANEN",
			SectorEconomy2Id : 83,
			Seq : "016000",
			
			
		},
		{
			Code : "017000",
			Name : "PERBURUAN, PENANGKAPAN DAN PENANGKARAN TUMBUHAN/ SATWA LIAR",
			SectorEconomy2Id : 86,
			Seq : "017000",
			
			
		},
		{
			Code : "021100",
			Name : "PENGUSAHAAN HUTAN TANAMAN",
			SectorEconomy2Id : 86,
			Seq : "021100",
			
			
		},
		{
			Code : "021200",
			Name : "PENGUSAHAAN HUTAN ALAM",
			SectorEconomy2Id : 86,
			Seq : "021200",
			
			
		},
		{
			Code : "021300",
			Name : "PENGUSAHAAN HASIL HUTAN BUKAN KAYU",
			SectorEconomy2Id : 86,
			Seq : "021300",
			
			
		},
		{
			Code : "021400",
			Name : "PENGUSAHAAN PEMBIBITAN TANAMAN KEHUTANAN",
			SectorEconomy2Id : 86,
			Seq : "021400",
			
			
		},
		{
			Code : "022090",
			Name : "USAHA KEHUTANAN LAINNYA",
			SectorEconomy2Id : 86,
			Seq : "022090",
			
			
		},
		{
			Code : "024000",
			Name : "JASA PENUNJANG KEHUTANAN",
			SectorEconomy2Id : 86,
			Seq : "024000",
			
			
		},
		{
			Code : "031111",
			Name : "PENANGKAPAN IKAN TUNA",
			SectorEconomy2Id : 86,
			Seq : "031111",
			
			
		},
		{
			Code : "031119",
			Name : "PENANGKAPAN IKAN LAINNYA",
			SectorEconomy2Id : 86,
			Seq : "031119",
			
			
		},
		{
			Code : "031121",
			Name : "PENANGKAPAN UDANG LAUT",
			SectorEconomy2Id : 86,
			Seq : "031121",
			
			
		},
		{
			Code : "031129",
			Name : "PENANGKAPAN CRUSTACEA LAINNYA DI LAUT",
			SectorEconomy2Id : 86,
			Seq : "031129",
			
			
		},
		{
			Code : "031190",
			Name : "PENANGKAPAN BIOTA AIR LAINNYA DI LAUT",
			SectorEconomy2Id : 86,
			Seq : "031190",
			
			
		},
		{
			Code : "031210",
			Name : "PENANGKAPAN PISCES/IKAN BERSIRIP DI PERAIRAN UMUM",
			SectorEconomy2Id : 86,
			Seq : "031210",
			
			
		},
		{
			Code : "031290",
			Name : "PENANGKAPAN BIOTA AIR LAINNYA DI PERAIRAN UMUM",
			SectorEconomy2Id : 86,
			Seq : "031290",
			
			
		},
		{
			Code : "031300",
			Name : "JASA PENANGKAPAN IKAN DI LAUT",
			SectorEconomy2Id : 86,
			Seq : "031300",
			
			
		},
		{
			Code : "031400",
			Name : "JASA PENANGKAPAN IKAN DI PERAIRAN UMUM",
			SectorEconomy2Id : 86,
			Seq : "031400",
			
			
		},
		{
			Code : "032101",
			Name : "BUDIDAYA BIOTA LAUT UDANG",
			SectorEconomy2Id : 85,
			Seq : "032101",
			
			
		},
		{
			Code : "032102",
			Name : "BUDIDAYA BIOTA LAUT RUMPUT LAUT",
			SectorEconomy2Id : 85,
			Seq : "032102",
			
			
		},
		{
			Code : "032109",
			Name : "BUDIDAYA BIOTA LAUT LAINNYA",
			SectorEconomy2Id : 85,
			Seq : "032109",
			
			
		},
		{
			Code : "032201",
			Name : "BUDIDAYA BIOTA AIR TAWAR UDANG",
			SectorEconomy2Id : 85,
			Seq : "032201",
			
			
		},
		{
			Code : "032202",
			Name : "PEMBENIHAN IKAN AIR TAWAR",
			SectorEconomy2Id : 86,
			Seq : "032202",
			
			
		},
		{
			Code : "032209",
			Name : "BUDIDAYA BIOTA AIR TAWAR LAINNYA",
			SectorEconomy2Id : 85,
			Seq : "032209",
			
			
		},
		{
			Code : "032300",
			Name : "JASA BUDIDAYA IKAN LAUT",
			SectorEconomy2Id : 85,
			Seq : "032300",
			
			
		},
		{
			Code : "032400",
			Name : "JASA BUDIDAYA IKAN AIR TAWAR",
			SectorEconomy2Id : 85,
			Seq : "032400",
			
			
		},
		{
			Code : "032501",
			Name : "BUDIDAYA BIOTA AIR PAYAU UDANG",
			SectorEconomy2Id : 85,
			Seq : "032501",
			
			
		},
		{
			Code : "032509",
			Name : "BUDIDAYA BIOTA AIR PAYAU LAINNYA",
			SectorEconomy2Id : 85,
			Seq : "032509",
			
			
		},
		{
			Code : "032600",
			Name : "JASA BUDIDAYA IKAN AIR PAYAU",
			SectorEconomy2Id : 85,
			Seq : "032600",
			
			
		},
		{
			Code : "050000",
			Name : "PERTAMBANGAN BATU BARA DAN LIGNIT",
			SectorEconomy2Id : 92,
			Seq : "050000",
			
			
		},
		{
			Code : "060001",
			Name : "PERTAMBANGAN MINYAK BUMI DAN GAS ALAM",
			SectorEconomy2Id : 92,
			Seq : "060001",
			
			
		},
		{
			Code : "060002",
			Name : "PENGUSAHAAN TENAGA PANAS BUMI",
			SectorEconomy2Id : 92,
			Seq : "060002",
			
			
		},
		{
			Code : "071000",
			Name : "PERTAMBANGAN PASIR BESI DAN BIJIH BESI",
			SectorEconomy2Id : 91,
			Seq : "071000",
			
			
		},
		{
			Code : "072100",
			Name : "PERTAMBANGAN BIJIH URANIUM DAN THORIUM",
			SectorEconomy2Id : 91,
			Seq : "072100",
			
			
		},
		{
			Code : "072910",
			Name : "PERTAMBANGAN BIJIH TIMAH",
			SectorEconomy2Id : 91,
			Seq : "072910",
			
			
		},
		{
			Code : "072930",
			Name : "PERTAMBANGAN BIJIH BAUKSIT/ALUMINIUM",
			SectorEconomy2Id : 91,
			Seq : "072930",
			
			
		},
		{
			Code : "072940",
			Name : "PERTAMBANGAN BIJIH TEMBAGA",
			SectorEconomy2Id : 91,
			Seq : "072940",
			
			
		},
		{
			Code : "072950",
			Name : "PERTAMBANGAN BIJIH NIKEL",
			SectorEconomy2Id : 91,
			Seq : "072950",
			
			
		},
		{
			Code : "072990",
			Name : "PERTAMBANGAN BAHAN GALIAN LAINNYA YANG TIDAK MENGANDUNG BIJIH BESI",
			SectorEconomy2Id : 91,
			Seq : "072990",
			
			
		},
		{
			Code : "073011",
			Name : "PERTAMBANGAN EMAS",
			SectorEconomy2Id : 92,
			Seq : "073011",
			
			
		},
		{
			Code : "073012",
			Name : "PERTAMBANGAN PERAK",
			SectorEconomy2Id : 92,
			Seq : "073012",
			
			
		},
		{
			Code : "073090",
			Name : "PERTAMBANGAN BIJIH LOGAM MULIA LAINNYA",
			SectorEconomy2Id : 91,
			Seq : "073090",
			
			
		},
		{
			Code : "081000",
			Name : "PENGGALIAN BATU, PASIR DAN TANAH LIAT",
			SectorEconomy2Id : 92,
			Seq : "081000",
			
			
		},
		{
			Code : "089100",
			Name : "PERTAMBANGAN MINERAL, BAHAN KIMIA DAN BAHAN PUPUK",
			SectorEconomy2Id : 92,
			Seq : "089100",
			
			
		},
		{
			Code : "089300",
			Name : "EKSTRAKSI GARAM",
			SectorEconomy2Id : 92,
			Seq : "089300",
			
			
		},
		{
			Code : "089900",
			Name : "PERTAMBANGAN DAN PENGGALIAN LAINNYA YTDL",
			SectorEconomy2Id : 92,
			Seq : "089900",
			
			
		},
		{
			Code : "091000",
			Name : "AKTIVITAS PENUNJANG PERTAMBANGAN MINYAK BUMI DAN GAS ALAM",
			SectorEconomy2Id : 92,
			Seq : "091000",
			
			
		},
		{
			Code : "099000",
			Name : "AKTIVITAS PENUNJANG PERTAMBANGAN DAN PENGGALIAN LAINNYA",
			SectorEconomy2Id : 92,
			Seq : "099000",
			
			
		},
		{
			Code : "101000",
			Name : "INDUSTRI PENGOLAHAN DAN PENGAWETAN DAGING",
			SectorEconomy2Id : 95,
			Seq : "101000",
			
			
		},
		{
			Code : "102000",
			Name : "INDUSTRI PENGOLAHAN DAN PENGAWETAN IKAN DAN BIOTA AIR",
			SectorEconomy2Id : 95,
			Seq : "102000",
			
			
		},
		{
			Code : "103001",
			Name : "INDUSTRI TEMPE & TAHU KEDELAI",
			SectorEconomy2Id : 96,
			Seq : "103001",
			
			
		},
		{
			Code : "103009",
			Name : "INDUSTRI PENGOLAHAN DAN PENGAWETAN LAINNYA BUAH-BUAHAN DAN SAYURAN",
			SectorEconomy2Id : 96,
			Seq : "103009",
			
			
		},
		{
			Code : "104100",
			Name : "INDUSTRI MINYAK DAN LEMAK NABATI DAN HEWANI",
			SectorEconomy2Id : 96,
			Seq : "104100",
			
			
		},
		{
			Code : "104210",
			Name : "INDUSTRI KOPRA, TEPUNG & PELET KELAPA",
			SectorEconomy2Id : 96,
			Seq : "104210",
			
			
		},
		{
			Code : "104230",
			Name : "INDUSTRI MINYAK MENTAH KELAPA & MINYAK GORENG KELAPA",
			SectorEconomy2Id : 96,
			Seq : "104230",
			
			
		},
		{
			Code : "104300",
			Name : "INDUSTRI MINYAK MENTAH/MURNI KELAPA SAWIT (CRUDE PALM OIL) DAN MINYAK GORENG KELAPA SAWIT",
			SectorEconomy2Id : 96,
			Seq : "104300",
			
			
		},
		{
			Code : "104900",
			Name : "INDUSTRI MINYAK MENTAH DAN LEMAK NABATI DAN HEWANI LAINNYA",
			SectorEconomy2Id : 95,
			Seq : "104900",
			
			
		},
		{
			Code : "105000",
			Name : "INDUSTRI PENGOLAHAN SUSU, PRODUK DARI SUSU DAN ES KRIM",
			SectorEconomy2Id : 95,
			Seq : "105000",
			
			
		},
		{
			Code : "106100",
			Name : "INDUSTRI PENGGILINGAN SERELIA DAN BIJI-BIJIAN LAINNYA (BUKAN BERAS DAN JAGUNG)",
			SectorEconomy2Id : 96,
			Seq : "106100",
			
			
		},
		{
			Code : "106200",
			Name : "INDUSTRI PATI DAN PRODUK PATI (BUKAN BERAS DAN JAGUNG)",
			SectorEconomy2Id : 96,
			Seq : "106200",
			
			
		},
		{
			Code : "106300",
			Name : "INDUSTRI PENGGILINGAN BERAS DAN JAGUNG DAN INDUSTRI TEPUNG BERAS DAN JAGUNG",
			SectorEconomy2Id : 96,
			Seq : "106300",
			
			
		},
		{
			Code : "107100",
			Name : "INDUSTRI PRODUK ROTI DAN KUE",
			SectorEconomy2Id : 94,
			Seq : "107100",
			
			
		},
		{
			Code : "107200",
			Name : "INDUSTRI GULA",
			SectorEconomy2Id : 94,
			Seq : "107200",
			
			
		},
		{
			Code : "107300",
			Name : "INDUSTRI KAKAO, COKELAT DAN KEMBANG GULA",
			SectorEconomy2Id : 94,
			Seq : "107300",
			
			
		},
		{
			Code : "107400",
			Name : "INDUSTRI MAKARONI, MIE DAN PRODUK SEJENISNYA",
			SectorEconomy2Id : 94,
			Seq : "107400",
			
			
		},
		{
			Code : "107610",
			Name : "INDUSTRI PENGOLAHAN KOPI",
			SectorEconomy2Id : 94,
			Seq : "107610",
			
			
		},
		{
			Code : "107630",
			Name : "INDUSTRI PENGOLAHAN TEH",
			SectorEconomy2Id : 94,
			Seq : "107630",
			
			
		},
		{
			Code : "107710",
			Name : "INDUSTRI KECAP",
			SectorEconomy2Id : 94,
			Seq : "107710",
			
			
		},
		{
			Code : "107900",
			Name : "INDUSTRI PRODUK MAKANAN LAINNYA",
			SectorEconomy2Id : 94,
			Seq : "107900",
			
			
		},
		{
			Code : "108000",
			Name : "INDUSTRI MAKANAN HEWAN",
			SectorEconomy2Id : 95,
			Seq : "108000",
			
			
		},
		{
			Code : "110000",
			Name : "INDUSTRI MINUMAN",
			SectorEconomy2Id : 94,
			Seq : "110000",
			
			
		},
		{
			Code : "120100",
			Name : "INDUSTRI ROKOK DAN PRODUK TEMBAKAU LAINNYA",
			SectorEconomy2Id : 96,
			Seq : "120100",
			
			
		},
		{
			Code : "120900",
			Name : "INDUSTRI PENGOLAHAN TEMBAKAU LAINNYA",
			SectorEconomy2Id : 96,
			Seq : "120900",
			
			
		},
		{
			Code : "131000",
			Name : "INDUSTRI PEMINTALAN, PENENUNAN DAN PENYELESAIAN AKHIR TEKSTIL",
			SectorEconomy2Id : 97,
			Seq : "131000",
			
			
		},
		{
			Code : "139000",
			Name : "INDUSTRI TEKSTIL LAINNYA",
			SectorEconomy2Id : 97,
			Seq : "139000",
			
			
		},
		{
			Code : "141000",
			Name : "INDUSTRI PAKAIAN JADI DAN PERLENGKAPANNYA, BUKAN PAKAIAN JADI DARI KULIT BERBULU",
			SectorEconomy2Id : 97,
			Seq : "141000",
			
			
		},
		{
			Code : "142000",
			Name : "INDUSTRI PAKAIAN JADI DAN BARANG DARI KULIT BERBULU",
			SectorEconomy2Id : 97,
			Seq : "142000",
			
			
		},
		{
			Code : "143000",
			Name : "INDUSTRI PAKAIAN JADI RAJUTAN DAN SULAMAN/BORDIR",
			SectorEconomy2Id : 97,
			Seq : "143000",
			
			
		},
		{
			Code : "151000",
			Name : "INDUSTRI KULIT DAN BARANG DARI KULIT, TERMASUK KULIT BUATAN",
			SectorEconomy2Id : 95,
			Seq : "151000",
			
			
		},
		{
			Code : "152000",
			Name : "INDUSTRI ALAS KAKI",
			SectorEconomy2Id : 97,
			Seq : "152000",
			
			
		},
		{
			Code : "161000",
			Name : "INDUSTRI PENGGERGAJIAN DAN PENGAWETAN KAYU, ROTAN, BAMBU DAN SEJENISNYA",
			SectorEconomy2Id : 98,
			Seq : "161000",
			
			
		},
		{
			Code : "162100",
			Name : "INDUSTRI KAYU LAPIS, VENEER DAN SEJENISNYA",
			SectorEconomy2Id : 98,
			Seq : "162100",
			
			
		},
		{
			Code : "162900",
			Name : "INDUSTRI BARANG LAINNYA DARI KAYU; INDUSTRI BARANG DARI GABUS DAN BARANG ANYAMAN DARI JERAMI, ROTAN, BAMBU DAN SEJENISNYA",
			SectorEconomy2Id : 98,
			Seq : "162900",
			
			
		},
		{
			Code : "170100",
			Name : "INDUSTRI BUBUR KERTAS, KERTAS DAN PAPAN KERTAS",
			SectorEconomy2Id : 98,
			Seq : "170100",
			
			
		},
		{
			Code : "170200",
			Name : "INDUSTRI KERTAS DAN PAPAN KERTAS BERGELOMBANG DAN WADAH DARI KERTAS DAN PAPAN KERTAS",
			SectorEconomy2Id : 98,
			Seq : "170200",
			
			
		},
		{
			Code : "170900",
			Name : "INDUSTRI BARANG DARI KERTAS DAN PAPAN KERTAS LAINNYA",
			SectorEconomy2Id : 98,
			Seq : "170900",
			
			
		},
		{
			Code : "181000",
			Name : "INDUSTRI PENCETAKAN DAN KEGIATAN YBDI",
			SectorEconomy2Id : 102,
			Seq : "181000",
			
			
		},
		{
			Code : "182000",
			Name : "REPRODUKSI MEDIA REKAMAN",
			SectorEconomy2Id : 102,
			Seq : "182000",
			
			
		},
		{
			Code : "191000",
			Name : "INDUSTRI PRODUK DARI BATU BARA",
			SectorEconomy2Id : 93,
			Seq : "191000",
			
			
		},
		{
			Code : "192100",
			Name : "INDUSTRI BAHAN BAKAR DAN MINYAK PELUMAS HASIL PENGILANGAN MINYAK BUMI",
			SectorEconomy2Id : 93,
			Seq : "192100",
			
			
		},
		{
			Code : "192900",
			Name : "INDUSTRI BRIKET BATU BARA",
			SectorEconomy2Id : 93,
			Seq : "192900",
			
			
		},
		{
			Code : "201100",
			Name : "INDUSTRI KIMIA DASAR",
			SectorEconomy2Id : 99,
			Seq : "201100",
			
			
		},
		{
			Code : "201200",
			Name : "INDUSTRI PUPUK DAN BAHAN SENYAWA NITROGEN",
			SectorEconomy2Id : 99,
			Seq : "201200",
			
			
		},
		{
			Code : "201300",
			Name : "INDUSTRI PLASTIK DAN KARET BUATAN DALAM BENTUK DASAR",
			SectorEconomy2Id : 98,
			Seq : "201300",
			
			
		},
		{
			Code : "202100",
			Name : "INDUSTRI PESTISIDA DAN PRODUK AGROKIMIA LAINNYA",
			SectorEconomy2Id : 99,
			Seq : "202100",
			
			
		},
		{
			Code : "202200",
			Name : "INDUSTRI CAT DAN TINTA CETAK, PERNIS DAN BAHAN PELAPISAN SEJENISNYA DAN LAK",
			SectorEconomy2Id : 99,
			Seq : "202200",
			
			
		},
		{
			Code : "202300",
			Name : "INDUSTRI SABUN DAN DETERJEN, BAHAN PEMBERSIH DAN PENGILAP, PARFUM DAN KOSMETIK",
			SectorEconomy2Id : 99,
			Seq : "202300",
			
			
		},
		{
			Code : "202940",
			Name : "INDUSTRI MINYAK ATSIRI",
			SectorEconomy2Id : 99,
			Seq : "202940",
			
			
		},
		{
			Code : "202990",
			Name : "INDUSTRI BARANG KIMIA LAINNYA YTDL",
			SectorEconomy2Id : 99,
			Seq : "202990",
			
			
		},
		{
			Code : "203000",
			Name : "INDUSTRI SERAT BUATAN",
			SectorEconomy2Id : 102,
			Seq : "203000",
			
			
		},
		{
			Code : "210000",
			Name : "INDUSTRI FARMASI, PRODUK OBAT KIMIA DAN OBAT TRADISIONAL",
			SectorEconomy2Id : 99,
			Seq : "210000",
			
			
		},
		{
			Code : "221210",
			Name : "INDUSTRI PENGASAPAN KARET",
			SectorEconomy2Id : 98,
			Seq : "221210",
			
			
		},
		{
			Code : "221220",
			Name : "INDUSTRI REMILLING KARET",
			SectorEconomy2Id : 98,
			Seq : "221220",
			
			
		},
		{
			Code : "221230",
			Name : "INDUSTRI KARET REMAH (CRUMB RUBBER)",
			SectorEconomy2Id : 98,
			Seq : "221230",
			
			
		},
		{
			Code : "221900",
			Name : "INDUSTRI BARANG DARI KARET LAINNYA",
			SectorEconomy2Id : 98,
			Seq : "221900",
			
			
		},
		{
			Code : "222000",
			Name : "INDUSTRI BARANG DARI PLASTIK",
			SectorEconomy2Id : 98,
			Seq : "222000",
			
			
		},
		{
			Code : "231000",
			Name : "INDUSTRI KACA DAN BARANG DARI KACA",
			SectorEconomy2Id : 98,
			Seq : "231000",
			
			
		},
		{
			Code : "239200",
			Name : "INDUSTRI BAHAN BANGUNAN DARI TANAH LIAT/KERAMIK",
			SectorEconomy2Id : 98,
			Seq : "239200",
			
			
		},
		{
			Code : "239301",
			Name : "INDUSTRI BARANG PORSELEN BUKAN BAHAN BANGUNAN",
			SectorEconomy2Id : 98,
			Seq : "239301",
			
			
		},
		{
			Code : "239302",
			Name : "INDUSTRI BARANG TANAH LIAT/KERAMIK BUKAN BAHAN BANGUNAN",
			SectorEconomy2Id : 98,
			Seq : "239302",
			
			
		},
		{
			Code : "239400",
			Name : "INDUSTRI SEMEN, KAPUR DAN GIPS",
			SectorEconomy2Id : 93,
			Seq : "239400",
			
			
		},
		{
			Code : "239600",
			Name : "INDUSTRI BARANG DARI BATU",
			SectorEconomy2Id : 93,
			Seq : "239600",
			
			
		},
		{
			Code : "239900",
			Name : "INDUSTRI BARANG GALIAN BUKAN LOGAM LAINNYA YTDL",
			SectorEconomy2Id : 93,
			Seq : "239900",
			
			
		},
		{
			Code : "241000",
			Name : "INDUSTRI LOGAM DASAR BESI DAN BAJA",
			SectorEconomy2Id : 93,
			Seq : "241000",
			
			
		},
		{
			Code : "242060",
			Name : "INDUSTRI PENGOLAHAN URANIUM DAN BIJIH URANIUM",
			SectorEconomy2Id : 93,
			Seq : "242060",
			
			
		},
		{
			Code : "242090",
			Name : "INDUSTRI LOGAM DASAR MULIA DAN LOGAM DASAR BUKAN BESI LAINNYA",
			SectorEconomy2Id : 93,
			Seq : "242090",
			
			
		},
		{
			Code : "243100",
			Name : "INDUSTRI PENGECORAN BESI DAN BAJA",
			SectorEconomy2Id : 93,
			Seq : "243100",
			
			
		},
		{
			Code : "243200",
			Name : "INDUSTRI PENGECORAN LOGAM BUKAN BESI DAN BAJA",
			SectorEconomy2Id : 93,
			Seq : "243200",
			
			
		},
		{
			Code : "251000",
			Name : "INDUSTRI BARANG LOGAM SIAP PASANG UNTUK BANGUNAN, TANGKI, TANDON AIR DAN GENERATOR UAP",
			SectorEconomy2Id : 93,
			Seq : "251000",
			
			
		},
		{
			Code : "259300",
			Name : "INDUSTRI ALAT POTONG, PERKAKAS TANGAN DAN PERALATAN UMUM",
			SectorEconomy2Id : 102,
			Seq : "259300",
			
			
		},
		{
			Code : "259900",
			Name : "INDUSTRI BARANG LOGAM LAINNYA YTDL",
			SectorEconomy2Id : 93,
			Seq : "259900",
			
			
		},
		{
			Code : "261000",
			Name : "INDUSTRI KOMPONEN DAN PAPAN ELEKTRONIK",
			SectorEconomy2Id : 101,
			Seq : "261000",
			
			
		},
		{
			Code : "262000",
			Name : "INDUSTRI KOMPUTER DAN PERLENGKAPANNYA",
			SectorEconomy2Id : 101,
			Seq : "262000",
			
			
		},
		{
			Code : "263000",
			Name : "INDUSTRI PERALATAN KOMUNIKASI",
			SectorEconomy2Id : 101,
			Seq : "263000",
			
			
		},
		{
			Code : "264000",
			Name : "INDUSTRI PERALATAN AUDIO DAN VIDEO ELEKTRONIK",
			SectorEconomy2Id : 101,
			Seq : "264000",
			
			
		},
		{
			Code : "265100",
			Name : "INDUSTRI ALAT UKUR, ALAT UJI, PERALATAN NAVIGASI DAN KONTROL",
			SectorEconomy2Id : 102,
			Seq : "265100",
			
			
		},
		{
			Code : "265200",
			Name : "INDUSTRI ALAT UKUR WAKTU",
			SectorEconomy2Id : 102,
			Seq : "265200",
			
			
		},
		{
			Code : "266000",
			Name : "INDUSTRI PERALATAN IRADIASI, ELEKTROMEDIKAL DAN ELEKTROTERAPI",
			SectorEconomy2Id : 101,
			Seq : "266000",
			
			
		},
		{
			Code : "267000",
			Name : "INDUSTRI PERALATAN FOTOGRAFI DAN INSTRUMEN OPTIK BUKAN KACA MATA",
			SectorEconomy2Id : 98,
			Seq : "267000",
			
			
		},
		{
			Code : "269000",
			Name : "INDUSTRI KOMPUTER, BARANG ELEKTRONIK DAN OPTIK LAINNYA",
			SectorEconomy2Id : 101,
			Seq : "269000",
			
			
		},
		{
			Code : "271100",
			Name : "INDUSTRI MOTOR LISTRIK, GENERATOR DAN TRANSFORMATOR",
			SectorEconomy2Id : 101,
			Seq : "271100",
			
			
		},
		{
			Code : "271200",
			Name : "INDUSTRI PERALATAN PENGONTROL DAN PENDISTRIBUSIAN LISTRIK",
			SectorEconomy2Id : 101,
			Seq : "271200",
			
			
		},
		{
			Code : "272000",
			Name : "INDUSTRI BATU BATERAI DAN AKUMULATOR LISTRIK",
			SectorEconomy2Id : 101,
			Seq : "272000",
			
			
		},
		{
			Code : "273000",
			Name : "INDUSTRI KABEL DAN PERLENGKAPANNYA",
			SectorEconomy2Id : 101,
			Seq : "273000",
			
			
		},
		{
			Code : "274000",
			Name : "INDUSTRI PERALATAN PENERANGAN LISTRIK (TERMASUK PERALATAN PENERANGAN BUKAN LISTRIK)",
			SectorEconomy2Id : 101,
			Seq : "274000",
			
			
		},
		{
			Code : "275000",
			Name : "INDUSTRI PERALATAN RUMAH TANGGA",
			SectorEconomy2Id : 101,
			Seq : "275000",
			
			
		},
		{
			Code : "279000",
			Name : "INDUSTRI PERALATAN LISTRIK LAINNYA",
			SectorEconomy2Id : 101,
			Seq : "279000",
			
			
		},
		{
			Code : "281000",
			Name : "INDUSTRI MESIN UNTUK KEPERLUAN UMUM",
			SectorEconomy2Id : 93,
			Seq : "281000",
			
			
		},
		{
			Code : "282100",
			Name : "INDUSTRI MESIN PERTANIAN DAN KEHUTANAN",
			SectorEconomy2Id : 93,
			Seq : "282100",
			
			
		},
		{
			Code : "282400",
			Name : "INDUSTRI MESIN PENAMBANGAN, PENGGALIAN DAN KONSTRUKSI",
			SectorEconomy2Id : 93,
			Seq : "282400",
			
			
		},
		{
			Code : "282500",
			Name : "INDUSTRI MESIN PENGOLAHAN MAKANAN, MINUMAN DAN TEMBAKAU",
			SectorEconomy2Id : 93,
			Seq : "282500",
			
			
		},
		{
			Code : "282600",
			Name : "INDUSTRI MESIN TEKSTIL, PAKAIAN JADI DAN PRODUK KULIT",
			SectorEconomy2Id : 93,
			Seq : "282600",
			
			
		},
		{
			Code : "282900",
			Name : "INDUSTRI MESIN KEPERLUAN KHUSUS LAINNYA",
			SectorEconomy2Id : 93,
			Seq : "282900",
			
			
		},
		{
			Code : "291000",
			Name : "INDUSTRI KENDARAAN BERMOTOR RODA EMPAT ATAU LEBIH",
			SectorEconomy2Id : 100,
			Seq : "291000",
			
			
		},
		{
			Code : "292000",
			Name : "INDUSTRI KAROSERI KENDARAAN BERMOTOR RODA EMPAT ATAU LEBIH DAN INDUSTRI TRAILER DAN SEMI TRAILER",
			SectorEconomy2Id : 100,
			Seq : "292000",
			
			
		},
		{
			Code : "293000",
			Name : "INDUSTRI SUKU CADANG DAN AKSESORI KENDARAAN BERMOTOR RODA EMPAT ATAU LEBIH",
			SectorEconomy2Id : 100,
			Seq : "293000",
			
			
		},
		{
			Code : "301000",
			Name : "INDUSTRI PEMBUATAN KAPAL DAN PERAHU",
			SectorEconomy2Id : 100,
			Seq : "301000",
			
			
		},
		{
			Code : "302000",
			Name : "INDUSTRI LOKOMOTIF DAN GERBONG KERETA",
			SectorEconomy2Id : 100,
			Seq : "302000",
			
			
		},
		{
			Code : "303000",
			Name : "INDUSTRI PESAWAT TERBANG DAN PERLENGKAPANNYA",
			SectorEconomy2Id : 100,
			Seq : "303000",
			
			
		},
		{
			Code : "309110",
			Name : "INDUSTRI SEPEDA MOTOR RODA DUA DAN TIGA",
			SectorEconomy2Id : 100,
			Seq : "309110",
			
			
		},
		{
			Code : "309900",
			Name : "INDUSTRI ALAT ANGKUTAN LAINNYA YTDL",
			SectorEconomy2Id : 100,
			Seq : "309900",
			
			
		},
		{
			Code : "310000",
			Name : "INDUSTRI FURNITUR",
			SectorEconomy2Id : 102,
			Seq : "310000",
			
			
		},
		{
			Code : "320000",
			Name : "INDUSTRI PENGOLAHAN LAINNYA",
			SectorEconomy2Id : 102,
			Seq : "320000",
			
			
		},
		{
			Code : "330000",
			Name : "REPARASI DAN PEMASANGAN MESIN DAN PERALATAN",
			SectorEconomy2Id : 93,
			Seq : "330000",
			
			
		},
		{
			Code : "351001",
			Name : "KETENAGALISTRIKAN PEDESAAN",
			SectorEconomy2Id : 2,
			Seq : "351001",
			
			
		},
		{
			Code : "351002",
			Name : "KETENAGALISTRIKAN LAINNYA",
			SectorEconomy2Id : 3,
			Seq : "351002",
			
			
		},
		{
			Code : "352000",
			Name : "PENGADAAN DAN DISTRIBUSI GAS ALAM DAN BUATAN",
			SectorEconomy2Id : 4,
			Seq : "352000",
			
			
		},
		{
			Code : "353000",
			Name : "PENGADAAN UAP/AIR PANAS, UDARA DINGIN DAN PRODUKSI ES",
			SectorEconomy2Id : 5,
			Seq : "353000",
			
			
		},
		{
			Code : "360000",
			Name : "PENGELOLAAN AIR",
			SectorEconomy2Id : 6,
			Seq : "360000",
			
			
		},
		{
			Code : "370000",
			Name : "PENGELOLAAN AIR LIMBAH",
			SectorEconomy2Id : 7,
			Seq : "370000",
			
			
		},
		{
			Code : "380000",
			Name : "PENGELOLAAN DAN DAUR ULANG SAMPAH",
			SectorEconomy2Id : 8,
			Seq : "380000",
			
			
		},
		{
			Code : "390000",
			Name : "AKTIVITAS REMEDIASI DAN PENGELOLAAN SAMPAH LAINNYA",
			SectorEconomy2Id : 9,
			Seq : "390000",
			
			
		},
		{
			Code : "410111",
			Name : "KONSTRUKSI PERUMAHAN SEDERHANA BANK TABUNGAN NEGARA",
			SectorEconomy2Id : 103,
			Seq : "410111",
			
			
		},
		{
			Code : "410112",
			Name : "KONSTRUKSI PERUMAHAN SEDERHANA PERUMNAS",
			SectorEconomy2Id : 103,
			Seq : "410112",
			
			
		},
		{
			Code : "410113",
			Name : "KONSTRUKSI PERUMAHAN SEDERHANA LAINNYA TIPE S.D. 21",
			SectorEconomy2Id : 103,
			Seq : "410113",
			
			
		},
		{
			Code : "410114",
			Name : "KONSTRUKSI PERUMAHAN SEDERHANA LAINNYA TIPE 22 S.D. 70",
			SectorEconomy2Id : 103,
			Seq : "410114",
			
			
		},
		{
			Code : "410115",
			Name : "KONSTRUKSI PERUMAHAN MENENGAH, BESAR, MEWAH (TIPE DIATAS 70)",
			SectorEconomy2Id : 103,
			Seq : "410115",
			
			
		},
		{
			Code : "410119",
			Name : "KONSTRUKSI GEDUNG TEMPAT TINGGAL LAINNYA",
			SectorEconomy2Id : 103,
			Seq : "410119",
			
			
		},
		{
			Code : "410120",
			Name : "KONSTRUKSI GEDUNG PERKANTORAN",
			SectorEconomy2Id : 103,
			Seq : "410120",
			
			
		},
		{
			Code : "410130",
			Name : "KONSTRUKSI GEDUNG INDUSTRI",
			SectorEconomy2Id : 103,
			Seq : "410130",
			
			
		},
		{
			Code : "410141",
			Name : "KONSTRUKSI GEDUNG PERBELANJAAN PASAR INPRES",
			SectorEconomy2Id : 103,
			Seq : "410141",
			
			
		},
		{
			Code : "410149",
			Name : "KONSTRUKSI GEDUNG PERBELANJAAN LAINNYA",
			SectorEconomy2Id : 103,
			Seq : "410149",
			
			
		},
		{
			Code : "410190",
			Name : "KONSTRUKSI GEDUNG LAINNYA",
			SectorEconomy2Id : 103,
			Seq : "410190",
			
			
		},
		{
			Code : "421101",
			Name : "KONSTRUKSI JALAN TOL",
			SectorEconomy2Id : 104,
			Seq : "421101",
			
			
		},
		{
			Code : "421102",
			Name : "KONSTRUKSI JALAN RAYA SELAIN TOL",
			SectorEconomy2Id : 104,
			Seq : "421102",
			
			
		},
		{
			Code : "421103",
			Name : "KONSTRUKSI JEMBATAN DAN JALAN LAYANG",
			SectorEconomy2Id : 104,
			Seq : "421103",
			
			
		},
		{
			Code : "421104",
			Name : "KONSTRUKSI JALAN REL DAN JEMBATAN REL",
			SectorEconomy2Id : 104,
			Seq : "421104",
			
			
		},
		{
			Code : "421109",
			Name : "KONSTRUKSI JALAN RAYA LAINNYA",
			SectorEconomy2Id : 104,
			Seq : "421109",
			
			
		},
		{
			Code : "422110",
			Name : "KONSTRUKSI JARINGAN IRIGASI",
			SectorEconomy2Id : 104,
			Seq : "422110",
			
			
		},
		{
			Code : "422131",
			Name : "KONSTRUKSI BANGUNAN LISTRIK PEDESAAN",
			SectorEconomy2Id : 104,
			Seq : "422131",
			
			
		},
		{
			Code : "422139",
			Name : "KONSTRUKSI BANGUNAN ELEKTRIKAL DAN KOMUNIKASI LAINNYA",
			SectorEconomy2Id : 104,
			Seq : "422139",
			
			
		},
		{
			Code : "422190",
			Name : "KONSTRUKSI JARINGAN ELEKTRIKAL DAN TELEKOMUNIKASI LAINNYA",
			SectorEconomy2Id : 103,
			Seq : "422190",
			
			
		},
		{
			Code : "429120",
			Name : "KONSTRUKSI BANGUNAN PELABUHAN BUKAN PERIKANAN",
			SectorEconomy2Id : 103,
			Seq : "429120",
			
			
		},
		{
			Code : "429190",
			Name : "KONSTRUKSI BANGUNAN SIPIL LAINNYA YTDL",
			SectorEconomy2Id : 103,
			Seq : "429190",
			
			
		},
		{
			Code : "431201",
			Name : "PENYIAPAN TANAH PEMUKIMAN TRANSMIGRASI (PTPT)",
			SectorEconomy2Id : 104,
			Seq : "431201",
			
			
		},
		{
			Code : "431202",
			Name : "PENCETAKAN LAHAN SAWAH",
			SectorEconomy2Id : 104,
			Seq : "431202",
			
			
		},
		{
			Code : "431209",
			Name : "PENYIAPAN LAHAN LAINNYA DAN PEMBONGKARAN",
			SectorEconomy2Id : 104,
			Seq : "431209",
			
			
		},
		{
			Code : "432000",
			Name : "INSTALASI SISTEM KELISTRIKAN, AIR (PIPA) DAN INSTALASI KONSTRUKSI LAINNYA",
			SectorEconomy2Id : 104,
			Seq : "432000",
			
			
		},
		{
			Code : "433000",
			Name : "PENYELESAIAN KONSTRUKSI BANGUNAN",
			SectorEconomy2Id : 103,
			Seq : "433000",
			
			
		},
		{
			Code : "439050",
			Name : "PENYEWAAN ALAT KONSTRUKSI DENGAN OPERATOR",
			SectorEconomy2Id : 104,
			Seq : "439050",
			
			
		},
		{
			Code : "439090",
			Name : "KONSTRUKSI KHUSUS LAINNYA YTDL",
			SectorEconomy2Id : 104,
			Seq : "439090",
			
			
		},
		{
			Code : "451000",
			Name : "PERDAGANGAN MOBIL",
			SectorEconomy2Id : 112,
			Seq : "451000",
			
			
		},
		{
			Code : "452000",
			Name : "REPARASI DAN PERAWATAN MOBIL",
			SectorEconomy2Id : 112,
			Seq : "452000",
			
			
		},
		{
			Code : "453000",
			Name : "PERDAGANGAN SUKU CADANG DAN AKSESORI MOBIL",
			SectorEconomy2Id : 112,
			Seq : "453000",
			
			
		},
		{
			Code : "454001",
			Name : "PERDAGANGAN SEPEDA MOTOR",
			SectorEconomy2Id : 112,
			Seq : "454001",
			
			
		},
		{
			Code : "454002",
			Name : "PERDAGANGAN SUKU CADANG SEPEDA MOTOR DAN AKSESORINYA",
			SectorEconomy2Id : 112,
			Seq : "454002",
			
			
		},
		{
			Code : "454003",
			Name : "REPARASI DAN PERAWATAN SEPEDA MOTOR",
			SectorEconomy2Id : 112,
			Seq : "454003",
			
			
		},
		{
			Code : "461000",
			Name : "PERDAGANGAN BESAR ATAS DASAR BALAS JASA (FEE) ATAU KONTRAK",
			SectorEconomy2Id : 108,
			Seq : "461000",
			
			
		},
		{
			Code : "462011",
			Name : "PERDAGANGAN BESAR JAGUNG",
			SectorEconomy2Id : 105,
			Seq : "462011",
			
			
		},
		{
			Code : "462019",
			Name : "PERDAGANGAN BESAR PADI DAN PALAWIJA LAINNYA",
			SectorEconomy2Id : 105,
			Seq : "462019",
			
			
		},
		{
			Code : "462020",
			Name : "PERDAGANGAN BESAR BUAH YANG MENGANDUNG MINYAK",
			SectorEconomy2Id : 105,
			Seq : "462020",
			
			
		},
		{
			Code : "462040",
			Name : "PERDAGANGAN BESAR TEMBAKAU RAJANGAN",
			SectorEconomy2Id : 106,
			Seq : "462040",
			
			
		},
		{
			Code : "462050",
			Name : "PERDAGANGAN BESAR BINATANG HIDUP",
			SectorEconomy2Id : 106,
			Seq : "462050",
			
			
		},
		{
			Code : "462060",
			Name : "PERDAGANGAN BESAR HASIL PERIKANAN",
			SectorEconomy2Id : 106,
			Seq : "462060",
			
			
		},
		{
			Code : "462071",
			Name : "PERDAGANGAN KAYU",
			SectorEconomy2Id : 112,
			Seq : "462071",
			
			
		},
		{
			Code : "462079",
			Name : "PERDAGANGAN BESAR HASIL KEHUTANAN DAN PERBURUAN LAINNYA",
			SectorEconomy2Id : 106,
			Seq : "462079",
			
			
		},
		{
			Code : "462080",
			Name : "PERDAGANGAN BESAR KULIT DAN KULIT JANGAT",
			SectorEconomy2Id : 106,
			Seq : "462080",
			
			
		},
		{
			Code : "462091",
			Name : "PERDAGANGAN KARET",
			SectorEconomy2Id : 112,
			Seq : "462091",
			
			
		},
		{
			Code : "462092",
			Name : "PERDAGANGAN CENGKEH",
			SectorEconomy2Id : 112,
			Seq : "462092",
			
			
		},
		{
			Code : "462093",
			Name : "PERDAGANGAN LADA",
			SectorEconomy2Id : 112,
			Seq : "462093",
			
			
		},
		{
			Code : "462094",
			Name : "PERDAGANGAN KAPAS",
			SectorEconomy2Id : 112,
			Seq : "462094",
			
			
		},
		{
			Code : "462095",
			Name : "PERDAGANGAN BIJI KELAPA SAWIT",
			SectorEconomy2Id : 112,
			Seq : "462095",
			
			
		},
		{
			Code : "462099",
			Name : "PERDAGANGAN BESAR HASIL PERTANIAN DAN HEWAN HIDUP LAINNYA",
			SectorEconomy2Id : 106,
			Seq : "462099",
			
			
		},
		{
			Code : "463110",
			Name : "PERDAGANGAN BESAR BERAS",
			SectorEconomy2Id : 105,
			Seq : "463110",
			
			
		},
		{
			Code : "463141",
			Name : "PERDAGANGAN BESAR KOPI",
			SectorEconomy2Id : 105,
			Seq : "463141",
			
			
		},
		{
			Code : "463142",
			Name : "PERDAGANGAN BESAR TEH",
			SectorEconomy2Id : 105,
			Seq : "463142",
			
			
		},
		{
			Code : "463150",
			Name : "PERDAGANGAN BESAR MINYAK DAN LEMAK NABATI",
			SectorEconomy2Id : 105,
			Seq : "463150",
			
			
		},
		{
			Code : "463190",
			Name : "PERDAGANGAN BESAR BAHAN MAKANAN DAN MINUMAN HASIL PERTANIAN LAINNYA",
			SectorEconomy2Id : 105,
			Seq : "463190",
			
			
		},
		{
			Code : "463201",
			Name : "PERDAGANGAN BESER UDANG OLAHAN",
			SectorEconomy2Id : 112,
			Seq : "463201",
			
			
		},
		{
			Code : "463209",
			Name : "PERDAGANGAN BESAR BAHAN MAKANAN DAN MINUMAN HASIL PETERNAKAN DAN PERIKANAN LAINNYA",
			SectorEconomy2Id : 105,
			Seq : "463209",
			
			
		},
		{
			Code : "463301",
			Name : "PERDAGANGAN BESAR GULA, COKLAT DAN KEMBANG GULA",
			SectorEconomy2Id : 105,
			Seq : "463301",
			
			
		},
		{
			Code : "463302",
			Name : "PERDAGANGAN BESAR ROKOK DAN TEMBAKAU",
			SectorEconomy2Id : 106,
			Seq : "463302",
			
			
		},
		{
			Code : "463309",
			Name : "PERDAGANGAN BESAR MAKANAN DAN MINUMAN LAINNYA",
			SectorEconomy2Id : 105,
			Seq : "463309",
			
			
		},
		{
			Code : "464110",
			Name : "PERDAGANGAN BESAR TEKSTIL",
			SectorEconomy2Id : 107,
			Seq : "464110",
			
			
		},
		{
			Code : "464120",
			Name : "PERDAGANGAN BESAR PAKAIAN",
			SectorEconomy2Id : 107,
			Seq : "464120",
			
			
		},
		{
			Code : "464130",
			Name : "PERDAGANGAN BESAR ALAS KAKI",
			SectorEconomy2Id : 107,
			Seq : "464130",
			
			
		},
		{
			Code : "464190",
			Name : "PERDAGANGAN BESAR TEKSTIL, PAKAIAN DAN ALAS KAKI LAINNYA",
			SectorEconomy2Id : 107,
			Seq : "464190",
			
			
		},
		{
			Code : "464900",
			Name : "PERDAGANGAN BESAR BARANG KEPERLUAN RUMAH TANGGA LAINNYA",
			SectorEconomy2Id : 108,
			Seq : "464900",
			
			
		},
		{
			Code : "465000",
			Name : "PERDAGANGAN BESAR MESIN, PERALATAN DAN PERLENGKAPANNYA",
			SectorEconomy2Id : 108,
			Seq : "465000",
			
			
		},
		{
			Code : "466100",
			Name : "PERDAGANGAN BESAR BAHAN BAKAR PADAT, CAIR DAN GAS DAN PRODUK YBDI",
			SectorEconomy2Id : 106,
			Seq : "466100",
			
			
		},
		{
			Code : "466200",
			Name : "PERDAGANGAN BESAR LOGAM DAN BIJIH LOGAM",
			SectorEconomy2Id : 106,
			Seq : "466200",
			
			
		},
		{
			Code : "466301",
			Name : "PERDAGANGAN BESAR BAHAN KONSTRUKSI DARI KAYU",
			SectorEconomy2Id : 108,
			Seq : "466301",
			
			
		},
		{
			Code : "466309",
			Name : "PERDAGANGAN BESAR BAHAN KONSTRUKSI LAINNYA",
			SectorEconomy2Id : 108,
			Seq : "466309",
			
			
		},
		{
			Code : "466920",
			Name : "PERDAGANGAN BESAR PUPUK DAN PRODUK AGROKIMIA",
			SectorEconomy2Id : 108,
			Seq : "466920",
			
			
		},
		{
			Code : "466930",
			Name : "PERDAGANGAN BESAR ALAT LABORATORIUM, FARMASI DAN KEDOKTERAN",
			SectorEconomy2Id : 108,
			Seq : "466930",
			
			
		},
		{
			Code : "466950",
			Name : "PERDAGANGAN BESAR KERTAS DAN KARTON",
			SectorEconomy2Id : 108,
			Seq : "466950",
			
			
		},
		{
			Code : "466970",
			Name : "PERDAGANGAN BESAR BARANG BEKAS DAN SISA-SISA TAK TERPAKAI (SCRAP)",
			SectorEconomy2Id : 108,
			Seq : "466970",
			
			
		},
		{
			Code : "466990",
			Name : "PERDAGANGAN BESAR PRODUK LAINNYA YTDL",
			SectorEconomy2Id : 108,
			Seq : "466990",
			
			
		},
		{
			Code : "471100",
			Name : "PERDAGANGAN ECERAN YANG UTAMANYA MAKANAN, MINUMAN ATAU TEMBAKAU DI TOKO",
			SectorEconomy2Id : 111,
			Seq : "471100",
			
			
		},
		{
			Code : "471900",
			Name : "PERDAGANGAN ECERAN BERBAGAI MACAM BARANG YANG DIDOMINASI OLEH BARANG BUKAN MAKANAN DAN TEMBAKAU DI TOKO",
			SectorEconomy2Id : 111,
			Seq : "471900",
			
			
		},
		{
			Code : "472001",
			Name : "PERDAGANGAN ECERAN KHUSUS KOMODITI MAKANAN DARI HASIL PERTANIAN DI TOKO",
			SectorEconomy2Id : 109,
			Seq : "472001",
			
			
		},
		{
			Code : "472009",
			Name : "PERDAGANGAN ECERAN KHUSUS MAKANAN, MINUMAN DAN TEMBAKAU LAINNYA DI TOKO",
			SectorEconomy2Id : 109,
			Seq : "472009",
			
			
		},
		{
			Code : "473000",
			Name : "PERDAGANGAN ECERAN KHUSUS BAHAN BAKAR KENDARAAN BERMOTOR",
			SectorEconomy2Id : 109,
			Seq : "473000",
			
			
		},
		{
			Code : "474000",
			Name : "PERDAGANGAN ECERAN KHUSUS PERALATAN INFORMASI DAN KOMUNIKASI DI TOKO",
			SectorEconomy2Id : 109,
			Seq : "474000",
			
			
		},
		{
			Code : "475100",
			Name : "PERDAGANGAN ECERAN KHUSUS TEKSTIL DI TOKO",
			SectorEconomy2Id : 109,
			Seq : "475100",
			
			
		},
		{
			Code : "475200",
			Name : "PERDAGANGAN ECERAN KHUSUS BARANG DAN BAHAN BANGUNAN, CAT DAN KACA DI TOKO",
			SectorEconomy2Id : 109,
			Seq : "475200",
			
			
		},
		{
			Code : "475900",
			Name : "PERDAGANGAN ECERAN KHUSUS FURNITUR, PERALATAN LISTRIK RUMAH TANGGA, PERALATAN PENERANGAN DAN PERALATAN RUMAH TANGGA LAINNYA DI TOKO",
			SectorEconomy2Id : 109,
			Seq : "475900",
			
			
		},
		{
			Code : "476000",
			Name : "PERDAGANGAN ECERAN KHUSUS BARANG BUDAYA DAN REKREASI DI TOKO KHUSUS",
			SectorEconomy2Id : 109,
			Seq : "476000",
			
			
		},
		{
			Code : "477100",
			Name : "PERDAGANGAN ECERAN KHUSUS PAKAIAN, ALAS KAKI DAN BARANG DARI KULIT DI TOKO",
			SectorEconomy2Id : 109,
			Seq : "477100",
			
			
		},
		{
			Code : "477200",
			Name : "PERDAGANGAN ECERAN KHUSUS BAHAN KIMIA, BARANG FARMASI, ALAT KEDOKTERAN, PARFUM DAN KOSMETIK DI TOKO",
			SectorEconomy2Id : 109,
			Seq : "477200",
			
			
		},
		{
			Code : "477300",
			Name : "PERDAGANGAN ECERAN KHUSUS BARANG BARU LAINNYA DI TOKO",
			SectorEconomy2Id : 109,
			Seq : "477300",
			
			
		},
		{
			Code : "477400",
			Name : "PERDAGANGAN ECERAN KHUSUS BARANG BEKAS DI TOKO",
			SectorEconomy2Id : 109,
			Seq : "477400",
			
			
		},
		{
			Code : "477700",
			Name : "PERDAGANGAN ECERAN BAHAN BAKAR BUKAN BAHAN BAKAR UNTUK KENDARAAN BERMOTOR DI TOKO",
			SectorEconomy2Id : 111,
			Seq : "477700",
			
			
		},
		{
			Code : "477800",
			Name : "PERDAGANGAN ECERAN BARANG KERAJINAN DAN LUKISAN DI TOKO",
			SectorEconomy2Id : 111,
			Seq : "477800",
			
			
		},
		{
			Code : "477900",
			Name : "PERDAGANGAN ECERAN KHUSUS BARANG LAINNYA YTDL",
			SectorEconomy2Id : 109,
			Seq : "477900",
			
			
		},
		{
			Code : "478100",
			Name : "PERDAGANGAN ECERAN KAKI LIMA DAN LOS PASAR KOMODITI HASIL PERTANIAN",
			SectorEconomy2Id : 110,
			Seq : "478100",
			
			
		},
		{
			Code : "478200",
			Name : "PERDAGANGAN ECERAN KAKI LIMA DAN LOS PASAR MAKANAN, MINUMAN DAN PRODUK TEMBAKAU HASIL INDUSTRI PENGOLAHAN",
			SectorEconomy2Id : 110,
			Seq : "478200",
			
			
		},
		{
			Code : "478300",
			Name : "PERDAGANGAN ECERAN KAKI LIMA DAN LOS PASAR TEKSTIL, PAKAIAN DAN ALAS KAKI",
			SectorEconomy2Id : 110,
			Seq : "478300",
			
			
		},
		{
			Code : "478400",
			Name : "PERDAGANGAN ECERAN KAKI LIMA DAN LOS PASAR BAHAN KIMIA, FARMASI, KOSMETIK DAN YBDI",
			SectorEconomy2Id : 110,
			Seq : "478400",
			
			
		},
		{
			Code : "478600",
			Name : "PERDAGANGAN ECERAN KAKI LIMA DAN LOS PASAR PERLENGKAPAN RUMAH TANGGA",
			SectorEconomy2Id : 110,
			Seq : "478600",
			
			
		},
		{
			Code : "478700",
			Name : "PERDAGANGAN ECERAN KAKI LIMA DAN LOS PASAR KERTAS, BARANG DARI KERTAS, ALAT TULIS, BARANG CETAKAN, ALAT OLAHRAGA, ALAT MUSIK, ALAT FOTOGRAFI DAN KOMPUTER",
			SectorEconomy2Id : 110,
			Seq : "478700",
			
			
		},
		{
			Code : "478800",
			Name : "PERDAGANGAN ECERAN KAKI LIMA DAN LOS PASAR BARANG KERAJINAN, MAINAN ANAK-ANAK DAN LUKISAN",
			SectorEconomy2Id : 110,
			Seq : "478800",
			
			
		},
		{
			Code : "478920",
			Name : "PERDAGANGAN ECERAN KAKI LIMA DAN LOS PASAR BAHAN BAKAR MINYAK, GAS, MINYAK PELUMAS DAN BAHAN BAKAR LAINNYA",
			SectorEconomy2Id : 110,
			Seq : "478920",
			
			
		},
		{
			Code : "478940",
			Name : "PERDAGANGAN ECERAN KAKI LIMA DAN LOS PASAR BARANG BEKAS PERLENGKAPAN RUMAH TANGGA",
			SectorEconomy2Id : 110,
			Seq : "478940",
			
			
		},
		{
			Code : "478990",
			Name : "PERDAGANGAN ECERAN KAKI LIMA DAN LOS PASAR BARANG LAINNYA",
			SectorEconomy2Id : 110,
			Seq : "478990",
			
			
		},
		{
			Code : "479100",
			Name : "PERDAGANGAN ECERAN MELALUI PEMESANAN POS ATAU INTERNET",
			SectorEconomy2Id : 111,
			Seq : "479100",
			
			
		},
		{
			Code : "479900",
			Name : "PERDAGANGAN ECERAN BUKAN DI TOKO, KIOS, KAKI LIMA DAN LOS PASAR LAINNYA",
			SectorEconomy2Id : 110,
			Seq : "479900",
			
			
		},
		{
			Code : "491000",
			Name : "ANGKUTAN JALAN REL",
			SectorEconomy2Id : 113,
			Seq : "491000",
			
			
		},
		{
			Code : "492100",
			Name : "ANGKUTAN BUS BERTRAYEK",
			SectorEconomy2Id : 113,
			Seq : "492100",
			
			
		},
		{
			Code : "492210",
			Name : "ANGKUTAN BUS PARIWISATA",
			SectorEconomy2Id : 113,
			Seq : "492210",
			
			
		},
		{
			Code : "492290",
			Name : "ANGKUTAN BUS TIDAK BERTRAYEK LAINNYA",
			SectorEconomy2Id : 113,
			Seq : "492290",
			
			
		},
		{
			Code : "493000",
			Name : "ANGKUTAN MELALUI SALURAN PIPA",
			SectorEconomy2Id : 116,
			Seq : "493000",
			
			
		},
		{
			Code : "494100",
			Name : "ANGKUTAN DARAT BUKAN BUS UNTUK PENUMPANG, BERTRAYEK",
			SectorEconomy2Id : 113,
			Seq : "494100",
			
			
		},
		{
			Code : "494200",
			Name : "ANGKUTAN DARAT LAINNYA UNTUK PENUMPANG",
			SectorEconomy2Id : 113,
			Seq : "494200",
			
			
		},
		{
			Code : "494300",
			Name : "ANGKUTAN DARAT UNTUK BARANG",
			SectorEconomy2Id : 113,
			Seq : "494300",
			
			
		},
		{
			Code : "494501",
			Name : "ANGKUTAN JALAN REL WISATA",
			SectorEconomy2Id : 113,
			Seq : "494501",
			
			
		},
		{
			Code : "494509",
			Name : "ANGKUTAN JALAN REL LAINNYA",
			SectorEconomy2Id : 113,
			Seq : "494509",
			
			
		},
		{
			Code : "501100",
			Name : "ANGKUTAN LAUT DALAM NEGERI UNTUK PENUMPANG",
			SectorEconomy2Id : 114,
			Seq : "501100",
			
			
		},
		{
			Code : "501130",
			Name : "ANGKUTAN LAUT UNTUK WISATA",
			SectorEconomy2Id : 114,
			Seq : "501130",
			
			
		},
		{
			Code : "501190",
			Name : "ANGKUTAN LAUT DALAM NEGERI UNTUK PENUMPANG SELAIN WISATA",
			SectorEconomy2Id : 114,
			Seq : "501190",
			
			
		},
		{
			Code : "501200",
			Name : "ANGKUTAN LAUT LUAR NEGERI UNTUK PENUMPANG",
			SectorEconomy2Id : 114,
			Seq : "501200",
			
			
		},
		{
			Code : "501300",
			Name : "ANGKUTAN LAUT DALAM NEGERI UNTUK BARANG",
			SectorEconomy2Id : 114,
			Seq : "501300",
			
			
		},
		{
			Code : "501400",
			Name : "ANGKUTAN LAUT LUAR NEGERI UNTUK BARANG",
			SectorEconomy2Id : 114,
			Seq : "501400",
			
			
		},
		{
			Code : "502101",
			Name : "ANGKUTAN SUNGAI DAN DANAU UNTUK WISATA DAN YBDI",
			SectorEconomy2Id : 114,
			Seq : "502101",
			
			
		},
		{
			Code : "502102",
			Name : "ANGKUTAN PENYEBERANGAN UNTUK PENUMPANG",
			SectorEconomy2Id : 114,
			Seq : "502102",
			
			
		},
		{
			Code : "502200",
			Name : "ANGKUTAN SUNGAI, DANAU DAN PENYEBERANGAN UNTUK BARANG",
			SectorEconomy2Id : 114,
			Seq : "502200",
			
			
		},
		{
			Code : "511001",
			Name : "ANGKUTAN UDARA BERJADWAL UNTUK PENUMPANG",
			SectorEconomy2Id : 115,
			Seq : "511001",
			
			
		},
		{
			Code : "511002",
			Name : "ANGKUTAN UDARA TIDAK BERJADWAL UNTUK PENUMPANG",
			SectorEconomy2Id : 115,
			Seq : "511002",
			
			
		},
		{
			Code : "511009",
			Name : "ANGKUTAN UDARA UNTUK PENUMPANG LAINNYA",
			SectorEconomy2Id : 115,
			Seq : "511009",
			
			
		},
		{
			Code : "512000",
			Name : "ANGKUTAN UDARA UNTUK BARANG",
			SectorEconomy2Id : 115,
			Seq : "512000",
			
			
		},
		{
			Code : "521000",
			Name : "PERGUDANGAN DAN PENYIMPANAN",
			SectorEconomy2Id : 116,
			Seq : "521000",
			
			
		},
		{
			Code : "522000",
			Name : "AKTIVITAS PENUNJANG ANGKUTAN",
			SectorEconomy2Id : 116,
			Seq : "522000",
			
			
		},
		{
			Code : "530000",
			Name : "AKTIVITAS POS DAN KURIR",
			SectorEconomy2Id : 116,
			Seq : "530000",
			
			
		},
		{
			Code : "551100",
			Name : "HOTEL BINTANG",
			SectorEconomy2Id : 10,
			Seq : "551100",
			
			
		},
		{
			Code : "551200",
			Name : "HOTEL MELATI",
			SectorEconomy2Id : 11,
			Seq : "551200",
			
			
		},
		{
			Code : "559000",
			Name : "PENYEDIAAN AKOMODASI LAINNYA",
			SectorEconomy2Id : 12,
			Seq : "559000",
			
			
		},
		{
			Code : "561001",
			Name : "RESTORAN DAN RUMAH MAKAN",
			SectorEconomy2Id : 13,
			Seq : "561001",
			
			
		},
		{
			Code : "561009",
			Name : "PENYEDIAAN MAKANAN DAN MINUMAN LAINNYA",
			SectorEconomy2Id : 14,
			Seq : "561009",
			
			
		},
		{
			Code : "580000",
			Name : "AKTIVITAS PENERBITAN",
			SectorEconomy2Id : 15,
			Seq : "580000",
			
			
		},
		{
			Code : "591000",
			Name : "AKTIVITAS PRODUKSI GAMBAR BERGERAK, VIDEO DAN PROGRAM TELEVISI",
			SectorEconomy2Id : 16,
			Seq : "591000",
			
			
		},
		{
			Code : "592000",
			Name : "AKTIVITAS PEREKAMAN SUARA DAN PENERBITAN MUSIK",
			SectorEconomy2Id : 17,
			Seq : "592000",
			
			
		},
		{
			Code : "600000",
			Name : "AKTIVITAS PENYIARAN DAN PEMROGRAMAN",
			SectorEconomy2Id : 18,
			Seq : "600000",
			
			
		},
		{
			Code : "610001",
			Name : "AKTIVITAS TELEKOMUNIKASI DENGAN KABEL, TANPA KABEL DAN SATELIT",
			SectorEconomy2Id : 19,
			Seq : "610001",
			
			
		},
		{
			Code : "610002",
			Name : "JASA NILAI TAMBAH TELEPONI DAN JASA MULTIMEDIA",
			SectorEconomy2Id : 20,
			Seq : "610002",
			
			
		},
		{
			Code : "610009",
			Name : "AKTIVITAS TELEKOMUNIKASI LAINNYA YTDL",
			SectorEconomy2Id : 21,
			Seq : "610009",
			
			
		},
		{
			Code : "620100",
			Name : "AKTIVITAS PEMROGRAMAN KOMPUTER",
			SectorEconomy2Id : 22,
			Seq : "620100",
			
			
		},
		{
			Code : "620200",
			Name : "AKTIVITAS KONSULTASI KOMPUTER DAN MANAJEMEN FASILITAS KOMPUTER",
			SectorEconomy2Id : 23,
			Seq : "620200",
			
			
		},
		{
			Code : "631110",
			Name : "AKTIVITAS PENGOLAHAN DATA",
			SectorEconomy2Id : 24,
			Seq : "631110",
			
			
		},
		{
			Code : "631120",
			Name : "AKTIVITAS HOSTING DAN YBDI",
			SectorEconomy2Id : 25,
			Seq : "631120",
			
			
		},
		{
			Code : "631210",
			Name : "PORTAL WEB DAN/ATAU PLATFORM DIGITAL TANPA TUJUAN KOMERSIAL",
			SectorEconomy2Id : 26,
			Seq : "631210",
			
			
		},
		{
			Code : "631220",
			Name : "PORTAL WEB DAN/ATAU PLATFORM DIGITAL DENGAN TUJUAN KOMERSIAL",
			SectorEconomy2Id : 27,
			Seq : "631220",
			
			
		},
		{
			Code : "639100",
			Name : "AKTIVITAS KANTOR BERITA",
			SectorEconomy2Id : 28,
			Seq : "639100",
			
			
		},
		{
			Code : "639900",
			Name : "AKTIVITAS JASA INFORMASI LAINNYA YTDL",
			SectorEconomy2Id : 29,
			Seq : "639900",
			
			
		},
		{
			Code : "641000",
			Name : "PERANTARA MONETER",
			SectorEconomy2Id : 30,
			Seq : "641000",
			
			
		},
		{
			Code : "649100",
			Name : "SEWA GUNA USAHA DENGAN HAK OPSI",
			SectorEconomy2Id : 31,
			Seq : "649100",
			
			
		},
		{
			Code : "649900",
			Name : "AKTIVITAS JASA KEUANGAN LAINNYA YTDL, BUKAN ASURANSI DAN DANA PENSIUN",
			SectorEconomy2Id : 32,
			Seq : "649900",
			
			
		},
		{
			Code : "650000",
			Name : "ASURANSI, REASURANSI DAN DANA PENSIUN, BUKAN JAMINAN SOSIAL WAJIB",
			SectorEconomy2Id : 33,
			Seq : "650000",
			
			
		},
		{
			Code : "661001",
			Name : "KEGIATAN PENUKARAN VALUTA ASING (MONEY CHANGER)",
			SectorEconomy2Id : 34,
			Seq : "661001",
			
			
		},
		{
			Code : "661009",
			Name : "AKTIVITAS PENUNJANG JASA KEUANGAN LAINNYA",
			SectorEconomy2Id : 35,
			Seq : "661009",
			
			
		},
		{
			Code : "662000",
			Name : "AKTIVITAS PENUNJANG ASURANSI DAN DANA PENSIUN",
			SectorEconomy2Id : 36,
			Seq : "662000",
			
			
		},
		{
			Code : "681101",
			Name : "REAL ESTATE PERUMAHAN SEDERHANA PERUMNAS",
			SectorEconomy2Id : 37,
			Seq : "681101",
			
			
		},
		{
			Code : "681102",
			Name : "REAL ESTATE PERUMAHAN SEDERHANA PERUMNAS TIPE 21",
			SectorEconomy2Id : 38,
			Seq : "681102",
			
			
		},
		{
			Code : "681103",
			Name : "REAL ESTATE PERUMAHAN SEDERHANA PERUMNAS TIPE 22 S.D. 70",
			SectorEconomy2Id : 39,
			Seq : "681103",
			
			
		},
		{
			Code : "681104",
			Name : "REAL ESTATE PERUMAHAN MENENGAH, BESAR ATAU MEWAH (TIPE DIATAS 70)",
			SectorEconomy2Id : 40,
			Seq : "681104",
			
			
		},
		{
			Code : "681105",
			Name : "REAL ESTATE PERUMAHAN FLAT / APARTEMEN",
			SectorEconomy2Id : 41,
			Seq : "681105",
			
			
		},
		{
			Code : "681106",
			Name : "REAL ESTATE GEDUNG PERBELANJAAN (MAL, PLAZA)",
			SectorEconomy2Id : 42,
			Seq : "681106",
			
			
		},
		{
			Code : "681107",
			Name : "REAL ESTATE GEDUNG PERKANTORAN",
			SectorEconomy2Id : 43,
			Seq : "681107",
			
			
		},
		{
			Code : "681108",
			Name : "REAL ESTATE GEDUNG RUMAH TOKO (RUKO) ATAU RUMAH KANTOR (RUKAN)",
			SectorEconomy2Id : 44,
			Seq : "681108",
			
			
		},
		{
			Code : "681109",
			Name : "REAL ESTATE LAINNYA",
			SectorEconomy2Id : 45,
			Seq : "681109",
			
			
		},
		{
			Code : "681200",
			Name : "KAWASAN PARIWISATA",
			SectorEconomy2Id : 46,
			Seq : "681200",
			
			
		},
		{
			Code : "681300",
			Name : "KAWASAN INDUSTRI",
			SectorEconomy2Id : 47,
			Seq : "681300",
			
			
		},
		{
			Code : "682000",
			Name : "REAL ESTAT ATAS DASAR BALAS JASA (FEE) ATAU KONTRAK",
			SectorEconomy2Id : 48,
			Seq : "682000",
			
			
		},
		{
			Code : "690000",
			Name : "AKTIVITAS HUKUM DAN AKUNTANSI",
			SectorEconomy2Id : 49,
			Seq : "690000",
			
			
		},
		{
			Code : "702010",
			Name : "AKTIVITAS KONSULTASI PARIWISATA",
			SectorEconomy2Id : 50,
			Seq : "702010",
			
			
		},
		{
			Code : "702090",
			Name : "AKTIVITAS KANTOR PUSAT DAN KONSULTASI MANAJEMEN LAINNYA",
			SectorEconomy2Id : 51,
			Seq : "702090",
			
			
		},
		{
			Code : "710000",
			Name : "AKTIVITAS ARSITEKTUR DAN KEINSINYURAN; ANALISIS DAN UJI TEKNIS",
			SectorEconomy2Id : 52,
			Seq : "710000",
			
			
		},
		{
			Code : "721000",
			Name : "PENELITIAN DAN PENGEMBANGAN ILMU PENGETAHUAN ALAM DAN ILMU TEKNOLOGI DAN REKAYASA",
			SectorEconomy2Id : 53,
			Seq : "721000",
			
			
		},
		{
			Code : "722000",
			Name : "PENELITIAN DAN PENGEMBANGAN ILMU PENGETAHUAN SOSIAL DAN HUMANIORA",
			SectorEconomy2Id : 54,
			Seq : "722000",
			
			
		},
		{
			Code : "730000",
			Name : "PERIKLANAN DAN PENELITIAN PASAR",
			SectorEconomy2Id : 55,
			Seq : "730000",
			
			
		},
		{
			Code : "740000",
			Name : "AKTIVITAS PROFESIONAL, ILMIAH DAN TEKNIS LAINNYA",
			SectorEconomy2Id : 56,
			Seq : "740000",
			
			
		},
		{
			Code : "750000",
			Name : "AKTIVITAS KESEHATAN HEWAN",
			SectorEconomy2Id : 57,
			Seq : "750000",
			
			
		},
		{
			Code : "771000",
			Name : "AKTIVITAS PENYEWAAN DAN SEWA GUNA USAHA TANPA HAK OPSI MOBIL, BUS, TRUK DAN SEJENISNYA",
			SectorEconomy2Id : 117,
			Seq : "771000",
			
			
		},
		{
			Code : "772000",
			Name : "AKTIVITAS PENYEWAAN DAN SEWA GUNA USAHA TANPA HAK OPSI BARANG PRIBADI DAN RUMAH TANGGA",
			SectorEconomy2Id : 117,
			Seq : "772000",
			
			
		},
		{
			Code : "773020",
			Name : "AKTIVITAS PENYEWAAN DAN SEWA GUNA USAHA TANPA HAK OPSI ALAT TRANSPORTASI DARAT BUKAN KENDARAAN BERMOTOR RODA EMPAT ATAU LEBIH",
			SectorEconomy2Id : 117,
			Seq : "773020",
			
			
		},
		{
			Code : "773030",
			Name : "AKTIVITAS PENYEWAAN DAN SEWA GUNA USAHA TANPA HAK OPSI ALAT TRANSPORTASI AIR",
			SectorEconomy2Id : 117,
			Seq : "773030",
			
			
		},
		{
			Code : "773040",
			Name : "AKTIVITAS PENYEWAAN DAN SEWA GUNA USAHA TANPA HAK OPSI ALAT TRANSPORTASI UDARA",
			SectorEconomy2Id : 117,
			Seq : "773040",
			
			
		},
		{
			Code : "773050",
			Name : "AKTIVITAS PENYEWAAN DAN SEWA GUNA USAHA TANPA HAK OPSI MESIN PERTANIAN DAN PERALATANNYA",
			SectorEconomy2Id : 117,
			Seq : "773050",
			
			
		},
		{
			Code : "773060",
			Name : "AKTIVITAS PENYEWAAN DAN SEWA GUNA USAHA TANPA HAK OPSI MESIN DAN PERALATAN KONSTRUKSI DAN TEKNIK SIPIL",
			SectorEconomy2Id : 117,
			Seq : "773060",
			
			
		},
		{
			Code : "773070",
			Name : "AKTIVITAS PENYEWAAN DAN SEWA GUNA USAHA TANPA HAK OPSI MESIN KANTOR DAN PERALATANNYA",
			SectorEconomy2Id : 117,
			Seq : "773070",
			
			
		},
		{
			Code : "773090",
			Name : "AKTIVITAS PENYEWAAN DAN SEWA GUNA USAHA TANPA HAK OPSI MESIN, PERALATAN DAN BARANG BERWUJUD LAINNYA YTDL",
			SectorEconomy2Id : 117,
			Seq : "773090",
			
			
		},
		{
			Code : "780000",
			Name : "AKTIVITAS KETENAGAKERJAAN",
			SectorEconomy2Id : 118,
			Seq : "780000",
			
			
		},
		{
			Code : "791110",
			Name : "AKTIVITAS AGEN PERJALANAN WISATA",
			SectorEconomy2Id : 118,
			Seq : "791110",
			
			
		},
		{
			Code : "791120",
			Name : "AKTIVITAS AGEN PERJALANAN BUKAN WISATA",
			SectorEconomy2Id : 118,
			Seq : "791120",
			
			
		},
		{
			Code : "791200",
			Name : "AKTIVITAS BIRO PERJALANAN WISATA",
			SectorEconomy2Id : 118,
			Seq : "791200",
			
			
		},
		{
			Code : "799000",
			Name : "JASA RESERVASI LAINNYA DAN KEGIATAN YBDI",
			SectorEconomy2Id : 118,
			Seq : "799000",
			
			
		},
		{
			Code : "823000",
			Name : "PENYELENGGARA KONVENSI DAN PAMERAN DAGANG",
			SectorEconomy2Id : 118,
			Seq : "823000",
			
			
		},
		{
			Code : "829000",
			Name : "AKTIVITAS JASA PENUNJANG USAHA YTDL",
			SectorEconomy2Id : 118,
			Seq : "829000",
			
			
		},
		{
			Code : "841000",
			Name : "ADMINISTRASI PEMERINTAHAN DAN KEBIJAKAN EKONOMI DAN SOSIAL",
			SectorEconomy2Id : 58,
			Seq : "841000",
			
			
		},
		{
			Code : "842000",
			Name : "PENYEDIAAN LAYANAN UNTUK MASYARAKAT DALAM BIDANG HUBUNGAN LUAR NEGERI, PERTAHANAN, KEAMANAN DAN KETERTIBAN",
			SectorEconomy2Id : 59,
			Seq : "842000",
			
			
		},
		{
			Code : "843000",
			Name : "JAMINAN SOSIAL WAJIB",
			SectorEconomy2Id : 60,
			Seq : "843000",
			
			
		},
		{
			Code : "851000",
			Name : "PENDIDIKAN DASAR DAN PENDIDIKAN ANAK USIA DINI",
			SectorEconomy2Id : 61,
			Seq : "851000",
			
			
		},
		{
			Code : "852000",
			Name : "PENDIDIKAN MENENGAH",
			SectorEconomy2Id : 62,
			Seq : "852000",
			
			
		},
		{
			Code : "853000",
			Name : "PENDIDIKAN TINGGI",
			SectorEconomy2Id : 63,
			Seq : "853000",
			
			
		},
		{
			Code : "854000",
			Name : "PENDIDIKAN LAINNYA",
			SectorEconomy2Id : 64,
			Seq : "854000",
			
			
		},
		{
			Code : "855000",
			Name : "KEGIATAN PENUNJANG PENDIDIKAN",
			SectorEconomy2Id : 65,
			Seq : "855000",
			
			
		},
		{
			Code : "861000",
			Name : "AKTIVITAS RUMAH SAKIT",
			SectorEconomy2Id : 66,
			Seq : "861000",
			
			
		},
		{
			Code : "862000",
			Name : "AKTIVITAS PRAKTIK DOKTER DAN DOKTER GIGI",
			SectorEconomy2Id : 67,
			Seq : "862000",
			
			
		},
		{
			Code : "869000",
			Name : "AKTIVITAS PELAYANAN KESEHATAN MANUSIA LAINNYA",
			SectorEconomy2Id : 68,
			Seq : "869000",
			
			
		},
		{
			Code : "870000",
			Name : "AKTIVITAS SOSIAL",
			SectorEconomy2Id : 69,
			Seq : "870000",
			
			
		},
		{
			Code : "900001",
			Name : "JASA IMPRESARIAT BIDANG SENI",
			SectorEconomy2Id : 70,
			Seq : "900001",
			
			
		},
		{
			Code : "900009",
			Name : "AKTIVITAS HIBURAN, SENI DAN KREATIVITAS LAINNYA",
			SectorEconomy2Id : 71,
			Seq : "900009",
			
			
		},
		{
			Code : "910100",
			Name : "PERPUSTAKAAN DAN ARSIP",
			SectorEconomy2Id : 72,
			Seq : "910100",
			
			
		},
		{
			Code : "910200",
			Name : "MUSEUM DAN OPERASIONAL BANGUNAN DAN SITUS BERSEJARAH",
			SectorEconomy2Id : 73,
			Seq : "910200",
			
			
		},
		{
			Code : "930000",
			Name : "AKTIVITAS OLAHRAGA DAN REKREASI LAINNYA",
			SectorEconomy2Id : 74,
			Seq : "930000",
			
			
		},
		{
			Code : "941000",
			Name : "AKTIVITAS ORGANISASI BISNIS, PENGUSAHA DAN PROFESI",
			SectorEconomy2Id : 75,
			Seq : "941000",
			
			
		},
		{
			Code : "942000",
			Name : "AKTIVITAS ORGANISASI BURUH",
			SectorEconomy2Id : 76,
			Seq : "942000",
			
			
		},
		{
			Code : "949000",
			Name : "AKTIVITAS ORGANISASI KEANGGOTAAN LAINNYA YTDL",
			SectorEconomy2Id : 77,
			Seq : "949000",
			
			
		},
		{
			Code : "950000",
			Name : "REPARASI KOMPUTER DAN BARANG KEPERLUAN PRIBADI DAN PERLENGKAPAN RUMAH TANGGA",
			SectorEconomy2Id : 78,
			Seq : "950000",
			
			
		},
		{
			Code : "960001",
			Name : "AKTIVITAS PANTI PIJAT DAN SPA",
			SectorEconomy2Id : 79,
			Seq : "960001",
			
			
		},
		{
			Code : "960009",
			Name : "AKTIVITAS JASA PERORANGAN LAINNYA",
			SectorEconomy2Id : 80,
			Seq : "960009",
			
			
		},
		{
			Code : "970000",
			Name : "AKTIVITAS RUMAH TANGGA SEBAGAI PEMBERI KERJA DARI PERSONIL DOMESTIK",
			SectorEconomy2Id : 81,
			Seq : "970000",
			
			
		},
		{
			Code : "990000",
			Name : "AKTIVITAS BADAN INTERNASIONAL DAN BADAN EKSTRA INTERNASIONAL LAINNYA",
			SectorEconomy2Id : 82,
			Seq : "990000",
			
			
		},
	}

	var sectorEconomyOjk = []models.SectorEconomyOjk{
		{
			SectorEconomy3Id : 1,
			Code : "001110",
			Name : "RUMAH TANGGA UNTUK PEMILIKAN RUMAH TINGGAL S.D. TIPE 21",
			
		},
		{
			SectorEconomy3Id : 2,
			Code : "001120",
			Name : "RUMAH TANGGA UNTUK PEMILIKAN RUMAH TINGGAL TIPE DIATAS 21 S.D. 70",
			
		},
		{
			SectorEconomy3Id : 3,
			Code : "001130",
			Name : "RUMAH TANGGA UNTUK PEMILIKAN RUMAH TINGGAL TIPE DIATAS 70",
			
		},
		{
			SectorEconomy3Id : 4,
			Code : "001210",
			Name : "RUMAH TANGGA UNTUK PEMILIKAN FLAT ATAU APARTEMEN S.D. TIPE 21",
			
		},
		{
			SectorEconomy3Id : 5,
			Code : "001220",
			Name : "RUMAH TANGGA UNTUK PEMILIKAN FLAT ATAU APARTEMEN TIPE DIATAS 21 S.D. 70",
			
		},
		{
			SectorEconomy3Id : 6,
			Code : "001230",
			Name : "RUMAH TANGGA UNTUK PEMILIKAN FLAT ATAU APARTEMEN TIPE DIATAS 70",
			
		},
		{
			SectorEconomy3Id : 7,
			Code : "001300",
			Name : "RUMAH TANGGA UNTUK PEMILIKAN RUMAH TOKO (RUKO) ATAU RUMAH KANTOR (RUKAN)",
			
		},
		{
			SectorEconomy3Id : 8,
			Code : "002100",
			Name : "RUMAH TANGGA UNTUK PEMILIKAN MOBIL RODA EMPAT",
			
		},
		{
			SectorEconomy3Id : 9,
			Code : "002200",
			Name : "RUMAH TANGGA UNTUK PEMILIKAN SEPEDA BERMOTOR",
			
		},
		{
			SectorEconomy3Id : 10,
			Code : "002300",
			Name : "RUMAH TANGGA UNTUK PEMILIKAN TRUK DAN KENDARAAN BERMOTOR RODA ENAM ATAU LEBIH",
			
		},
		{
			SectorEconomy3Id : 11,
			Code : "002900",
			Name : "RUMAH TANGGA UNTUK PEMILIKAN KENDARAAN BERMOTOR LAINNYA",
			
		},
		{
			SectorEconomy3Id : 12,
			Code : "003100",
			Name : "RUMAH TANGGA UNTUK PEMILIKAN FURNITUR DAN PERALATAN RUMAH TANGGA",
			
		},
		{
			SectorEconomy3Id : 13,
			Code : "003200",
			Name : "RUMAH TANGGA UNTUK PEMILIKAN TELEVISI, RADIO, DAN ALAT ELEKTRONIK",
			
		},
		{
			SectorEconomy3Id : 14,
			Code : "003300",
			Name : "RUMAH TANGGA UNTUK PEMILIKAN KOMPUTER DAN ALAT KOMUNIKASI",
			
		},
		{
			SectorEconomy3Id : 15,
			Code : "003900",
			Name : "RUMAH TANGGA UNTUK PEMILIKAN PERALATAN LAINNYA",
			
		},
		{
			SectorEconomy3Id : 16,
			Code : "004120",
			Name : "RUMAH TANGGA UNTUK KEPERLUAN MULTIGUNA BERAGUNAN RUMAH TINGGAL S.D TIPE 21",
			
		},
		{
			SectorEconomy3Id : 17,
			Code : "004130",
			Name : "RUMAH TANGGA UNTUK KEPERLUAN MULTIGUNA BERAGUNAN RUMAH TINGGAL TIPE DIATAS 21 S.D 70",
			
		},
		{
			SectorEconomy3Id : 18,
			Code : "004140",
			Name : "RUMAH TANGGA UNTUK KEPERLUAN MULTIGUNA BERAGUNAN RUMAH TINGGAL TIPE DIATAS 70",
			
		},
		{
			SectorEconomy3Id : 19,
			Code : "004150",
			Name : "RUMAH TANGGA UNTUK KEPERLUAN MULTIGUNA BERAGUNAN APARTEMEN S.D TIPE 21",
			
		},
		{
			SectorEconomy3Id : 20,
			Code : "004160",
			Name : "RUMAH TANGGA UNTUK KEPERLUAN MULTIGUNA BERAGUNAN APARTEMEN TIPE 22 S.D 70",
			
		},
		{
			SectorEconomy3Id : 21,
			Code : "004170",
			Name : "RUMAH TANGGA UNTUK KEPERLUAN MULTIGUNA BERAGUNAN APARTEMEN TIPE DIATAS 70",
			
		},
		{
			SectorEconomy3Id : 22,
			Code : "004180",
			Name : "RUMAH TANGGA UNTUK KEPERLUAN MULTIGUNA BERAGUNAN RUKO/RUKAN",
			
		},
		{
			SectorEconomy3Id : 23,
			Code : "004190",
			Name : "RUMAH TANGGA UNTUK KEPERLUAN MULTIGUNA LAINNYA",
			
		},
		{
			SectorEconomy3Id : 24,
			Code : "004900",
			Name : "RUMAH TANGGA UNTUK KEPERLUAN YANG TIDAK DIKLASIFIKASIKAN DI TEMPAT LAIN",
			
		},
		{
			SectorEconomy3Id : 25,
			Code : "009000",
			Name : "BUKAN LAPANGAN USAHA LAINNYA",
			
		},
		{
			SectorEconomy3Id : 26,
			Code : "011110",
			Name : "PERTANIAN JAGUNG",
			
		},
		{
			SectorEconomy3Id : 27,
			Code : "011130",
			Name : "PERTANIAN KEDELAI",
			
		},
		{
			SectorEconomy3Id : 28,
			Code : "011140",
			Name : "PERTANIAN KACANG TANAH",
			
		},
		{
			SectorEconomy3Id : 29,
			Code : "011190",
			Name : "PERTANIAN SEREALIA LAINNYA, ANEKA KACANG DAN BIJI-BIJIAN PENGHASIL MINYAK LAINNYA",
			
		},
		{
			SectorEconomy3Id : 30,
			Code : "011200",
			Name : "PERTANIAN PADI",
			
		},
		{
			SectorEconomy3Id : 31,
			Code : "011301",
			Name : "PERTANIAN HORTIKULTURA BAWANG MERAH",
			
		},
		{
			SectorEconomy3Id : 32,
			Code : "011302",
			Name : "PERTANIAN ANEKA UMBI PALAWIJA",
			
		},
		{
			SectorEconomy3Id : 33,
			Code : "011303",
			Name : "PERTANIAN BIT GULA DAN TANAMAN PEMANIS BUKAN TEBU",
			
		},
		{
			SectorEconomy3Id : 34,
			Code : "011309",
			Name : "PERTANIAN SAYURAN, BUAH DAN ANEKA UMBI LAINNYA",
			
		},
		{
			SectorEconomy3Id : 35,
			Code : "011400",
			Name : "PERKEBUNAN TEBU",
			
		},
		{
			SectorEconomy3Id : 36,
			Code : "011500",
			Name : "PERKEBUNAN TEMBAKAU",
			
		},
		{
			SectorEconomy3Id : 37,
			Code : "011600",
			Name : "PERTANIAN TANAMAN BERSERAT",
			
		},
		{
			SectorEconomy3Id : 38,
			Code : "011909",
			Name : "PERTANIAN TANAMAN SEMUSIM LAINNYA YTDL",
			
		},
		{
			SectorEconomy3Id : 39,
			Code : "011930",
			Name : "PERTANIAN TANAMAN BUNGA",
			
		},
		{
			SectorEconomy3Id : 40,
			Code : "011940",
			Name : "PERTANIAN PEMBIBITAN TANAMAN BUNGA",
			
		},
		{
			SectorEconomy3Id : 41,
			Code : "012201",
			Name : "PERTANIAN BUAH PISANG",
			
		},
		{
			SectorEconomy3Id : 42,
			Code : "012209",
			Name : "PERTANIAN BUAH-BUAHAN TROPIS DAN SUBTROPIS LAINNYA",
			
		},
		{
			SectorEconomy3Id : 43,
			Code : "012300",
			Name : "PERTANIAN BUAH JERUK",
			
		},
		{
			SectorEconomy3Id : 44,
			Code : "012400",
			Name : "PERTANIAN BUAH APEL DAN BUAH BATU (POME AND STONE FRUITS)",
			
		},
		{
			SectorEconomy3Id : 45,
			Code : "012500",
			Name : "PERTANIAN SAYURAN DAN BUAH SEMAK DAN BUAH BIJI KACANG-KACANGAN LAINNYA",
			
		},
		{
			SectorEconomy3Id : 46,
			Code : "012610",
			Name : "PERKEBUNAN BUAH KELAPA",
			
		},
		{
			SectorEconomy3Id : 47,
			Code : "012620",
			Name : "PERKEBUNAN BUAH KELAPA SAWIT",
			
		},
		{
			SectorEconomy3Id : 48,
			Code : "012690",
			Name : "PERKEBUNAN BUAH OLEAGINOUS LAINNYA",
			
		},
		{
			SectorEconomy3Id : 49,
			Code : "012701",
			Name : "PERKEBUNAN TANAMAN KOPI",
			
		},
		{
			SectorEconomy3Id : 50,
			Code : "012702",
			Name : "PERKEBUNAN TANAMAN TEH",
			
		},
		{
			SectorEconomy3Id : 51,
			Code : "012703",
			Name : "PERKEBUNAN TANAMAN COKLAT (KAKAO)",
			
		},
		{
			SectorEconomy3Id : 52,
			Code : "012709",
			Name : "PERTANIAN TANAMAN UNTUK BAHAN MINUMAN LAINNYA",
			
		},
		{
			SectorEconomy3Id : 53,
			Code : "012810",
			Name : "PERKEBUNAN LADA",
			
		},
		{
			SectorEconomy3Id : 54,
			Code : "012820",
			Name : "PERKEBUNAN CENGKEH",
			
		},
		{
			SectorEconomy3Id : 55,
			Code : "012830",
			Name : "PERTANIAN CABAI",
			
		},
		{
			SectorEconomy3Id : 56,
			Code : "012840",
			Name : "PERKEBUNAN TANAMAN AROMATIK/PENYEGAR",
			
		},
		{
			SectorEconomy3Id : 57,
			Code : "012850",
			Name : "PERKEBUNAN TANAMAN OBAT / BAHAN FARMASI",
			
		},
		{
			SectorEconomy3Id : 58,
			Code : "012891",
			Name : "PERKEBUNAN TANAMAN REMPAH PANILI",
			
		},
		{
			SectorEconomy3Id : 59,
			Code : "012892",
			Name : "PERKEBUNAN TANAMAN REMPAH PALA",
			
		},
		{
			SectorEconomy3Id : 60,
			Code : "012899",
			Name : "PERKEBUNAN TANAMAN REMPAH YANG TIDAK DIKLASIFIKASIKAN DI TEMPAT LAIN",
			
		},
		{
			SectorEconomy3Id : 61,
			Code : "012910",
			Name : "PERKEBUNAN KARET DAN TANAMAN PENGHASIL GETAH LAINNYA",
			
		},
		{
			SectorEconomy3Id : 62,
			Code : "012990",
			Name : "PERTANIAN CEMARA DAN TANAMAN TAHUNAN LAINNYA",
			
		},
		{
			SectorEconomy3Id : 63,
			Code : "013010",
			Name : "PERTANIAN TANAMAN HIAS",
			
		},
		{
			SectorEconomy3Id : 64,
			Code : "013020",
			Name : "PERTANIAN PENGEMBANGBIAKAN TANAMAN",
			
		},
		{
			SectorEconomy3Id : 65,
			Code : "014110",
			Name : "PEMBIBITAN DAN BUDIDAYA SAPI POTONG",
			
		},
		{
			SectorEconomy3Id : 66,
			Code : "014120",
			Name : "PEMBIBITAN DAN BUDIDAYA SAPI PERAH",
			
		},
		{
			SectorEconomy3Id : 67,
			Code : "014130",
			Name : "PEMBIBITAN DAN BUDIDAYA KERBAU POTONG",
			
		},
		{
			SectorEconomy3Id : 68,
			Code : "014140",
			Name : "PEMBIBITAN DAN BUDIDAYA KERBAU PERAH",
			
		},
		{
			SectorEconomy3Id : 69,
			Code : "014400",
			Name : "PETERNAKAN DOMBA DAN KAMBING",
			
		},
		{
			SectorEconomy3Id : 70,
			Code : "014500",
			Name : "PETERNAKAN BABI",
			
		},
		{
			SectorEconomy3Id : 71,
			Code : "014600",
			Name : "PETERNAKAN UNGGAS",
			
		},
		{
			SectorEconomy3Id : 72,
			Code : "014900",
			Name : "PETERNAKAN LAINNYA",
			
		},
		{
			SectorEconomy3Id : 73,
			Code : "016000",
			Name : "JASA PENUNJANG PERTANIAN DAN PASCA PANEN",
			
		},
		{
			SectorEconomy3Id : 74,
			Code : "017000",
			Name : "PERBURUAN, PENANGKAPAN DAN PENANGKARAN TUMBUHAN/ SATWA LIAR",
			
		},
		{
			SectorEconomy3Id : 75,
			Code : "021100",
			Name : "PENGUSAHAAN HUTAN TANAMAN",
			
		},
		{
			SectorEconomy3Id : 76,
			Code : "021200",
			Name : "PENGUSAHAAN HUTAN ALAM",
			
		},
		{
			SectorEconomy3Id : 77,
			Code : "021300",
			Name : "PENGUSAHAAN HASIL HUTAN BUKAN KAYU",
			
		},
		{
			SectorEconomy3Id : 78,
			Code : "021400",
			Name : "PENGUSAHAAN PEMBIBITAN TANAMAN KEHUTANAN",
			
		},
		{
			SectorEconomy3Id : 79,
			Code : "022090",
			Name : "USAHA KEHUTANAN LAINNYA",
			
		},
		{
			SectorEconomy3Id : 80,
			Code : "024000",
			Name : "JASA PENUNJANG KEHUTANAN",
			
		},
		{
			SectorEconomy3Id : 81,
			Code : "031111",
			Name : "PENANGKAPAN IKAN TUNA",
			
		},
		{
			SectorEconomy3Id : 82,
			Code : "031119",
			Name : "PENANGKAPAN IKAN LAINNYA",
			
		},
		{
			SectorEconomy3Id : 83,
			Code : "031121",
			Name : "PENANGKAPAN UDANG LAUT",
			
		},
		{
			SectorEconomy3Id : 84,
			Code : "031129",
			Name : "PENANGKAPAN CRUSTACEA LAINNYA DI LAUT",
			
		},
		{
			SectorEconomy3Id : 85,
			Code : "031190",
			Name : "PENANGKAPAN BIOTA AIR LAINNYA DI LAUT",
			
		},
		{
			SectorEconomy3Id : 86,
			Code : "031210",
			Name : "PENANGKAPAN PISCES/IKAN BERSIRIP DI PERAIRAN UMUM",
			
		},
		{
			SectorEconomy3Id : 87,
			Code : "031290",
			Name : "PENANGKAPAN BIOTA AIR LAINNYA DI PERAIRAN UMUM",
			
		},
		{
			SectorEconomy3Id : 88,
			Code : "031300",
			Name : "JASA PENANGKAPAN IKAN DI LAUT",
			
		},
		{
			SectorEconomy3Id : 89,
			Code : "031400",
			Name : "JASA PENANGKAPAN IKAN DI PERAIRAN UMUM",
			
		},
		{
			SectorEconomy3Id : 90,
			Code : "032101",
			Name : "BUDIDAYA BIOTA LAUT UDANG",
			
		},
		{
			SectorEconomy3Id : 91,
			Code : "032102",
			Name : "BUDIDAYA BIOTA LAUT RUMPUT LAUT",
			
		},
		{
			SectorEconomy3Id : 92,
			Code : "032109",
			Name : "BUDIDAYA BIOTA LAUT LAINNYA",
			
		},
		{
			SectorEconomy3Id : 93,
			Code : "032201",
			Name : "BUDIDAYA BIOTA AIR TAWAR UDANG",
			
		},
		{
			SectorEconomy3Id : 94,
			Code : "032202",
			Name : "PEMBENIHAN IKAN AIR TAWAR",
			
		},
		{
			SectorEconomy3Id : 95,
			Code : "032209",
			Name : "BUDIDAYA BIOTA AIR TAWAR LAINNYA",
			
		},
		{
			SectorEconomy3Id : 96,
			Code : "032300",
			Name : "JASA BUDIDAYA IKAN LAUT",
			
		},
		{
			SectorEconomy3Id : 97,
			Code : "032400",
			Name : "JASA BUDIDAYA IKAN AIR TAWAR",
			
		},
		{
			SectorEconomy3Id : 98,
			Code : "032501",
			Name : "BUDIDAYA BIOTA AIR PAYAU UDANG",
			
		},
		{
			SectorEconomy3Id : 99,
			Code : "032509",
			Name : "BUDIDAYA BIOTA AIR PAYAU LAINNYA",
			
		},
		{
			SectorEconomy3Id : 100,
			Code : "032600",
			Name : "JASA BUDIDAYA IKAN AIR PAYAU",
			
		},
		{
			SectorEconomy3Id : 101,
			Code : "050000",
			Name : "PERTAMBANGAN BATU BARA DAN LIGNIT",
			
		},
		{
			SectorEconomy3Id : 102,
			Code : "060001",
			Name : "PERTAMBANGAN MINYAK BUMI DAN GAS ALAM",
			
		},
		{
			SectorEconomy3Id : 103,
			Code : "060002",
			Name : "PENGUSAHAAN TENAGA PANAS BUMI",
			
		},
		{
			SectorEconomy3Id : 104,
			Code : "071000",
			Name : "PERTAMBANGAN PASIR BESI DAN BIJIH BESI",
			
		},
		{
			SectorEconomy3Id : 105,
			Code : "072100",
			Name : "PERTAMBANGAN BIJIH URANIUM DAN THORIUM",
			
		},
		{
			SectorEconomy3Id : 106,
			Code : "072910",
			Name : "PERTAMBANGAN BIJIH TIMAH",
			
		},
		{
			SectorEconomy3Id : 107,
			Code : "072930",
			Name : "PERTAMBANGAN BIJIH BAUKSIT/ALUMINIUM",
			
		},
		{
			SectorEconomy3Id : 108,
			Code : "072940",
			Name : "PERTAMBANGAN BIJIH TEMBAGA",
			
		},
		{
			SectorEconomy3Id : 109,
			Code : "072950",
			Name : "PERTAMBANGAN BIJIH NIKEL",
			
		},
		{
			SectorEconomy3Id : 110,
			Code : "072990",
			Name : "PERTAMBANGAN BAHAN GALIAN LAINNYA YANG TIDAK MENGANDUNG BIJIH BESI",
			
		},
		{
			SectorEconomy3Id : 111,
			Code : "073011",
			Name : "PERTAMBANGAN EMAS",
			
		},
		{
			SectorEconomy3Id : 112,
			Code : "073012",
			Name : "PERTAMBANGAN PERAK",
			
		},
		{
			SectorEconomy3Id : 113,
			Code : "073090",
			Name : "PERTAMBANGAN BIJIH LOGAM MULIA LAINNYA",
			
		},
		{
			SectorEconomy3Id : 114,
			Code : "081000",
			Name : "PENGGALIAN BATU, PASIR DAN TANAH LIAT",
			
		},
		{
			SectorEconomy3Id : 115,
			Code : "089100",
			Name : "PERTAMBANGAN MINERAL, BAHAN KIMIA DAN BAHAN PUPUK",
			
		},
		{
			SectorEconomy3Id : 116,
			Code : "089300",
			Name : "EKSTRAKSI GARAM",
			
		},
		{
			SectorEconomy3Id : 117,
			Code : "089900",
			Name : "PERTAMBANGAN DAN PENGGALIAN LAINNYA YTDL",
			
		},
		{
			SectorEconomy3Id : 118,
			Code : "091000",
			Name : "AKTIVITAS PENUNJANG PERTAMBANGAN MINYAK BUMI DAN GAS ALAM",
			
		},
		{
			SectorEconomy3Id : 119,
			Code : "099000",
			Name : "AKTIVITAS PENUNJANG PERTAMBANGAN DAN PENGGALIAN LAINNYA",
			
		},
		{
			SectorEconomy3Id : 120,
			Code : "101000",
			Name : "INDUSTRI PENGOLAHAN DAN PENGAWETAN DAGING",
			
		},
		{
			SectorEconomy3Id : 121,
			Code : "102000",
			Name : "INDUSTRI PENGOLAHAN DAN PENGAWETAN IKAN DAN BIOTA AIR",
			
		},
		{
			SectorEconomy3Id : 122,
			Code : "103001",
			Name : "INDUSTRI TEMPE & TAHU KEDELAI",
			
		},
		{
			SectorEconomy3Id : 123,
			Code : "103009",
			Name : "INDUSTRI PENGOLAHAN DAN PENGAWETAN LAINNYA BUAH-BUAHAN DAN SAYURAN",
			
		},
		{
			SectorEconomy3Id : 124,
			Code : "104100",
			Name : "INDUSTRI MINYAK DAN LEMAK NABATI DAN HEWANI",
			
		},
		{
			SectorEconomy3Id : 125,
			Code : "104210",
			Name : "INDUSTRI KOPRA, TEPUNG & PELET KELAPA",
			
		},
		{
			SectorEconomy3Id : 126,
			Code : "104230",
			Name : "INDUSTRI MINYAK MENTAH KELAPA & MINYAK GORENG KELAPA",
			
		},
		{
			SectorEconomy3Id : 127,
			Code : "104300",
			Name : "INDUSTRI MINYAK MENTAH/MURNI KELAPA SAWIT (CRUDE PALM OIL) DAN MINYAK GORENG KELAPA SAWIT",
			
		},
		{
			SectorEconomy3Id : 128,
			Code : "104900",
			Name : "INDUSTRI MINYAK MENTAH DAN LEMAK NABATI DAN HEWANI LAINNYA",
			
		},
		{
			SectorEconomy3Id : 129,
			Code : "105000",
			Name : "INDUSTRI PENGOLAHAN SUSU, PRODUK DARI SUSU DAN ES KRIM",
			
		},
		{
			SectorEconomy3Id : 130,
			Code : "106100",
			Name : "INDUSTRI PENGGILINGAN SERELIA DAN BIJI-BIJIAN LAINNYA (BUKAN BERAS DAN JAGUNG)",
			
		},
		{
			SectorEconomy3Id : 131,
			Code : "106200",
			Name : "INDUSTRI PATI DAN PRODUK PATI (BUKAN BERAS DAN JAGUNG)",
			
		},
		{
			SectorEconomy3Id : 132,
			Code : "106300",
			Name : "INDUSTRI PENGGILINGAN BERAS DAN JAGUNG DAN INDUSTRI TEPUNG BERAS DAN JAGUNG",
			
		},
		{
			SectorEconomy3Id : 133,
			Code : "107100",
			Name : "INDUSTRI PRODUK ROTI DAN KUE",
			
		},
		{
			SectorEconomy3Id : 134,
			Code : "107200",
			Name : "INDUSTRI GULA",
			
		},
		{
			SectorEconomy3Id : 135,
			Code : "107300",
			Name : "INDUSTRI KAKAO, COKELAT DAN KEMBANG GULA",
			
		},
		{
			SectorEconomy3Id : 136,
			Code : "107400",
			Name : "INDUSTRI MAKARONI, MIE DAN PRODUK SEJENISNYA",
			
		},
		{
			SectorEconomy3Id : 137,
			Code : "107610",
			Name : "INDUSTRI PENGOLAHAN KOPI",
			
		},
		{
			SectorEconomy3Id : 138,
			Code : "107630",
			Name : "INDUSTRI PENGOLAHAN TEH",
			
		},
		{
			SectorEconomy3Id : 139,
			Code : "107710",
			Name : "INDUSTRI KECAP",
			
		},
		{
			SectorEconomy3Id : 140,
			Code : "107900",
			Name : "INDUSTRI PRODUK MAKANAN LAINNYA",
			
		},
		{
			SectorEconomy3Id : 141,
			Code : "108000",
			Name : "INDUSTRI MAKANAN HEWAN",
			
		},
		{
			SectorEconomy3Id : 142,
			Code : "110000",
			Name : "INDUSTRI MINUMAN",
			
		},
		{
			SectorEconomy3Id : 143,
			Code : "120100",
			Name : "INDUSTRI ROKOK DAN PRODUK TEMBAKAU LAINNYA",
			
		},
		{
			SectorEconomy3Id : 144,
			Code : "120900",
			Name : "INDUSTRI PENGOLAHAN TEMBAKAU LAINNYA",
			
		},
		{
			SectorEconomy3Id : 145,
			Code : "131000",
			Name : "INDUSTRI PEMINTALAN, PENENUNAN DAN PENYELESAIAN AKHIR TEKSTIL",
			
		},
		{
			SectorEconomy3Id : 146,
			Code : "139000",
			Name : "INDUSTRI TEKSTIL LAINNYA",
			
		},
		{
			SectorEconomy3Id : 147,
			Code : "141000",
			Name : "INDUSTRI PAKAIAN JADI DAN PERLENGKAPANNYA, BUKAN PAKAIAN JADI DARI KULIT BERBULU",
			
		},
		{
			SectorEconomy3Id : 148,
			Code : "142000",
			Name : "INDUSTRI PAKAIAN JADI DAN BARANG DARI KULIT BERBULU",
			
		},
		{
			SectorEconomy3Id : 149,
			Code : "143000",
			Name : "INDUSTRI PAKAIAN JADI RAJUTAN DAN SULAMAN/BORDIR",
			
		},
		{
			SectorEconomy3Id : 150,
			Code : "151000",
			Name : "INDUSTRI KULIT DAN BARANG DARI KULIT, TERMASUK KULIT BUATAN",
			
		},
		{
			SectorEconomy3Id : 151,
			Code : "152000",
			Name : "INDUSTRI ALAS KAKI",
			
		},
		{
			SectorEconomy3Id : 152,
			Code : "161000",
			Name : "INDUSTRI PENGGERGAJIAN DAN PENGAWETAN KAYU, ROTAN, BAMBU DAN SEJENISNYA",
			
		},
		{
			SectorEconomy3Id : 153,
			Code : "162100",
			Name : "INDUSTRI KAYU LAPIS, VENEER DAN SEJENISNYA",
			
		},
		{
			SectorEconomy3Id : 154,
			Code : "162900",
			Name : "INDUSTRI BARANG LAINNYA DARI KAYU; INDUSTRI BARANG DARI GABUS DAN BARANG ANYAMAN DARI JERAMI, ROTAN, BAMBU DAN SEJENISNYA",
			
		},
		{
			SectorEconomy3Id : 155,
			Code : "170100",
			Name : "INDUSTRI BUBUR KERTAS, KERTAS DAN PAPAN KERTAS",
			
		},
		{
			SectorEconomy3Id : 156,
			Code : "170200",
			Name : "INDUSTRI KERTAS DAN PAPAN KERTAS BERGELOMBANG DAN WADAH DARI KERTAS DAN PAPAN KERTAS",
			
		},
		{
			SectorEconomy3Id : 157,
			Code : "170900",
			Name : "INDUSTRI BARANG DARI KERTAS DAN PAPAN KERTAS LAINNYA",
			
		},
		{
			SectorEconomy3Id : 158,
			Code : "181000",
			Name : "INDUSTRI PENCETAKAN DAN KEGIATAN YBDI",
			
		},
		{
			SectorEconomy3Id : 159,
			Code : "182000",
			Name : "REPRODUKSI MEDIA REKAMAN",
			
		},
		{
			SectorEconomy3Id : 160,
			Code : "191000",
			Name : "INDUSTRI PRODUK DARI BATU BARA",
			
		},
		{
			SectorEconomy3Id : 161,
			Code : "192100",
			Name : "INDUSTRI BAHAN BAKAR DAN MINYAK PELUMAS HASIL PENGILANGAN MINYAK BUMI",
			
		},
		{
			SectorEconomy3Id : 162,
			Code : "192900",
			Name : "INDUSTRI BRIKET BATU BARA",
			
		},
		{
			SectorEconomy3Id : 163,
			Code : "201100",
			Name : "INDUSTRI KIMIA DASAR",
			
		},
		{
			SectorEconomy3Id : 164,
			Code : "201200",
			Name : "INDUSTRI PUPUK DAN BAHAN SENYAWA NITROGEN",
			
		},
		{
			SectorEconomy3Id : 165,
			Code : "201300",
			Name : "INDUSTRI PLASTIK DAN KARET BUATAN DALAM BENTUK DASAR",
			
		},
		{
			SectorEconomy3Id : 166,
			Code : "202100",
			Name : "INDUSTRI PESTISIDA DAN PRODUK AGROKIMIA LAINNYA",
			
		},
		{
			SectorEconomy3Id : 167,
			Code : "202200",
			Name : "INDUSTRI CAT DAN TINTA CETAK, PERNIS DAN BAHAN PELAPISAN SEJENISNYA DAN LAK",
			
		},
		{
			SectorEconomy3Id : 168,
			Code : "202300",
			Name : "INDUSTRI SABUN DAN DETERJEN, BAHAN PEMBERSIH DAN PENGILAP, PARFUM DAN KOSMETIK",
			
		},
		{
			SectorEconomy3Id : 169,
			Code : "202940",
			Name : "INDUSTRI MINYAK ATSIRI",
			
		},
		{
			SectorEconomy3Id : 170,
			Code : "202990",
			Name : "INDUSTRI BARANG KIMIA LAINNYA YTDL",
			
		},
		{
			SectorEconomy3Id : 171,
			Code : "203000",
			Name : "INDUSTRI SERAT BUATAN",
			
		},
		{
			SectorEconomy3Id : 172,
			Code : "210000",
			Name : "INDUSTRI FARMASI, PRODUK OBAT KIMIA DAN OBAT TRADISIONAL",
			
		},
		{
			SectorEconomy3Id : 173,
			Code : "221210",
			Name : "INDUSTRI PENGASAPAN KARET",
			
		},
		{
			SectorEconomy3Id : 174,
			Code : "221220",
			Name : "INDUSTRI REMILLING KARET",
			
		},
		{
			SectorEconomy3Id : 175,
			Code : "221230",
			Name : "INDUSTRI KARET REMAH (CRUMB RUBBER)",
			
		},
		{
			SectorEconomy3Id : 176,
			Code : "221900",
			Name : "INDUSTRI BARANG DARI KARET LAINNYA",
			
		},
		{
			SectorEconomy3Id : 177,
			Code : "222000",
			Name : "INDUSTRI BARANG DARI PLASTIK",
			
		},
		{
			SectorEconomy3Id : 178,
			Code : "231000",
			Name : "INDUSTRI KACA DAN BARANG DARI KACA",
			
		},
		{
			SectorEconomy3Id : 179,
			Code : "239200",
			Name : "INDUSTRI BAHAN BANGUNAN DARI TANAH LIAT/KERAMIK",
			
		},
		{
			SectorEconomy3Id : 180,
			Code : "239301",
			Name : "INDUSTRI BARANG PORSELEN BUKAN BAHAN BANGUNAN",
			
		},
		{
			SectorEconomy3Id : 181,
			Code : "239302",
			Name : "INDUSTRI BARANG TANAH LIAT/KERAMIK BUKAN BAHAN BANGUNAN",
			
		},
		{
			SectorEconomy3Id : 182,
			Code : "239400",
			Name : "INDUSTRI SEMEN, KAPUR DAN GIPS",
			
		},
		{
			SectorEconomy3Id : 183,
			Code : "239600",
			Name : "INDUSTRI BARANG DARI BATU",
			
		},
		{
			SectorEconomy3Id : 184,
			Code : "239900",
			Name : "INDUSTRI BARANG GALIAN BUKAN LOGAM LAINNYA YTDL",
			
		},
		{
			SectorEconomy3Id : 185,
			Code : "241000",
			Name : "INDUSTRI LOGAM DASAR BESI DAN BAJA",
			
		},
		{
			SectorEconomy3Id : 186,
			Code : "242060",
			Name : "INDUSTRI PENGOLAHAN URANIUM DAN BIJIH URANIUM",
			
		},
		{
			SectorEconomy3Id : 187,
			Code : "242090",
			Name : "INDUSTRI LOGAM DASAR MULIA DAN LOGAM DASAR BUKAN BESI LAINNYA",
			
		},
		{
			SectorEconomy3Id : 188,
			Code : "243100",
			Name : "INDUSTRI PENGECORAN BESI DAN BAJA",
			
		},
		{
			SectorEconomy3Id : 189,
			Code : "243200",
			Name : "INDUSTRI PENGECORAN LOGAM BUKAN BESI DAN BAJA",
			
		},
		{
			SectorEconomy3Id : 190,
			Code : "251000",
			Name : "INDUSTRI BARANG LOGAM SIAP PASANG UNTUK BANGUNAN, TANGKI, TANDON AIR DAN GENERATOR UAP",
			
		},
		{
			SectorEconomy3Id : 191,
			Code : "259300",
			Name : "INDUSTRI ALAT POTONG, PERKAKAS TANGAN DAN PERALATAN UMUM",
			
		},
		{
			SectorEconomy3Id : 192,
			Code : "259900",
			Name : "INDUSTRI BARANG LOGAM LAINNYA YTDL",
			
		},
		{
			SectorEconomy3Id : 193,
			Code : "261000",
			Name : "INDUSTRI KOMPONEN DAN PAPAN ELEKTRONIK",
			
		},
		{
			SectorEconomy3Id : 194,
			Code : "262000",
			Name : "INDUSTRI KOMPUTER DAN PERLENGKAPANNYA",
			
		},
		{
			SectorEconomy3Id : 195,
			Code : "263000",
			Name : "INDUSTRI PERALATAN KOMUNIKASI",
			
		},
		{
			SectorEconomy3Id : 196,
			Code : "264000",
			Name : "INDUSTRI PERALATAN AUDIO DAN VIDEO ELEKTRONIK",
			
		},
		{
			SectorEconomy3Id : 197,
			Code : "265100",
			Name : "INDUSTRI ALAT UKUR, ALAT UJI, PERALATAN NAVIGASI DAN KONTROL",
			
		},
		{
			SectorEconomy3Id : 198,
			Code : "265200",
			Name : "INDUSTRI ALAT UKUR WAKTU",
			
		},
		{
			SectorEconomy3Id : 199,
			Code : "266000",
			Name : "INDUSTRI PERALATAN IRADIASI, ELEKTROMEDIKAL DAN ELEKTROTERAPI",
			
		},
		{
			SectorEconomy3Id : 200,
			Code : "267000",
			Name : "INDUSTRI PERALATAN FOTOGRAFI DAN INSTRUMEN OPTIK BUKAN KACA MATA",
			
		},
		{
			SectorEconomy3Id : 201,
			Code : "269000",
			Name : "INDUSTRI KOMPUTER, BARANG ELEKTRONIK DAN OPTIK LAINNYA",
			
		},
		{
			SectorEconomy3Id : 202,
			Code : "271100",
			Name : "INDUSTRI MOTOR LISTRIK, GENERATOR DAN TRANSFORMATOR",
			
		},
		{
			SectorEconomy3Id : 203,
			Code : "271200",
			Name : "INDUSTRI PERALATAN PENGONTROL DAN PENDISTRIBUSIAN LISTRIK",
			
		},
		{
			SectorEconomy3Id : 204,
			Code : "272000",
			Name : "INDUSTRI BATU BATERAI DAN AKUMULATOR LISTRIK",
			
		},
		{
			SectorEconomy3Id : 205,
			Code : "273000",
			Name : "INDUSTRI KABEL DAN PERLENGKAPANNYA",
			
		},
		{
			SectorEconomy3Id : 206,
			Code : "274000",
			Name : "INDUSTRI PERALATAN PENERANGAN LISTRIK (TERMASUK PERALATAN PENERANGAN BUKAN LISTRIK)",
			
		},
		{
			SectorEconomy3Id : 207,
			Code : "275000",
			Name : "INDUSTRI PERALATAN RUMAH TANGGA",
			
		},
		{
			SectorEconomy3Id : 208,
			Code : "279000",
			Name : "INDUSTRI PERALATAN LISTRIK LAINNYA",
			
		},
		{
			SectorEconomy3Id : 209,
			Code : "281000",
			Name : "INDUSTRI MESIN UNTUK KEPERLUAN UMUM",
			
		},
		{
			SectorEconomy3Id : 210,
			Code : "282100",
			Name : "INDUSTRI MESIN PERTANIAN DAN KEHUTANAN",
			
		},
		{
			SectorEconomy3Id : 211,
			Code : "282400",
			Name : "INDUSTRI MESIN PENAMBANGAN, PENGGALIAN DAN KONSTRUKSI",
			
		},
		{
			SectorEconomy3Id : 212,
			Code : "282500",
			Name : "INDUSTRI MESIN PENGOLAHAN MAKANAN, MINUMAN DAN TEMBAKAU",
			
		},
		{
			SectorEconomy3Id : 213,
			Code : "282600",
			Name : "INDUSTRI MESIN TEKSTIL, PAKAIAN JADI DAN PRODUK KULIT",
			
		},
		{
			SectorEconomy3Id : 214,
			Code : "282900",
			Name : "INDUSTRI MESIN KEPERLUAN KHUSUS LAINNYA",
			
		},
		{
			SectorEconomy3Id : 215,
			Code : "291000",
			Name : "INDUSTRI KENDARAAN BERMOTOR RODA EMPAT ATAU LEBIH",
			
		},
		{
			SectorEconomy3Id : 216,
			Code : "292000",
			Name : "INDUSTRI KAROSERI KENDARAAN BERMOTOR RODA EMPAT ATAU LEBIH DAN INDUSTRI TRAILER DAN SEMI TRAILER",
			
		},
		{
			SectorEconomy3Id : 217,
			Code : "293000",
			Name : "INDUSTRI SUKU CADANG DAN AKSESORI KENDARAAN BERMOTOR RODA EMPAT ATAU LEBIH",
			
		},
		{
			SectorEconomy3Id : 218,
			Code : "301000",
			Name : "INDUSTRI PEMBUATAN KAPAL DAN PERAHU",
			
		},
		{
			SectorEconomy3Id : 219,
			Code : "302000",
			Name : "INDUSTRI LOKOMOTIF DAN GERBONG KERETA",
			
		},
		{
			SectorEconomy3Id : 220,
			Code : "303000",
			Name : "INDUSTRI PESAWAT TERBANG DAN PERLENGKAPANNYA",
			
		},
		{
			SectorEconomy3Id : 221,
			Code : "309110",
			Name : "INDUSTRI SEPEDA MOTOR RODA DUA DAN TIGA",
			
		},
		{
			SectorEconomy3Id : 222,
			Code : "309900",
			Name : "INDUSTRI ALAT ANGKUTAN LAINNYA YTDL",
			
		},
		{
			SectorEconomy3Id : 223,
			Code : "310000",
			Name : "INDUSTRI FURNITUR",
			
		},
		{
			SectorEconomy3Id : 224,
			Code : "320000",
			Name : "INDUSTRI PENGOLAHAN LAINNYA",
			
		},
		{
			SectorEconomy3Id : 225,
			Code : "330000",
			Name : "REPARASI DAN PEMASANGAN MESIN DAN PERALATAN",
			
		},
		{
			SectorEconomy3Id : 226,
			Code : "351001",
			Name : "KETENAGALISTRIKAN PEDESAAN",
			
		},
		{
			SectorEconomy3Id : 227,
			Code : "351002",
			Name : "KETENAGALISTRIKAN LAINNYA",
			
		},
		{
			SectorEconomy3Id : 228,
			Code : "352000",
			Name : "PENGADAAN DAN DISTRIBUSI GAS ALAM DAN BUATAN",
			
		},
		{
			SectorEconomy3Id : 229,
			Code : "353000",
			Name : "PENGADAAN UAP/AIR PANAS, UDARA DINGIN DAN PRODUKSI ES",
			
		},
		{
			SectorEconomy3Id : 230,
			Code : "360000",
			Name : "PENGELOLAAN AIR",
			
		},
		{
			SectorEconomy3Id : 231,
			Code : "370000",
			Name : "PENGELOLAAN AIR LIMBAH",
			
		},
		{
			SectorEconomy3Id : 232,
			Code : "380000",
			Name : "PENGELOLAAN DAN DAUR ULANG SAMPAH",
			
		},
		{
			SectorEconomy3Id : 233,
			Code : "390000",
			Name : "AKTIVITAS REMEDIASI DAN PENGELOLAAN SAMPAH LAINNYA",
			
		},
		{
			SectorEconomy3Id : 234,
			Code : "410111",
			Name : "KONSTRUKSI PERUMAHAN SEDERHANA BANK TABUNGAN NEGARA",
			
		},
		{
			SectorEconomy3Id : 235,
			Code : "410112",
			Name : "KONSTRUKSI PERUMAHAN SEDERHANA PERUMNAS",
			
		},
		{
			SectorEconomy3Id : 236,
			Code : "410113",
			Name : "KONSTRUKSI PERUMAHAN SEDERHANA LAINNYA TIPE S.D. 21",
			
		},
		{
			SectorEconomy3Id : 237,
			Code : "410114",
			Name : "KONSTRUKSI PERUMAHAN SEDERHANA LAINNYA TIPE 22 S.D. 70",
			
		},
		{
			SectorEconomy3Id : 238,
			Code : "410115",
			Name : "KONSTRUKSI PERUMAHAN MENENGAH, BESAR, MEWAH (TIPE DIATAS 70)",
			
		},
		{
			SectorEconomy3Id : 239,
			Code : "410119",
			Name : "KONSTRUKSI GEDUNG TEMPAT TINGGAL LAINNYA",
			
		},
		{
			SectorEconomy3Id : 240,
			Code : "410120",
			Name : "KONSTRUKSI GEDUNG PERKANTORAN",
			
		},
		{
			SectorEconomy3Id : 241,
			Code : "410130",
			Name : "KONSTRUKSI GEDUNG INDUSTRI",
			
		},
		{
			SectorEconomy3Id : 242,
			Code : "410141",
			Name : "KONSTRUKSI GEDUNG PERBELANJAAN PASAR INPRES",
			
		},
		{
			SectorEconomy3Id : 243,
			Code : "410149",
			Name : "KONSTRUKSI GEDUNG PERBELANJAAN LAINNYA",
			
		},
		{
			SectorEconomy3Id : 244,
			Code : "410190",
			Name : "KONSTRUKSI GEDUNG LAINNYA",
			
		},
		{
			SectorEconomy3Id : 245,
			Code : "421101",
			Name : "KONSTRUKSI JALAN TOL",
			
		},
		{
			SectorEconomy3Id : 246,
			Code : "421102",
			Name : "KONSTRUKSI JALAN RAYA SELAIN TOL",
			
		},
		{
			SectorEconomy3Id : 247,
			Code : "421103",
			Name : "KONSTRUKSI JEMBATAN DAN JALAN LAYANG",
			
		},
		{
			SectorEconomy3Id : 248,
			Code : "421104",
			Name : "KONSTRUKSI JALAN REL DAN JEMBATAN REL",
			
		},
		{
			SectorEconomy3Id : 249,
			Code : "421109",
			Name : "KONSTRUKSI JALAN RAYA LAINNYA",
			
		},
		{
			SectorEconomy3Id : 250,
			Code : "422110",
			Name : "KONSTRUKSI JARINGAN IRIGASI",
			
		},
		{
			SectorEconomy3Id : 251,
			Code : "422131",
			Name : "KONSTRUKSI BANGUNAN LISTRIK PEDESAAN",
			
		},
		{
			SectorEconomy3Id : 252,
			Code : "422139",
			Name : "KONSTRUKSI BANGUNAN ELEKTRIKAL DAN KOMUNIKASI LAINNYA",
			
		},
		{
			SectorEconomy3Id : 253,
			Code : "422190",
			Name : "KONSTRUKSI JARINGAN ELEKTRIKAL DAN TELEKOMUNIKASI LAINNYA",
			
		},
		{
			SectorEconomy3Id : 254,
			Code : "429120",
			Name : "KONSTRUKSI BANGUNAN PELABUHAN BUKAN PERIKANAN",
			
		},
		{
			SectorEconomy3Id : 255,
			Code : "429190",
			Name : "KONSTRUKSI BANGUNAN SIPIL LAINNYA YTDL",
			
		},
		{
			SectorEconomy3Id : 256,
			Code : "431201",
			Name : "PENYIAPAN TANAH PEMUKIMAN TRANSMIGRASI (PTPT)",
			
		},
		{
			SectorEconomy3Id : 257,
			Code : "431202",
			Name : "PENCETAKAN LAHAN SAWAH",
			
		},
		{
			SectorEconomy3Id : 258,
			Code : "431209",
			Name : "PENYIAPAN LAHAN LAINNYA DAN PEMBONGKARAN",
			
		},
		{
			SectorEconomy3Id : 259,
			Code : "432000",
			Name : "INSTALASI SISTEM KELISTRIKAN, AIR (PIPA) DAN INSTALASI KONSTRUKSI LAINNYA",
			
		},
		{
			SectorEconomy3Id : 260,
			Code : "433000",
			Name : "PENYELESAIAN KONSTRUKSI BANGUNAN",
			
		},
		{
			SectorEconomy3Id : 261,
			Code : "439050",
			Name : "PENYEWAAN ALAT KONSTRUKSI DENGAN OPERATOR",
			
		},
		{
			SectorEconomy3Id : 262,
			Code : "439090",
			Name : "KONSTRUKSI KHUSUS LAINNYA YTDL",
			
		},
		{
			SectorEconomy3Id : 263,
			Code : "451000",
			Name : "PERDAGANGAN MOBIL",
			
		},
		{
			SectorEconomy3Id : 264,
			Code : "452000",
			Name : "REPARASI DAN PERAWATAN MOBIL",
			
		},
		{
			SectorEconomy3Id : 265,
			Code : "453000",
			Name : "PERDAGANGAN SUKU CADANG DAN AKSESORI MOBIL",
			
		},
		{
			SectorEconomy3Id : 266,
			Code : "454001",
			Name : "PERDAGANGAN SEPEDA MOTOR",
			
		},
		{
			SectorEconomy3Id : 267,
			Code : "454002",
			Name : "PERDAGANGAN SUKU CADANG SEPEDA MOTOR DAN AKSESORINYA",
			
		},
		{
			SectorEconomy3Id : 268,
			Code : "454003",
			Name : "REPARASI DAN PERAWATAN SEPEDA MOTOR",
			
		},
		{
			SectorEconomy3Id : 269,
			Code : "461000",
			Name : "PERDAGANGAN BESAR ATAS DASAR BALAS JASA (FEE) ATAU KONTRAK",
			
		},
		{
			SectorEconomy3Id : 270,
			Code : "462011",
			Name : "PERDAGANGAN BESAR JAGUNG",
			
		},
		{
			SectorEconomy3Id : 271,
			Code : "462019",
			Name : "PERDAGANGAN BESAR PADI DAN PALAWIJA LAINNYA",
			
		},
		{
			SectorEconomy3Id : 272,
			Code : "462020",
			Name : "PERDAGANGAN BESAR BUAH YANG MENGANDUNG MINYAK",
			
		},
		{
			SectorEconomy3Id : 273,
			Code : "462040",
			Name : "PERDAGANGAN BESAR TEMBAKAU RAJANGAN",
			
		},
		{
			SectorEconomy3Id : 274,
			Code : "462050",
			Name : "PERDAGANGAN BESAR BINATANG HIDUP",
			
		},
		{
			SectorEconomy3Id : 275,
			Code : "462060",
			Name : "PERDAGANGAN BESAR HASIL PERIKANAN",
			
		},
		{
			SectorEconomy3Id : 276,
			Code : "462071",
			Name : "PERDAGANGAN KAYU",
			
		},
		{
			SectorEconomy3Id : 277,
			Code : "462079",
			Name : "PERDAGANGAN BESAR HASIL KEHUTANAN DAN PERBURUAN LAINNYA",
			
		},
		{
			SectorEconomy3Id : 278,
			Code : "462080",
			Name : "PERDAGANGAN BESAR KULIT DAN KULIT JANGAT",
			
		},
		{
			SectorEconomy3Id : 279,
			Code : "462091",
			Name : "PERDAGANGAN KARET",
			
		},
		{
			SectorEconomy3Id : 280,
			Code : "462092",
			Name : "PERDAGANGAN CENGKEH",
			
		},
		{
			SectorEconomy3Id : 281,
			Code : "462093",
			Name : "PERDAGANGAN LADA",
			
		},
		{
			SectorEconomy3Id : 282,
			Code : "462094",
			Name : "PERDAGANGAN KAPAS",
			
		},
		{
			SectorEconomy3Id : 283,
			Code : "462095",
			Name : "PERDAGANGAN BIJI KELAPA SAWIT",
			
		},
		{
			SectorEconomy3Id : 284,
			Code : "462099",
			Name : "PERDAGANGAN BESAR HASIL PERTANIAN DAN HEWAN HIDUP LAINNYA",
			
		},
		{
			SectorEconomy3Id : 285,
			Code : "463110",
			Name : "PERDAGANGAN BESAR BERAS",
			
		},
		{
			SectorEconomy3Id : 286,
			Code : "463141",
			Name : "PERDAGANGAN BESAR KOPI",
			
		},
		{
			SectorEconomy3Id : 287,
			Code : "463142",
			Name : "PERDAGANGAN BESAR TEH",
			
		},
		{
			SectorEconomy3Id : 288,
			Code : "463150",
			Name : "PERDAGANGAN BESAR MINYAK DAN LEMAK NABATI",
			
		},
		{
			SectorEconomy3Id : 289,
			Code : "463190",
			Name : "PERDAGANGAN BESAR BAHAN MAKANAN DAN MINUMAN HASIL PERTANIAN LAINNYA",
			
		},
		{
			SectorEconomy3Id : 290,
			Code : "463201",
			Name : "PERDAGANGAN BESER UDANG OLAHAN",
			
		},
		{
			SectorEconomy3Id : 291,
			Code : "463209",
			Name : "PERDAGANGAN BESAR BAHAN MAKANAN DAN MINUMAN HASIL PETERNAKAN DAN PERIKANAN LAINNYA",
			
		},
		{
			SectorEconomy3Id : 292,
			Code : "463301",
			Name : "PERDAGANGAN BESAR GULA, COKLAT DAN KEMBANG GULA",
			
		},
		{
			SectorEconomy3Id : 293,
			Code : "463302",
			Name : "PERDAGANGAN BESAR ROKOK DAN TEMBAKAU",
			
		},
		{
			SectorEconomy3Id : 294,
			Code : "463309",
			Name : "PERDAGANGAN BESAR MAKANAN DAN MINUMAN LAINNYA",
			
		},
		{
			SectorEconomy3Id : 295,
			Code : "464110",
			Name : "PERDAGANGAN BESAR TEKSTIL",
			
		},
		{
			SectorEconomy3Id : 296,
			Code : "464120",
			Name : "PERDAGANGAN BESAR PAKAIAN",
			
		},
		{
			SectorEconomy3Id : 297,
			Code : "464130",
			Name : "PERDAGANGAN BESAR ALAS KAKI",
			
		},
		{
			SectorEconomy3Id : 298,
			Code : "464190",
			Name : "PERDAGANGAN BESAR TEKSTIL, PAKAIAN DAN ALAS KAKI LAINNYA",
			
		},
		{
			SectorEconomy3Id : 299,
			Code : "464900",
			Name : "PERDAGANGAN BESAR BARANG KEPERLUAN RUMAH TANGGA LAINNYA",
			
		},
		{
			SectorEconomy3Id : 300,
			Code : "465000",
			Name : "PERDAGANGAN BESAR MESIN, PERALATAN DAN PERLENGKAPANNYA",
			
		},
		{
			SectorEconomy3Id : 301,
			Code : "466100",
			Name : "PERDAGANGAN BESAR BAHAN BAKAR PADAT, CAIR DAN GAS DAN PRODUK YBDI",
			
		},
		{
			SectorEconomy3Id : 302,
			Code : "466200",
			Name : "PERDAGANGAN BESAR LOGAM DAN BIJIH LOGAM",
			
		},
		{
			SectorEconomy3Id : 303,
			Code : "466301",
			Name : "PERDAGANGAN BESAR BAHAN KONSTRUKSI DARI KAYU",
			
		},
		{
			SectorEconomy3Id : 304,
			Code : "466309",
			Name : "PERDAGANGAN BESAR BAHAN KONSTRUKSI LAINNYA",
			
		},
		{
			SectorEconomy3Id : 305,
			Code : "466920",
			Name : "PERDAGANGAN BESAR PUPUK DAN PRODUK AGROKIMIA",
			
		},
		{
			SectorEconomy3Id : 306,
			Code : "466930",
			Name : "PERDAGANGAN BESAR ALAT LABORATORIUM, FARMASI DAN KEDOKTERAN",
			
		},
		{
			SectorEconomy3Id : 307,
			Code : "466950",
			Name : "PERDAGANGAN BESAR KERTAS DAN KARTON",
			
		},
		{
			SectorEconomy3Id : 308,
			Code : "466970",
			Name : "PERDAGANGAN BESAR BARANG BEKAS DAN SISA-SISA TAK TERPAKAI (SCRAP)",
			
		},
		{
			SectorEconomy3Id : 309,
			Code : "466990",
			Name : "PERDAGANGAN BESAR PRODUK LAINNYA YTDL",
			
		},
		{
			SectorEconomy3Id : 310,
			Code : "471100",
			Name : "PERDAGANGAN ECERAN YANG UTAMANYA MAKANAN, MINUMAN ATAU TEMBAKAU DI TOKO",
			
		},
		{
			SectorEconomy3Id : 311,
			Code : "471900",
			Name : "PERDAGANGAN ECERAN BERBAGAI MACAM BARANG YANG DIDOMINASI OLEH BARANG BUKAN MAKANAN DAN TEMBAKAU DI TOKO",
			
		},
		{
			SectorEconomy3Id : 312,
			Code : "472001",
			Name : "PERDAGANGAN ECERAN KHUSUS KOMODITI MAKANAN DARI HASIL PERTANIAN DI TOKO",
			
		},
		{
			SectorEconomy3Id : 313,
			Code : "472009",
			Name : "PERDAGANGAN ECERAN KHUSUS MAKANAN, MINUMAN DAN TEMBAKAU LAINNYA DI TOKO",
			
		},
		{
			SectorEconomy3Id : 314,
			Code : "473000",
			Name : "PERDAGANGAN ECERAN KHUSUS BAHAN BAKAR KENDARAAN BERMOTOR",
			
		},
		{
			SectorEconomy3Id : 315,
			Code : "474000",
			Name : "PERDAGANGAN ECERAN KHUSUS PERALATAN INFORMASI DAN KOMUNIKASI DI TOKO",
			
		},
		{
			SectorEconomy3Id : 316,
			Code : "475100",
			Name : "PERDAGANGAN ECERAN KHUSUS TEKSTIL DI TOKO",
			
		},
		{
			SectorEconomy3Id : 317,
			Code : "475200",
			Name : "PERDAGANGAN ECERAN KHUSUS BARANG DAN BAHAN BANGUNAN, CAT DAN KACA DI TOKO",
			
		},
		{
			SectorEconomy3Id : 318,
			Code : "475900",
			Name : "PERDAGANGAN ECERAN KHUSUS FURNITUR, PERALATAN LISTRIK RUMAH TANGGA, PERALATAN PENERANGAN DAN PERALATAN RUMAH TANGGA LAINNYA DI TOKO",
			
		},
		{
			SectorEconomy3Id : 319,
			Code : "476000",
			Name : "PERDAGANGAN ECERAN KHUSUS BARANG BUDAYA DAN REKREASI DI TOKO KHUSUS",
			
		},
		{
			SectorEconomy3Id : 320,
			Code : "477100",
			Name : "PERDAGANGAN ECERAN KHUSUS PAKAIAN, ALAS KAKI DAN BARANG DARI KULIT DI TOKO",
			
		},
		{
			SectorEconomy3Id : 321,
			Code : "477200",
			Name : "PERDAGANGAN ECERAN KHUSUS BAHAN KIMIA, BARANG FARMASI, ALAT KEDOKTERAN, PARFUM DAN KOSMETIK DI TOKO",
			
		},
		{
			SectorEconomy3Id : 322,
			Code : "477300",
			Name : "PERDAGANGAN ECERAN KHUSUS BARANG BARU LAINNYA DI TOKO",
			
		},
		{
			SectorEconomy3Id : 323,
			Code : "477400",
			Name : "PERDAGANGAN ECERAN KHUSUS BARANG BEKAS DI TOKO",
			
		},
		{
			SectorEconomy3Id : 324,
			Code : "477700",
			Name : "PERDAGANGAN ECERAN BAHAN BAKAR BUKAN BAHAN BAKAR UNTUK KENDARAAN BERMOTOR DI TOKO",
			
		},
		{
			SectorEconomy3Id : 325,
			Code : "477800",
			Name : "PERDAGANGAN ECERAN BARANG KERAJINAN DAN LUKISAN DI TOKO",
			
		},
		{
			SectorEconomy3Id : 326,
			Code : "477900",
			Name : "PERDAGANGAN ECERAN KHUSUS BARANG LAINNYA YTDL",
			
		},
		{
			SectorEconomy3Id : 327,
			Code : "478100",
			Name : "PERDAGANGAN ECERAN KAKI LIMA DAN LOS PASAR KOMODITI HASIL PERTANIAN",
			
		},
		{
			SectorEconomy3Id : 328,
			Code : "478200",
			Name : "PERDAGANGAN ECERAN KAKI LIMA DAN LOS PASAR MAKANAN, MINUMAN DAN PRODUK TEMBAKAU HASIL INDUSTRI PENGOLAHAN",
			
		},
		{
			SectorEconomy3Id : 329,
			Code : "478300",
			Name : "PERDAGANGAN ECERAN KAKI LIMA DAN LOS PASAR TEKSTIL, PAKAIAN DAN ALAS KAKI",
			
		},
		{
			SectorEconomy3Id : 330,
			Code : "478400",
			Name : "PERDAGANGAN ECERAN KAKI LIMA DAN LOS PASAR BAHAN KIMIA, FARMASI, KOSMETIK DAN YBDI",
			
		},
		{
			SectorEconomy3Id : 331,
			Code : "478600",
			Name : "PERDAGANGAN ECERAN KAKI LIMA DAN LOS PASAR PERLENGKAPAN RUMAH TANGGA",
			
		},
		{
			SectorEconomy3Id : 332,
			Code : "478700",
			Name : "PERDAGANGAN ECERAN KAKI LIMA DAN LOS PASAR KERTAS, BARANG DARI KERTAS, ALAT TULIS, BARANG CETAKAN, ALAT OLAHRAGA, ALAT MUSIK, ALAT FOTOGRAFI DAN KOMPUTER",
			
		},
		{
			SectorEconomy3Id : 333,
			Code : "478800",
			Name : "PERDAGANGAN ECERAN KAKI LIMA DAN LOS PASAR BARANG KERAJINAN, MAINAN ANAK-ANAK DAN LUKISAN",
			
		},
		{
			SectorEconomy3Id : 334,
			Code : "478920",
			Name : "PERDAGANGAN ECERAN KAKI LIMA DAN LOS PASAR BAHAN BAKAR MINYAK, GAS, MINYAK PELUMAS DAN BAHAN BAKAR LAINNYA",
			
		},
		{
			SectorEconomy3Id : 335,
			Code : "478940",
			Name : "PERDAGANGAN ECERAN KAKI LIMA DAN LOS PASAR BARANG BEKAS PERLENGKAPAN RUMAH TANGGA",
			
		},
		{
			SectorEconomy3Id : 336,
			Code : "478990",
			Name : "PERDAGANGAN ECERAN KAKI LIMA DAN LOS PASAR BARANG LAINNYA",
			
		},
		{
			SectorEconomy3Id : 337,
			Code : "479100",
			Name : "PERDAGANGAN ECERAN MELALUI PEMESANAN POS ATAU INTERNET",
			
		},
		{
			SectorEconomy3Id : 338,
			Code : "479900",
			Name : "PERDAGANGAN ECERAN BUKAN DI TOKO, KIOS, KAKI LIMA DAN LOS PASAR LAINNYA",
			
		},
		{
			SectorEconomy3Id : 339,
			Code : "491000",
			Name : "ANGKUTAN JALAN REL",
			
		},
		{
			SectorEconomy3Id : 340,
			Code : "492100",
			Name : "ANGKUTAN BUS BERTRAYEK",
			
		},
		{
			SectorEconomy3Id : 341,
			Code : "492210",
			Name : "ANGKUTAN BUS PARIWISATA",
			
		},
		{
			SectorEconomy3Id : 342,
			Code : "492290",
			Name : "ANGKUTAN BUS TIDAK BERTRAYEK LAINNYA",
			
		},
		{
			SectorEconomy3Id : 343,
			Code : "493000",
			Name : "ANGKUTAN MELALUI SALURAN PIPA",
			
		},
		{
			SectorEconomy3Id : 344,
			Code : "494100",
			Name : "ANGKUTAN DARAT BUKAN BUS UNTUK PENUMPANG, BERTRAYEK",
			
		},
		{
			SectorEconomy3Id : 345,
			Code : "494200",
			Name : "ANGKUTAN DARAT LAINNYA UNTUK PENUMPANG",
			
		},
		{
			SectorEconomy3Id : 346,
			Code : "494300",
			Name : "ANGKUTAN DARAT UNTUK BARANG",
			
		},
		{
			SectorEconomy3Id : 347,
			Code : "494501",
			Name : "ANGKUTAN JALAN REL WISATA",
			
		},
		{
			SectorEconomy3Id : 348,
			Code : "494509",
			Name : "ANGKUTAN JALAN REL LAINNYA",
			
		},
		{
			SectorEconomy3Id : 349,
			Code : "501100",
			Name : "ANGKUTAN LAUT DALAM NEGERI UNTUK PENUMPANG",
			
		},
		{
			SectorEconomy3Id : 350,
			Code : "501130",
			Name : "ANGKUTAN LAUT UNTUK WISATA",
			
		},
		{
			SectorEconomy3Id : 351,
			Code : "501190",
			Name : "ANGKUTAN LAUT DALAM NEGERI UNTUK PENUMPANG SELAIN WISATA",
			
		},
		{
			SectorEconomy3Id : 352,
			Code : "501200",
			Name : "ANGKUTAN LAUT LUAR NEGERI UNTUK PENUMPANG",
			
		},
		{
			SectorEconomy3Id : 353,
			Code : "501300",
			Name : "ANGKUTAN LAUT DALAM NEGERI UNTUK BARANG",
			
		},
		{
			SectorEconomy3Id : 354,
			Code : "501400",
			Name : "ANGKUTAN LAUT LUAR NEGERI UNTUK BARANG",
			
		},
		{
			SectorEconomy3Id : 355,
			Code : "502101",
			Name : "ANGKUTAN SUNGAI DAN DANAU UNTUK WISATA DAN YBDI",
			
		},
		{
			SectorEconomy3Id : 356,
			Code : "502102",
			Name : "ANGKUTAN PENYEBERANGAN UNTUK PENUMPANG",
			
		},
		{
			SectorEconomy3Id : 357,
			Code : "502200",
			Name : "ANGKUTAN SUNGAI, DANAU DAN PENYEBERANGAN UNTUK BARANG",
			
		},
		{
			SectorEconomy3Id : 358,
			Code : "511001",
			Name : "ANGKUTAN UDARA BERJADWAL UNTUK PENUMPANG",
			
		},
		{
			SectorEconomy3Id : 359,
			Code : "511002",
			Name : "ANGKUTAN UDARA TIDAK BERJADWAL UNTUK PENUMPANG",
			
		},
		{
			SectorEconomy3Id : 360,
			Code : "511009",
			Name : "ANGKUTAN UDARA UNTUK PENUMPANG LAINNYA",
			
		},
		{
			SectorEconomy3Id : 361,
			Code : "512000",
			Name : "ANGKUTAN UDARA UNTUK BARANG",
			
		},
		{
			SectorEconomy3Id : 362,
			Code : "521000",
			Name : "PERGUDANGAN DAN PENYIMPANAN",
			
		},
		{
			SectorEconomy3Id : 363,
			Code : "522000",
			Name : "AKTIVITAS PENUNJANG ANGKUTAN",
			
		},
		{
			SectorEconomy3Id : 364,
			Code : "530000",
			Name : "AKTIVITAS POS DAN KURIR",
			
		},
		{
			SectorEconomy3Id : 365,
			Code : "551100",
			Name : "HOTEL BINTANG",
			
		},
		{
			SectorEconomy3Id : 366,
			Code : "551200",
			Name : "HOTEL MELATI",
			
		},
		{
			SectorEconomy3Id : 367,
			Code : "559000",
			Name : "PENYEDIAAN AKOMODASI LAINNYA",
			
		},
		{
			SectorEconomy3Id : 368,
			Code : "561001",
			Name : "RESTORAN DAN RUMAH MAKAN",
			
		},
		{
			SectorEconomy3Id : 369,
			Code : "561009",
			Name : "PENYEDIAAN MAKANAN DAN MINUMAN LAINNYA",
			
		},
		{
			SectorEconomy3Id : 370,
			Code : "580000",
			Name : "AKTIVITAS PENERBITAN",
			
		},
		{
			SectorEconomy3Id : 371,
			Code : "591000",
			Name : "AKTIVITAS PRODUKSI GAMBAR BERGERAK, VIDEO DAN PROGRAM TELEVISI",
			
		},
		{
			SectorEconomy3Id : 372,
			Code : "592000",
			Name : "AKTIVITAS PEREKAMAN SUARA DAN PENERBITAN MUSIK",
			
		},
		{
			SectorEconomy3Id : 373,
			Code : "600000",
			Name : "AKTIVITAS PENYIARAN DAN PEMROGRAMAN",
			
		},
		{
			SectorEconomy3Id : 374,
			Code : "610001",
			Name : "AKTIVITAS TELEKOMUNIKASI DENGAN KABEL, TANPA KABEL DAN SATELIT",
			
		},
		{
			SectorEconomy3Id : 375,
			Code : "610002",
			Name : "JASA NILAI TAMBAH TELEPONI DAN JASA MULTIMEDIA",
			
		},
		{
			SectorEconomy3Id : 376,
			Code : "610009",
			Name : "AKTIVITAS TELEKOMUNIKASI LAINNYA YTDL",
			
		},
		{
			SectorEconomy3Id : 377,
			Code : "620100",
			Name : "AKTIVITAS PEMROGRAMAN KOMPUTER",
			
		},
		{
			SectorEconomy3Id : 378,
			Code : "620200",
			Name : "AKTIVITAS KONSULTASI KOMPUTER DAN MANAJEMEN FASILITAS KOMPUTER",
			
		},
		{
			SectorEconomy3Id : 379,
			Code : "631110",
			Name : "AKTIVITAS PENGOLAHAN DATA",
			
		},
		{
			SectorEconomy3Id : 380,
			Code : "631120",
			Name : "AKTIVITAS HOSTING DAN YBDI",
			
		},
		{
			SectorEconomy3Id : 381,
			Code : "631210",
			Name : "PORTAL WEB DAN/ATAU PLATFORM DIGITAL TANPA TUJUAN KOMERSIAL",
			
		},
		{
			SectorEconomy3Id : 382,
			Code : "631220",
			Name : "PORTAL WEB DAN/ATAU PLATFORM DIGITAL DENGAN TUJUAN KOMERSIAL",
			
		},
		{
			SectorEconomy3Id : 383,
			Code : "639100",
			Name : "AKTIVITAS KANTOR BERITA",
			
		},
		{
			SectorEconomy3Id : 384,
			Code : "639900",
			Name : "AKTIVITAS JASA INFORMASI LAINNYA YTDL",
			
		},
		{
			SectorEconomy3Id : 385,
			Code : "641000",
			Name : "PERANTARA MONETER",
			
		},
		{
			SectorEconomy3Id : 386,
			Code : "649100",
			Name : "SEWA GUNA USAHA DENGAN HAK OPSI",
			
		},
		{
			SectorEconomy3Id : 387,
			Code : "649900",
			Name : "AKTIVITAS JASA KEUANGAN LAINNYA YTDL, BUKAN ASURANSI DAN DANA PENSIUN",
			
		},
		{
			SectorEconomy3Id : 388,
			Code : "650000",
			Name : "ASURANSI, REASURANSI DAN DANA PENSIUN, BUKAN JAMINAN SOSIAL WAJIB",
			
		},
		{
			SectorEconomy3Id : 389,
			Code : "661001",
			Name : "KEGIATAN PENUKARAN VALUTA ASING (MONEY CHANGER)",
			
		},
		{
			SectorEconomy3Id : 390,
			Code : "661009",
			Name : "AKTIVITAS PENUNJANG JASA KEUANGAN LAINNYA",
			
		},
		{
			SectorEconomy3Id : 391,
			Code : "662000",
			Name : "AKTIVITAS PENUNJANG ASURANSI DAN DANA PENSIUN",
			
		},
		{
			SectorEconomy3Id : 392,
			Code : "681101",
			Name : "REAL ESTATE PERUMAHAN SEDERHANA PERUMNAS",
			
		},
		{
			SectorEconomy3Id : 393,
			Code : "681102",
			Name : "REAL ESTATE PERUMAHAN SEDERHANA PERUMNAS TIPE 21",
			
		},
		{
			SectorEconomy3Id : 394,
			Code : "681103",
			Name : "REAL ESTATE PERUMAHAN SEDERHANA PERUMNAS TIPE 22 S.D. 70",
			
		},
		{
			SectorEconomy3Id : 395,
			Code : "681104",
			Name : "REAL ESTATE PERUMAHAN MENENGAH, BESAR ATAU MEWAH (TIPE DIATAS 70)",
			
		},
		{
			SectorEconomy3Id : 396,
			Code : "681105",
			Name : "REAL ESTATE PERUMAHAN FLAT / APARTEMEN",
			
		},
		{
			SectorEconomy3Id : 397,
			Code : "681106",
			Name : "REAL ESTATE GEDUNG PERBELANJAAN (MAL, PLAZA)",
			
		},
		{
			SectorEconomy3Id : 398,
			Code : "681107",
			Name : "REAL ESTATE GEDUNG PERKANTORAN",
			
		},
		{
			SectorEconomy3Id : 399,
			Code : "681108",
			Name : "REAL ESTATE GEDUNG RUMAH TOKO (RUKO) ATAU RUMAH KANTOR (RUKAN)",
			
		},
		{
			SectorEconomy3Id : 400,
			Code : "681109",
			Name : "REAL ESTATE LAINNYA",
			
		},
		{
			SectorEconomy3Id : 401,
			Code : "681200",
			Name : "KAWASAN PARIWISATA",
			
		},
		{
			SectorEconomy3Id : 402,
			Code : "681300",
			Name : "KAWASAN INDUSTRI",
			
		},
		{
			SectorEconomy3Id : 403,
			Code : "682000",
			Name : "REAL ESTAT ATAS DASAR BALAS JASA (FEE) ATAU KONTRAK",
			
		},
		{
			SectorEconomy3Id : 404,
			Code : "690000",
			Name : "AKTIVITAS HUKUM DAN AKUNTANSI",
			
		},
		{
			SectorEconomy3Id : 405,
			Code : "702010",
			Name : "AKTIVITAS KONSULTASI PARIWISATA",
			
		},
		{
			SectorEconomy3Id : 406,
			Code : "702090",
			Name : "AKTIVITAS KANTOR PUSAT DAN KONSULTASI MANAJEMEN LAINNYA",
			
		},
		{
			SectorEconomy3Id : 407,
			Code : "710000",
			Name : "AKTIVITAS ARSITEKTUR DAN KEINSINYURAN; ANALISIS DAN UJI TEKNIS",
			
		},
		{
			SectorEconomy3Id : 408,
			Code : "721000",
			Name : "PENELITIAN DAN PENGEMBANGAN ILMU PENGETAHUAN ALAM DAN ILMU TEKNOLOGI DAN REKAYASA",
			
		},
		{
			SectorEconomy3Id : 409,
			Code : "722000",
			Name : "PENELITIAN DAN PENGEMBANGAN ILMU PENGETAHUAN SOSIAL DAN HUMANIORA",
			
		},
		{
			SectorEconomy3Id : 410,
			Code : "730000",
			Name : "PERIKLANAN DAN PENELITIAN PASAR",
			
		},
		{
			SectorEconomy3Id : 411,
			Code : "740000",
			Name : "AKTIVITAS PROFESIONAL, ILMIAH DAN TEKNIS LAINNYA",
			
		},
		{
			SectorEconomy3Id : 412,
			Code : "750000",
			Name : "AKTIVITAS KESEHATAN HEWAN",
			
		},
		{
			SectorEconomy3Id : 413,
			Code : "771000",
			Name : "AKTIVITAS PENYEWAAN DAN SEWA GUNA USAHA TANPA HAK OPSI MOBIL, BUS, TRUK DAN SEJENISNYA",
			
		},
		{
			SectorEconomy3Id : 414,
			Code : "772000",
			Name : "AKTIVITAS PENYEWAAN DAN SEWA GUNA USAHA TANPA HAK OPSI BARANG PRIBADI DAN RUMAH TANGGA",
			
		},
		{
			SectorEconomy3Id : 415,
			Code : "773020",
			Name : "AKTIVITAS PENYEWAAN DAN SEWA GUNA USAHA TANPA HAK OPSI ALAT TRANSPORTASI DARAT BUKAN KENDARAAN BERMOTOR RODA EMPAT ATAU LEBIH",
			
		},
		{
			SectorEconomy3Id : 416,
			Code : "773030",
			Name : "AKTIVITAS PENYEWAAN DAN SEWA GUNA USAHA TANPA HAK OPSI ALAT TRANSPORTASI AIR",
			
		},
		{
			SectorEconomy3Id : 417,
			Code : "773040",
			Name : "AKTIVITAS PENYEWAAN DAN SEWA GUNA USAHA TANPA HAK OPSI ALAT TRANSPORTASI UDARA",
			
		},
		{
			SectorEconomy3Id : 418,
			Code : "773050",
			Name : "AKTIVITAS PENYEWAAN DAN SEWA GUNA USAHA TANPA HAK OPSI MESIN PERTANIAN DAN PERALATANNYA",
			
		},
		{
			SectorEconomy3Id : 419,
			Code : "773060",
			Name : "AKTIVITAS PENYEWAAN DAN SEWA GUNA USAHA TANPA HAK OPSI MESIN DAN PERALATAN KONSTRUKSI DAN TEKNIK SIPIL",
			
		},
		{
			SectorEconomy3Id : 420,
			Code : "773070",
			Name : "AKTIVITAS PENYEWAAN DAN SEWA GUNA USAHA TANPA HAK OPSI MESIN KANTOR DAN PERALATANNYA",
			
		},
		{
			SectorEconomy3Id : 421,
			Code : "773090",
			Name : "AKTIVITAS PENYEWAAN DAN SEWA GUNA USAHA TANPA HAK OPSI MESIN, PERALATAN DAN BARANG BERWUJUD LAINNYA YTDL",
			
		},
		{
			SectorEconomy3Id : 422,
			Code : "780000",
			Name : "AKTIVITAS KETENAGAKERJAAN",
			
		},
		{
			SectorEconomy3Id : 423,
			Code : "791110",
			Name : "AKTIVITAS AGEN PERJALANAN WISATA",
			
		},
		{
			SectorEconomy3Id : 424,
			Code : "791120",
			Name : "AKTIVITAS AGEN PERJALANAN BUKAN WISATA",
			
		},
		{
			SectorEconomy3Id : 425,
			Code : "791200",
			Name : "AKTIVITAS BIRO PERJALANAN WISATA",
			
		},
		{
			SectorEconomy3Id : 426,
			Code : "799000",
			Name : "JASA RESERVASI LAINNYA DAN KEGIATAN YBDI",
			
		},
		{
			SectorEconomy3Id : 427,
			Code : "823000",
			Name : "PENYELENGGARA KONVENSI DAN PAMERAN DAGANG",
			
		},
		{
			SectorEconomy3Id : 428,
			Code : "829000",
			Name : "AKTIVITAS JASA PENUNJANG USAHA YTDL",
			
		},
		{
			SectorEconomy3Id : 429,
			Code : "841000",
			Name : "ADMINISTRASI PEMERINTAHAN DAN KEBIJAKAN EKONOMI DAN SOSIAL",
			
		},
		{
			SectorEconomy3Id : 430,
			Code : "842000",
			Name : "PENYEDIAAN LAYANAN UNTUK MASYARAKAT DALAM BIDANG HUBUNGAN LUAR NEGERI, PERTAHANAN, KEAMANAN DAN KETERTIBAN",
			
		},
		{
			SectorEconomy3Id : 431,
			Code : "843000",
			Name : "JAMINAN SOSIAL WAJIB",
			
		},
		{
			SectorEconomy3Id : 432,
			Code : "851000",
			Name : "PENDIDIKAN DASAR DAN PENDIDIKAN ANAK USIA DINI",
			
		},
		{
			SectorEconomy3Id : 433,
			Code : "852000",
			Name : "PENDIDIKAN MENENGAH",
			
		},
		{
			SectorEconomy3Id : 434,
			Code : "853000",
			Name : "PENDIDIKAN TINGGI",
			
		},
		{
			SectorEconomy3Id : 435,
			Code : "854000",
			Name : "PENDIDIKAN LAINNYA",
			
		},
		{
			SectorEconomy3Id : 436,
			Code : "855000",
			Name : "KEGIATAN PENUNJANG PENDIDIKAN",
			
		},
		{
			SectorEconomy3Id : 437,
			Code : "861000",
			Name : "AKTIVITAS RUMAH SAKIT",
			
		},
		{
			SectorEconomy3Id : 438,
			Code : "862000",
			Name : "AKTIVITAS PRAKTIK DOKTER DAN DOKTER GIGI",
			
		},
		{
			SectorEconomy3Id : 439,
			Code : "869000",
			Name : "AKTIVITAS PELAYANAN KESEHATAN MANUSIA LAINNYA",
			
		},
		{
			SectorEconomy3Id : 440,
			Code : "870000",
			Name : "AKTIVITAS SOSIAL",
			
		},
		{
			SectorEconomy3Id : 441,
			Code : "900001",
			Name : "JASA IMPRESARIAT BIDANG SENI",
			
		},
		{
			SectorEconomy3Id : 442,
			Code : "900009",
			Name : "AKTIVITAS HIBURAN, SENI DAN KREATIVITAS LAINNYA",
			
		},
		{
			SectorEconomy3Id : 443,
			Code : "910100",
			Name : "PERPUSTAKAAN DAN ARSIP",
			
		},
		{
			SectorEconomy3Id : 444,
			Code : "910200",
			Name : "MUSEUM DAN OPERASIONAL BANGUNAN DAN SITUS BERSEJARAH",
			
		},
		{
			SectorEconomy3Id : 445,
			Code : "930000",
			Name : "AKTIVITAS OLAHRAGA DAN REKREASI LAINNYA",
			
		},
		{
			SectorEconomy3Id : 446,
			Code : "941000",
			Name : "AKTIVITAS ORGANISASI BISNIS, PENGUSAHA DAN PROFESI",
			
		},
		{
			SectorEconomy3Id : 447,
			Code : "942000",
			Name : "AKTIVITAS ORGANISASI BURUH",
			
		},
		{
			SectorEconomy3Id : 448,
			Code : "949000",
			Name : "AKTIVITAS ORGANISASI KEANGGOTAAN LAINNYA YTDL",
			
		},
		{
			SectorEconomy3Id : 449,
			Code : "950000",
			Name : "REPARASI KOMPUTER DAN BARANG KEPERLUAN PRIBADI DAN PERLENGKAPAN RUMAH TANGGA",
			
		},
		{
			SectorEconomy3Id : 450,
			Code : "960001",
			Name : "AKTIVITAS PANTI PIJAT DAN SPA",
			
		},
		{
			SectorEconomy3Id : 451,
			Code : "960009",
			Name : "AKTIVITAS JASA PERORANGAN LAINNYA",
			
		},
		{
			SectorEconomy3Id : 452,
			Code : "970000",
			Name : "AKTIVITAS RUMAH TANGGA SEBAGAI PEMBERI KERJA DARI PERSONIL DOMESTIK",
			
		},
		{
			SectorEconomy3Id : 453,
			Code : "990000",
			Name : "AKTIVITAS BADAN INTERNASIONAL DAN BADAN EKSTRA INTERNASIONAL LAINNYA",
			
		},
	}

	var LokasiPabrik = []models.LokasiPabrik{
		{
			Code : "0000 ",
			Name : "Lainnya Tidak Terdefinisi",
		},
		{
			Code : "0100rn ",
			Name : "Provinsi Jawa Baratrn",
		},
		{
			Code : "0102 ",
			Name : "Kab. Bekasi",
		},
		{
			Code : "0103 ",
			Name : "Kab. Purwakarta",
		},
		{
			Code : "0106 ",
			Name : "Kab. Karawang",
		},
		{
			Code : "0108 ",
			Name : "Kab. Bogor",
		},
		{
			Code : "0109 ",
			Name : "Kab. Sukabumi",
		},
		{
			Code : "0110 ",
			Name : "Kab. Cianjur",
		},
		{
			Code : "0111 ",
			Name : "Kab.Bandung",
		},
		{
			Code : "0112 ",
			Name : "Kab.Sumedang",
		},
		{
			Code : "0113 ",
			Name : "Kab.Tasikmalaya",
		},
		{
			Code : "0114 ",
			Name : "Kab.Garut",
		},
		{
			Code : "0115 ",
			Name : "Kab.Ciamis",
		},
		{
			Code : "0116 ",
			Name : "Kab.Cirebon",
		},
		{
			Code : "0117 ",
			Name : "Kab.Kuningan",
		},
		{
			Code : "0118 ",
			Name : "Kab.Indramayu",
		},
		{
			Code : "0119 ",
			Name : "Kab.Majalengka",
		},
		{
			Code : "0121 ",
			Name : "Kab.Subang",
		},
		{
			Code : "0122 ",
			Name : "Kab.Bandung Barat",
		},
		{
			Code : "0123 ",
			Name : "Kab.Pangandaran ",
		},
		{
			Code : "0180 ",
			Name : "Kota Banjar ",
		},
		{
			Code : "0191 ",
			Name : "Kota Bandung ",
		},
		{
			Code : "0192 ",
			Name : "Kota Bogor ",
		},
		{
			Code : "0193 ",
			Name : "Kota Sukabumi ",
		},
		{
			Code : "0194 ",
			Name : "Kota Cirebon ",
		},
		{
			Code : "0195 ",
			Name : "Kota Tasikmalaya ",
		},
		{
			Code : "0196 ",
			Name : "Kota Cimahi ",
		},
		{
			Code : "0197 ",
			Name : "Kota Depok ",
		},
		{
			Code : "0198 ",
			Name : "Kota Bekasi ",
		},
		{
			Code : "0200 ",
			Name : "Provinsi Banten",
		},
		{
			Code : "0201 ",
			Name : "Kab.Lebak",
		},
		{
			Code : "0202 ",
			Name : "Kab.Pandeglang",
		},
		{
			Code : "0203 ",
			Name : "Kab.Serang",
		},
		{
			Code : "0204 ",
			Name : "Kab.Tangerang",
		},
		{
			Code : "0291 ",
			Name : "Kota Cilegon",
		},
		{
			Code : "0292 ",
			Name : "Kota Tangerang",
		},
		{
			Code : "0293 ",
			Name : "Kota Serang",
		},
		{
			Code : "0294 ",
			Name : "Kota Tangerang Selatan",
		},
		{
			Code : "0300 ",
			Name : "Provinsi DKI Jakarta Raya",
		},
		{
			Code : "0391 ",
			Name : "Wil.Kota Jakarta Pusat",
		},
		{
			Code : "0392 ",
			Name : "Wil.Kota Jakarta Utara",
		},
		{
			Code : "0393",
			Name : "Wil.Kota Jakarta Barat",
		},
		{
			Code : "0394",
			Name : "Wil.Kota Jakarta Selatan",
		},
		{
			Code : "0395",
			Name : "Wil.Kota Jakarta Timur",
		},
		{
			Code : "0396",
			Name : "Wil.Kab.Administrasi Kepulauan Seribu",
		},
		{
			Code : "0500",
			Name : "Daerah Istimewa Yogyakarta",
		},
		{
			Code : "0501",
			Name : "Kab.Bantul",
		},
		{
			Code : "0502",
			Name : "Kab.Sleman",
		},
		{
			Code : "0503",
			Name : "Kab.Gunung Kidul",
		},
		{
			Code : "0504",
			Name : "Kab.Kulon Progo",
		},
		{
			Code : "0591",
			Name : "Kota Yogyakarta",
		},
		{
			Code : "0900",
			Name : "Provinsi Jawa Tengah",
		},
		{
			Code : "0901",
			Name : "Kab.Semarang",
		},
		{
			Code : "0902",
			Name : "Kab.Kendal",
		},
		{
			Code : "0903",
			Name : "Kab.Demak",
		},
		{
			Code : "0904",
			Name : "Kab.Grobogan",
		},
		{
			Code : "0905",
			Name : "Kab.Pekalongan",
		},
		{
			Code : "0906",
			Name : "Kab.Tegal",
		},
		{
			Code : "0907",
			Name : "Kab.Brebes",
		},
		{
			Code : "0908",
			Name : "Kab.Pati",
		},
		{
			Code : "0909",
			Name : "Kab.Kudus",
		},
		{
			Code : "0910",
			Name : "Kab.Pemalang",
		},
		{
			Code : "0911",
			Name : "Kab.Jepara",
		},
		{
			Code : "0912",
			Name : "Kab.Rembang",
		},
		{
			Code : "0913",
			Name : "Kab.Blora",
		},
		{
			Code : "0914",
			Name : "Kab.Banyumas",
		},
		{
			Code : "0915",
			Name : "Kab.Cilacap",
		},
		{
			Code : "0916",
			Name : "Kab.Purbalingga",
		},
		{
			Code : "0917",
			Name : "Kab.Banjarnegara",
		},
		{
			Code : "0918",
			Name : "Kab.Magelang",
		},
		{
			Code : "0919",
			Name : "Kab.Temanggung",
		},
		{
			Code : "0920",
			Name : "Kab.Wonosobo",
		},
		{
			Code : "0921",
			Name : "Kab.Purworejo",
		},
		{
			Code : "0922",
			Name : "Kab.Kebumen",
		},
		{
			Code : "0923",
			Name : "Kab.Klaten",
		},
		{
			Code : "0924",
			Name : "Kab.Boyolali",
		},
		{
			Code : "0925",
			Name : "Kab.Sragen",
		},
		{
			Code : "0926",
			Name : "Kab.Sukoharjo",
		},
		{
			Code : "0927",
			Name : "Kab.Karanganyar",
		},
		{
			Code : "0928",
			Name : "Kab.Wonogiri",
		},
		{
			Code : "0929",
			Name : "Kab.Batang",
		},
		{
			Code : "0991",
			Name : "Kota Semarang",
		},
		{
			Code : "0992",
			Name : "Kota Salatiga",
		},
		{
			Code : "0993",
			Name : "Kota Pekalongan",
		},
		{
			Code : "0994",
			Name : "Kota Tegal",
		},
		{
			Code : "0995",
			Name : "Kota Magelang",
		},
		{
			Code : "0996",
			Name : "Kota Surakarta/Solo",
		},
		{
			Code : "101",
			Name : "Kab. Tangerang, Jawa Barat",
		},
		{
			Code : "102",
			Name : "Kab. Bekasi, Jawa Barat",
		},
		{
			Code : "103",
			Name : "Kab. Purwakarta, Jawa Barat",
		},
		{
			Code : "104",
			Name : "Kab. Serang, Jawa Barat",
		},
		{
			Code : "105",
			Name : "Kab. Pandeglang, Jawa Barat",
		},
		{
			Code : "106",
			Name : "Kab. Karawang, Jawa Barat",
		},
		{
			Code : "107",
			Name : "Kab. Lebak, Jawa Barat",
		},
		{
			Code : "108",
			Name : "Kab. Bogor, Jawa Barat",
		},
		{
			Code : "109",
			Name : "Kab. Sukabumi, Jawa Barat",
		},
		{
			Code : "110",
			Name : "Kab. Cianjur, Jawa Barat",
		},
		{
			Code : "111",
			Name : "Kab. Bandung, Jawa Barat",
		},
		{
			Code : "112",
			Name : "Kab. Sumedang, Jawa Barat",
		},
		{
			Code : "113",
			Name : "Kab. Tasikmalaya, Jawa Barat",
		},
		{
			Code : "114",
			Name : "Kab. Garut, Jawa Barat",
		},
		{
			Code : "115",
			Name : "Kab. Ciamis, Jawa Barat",
		},
		{
			Code : "116",
			Name : "Kab. Cirebon, Jawa Barat",
		},
		{
			Code : "117",
			Name : "Kab. Kuningan, Jawa Barat",
		},
		{
			Code : "118",
			Name : "Kab. Indramayu, Jawa Barat",
		},
		{
			Code : "119",
			Name : "Kab. Majalengka, Jawa Barat",
		},
		{
			Code : "1200",
			Name : "Provinsi Jawa Timur",
		},
		{
			Code : "1201",
			Name : "Kab. Gresik, Jawa Timur",
		},
		{
			Code : "1202",
			Name : "Kab. Sidoarjo, Jawa Timur",
		},
		{
			Code : "1203",
			Name : "Kab. Mojokerto, Jawa Timur",
		},
		{
			Code : "1204",
			Name : "Kab. Jombang, Jawa Timur",
		},
		{
			Code : "1205",
			Name : "Kab. Sampang, Jawa Timur",
		},
		{
			Code : "1206",
			Name : "Kab. Pemekasan, Jawa Timur",
		},
		{
			Code : "1207",
			Name : "Kab. Sumenep, Jawa Timur",
		},
		{
			Code : "1208",
			Name : "Kab. Bangkalan, Jawa Timur",
		},
		{
			Code : "1209",
			Name : "Kab. Bondowoso, Jawa Timur",
		},
		{
			Code : "121",
			Name : "Kab. Subang, Jawa Barat",
		},
		{
			Code : "1211",
			Name : "Kab. Banyuwangi, Jawa Timur",
		},
		{
			Code : "1212",
			Name : "Kab. Jember, Jawa Timur",
		},
		{
			Code : "1213",
			Name : "Kab. Malang, Jawa Timur",
		},
		{
			Code : "1214",
			Name : "Kab. Pasuruan, Jawa Timur",
		},
		{
			Code : "1215",
			Name : "Kab. Probolinggo, Jawa Timur",
		},
		{
			Code : "1216",
			Name : "Kab. Lumajang, Jawa Timur",
		},
		{
			Code : "1217",
			Name : "Kab. Kediri, Jawa Timur",
		},
		{
			Code : "1218",
			Name : "Kab. Nganjuk, Jawa Timur",
		},
		{
			Code : "1219",
			Name : "Kab. Tulungagung, Jawa Timur",
		},
		{
			Code : "1220",
			Name : "Kab. Trenggalek, Jawa Timur",
		},
		{
			Code : "1221",
			Name : "Kab. Blitar, Jawa Timur",
		},
		{
			Code : "1222",
			Name : "Kab. Madiun, Jawa Timur",
		},
		{
			Code : "1223",
			Name : "Kab. Ngawi, Jawa Timur",
		},
		{
			Code : "1224",
			Name : "Kab. Magetan, Jawa Timur",
		},
		{
			Code : "1225",
			Name : "Kab. Ponorogo, Jawa Timur",
		},
		{
			Code : "1226",
			Name : "Kab. Pacitan, Jawa Timur",
		},
		{
			Code : "1227",
			Name : "Kab. Bojonegoro, Jawa Timur",
		},
		{
			Code : "1228",
			Name : "Kab. Tuban, Jawa Timur",
		},
		{
			Code : "1229",
			Name : "Kab. Lamongan, Jawa Timur",
		},
		{
			Code : "1230",
			Name : "Kab. Situbondo, Jawa Timur",
		},
		{
			Code : "1271",
			Name : "Kota Batu",
		},
		{
			Code : "1291",
			Name : "Kod. Surabaya, Jawa Timur",
		},
		{
			Code : "1292",
			Name : "Kod. Mojokerto, Jawa Timur",
		},
		{
			Code : "1293",
			Name : "Kod. Malang, Jawa Timur",
		},
		{
			Code : "1294",
			Name : "Kod. Pasuruan, Jawa Timur",
		},
		{
			Code : "1295",
			Name : "Kod. Probolinggo, Jawa Timur",
		},
		{
			Code : "1296",
			Name : "Kod. Blitar, Jawa Timur",
		},
		{
			Code : "1297",
			Name : "Kod. Kediri, Jawa Timur",
		},
		{
			Code : "1298",
			Name : "Kod. Madiun, Jawa Timur",
		},
		{
			Code : "1299",
			Name : "Kotif Jember, Jawa Timur",
		},
		{
			Code : "180",
			Name : "Kotif Banjar, Jawa Barat",
		},
		{
			Code : "190",
			Name : "Kotif Cilegon, Jawa Barat",
		},
		{
			Code : "191",
			Name : "Kod. Bandung, Jawa Barat",
		},
		{
			Code : "192",
			Name : "Kod. Bogor, Jawa Barat",
		},
		{
			Code : "193",
			Name : "Kod. Sukabumi, Jawa Barat",
		},
		{
			Code : "194",
			Name : "Kod. Cirebon, Jawa Barat",
		},
		{
			Code : "195",
			Name : "Kotif Tasikmalaya, Jawa Barat",
		},
		{
			Code : "196",
			Name : "Kotif Cimahi, Jawa Barat",
		},
		{
			Code : "197",
			Name : "Kotif Depok, Jawa Barat",
		},
		{
			Code : "198",
			Name : "Kotif Bekasi, Jawa Barat",
		},
		{
			Code : "199",
			Name : "Kotif Tangerang, Jawa Barat",
		},
		{
			Code : "2300",
			Name : "Provinsi Bengkulu",
		},
		{
			Code : "2301",
			Name : "Kab. Bengkulu Selatan, Bengkul",
		},
		{
			Code : "2302",
			Name : "Kab. Bengkulu Utara, Bengkulu",
		},
		{
			Code : "2303",
			Name : "Kab. Rejang Lebong, Bengkulu",
		},
		{
			Code : "2304",
			Name : "Kab. Lebong",
		},
		{
			Code : "2305",
			Name : "Kab. Kepahiang",
		},
		{
			Code : "2306",
			Name : "Kab.Mukomuko",
		},
		{
			Code : "2307",
			Name : "Kab.Seluma",
		},
		{
			Code : "2308",
			Name : "Kab.Kaur",
		},
		{
			Code : "2309",
			Name : "Kab.Bengkulu Tengah",
		},
		{
			Code : "2391",
			Name : "Kod. Bengkulu, Bengkulu",
		},
		{
			Code : "3100",
			Name : "Provinsi Jambi",
		},
		{
			Code : "3101",
			Name : "Kab. Batanghari, Jambi",
		},
		{
			Code : "3102",
			Name : "Kab. Tanjung Jabung, Jambi",
		},
		{
			Code : "3103",
			Name : "Kab. Muara Bungo-Tebo, Jambi",
		},
		{
			Code : "3104",
			Name : "Kab. Sarolangun / Bang, Jambi",
		},
		{
			Code : "3105",
			Name : "Kab. Kerinci, Jambi",
		},
		{
			Code : "3106",
			Name : "Kab. Muara Jambi, Jambi",
		},
		{
			Code : "3107",
			Name : "Kab. Tanjung Jabung Barat, Jam",
		},
		{
			Code : "3108",
			Name : "Kab. Tanjung Jabung Timur, Jam",
		},
		{
			Code : "3109",
			Name : "Kab. Tebo, Jambi",
		},
		{
			Code : "3110",
			Name : "Kab. Bungo, Jambi",
		},
		{
			Code : "3111",
			Name : "Kab. Merangin, Jambi",
		},
		{
			Code : "3112",
			Name : "Kab.Bungo",
		},
		{
			Code : "3191",
			Name : "Kod. Jambi, Jambi",
		},
		{
			Code : "3192",
			Name : "Kota Sungai Penuh",
		},
		{
			Code : "3200",
			Name : "Provinsi Nanggroe Aceh Darussalam",
		},
		{
			Code : "3201",
			Name : "Kab. Aceh Besar, Aceh",
		},
		{
			Code : "3202",
			Name : "Kab. Pidie, Aceh",
		},
		{
			Code : "3203",
			Name : "Kab. Aceh Utara, Aceh",
		},
		{
			Code : "3204",
			Name : "Kab. Aceh Timur, Aceh",
		},
		{
			Code : "3205",
			Name : "Kab. Aceh Selatan, Aceh",
		},
		{
			Code : "3206",
			Name : "Kab. Aceh Barat, Aceh",
		},
		{
			Code : "3207",
			Name : "Kab. Aceh Tengah, Aceh",
		},
		{
			Code : "3208",
			Name : "Kab. Aceh Tenggara, Aceh",
		},
		{
			Code : "3209",
			Name : "Kab. Aceh Singkil, Aceh",
		},
		{
			Code : "3210",
			Name : "Kab.Aceh Jeumpa/Bireuen",
		},
		{
			Code : "3211",
			Name : "Kab. Aceh Tamiang",
		},
		{
			Code : "3212",
			Name : "Kab. Gayo Luwes",
		},
		{
			Code : "3213",
			Name : "Kab. Aceh Barat Daya",
		},
		{
			Code : "3214",
			Name : "Kab.Aceh Jaya",
		},
		{
			Code : "3215",
			Name : "Kab.Nagan Raya",
		},
		{
			Code : "3216",
			Name : "Kab.Simeuleu",
		},
		{
			Code : "3217",
			Name : "Kab.Bener Meriah",
		},
		{
			Code : "3218",
			Name : "Kab.Pidie Jaya",
		},
		{
			Code : "3219",
			Name : "Kab.Subulussalam",
		},
		{
			Code : "3291",
			Name : "Kod. Banda Aceh, Aceh",
		},
		{
			Code : "3292",
			Name : "Kod. Sabang, Aceh",
		},
		{
			Code : "3293",
			Name : "Kotif Lhokseumawe, Aceh",
		},
		{
			Code : "3294",
			Name : "Kotif Langsa, Aceh",
		},
		{
			Code : "3295",
			Name : "Kotif Simeulue, Aceh",
		},
		{
			Code : "3300",
			Name : "Provinsi Sumatera Utara",
		},
		{
			Code : "3301",
			Name : "Kab. Deli Serdang, Sumatera Ut",
		},
		{
			Code : "3302",
			Name : "Kab. Langkat, Sumatera Utara",
		},
		{
			Code : "3303",
			Name : "Kab. Karo, Sumatera Utara",
		},
		{
			Code : "3304",
			Name : "Kab. Simalungun, Sumatera Utar",
		},
		{
			Code : "3305",
			Name : "Kab. Labuhan Batu, Sumatera Ut",
		},
		{
			Code : "3306",
			Name : "Kab. Asahan, Sumatera Utara",
		},
		{
			Code : "3307",
			Name : "Kab. Dairi, Sumatera Utara",
		},
		{
			Code : "3308",
			Name : "Kab. Tapanuli Utara, Sumatera ",
		},
		{
			Code : "3309",
			Name : "Kab. Tapanuli Tengah, Sumatera",
		},
		{
			Code : "3310",
			Name : "Kab. Tapanuli Selatan, Sumater",
		},
		{
			Code : "3311",
			Name : "Kab. Nias, Sumatera Utara",
		},
		{
			Code : "3312",
			Name : "Kotif. Rantau Prapat, Sumatera",
		},
		{
			Code : "3313",
			Name : "Kab. Toba Samosir, Sumatera Ut",
		},
		{
			Code : "3314",
			Name : "Kab. Mandailing Natal, Sumater",
		},
		{
			Code : "3315",
			Name : "Kab.Nias Selatan",
		},
		{
			Code : "3316",
			Name : "Kab.Humbang Hasundutan",
		},
		{
			Code : "3317",
			Name : "Kab.Pakpak Bharat",
		},
		{
			Code : "3318",
			Name : "Kab.Samosir",
		},
		{
			Code : "3319",
			Name : "Kab.Serdang Bedagai",
		},
		{
			Code : "3321",
			Name : "Kab.Batu Bara",
		},
		{
			Code : "3322",
			Name : "Kab.Padang Lawas",
		},
		{
			Code : "3323",
			Name : "Kab.Padang Lawas Utara",
		},
		{
			Code : "3324",
			Name : "Kab.Labuanbatu Selatan",
		},
		{
			Code : "3325",
			Name : "Kab.Labuanbatu Utara",
		},
		{
			Code : "3326",
			Name : "Kab.Nias Barat",
		},
		{
			Code : "3327",
			Name : "Kab.Nias Utara",
		},
		{
			Code : "3391",
			Name : "Kod. Tebing Tinggi, Sumatera U",
		},
		{
			Code : "3392",
			Name : "Kod. Binjai, Sumatera Utara",
		},
		{
			Code : "3393",
			Name : "Kod. Pematangsiantar, Sumatera",
		},
		{
			Code : "3394",
			Name : "Kod. Tanjung Balai, Sumatera U",
		},
		{
			Code : "3395",
			Name : "Kod. Sibolga, Sumatera Utara",
		},
		{
			Code : "3396",
			Name : "Kod. Medan, Sumatera Utara",
		},
		{
			Code : "3397",
			Name : "Kota Gunung Sitoli",
		},
		{
			Code : "3398",
			Name : "Kotif Kisaran, Sumatera Utara",
		},
		{
			Code : "3399",
			Name : "Kotif Padang Sidempuan, Sumate",
		},
		{
			Code : "3400",
			Name : "Provinsi Sumatera Barat",
		},
		{
			Code : "3401",
			Name : "Kab. Agam, Sumatera Barat",
		},
		{
			Code : "3402",
			Name : "Kab. Pasaman, Sumatera Barat",
		},
		{
			Code : "3403",
			Name : "Kab. Limapuluh Kota, Sumatera ",
		},
		{
			Code : "3404",
			Name : "Kab. Solok, Sumatera Barat",
		},
		{
			Code : "3405",
			Name : "Kab. Padang/Pariaman, Sumatera",
		},
		{
			Code : "3406",
			Name : "Kab. Pesisir Selatan, Sumatera",
		},
		{
			Code : "3407",
			Name : "Kab. Tanah Datar, Sumatera Bar",
		},
		{
			Code : "3408",
			Name : "Kab. Sawahlunto/Sijunj, Sumate",
		},
		{
			Code : "3409",
			Name : "Kab. Mentawai, Sumatera Barat",
		},
		{
			Code : "3410",
			Name : "Kab.Pasaman Barat",
		},
		{
			Code : "3412",
			Name : "Kab.Solok",
		},
		{
			Code : "3491",
			Name : "Kod. Bukittinggi, Sumatera Bar",
		},
		{
			Code : "3492",
			Name : "Kod. Padang, Sumatera Barat",
		},
		{
			Code : "3493",
			Name : "Kod. Sawahlunto, Sumatera Bara",
		},
		{
			Code : "3494",
			Name : "Kod. Padangpanjang, Sumatera B",
		},
		{
			Code : "3495",
			Name : "Kod. Solok, Sumatera Barat",
		},
		{
			Code : "3496",
			Name : "Kod. Payakumbuh, Sumatera Bara",
		},
		{
			Code : "3497",
			Name : "Kotif Pariaman, Sumatera Barat",
		},
		{
			Code : "3500",
			Name : "Provinsi Riau",
		},
		{
			Code : "3501",
			Name : "Kab. Kampar, Riau",
		},
		{
			Code : "3502",
			Name : "Kab. Bengkalis, Riau",
		},
		{
			Code : "3503",
			Name : "Kab. Riau Kepulauan, Riau",
		},
		{
			Code : "3504",
			Name : "Kab. Indragiri Hulu, Riau",
		},
		{
			Code : "3505",
			Name : "Kab. Indragiri Hilir, Riau",
		},
		{
			Code : "3506",
			Name : "Kab. Karimun, Riau",
		},
		{
			Code : "3507",
			Name : "Kab. Natuna, Riau",
		},
		{
			Code : "3508",
			Name : "Kab. Rokan Hulu, Riau",
		},
		{
			Code : "3509",
			Name : "Kab. Rokan Hilir, Riau",
		},
		{
			Code : "3510",
			Name : "Kab. Pelalawan, Riau",
		},
		{
			Code : "3511",
			Name : "Kab. Siak, Riau",
		},
		{
			Code : "3512",
			Name : "Kab. Kuantan Singingi",
		},
		{
			Code : "3513",
			Name : "Kab.Kepulauan Meranti",
		},
		{
			Code : "3591",
			Name : "Kod. Pekanbaru, Riau",
		},
		{
			Code : "3592",
			Name : "Kotif Dumai, Riau",
		},
		{
			Code : "3593",
			Name : "Kotif Tanjungpinang, Riau",
		},
		{
			Code : "3594",
			Name : "Kotif Pulau Batam, Riau",
		},
		{
			Code : "3600",
			Name : "Provinsi Sumatera Selatan",
		},
		{
			Code : "3604",
			Name : "Kab. Belitung, Sumatera Selata",
		},
		{
			Code : "3605",
			Name : "Kab. Bangka, Sumatera Selatan",
		},
		{
			Code : "3606",
			Name : "Kab. Musi/Banyuasin, Sumatera ",
		},
		{
			Code : "3607",
			Name : "Kab. Ogan Komering Ulu, Sumate",
		},
		{
			Code : "3608",
			Name : "Kab.Lematang Ilir Ogan, Sumate",
		},
		{
			Code : "3609",
			Name : "Kab. Lahat, Sumatera Selatan",
		},
		{
			Code : "3610",
			Name : "Kab. Musi Rawas, Sumatera Sela",
		},
		{
			Code : "3611",
			Name : "Kab.Ogan Komering Ilir, Sumate",
		},
		{
			Code : "3613",
			Name : "Kab.Banyuasin",
		},
		{
			Code : "3614",
			Name : "Kab.Ogan Komering Ulu Selatan",
		},
		{
			Code : "3615",
			Name : "Kab.Ogan Komering Ulu Timur",
		},
		{
			Code : "3616",
			Name : "Kab.Ogan Ilir",
		},
		{
			Code : "3617",
			Name : "Kab.Empat Lawang",
		},
		{
			Code : "3618",
			Name : "Kab.Musi Rawas Utara",
		},
		{
			Code : "3619",
			Name : "Kab.Penukal Abab Lematang Ilir",
		},
		{
			Code : "3688",
			Name : "Prov.Sumatera Selatan,Kab/Kota Lainnya",
		},
		{
			Code : "3691",
			Name : "Kod. Palembang, Sumatera Selat",
		},
		{
			Code : "3692",
			Name : "Kod. Pangkal Pinang, Sumatera ",
		},
		{
			Code : "3693",
			Name : "Kotif Lubuklinggau, Sumatera S",
		},
		{
			Code : "3694",
			Name : "Kotif Prabumulih, Sumatera Sel",
		},
		{
			Code : "3695",
			Name : "Kotif Baturaja, Sumatera Selat",
		},
		{
			Code : "3697",
			Name : "Kotif Pagar Alam, Sumatera Sel",
		},
		{
			Code : "3700",
			Name : "Provinsi Kep.Bangka Belitung",
		},
		{
			Code : "3701",
			Name : "Kab.Bangka",
		},
		{
			Code : "3702",
			Name : "Kab.Belitung",
		},
		{
			Code : "3703",
			Name : "Kab.Bangka Barat",
		},
		{
			Code : "3704",
			Name : "Kab.Bangka Selatan",
		},
		{
			Code : "3705",
			Name : "Kab.Bangka Tengah",
		},
		{
			Code : "3706",
			Name : "Kab.Belitung Timur",
		},
		{
			Code : "3707",
			Name : "Kota Pangkal Pinang",
		},
		{
			Code : "3800",
			Name : "Provinsi Kep.Riau",
		},
		{
			Code : "3801",
			Name : "Kab.Karimun",
		},
		{
			Code : "3802",
			Name : "Kab.Lingga",
		},
		{
			Code : "3803",
			Name : "Kab.Natuna",
		},
		{
			Code : "3804",
			Name : "Kab.Bintan (d/h Kabupaten Kepulauan Riau)",
		},
		{
			Code : "3805",
			Name : "Kab.Kepulauan Anambas",
		},
		{
			Code : "3891",
			Name : "Kota Tanjung Pinang",
		},
		{
			Code : "3892",
			Name : "Kota Batam",
		},
		{
			Code : "3900",
			Name : "Provinsi Lampung",
		},
		{
			Code : "3901",
			Name : "Kab. Lampung Selatan, Lampung",
		},
		{
			Code : "3902",
			Name : "Kab. Lampung Tengah, Lampung",
		},
		{
			Code : "3903",
			Name : "Kab. Lampung Utara, Lampung",
		},
		{
			Code : "3904",
			Name : "Kab. Lampung Barat, Lampung",
		},
		{
			Code : "3905",
			Name : "Kab. Tulang Bawang, Lampung",
		},
		{
			Code : "3906",
			Name : "Kab. Tanggamus, Lampung",
		},
		{
			Code : "3907",
			Name : "Kab. Lampung Timur, Lampung",
		},
		{
			Code : "3908",
			Name : "Kab. Way Kanan, Lampung",
		},
		{
			Code : "3909",
			Name : "Kab.Pesawaran",
		},
		{
			Code : "391",
			Name : "Wil. Jakarta Pusat, Jakarta",
		},
		{
			Code : "3910",
			Name : "Kab.Pringsewu",
		},
		{
			Code : "3911",
			Name : "Kab.Tulang Bawang Barat",
		},
		{
			Code : "3912",
			Name : "Kab.Mesuji",
		},
		{
			Code : "3913",
			Name : "Kab.Pesisir Barat",
		},
		{
			Code : "392",
			Name : "Wil. Jakarta Utara, Jakarta",
		},
		{
			Code : "393",
			Name : "Wil. Jakarta Barat, Jakarta",
		},
		{
			Code : "394",
			Name : "Wil. Jakarta Selatan, Jakarta",
		},
		{
			Code : "395",
			Name : "Wil. Jakarta Timur, Jakarta",
		},
		{
			Code : "3991",
			Name : "Kod. Bandar Lampung, Lampung",
		},
		{
			Code : "3992",
			Name : "Kotif Metro, Lampung",
		},
		{
			Code : "4411",
			Name : "Kab.Dharmasraya",
		},
		{
			Code : "501",
			Name : "Kab. Bantul, Yogyakarta",
		},
		{
			Code : "502",
			Name : "Kab. Sleman, Yogyakarta",
		},
		{
			Code : "503",
			Name : "Kab. Gunung Kidul, Yogyakarta",
		},
		{
			Code : "504",
			Name : "Kab. Kulon Progo, Yogyakarta",
		},
		{
			Code : "5100",
			Name : "Provinsi Kalimantan Selatan",
		},
		{
			Code : "5101",
			Name : "Kab. Banjar, Kalimantan Selata",
		},
		{
			Code : "5102",
			Name : "Kab. Tanah Laut, Kalimantan Se",
		},
		{
			Code : "5103",
			Name : "Kab. Tapin, Kalimantan Selatan",
		},
		{
			Code : "5104",
			Name : "Kab.Hulu Sungai Selata, Kal-Se",
		},
		{
			Code : "5105",
			Name : "Kab.Hulu Sungai Tengah, Kal-Se",
		},
		{
			Code : "5106",
			Name : "Kab.Hulu Sungai Utara, Kal-Sel",
		},
		{
			Code : "5107",
			Name : "Kab. Barito Kuala, Kalimantan ",
		},
		{
			Code : "5108",
			Name : "Kab. Kota Baru, Kalimantan Sel",
		},
		{
			Code : "5109",
			Name : "Kab. Tobalong, Kalimantan Sela",
		},
		{
			Code : "5110",
			Name : "Kab.Tanah Bumbu",
		},
		{
			Code : "5111",
			Name : "Kab.Balangan",
		},
		{
			Code : "5191",
			Name : "Kod. Banjarmasin, Kalimantan S",
		},
		{
			Code : "5192",
			Name : "Kotif Banjarbaru, Kalimantan S",
		},
		{
			Code : "5300",
			Name : "Provinsi Kalimantan Barat",
		},
		{
			Code : "5301",
			Name : "Kab. Pontianak, Kalimatan Bara",
		},
		{
			Code : "5302",
			Name : "Kab. Sambas, Kalimatan Barat",
		},
		{
			Code : "5303",
			Name : "Kab. Ketapang, Kalimatan Barat",
		},
		{
			Code : "5304",
			Name : "Kab. Sanggau, Kalimatan Barat",
		},
		{
			Code : "5305",
			Name : "Kab. Sintang, Kalimatan Barat",
		},
		{
			Code : "5306",
			Name : "Kab. Kapuas Hulu, Kalimatan Ba",
		},
		{
			Code : "5307",
			Name : "Kab. Bengkayang, Kalimatan Bar",
		},
		{
			Code : "5308",
			Name : "Kab. Landak, Kalimatan Barat",
		},
		{
			Code : "5309",
			Name : "Kab.Sekadau",
		},
		{
			Code : "5310",
			Name : "Kab.Melawi",
		},
		{
			Code : "5311",
			Name : "Kab.Kayong Utara",
		},
		{
			Code : "5312",
			Name : "Kab.Kubu Raya",
		},
		{
			Code : "5391",
			Name : "Kod. Pontianak, Kalimatan Bara",
		},
		{
			Code : "5392",
			Name : "Kotif Singkawang, Kalimatan Ba",
		},
		{
			Code : "5400",
			Name : "Provinsi Kalimantan Timur",
		},
		{
			Code : "5401",
			Name : "Kab. Kutai, Kalimantan Timur",
		},
		{
			Code : "5402",
			Name : "Kab. Berau, Kalimantan Timur",
		},
		{
			Code : "5403",
			Name : "Kab. Tanah Pasir, Kalimantan T",
		},
		{
			Code : "5404",
			Name : "Kab. Bulungan, Kalimantan Timu",
		},
		{
			Code : "5405",
			Name : "Kab. Kutai Barat, Kalimantan T",
		},
		{
			Code : "5406",
			Name : "Kab. Kutai Timur, Kalimantan T",
		},
		{
			Code : "5407",
			Name : "Kab. Malinau, Kalimantan Timur",
		},
		{
			Code : "5408",
			Name : "Kab. Nunukan, Kalimantan Timur",
		},
		{
			Code : "5409",
			Name : "Kab.Nunukan",
		},
		{
			Code : "5410",
			Name : "Kab.Malinau",
		},
		{
			Code : "5411",
			Name : "Kab.Penajam Paser Utara",
		},
		{
			Code : "5412",
			Name : "Kab.Tana Tidung",
		},
		{
			Code : "5413",
			Name : "kab.Mahakam Ulu",
		},
		{
			Code : "5491",
			Name : "Kod. Samarinda, Kalimantan Tim",
		},
		{
			Code : "5492",
			Name : "Kod. Balikpapan, Kalimantan Ti",
		},
		{
			Code : "5493",
			Name : "Kotif Tarakan, Kalimantan Timu",
		},
		{
			Code : "5494",
			Name : "Kotif Bontang, Kalimantan Timu",
		},
		{
			Code : "5500",
			Name : "Provinsi Kalimantan Utara",
		},
		{
			Code : "5800",
			Name : "Provinsi Kalimantan Tengah",
		},
		{
			Code : "5801",
			Name : "Kab. Kapuas, Kalimantan Tengah",
		},
		{
			Code : "5802",
			Name : "Kab.Kotawaringin Barat",
		},
		{
			Code : "5803",
			Name : "Kab. Kotawaringin Barat, Kal-T",
		},
		{
			Code : "5804",
			Name : "Kab.Murung Raya",
		},
		{
			Code : "5805",
			Name : "Kab.Barito Timur",
		},
		{
			Code : "5806",
			Name : "Kab. Barito Selatan, Kal-Teng",
		},
		{
			Code : "5807",
			Name : "Kab.Gunung Mas",
		},
		{
			Code : "5808",
			Name : "Kab. Barito Utara, Kal-Teng",
		},
		{
			Code : "5809",
			Name : "Kab.Pulang Pisau",
		},
		{
			Code : "5810",
			Name : "Kab.Seruyan",
		},
		{
			Code : "5811",
			Name : "Kab.Katingan",
		},
		{
			Code : "5812",
			Name : "Kab.Sukamara",
		},
		{
			Code : "5813",
			Name : "Kab.Lamandu",
		},
		{
			Code : "5892",
			Name : "Kod. Palangkaraya, Kal-Teng",
		},
		{
			Code : "591",
			Name : "Kod. Yogyakarta, Yogyakarta",
		},
		{
			Code : "6000",
			Name : "Provinsi Sulawesi Tengah",
		},
		{
			Code : "6001",
			Name : "Kab. Donggala, Sulawesi Tengah",
		},
		{
			Code : "6002",
			Name : "Kab. Poso, Sulawesi Tengah",
		},
		{
			Code : "6003",
			Name : "Kab. Banggai, Sulawesi Tengah",
		},
		{
			Code : "6004",
			Name : "Kab. Toli-toli, Sulawesi Tenga",
		},
		{
			Code : "6005",
			Name : "Kab.Banggai Kepulauan, Sulawes",
		},
		{
			Code : "6006",
			Name : "Kab. Morowali, Sulawesi Tengah",
		},
		{
			Code : "6007",
			Name : "Kab. Buol, Sulawesi Tengah",
		},
		{
			Code : "6008",
			Name : "Kab.Tojo Una-Una",
		},
		{
			Code : "6009",
			Name : "Kab.Parigi Moutong",
		},
		{
			Code : "6010",
			Name : "Kab.Sigi",
		},
		{
			Code : "6011",
			Name : "Kab.Banggai Laut",
		},
		{
			Code : "6012",
			Name : "Kab.Morowali Utara",
		},
		{
			Code : "6091",
			Name : "Kotif Palu, Sulawesi Tengah",
		},
		{
			Code : "6100",
			Name : "Provinsi Sulawesi Selatan",
		},
		{
			Code : "6101",
			Name : "Kab. Pinrang, Sulawesi Selatan",
		},
		{
			Code : "6102",
			Name : "Kab. Gowa, Sulawesi Selatan",
		},
		{
			Code : "6103",
			Name : "Kab. Wajo, Sulawesi Selatan",
		},
		{
			Code : "6104",
			Name : "Kab. Mamuju, Sulawesi Selatan",
		},
		{
			Code : "6105",
			Name : "Kab. Bone, Sulawesi Selatan",
		},
		{
			Code : "6106",
			Name : "Kab. Tana Toraja, Sulawesi Sel",
		},
		{
			Code : "6107",
			Name : "Kab. Maros, Sulawesi Selatan",
		},
		{
			Code : "6108",
			Name : "Kab. Majene, Sulawesi Selatan",
		},
		{
			Code : "6109",
			Name : "Kab. Luwu, Sulawesi Selatan",
		},
		{
			Code : "6110",
			Name : "Kab. Sinjai, Sulawesi Selatan",
		},
		{
			Code : "6111",
			Name : "Kab. Bulukumba, Sulawesi Selat",
		},
		{
			Code : "6112",
			Name : "Kab. Bantaeng, Sulawesi Selata",
		},
		{
			Code : "6113",
			Name : "Kab. Jeneponto, Sulawesi Selat",
		},
		{
			Code : "6114",
			Name : "Kab. Selayar, Sulawesi Selatan",
		},
		{
			Code : "6115",
			Name : "Kab. Takalar, Sulawesi Selatan",
		},
		{
			Code : "6116",
			Name : "Kab. Barru, Sulawesi Selatan",
		},
		{
			Code : "6117",
			Name : "Kab. Sindenreng Rappan, Sulawe",
		},
		{
			Code : "6118",
			Name : "Kab. Pangkajene Kepula, Sulawe",
		},
		{
			Code : "6119",
			Name : "Kab. Soppeng, Sulawesi Selatan",
		},
		{
			Code : "6120",
			Name : "Kab. Polewali Mamasa, Sulawesi",
		},
		{
			Code : "6121",
			Name : "Kab. Enrekang, Sulawesi Selata",
		},
		{
			Code : "6122",
			Name : "Kab. Luwu Selatan, Sulawesi Se",
		},
		{
			Code : "6124",
			Name : "Kab.Luwu Utara",
		},
		{
			Code : "6125",
			Name : "Kab.Toraja Utara",
		},
		{
			Code : "6191",
			Name : "Kod. Ujungpandang, Sulawesi Se",
		},
		{
			Code : "6192",
			Name : "Kod. Pare-pare, Sulawesi Selat",
		},
		{
			Code : "6193",
			Name : "Kotif Palopo, Sulawesi Selatan",
		},
		{
			Code : "6194",
			Name : "Kotif Watampone, Sulawesi Sela",
		},
		{
			Code : "6200",
			Name : "Provinsi Sulawesi Utara",
		},
		{
			Code : "6201",
			Name : "Kab. Gorontalo, Sulawesi Utara",
		},
		{
			Code : "6202",
			Name : "Kab. Minahasa, Sulawesi Utara",
		},
		{
			Code : "6203",
			Name : "Kab. Bolaang Mongondow, Sul-Ut",
		},
		{
			Code : "6204",
			Name : "Kab. Sangihe Talaud, Sulawesi ",
		},
		{
			Code : "6205",
			Name : "kab. Bitung Sulawesi Utara",
		},
		{
			Code : "6206",
			Name : "Kab. Bualemo, Sulawesi Utara",
		},
		{
			Code : "6207",
			Name : "Kab.Minahasa Utara",
		},
		{
			Code : "6209",
			Name : "Kab.Minahasa Tenggara",
		},
		{
			Code : "6210",
			Name : "Kab.Bolaang Mongondow Utara",
		},
		{
			Code : "6211",
			Name : "Kab.Kepulauan Sitaro",
		},
		{
			Code : "6212",
			Name : "Kab.Bolaang Mongondow Selatan",
		},
		{
			Code : "6213",
			Name : "Kab.Bolaang Mongondow Timur",
		},
		{
			Code : "6291",
			Name : "Kod. Manado, Sulawesi Utara",
		},
		{
			Code : "6292",
			Name : "Kod. Gorontalo, Sulawesi Utara",
		},
		{
			Code : "6293",
			Name : "Kod. Bitung, Sulawesi Utara",
		},
		{
			Code : "6294",
			Name : "Kota Tomohon",
		},
		{
			Code : "6300",
			Name : "Provinsi Gorontalo",
		},
		{
			Code : "6301",
			Name : "Kab.Gorontalo",
		},
		{
			Code : "6303",
			Name : "Kab.Bonebolango",
		},
		{
			Code : "6304",
			Name : "Kab.Pohuwato",
		},
		{
			Code : "6305",
			Name : "Kab.Gorontalo Utara",
		},
		{
			Code : "6391",
			Name : "Kota Gorontalo",
		},
		{
			Code : "6400",
			Name : "Provinsi Sulawesi Barat",
		},
		{
			Code : "6401",
			Name : "Kab.Polewali Mandar",
		},
		{
			Code : "6402",
			Name : "Kab.Majene",
		},
		{
			Code : "6403",
			Name : "Kab.Mamasa",
		},
		{
			Code : "6404",
			Name : "Kab.Mamuju Utara",
		},
		{
			Code : "6405",
			Name : "Kab.Mamuju Tengah",
		},
		{
			Code : "6406",
			Name : "Kab.Mamuju",
		},
		{
			Code : "6900",
			Name : "Provinsi Sulawesi Tenggara",
		},
		{
			Code : "6901",
			Name : "Kab. Buton, Sulawesi Tenggara",
		},
		{
			Code : "6902",
			Name : "Kab. Kendari, Sulawesi Tenggar",
		},
		{
			Code : "6903",
			Name : "Kab. Muna, Sulawesi Tenggara",
		},
		{
			Code : "6904",
			Name : "Kab. Kolaka, Sulawesi Tenggara",
		},
		{
			Code : "6905",
			Name : "Kab.Wakatobi",
		},
		{
			Code : "6906",
			Name : "Kab.Konawe",
		},
		{
			Code : "6907",
			Name : "Kab.Konawe Selatan",
		},
		{
			Code : "6908",
			Name : "Kab.Bombana",
		},
		{
			Code : "6909",
			Name : "Kab.Kolaka Utara",
		},
		{
			Code : "6910",
			Name : "Kab.Buton Utara",
		},
		{
			Code : "6911",
			Name : "Kab.Konawe Utara",
		},
		{
			Code : "6912",
			Name : "Kab.Kaloka Timur",
		},
		{
			Code : "6913",
			Name : "Kab.Konawe Kepulauan",
		},
		{
			Code : "6914",
			Name : "Kab.Buton Selatan",
		},
		{
			Code : "6915",
			Name : "Kab.Buton Tengah",
		},
		{
			Code : "6916",
			Name : "Kab.Muna Barat",
		},
		{
			Code : "6990",
			Name : "Kotif Bau-Bau, Sulawesi Tengga",
		},
		{
			Code : "6991",
			Name : "Kotif Kendari, Sulawesi Tengga",
		},
		{
			Code : "7100",
			Name : "Provinsi Nusa Tenggara Barat",
		},
		{
			Code : "7101",
			Name : "Kab. Lombok Barat, NTB",
		},
		{
			Code : "7102",
			Name : "Kab. Lombok Tengah, NTB",
		},
		{
			Code : "7103",
			Name : "Kab. Lombok Timur, NTB",
		},
		{
			Code : "7104",
			Name : "Kab. Sumbawa, NTB",
		},
		{
			Code : "7105",
			Name : "Kab. Bima, NTB",
		},
		{
			Code : "7106",
			Name : "Kab. Dompu, NTB",
		},
		{
			Code : "7107",
			Name : "Kab.Sumbawa Barat",
		},
		{
			Code : "7108",
			Name : "Kab.Lombok Utara",
		},
		{
			Code : "7191",
			Name : "Kotif Mataram, NTB",
		},
		{
			Code : "7192",
			Name : "Kota Bima",
		},
		{
			Code : "7200",
			Name : "Provinsi Bali",
		},
		{
			Code : "7201",
			Name : "Kab. Buleleng, Bali",
		},
		{
			Code : "7202",
			Name : "Kab. Jembrana, Bali",
		},
		{
			Code : "7203",
			Name : "Kab. Tabanan, Bali",
		},
		{
			Code : "7204",
			Name : "Kab. Badung, Bali",
		},
		{
			Code : "7205",
			Name : "Kab. Gianyar, Bali",
		},
		{
			Code : "7206",
			Name : "Kab. Klungkung, Bali",
		},
		{
			Code : "7207",
			Name : "Kab. Bangli, Bali",
		},
		{
			Code : "7208",
			Name : "Kab. Karangasem, Bali",
		},
		{
			Code : "7291",
			Name : "Kotif Denpasar, Bali",
		},
		{
			Code : "7400",
			Name : "Provinsi Nusa Tenggara Timur",
		},
		{
			Code : "7401",
			Name : "Kab. Kupang, NTT",
		},
		{
			Code : "7402",
			Name : "Kab. Timor Tengah Sela, NTT",
		},
		{
			Code : "7403",
			Name : "Kab. Timor Tengah Utar, NTT",
		},
		{
			Code : "7404",
			Name : "Kab. Belu, NTT",
		},
		{
			Code : "7405",
			Name : "Kab. Alor, NTT",
		},
		{
			Code : "7406",
			Name : "Kab. Flores Timur, NTT",
		},
		{
			Code : "7407",
			Name : "Kab. Sikka, NTT",
		},
		{
			Code : "7408",
			Name : "Kab. Ende, NTT",
		},
		{
			Code : "7409",
			Name : "Kab. Ngada, NTT",
		},
		{
			Code : "7410",
			Name : "Kab. Manggarai, NTT",
		},
		{
			Code : "7411",
			Name : "Kab. Sumba Timur, NTT",
		},
		{
			Code : "7412",
			Name : "Kab. Sumba Barat, NTT",
		},
		{
			Code : "7413",
			Name : "Kab. Lembata, NTT",
		},
		{
			Code : "7414",
			Name : "Kab.Rote Ndao",
		},
		{
			Code : "7415",
			Name : "Kab.Manggarai Barat",
		},
		{
			Code : "7416",
			Name : "Kab.Sumba Tengah",
		},
		{
			Code : "7417",
			Name : "Kab.Sumba Barat Daya",
		},
		{
			Code : "7418",
			Name : "Kab.Manggarai Timur",
		},
		{
			Code : "7419",
			Name : "Kab.Nagekeo",
		},
		{
			Code : "7420",
			Name : "Kab.Sabu Raijua",
		},
		{
			Code : "7421",
			Name : "Kab.Malaka",
		},
		{
			Code : "7491",
			Name : "Kotif Kupang, NTT",
		},
		{
			Code : "7501",
			Name : "Kab. Dilli, Timor Timur",
		},
		{
			Code : "7502",
			Name : "Kab. Baucau, Timor Timur",
		},
		{
			Code : "7503",
			Name : "Kab. Manatuto, Timor Timur",
		},
		{
			Code : "7504",
			Name : "Kab. Lautem/Lospalos, Timor Ti",
		},
		{
			Code : "7505",
			Name : "Kab. Viqueque, Timor Timur",
		},
		{
			Code : "7506",
			Name : "Kab. Ainaro, Timor Timur",
		},
		{
			Code : "7507",
			Name : "Kab. Manufahi/Same, Timor Timu",
		},
		{
			Code : "7508",
			Name : "Kab. Cova-Lima/Suai, Timor Tim",
		},
		{
			Code : "7509",
			Name : "Kab. Ambeno/P.Makasar, Timor T",
		},
		{
			Code : "7510",
			Name : "Kab. Bobonaro/Mahana, Timor Ti",
		},
		{
			Code : "7511",
			Name : "Kab. Liquica, Timor Timur",
		},
		{
			Code : "7512",
			Name : "Kab. Ermera, Timor Timur",
		},
		{
			Code : "7513",
			Name : "Kab. Aileu, Timor Timur",
		},
		{
			Code : "7590",
			Name : "Kotif Dili, Timor Timur",
		},
		{
			Code : "8100",
			Name : "Provinsi Maluku",
		},
		{
			Code : "8101",
			Name : "Kab.Maluku Tengah",
		},
		{
			Code : "8102",
			Name : "Kab. Maluku Tenggara, Maluku",
		},
		{
			Code : "8103",
			Name : "Kab.Maluku Tenggara Barat",
		},
		{
			Code : "8104",
			Name : "Kab.Buru",
		},
		{
			Code : "8105",
			Name : "Kab.Seram Bagian Barat",
		},
		{
			Code : "8106",
			Name : "Kab.Seram Bagian, Timur",
		},
		{
			Code : "8107",
			Name : "Kab.Kepulauan Aru",
		},
		{
			Code : "8108",
			Name : "Kab.Maluku Barat Daya",
		},
		{
			Code : "8109",
			Name : "Kab.Buru Selatan",
		},
		{
			Code : "8191",
			Name : "Kod. Ambon, Maluku",
		},
		{
			Code : "8192",
			Name : "Kodya Ternate, Maluku",
		},
		{
			Code : "8193",
			Name : "Kotif Halmahera Tengah, Maluku",
		},
		{
			Code : "8200",
			Name : "Provinsi Papua",
		},
		{
			Code : "8201",
			Name : "Kab. Jayapura, Irian Jaya",
		},
		{
			Code : "8202",
			Name : "Kab. Teluk Cendrawasih/Biak Nu",
		},
		{
			Code : "8204",
			Name : "Kab. Sorong, Irian Jaya",
		},
		{
			Code : "8205",
			Name : "Kab. Fak-Fak, Irian Jaya",
		},
		{
			Code : "8209",
			Name : "Kab. Manokwari, Irian Jaya",
		},
		{
			Code : "8210",
			Name : "Kab. Yapen-Waropen, Irian Jaya",
		},
		{
			Code : "8211",
			Name : "Kab. Merauke, Irian Jaya",
		},
		{
			Code : "8212",
			Name : "Kab. Paniai, Irian Jaya",
		},
		{
			Code : "8213",
			Name : "Kab. Jayawijaya, Irian Jaya",
		},
		{
			Code : "8214",
			Name : "Kab. Nabire",
		},
		{
			Code : "8215",
			Name : "Kab. Mimika",
		},
		{
			Code : "8216",
			Name : "Kab. Puncak Jaya",
		},
		{
			Code : "8217",
			Name : "Kab.Sarmi",
		},
		{
			Code : "8218",
			Name : "Kab.Keerom",
		},
		{
			Code : "8221",
			Name : "Kab.Pegunungan Bintang",
		},
		{
			Code : "8222",
			Name : "Kab.Yahukimo",
		},
		{
			Code : "8223",
			Name : "Kab.Tolikara",
		},
		{
			Code : "8224",
			Name : "Kab.Waropen",
		},
		{
			Code : "8226",
			Name : "Kab.Boven Digoel",
		},
		{
			Code : "8227",
			Name : "Kab.Mappi",
		},
		{
			Code : "8228",
			Name : "Kab.Asmat",
		},
		{
			Code : "8231",
			Name : "Kab.Supiori",
		},
		{
			Code : "8232",
			Name : "Kab.Mamberamo Raya",
		},
		{
			Code : "8233",
			Name : "Kab.Dogiyai",
		},
		{
			Code : "8234",
			Name : "Kab.Lanny Jaya",
		},
		{
			Code : "8235",
			Name : "Kab.Mamberamo Tengah",
		},
		{
			Code : "8236",
			Name : "Kab.Nduga",
		},
		{
			Code : "8237",
			Name : "Kab.Yalimo",
		},
		{
			Code : "8238",
			Name : "Kab.Puncak",
		},
		{
			Code : "8239",
			Name : "Kab.Intan Jaya",
		},
		{
			Code : "8240",
			Name : "Kab.Deiya",
		},
		{
			Code : "8291",
			Name : "Kotif Jayapura, Irian Jaya",
		},
		{
			Code : "8292",
			Name : "Kodya Sorong, Irian Jaya",
		},
		{
			Code : "8300",
			Name : "Provinsi Maluku Utara",
		},
		{
			Code : "8301",
			Name : "Kab. Maluku Utara, Irian Jaya",
		},
		{
			Code : "8302",
			Name : "Kab. Halmahera Tengah, Irian J",
		},
		{
			Code : "8303",
			Name : "Kab.Halmahera Utara",
		},
		{
			Code : "8304",
			Name : "Kab.Halmahera Timur",
		},
		{
			Code : "8305",
			Name : "Kab.Halmahera Barat",
		},
		{
			Code : "8306",
			Name : "Kab.Halmahera Selatan",
		},
		{
			Code : "8307",
			Name : "Kab.Kepulauan Sula",
		},
		{
			Code : "8308",
			Name : "Kab.Pulau Morotai",
		},
		{
			Code : "8309",
			Name : "Kab.Pulau Taliabu",
		},
		{
			Code : "8390",
			Name : "Kota Ternate",
		},
		{
			Code : "8391",
			Name : "KotaTidore Kepulauan",
		},
		{
			Code : "8400",
			Name : "Provinsi Papua Barat",
		},
		{
			Code : "8401",
			Name : "Kab.Sorong",
		},
		{
			Code : "8402",
			Name : "Kab.Fak-Fak",
		},
		{
			Code : "8403",
			Name : "Kab.Manokwari",
		},
		{
			Code : "8404",
			Name : "Kab.Sorong Selatan",
		},
		{
			Code : "8405",
			Name : "Kab.Raja Ampat",
		},
		{
			Code : "8406",
			Name : "Kab.Kaimana",
		},
		{
			Code : "8407",
			Name : "Kab.Teluk Bintuni",
		},
		{
			Code : "8408",
			Name : "Kab.Teluk Wondama",
		},
		{
			Code : "8409",
			Name : "Kab.Tembrauw",
		},
		{
			Code : "8410",
			Name : "Kab.Maybrat",
		},
		{
			Code : "8411",
			Name : "Kab.Pegunungan Arfak",
		},
		{
			Code : "8412",
			Name : "Kab.Manokwari Selatan",
		},
		{
			Code : "8491",
			Name : "Kota Sorong",
		},
		{
			Code : "901",
			Name : "Kab. Semarang, Jawa Tengah",
		},
		{
			Code : "902",
			Name : "Kab. Kendal, Jawa Tengah",
		},
		{
			Code : "903",
			Name : "Kab. Demak, Jawa Tengah",
		},
		{
			Code : "904",
			Name : "Kab. Grobogan, Jawa Tengah",
		},
		{
			Code : "905",
			Name : "Kab. Pekalongan, Jawa Tengah",
		},
		{
			Code : "906",
			Name : "Kab. Tegal, Jawa Tengah",
		},
		{
			Code : "907",
			Name : "Kab. Brebes, Jawa Tengah",
		},
		{
			Code : "908",
			Name : "Kab. Pati, Jawa Tengah",
		},
		{
			Code : "909",
			Name : "Kab. Kudus, Jawa Tengah",
		},
		{
			Code : "910",
			Name : "Kab. Pemalang, Jawa Tengah",
		},
		{
			Code : "911",
			Name : "Kab. Jepara, Jawa Tengah",
		},
		{
			Code : "912",
			Name : "Kab. Rembang, Jawa Tengah",
		},
		{
			Code : "914",
			Name : "Kab. Banyumas, Jawa Tengah",
		},
		{
			Code : "915",
			Name : "Kab. Cilacap, Jawa Tengah",
		},
		{
			Code : "916",
			Name : "Kab. Purbalingga, Jawa Tengah",
		},
		{
			Code : "917",
			Name : "Kab. Banjarnegara, Jawa Tengah",
		},
		{
			Code : "918",
			Name : "Kab. Magelang, Jawa Tengah",
		},
		{
			Code : "919",
			Name : "Kab. Temanggung, Jawa Tengah",
		},
		{
			Code : "920",
			Name : "Kab. Wonosobo, Jawa Tengah",
		},
		{
			Code : "921",
			Name : "Kab. Purworejo, Jawa Tengah",
		},
		{
			Code : "922",
			Name : "Kab. Kebumen, Jawa Tengah",
		},
		{
			Code : "923",
			Name : "Kab. Klaten, Jawa Tengah",
		},
		{
			Code : "925",
			Name : "Kab. Sragen, Jawa Tengah",
		},
		{
			Code : "926",
			Name : "Kab. Sukoharjo, Jawa Tengah",
		},
		{
			Code : "927",
			Name : "Kab. Karanganyar, Jawa Tengah",
		},
		{
			Code : "928",
			Name : "Kab. Wonogiri, Jawa Tengah",
		},
		{
			Code : "929",
			Name : "Kab. Batang, Jawa Tengah",
		},
		{
			Code : "991",
			Name : "Kod. Semarang, Jawa Tengah",
		},
		{
			Code : "992",
			Name : "Kod. Salatiga, Jawa Tengah",
		},
		{
			Code : "993",
			Name : "Kod. Pekalongan, Jawa Tengah",
		},
		{
			Code : "994",
			Name : "Kod. Tegal, Jawa Tengah",
		},
		{
			Code : "995",
			Name : "Kod. Magelang, Jawa Tengah",
		},
		{
			Code : "996",
			Name : "Kod. Surakarta, Jawa Tengah",
		},
		{
			Code : "997",
			Name : "Kotif Klaten, Jawa Tengah",
		},
		{
			Code : "998",
			Name : "Kotif Cilacap, Jawa Tengah",
		},
		{
			Code : "999",
			Name : "Kotif Purwokerto, Jawa Tengah",
		},
		{
			Code : "9999",
			Name : "DI LUAR INDONESIA ( LUAR NEGERI )",
		},
	}

	var LokasiDati2 = []models.LokasiDati2{
		{
			Code : "0000 ",
			Name : "Lainnya Tidak Terdefinisi",
		},
		{
			Code : "0100rn ",
			Name : "Provinsi Jawa Baratrn",
		},
		{
			Code : "0102 ",
			Name : "Kab. Bekasi",
		},
		{
			Code : "0103 ",
			Name : "Kab. Purwakarta",
		},
		{
			Code : "0106 ",
			Name : "Kab. Karawang",
		},
		{
			Code : "0108 ",
			Name : "Kab. Bogor",
		},
		{
			Code : "0109 ",
			Name : "Kab. Sukabumi",
		},
		{
			Code : "0110 ",
			Name : "Kab. Cianjur",
		},
		{
			Code : "0111 ",
			Name : "Kab.Bandung",
		},
		{
			Code : "0112 ",
			Name : "Kab.Sumedang",
		},
		{
			Code : "0113 ",
			Name : "Kab.Tasikmalaya",
		},
		{
			Code : "0114 ",
			Name : "Kab.Garut",
		},
		{
			Code : "0115 ",
			Name : "Kab.Ciamis",
		},
		{
			Code : "0116 ",
			Name : "Kab.Cirebon",
		},
		{
			Code : "0117 ",
			Name : "Kab.Kuningan",
		},
		{
			Code : "0118 ",
			Name : "Kab.Indramayu",
		},
		{
			Code : "0119 ",
			Name : "Kab.Majalengka",
		},
		{
			Code : "0121 ",
			Name : "Kab.Subang",
		},
		{
			Code : "0122 ",
			Name : "Kab.Bandung Barat",
		},
		{
			Code : "0123 ",
			Name : "Kab.Pangandaran ",
		},
		{
			Code : "0180 ",
			Name : "Kota Banjar ",
		},
		{
			Code : "0191 ",
			Name : "Kota Bandung ",
		},
		{
			Code : "0192 ",
			Name : "Kota Bogor ",
		},
		{
			Code : "0193 ",
			Name : "Kota Sukabumi ",
		},
		{
			Code : "0194 ",
			Name : "Kota Cirebon ",
		},
		{
			Code : "0195 ",
			Name : "Kota Tasikmalaya ",
		},
		{
			Code : "0196 ",
			Name : "Kota Cimahi ",
		},
		{
			Code : "0197 ",
			Name : "Kota Depok ",
		},
		{
			Code : "0198 ",
			Name : "Kota Bekasi ",
		},
		{
			Code : "0200 ",
			Name : "Provinsi Banten",
		},
		{
			Code : "0201 ",
			Name : "Kab.Lebak",
		},
		{
			Code : "0202 ",
			Name : "Kab.Pandeglang",
		},
		{
			Code : "0203 ",
			Name : "Kab.Serang",
		},
		{
			Code : "0204 ",
			Name : "Kab.Tangerang",
		},
		{
			Code : "0291 ",
			Name : "Kota Cilegon",
		},
		{
			Code : "0292 ",
			Name : "Kota Tangerang",
		},
		{
			Code : "0293 ",
			Name : "Kota Serang",
		},
		{
			Code : "0294 ",
			Name : "Kota Tangerang Selatan",
		},
		{
			Code : "0300 ",
			Name : "Provinsi DKI Jakarta Raya",
		},
		{
			Code : "0391 ",
			Name : "Wil.Kota Jakarta Pusat",
		},
		{
			Code : "0392 ",
			Name : "Wil.Kota Jakarta Utara",
		},
		{
			Code : "0393",
			Name : "Wil.Kota Jakarta Barat",
		},
		{
			Code : "0394",
			Name : "Wil.Kota Jakarta Selatan",
		},
		{
			Code : "0395",
			Name : "Wil.Kota Jakarta Timur",
		},
		{
			Code : "0396",
			Name : "Wil.Kab.Administrasi Kepulauan Seribu",
		},
		{
			Code : "0500",
			Name : "Daerah Istimewa Yogyakarta",
		},
		{
			Code : "0501",
			Name : "Kab.Bantul",
		},
		{
			Code : "0502",
			Name : "Kab.Sleman",
		},
		{
			Code : "0503",
			Name : "Kab.Gunung Kidul",
		},
		{
			Code : "0504",
			Name : "Kab.Kulon Progo",
		},
		{
			Code : "0591",
			Name : "Kota Yogyakarta",
		},
		{
			Code : "0900",
			Name : "Provinsi Jawa Tengah",
		},
		{
			Code : "0901",
			Name : "Kab.Semarang",
		},
		{
			Code : "0902",
			Name : "Kab.Kendal",
		},
		{
			Code : "0903",
			Name : "Kab.Demak",
		},
		{
			Code : "0904",
			Name : "Kab.Grobogan",
		},
		{
			Code : "0905",
			Name : "Kab.Pekalongan",
		},
		{
			Code : "0906",
			Name : "Kab.Tegal",
		},
		{
			Code : "0907",
			Name : "Kab.Brebes",
		},
		{
			Code : "0908",
			Name : "Kab.Pati",
		},
		{
			Code : "0909",
			Name : "Kab.Kudus",
		},
		{
			Code : "0910",
			Name : "Kab.Pemalang",
		},
		{
			Code : "0911",
			Name : "Kab.Jepara",
		},
		{
			Code : "0912",
			Name : "Kab.Rembang",
		},
		{
			Code : "0913",
			Name : "Kab.Blora",
		},
		{
			Code : "0914",
			Name : "Kab.Banyumas",
		},
		{
			Code : "0915",
			Name : "Kab.Cilacap",
		},
		{
			Code : "0916",
			Name : "Kab.Purbalingga",
		},
		{
			Code : "0917",
			Name : "Kab.Banjarnegara",
		},
		{
			Code : "0918",
			Name : "Kab.Magelang",
		},
		{
			Code : "0919",
			Name : "Kab.Temanggung",
		},
		{
			Code : "0920",
			Name : "Kab.Wonosobo",
		},
		{
			Code : "0921",
			Name : "Kab.Purworejo",
		},
		{
			Code : "0922",
			Name : "Kab.Kebumen",
		},
		{
			Code : "0923",
			Name : "Kab.Klaten",
		},
		{
			Code : "0924",
			Name : "Kab.Boyolali",
		},
		{
			Code : "0925",
			Name : "Kab.Sragen",
		},
		{
			Code : "0926",
			Name : "Kab.Sukoharjo",
		},
		{
			Code : "0927",
			Name : "Kab.Karanganyar",
		},
		{
			Code : "0928",
			Name : "Kab.Wonogiri",
		},
		{
			Code : "0929",
			Name : "Kab.Batang",
		},
		{
			Code : "0991",
			Name : "Kota Semarang",
		},
		{
			Code : "0992",
			Name : "Kota Salatiga",
		},
		{
			Code : "0993",
			Name : "Kota Pekalongan",
		},
		{
			Code : "0994",
			Name : "Kota Tegal",
		},
		{
			Code : "0995",
			Name : "Kota Magelang",
		},
		{
			Code : "0996",
			Name : "Kota Surakarta/Solo",
		},
		{
			Code : "101",
			Name : "Kab. Tangerang, Jawa Barat",
		},
		{
			Code : "102",
			Name : "Kab. Bekasi, Jawa Barat",
		},
		{
			Code : "103",
			Name : "Kab. Purwakarta, Jawa Barat",
		},
		{
			Code : "104",
			Name : "Kab. Serang, Jawa Barat",
		},
		{
			Code : "105",
			Name : "Kab. Pandeglang, Jawa Barat",
		},
		{
			Code : "106",
			Name : "Kab. Karawang, Jawa Barat",
		},
		{
			Code : "107",
			Name : "Kab. Lebak, Jawa Barat",
		},
		{
			Code : "108",
			Name : "Kab. Bogor, Jawa Barat",
		},
		{
			Code : "109",
			Name : "Kab. Sukabumi, Jawa Barat",
		},
		{
			Code : "110",
			Name : "Kab. Cianjur, Jawa Barat",
		},
		{
			Code : "111",
			Name : "Kab. Bandung, Jawa Barat",
		},
		{
			Code : "112",
			Name : "Kab. Sumedang, Jawa Barat",
		},
		{
			Code : "113",
			Name : "Kab. Tasikmalaya, Jawa Barat",
		},
		{
			Code : "114",
			Name : "Kab. Garut, Jawa Barat",
		},
		{
			Code : "115",
			Name : "Kab. Ciamis, Jawa Barat",
		},
		{
			Code : "116",
			Name : "Kab. Cirebon, Jawa Barat",
		},
		{
			Code : "117",
			Name : "Kab. Kuningan, Jawa Barat",
		},
		{
			Code : "118",
			Name : "Kab. Indramayu, Jawa Barat",
		},
		{
			Code : "119",
			Name : "Kab. Majalengka, Jawa Barat",
		},
		{
			Code : "1200",
			Name : "Provinsi Jawa Timur",
		},
		{
			Code : "1201",
			Name : "Kab. Gresik, Jawa Timur",
		},
		{
			Code : "1202",
			Name : "Kab. Sidoarjo, Jawa Timur",
		},
		{
			Code : "1203",
			Name : "Kab. Mojokerto, Jawa Timur",
		},
		{
			Code : "1204",
			Name : "Kab. Jombang, Jawa Timur",
		},
		{
			Code : "1205",
			Name : "Kab. Sampang, Jawa Timur",
		},
		{
			Code : "1206",
			Name : "Kab. Pemekasan, Jawa Timur",
		},
		{
			Code : "1207",
			Name : "Kab. Sumenep, Jawa Timur",
		},
		{
			Code : "1208",
			Name : "Kab. Bangkalan, Jawa Timur",
		},
		{
			Code : "1209",
			Name : "Kab. Bondowoso, Jawa Timur",
		},
		{
			Code : "121",
			Name : "Kab. Subang, Jawa Barat",
		},
		{
			Code : "1211",
			Name : "Kab. Banyuwangi, Jawa Timur",
		},
		{
			Code : "1212",
			Name : "Kab. Jember, Jawa Timur",
		},
		{
			Code : "1213",
			Name : "Kab. Malang, Jawa Timur",
		},
		{
			Code : "1214",
			Name : "Kab. Pasuruan, Jawa Timur",
		},
		{
			Code : "1215",
			Name : "Kab. Probolinggo, Jawa Timur",
		},
		{
			Code : "1216",
			Name : "Kab. Lumajang, Jawa Timur",
		},
		{
			Code : "1217",
			Name : "Kab. Kediri, Jawa Timur",
		},
		{
			Code : "1218",
			Name : "Kab. Nganjuk, Jawa Timur",
		},
		{
			Code : "1219",
			Name : "Kab. Tulungagung, Jawa Timur",
		},
		{
			Code : "1220",
			Name : "Kab. Trenggalek, Jawa Timur",
		},
		{
			Code : "1221",
			Name : "Kab. Blitar, Jawa Timur",
		},
		{
			Code : "1222",
			Name : "Kab. Madiun, Jawa Timur",
		},
		{
			Code : "1223",
			Name : "Kab. Ngawi, Jawa Timur",
		},
		{
			Code : "1224",
			Name : "Kab. Magetan, Jawa Timur",
		},
		{
			Code : "1225",
			Name : "Kab. Ponorogo, Jawa Timur",
		},
		{
			Code : "1226",
			Name : "Kab. Pacitan, Jawa Timur",
		},
		{
			Code : "1227",
			Name : "Kab. Bojonegoro, Jawa Timur",
		},
		{
			Code : "1228",
			Name : "Kab. Tuban, Jawa Timur",
		},
		{
			Code : "1229",
			Name : "Kab. Lamongan, Jawa Timur",
		},
		{
			Code : "1230",
			Name : "Kab. Situbondo, Jawa Timur",
		},
		{
			Code : "1271",
			Name : "Kota Batu",
		},
		{
			Code : "1291",
			Name : "Kod. Surabaya, Jawa Timur",
		},
		{
			Code : "1292",
			Name : "Kod. Mojokerto, Jawa Timur",
		},
		{
			Code : "1293",
			Name : "Kod. Malang, Jawa Timur",
		},
		{
			Code : "1294",
			Name : "Kod. Pasuruan, Jawa Timur",
		},
		{
			Code : "1295",
			Name : "Kod. Probolinggo, Jawa Timur",
		},
		{
			Code : "1296",
			Name : "Kod. Blitar, Jawa Timur",
		},
		{
			Code : "1297",
			Name : "Kod. Kediri, Jawa Timur",
		},
		{
			Code : "1298",
			Name : "Kod. Madiun, Jawa Timur",
		},
		{
			Code : "1299",
			Name : "Kotif Jember, Jawa Timur",
		},
		{
			Code : "180",
			Name : "Kotif Banjar, Jawa Barat",
		},
		{
			Code : "190",
			Name : "Kotif Cilegon, Jawa Barat",
		},
		{
			Code : "191",
			Name : "Kod. Bandung, Jawa Barat",
		},
		{
			Code : "192",
			Name : "Kod. Bogor, Jawa Barat",
		},
		{
			Code : "193",
			Name : "Kod. Sukabumi, Jawa Barat",
		},
		{
			Code : "194",
			Name : "Kod. Cirebon, Jawa Barat",
		},
		{
			Code : "195",
			Name : "Kotif Tasikmalaya, Jawa Barat",
		},
		{
			Code : "196",
			Name : "Kotif Cimahi, Jawa Barat",
		},
		{
			Code : "197",
			Name : "Kotif Depok, Jawa Barat",
		},
		{
			Code : "198",
			Name : "Kotif Bekasi, Jawa Barat",
		},
		{
			Code : "199",
			Name : "Kotif Tangerang, Jawa Barat",
		},
		{
			Code : "2300",
			Name : "Provinsi Bengkulu",
		},
		{
			Code : "2301",
			Name : "Kab. Bengkulu Selatan, Bengkul",
		},
		{
			Code : "2302",
			Name : "Kab. Bengkulu Utara, Bengkulu",
		},
		{
			Code : "2303",
			Name : "Kab. Rejang Lebong, Bengkulu",
		},
		{
			Code : "2304",
			Name : "Kab. Lebong",
		},
		{
			Code : "2305",
			Name : "Kab. Kepahiang",
		},
		{
			Code : "2306",
			Name : "Kab.Mukomuko",
		},
		{
			Code : "2307",
			Name : "Kab.Seluma",
		},
		{
			Code : "2308",
			Name : "Kab.Kaur",
		},
		{
			Code : "2309",
			Name : "Kab.Bengkulu Tengah",
		},
		{
			Code : "2391",
			Name : "Kod. Bengkulu, Bengkulu",
		},
		{
			Code : "3100",
			Name : "Provinsi Jambi",
		},
		{
			Code : "3101",
			Name : "Kab. Batanghari, Jambi",
		},
		{
			Code : "3102",
			Name : "Kab. Tanjung Jabung, Jambi",
		},
		{
			Code : "3103",
			Name : "Kab. Muara Bungo-Tebo, Jambi",
		},
		{
			Code : "3104",
			Name : "Kab. Sarolangun / Bang, Jambi",
		},
		{
			Code : "3105",
			Name : "Kab. Kerinci, Jambi",
		},
		{
			Code : "3106",
			Name : "Kab. Muara Jambi, Jambi",
		},
		{
			Code : "3107",
			Name : "Kab. Tanjung Jabung Barat, Jam",
		},
		{
			Code : "3108",
			Name : "Kab. Tanjung Jabung Timur, Jam",
		},
		{
			Code : "3109",
			Name : "Kab. Tebo, Jambi",
		},
		{
			Code : "3110",
			Name : "Kab. Bungo, Jambi",
		},
		{
			Code : "3111",
			Name : "Kab. Merangin, Jambi",
		},
		{
			Code : "3112",
			Name : "Kab.Bungo",
		},
		{
			Code : "3191",
			Name : "Kod. Jambi, Jambi",
		},
		{
			Code : "3192",
			Name : "Kota Sungai Penuh",
		},
		{
			Code : "3200",
			Name : "Provinsi Nanggroe Aceh Darussalam",
		},
		{
			Code : "3201",
			Name : "Kab. Aceh Besar, Aceh",
		},
		{
			Code : "3202",
			Name : "Kab. Pidie, Aceh",
		},
		{
			Code : "3203",
			Name : "Kab. Aceh Utara, Aceh",
		},
		{
			Code : "3204",
			Name : "Kab. Aceh Timur, Aceh",
		},
		{
			Code : "3205",
			Name : "Kab. Aceh Selatan, Aceh",
		},
		{
			Code : "3206",
			Name : "Kab. Aceh Barat, Aceh",
		},
		{
			Code : "3207",
			Name : "Kab. Aceh Tengah, Aceh",
		},
		{
			Code : "3208",
			Name : "Kab. Aceh Tenggara, Aceh",
		},
		{
			Code : "3209",
			Name : "Kab. Aceh Singkil, Aceh",
		},
		{
			Code : "3210",
			Name : "Kab.Aceh Jeumpa/Bireuen",
		},
		{
			Code : "3211",
			Name : "Kab. Aceh Tamiang",
		},
		{
			Code : "3212",
			Name : "Kab. Gayo Luwes",
		},
		{
			Code : "3213",
			Name : "Kab. Aceh Barat Daya",
		},
		{
			Code : "3214",
			Name : "Kab.Aceh Jaya",
		},
		{
			Code : "3215",
			Name : "Kab.Nagan Raya",
		},
		{
			Code : "3216",
			Name : "Kab.Simeuleu",
		},
		{
			Code : "3217",
			Name : "Kab.Bener Meriah",
		},
		{
			Code : "3218",
			Name : "Kab.Pidie Jaya",
		},
		{
			Code : "3219",
			Name : "Kab.Subulussalam",
		},
		{
			Code : "3291",
			Name : "Kod. Banda Aceh, Aceh",
		},
		{
			Code : "3292",
			Name : "Kod. Sabang, Aceh",
		},
		{
			Code : "3293",
			Name : "Kotif Lhokseumawe, Aceh",
		},
		{
			Code : "3294",
			Name : "Kotif Langsa, Aceh",
		},
		{
			Code : "3295",
			Name : "Kotif Simeulue, Aceh",
		},
		{
			Code : "3300",
			Name : "Provinsi Sumatera Utara",
		},
		{
			Code : "3301",
			Name : "Kab. Deli Serdang, Sumatera Ut",
		},
		{
			Code : "3302",
			Name : "Kab. Langkat, Sumatera Utara",
		},
		{
			Code : "3303",
			Name : "Kab. Karo, Sumatera Utara",
		},
		{
			Code : "3304",
			Name : "Kab. Simalungun, Sumatera Utar",
		},
		{
			Code : "3305",
			Name : "Kab. Labuhan Batu, Sumatera Ut",
		},
		{
			Code : "3306",
			Name : "Kab. Asahan, Sumatera Utara",
		},
		{
			Code : "3307",
			Name : "Kab. Dairi, Sumatera Utara",
		},
		{
			Code : "3308",
			Name : "Kab. Tapanuli Utara, Sumatera ",
		},
		{
			Code : "3309",
			Name : "Kab. Tapanuli Tengah, Sumatera",
		},
		{
			Code : "3310",
			Name : "Kab. Tapanuli Selatan, Sumater",
		},
		{
			Code : "3311",
			Name : "Kab. Nias, Sumatera Utara",
		},
		{
			Code : "3312",
			Name : "Kotif. Rantau Prapat, Sumatera",
		},
		{
			Code : "3313",
			Name : "Kab. Toba Samosir, Sumatera Ut",
		},
		{
			Code : "3314",
			Name : "Kab. Mandailing Natal, Sumater",
		},
		{
			Code : "3315",
			Name : "Kab.Nias Selatan",
		},
		{
			Code : "3316",
			Name : "Kab.Humbang Hasundutan",
		},
		{
			Code : "3317",
			Name : "Kab.Pakpak Bharat",
		},
		{
			Code : "3318",
			Name : "Kab.Samosir",
		},
		{
			Code : "3319",
			Name : "Kab.Serdang Bedagai",
		},
		{
			Code : "3321",
			Name : "Kab.Batu Bara",
		},
		{
			Code : "3322",
			Name : "Kab.Padang Lawas",
		},
		{
			Code : "3323",
			Name : "Kab.Padang Lawas Utara",
		},
		{
			Code : "3324",
			Name : "Kab.Labuanbatu Selatan",
		},
		{
			Code : "3325",
			Name : "Kab.Labuanbatu Utara",
		},
		{
			Code : "3326",
			Name : "Kab.Nias Barat",
		},
		{
			Code : "3327",
			Name : "Kab.Nias Utara",
		},
		{
			Code : "3391",
			Name : "Kod. Tebing Tinggi, Sumatera U",
		},
		{
			Code : "3392",
			Name : "Kod. Binjai, Sumatera Utara",
		},
		{
			Code : "3393",
			Name : "Kod. Pematangsiantar, Sumatera",
		},
		{
			Code : "3394",
			Name : "Kod. Tanjung Balai, Sumatera U",
		},
		{
			Code : "3395",
			Name : "Kod. Sibolga, Sumatera Utara",
		},
		{
			Code : "3396",
			Name : "Kod. Medan, Sumatera Utara",
		},
		{
			Code : "3397",
			Name : "Kota Gunung Sitoli",
		},
		{
			Code : "3398",
			Name : "Kotif Kisaran, Sumatera Utara",
		},
		{
			Code : "3399",
			Name : "Kotif Padang Sidempuan, Sumate",
		},
		{
			Code : "3400",
			Name : "Provinsi Sumatera Barat",
		},
		{
			Code : "3401",
			Name : "Kab. Agam, Sumatera Barat",
		},
		{
			Code : "3402",
			Name : "Kab. Pasaman, Sumatera Barat",
		},
		{
			Code : "3403",
			Name : "Kab. Limapuluh Kota, Sumatera ",
		},
		{
			Code : "3404",
			Name : "Kab. Solok, Sumatera Barat",
		},
		{
			Code : "3405",
			Name : "Kab. Padang/Pariaman, Sumatera",
		},
		{
			Code : "3406",
			Name : "Kab. Pesisir Selatan, Sumatera",
		},
		{
			Code : "3407",
			Name : "Kab. Tanah Datar, Sumatera Bar",
		},
		{
			Code : "3408",
			Name : "Kab. Sawahlunto/Sijunj, Sumate",
		},
		{
			Code : "3409",
			Name : "Kab. Mentawai, Sumatera Barat",
		},
		{
			Code : "3410",
			Name : "Kab.Pasaman Barat",
		},
		{
			Code : "3412",
			Name : "Kab.Solok",
		},
		{
			Code : "3491",
			Name : "Kod. Bukittinggi, Sumatera Bar",
		},
		{
			Code : "3492",
			Name : "Kod. Padang, Sumatera Barat",
		},
		{
			Code : "3493",
			Name : "Kod. Sawahlunto, Sumatera Bara",
		},
		{
			Code : "3494",
			Name : "Kod. Padangpanjang, Sumatera B",
		},
		{
			Code : "3495",
			Name : "Kod. Solok, Sumatera Barat",
		},
		{
			Code : "3496",
			Name : "Kod. Payakumbuh, Sumatera Bara",
		},
		{
			Code : "3497",
			Name : "Kotif Pariaman, Sumatera Barat",
		},
		{
			Code : "3500",
			Name : "Provinsi Riau",
		},
		{
			Code : "3501",
			Name : "Kab. Kampar, Riau",
		},
		{
			Code : "3502",
			Name : "Kab. Bengkalis, Riau",
		},
		{
			Code : "3503",
			Name : "Kab. Riau Kepulauan, Riau",
		},
		{
			Code : "3504",
			Name : "Kab. Indragiri Hulu, Riau",
		},
		{
			Code : "3505",
			Name : "Kab. Indragiri Hilir, Riau",
		},
		{
			Code : "3506",
			Name : "Kab. Karimun, Riau",
		},
		{
			Code : "3507",
			Name : "Kab. Natuna, Riau",
		},
		{
			Code : "3508",
			Name : "Kab. Rokan Hulu, Riau",
		},
		{
			Code : "3509",
			Name : "Kab. Rokan Hilir, Riau",
		},
		{
			Code : "3510",
			Name : "Kab. Pelalawan, Riau",
		},
		{
			Code : "3511",
			Name : "Kab. Siak, Riau",
		},
		{
			Code : "3512",
			Name : "Kab. Kuantan Singingi",
		},
		{
			Code : "3513",
			Name : "Kab.Kepulauan Meranti",
		},
		{
			Code : "3591",
			Name : "Kod. Pekanbaru, Riau",
		},
		{
			Code : "3592",
			Name : "Kotif Dumai, Riau",
		},
		{
			Code : "3593",
			Name : "Kotif Tanjungpinang, Riau",
		},
		{
			Code : "3594",
			Name : "Kotif Pulau Batam, Riau",
		},
		{
			Code : "3600",
			Name : "Provinsi Sumatera Selatan",
		},
		{
			Code : "3604",
			Name : "Kab. Belitung, Sumatera Selata",
		},
		{
			Code : "3605",
			Name : "Kab. Bangka, Sumatera Selatan",
		},
		{
			Code : "3606",
			Name : "Kab. Musi/Banyuasin, Sumatera ",
		},
		{
			Code : "3607",
			Name : "Kab. Ogan Komering Ulu, Sumate",
		},
		{
			Code : "3608",
			Name : "Kab.Lematang Ilir Ogan, Sumate",
		},
		{
			Code : "3609",
			Name : "Kab. Lahat, Sumatera Selatan",
		},
		{
			Code : "3610",
			Name : "Kab. Musi Rawas, Sumatera Sela",
		},
		{
			Code : "3611",
			Name : "Kab.Ogan Komering Ilir, Sumate",
		},
		{
			Code : "3613",
			Name : "Kab.Banyuasin",
		},
		{
			Code : "3614",
			Name : "Kab.Ogan Komering Ulu Selatan",
		},
		{
			Code : "3615",
			Name : "Kab.Ogan Komering Ulu Timur",
		},
		{
			Code : "3616",
			Name : "Kab.Ogan Ilir",
		},
		{
			Code : "3617",
			Name : "Kab.Empat Lawang",
		},
		{
			Code : "3618",
			Name : "Kab.Musi Rawas Utara",
		},
		{
			Code : "3619",
			Name : "Kab.Penukal Abab Lematang Ilir",
		},
		{
			Code : "3688",
			Name : "Prov.Sumatera Selatan,Kab/Kota Lainnya",
		},
		{
			Code : "3691",
			Name : "Kod. Palembang, Sumatera Selat",
		},
		{
			Code : "3692",
			Name : "Kod. Pangkal Pinang, Sumatera ",
		},
		{
			Code : "3693",
			Name : "Kotif Lubuklinggau, Sumatera S",
		},
		{
			Code : "3694",
			Name : "Kotif Prabumulih, Sumatera Sel",
		},
		{
			Code : "3695",
			Name : "Kotif Baturaja, Sumatera Selat",
		},
		{
			Code : "3697",
			Name : "Kotif Pagar Alam, Sumatera Sel",
		},
		{
			Code : "3700",
			Name : "Provinsi Kep.Bangka Belitung",
		},
		{
			Code : "3701",
			Name : "Kab.Bangka",
		},
		{
			Code : "3702",
			Name : "Kab.Belitung",
		},
		{
			Code : "3703",
			Name : "Kab.Bangka Barat",
		},
		{
			Code : "3704",
			Name : "Kab.Bangka Selatan",
		},
		{
			Code : "3705",
			Name : "Kab.Bangka Tengah",
		},
		{
			Code : "3706",
			Name : "Kab.Belitung Timur",
		},
		{
			Code : "3707",
			Name : "Kota Pangkal Pinang",
		},
		{
			Code : "3800",
			Name : "Provinsi Kep.Riau",
		},
		{
			Code : "3801",
			Name : "Kab.Karimun",
		},
		{
			Code : "3802",
			Name : "Kab.Lingga",
		},
		{
			Code : "3803",
			Name : "Kab.Natuna",
		},
		{
			Code : "3804",
			Name : "Kab.Bintan (d/h Kabupaten Kepulauan Riau)",
		},
		{
			Code : "3805",
			Name : "Kab.Kepulauan Anambas",
		},
		{
			Code : "3891",
			Name : "Kota Tanjung Pinang",
		},
		{
			Code : "3892",
			Name : "Kota Batam",
		},
		{
			Code : "3900",
			Name : "Provinsi Lampung",
		},
		{
			Code : "3901",
			Name : "Kab. Lampung Selatan, Lampung",
		},
		{
			Code : "3902",
			Name : "Kab. Lampung Tengah, Lampung",
		},
		{
			Code : "3903",
			Name : "Kab. Lampung Utara, Lampung",
		},
		{
			Code : "3904",
			Name : "Kab. Lampung Barat, Lampung",
		},
		{
			Code : "3905",
			Name : "Kab. Tulang Bawang, Lampung",
		},
		{
			Code : "3906",
			Name : "Kab. Tanggamus, Lampung",
		},
		{
			Code : "3907",
			Name : "Kab. Lampung Timur, Lampung",
		},
		{
			Code : "3908",
			Name : "Kab. Way Kanan, Lampung",
		},
		{
			Code : "3909",
			Name : "Kab.Pesawaran",
		},
		{
			Code : "391",
			Name : "Wil. Jakarta Pusat, Jakarta",
		},
		{
			Code : "3910",
			Name : "Kab.Pringsewu",
		},
		{
			Code : "3911",
			Name : "Kab.Tulang Bawang Barat",
		},
		{
			Code : "3912",
			Name : "Kab.Mesuji",
		},
		{
			Code : "3913",
			Name : "Kab.Pesisir Barat",
		},
		{
			Code : "392",
			Name : "Wil. Jakarta Utara, Jakarta",
		},
		{
			Code : "393",
			Name : "Wil. Jakarta Barat, Jakarta",
		},
		{
			Code : "394",
			Name : "Wil. Jakarta Selatan, Jakarta",
		},
		{
			Code : "395",
			Name : "Wil. Jakarta Timur, Jakarta",
		},
		{
			Code : "3991",
			Name : "Kod. Bandar Lampung, Lampung",
		},
		{
			Code : "3992",
			Name : "Kotif Metro, Lampung",
		},
		{
			Code : "4411",
			Name : "Kab.Dharmasraya",
		},
		{
			Code : "501",
			Name : "Kab. Bantul, Yogyakarta",
		},
		{
			Code : "502",
			Name : "Kab. Sleman, Yogyakarta",
		},
		{
			Code : "503",
			Name : "Kab. Gunung Kidul, Yogyakarta",
		},
		{
			Code : "504",
			Name : "Kab. Kulon Progo, Yogyakarta",
		},
		{
			Code : "5100",
			Name : "Provinsi Kalimantan Selatan",
		},
		{
			Code : "5101",
			Name : "Kab. Banjar, Kalimantan Selata",
		},
		{
			Code : "5102",
			Name : "Kab. Tanah Laut, Kalimantan Se",
		},
		{
			Code : "5103",
			Name : "Kab. Tapin, Kalimantan Selatan",
		},
		{
			Code : "5104",
			Name : "Kab.Hulu Sungai Selata, Kal-Se",
		},
		{
			Code : "5105",
			Name : "Kab.Hulu Sungai Tengah, Kal-Se",
		},
		{
			Code : "5106",
			Name : "Kab.Hulu Sungai Utara, Kal-Sel",
		},
		{
			Code : "5107",
			Name : "Kab. Barito Kuala, Kalimantan ",
		},
		{
			Code : "5108",
			Name : "Kab. Kota Baru, Kalimantan Sel",
		},
		{
			Code : "5109",
			Name : "Kab. Tobalong, Kalimantan Sela",
		},
		{
			Code : "5110",
			Name : "Kab.Tanah Bumbu",
		},
		{
			Code : "5111",
			Name : "Kab.Balangan",
		},
		{
			Code : "5191",
			Name : "Kod. Banjarmasin, Kalimantan S",
		},
		{
			Code : "5192",
			Name : "Kotif Banjarbaru, Kalimantan S",
		},
		{
			Code : "5300",
			Name : "Provinsi Kalimantan Barat",
		},
		{
			Code : "5301",
			Name : "Kab. Pontianak, Kalimatan Bara",
		},
		{
			Code : "5302",
			Name : "Kab. Sambas, Kalimatan Barat",
		},
		{
			Code : "5303",
			Name : "Kab. Ketapang, Kalimatan Barat",
		},
		{
			Code : "5304",
			Name : "Kab. Sanggau, Kalimatan Barat",
		},
		{
			Code : "5305",
			Name : "Kab. Sintang, Kalimatan Barat",
		},
		{
			Code : "5306",
			Name : "Kab. Kapuas Hulu, Kalimatan Ba",
		},
		{
			Code : "5307",
			Name : "Kab. Bengkayang, Kalimatan Bar",
		},
		{
			Code : "5308",
			Name : "Kab. Landak, Kalimatan Barat",
		},
		{
			Code : "5309",
			Name : "Kab.Sekadau",
		},
		{
			Code : "5310",
			Name : "Kab.Melawi",
		},
		{
			Code : "5311",
			Name : "Kab.Kayong Utara",
		},
		{
			Code : "5312",
			Name : "Kab.Kubu Raya",
		},
		{
			Code : "5391",
			Name : "Kod. Pontianak, Kalimatan Bara",
		},
		{
			Code : "5392",
			Name : "Kotif Singkawang, Kalimatan Ba",
		},
		{
			Code : "5400",
			Name : "Provinsi Kalimantan Timur",
		},
		{
			Code : "5401",
			Name : "Kab. Kutai, Kalimantan Timur",
		},
		{
			Code : "5402",
			Name : "Kab. Berau, Kalimantan Timur",
		},
		{
			Code : "5403",
			Name : "Kab. Tanah Pasir, Kalimantan T",
		},
		{
			Code : "5404",
			Name : "Kab. Bulungan, Kalimantan Timu",
		},
		{
			Code : "5405",
			Name : "Kab. Kutai Barat, Kalimantan T",
		},
		{
			Code : "5406",
			Name : "Kab. Kutai Timur, Kalimantan T",
		},
		{
			Code : "5407",
			Name : "Kab. Malinau, Kalimantan Timur",
		},
		{
			Code : "5408",
			Name : "Kab. Nunukan, Kalimantan Timur",
		},
		{
			Code : "5409",
			Name : "Kab.Nunukan",
		},
		{
			Code : "5410",
			Name : "Kab.Malinau",
		},
		{
			Code : "5411",
			Name : "Kab.Penajam Paser Utara",
		},
		{
			Code : "5412",
			Name : "Kab.Tana Tidung",
		},
		{
			Code : "5413",
			Name : "kab.Mahakam Ulu",
		},
		{
			Code : "5491",
			Name : "Kod. Samarinda, Kalimantan Tim",
		},
		{
			Code : "5492",
			Name : "Kod. Balikpapan, Kalimantan Ti",
		},
		{
			Code : "5493",
			Name : "Kotif Tarakan, Kalimantan Timu",
		},
		{
			Code : "5494",
			Name : "Kotif Bontang, Kalimantan Timu",
		},
		{
			Code : "5500",
			Name : "Provinsi Kalimantan Utara",
		},
		{
			Code : "5800",
			Name : "Provinsi Kalimantan Tengah",
		},
		{
			Code : "5801",
			Name : "Kab. Kapuas, Kalimantan Tengah",
		},
		{
			Code : "5802",
			Name : "Kab.Kotawaringin Barat",
		},
		{
			Code : "5803",
			Name : "Kab. Kotawaringin Barat, Kal-T",
		},
		{
			Code : "5804",
			Name : "Kab.Murung Raya",
		},
		{
			Code : "5805",
			Name : "Kab.Barito Timur",
		},
		{
			Code : "5806",
			Name : "Kab. Barito Selatan, Kal-Teng",
		},
		{
			Code : "5807",
			Name : "Kab.Gunung Mas",
		},
		{
			Code : "5808",
			Name : "Kab. Barito Utara, Kal-Teng",
		},
		{
			Code : "5809",
			Name : "Kab.Pulang Pisau",
		},
		{
			Code : "5810",
			Name : "Kab.Seruyan",
		},
		{
			Code : "5811",
			Name : "Kab.Katingan",
		},
		{
			Code : "5812",
			Name : "Kab.Sukamara",
		},
		{
			Code : "5813",
			Name : "Kab.Lamandu",
		},
		{
			Code : "5892",
			Name : "Kod. Palangkaraya, Kal-Teng",
		},
		{
			Code : "591",
			Name : "Kod. Yogyakarta, Yogyakarta",
		},
		{
			Code : "6000",
			Name : "Provinsi Sulawesi Tengah",
		},
		{
			Code : "6001",
			Name : "Kab. Donggala, Sulawesi Tengah",
		},
		{
			Code : "6002",
			Name : "Kab. Poso, Sulawesi Tengah",
		},
		{
			Code : "6003",
			Name : "Kab. Banggai, Sulawesi Tengah",
		},
		{
			Code : "6004",
			Name : "Kab. Toli-toli, Sulawesi Tenga",
		},
		{
			Code : "6005",
			Name : "Kab.Banggai Kepulauan, Sulawes",
		},
		{
			Code : "6006",
			Name : "Kab. Morowali, Sulawesi Tengah",
		},
		{
			Code : "6007",
			Name : "Kab. Buol, Sulawesi Tengah",
		},
		{
			Code : "6008",
			Name : "Kab.Tojo Una-Una",
		},
		{
			Code : "6009",
			Name : "Kab.Parigi Moutong",
		},
		{
			Code : "6010",
			Name : "Kab.Sigi",
		},
		{
			Code : "6011",
			Name : "Kab.Banggai Laut",
		},
		{
			Code : "6012",
			Name : "Kab.Morowali Utara",
		},
		{
			Code : "6091",
			Name : "Kotif Palu, Sulawesi Tengah",
		},
		{
			Code : "6100",
			Name : "Provinsi Sulawesi Selatan",
		},
		{
			Code : "6101",
			Name : "Kab. Pinrang, Sulawesi Selatan",
		},
		{
			Code : "6102",
			Name : "Kab. Gowa, Sulawesi Selatan",
		},
		{
			Code : "6103",
			Name : "Kab. Wajo, Sulawesi Selatan",
		},
		{
			Code : "6104",
			Name : "Kab. Mamuju, Sulawesi Selatan",
		},
		{
			Code : "6105",
			Name : "Kab. Bone, Sulawesi Selatan",
		},
		{
			Code : "6106",
			Name : "Kab. Tana Toraja, Sulawesi Sel",
		},
		{
			Code : "6107",
			Name : "Kab. Maros, Sulawesi Selatan",
		},
		{
			Code : "6108",
			Name : "Kab. Majene, Sulawesi Selatan",
		},
		{
			Code : "6109",
			Name : "Kab. Luwu, Sulawesi Selatan",
		},
		{
			Code : "6110",
			Name : "Kab. Sinjai, Sulawesi Selatan",
		},
		{
			Code : "6111",
			Name : "Kab. Bulukumba, Sulawesi Selat",
		},
		{
			Code : "6112",
			Name : "Kab. Bantaeng, Sulawesi Selata",
		},
		{
			Code : "6113",
			Name : "Kab. Jeneponto, Sulawesi Selat",
		},
		{
			Code : "6114",
			Name : "Kab. Selayar, Sulawesi Selatan",
		},
		{
			Code : "6115",
			Name : "Kab. Takalar, Sulawesi Selatan",
		},
		{
			Code : "6116",
			Name : "Kab. Barru, Sulawesi Selatan",
		},
		{
			Code : "6117",
			Name : "Kab. Sindenreng Rappan, Sulawe",
		},
		{
			Code : "6118",
			Name : "Kab. Pangkajene Kepula, Sulawe",
		},
		{
			Code : "6119",
			Name : "Kab. Soppeng, Sulawesi Selatan",
		},
		{
			Code : "6120",
			Name : "Kab. Polewali Mamasa, Sulawesi",
		},
		{
			Code : "6121",
			Name : "Kab. Enrekang, Sulawesi Selata",
		},
		{
			Code : "6122",
			Name : "Kab. Luwu Selatan, Sulawesi Se",
		},
		{
			Code : "6124",
			Name : "Kab.Luwu Utara",
		},
		{
			Code : "6125",
			Name : "Kab.Toraja Utara",
		},
		{
			Code : "6191",
			Name : "Kod. Ujungpandang, Sulawesi Se",
		},
		{
			Code : "6192",
			Name : "Kod. Pare-pare, Sulawesi Selat",
		},
		{
			Code : "6193",
			Name : "Kotif Palopo, Sulawesi Selatan",
		},
		{
			Code : "6194",
			Name : "Kotif Watampone, Sulawesi Sela",
		},
		{
			Code : "6200",
			Name : "Provinsi Sulawesi Utara",
		},
		{
			Code : "6201",
			Name : "Kab. Gorontalo, Sulawesi Utara",
		},
		{
			Code : "6202",
			Name : "Kab. Minahasa, Sulawesi Utara",
		},
		{
			Code : "6203",
			Name : "Kab. Bolaang Mongondow, Sul-Ut",
		},
		{
			Code : "6204",
			Name : "Kab. Sangihe Talaud, Sulawesi ",
		},
		{
			Code : "6205",
			Name : "kab. Bitung Sulawesi Utara",
		},
		{
			Code : "6206",
			Name : "Kab. Bualemo, Sulawesi Utara",
		},
		{
			Code : "6207",
			Name : "Kab.Minahasa Utara",
		},
		{
			Code : "6209",
			Name : "Kab.Minahasa Tenggara",
		},
		{
			Code : "6210",
			Name : "Kab.Bolaang Mongondow Utara",
		},
		{
			Code : "6211",
			Name : "Kab.Kepulauan Sitaro",
		},
		{
			Code : "6212",
			Name : "Kab.Bolaang Mongondow Selatan",
		},
		{
			Code : "6213",
			Name : "Kab.Bolaang Mongondow Timur",
		},
		{
			Code : "6291",
			Name : "Kod. Manado, Sulawesi Utara",
		},
		{
			Code : "6292",
			Name : "Kod. Gorontalo, Sulawesi Utara",
		},
		{
			Code : "6293",
			Name : "Kod. Bitung, Sulawesi Utara",
		},
		{
			Code : "6294",
			Name : "Kota Tomohon",
		},
		{
			Code : "6300",
			Name : "Provinsi Gorontalo",
		},
		{
			Code : "6301",
			Name : "Kab.Gorontalo",
		},
		{
			Code : "6303",
			Name : "Kab.Bonebolango",
		},
		{
			Code : "6304",
			Name : "Kab.Pohuwato",
		},
		{
			Code : "6305",
			Name : "Kab.Gorontalo Utara",
		},
		{
			Code : "6391",
			Name : "Kota Gorontalo",
		},
		{
			Code : "6400",
			Name : "Provinsi Sulawesi Barat",
		},
		{
			Code : "6401",
			Name : "Kab.Polewali Mandar",
		},
		{
			Code : "6402",
			Name : "Kab.Majene",
		},
		{
			Code : "6403",
			Name : "Kab.Mamasa",
		},
		{
			Code : "6404",
			Name : "Kab.Mamuju Utara",
		},
		{
			Code : "6405",
			Name : "Kab.Mamuju Tengah",
		},
		{
			Code : "6406",
			Name : "Kab.Mamuju",
		},
		{
			Code : "6900",
			Name : "Provinsi Sulawesi Tenggara",
		},
		{
			Code : "6901",
			Name : "Kab. Buton, Sulawesi Tenggara",
		},
		{
			Code : "6902",
			Name : "Kab. Kendari, Sulawesi Tenggar",
		},
		{
			Code : "6903",
			Name : "Kab. Muna, Sulawesi Tenggara",
		},
		{
			Code : "6904",
			Name : "Kab. Kolaka, Sulawesi Tenggara",
		},
		{
			Code : "6905",
			Name : "Kab.Wakatobi",
		},
		{
			Code : "6906",
			Name : "Kab.Konawe",
		},
		{
			Code : "6907",
			Name : "Kab.Konawe Selatan",
		},
		{
			Code : "6908",
			Name : "Kab.Bombana",
		},
		{
			Code : "6909",
			Name : "Kab.Kolaka Utara",
		},
		{
			Code : "6910",
			Name : "Kab.Buton Utara",
		},
		{
			Code : "6911",
			Name : "Kab.Konawe Utara",
		},
		{
			Code : "6912",
			Name : "Kab.Kaloka Timur",
		},
		{
			Code : "6913",
			Name : "Kab.Konawe Kepulauan",
		},
		{
			Code : "6914",
			Name : "Kab.Buton Selatan",
		},
		{
			Code : "6915",
			Name : "Kab.Buton Tengah",
		},
		{
			Code : "6916",
			Name : "Kab.Muna Barat",
		},
		{
			Code : "6990",
			Name : "Kotif Bau-Bau, Sulawesi Tengga",
		},
		{
			Code : "6991",
			Name : "Kotif Kendari, Sulawesi Tengga",
		},
		{
			Code : "7100",
			Name : "Provinsi Nusa Tenggara Barat",
		},
		{
			Code : "7101",
			Name : "Kab. Lombok Barat, NTB",
		},
		{
			Code : "7102",
			Name : "Kab. Lombok Tengah, NTB",
		},
		{
			Code : "7103",
			Name : "Kab. Lombok Timur, NTB",
		},
		{
			Code : "7104",
			Name : "Kab. Sumbawa, NTB",
		},
		{
			Code : "7105",
			Name : "Kab. Bima, NTB",
		},
		{
			Code : "7106",
			Name : "Kab. Dompu, NTB",
		},
		{
			Code : "7107",
			Name : "Kab.Sumbawa Barat",
		},
		{
			Code : "7108",
			Name : "Kab.Lombok Utara",
		},
		{
			Code : "7191",
			Name : "Kotif Mataram, NTB",
		},
		{
			Code : "7192",
			Name : "Kota Bima",
		},
		{
			Code : "7200",
			Name : "Provinsi Bali",
		},
		{
			Code : "7201",
			Name : "Kab. Buleleng, Bali",
		},
		{
			Code : "7202",
			Name : "Kab. Jembrana, Bali",
		},
		{
			Code : "7203",
			Name : "Kab. Tabanan, Bali",
		},
		{
			Code : "7204",
			Name : "Kab. Badung, Bali",
		},
		{
			Code : "7205",
			Name : "Kab. Gianyar, Bali",
		},
		{
			Code : "7206",
			Name : "Kab. Klungkung, Bali",
		},
		{
			Code : "7207",
			Name : "Kab. Bangli, Bali",
		},
		{
			Code : "7208",
			Name : "Kab. Karangasem, Bali",
		},
		{
			Code : "7291",
			Name : "Kotif Denpasar, Bali",
		},
		{
			Code : "7400",
			Name : "Provinsi Nusa Tenggara Timur",
		},
		{
			Code : "7401",
			Name : "Kab. Kupang, NTT",
		},
		{
			Code : "7402",
			Name : "Kab. Timor Tengah Sela, NTT",
		},
		{
			Code : "7403",
			Name : "Kab. Timor Tengah Utar, NTT",
		},
		{
			Code : "7404",
			Name : "Kab. Belu, NTT",
		},
		{
			Code : "7405",
			Name : "Kab. Alor, NTT",
		},
		{
			Code : "7406",
			Name : "Kab. Flores Timur, NTT",
		},
		{
			Code : "7407",
			Name : "Kab. Sikka, NTT",
		},
		{
			Code : "7408",
			Name : "Kab. Ende, NTT",
		},
		{
			Code : "7409",
			Name : "Kab. Ngada, NTT",
		},
		{
			Code : "7410",
			Name : "Kab. Manggarai, NTT",
		},
		{
			Code : "7411",
			Name : "Kab. Sumba Timur, NTT",
		},
		{
			Code : "7412",
			Name : "Kab. Sumba Barat, NTT",
		},
		{
			Code : "7413",
			Name : "Kab. Lembata, NTT",
		},
		{
			Code : "7414",
			Name : "Kab.Rote Ndao",
		},
		{
			Code : "7415",
			Name : "Kab.Manggarai Barat",
		},
		{
			Code : "7416",
			Name : "Kab.Sumba Tengah",
		},
		{
			Code : "7417",
			Name : "Kab.Sumba Barat Daya",
		},
		{
			Code : "7418",
			Name : "Kab.Manggarai Timur",
		},
		{
			Code : "7419",
			Name : "Kab.Nagekeo",
		},
		{
			Code : "7420",
			Name : "Kab.Sabu Raijua",
		},
		{
			Code : "7421",
			Name : "Kab.Malaka",
		},
		{
			Code : "7491",
			Name : "Kotif Kupang, NTT",
		},
		{
			Code : "7501",
			Name : "Kab. Dilli, Timor Timur",
		},
		{
			Code : "7502",
			Name : "Kab. Baucau, Timor Timur",
		},
		{
			Code : "7503",
			Name : "Kab. Manatuto, Timor Timur",
		},
		{
			Code : "7504",
			Name : "Kab. Lautem/Lospalos, Timor Ti",
		},
		{
			Code : "7505",
			Name : "Kab. Viqueque, Timor Timur",
		},
		{
			Code : "7506",
			Name : "Kab. Ainaro, Timor Timur",
		},
		{
			Code : "7507",
			Name : "Kab. Manufahi/Same, Timor Timu",
		},
		{
			Code : "7508",
			Name : "Kab. Cova-Lima/Suai, Timor Tim",
		},
		{
			Code : "7509",
			Name : "Kab. Ambeno/P.Makasar, Timor T",
		},
		{
			Code : "7510",
			Name : "Kab. Bobonaro/Mahana, Timor Ti",
		},
		{
			Code : "7511",
			Name : "Kab. Liquica, Timor Timur",
		},
		{
			Code : "7512",
			Name : "Kab. Ermera, Timor Timur",
		},
		{
			Code : "7513",
			Name : "Kab. Aileu, Timor Timur",
		},
		{
			Code : "7590",
			Name : "Kotif Dili, Timor Timur",
		},
		{
			Code : "8100",
			Name : "Provinsi Maluku",
		},
		{
			Code : "8101",
			Name : "Kab.Maluku Tengah",
		},
		{
			Code : "8102",
			Name : "Kab. Maluku Tenggara, Maluku",
		},
		{
			Code : "8103",
			Name : "Kab.Maluku Tenggara Barat",
		},
		{
			Code : "8104",
			Name : "Kab.Buru",
		},
		{
			Code : "8105",
			Name : "Kab.Seram Bagian Barat",
		},
		{
			Code : "8106",
			Name : "Kab.Seram Bagian, Timur",
		},
		{
			Code : "8107",
			Name : "Kab.Kepulauan Aru",
		},
		{
			Code : "8108",
			Name : "Kab.Maluku Barat Daya",
		},
		{
			Code : "8109",
			Name : "Kab.Buru Selatan",
		},
		{
			Code : "8191",
			Name : "Kod. Ambon, Maluku",
		},
		{
			Code : "8192",
			Name : "Kodya Ternate, Maluku",
		},
		{
			Code : "8193",
			Name : "Kotif Halmahera Tengah, Maluku",
		},
		{
			Code : "8200",
			Name : "Provinsi Papua",
		},
		{
			Code : "8201",
			Name : "Kab. Jayapura, Irian Jaya",
		},
		{
			Code : "8202",
			Name : "Kab. Teluk Cendrawasih/Biak Nu",
		},
		{
			Code : "8204",
			Name : "Kab. Sorong, Irian Jaya",
		},
		{
			Code : "8205",
			Name : "Kab. Fak-Fak, Irian Jaya",
		},
		{
			Code : "8209",
			Name : "Kab. Manokwari, Irian Jaya",
		},
		{
			Code : "8210",
			Name : "Kab. Yapen-Waropen, Irian Jaya",
		},
		{
			Code : "8211",
			Name : "Kab. Merauke, Irian Jaya",
		},
		{
			Code : "8212",
			Name : "Kab. Paniai, Irian Jaya",
		},
		{
			Code : "8213",
			Name : "Kab. Jayawijaya, Irian Jaya",
		},
		{
			Code : "8214",
			Name : "Kab. Nabire",
		},
		{
			Code : "8215",
			Name : "Kab. Mimika",
		},
		{
			Code : "8216",
			Name : "Kab. Puncak Jaya",
		},
		{
			Code : "8217",
			Name : "Kab.Sarmi",
		},
		{
			Code : "8218",
			Name : "Kab.Keerom",
		},
		{
			Code : "8221",
			Name : "Kab.Pegunungan Bintang",
		},
		{
			Code : "8222",
			Name : "Kab.Yahukimo",
		},
		{
			Code : "8223",
			Name : "Kab.Tolikara",
		},
		{
			Code : "8224",
			Name : "Kab.Waropen",
		},
		{
			Code : "8226",
			Name : "Kab.Boven Digoel",
		},
		{
			Code : "8227",
			Name : "Kab.Mappi",
		},
		{
			Code : "8228",
			Name : "Kab.Asmat",
		},
		{
			Code : "8231",
			Name : "Kab.Supiori",
		},
		{
			Code : "8232",
			Name : "Kab.Mamberamo Raya",
		},
		{
			Code : "8233",
			Name : "Kab.Dogiyai",
		},
		{
			Code : "8234",
			Name : "Kab.Lanny Jaya",
		},
		{
			Code : "8235",
			Name : "Kab.Mamberamo Tengah",
		},
		{
			Code : "8236",
			Name : "Kab.Nduga",
		},
		{
			Code : "8237",
			Name : "Kab.Yalimo",
		},
		{
			Code : "8238",
			Name : "Kab.Puncak",
		},
		{
			Code : "8239",
			Name : "Kab.Intan Jaya",
		},
		{
			Code : "8240",
			Name : "Kab.Deiya",
		},
		{
			Code : "8291",
			Name : "Kotif Jayapura, Irian Jaya",
		},
		{
			Code : "8292",
			Name : "Kodya Sorong, Irian Jaya",
		},
		{
			Code : "8300",
			Name : "Provinsi Maluku Utara",
		},
		{
			Code : "8301",
			Name : "Kab. Maluku Utara, Irian Jaya",
		},
		{
			Code : "8302",
			Name : "Kab. Halmahera Tengah, Irian J",
		},
		{
			Code : "8303",
			Name : "Kab.Halmahera Utara",
		},
		{
			Code : "8304",
			Name : "Kab.Halmahera Timur",
		},
		{
			Code : "8305",
			Name : "Kab.Halmahera Barat",
		},
		{
			Code : "8306",
			Name : "Kab.Halmahera Selatan",
		},
		{
			Code : "8307",
			Name : "Kab.Kepulauan Sula",
		},
		{
			Code : "8308",
			Name : "Kab.Pulau Morotai",
		},
		{
			Code : "8309",
			Name : "Kab.Pulau Taliabu",
		},
		{
			Code : "8390",
			Name : "Kota Ternate",
		},
		{
			Code : "8391",
			Name : "KotaTidore Kepulauan",
		},
		{
			Code : "8400",
			Name : "Provinsi Papua Barat",
		},
		{
			Code : "8401",
			Name : "Kab.Sorong",
		},
		{
			Code : "8402",
			Name : "Kab.Fak-Fak",
		},
		{
			Code : "8403",
			Name : "Kab.Manokwari",
		},
		{
			Code : "8404",
			Name : "Kab.Sorong Selatan",
		},
		{
			Code : "8405",
			Name : "Kab.Raja Ampat",
		},
		{
			Code : "8406",
			Name : "Kab.Kaimana",
		},
		{
			Code : "8407",
			Name : "Kab.Teluk Bintuni",
		},
		{
			Code : "8408",
			Name : "Kab.Teluk Wondama",
		},
		{
			Code : "8409",
			Name : "Kab.Tembrauw",
		},
		{
			Code : "8410",
			Name : "Kab.Maybrat",
		},
		{
			Code : "8411",
			Name : "Kab.Pegunungan Arfak",
		},
		{
			Code : "8412",
			Name : "Kab.Manokwari Selatan",
		},
		{
			Code : "8491",
			Name : "Kota Sorong",
		},
		{
			Code : "901",
			Name : "Kab. Semarang, Jawa Tengah",
		},
		{
			Code : "902",
			Name : "Kab. Kendal, Jawa Tengah",
		},
		{
			Code : "903",
			Name : "Kab. Demak, Jawa Tengah",
		},
		{
			Code : "904",
			Name : "Kab. Grobogan, Jawa Tengah",
		},
		{
			Code : "905",
			Name : "Kab. Pekalongan, Jawa Tengah",
		},
		{
			Code : "906",
			Name : "Kab. Tegal, Jawa Tengah",
		},
		{
			Code : "907",
			Name : "Kab. Brebes, Jawa Tengah",
		},
		{
			Code : "908",
			Name : "Kab. Pati, Jawa Tengah",
		},
		{
			Code : "909",
			Name : "Kab. Kudus, Jawa Tengah",
		},
		{
			Code : "910",
			Name : "Kab. Pemalang, Jawa Tengah",
		},
		{
			Code : "911",
			Name : "Kab. Jepara, Jawa Tengah",
		},
		{
			Code : "912",
			Name : "Kab. Rembang, Jawa Tengah",
		},
		{
			Code : "914",
			Name : "Kab. Banyumas, Jawa Tengah",
		},
		{
			Code : "915",
			Name : "Kab. Cilacap, Jawa Tengah",
		},
		{
			Code : "916",
			Name : "Kab. Purbalingga, Jawa Tengah",
		},
		{
			Code : "917",
			Name : "Kab. Banjarnegara, Jawa Tengah",
		},
		{
			Code : "918",
			Name : "Kab. Magelang, Jawa Tengah",
		},
		{
			Code : "919",
			Name : "Kab. Temanggung, Jawa Tengah",
		},
		{
			Code : "920",
			Name : "Kab. Wonosobo, Jawa Tengah",
		},
		{
			Code : "921",
			Name : "Kab. Purworejo, Jawa Tengah",
		},
		{
			Code : "922",
			Name : "Kab. Kebumen, Jawa Tengah",
		},
		{
			Code : "923",
			Name : "Kab. Klaten, Jawa Tengah",
		},
		{
			Code : "925",
			Name : "Kab. Sragen, Jawa Tengah",
		},
		{
			Code : "926",
			Name : "Kab. Sukoharjo, Jawa Tengah",
		},
		{
			Code : "927",
			Name : "Kab. Karanganyar, Jawa Tengah",
		},
		{
			Code : "928",
			Name : "Kab. Wonogiri, Jawa Tengah",
		},
		{
			Code : "929",
			Name : "Kab. Batang, Jawa Tengah",
		},
		{
			Code : "991",
			Name : "Kod. Semarang, Jawa Tengah",
		},
		{
			Code : "992",
			Name : "Kod. Salatiga, Jawa Tengah",
		},
		{
			Code : "993",
			Name : "Kod. Pekalongan, Jawa Tengah",
		},
		{
			Code : "994",
			Name : "Kod. Tegal, Jawa Tengah",
		},
		{
			Code : "995",
			Name : "Kod. Magelang, Jawa Tengah",
		},
		{
			Code : "996",
			Name : "Kod. Surakarta, Jawa Tengah",
		},
		{
			Code : "997",
			Name : "Kotif Klaten, Jawa Tengah",
		},
		{
			Code : "998",
			Name : "Kotif Cilacap, Jawa Tengah",
		},
		{
			Code : "999",
			Name : "Kotif Purwokerto, Jawa Tengah",
		},
		{
			Code : "9999",
			Name : "DI LUAR INDONESIA ( LUAR NEGERI )",
		},
	}

	var HubunganNasabahBank = []models.HubunganNasabahBank{
		{
			Code : "100",
			Name : "Perseorangan yg memiliki saham Bank >= 10% dari modal Bank",
			
		},
		{
			Code : "1000",
			Name : "Anak persh. Bank dengan kepemilikan 25% dari modal",
			
		},
		{
			Code : "200",
			Name : "Persh. yg memiliki saham Bank >= 10% dari modal Bank",
			
		},
		{
			Code : "2000",
			Name : "Tidak terkait dengan bank",
			
		},
		{
			Code : "300",
			Name : "Anggota dewan komisaris Bank",
			
		},
		{
			Code : "400",
			Name : "Anggota direksi Bank",
			
		},
		{
			Code : "501",
			Name : "Keluarga pemegang saham perorangan yang memiliki >= 10%",
			
		},
		{
			Code : "502",
			Name : "Keluarga anggota Dewan Komisaris Bank",
			
		},
		{
			Code : "503",
			Name : "Keluarga pemegang saham Anggota direksi Bank",
			
		},
		{
			Code : "600",
			Name : "Perseorangan memiliki >= 25% saham persh. yg memiliki >= 10%",
			
		},
		{
			Code : "700",
			Name : "Pejabat Bank yg mempunyai fungsi eksekutif",
			
		},
		{
			Code : "800",
			Name : "Persh. No.1 s/d 7 dengan kepemiliki 10% dari modal disetor",
			
		},
		{
			Code : "900",
			Name : "Persh. No.1 s/d 7 yang didalamnya terdapat pengaruh operasio",
			
		},
	}

	var HubunganKeluarga = []models.HubunganKeluarga{
		{
			Code : "1",
			Name : "Orang tua kandung/tiri/angkat",
		},
		{
			Code : "10",
			Name : "Saudara kandung/tiri/angkat dari orang tua",
		},
		{
			Code : "11",
			Name : "Mertua",
		},
		{
			Code : "2",
			Name : "Saudara kandung/tiri/angkat",
		},
		{
			Code : "3",
			Name : "Suami /Isteri",
		},
		{
			Code : "4",
			Name : "Anak kandung/tiri/angkat",
		},
		{
			Code : "5",
			Name : "Suami/isteri dari anak kandung/tiri/angkat",
		},
		{
			Code : "6",
			Name : "Kakek/nenek kandung/tiri/angkat",
		},
		{
			Code : "7",
			Name : "Cucu kandung/tiri/angkat",
		},
		{
			Code : "8",
			Name : "Saudara kandung/tiri/angkat dari suami/isteri",
		},
		{
			Code : "9",
			Name : "Suami/isteri dari saudara kandung/tiri/angkat",
		},
	}

	connection.DB.Save(&sectorEconomy1)
	connection.DB.Save(&sectorEconomy2)
	connection.DB.Save(&sectorEconomy3)
	connection.DB.Save(&sectorEconomyOjk)

	connection.DB.Save(&LokasiPabrik)
	connection.DB.Save(&LokasiDati2)
	connection.DB.Save(&HubunganNasabahBank)
	connection.DB.Save(&HubunganKeluarga,)

}

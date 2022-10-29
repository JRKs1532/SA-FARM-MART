package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("sa-65.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	database.AutoMigrate(
		&User{},
		&Gender{},
		&Position{},
		&Education{},
		&Employee{},

		&Province{},
		&Member{},

		&Typeproduct{},
		&Manufacturer{},
		&Product{},

		&Lot{},
		&Shelfproduct{},
		&Stock{},

		&Cart{},
		&Order{},

		&Paymenttype{},
		&Receipt{},
	)

	db = database

	password, err := bcrypt.GenerateFromPassword([]byte("123456"), 14)
	password2, err := bcrypt.GenerateFromPassword([]byte("56789"), 14)

	db.Model(&User{}).Create(&User{
		Name:     "Gunner",
		Email:    "gunner@email.com",
		Password: string(password),
	})
	var gunner User
	db.Raw("SELECT * FROM users WHERE email = ?", "gunner@email.com").Scan(&gunner)

	var target User
	db.Model(&User{}).Find(&target, db.Where("email = ?", "gunner@email.com"))

	//-----Gender
	male := Gender{
		Gender: "Male",
	}
	db.Model(&Gender{}).Create(&male)

	female := Gender{
		Gender: "Female",
	}
	db.Model(&Gender{}).Create(&female)

	//---Education
	bachelor := Education{
		Education: "Bachelor Degrees",
	}
	db.Model(&Education{}).Create(&bachelor)

	senior := Education{
		Education: "Senior High School",
	}
	db.Model(&Education{}).Create(&senior)

	junior := Education{
		Education: "Junior High School",
	}
	db.Model(&Education{}).Create(&junior)

	//---Position
	manager := Position{
		Position: "Manager",
	}
	db.Model(&Position{}).Create(&manager)

	salesperson := Position{
		Position: "Salesperson",
	}
	db.Model(&Position{}).Create(&salesperson)

	//-----Employee 1
	db.Model(&Employee{}).Create(&Employee{
		StartJob:  time.Now(),
		FirstName: "Bastien",
		LastName:  "Adel",
		Telephone: "0912345671",
		Email:     "bastien@email.com",
		Slary:     5000,
		Password:  string(password),
		User:      gunner,
		Gender:    male,
		Position:  manager,
		Education: bachelor,
	})

	//-----Employee 2
	db.Model(&Employee{}).Create(&Employee{
		StartJob:  time.Now(),
		FirstName: "Cosmo",
		LastName:  "Mable",
		Telephone: "0912345672",
		Email:     "cosmo@email.com",
		Slary:     9000,
		Password:  string(password2),
		User:      gunner,
		Gender:    male,
		Position:  salesperson,
		Education: junior,
	})

	var bastien Employee
	db.Raw("SELECT * FROM employees WHERE email = ?", "bastien@email.com").Scan(&bastien)
	var cosmo Employee
	db.Raw("SELECT * FROM employees WHERE email = ?", "cosmo@email.com").Scan(&cosmo)

	//=======================================================================
	//---Province Data
	bangkok := Province{
		Name: "Bangkok",
	}
	db.Model(&Province{}).Create(&bangkok)

	krabi := Province{
		Name: "Krabi",
	}
	db.Model(&Province{}).Create(&krabi)

	kanchanaburi := Province{
		Name: "Kanchanaburi",
	}
	db.Model(&Province{}).Create(&kanchanaburi)

	kalasin := Province{
		Name: "Kalasin",
	}
	db.Model(&Province{}).Create(&kalasin)

	kamphaeng_phet := Province{
		Name: "Kamphaeng Phet",
	}
	db.Model(&Province{}).Create(&kamphaeng_phet)

	khon_kaen := Province{
		Name: "Khon Kaen",
	}
	db.Model(&Province{}).Create(&khon_kaen)

	chanthaburi := Province{
		Name: "Chanthaburi",
	}
	db.Model(&Province{}).Create(&chanthaburi)

	chachoengsao := Province{
		Name: "Chachoengsao",
	}
	db.Model(&Province{}).Create(&chachoengsao)

	chonburi := Province{
		Name: "Chonburi",
	}
	db.Model(&Province{}).Create(&chonburi)

	chainat := Province{
		Name: "Chainat",
	}
	db.Model(&Province{}).Create(&chainat)

	chaiyaphum := Province{
		Name: "Chaiyaphum",
	}
	db.Model(&Province{}).Create(&chaiyaphum)

	chumphon := Province{
		Name: "Chumphon",
	}
	db.Model(&Province{}).Create(&chumphon)

	chiang_rai := Province{
		Name: "Chiang Rai",
	}
	db.Model(&Province{}).Create(&chiang_rai)

	chiang_mai := Province{
		Name: "Chiang Mai",
	}
	db.Model(&Province{}).Create(&chiang_mai)

	trang := Province{
		Name: "Trang",
	}
	db.Model(&Province{}).Create(&trang)

	trat := Province{
		Name: "Trat",
	}
	db.Model(&Province{}).Create(&trat)

	tak := Province{
		Name: "Tak",
	}
	db.Model(&Province{}).Create(&tak)

	nakhon_nayok := Province{
		Name: "Nakhon Nayok",
	}
	db.Model(&Province{}).Create(&nakhon_nayok)

	nakhon_pathom := Province{
		Name: "Nakhon Pathom",
	}
	db.Model(&Province{}).Create(&nakhon_pathom)

	nakhon_phanom := Province{
		Name: "Nakhon Phanom",
	}
	db.Model(&Province{}).Create(&nakhon_phanom)

	nakhon_ratchasima := Province{
		Name: "Nakhon Ratchasima",
	}
	db.Model(&Province{}).Create(&nakhon_ratchasima)

	nakhon_si_thammarat := Province{
		Name: "Nakhon Si Thammarat",
	}
	db.Model(&Province{}).Create(&nakhon_si_thammarat)

	nakhon_sawan := Province{
		Name: "Nakhon Sawan",
	}
	db.Model(&Province{}).Create(&nakhon_sawan)

	nonthaburi := Province{
		Name: "Nonthaburi",
	}
	db.Model(&Province{}).Create(&nonthaburi)

	narathiwat := Province{
		Name: "Narathiwat",
	}
	db.Model(&Province{}).Create(&narathiwat)

	nan := Province{
		Name: "Nan",
	}
	db.Model(&Province{}).Create(&nan)

	bueng_kan := Province{
		Name: "Bueng Kan",
	}
	db.Model(&Province{}).Create(&bueng_kan)

	buriram := Province{
		Name: "Buriram",
	}
	db.Model(&Province{}).Create(&buriram)

	pathum_thani := Province{
		Name: "Pathum Thani",
	}
	db.Model(&Province{}).Create(&pathum_thani)

	prachuap_khiri_khan := Province{
		Name: "Prachuap Khiri Khan",
	}
	db.Model(&Province{}).Create(&prachuap_khiri_khan)

	prachinburi := Province{
		Name: "Prachinburi",
	}
	db.Model(&Province{}).Create(&prachinburi)

	pattani := Province{
		Name: "Pattani",
	}
	db.Model(&Province{}).Create(&pattani)

	phra_nakhon_si_ayutthaya := Province{
		Name: "Phra Nakhon Si Ayutthaya",
	}
	db.Model(&Province{}).Create(&phra_nakhon_si_ayutthaya)

	phayao := Province{
		Name: "Phayao",
	}
	db.Model(&Province{}).Create(&phayao)

	phang_nga := Province{
		Name: "Phang Nga",
	}
	db.Model(&Province{}).Create(&phang_nga)

	phatthalung := Province{
		Name: "Phatthalung",
	}
	db.Model(&Province{}).Create(&phatthalung)

	phichit := Province{
		Name: "Phichit ",
	}
	db.Model(&Province{}).Create(&phichit)

	phitsanulok := Province{
		Name: "Phitsanulok ",
	}
	db.Model(&Province{}).Create(&phitsanulok)

	phetchaburi := Province{
		Name: "Phetchaburi",
	}
	db.Model(&Province{}).Create(&phetchaburi)

	phetchabun := Province{
		Name: "Phetchabun",
	}
	db.Model(&Province{}).Create(&phetchabun)

	phrae := Province{
		Name: "Phrae",
	}
	db.Model(&Province{}).Create(&phrae)

	phuket := Province{
		Name: "Phuket",
	}
	db.Model(&Province{}).Create(&phuket)

	maha_sarakham := Province{
		Name: "Maha Sarakham ",
	}
	db.Model(&Province{}).Create(&maha_sarakham)

	mukdahan := Province{
		Name: "Mukdahan",
	}
	db.Model(&Province{}).Create(&mukdahan)

	mae_hong_son := Province{
		Name: "Mae Hong Son",
	}
	db.Model(&Province{}).Create(&mae_hong_son)

	yasothon := Province{
		Name: "Yasothon",
	}
	db.Model(&Province{}).Create(&yasothon)

	yala := Province{
		Name: "Yala",
	}
	db.Model(&Province{}).Create(&yala)

	roi_et := Province{
		Name: "Roi Et",
	}
	db.Model(&Province{}).Create(&roi_et)

	ranong := Province{
		Name: "Ranong",
	}
	db.Model(&Province{}).Create(&ranong)

	rayong := Province{
		Name: "Rayong",
	}
	db.Model(&Province{}).Create(&rayong)

	ratchaburi := Province{
		Name: "Ratchaburi",
	}
	db.Model(&Province{}).Create(&ratchaburi)

	lopburi := Province{
		Name: "Lopburi",
	}
	db.Model(&Province{}).Create(&lopburi)

	lampang := Province{
		Name: "Lampang",
	}
	db.Model(&Province{}).Create(&lampang)

	lamphun := Province{
		Name: "Lamphun",
	}
	db.Model(&Province{}).Create(&lamphun)

	loei := Province{
		Name: "Loei",
	}
	db.Model(&Province{}).Create(&loei)

	sisaket := Province{
		Name: "Sisaket",
	}
	db.Model(&Province{}).Create(&sisaket)

	sakon_nakhon := Province{
		Name: "Sakon Nakhon",
	}
	db.Model(&Province{}).Create(&sakon_nakhon)

	songkhla := Province{
		Name: "Songkhla",
	}
	db.Model(&Province{}).Create(&songkhla)

	satun := Province{
		Name: "Stun",
	}
	db.Model(&Province{}).Create(&satun)

	samut_prakan := Province{
		Name: "Samut Prakan",
	}
	db.Model(&Province{}).Create(&samut_prakan)

	samut_sakhon := Province{
		Name: "Samut Sakhon",
	}
	db.Model(&Province{}).Create(&samut_sakhon)

	sa_kaeo := Province{
		Name: "Sa Kaeo",
	}
	db.Model(&Province{}).Create(&sa_kaeo)

	saraburi := Province{
		Name: "Saraburi",
	}
	db.Model(&Province{}).Create(&saraburi)

	sing_buri := Province{
		Name: "Sing Buri",
	}
	db.Model(&Province{}).Create(&sing_buri)

	sukhothai := Province{
		Name: "Sukhothai",
	}
	db.Model(&Province{}).Create(&sukhothai)

	suphan_buri := Province{
		Name: "Suphan Buri",
	}
	db.Model(&Province{}).Create(&suphan_buri)

	surat_thani := Province{
		Name: "Surat Thani",
	}
	db.Model(&Province{}).Create(&surat_thani)

	surin := Province{
		Name: "Surin",
	}
	db.Model(&Province{}).Create(&surin)

	nong_khai := Province{
		Name: "Nong Khai",
	}
	db.Model(&Province{}).Create(&nong_khai)

	nong_bua_lamphu := Province{
		Name: "Nong Bua Lamphu",
	}
	db.Model(&Province{}).Create(&nong_bua_lamphu)

	ang_thong := Province{
		Name: "Ang Thong",
	}
	db.Model(&Province{}).Create(&ang_thong)

	amnat_charoen := Province{
		Name: "Amnat Charoen",
	}
	db.Model(&Province{}).Create(&amnat_charoen)

	udon_thani := Province{
		Name: "Udon Thani",
	}
	db.Model(&Province{}).Create(&udon_thani)

	uttaradit := Province{
		Name: "Uttaradit",
	}
	db.Model(&Province{}).Create(&uttaradit)

	uthai_thani := Province{
		Name: "Uthai Thani",
	}
	db.Model(&Province{}).Create(&uthai_thani)

	ubon_ratchathani := Province{
		Name: "Ubon Ratchathani ",
	}
	db.Model(&Province{}).Create(&ubon_ratchathani)

	//---Member 1
	lisa := Member{
		FirstName:     "Lisa",
		LastName:      "Kim",
		Age:           25,
		Gender:        female,
		Date_Of_Birth: time.Date(1997, 3, 27, 0, 0, 0, 0, time.Now().UTC().Location()),
		Province:      sisaket,
		Telephone:     "0891341594",
		Employee:      bastien,
	}
	db.Model(&Member{}).Create(&lisa)

	//---Member 2
	tom := Member{
		FirstName:     "Tom",
		LastName:      "William",
		Age:           11,
		Gender:        male,
		Date_Of_Birth: time.Date(2011, 11, 11, 0, 0, 0, 0, time.Now().UTC().Location()),
		Province:      phra_nakhon_si_ayutthaya,
		Telephone:     "0985547553",
		Employee:      bastien,
	}
	db.Model(&Member{}).Create(&tom)

	// ----Member 3
	alice := Member{
		FirstName:     "Alice",
		LastName:      "Barber",
		Age:           23,
		Gender:        female,
		Date_Of_Birth: time.Date(1999, 6, 4, 0, 0, 0, 0, time.Now().UTC().Location()),
		Province:      sing_buri,
		Telephone:     "0785412335",
		Employee:      cosmo,
	}
	db.Model(&Member{}).Create(&alice)

	// ---Member 4
	bieber := Member{
		FirstName:     "Bieber",
		LastName:      "Hiddleson",
		Age:           28,
		Gender:        male,
		Date_Of_Birth: time.Date(1994, 6, 23, 0, 0, 0, 0, time.Now().UTC().Location()),
		Province:      nakhon_pathom,
		Telephone:     "0654252198",
		Employee:      cosmo,
	}
	db.Model(&Member{}).Create(&bieber)
	//=======================================================================
	// --- TypeProduct Data
	drink := Typeproduct{
		Typeproduct_name: "Drink",
	}
	db.Model(&Typeproduct{}).Create(&drink)

	snack := Typeproduct{
		Typeproduct_name: "Snack",
	}
	db.Model(&Typeproduct{}).Create(&snack)

	softener := Typeproduct{
		Typeproduct_name: "Softener",
	}
	db.Model(&Typeproduct{}).Create(&softener)

	// Manufacturer Data
	nestle := Manufacturer{
		Manufacturer_name: "NESTLE",
	}
	db.Model(&Manufacturer{}).Create(&nestle)

	lotte := Manufacturer{
		Manufacturer_name: "LOTTE",
	}
	db.Model(&Manufacturer{}).Create(&lotte)

	unilever := Manufacturer{
		Manufacturer_name: "UNILEVER",
	}
	db.Model(&Manufacturer{}).Create(&unilever)

	// Product 1
	purelife := Product{
		Product_name:  "Pure Life",
		Product_price: 10.00,
		Typeproduct:   drink,
		Manufacturer:  nestle,
		Employee:      bastien,
	}
	db.Model(&Product{}).Create(&purelife)

	// Product 2
	chocopie := Product{
		Product_name:  "Choco Pie",
		Product_price: 5.00,
		Typeproduct:   snack,
		Manufacturer:  lotte,
		Employee:      cosmo,
	}
	db.Model(&Product{}).Create(&chocopie)

	// Product 3
	comfort := Product{
		Product_name:  "Comfort",
		Product_price: 47.50,
		Typeproduct:   softener,
		Manufacturer:  unilever,
		Employee:      cosmo,
	}
	db.Model(&Product{}).Create(&comfort)

	//===============================================
	// Lot Data
	lot_1 := Lot{
		Expired: time.Now(),
	}
	db.Model(&Lot{}).Create(&lot_1)

	lot_2 := Lot{
		Expired: time.Now(),
	}
	db.Model(&Lot{}).Create(&lot_2)

	lot_3 := Lot{
		Expired: time.Now(),
	}
	db.Model(&Lot{}).Create(&lot_3)

	// Shelf Data
	shelf_A := Shelfproduct{
		Shelf_name: "Shelf_A",
	}
	db.Model(&Shelfproduct{}).Create(&shelf_A)

	shelf_B := Shelfproduct{
		Shelf_name: "Shelf_B",
	}
	db.Model(&Shelfproduct{}).Create(&shelf_B)

	shelf_C := Shelfproduct{
		Shelf_name: "Shelf_C",
	}
	db.Model(&Shelfproduct{}).Create(&shelf_C)

	// stock 1
	stock_PureLife := Stock{
		Product:        purelife,
		Stock_quantity: 200,
		Lot:            lot_1,
		Shelfproduct:   shelf_B,
		Employee:       cosmo,
	}
	db.Model(&Stock{}).Create(&stock_PureLife)

	// stock 2
	stock_Chocopie := Stock{
		Product:        chocopie,
		Stock_quantity: 300,
		Lot:            lot_3,
		Shelfproduct:   shelf_C,
		Employee:       cosmo,
	}
	db.Model(&Stock{}).Create(&stock_Chocopie)

	// stock 3
	stock_Comfort := Stock{
		Product:        comfort,
		Stock_quantity: 400,
		Lot:            lot_2,
		Shelfproduct:   shelf_A,
		Employee:       bastien,
	}
	db.Model(&Stock{}).Create(&stock_Comfort)

	//==============================================
	// Cart 1
	cart1 := Cart{
		Employee: cosmo,
		Member:   lisa,
	}
	db.Model(&Cart{}).Create(&cart1)

	// Cart 2
	cart2 := Cart{
		Employee: cosmo,
		Member:   alice,
	}
	db.Model(&Cart{}).Create(&cart2)

	// Cart 3
	cart3 := Cart{
		Employee: bastien,
		Member:   alice,
	}
	db.Model(&Cart{}).Create(&cart3)

	// Order 1
	order1 := Order{
		Product_quantity: 2,
		Product:          chocopie,
		Cart:             cart1,
	}
	db.Model(&Order{}).Create(&order1)

	// Order 2
	order2 := Order{
		Product_quantity: 5,
		Product:          purelife,
		Cart:             cart3,
	}
	db.Model(&Order{}).Create(&order2)

	// Order 3
	order3 := Order{
		Product_quantity: 1,
		Product:          comfort,
		Cart:             cart3,
	}
	db.Model(&Order{}).Create(&order3)

	//==============================================
	//payment_type 1
	cash := Paymenttype{
		Paymenttype: "Cash",
	}
	db.Model(&Paymenttype{}).Create(&cash)

	//payment_type 2
	banking := Paymenttype{
		Paymenttype: "Banking",
	}
	db.Model(&Paymenttype{}).Create(&banking)

	//payment_type 3
	promptPay := Paymenttype{
		Paymenttype: "PromptPay",
	}
	db.Model(&Paymenttype{}).Create(&promptPay)

	//payment_type 4
	credit := Paymenttype{
		Paymenttype: "Debit/Credit Card",
	}
	db.Model(&Paymenttype{}).Create(&credit)

	//Receipt
	db.Model(&Receipt{}).Create(&Receipt{
		ReceiptTime:          time.Now(),
		ReceiptSum:           145.00,
		ReceiptPaymentAmount: 150.00,
		Paymenttype:          cash,
		Cart:                 cart1,
		Employee:             cosmo,
		Member:               lisa,
	})
}

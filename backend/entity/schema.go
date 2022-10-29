package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string 
	Password string
	Employee []Employee `gorm:"foreignKey:UserID"`
}

type Gender struct {
	gorm.Model
	Gender   string
	Employee []Employee `gorm:"foreignKey:GenderID"`
	Member   []Member   `gorm:"foreignKey:GenderID"`
}

type Position struct {
	gorm.Model
	Position string
	Employee []Employee `gorm:"foreignKey:PositionID"`
}

type Education struct {
	gorm.Model
	Education string
	Employee  []Employee `gorm:"foreignKey:EducationID"`
}

type Employee struct {
	gorm.Model
	StartJob    time.Time
	FirstName   string
	LastName    string
	Telephone   string
	Email       string 
	Slary       int
	Password    string `json:"password"`
	UserID      *uint
	User        User `gorm:"references:id"`
	GenderID    *uint
	Gender      Gender `gorm:"references:id"`
	PositionID  *uint
	Position    Position `gorm:"references:id"`
	EducationID *uint
	Education   Education `gorm:"references:id"`
}

// --------------------------------------------------------------
type Province struct {
	gorm.Model
	Name    string
	Members []Member `gorm:"foreignKey:ProvinceID"`
}

type Member struct {
	gorm.Model
	FirstName string
	LastName  string
	Age       int

	GenderID *uint
	Gender   Gender `gorm:"references:id"`

	Date_Of_Birth time.Time

	ProvinceID *uint
	Province   Province `gorm:"references:id"`

	Telephone string `gorm:"UniqueIndex"`

	EmployeeID *uint
	Employee   Employee `gorm:"references:id"`
}

// -------------------------------------------------------------
type Typeproduct struct {
	gorm.Model
	Typeproduct_name string
	Products         []Product `gorm:"foreignKey:TypeproductID"`
}

type Manufacturer struct {
	gorm.Model
	Manufacturer_name string
	Products          []Product `gorm:"foreignKey:ManufacturerID"`
}

type Product struct {
	gorm.Model
	Product_name   string
	Product_price  float32
	TypeproductID  *uint
	Typeproduct    Typeproduct `gorm:"references:id"`
	ManufacturerID *uint
	Manufacturer   Manufacturer `gorm:"references:id"`
	EmployeeID     *uint
	Employee       Employee `gorm:"references:id"`
}

// -------------------------------------------------------------------
type Lot struct {
	gorm.Model
	Expired time.Time
	Stock   []Stock `gorm:"foreignKey:LotID"`
}

type Shelfproduct struct {
	gorm.Model
	Shelf_name string
	Stock      []Stock `gorm:"foreignKey:ShelfproductID"`
}

type Stock struct {
	gorm.Model
	ProductID      *uint
	Product        Product `gorm:"references:id"`
	Stock_quantity float32
	LotID          *uint
	Lot            Lot `gorm:"references:id"`
	ShelfproductID *uint
	Shelfproduct   Shelfproduct `gorm:"references:id"`
	EmployeeID     *uint
	Employee       Employee `gorm:"references:id"`
}

// ============================================================
type Cart struct {
	gorm.Model
	EmployeeID *uint
	Employee   Employee `gorm:"references:id"`
	MemberID   *uint
	Member     Member    `gorm:"references:id"`
	Order      []Order   `gorm:"foreignKey:CartID"`
	Receipt    []Receipt `gorm:"foreignKey:CartID"`
}

type Order struct {
	gorm.Model
	ProductID        *uint
	Product          Product `gorm:"references:id"`
	Product_quantity int8
	CartID           *uint
	Cart             Cart `gorm:"references:id"`
}

// ============================================================
type Paymenttype struct {
	gorm.Model
	Paymenttype string
	Receipt     []Receipt `gorm:"foreignKey:PaymenttypeID"`
}

// ตารางหลักของ ระบบชำระเงิน
type Receipt struct {
	gorm.Model
	ReceiptTime          time.Time
	ReceiptSum           float32
	ReceiptPaymentAmount float32

	PaymenttypeID *uint
	Paymenttype   Paymenttype `gorm:"references:id"`

	CartID *uint
	Cart   Cart `gorm:"references:id"`

	EmployeeID *uint
	Employee   Employee `gorm:"references:id"`

	MemberID *uint
	Member   Member `gorm:"references:id"`
}

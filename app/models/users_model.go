package models

import "time"

/**
 * ユーザー情報の構造体
 **/
type Users struct {
	EmployeeId        string    `xorm:"employee_id"`          /** 社員ID */
	UserLastName      string    `xorm:"user_last_name"`       /** ユーザー氏名（氏） */
	UserFirstName     string    `xorm:"user_first_name"`      /** ユーザー氏名（名） */
	UserLastNameKana  string    `xorm:"user_last_name_kana"`  /** ユーザー氏名カナ（氏） */
	UserFirstNamaKana string    `xorm:"user_first_nama_kana"` /** ユーザー氏名カナ（名） */
	DateOfJoin        string    `xorm:"date_of_join"`         /** 入社年月 */
	BirthDay          string    `xorm:"birth_day"`            /** 生年月日 */
	Age               int       `xorm:"age"`                  /** 年齢 */
	Gender            int       `xorm:"gender"`               /** 性別 */
	PhoneNumber       string    `xorm:"phone_number"`         /** 電話番号 */
	MailAddress       string    `xorm:"mail_address"`         /** メールアドレス */
	Zipcode           string    `xorm:"zipcode"`              /** 郵便番号 */
	Prefcode          string    `xorm:"prefcode"`             /** 都道府県コード */
	Prefecture        string    `xorm:"prefecture"`           /** 都道府県 */
	Municipalities    string    `xorm:"municipalities"`       /** 市区町村 */
	Building          string    `xorm:"building"`             /** 建物 */
	CreatedAuthor     string    `xorm:"created_author"`       /** 作成者 */
	CreatedDate       time.Time `xorm:"created_date"`         /** 作成日時 */
	UpdatedAuthor     string    `xorm:"updated_author"`       /** 更新者 */
	UpdatedDate       time.Time `xorm:"updated_date"`         /** 更新日時 */
}

/**
 * ユーザー情報のリクエスト
 */
type RequestUsers struct {
	UserLastName      string `json:"userLastName" binding:"required,max=30" xorm:"user_last_name"`   /** ユーザー氏名（氏） */
	UserFirstName     string `json:"userFirstName" binding:"required,max=30" xorm:"user_first_name"` /** ユーザー氏名（名） */
	UserLastNameKana  string `json:"userLastNameKana" binding:"max=30" xorm:"user_last_name_kana"`   /** ユーザー氏名カナ（氏） */
	UserFirstNamaKana string `json:"userFirstNamaKana" binding:"max=30" xorm:"user_first_nama_kana"` /** ユーザー氏名カナ（名） */
	DateOfJoin        string `json:"dateOfJoin" xorm:"date_of_join"`                                 /** 入社年月 */
	Age               int    `json:"age" xorm:"age"`                                                 /** 年齢 */
	Gender            int    `json:"gender" xorm:"gender"`                                           /** 性別 */
	PhoneNumber       string `json:"phoneNumber" binding:"max=21" xorm:"phone_number"`               /** 電話番号 */
	MailAddress       string `json:"mailAddress" binding:"max=254,email" xorm:"mail_address"`        /** メールアドレス */
	Zipcode           string `json:"zipcode" binding:"len=7" xorm:"zipcode"`                         /** 郵便番号 */
	Prefcode          string `json:"prefcode" binding:"len=2" xorm:"prefcode"`                       /** 都道府県コード */
	Prefecture        string `json:"prefecture" binding:"min=3,max=4" xorm:"prefecture"`             /** 都道府県 */
	Municipalities    string `json:"municipalities" binding:"max=100" xorm:"municipalities"`         /** 市区町村 */
	Building          string `json:"building" binding:"max=100" xorm:"building"`                     /** 建物 */
}

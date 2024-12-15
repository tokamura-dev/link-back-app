package authmodel

import "time"

/**
 * ログイン情報の構造体
 **/
type Login struct {
	EmployeeId      string     `xorm:"employee_id"`      /** 社員ID */
	Password        string     `xorm:"password"`         /** パスワード */
	DeleteFlg       int        `xorm:"delete_flg"`       /** 論理削除フラグ */
	CreatedAuthor   string     `xorm:"created_author"`   /** 作成者 */
	CreatedDatetime time.Time  `xorm:"created_datetime"` /** 作成日時 */
	UpdatedAuthor   string     `xorm:"updated_author"`   /** 更新者 */
	UpdatedDatetime *time.Time `xorm:"updated_datetime"` /** 更新日時 */
}

/**
 * サインアップのリクエスト
 **/
type RequestSignUp struct {
	Password          string `json:"password" binding:"required,max=30" xorm:"password"`             /** パスワード */
	UserLastName      string `json:"userLastName" binding:"required,max=30" xorm:"user_last_name"`   /** ユーザー氏名（氏） */
	UserFirstName     string `json:"userFirstName" binding:"required,max=30" xorm:"user_first_name"` /** ユーザー氏名（名） */
	UserLastNameKana  string `json:"userLastNameKana" binding:"max=30" xorm:"user_last_name_kana"`   /** ユーザー氏名カナ（氏） */
	UserFirstNamaKana string `json:"userFirstNamaKana" binding:"max=30" xorm:"user_first_nama_kana"` /** ユーザー氏名カナ（名） */
	DateOfJoin        string `json:"dateOfJoin" xorm:"date_of_join"`                                 /** 入社年月 */
	BirthDay          string `json:"birthDay" xorm:"birth_day"`
	Age               int    `json:"age" xorm:"age"`                                          /** 年齢 */
	Gender            int    `json:"gender" xorm:"gender"`                                    /** 性別 */
	PhoneNumber       string `json:"phoneNumber" binding:"max=21" xorm:"phone_number"`        /** 電話番号 */
	MailAddress       string `json:"mailAddress" binding:"max=254,email" xorm:"mail_address"` /** メールアドレス */
	Zipcode           string `json:"zipcode" binding:"len=7" xorm:"zipcode"`                  /** 郵便番号 */
	Prefcode          string `json:"prefcode" binding:"len=2" xorm:"prefcode"`                /** 都道府県コード */
	Prefecture        string `json:"prefecture" binding:"min=3,max=4" xorm:"prefecture"`      /** 都道府県 */
	Municipalities    string `json:"municipalities" binding:"max=100" xorm:"municipalities"`  /** 市区町村 */
	Building          string `json:"building" binding:"max=100" xorm:"building"`
}

/**
 * サインインのリクエスト
 **/
type RequstSignIn struct {
	EmployeeId string `json:"employeeId" binding:"required,max=30" xorm:"employee_id"` /** 社員ID */
	Password   string `json:"password" binding:"required,max=30" xorm:"password"`      /** パスワード */
}

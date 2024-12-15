package cryptoutil

import "golang.org/x/crypto/bcrypt"

/**
 * パスワードを暗号化する処理
 **/
func PasswordEncrypt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

/**
 * 暗号化されたパスワードと入力された平文のパスワードを比較する処理
 **/
func CompareHashAndPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

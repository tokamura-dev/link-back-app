package enum

type Div struct {
	Code  int
	Value string
}

/**
 * 論理削除区分
 * 0：未削除
 * 1：削除済
 **/
var LogicalDeleteDiv = struct {
	NotDeleted Div
	Deleted    Div
}{
	NotDeleted: Div{Code: 0, Value: "未削除"},
	Deleted:    Div{Code: 1, Value: "削除済"},
}

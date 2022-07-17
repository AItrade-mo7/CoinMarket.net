package okxInfo

type InstInfoType struct {
	Alias     string `bson:"Alias"`
	BaseCcy   string `bson:"BaseCcy"`
	Category  string `bson:"Category"`
	CtMult    string `bson:"CtMult"`
	CtType    string `bson:"CtType"`
	CtVal     string `bson:"CtVal"`
	CtValCcy  string `bson:"CtValCcy"`
	ExpTime   string `bson:"ExpTime"`
	InstID    string `bson:"InstID"`
	InstType  string `bson:"InstType"` // SWAP   SPOT
	Lever     string `bson:"Lever"`
	ListTime  string `bson:"ListTime"`
	LotSz     string `bson:"LotSz"`
	MinSz     string `bson:"MinSz"`
	OptType   string `bson:"OptType"`
	QuoteCcy  string `bson:"QuoteCcy"`
	SettleCcy string `bson:"SettleCcy"`
	State     string `bson:"State"`
	Stk       string `bson:"Stk"`
	TickSz    string `bson:"TickSz"`
	Uly       string `bson:"Uly"`
}

type InstListType struct {
	SWAP []InstInfoType
	SPOT []InstInfoType
}

// 当前产品频道
var InstList InstListType

// 预上线产品
var PreopenInstList InstListType

func ClearInst(instType string) {
	if instType == "SWAP" {
		InstList.SWAP = []InstInfoType{}
		PreopenInstList.SWAP = []InstInfoType{}
	}

	if instType == "SPOT" {
		InstList.SPOT = []InstInfoType{}
		PreopenInstList.SPOT = []InstInfoType{}
	}
}

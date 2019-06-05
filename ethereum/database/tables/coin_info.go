package tables

import "fmt"

type TableCoinInfo struct {
	Id          uint64  `json:"id"                gorm:"column:id;primary_key;AUTO_INCREMENT"`       //自增主键
	NameId      string  `json:"name_id"           gorm:"column:name_id;type:varchar(20)"`            //代币全名
	Name        string  `json:"name"              gorm:"column:name;type:varchar(20)"`               //代币全名
	Symbol      string  `json:"symbol"            gorm:"column:symbol;type:varchar(10)"`             //代币符号
	Decimals    uint8   `json:"decimals"          gorm:"column:decimals"`                            //代币位数
	Contract    string  `json:"contract"          gorm:"column:contract;type:char(20);unique_index"` //代币合约地址
	Transaction string  `json:"transaction"       gorm:"column:tx_hash;type:char(64)"`               //创建合约的交易哈希
	Supply      string  `json:"supply"            gorm:"column:supply"`                              //代币发行总量
	Price       float64 `json:"price"             gorm:"column:price"`                               //代币价格，单位是美元
	ICon        string  `json:"icon"              gorm:"column:icon"`                                //代币图标链接
}

func (t *TableCoinInfo) TableName() string {
	return "t_coins_info"
}

func (t *TableCoinInfo) SqlCommand() string {
	return fmt.Sprintf("REPLACE INTO `t_coins_info`(`name_id`,`name`,`symbol`,`decimals`,`contract`,`tx_hash`,`supply`,`price`,`icon`) "+
		"VALUES('%s','%s','%s','%d','%s','%s','%s','%f','%s') ",
		MysqlFormat(t.NameId), MysqlFormat(t.Name), MysqlFormat(t.Symbol), t.Decimals, t.Contract, t.Transaction, t.Supply, t.Price, MysqlFormat(t.ICon))
}

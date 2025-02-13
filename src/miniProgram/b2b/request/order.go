package request

type GetOrderReq struct {
	MchID      string `json:"mchid"` // 微信 b2b 支付配的商户号，和商户订单号配合使用
	OutTradeNo string `json:"out_trade_no,omitempty"`
	OrderId    string `json:"order_id,omitempty"`
}

type RefundReq struct {
	MchID        string `json:"mchid"`                   // 微信商户号
	OutTradeNo   string `json:"out_trade_no,omitempty"`  // 商户订单号
	OrderID      string `json:"order_id,omitempty"`      // B2b支付订单号
	OutRefundNo  string `json:"out_refund_no"`           // 商户退款单号
	RefundAmount int64  `json:"refund_amount"`           // 退款金额，单位为分，只能为整数，不能超过原订单支付金额。
	RefundFrom   int    `json:"refund_from"`             // 退款来源，枚举值 1：人工客服退款 2：用户自己退款 3：其他
	RefundReason int    `json:"refund_reason,omitempty"` // 退款原因，枚举值 0：暂无描述 1：产品问题 2：售后问题 3：意愿问题 4：价格问题 5：其他原因
	AppKey       string `json:"-" xml:"-"`               // 唯独这个接口一定要现网 appKey，覆盖计算
}

type GetRefundReq struct {
	MchID       string `json:"mchid"`                   // 微信商户号
	OutRefundNo string `json:"out_refund_no,omitempty"` // 商户退款单号
	RefundID    string `json:"refund_id,omitempty"`     // B2b支付退款单号
}

type AddProfitSharingAccountReq struct {
	ProfitSharingRelationType string `json:"profit_sharing_relation_type"` // 分账接收方关系类型
	PayeeType                 string `json:"payee_type"`                   // 分账接收方类型
	PayeeID                   string `json:"payee_id"`                     // 分账接收方标识
	PayeeName                 string `json:"payee_name,omitempty"`         // 分账接收方名称
	Env                       int    `json:"env"`                          // 环境类型
}

type DelProfitSharingAccountReq struct {
	PayeeType string `json:"payee_type"` // 分账接收方类型
	PayeeID   string `json:"payee_id"`   // 分账接收方标识
	Env       int    `json:"env"`        // 环境类型
}

type QueryProfitSharingAccountReq struct {
	Offset int `json:"offset"` // 偏移量
	Limit  int `json:"limit"`  // 查询数量
}

type CreateProfitSharingOrderReq struct {
	MchID           string `json:"mchid"`            // 子商户id
	OutTradeNo      string `json:"out_trade_no"`     // 支付单id
	ProfitFee       int64  `json:"profit_fee"`       // 分账费用, 单位:分，不超过支付单本身的金额
	ReceiverType    string `json:"receiver_type"`    // 分账接收方类型
	ReceiverAccount string `json:"receiver_account"` // 分账接收方账号
}

type QueryProfitSharingOrderReq struct {
	MchID           string `json:"mchid"`            // 商户号，必填
	OutTradeNo      string `json:"out_trade_no"`     // 支付单 id，必填
	ReceiverType    string `json:"receiver_type"`    // 分账接收方类型，必填
	ReceiverAccount string `json:"receiver_account"` // 分账接收方账号，必填
}

type QueryProfitSharingRemainAmtReq struct {
	MchID      string `json:"mchid"`        // 子商户id，必填
	OutTradeNo string `json:"out_trade_no"` // 订单 id，必填
	Env        int    `json:"env"`          // 环境参数，0标识正式（默认值），1 标识沙盒，否
}

type FinishProfitSharingOrderReq struct {
	MchID      string `json:"mchid"`        // 子商户id，必填
	OutTradeNo string `json:"out_trade_no"` // 订单 id，必填
}

type RefundProfitSharingReq struct {
	MchID       string `json:"mchid"`         // 商户号 id，必填
	OutTradeNo  string `json:"out_trade_no"`  // 订单 id，必填
	OutRefundNo string `json:"out_refund_no"` // 退款单 id，必填
	PayeeType   string `json:"payee_type"`    // 退款分账方类型，必填
	PayeeID     string `json:"payee_id"`      // 退款分账方 id，必填
	RefundAmt   int64  `json:"refund_amt"`    // 退款金额，必填
}

type QueryRefundProfitSharingOrderReq struct {
	MchID       string `json:"mchid"`         // 商户号 id，必填
	OutTradeNo  string `json:"out_trade_no"`  // 订单 id，必填
	OutRefundNo string `json:"out_refund_no"` // 退款单 id，必填
	PayeeType   string `json:"payee_type"`    // 退款分账方类型，必填
	PayeeID     string `json:"payee_id"`      // 退款分账方 id，必填
}

type DownloadBillReq struct {
	MchID    string `json:"mchid"`     // 商户号，必填
	BillDate string `json:"bill_date"` // 账单日期，格式:yyyymmdd，必填
}

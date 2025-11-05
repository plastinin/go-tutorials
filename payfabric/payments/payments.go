package payments

type PaymentMethod interface {
	Pay(usd int) int
	Cancel(id int)
}

type PaymentModule struct {
	paymentsInfo  map[int]PaymentInfo
	paymentMethod PaymentMethod
}

func NewPaymentModule(paymentMethod PaymentMethod) *PaymentModule {
	return &PaymentModule{
		paymentMethod: paymentMethod,
		paymentsInfo:  make(map[int]PaymentInfo),
	}
}

// Принимает:
//  1. Описание проводимой оплаты
//  2. Сумму оплаты
//
// Возвращает:
//
//	ID проведенной оплаты
func (p *PaymentModule) Pay(description string, usd int) int {

	id := p.paymentMethod.Pay(usd)

	info := PaymentInfo{
		Description: description,
		Usd:         usd,
		Cancelled:   false,
	}

	p.paymentsInfo[id] = info

	return id
}

// Принимает:
//  1. ID проведенной оплаты
//
// Возвращает:
//
//	Ничего
func (p *PaymentModule) Cancel(id int) {
	p.paymentMethod.Cancel(id)
	info, ok := p.paymentsInfo[id]
	if ok {
		info.Cancelled = true
		p.paymentsInfo[id] = info
	}
}

// Принимает:
//  1. ID проведенной оплаты
//
// Возвращает:
//
//	Информация о проведенной информации
func (p *PaymentModule) Info(id int) PaymentInfo {
	info, ok := p.paymentsInfo[id]
	if !ok {
		return PaymentInfo{}
	}
	return info
}

// Принимает:
//
//	Ничего
//
// Возвращает:
//
//	Информацию о всех операциях
func (p *PaymentModule) AllInfo() map[int]PaymentInfo {
	tempMap := make(map[int]PaymentInfo, len(p.paymentsInfo))
	for k, value := range p.paymentsInfo {
		tempMap[k] = value
	}
	return tempMap
}
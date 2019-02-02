package privat

import "encoding/xml"



func (api *Privat24Api) GetBalance(cardNumber string) {
	url := api.apiUrl + "/balance"

	data:=new(Data)
	items:= make([]interface{}, 4)

	paymentProp:=make([]Prop, 2)
	paymentProp=append(paymentProp, Prop{Name:"cardnum", Value:cardNumber}, Prop{Name:"country", Value:"UA"})

	payment:=new(Payment)
	payment.Properties=paymentProp

	items=append(items, new(Oper), new(Wait), new(Test), payment)

	data.Items=items

	queryXML:=new(QueryXML)
}

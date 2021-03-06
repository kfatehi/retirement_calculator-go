package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"retirement_calculator-go/analytics"
	"retirement_calculator-go/retcalc"
)

func perror(err error) {
	if err != nil {
		panic(err)
	}
}

func GetResponse() {
	url := "http://127.0.0.1:8080/recalc/"
	res, err := http.Get(url)
	perror(err)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	perror(err)
	var r retcalc.RetCalc

	err = json.Unmarshal(body, &r)

}

func main() {
	JsonObj := []byte(`{"Age":22, "Retirement_age":65, "Terminal_age":90, "Effective_tax_rate":0.3, "Returns_tax_rate":0.3, "N": 20000, 
						"Non_Taxable_contribution":17500, "Taxable_contribution": 0, "Non_Taxable_balance":0, "Taxable_balance": 0, 
						"Yearly_social_security_income":0, "Asset_volatility": 0.15, "Expected_rate_of_return": 0.07, "Inflation_rate":0.035}`)
	rc := retcalc.NewRetCalcFromJSON(JsonObj)
	a := rc.RunIncomes()
	fmt.Println(analytics.HistoCumulative(a, 200))
}

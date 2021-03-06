package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"retirement_calculator-go/analytics"
	"retirement_calculator-go/retcalc"

	"github.com/gorilla/mux"
)

const AllDataPrefix = "/alldata/"
const IncomesPrefix = "/incomes/"
const PathPrefix = "/paths/"
const InputPrefix = "/input/"

func RegisterHandlers() {
	r := mux.NewRouter()
	r.HandleFunc(InputPrefix, error_handler(RecalcFromWebInput)).Methods("POST")
	r.HandleFunc(PathPrefix, error_handler(SinglePath)).Methods("GET")
	r.HandleFunc(IncomesPrefix, error_handler(IncomesJSON)).Methods("GET")
	r.HandleFunc(AllDataPrefix, error_handler(Retcalc_basic)).Methods("GET")
	//r.HandleFunc(AllDataPrefix, error_handler(Retcalc_user_input)).Methods("POST")
	http.Handle(InputPrefix, r)
	http.Handle(AllDataPrefix, r)
	http.Handle(IncomesPrefix, r)
	http.Handle(PathPrefix, r)
}

// badRequest is handled by setting the status code in the reply to StatusBadRequest.
type badRequest struct{ error }

// notFound is handled by setting the status code in the reply to StatusNotFound.
type notFound struct{ error }

func error_handler(f func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err == nil {
			return
		}
		switch err.(type) {
		case badRequest:
			http.Error(w, err.Error(), http.StatusBadRequest)
		case notFound:
			http.Error(w, "task not found", http.StatusNotFound)
		default:
			log.Println(err)
			http.Error(w, "oops", http.StatusInternalServerError)
		}
	}
}

func RecalcFromWebInput(w http.ResponseWriter, r *http.Request) error {
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)

	myRetCalc := retcalc.NewRetCalcFromJSON(body)

	fmt.Println()
	fmt.Println("POST request recieved - RecalcFromWebInput()")
	fmt.Printf("Recived: %s\n", string(body))
	fmt.Println()
	incs := myRetCalc.RunIncomes()
	fmt.Println("RunIncomes() analystics:")
	fmt.Printf("Max: %f\n", analytics.MaxF64(incs))
	fmt.Printf("Min: %f\n", analytics.MinF64(incs))
	fmt.Printf("Avg: %f\n", analytics.AvgF64(incs))
	myRetCalc.ShowRetCalc()

	//return json.NewEncoder(w).Encode(retcalc.HistoFromSlice(myRetCalc.RunIncomes()))
	return json.NewEncoder(w).Encode(analytics.HistoCumulative(myRetCalc.RunIncomes(), 250))
}

func Retcalc_basic(w http.ResponseWriter, r *http.Request) error {
	rc := retcalc.NewRetCalc()
	return json.NewEncoder(w).Encode(rc)
}

/*
func Retcalc_user_input(w http.ResponseWriter, r *http.Request) error {
	req := retcalc.RetCalcWebInput{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		panic(err)
	}
	return json.NewEncoder(w).Encode(req)
}
*/
func IncomesJSON(w http.ResponseWriter, r *http.Request) error {
	rc := retcalc.NewRetCalc()
	return json.NewEncoder(w).Encode(retcalc.HistoFromSlice(rc.RunIncomes()))
}

func SinglePath(w http.ResponseWriter, r *http.Request) error {
	rc := retcalc.NewRetCalc()
	return json.NewEncoder(w).Encode(rc.PercentilePath(0.25))
}

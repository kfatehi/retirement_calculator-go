/*

main.go

My retirement calculator in Go

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.

Copyright 2014 Johnnydiabetic
*/

package main

import (
	"fmt"
	"net/http"
	"retirement_calculator-go/server"
	"strconv"
)

const listenPort = 8081

func main() {
	listenStr := ":" + strconv.Itoa(listenPort)
	server.RegisterHandlers()
	http.Handle("/", http.FileServer(http.Dir("./static")))
	fmt.Printf("RetCalc server on: Listening on port %d\n", listenPort)
	http.ListenAndServe(listenStr, nil)
}

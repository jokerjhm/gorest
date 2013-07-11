/*
Copyright Siyabonga Dlamini (siyabonga.dlamini@gmail.com).
All rights reserved.

Modified by HourlyHost, Inc.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions
are met:

 1. Redistributions of source code must retain the above copyright
    notice, this list of conditions and the following disclaimer.

 2. Redistributions in binary form must reproduce the above copyright
    notice, this list of conditions and the following disclaimer
    in the documentation and/or other materials provided with the
    distribution.

THIS SOFTWARE IS PROVIDED BY THE AUTHOR "AS IS" AND ANY EXPRESS OR
IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES
OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.
IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO,
PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS;
OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY,
WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR
OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF
ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

package gorest

import (
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

var MUX_ROOT = "/home/now/the/future/"
var RootPath = "http://localhost:8787" + MUX_ROOT
var globalTestScope *testing.T //This is used to do Asserts inside the service implementations

func TestInit(t *testing.T) {
	globalTestScope = t
	log.Println("Starting tests")
	log.SetOutput(ioutil.Discard) //Toggle comment in-out to see log output

	RegisterRealmAuthorizer("testing", TestingAuthorizer)
	RegisterServiceOnPath(MUX_ROOT, new(TypesService))
	RegisterServiceOnPath(MUX_ROOT, new(PathsService))
	http.Handle(MUX_ROOT, Handle())

	go http.ListenAndServe(":8787", nil)
}

func TestDataTransmition(t *testing.T) {
	testTypes(t)
}

func TestPaths(t *testing.T) {
	testPaths(t)
}

package transport

/*
	Sliver Implant Framework
	Copyright (C) 2021  Bishop Fox

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

import (
	"encoding/pem"
	"testing"

	"github.com/bishopfox/sliver/server/certs"
)

func TestMTLSCertificateValidation(t *testing.T) {
	certs.GenerateCertificateAuthority(certs.OperatorCA, "")
	opCert, _, err := certs.OperatorClientGenerateCertificate("mtls-test")
	if err != nil {
		t.Errorf("Failed to store ecc certificate %s", err)
		return
	}
	opCACert, _, err := certs.GetCertificateAuthorityPEM(certs.OperatorCA)
	if err != nil {
		t.Errorf("Failed to find operator CA %s", err)
		return
	}

	block, _ := pem.Decode(opCert)
	err = rootOnlyVerifyCertificate(string(opCACert), [][]byte{block.Bytes})
	if err != nil {
		t.Errorf("Operator certificate validation failed %s", err)
		return
	}
}

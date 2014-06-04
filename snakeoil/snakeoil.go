// Package snakeoil provides a self-signed certificate for serving via HTTPS for
// test purposes with no need for actual security.
//
// $ sudo make-ssl-cert generate-default-snakeoil --force-overwrite
//
// Note: I generated this pair using make-ssl-cert and then promptly replaced
// it with a new keypair.  This should be harmless.

package snakeoil

import (
	"io/ioutil"
	"os"
	"path"
)

// CertFileName writes its snakeoil cert to a temp file and returns the
// path to the tempfile.
func CertFileName() string {
	fileName := path.Join(os.TempDir(), "endless-snakeoil.pem")
	ioutil.WriteFile(fileName, []byte(Cert), 0777)
	return fileName
}

// KeyFileName writes its snakeoil key to a temp file and returns the
// path to the tempfile.
func KeyFileName() string {
	fileName := path.Join(os.TempDir(), "endless-snakeoil.key")
	ioutil.WriteFile(fileName, []byte(Key), 0777)
	return fileName
}

var Cert = `-----BEGIN CERTIFICATE-----
MIICxDCCAaygAwIBAgIJALZJDp2u07iGMA0GCSqGSIb3DQEBBQUAMBoxGDAWBgNV
BAMTD2h1YmVydHVzLWJpZ2VuZDAeFw0xNDA2MDQxNzU3MjFaFw0yNDA2MDExNzU3
MjFaMBoxGDAWBgNVBAMTD2h1YmVydHVzLWJpZ2VuZDCCASIwDQYJKoZIhvcNAQEB
BQADggEPADCCAQoCggEBAKV0FD0GrtZCSXPofWbyzIKMlHHl+X53hdgRY3QGjG6z
aWsrvoTEE4i689E1QF2SJmbkHWhV9/xF4aSgUrF9SpuXW9V7k1NVLkXhK2m0EMFN
6UvzCwb4VI30LGt8SgTSusiEGahMMdNBJwV7s+MeUVrZJKI9XzH3DNQBfLhMJ2xl
qXv2NKdXVuErXX2CAZfPhmMWlgyfysEvAtdqUSNaIBmsGAnMLVGB0A4UGo2jHsZ8
owHp/ftepBtGkikUvaRj67bC9BB17riTM6c94ztVImEIdl6fiunhyRjso3Dpi+Vv
ZLNma+1XrjfPsheaLatbrutMxNMUPu3If5bMTq4nSDkCAwEAAaMNMAswCQYDVR0T
BAIwADANBgkqhkiG9w0BAQUFAAOCAQEANcwRN3EOs/4kWG1ubvBDOoFlfOistb3u
Fz4ObbIflFprYCLaazsws587cQcnJG8zKjUlZqJ6U6/SdKE9uZlRnNpZXSHhC8i0
V4vvaNel/mv1hS661StjOGG71yJd7cJRjmGABkA2f8CzLM74jI68QpxtriutEDsG
KWumn1AN7PJpYXER80ivp8WgM68phl5OBENHzNfeLEyNv0+qJRLXKzYG8uPkDVb4
fSjeUaVb/1Km1PlnjMdamnsYriR9ko8KH7UdiolUeDVi4kecDJoULFDkutkvE1Px
ivmiajuflI1u2bUDmjyWhnGhgb2gXDBD+qsg4GZLH0jCnV1zyFH1yg==
-----END CERTIFICATE-----`

var Key = `-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCldBQ9Bq7WQklz
6H1m8syCjJRx5fl+d4XYEWN0Boxus2lrK76ExBOIuvPRNUBdkiZm5B1oVff8ReGk
oFKxfUqbl1vVe5NTVS5F4StptBDBTelL8wsG+FSN9CxrfEoE0rrIhBmoTDHTQScF
e7PjHlFa2SSiPV8x9wzUAXy4TCdsZal79jSnV1bhK119ggGXz4ZjFpYMn8rBLwLX
alEjWiAZrBgJzC1RgdAOFBqNox7GfKMB6f37XqQbRpIpFL2kY+u2wvQQde64kzOn
PeM7VSJhCHZen4rp4ckY7KNw6Yvlb2SzZmvtV643z7IXmi2rW67rTMTTFD7tyH+W
zE6uJ0g5AgMBAAECggEBAIkJaM3W4C60PnRTJrKC/WJPn4/q48ecpW39kPsDhYMW
9ISAec8rO+auuc0YpxQZPddQrw7AzaHUG30oEPXTCV1vcu+R58dIsQfN5RAqxQnm
RwnR5ttCCX59s+De08vE4lG8ICgl28rWlsRS3f6KGOX/HfmGjXcr4SMjjDJ1WN3M
zSHQxjClSj5slua0mx3F9R11xyFAxhzVOjVWDlAZ43PHRC+6mJte3u6wNkA2Ht1l
ayZ/Bm10NJJoiF5CBc+iMsscbwTA7lt06cZgJaXkCNbgkL9OpRdHJeaVAECaARWB
sCVKs75B+oOX1kjsnqwQwhv8gWOUwlbbFql7SaG85NECgYEA0h+1pw4yCarlYF6x
93ItkestD5v88oJChxPZknI+RvNg38usypXj0VFn4MqqwcHL0mZhmcgZKMknWHn/
rC9V4lIH6m5s812Kl8lTsBSrvRx++5ib6VlJR2n+RXlNPA7e4Syhx12ZcfhEBiub
qn1CQ3/rYv24EbWpL76P22PPmR0CgYEAyZOjv7hSfiTX1e6cV2zDEpfwv2IbH/sl
SjVKY+nnrzlhsUsaRFsiuLyN3iSJgjxy744/FUqSOAgpe+7wEvHIuuoaiQt+hyMf
Lyg4C5djyScLc0tQ0btSyAzfoTU8Z5wvCn/a2x+Wp0LD3Zuut1tbzwgQU+pf75Il
12UdLAdMnM0CgYBUGFbGPjsufVQB5sRJKUwtqzbEmYR5tkJT91DLeKeOE9fAma5V
AfpEitUNNW9zzlD5qvoC5v+SwDbcBS+bRPKVeokqfPljyRsvtmalARDexenYXfKA
SLi4OSaHvY14rObsrcUtmjtQTrFC2u8ZI7qCc07MtoiVpePJdPX0MxUhzQKBgQCp
8gxrNeZurXc+ySYMWwj34xu47uh5hQ7CSr2GEQ0g4NrpU604ljFm2Kku8VTxdS9d
omqy86TlEpSPBTpobmpSk859XjB1lHnVEy51L1SUoZN4x7XCrRC00o2z5yMfI0FR
s8t7VF16dMKXUorx0VdM67qAVSg+3dODjp8SoLDJvQKBgG/KlKA5VaMpjtOMrUQA
48b33X6Jy2In4MBOWTw49q9OChRqQEshB1/lEI4YRk1UbMk6S5vZSla7fqCthEUZ
KoIJ4Iqu0Qd20OYpAH/IxBjqUIdeCR4RSkiBwJui8ejb0zc2QpVqJ5e5s24qn7Gy
MuNGknVTu3N1u3Mg7/srx5UN
-----END PRIVATE KEY-----`

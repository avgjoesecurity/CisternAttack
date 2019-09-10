// Written by Ben Glass. Use to do cool stuff.

package main

import (
	//"io/ioutil"
	"fmt"
	"log"
	"net"
	"os"

	"encoding/base64"
	"golang.org/x/crypto/ssh"
)

var HOST_PRIVATE_KEY = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEA4FJCZ2TzIqqEwDTWxo/sWqrqBS3djclWT/JZyvXq2WHcKZed
WXByo8OQufhIa2Q75Q1bkZtxmoR1Eq2kMc+JZhmDCpjX9FJU0JwvyrDxVyMksl7s
PolSYF4DvsjZm9R+6/+RenTLdVZhUIcsiKS1q+1GkIvQfVr+0vf7Qpr1nD1XHRQ8
vKYGl8BWyTi5n/JWFt5UrgtoSkYx0RcJfn2D/MF3p6+TitMRtzLkKAOb2FjcsayU
CAGZZGH3KNzwweDLdvJmXUkD2nXfdIAV3bmUVQeDTjz7wBjgeXD4VBF7Yfl1+8xW
aEhmUlTR2UhwPOr6fX9KPlWL+cg63n0VOhM6SwIDAQABAoIBAEnxc6gnc6nysY22
HTIY0R1WB4wrgUOyoXv69DRd+cssYeln69tGoNNwrK1QtXaI9OFdFNkbd0gq73h3
uo8czTrVoAeHOHB2DzUwtuRGjVde40LU3WUD2R6g/vUNugWO90ON3AOUqliEX9Gu
rTRhruz8au1M2S3pJ+SdD7r1345fX82urUmede/cn7BdkBs02bM2+55DDv2+aR4s
Z9aOoyn7qG4bGii1NPanGb3vEjp5BrW1kP7FpoeoW58yKidFlS764/wteSrGoygv
i57zC+OyHZuk7zPma/84Gl57dgWD/s+JntofNz+O21Hjwlk0EG9PlfKrX3+Do9M1
QMu/N9kCgYEA+hBfJ7pDRWSPV87lJC3FPMzpwNgfrPSn1TnEw9Ylt+Pfwp/l1Mxl
oyFRzZsL1gDB6Q0T4pcEQ7UvwQn5lUg2JBd7o5N3sRBfRFoRrOnCbhH0QuornYSo
oYGIIGTYTAVTZ0JClLTHjjs15niemzHcuW+RbJnSjyI2wSo+NFXSGxcCgYEA5aVz
aJjf39pFzGa6rMsr9leECb4F1bQYdLwiPjpbBhi6z1QuFulhUCfih3WBtRFSIqru
3AbpJnUnuH10jcMqKyDw+/0Ncn9jDWaMduIKPfcpqBtTKYlJcxoJdcvFAu+OMv3u
sSxWy5E8ZHE6o46irqcLoWSlqM0+vr7EhmpQyu0CgYAD7hnwv+aqwofWM3Q6e1Ws
PPQ3ia4Fw9qXaU8EXOs+1p//qQ+tpoC6bZ2hg5xOArkJdqEnYIQ6pDBZ8lQv2nCv
ZIcW7QxmhgFZS6kXdJ0PqcyS+Kcy1JDp6OVAzRTxLOxApTPyrlysiZWxR5gGjGQR
rgU1Fb3PbhQ6OPW9UNP6hQKBgDvm73a+Qc97eS97MIB4Msx8QZXk+lKnOqXgmTIQ
OMQbWUdRlwAk+lxVOxLIhP/9N6aRXeMI7nI4pxbJqsh7uxkq65Ffjc1rCyRotoUg
KznEqhoOnp+8Duu2Q/d1IrvETxxf/o46mZ1aEh3FJ12YjDjwm+kKfufMOey5iDvj
qZV1AoGAX3Wlprf+sLtUivlBN6iwxpEK+efuwn/cLx9rEK8NGLfQXcQhgZb1CKpG
OtJo3ICseP7/m1ijLZKkUm0GbpKyyc50COyTEzdrgwhcfKX9ZfhD7FLo11V2EChe
5VzQAel7xGvm8kc3QA7zbcfqegqWDzbhVwQdwefTZVTxS5f91ro=
-----END RSA PRIVATE KEY-----`)

func handle(conn net.Conn, config *ssh.ServerConfig) {
	// From a standard TCP connection to an encrypted SSH connection
	sshConn, _, _, err := ssh.NewServerConn(conn, config)
	if err != nil {
		log.Printf("Failed to handshake (%s)", err)
		return
	}

	log.Println("Connection from", sshConn.RemoteAddr())
	sshConn.Close()
}

func main() {
	hostPrivateKeySigner, err := ssh.ParsePrivateKey(HOST_PRIVATE_KEY)
	if err != nil {
		panic(err)
	}

	config := &ssh.ServerConfig{
		//Define a function to run when a client attempts a password login
		PasswordCallback: func(c ssh.ConnMetadata, pass []byte) (*ssh.Permissions, error) {
			// Should use constant-time compare (or better, salt+hash) in a production setting.
			fmt.Println("User: " + string(c.User()) + " Password: " + string(pass))
			return nil, fmt.Errorf("password rejected for %q", c.User())
		},
		PublicKeyCallback: func(c ssh.ConnMetadata, key ssh.PublicKey) (*ssh.Permissions, error) {
			fmt.Println("User: " + string(c.User()) + " Public Key Type: " + string(key.Type()) + " Public Key: " + (base64.StdEncoding.EncodeToString(key.Marshal())))
			return nil, fmt.Errorf("key rejected for %q", c.User())
		},
	}
	config.AddHostKey(hostPrivateKeySigner)

	port := "2222"
	log.Println("Starting mock ssh server on " + port)
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	socket, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}

	for {
		conn, _ := socket.Accept()
		go handle(conn, config)
	}
}

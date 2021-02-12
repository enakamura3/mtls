# mtls

- [mtls](#mtls)
  - [About](#about)
  - [KEY, CSR and CRT](#key-csr-and-crt)
    - [CN - Common Name](#cn---common-name)
    - [SAN - Subject Alternative Name](#san---subject-alternative-name)
    - [Certificate info](#certificate-info)
  - [Client](#client)
    - [Client: CSR, KEY and CRT](#client-csr-key-and-crt)
  - [Server](#server)
    - [Server: CSR, KEY and CRT](#server-csr-key-and-crt)
  - [tests](#tests)
    - [server](#server-1)
    - [go client](#go-client)
    - [curl](#curl)

## About

Client and server using mtls

Source: 
> https://venilnoronha.io/a-step-by-step-guide-to-mtls-in-go

> https://kofo.dev/how-to-mtls-in-golang

> https://medium.com/jspoint/a-brief-overview-of-the-tcp-ip-model-ssl-tls-https-protocols-and-ssl-certificates-d5a6269fe29e

## KEY, CSR and CRT

- KEY: private key uset to sign an certificate
- CSR: certificate file request, the file created using the private key to send to an CA to sign
- CA: certificate authority, is an entity that issues digital certificates (the CSR file)
- CRT: the certificate file after CA sign using the CA private key

The KEY, CSR and CRT uses the PEM format.

When creating a KEY, you can fill some usefull info, or leave all in blank.

If you fill with some info, the certificate will be issued with thos info, like the CN (Common Name) info.

So, after creating the cert, you will need to add the host on /etc/hosts

Otherwise, you can have erros like:

```sh
curl: (51) SSL: certificate subject name 'some.company.wow' does not match target host name 'localhost'
```

### CN - Common Name

> The certificate is valid only if the request hostname matches the certificate common name

One of the validations is made checking the __hostname__ from the URL with the __common name/CN__ from server certificate.

If you are accessing: http://himate.com/good?no=good

The remote certificate must have in the CN: himate.com 

More info about the `Common Name` field

> https://support.dnsimple.com/articles/what-is-common-name/

> The Common Name (AKA CN) represents the server name protected by the SSL certificate. The certificate is valid only if the request hostname matches the certificate common name. Most web browsers display a warning message when connecting to an address that does not match the common name in the certificate.

> In the case of a single-name certificate, the common name consists of a single host name (e.g. example.com, www.example.com), or a wildcard name in case of a wildcard certificate (e.g. *.example.com).

Sample CN validation from curl -v:

```
* Server certificate:
*  subject: C=BR; ST=SP; L=Sao Paulo; O=some-company; OU=some-company-evil-section; CN=some.company.wow; emailAddress=iam@somecompany.com
*  start date: Jan 13 19:25:45 2021 GMT
*  expire date: Jan 13 19:25:45 2022 GMT
*  common name: some.company.wow (matched)
*  issuer: C=BR; ST=SP; L=Sao Paulo; O=some-company; OU=some-company-evil-section; CN=some.company.wow;
```

### SAN - Subject Alternative Name

> https://support.dnsimple.com/articles/what-is-common-name/#common-name-vs-subject-alternative-name

> The common name can only contain up to one entry: either a wildcard or non-wildcard name. It’s not possible to specify a list of names covered by an SSL certificate in the common name field.

> The Subject Alternative Name extension (also called Subject Alternate Name or SAN) was introduced to solve this limitation. The SAN allows issuance of multi-name SSL certificates.

> The ability to directly specify the content of a certificate SAN depends on the Certificate Authority and the specific product. Most certificate authorities have historically marketed multi-domain SSL certificates as a separate product. They’re generally charged at a higher rate than a standard single-name certificate.

> On the technical side, the SAN extension was introduced to integrate the common name. Since HTTPS was first introduced in 2000 (and defined by the RFC 2818), the use of the commonName field has been considered deprecated, because it’s ambiguous and untyped.

Sample SAN validation from curl -v:

```
* Server certificate:
*  subject: CN=viacep.com.br
*  start date: Mar 24 00:00:00 2020 GMT
*  expire date: Apr 18 23:59:59 2021 GMT
*  subjectAltName: host "viacep.com.br" matched cert's "viacep.com.br"
*  issuer: C=GB; ST=Greater Manchester; L=Salford; O=Sectigo Limited; CN=Sectigo RSA Domain Validation Secure Server CA
*  SSL certificate verify ok.
```

### Certificate info

To see the info from a certificate, you can use the openssl command:

```sh
openssl x509 -in CERTFILE -noout -text
```

Ex.:

```sh
openssl x509 -in server1/server.crt -noout -text
Certificate:
    Data:
        Version: 1 (0x0)
        Serial Number:
            01:5f:09:19:ee:66:cc:f5:4d:61:96:c0:03:24:a0:44:b0:be:da:53
        Signature Algorithm: sha256WithRSAEncryption
        Issuer: C = BR, ST = SP, L = Sao Paulo, O = some-company, OU = some-company-evil-section, CN = some.company.wow, emailAddress = iam@somecompany.com
        Validity
            Not Before: Jan 13 19:25:45 2021 GMT
            Not After : Jan 13 19:25:45 2022 GMT
        Subject: C = BR, ST = SP, L = Sao Paulo, O = some-company, OU = some-company-evil-section, CN = some.company.wow, emailAddress = iam@somecompany.com
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                RSA Public-Key: (2048 bit)
                Modulus:
                    00:af:39:55:1b:63:b8:87:ba:d2:38:b9:45:8a:95:
                    e3:55:ad:c6:75:55:dd:07:4a:79:85:1b:71:02:d3:
                    67:13:97:5d:a2:6c:50:40:c8:6e:c7:c6:97:2d:ee:
                    e0:43:4c:54:4d:d6:52:de:ec:1c:45:f2:f2:c7:a6:
                    5e:51:fb:9d:cf:0f:01:fd:a6:2c:d8:37:a8:0f:85:
                    0d:ca:30:fc:4a:24:d8:94:e2:8b:07:8b:88:90:4e:
                    04:e0:ff:f5:c7:22:d4:fa:82:6e:f4:f1:7f:d0:a5:
                    ec:37:8e:32:b1:da:08:1f:e2:30:f5:3a:8c:4a:33:
                    40:60:84:1a:c4:19:03:a1:d2:ff:e8:17:8c:df:1f:
                    e4:91:89:c2:d4:ea:bc:ed:56:68:ea:9c:e3:b1:4e:
                    5f:e4:8c:5e:11:07:74:3d:7a:35:57:c3:70:a5:e8:
                    ca:12:fd:71:a1:fb:ce:89:c9:98:ff:ba:7e:35:80:
                    bc:a6:51:81:d8:c0:aa:d9:f8:cd:c2:85:be:21:ce:
                    2d:5c:7e:91:74:65:cb:03:15:58:13:df:63:fb:9f:
                    96:9f:58:02:97:fb:61:cc:20:7f:40:32:27:be:c6:
                    70:7c:f7:53:5a:2f:db:6c:ae:ce:fb:ed:cb:79:86:
                    3b:63:28:59:60:55:e4:8b:8a:48:87:97:65:0a:37:
                    ca:7b
                Exponent: 65537 (0x10001)
    Signature Algorithm: sha256WithRSAEncryption
         7c:ae:af:eb:dc:b9:41:c5:9a:26:2f:4e:d6:5b:df:b7:3c:39:
         07:d5:81:71:a0:60:fe:c9:6f:1f:e5:0b:6c:e1:03:0a:fd:ba:
         3e:8b:7a:52:a7:45:f5:45:a4:3f:26:3c:98:c6:ce:c0:04:35:
         9a:6a:17:ca:d4:18:6c:1c:a3:10:29:c8:38:aa:66:2f:c1:99:
         d5:93:e7:8e:6d:04:2d:52:60:e5:51:48:85:49:b4:4f:24:26:
         15:c1:11:43:4b:42:ec:47:0e:30:6b:90:e7:c1:ef:90:c5:8e:
         2e:7c:c5:7a:47:ad:b0:3e:41:f9:ff:93:d4:a2:ce:f0:1e:16:
         bf:81:b8:68:ab:ce:5c:8a:fd:a3:82:38:a5:b9:30:7c:c1:6f:
         bc:ed:0e:ab:0b:eb:87:0b:26:ef:01:33:12:01:c8:80:e0:b4:
         88:28:20:03:4d:38:fe:1a:36:2a:b2:aa:56:e6:1f:0f:ba:cb:
         54:00:c7:30:d1:34:67:2e:d4:c0:ba:0f:04:50:fe:d4:a0:28:
         a4:fe:42:3f:26:89:18:b0:cc:44:33:0e:e8:78:90:a8:ca:26:
         cb:be:7e:26:9b:3d:7e:a1:0b:bd:54:66:21:5c:15:bf:89:d8:
         17:0a:23:f8:2d:eb:04:95:29:19:93:32:16:b5:17:03:93:37:
         02:dd:c1:58
```


## Client

https client using go 

To load an cert or key file, you can use `tls.LoadX509KeyPair` or `tls.X509KeyPair`.

The difference is the parameter format.

- LoadX509KeyPair: file name
- X509KeyPair: byte array

### Client: CSR, KEY and CRT 

The `client1` and `client2` directory already have the certificate and key created.

The following commands were used to create it.

- create private KEY and CSR file (no password required)

```sh
openssl req -new -newkey rsa:2048 -nodes -keyout client.key -out client.csr
```

- sign the csr file using your private key

```sh
openssl x509 -req -days 365 -in client.csr -signkey client.key -out client.crt
```


## Server

https server using go

### Server: CSR, KEY and CRT

The `server1` and `server2` directory already have the certificate and key created.

The following commands were used to create it.

- create private KEY and CSR file (no password required)

```sh
openssl req -new -newkey rsa:2048 -nodes -keyout server.key -out server.csr
```

- sign the csr file using your private key

```sh
openssl x509 -req -days 365 -in server.csr -signkey server.key -out server.crt
```

## tests

### server

We have 2 servers:

- server1: use CN = some.company.wow
- server2: use CN = localhost

start the server:

```sh
cd server1
go run server.go
```

or 

```sh
cd server2
go run server.go
```

If you start the `server1` directory, you will need to add an entry on /etc/hosts file.
 
```
127.0.0.1 some.company.wow
```

Because when I generated the private key, I filled in the fields with some fictitious values.

```
CN = some.company.wow
```

### go client

We have 2 clients:

- client1: call some.company.wow
- client2: call localhost

```sh
cd client1
go run client.go
```

or

```sh
cd client2
go run client.go
```

If an error like:

```
Get "https://some.company.wow:8443/hello": x509: certificate relies on legacy Common Name field, use SANs or temporarily enable Common Name matching with GODEBUG=x509ignoreCN=0
exit status 1
```

execute: 

```sh
export GODEBUG=x509ignoreCN=0
```

and run it again

```sh
go run client.go 
Hello, world!
```

### curl

- success

Calling `server1`

```sh
curl -v https://some.company.wow:8443/hello --key client1/client.key --cert client1/client.crt --cacert server1/server.crt

*   Trying 127.0.0.1...
* TCP_NODELAY set
* Connected to some.company.wow (127.0.0.1) port 8443 (#0)
* ALPN, offering h2
* ALPN, offering http/1.1
* successfully set certificate verify locations:
*   CAfile: server/server.crt
  CApath: /etc/ssl/certs
* TLSv1.3 (OUT), TLS handshake, Client hello (1):
* TLSv1.3 (IN), TLS handshake, Server hello (2):
* TLSv1.3 (IN), TLS Unknown, Certificate Status (22):
* TLSv1.3 (IN), TLS handshake, Unknown (8):
* TLSv1.3 (IN), TLS Unknown, Certificate Status (22):
* TLSv1.3 (IN), TLS handshake, Request CERT (13):
* TLSv1.3 (IN), TLS Unknown, Certificate Status (22):
* TLSv1.3 (IN), TLS handshake, Certificate (11):
* TLSv1.3 (IN), TLS Unknown, Certificate Status (22):
* TLSv1.3 (IN), TLS handshake, CERT verify (15):
* TLSv1.3 (IN), TLS Unknown, Certificate Status (22):
* TLSv1.3 (IN), TLS handshake, Finished (20):
* TLSv1.3 (OUT), TLS change cipher, Client hello (1):
* TLSv1.3 (OUT), TLS Unknown, Certificate Status (22):
* TLSv1.3 (OUT), TLS handshake, Certificate (11):
* TLSv1.3 (OUT), TLS Unknown, Certificate Status (22):
* TLSv1.3 (OUT), TLS handshake, CERT verify (15):
* TLSv1.3 (OUT), TLS Unknown, Certificate Status (22):
* TLSv1.3 (OUT), TLS handshake, Finished (20):
* SSL connection using TLSv1.3 / TLS_AES_128_GCM_SHA256
* ALPN, server accepted to use h2
* Server certificate:
*  subject: C=BR; ST=SP; L=Sao Paulo; O=some-company; OU=some-company-evil-section; CN=some.company.wow; emailAddress=iam@somecompany.com
*  start date: Jan 13 19:25:45 2021 GMT
*  expire date: Jan 13 19:25:45 2022 GMT
*  common name: some.company.wow (matched)
*  issuer: C=BR; ST=SP; L=Sao Paulo; O=some-company; OU=some-company-evil-section; CN=some.company.wow; emailAddress=iam@somecompany.com
*  SSL certificate verify ok.
* Using HTTP2, server supports multi-use
* Connection state changed (HTTP/2 confirmed)
* Copying HTTP/2 data in stream buffer to connection buffer after upgrade: len=0
* TLSv1.3 (OUT), TLS Unknown, Unknown (23):
* TLSv1.3 (OUT), TLS Unknown, Unknown (23):
* TLSv1.3 (OUT), TLS Unknown, Unknown (23):
* Using Stream ID: 1 (easy handle 0x5594b5f265c0)
* TLSv1.3 (OUT), TLS Unknown, Unknown (23):
> GET /hello HTTP/2
> Host: some.company.wow:8443
> User-Agent: curl/7.58.0
> Accept: */*
> 
* TLSv1.3 (IN), TLS Unknown, Certificate Status (22):
* TLSv1.3 (IN), TLS handshake, Newsession Ticket (4):
* TLSv1.3 (IN), TLS Unknown, Unknown (23):
* Connection state changed (MAX_CONCURRENT_STREAMS updated)!
* TLSv1.3 (OUT), TLS Unknown, Unknown (23):
* TLSv1.3 (IN), TLS Unknown, Unknown (23):
* TLSv1.3 (IN), TLS Unknown, Unknown (23):
* TLSv1.3 (IN), TLS Unknown, Unknown (23):
< HTTP/2 200 
< content-type: text/plain; charset=utf-8
< content-length: 14
< date: Thu, 14 Jan 2021 23:30:34 GMT
< 
Hello, world!
* Connection #0 to host some.company.wow left intact
```

Calling server2

```sh
curl -v https://localhost:8443/hello --key client2/client.key --cert client2/client.crt --cacert server2/server.crt*   Trying 127.0.0.1...
* TCP_NODELAY set
* Connected to localhost (127.0.0.1) port 8443 (#0)
* ALPN, offering h2
* ALPN, offering http/1.1
* successfully set certificate verify locations:
*   CAfile: server2/server.crt
  CApath: /etc/ssl/certs
* TLSv1.3 (OUT), TLS handshake, Client hello (1):
* TLSv1.3 (IN), TLS handshake, Server hello (2):
* TLSv1.3 (IN), TLS Unknown, Certificate Status (22):
* TLSv1.3 (IN), TLS handshake, Unknown (8):
* TLSv1.3 (IN), TLS Unknown, Certificate Status (22):
* TLSv1.3 (IN), TLS handshake, Request CERT (13):
* TLSv1.3 (IN), TLS Unknown, Certificate Status (22):
* TLSv1.3 (IN), TLS handshake, Certificate (11):
* TLSv1.3 (IN), TLS Unknown, Certificate Status (22):
* TLSv1.3 (IN), TLS handshake, CERT verify (15):
* TLSv1.3 (IN), TLS Unknown, Certificate Status (22):
* TLSv1.3 (IN), TLS handshake, Finished (20):
* TLSv1.3 (OUT), TLS change cipher, Client hello (1):
* TLSv1.3 (OUT), TLS Unknown, Certificate Status (22):
* TLSv1.3 (OUT), TLS handshake, Certificate (11):
* TLSv1.3 (OUT), TLS Unknown, Certificate Status (22):
* TLSv1.3 (OUT), TLS handshake, CERT verify (15):
* TLSv1.3 (OUT), TLS Unknown, Certificate Status (22):
* TLSv1.3 (OUT), TLS handshake, Finished (20):
* SSL connection using TLSv1.3 / TLS_AES_128_GCM_SHA256
* ALPN, server accepted to use h2
* Server certificate:
*  subject: C=AU; ST=Some-State; O=Internet Widgits Pty Ltd; CN=localhost
*  start date: Jan 14 23:20:31 2021 GMT
*  expire date: Jan 14 23:20:31 2022 GMT
*  common name: localhost (matched)
*  issuer: C=AU; ST=Some-State; O=Internet Widgits Pty Ltd; CN=localhost
*  SSL certificate verify ok.
* Using HTTP2, server supports multi-use
* Connection state changed (HTTP/2 confirmed)
* Copying HTTP/2 data in stream buffer to connection buffer after upgrade: len=0
* TLSv1.3 (OUT), TLS Unknown, Unknown (23):
* TLSv1.3 (OUT), TLS Unknown, Unknown (23):
* TLSv1.3 (OUT), TLS Unknown, Unknown (23):
* Using Stream ID: 1 (easy handle 0x55d0c30825c0)
* TLSv1.3 (OUT), TLS Unknown, Unknown (23):
> GET /hello HTTP/2
> Host: localhost:8443
> User-Agent: curl/7.58.0
> Accept: */*
> 
* TLSv1.3 (IN), TLS Unknown, Certificate Status (22):
* TLSv1.3 (IN), TLS handshake, Newsession Ticket (4):
* TLSv1.3 (IN), TLS Unknown, Unknown (23):
* Connection state changed (MAX_CONCURRENT_STREAMS updated)!
* TLSv1.3 (OUT), TLS Unknown, Unknown (23):
* TLSv1.3 (IN), TLS Unknown, Unknown (23):
* TLSv1.3 (IN), TLS Unknown, Unknown (23):
< HTTP/2 200 
< content-type: text/plain; charset=utf-8
< content-length: 14
< date: Thu, 14 Jan 2021 23:37:46 GMT
< 
Hello, world!
* Connection #0 to host localhost left intact
```

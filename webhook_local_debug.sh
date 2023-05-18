#!/usr/bin/env bash

if [ ! -d "$1/config/cert" ]
then
  mkdir -p "$1"/config/cert
fi

rm -rf "$1"/config/cert
cd "$1"/config/cert

cat << EOF > csr.conf
[ req ]
default_bits = 2048
prompt = no
default_md = sha256
req_extensions = req_ext
distinguished_name = dn

[ dn ]
C = CN
ST = Guangzhou
L = Shenzhen
CN = $2

[ req_ext ]
subjectAltName = @alt_names

[ alt_names ]
IP.1 = $2

[ v3_ext ]
authorityKeyIdentifier=keyid,issuer:always
basicConstraints=CA:FALSE
keyUsage=keyEncipherment,dataEncipherment
extendedKeyUsage=serverAuth,clientAuth
subjectAltName=@alt_names
EOF

openssl genrsa -out ca.key 2048
openssl req -x509 -new -nodes -key ca.key -subj "/CN=$2" -days 10000 -out ca.crt
openssl genrsa -out tls.key 2048
openssl req -new -SHA256 -newkey rsa:2048 -nodes -keyout tls.key -out tls.csr -subj "/C=CN/ST=Shanghai/L=Shanghai/O=/OU=/CN=$2"
openssl req -new -key tls.key -out tls.csr -config csr.conf
openssl x509 -req -in tls.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out tls.crt -days 10000 -extensions v3_ext -extfile csr.conf
cat ca.crt | base64 | tr -d '\n' >> a.txt

echo "success"

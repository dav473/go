- Generate private key(key public key is embeded in)
`openssl genrsa -out ./ca/ca.key 2048` 

- OpenSSL will generate a new CSR using the private key provided (server.key) and write the CSR to the specified output file (server.csr). This CSR can then be provided to a CA to obtain a certificate for your server.

- i remember to get the certificate from CA, server need to provide the public key, but why in above code, -key requires the private key
- Generation of the CSR: The CSR contains information about your server (like its domain name, organization details, etc.) and its public key. However, it also needs to be signed by the corresponding private key to prove that the requester has control over the public key. This signing process ensures the integrity of the information in the CSR.
- Private Key Requirement: OpenSSL requires the private key to generate the CSR and to sign it. The private key is used to create a digital signature on the CSR, which is a cryptographic proof that the CSR was generated by the entity possessing the private key and that the CSR's contents have not been tampered with.

`openssl req -new -key ./ca/ca.key -out ./ca/ca.csr`

-  OpenSSL will generate a new self-signed certificate using the private key provided (server.key) and write it to the specified output file (server.crt). This certificate will be valid for 100 years.
`openssl req -new -x509 -key ./ca/ca.key -out ./ca/ca.crt -days 36500`

-----------------------SERVER


`openssl genrsa -out ./server/server.key 2048`

`openssl req -new -key ./server/server.key -out ./server/server.csr -config ./openssl.cnf -extensions v3_req`



`openssl x509 -req -days 36500 -in ./server/server.csr -out ./server/server.pem -CA ./ca/ca.crt -CAkey ./ca/ca.key -CAcreateserial -extfile ./openssl.cnf -extensions v3_req`





  // const  PROTO_PATH = "/home/zhen/govpn/server/aweb/app/proto/hello.proto";
  // const protoDefinition = loadSync(
  //   PROTO_PATH,{
  //     keepCase: true,
  //     longs: String,
  //     enums: String,
  //     defaults: true,
  //     oneofs: true
  //   }
  // )
  // const hello_proto = grpc.loadPackageDefinition(protoDefinition)

  
  // const client = new (hello_proto as any).SayHello(
  //   "localhost:9090",
  //   grpc.credentials.createInsecure()
  // );

  // const stream =client.sayHello({}) ;
  // stream.on("data", (data) => {
  //   console.log(data)
  // })
# GO-WEB
GO1.13.4 web apllication using go

** Note: Please ignore method syntax, refer go official documentation for proper syntax.** 

Code repository for my Microservices POC using golang.

Documentation of Building Microservices with golang using go standard "http" and "Gorilla/Mux" packages. Please feel free to leave a comment.

## Introduction:
Everything in go is an http handler and it had a default handler called serveMux. "http" package as simple functions like "Listenand Serve(address, Handler)" which listens on the TCP network address and then calls handler to handle request. DefaultServeMux is used when the Handler is nil. 

## Custom Server:
Server is a struct in http package, once you created a new struct with properties then you call "ListenAndServe" function using the address of new server(s := &http.Server{}).. It required in order to set properties like "IdleTimeout, ReadTimeout, WriteTImeout etc… (check doc for more properties) 

## Custom serveMux:
You can create custom serveMux using "http.NewServeMux()" and register custom handlers using "Handle" method on it.

## Custom Handler:
Http Handler is an interface which implements ServeHTTP method. To create a custom handler, create an Interface and implement ServeHTTP method.

## Graceful shutdown:
Custom server as function shutdown to which we can pass a context with some timeout, so if there are any unfinished requests they can finish before shutdown.

## Struct field tags:
These tags are required in Marshaling and Unmarshaling data.\
Note: Use NewEncoder and Encode functions, as they are clean and performance efficient way.
- Struct field tags are king of annotations use to customize the fields. `json:"myName"`
- Can omit the field if it is empty. `json:"myName, omitempty"`
- Field can completely skipped. `json:" - "`

## Restful Service:
Representational State Transfer(REST) is an architectural approch/pattern to design web services, in which resources are manupilated using simple operations: GET,POST,PUT and DELETE.





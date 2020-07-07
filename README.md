# JSON Schema Validation

Follow below steps to validate the json document
## clone the code
- Copy the code into $GOPATH/src
- modify the host_document.json as per your host info
- update the validateJson.go file with the file path for 
host_schema.json and host_document.json

Note:
(the prefix (file://) and a full path to the file are required)

- Install the dependency package
```bash
go get "github.com/xeipuuv/gojsonschema"
``` 

- run below command
 
```bash
go test
```



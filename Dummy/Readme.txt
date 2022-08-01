Notes:

1. For project create folder name as Project-Name.
2. go mod init Project-name
3. go mod tidy
4. Get all dependancies
 examples,
    1. go get google.golang.org/protobuf/cmd/protoc-gen-go
    2. go get github.com/golang/protobuf
5. Create proto folder under project
6. Create .proto files with package proto [folder name]
7. Define go_package="./proto;gen"
8. For generating .pb files run from project folder
    protoc -I=proto --go_out=. proto/*.proto

9. For using this file in go files import as,
    import pb "project-folder/proto"
   and then use as pb.{message_name}

#### Preparation
```bash
# Install protoc-gen-go-grpc for generate service
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Or
brew install protoc-gen-go-grpc
```
#### Generate Proto
```bash
cd /common/model

# generate message
protoc --go_out=. *.proto

# generate service
protoc --go-grpc_out=. *.proto

comment, new version of proto
// mustEmbedUnimplementedGaragesServer()
```


#### Run CMD
```bash
# service garage
asx@MacMini github-mragungsetiaji@learn-go % go run grpc/services/service-garage/main.go
2022/03/21 20:25:48 Starting RPC server at :9001
2022/03/21 20:26:41 Adding garage id:"q001" name:"tuktuk tempe" coordinate:{latitude:45.123123 longitude:54.12313} for user n001
2022/03/21 20:26:41 Adding garage id:"x02" name:"Simba Ubi" coordinate:{latitude:45.123123 longitude:54.12313} for user n001

# service user
asx@MacMini github-mragungsetiaji@learn-go % go run grpc/services/service-user/main.go
2022/03/21 20:25:58 Starting RPC server at :9000
2022/03/21 20:26:41 Registering user id:"n001" name:"Tuk tuk" password:"tuktuk123" gender:MALE
2022/03/21 20:26:41 Registering user id:"n002" name:"Simba" password:"simba321" gender:MALE


# client
asx@MacMini github-mragungsetiaji@learn-go % go run grpc/client/main.go 

 ===========> user test
2022/03/21 20:26:41 [{"id":"n001","name":"Tuk tuk","password":"tuktuk123","gender":1},{"id":"n002","name":"Simba","password":"simba321","gender":1}]

 ===========> garage test A
2022/03/21 20:26:41 [{"id":"q001","name":"tuktuk tempe","coordinate":{"latitude":45.123123,"longitude":54.12313}},{"id":"x02","name":"Simba Ubi","coordinate":{"latitude":45.123123,"longitude":54.12313}}]
asx@MacMini github-mragungsetiaji@learn-go %
```
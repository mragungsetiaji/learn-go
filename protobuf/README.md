### Protobuf

#### Generate Proto
```bash
protoc --go_out=. *.proto
```

#### Run CMD
```bash
asx@MacMini github-mragungsetiaji@learn-go % go run protobuf/main.go  
# ==== Original
       &model.Star{state:impl.MessageState{NoUnkeyedLiterals:pragma.NoUnkeyedLiterals{}, DoNotCompare:pragma.DoNotCompare{}, DoNotCopy:pragma.DoNotCopy{}, atomicMessageInfo:(*impl.MessageInfo)(nil)}, sizeCache:0, unknownFields:[]uint8(nil), Id:"x1", Name:"Robert Pattison", Gender:1} 
# ==== As String
       id:"x1" name:"Robert Pattison" gender:MALE 
# ==== Original
       &model.Movie{state:impl.MessageState{NoUnkeyedLiterals:pragma.NoUnkeyedLiterals{}, DoNotCompare:pragma.DoNotCompare{}, DoNotCopy:pragma.DoNotCopy{}, atomicMessageInfo:(*impl.MessageInfo)(nil)}, sizeCache:0, unknownFields:[]uint8(nil), Id:"m1", Title:"The Batman", Year:2022} 
# ==== As String
       id:"m1" title:"The Batman" year:2022 
# ==== As JSON
       {"id":"m1","title":"The Batman","year":2022} 
# ==== As Proto Object
       &model.Movie{state:impl.MessageState{NoUnkeyedLiterals:pragma.NoUnkeyedLiterals{}, DoNotCompare:pragma.DoNotCompare{}, DoNotCopy:pragma.DoNotCopy{}, atomicMessageInfo:(*impl.MessageInfo)(0xc000126000)}, sizeCache:0, unknownFields:[]uint8(nil), Id:"m1", Title:"The Batman", Year:2022} 
# ==== Original
       &model.MovieListByStar{state:impl.MessageState{NoUnkeyedLiterals:pragma.NoUnkeyedLiterals{}, DoNotCompare:pragma.DoNotCompare{}, DoNotCopy:pragma.DoNotCopy{}, atomicMessageInfo:(*impl.MessageInfo)(nil)}, sizeCache:0, unknownFields:[]uint8(nil), List:map[string]*model.MovieList{"x1":(*model.MovieList)(0xc000020480)}} 
# ==== As String
       List:{key:"x1" value:{movies:{id:"m1" title:"The Batman" year:2022} movies:{id:"m2" title:"Breaking Dawn" year:2011}}} 
```
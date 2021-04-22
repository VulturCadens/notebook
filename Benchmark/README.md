# Benchmarking JSON versus Protocol Buffers

* Serialize a native Go structure into a JSON string by using the Go standard library and vice versa.

* Serialize a native Go structure into binary data by using Protocol Buffers and vice versa.

## Performance
```
$ go test -benchmem -bench .

goos: linux
goarch: amd64
pkg: benchmark
BenchmarkJsonToStruct-4       	   50000	 32910 ns/op	     552 B/op	    30 allocs/op
BenchmarkStructToJson-4       	  200000	  6209 ns/op	     576 B/op	     1 allocs/op
BenchmarkProtobufToStruct-4   	  200000	  6325 ns/op	    1360 B/op	    40 allocs/op
BenchmarkStructToProtobuf-4   	  500000	  2470 ns/op	     352 B/op	     1 allocs/op
PASS
ok  	benchmark	5.905s
```

| Operation | ns/op |
| :--- | ---: |
| JSON -> Struct | 32910 |
| ProtoBuf -> Struct | **6325** |
| Struct -> JSON | 6209 |
| Struct -> ProtoBuf | **2470** |

## Size
```
$ go run main.go 

Protocol Buffers: 326 bytes
JSON: 544 bytes
```

| Operation | Size |
| :--- | :--- |
| JSON | 544 bytes |
| ProtoBuf | **326 bytes** |

##  Structure

```
$ go run main.go -write
$ hexdump -C BOOK.json

00000000  7b 22 74 69 74 6c 65 22  3a 22 4c 65 73 20 4d 69  |{"title":"Les Mi|
00000010  73 c3 a9 72 61 62 6c 65  73 22 2c 22 61 75 74 68  |s..rables","auth|
00000020  6f 72 22 3a 22 56 69 63  74 6f 72 20 48 75 67 6f  |or":"Victor Hugo|
00000030  22 2c 22 63 6f 75 6e 74  72 79 22 3a 22 46 72 61  |","country":"Fra|
00000040  6e 63 65 22 2c 22 79 65  61 72 22 3a 31 38 36 32  |nce","year":1862|
00000050  2c 22 76 6f 6c 75 6d 65  73 22 3a 5b 7b 22 76 6f  |,"volumes":[{"vo|
00000060  6c 75 6d 65 22 3a 31 2c  22 74 69 74 6c 65 22 3a  |lume":1,"title":|
00000070  22 46 61 6e 74 69 6e 65  22 7d 2c 7b 22 76 6f 6c  |"Fantine"},{"vol|
00000080  75 6d 65 22 3a 32 2c 22  74 69 74 6c 65 22 3a 22  |ume":2,"title":"|
00000090  43 6f 73 65 74 74 65 22  7d 2c 7b 22 76 6f 6c 75  |Cosette"},{"volu|
000000a0  6d 65 22 3a 33 2c 22 74  69 74 6c 65 22 3a 22 4d  |me":3,"title":"M|
000000b0  61 72 69 75 73 22 7d 2c  7b 22 76 6f 6c 75 6d 65  |arius"},{"volume|
000000c0  22 3a 34 2c 22 74 69 74  6c 65 22 3a 22 54 68 65  |":4,"title":"The|
000000d0  20 49 64 79 6c 6c 20 69  6e 20 74 68 65 20 52 75  | Idyll in the Ru|
000000e0  65 20 50 6c 75 6d 65 74  20 61 6e 64 20 74 68 65  |e Plumet and the|
000000f0  20 45 70 69 63 20 69 6e  20 74 68 65 20 52 75 65  | Epic in the Rue|
00000100  20 53 74 2e 20 44 65 6e  69 73 22 7d 2c 7b 22 76  | St. Denis"},{"v|
00000110  6f 6c 75 6d 65 22 3a 35  2c 22 74 69 74 6c 65 22  |olume":5,"title"|
00000120  3a 22 4a 65 61 6e 20 56  61 6c 6a 65 61 6e 22 7d  |:"Jean Valjean"}|
00000130  5d 2c 22 63 68 61 72 61  63 74 65 72 73 22 3a 7b  |],"characters":{|
00000140  22 6d 61 6a 6f 72 22 3a  5b 22 4a 65 61 6e 20 56  |"major":["Jean V|
00000150  61 6c 6a 65 61 6e 22 2c  22 4a 61 76 65 72 74 22  |aljean","Javert"|
00000160  2c 22 46 61 6e 74 69 6e  65 22 2c 22 43 6f 73 65  |,"Fantine","Cose|
00000170  74 74 65 22 2c 22 4d 61  72 69 75 73 20 50 6f 6e  |tte","Marius Pon|
00000180  74 6d 65 72 63 79 22 2c  22 c3 89 70 6f 6e 69 6e  |tmercy","..ponin|
00000190  65 22 5d 2c 22 6d 69 6e  6f 72 22 3a 5b 22 4d 6f  |e"],"minor":["Mo|
000001a0  6e 73 69 65 75 72 20 54  68 c3 a9 6e 61 72 64 69  |nsieur Th..nardi|
000001b0  65 72 22 2c 22 4d 61 64  61 6d 65 20 54 68 c3 a9  |er","Madame Th..|
000001c0  6e 61 72 64 69 65 72 22  2c 22 45 6e 6a 6f 6c 72  |nardier","Enjolr|
000001d0  61 73 22 2c 22 47 61 76  72 6f 63 68 65 22 2c 22  |as","Gavroche","|
000001e0  42 69 73 68 6f 70 20 4d  79 72 69 65 6c 22 5d 7d  |Bishop Myriel"]}|
000001f0  2c 22 74 72 61 6e 73 6c  61 74 69 6f 6e 73 22 3a  |,"translations":|
00000200  5b 31 38 36 32 2c 31 38  38 37 2c 31 39 37 36 2c  |[1862,1887,1976,|
00000210  31 39 38 37 2c 32 30 30  38 2c 32 30 31 33 5d 7d  |1987,2008,2013]}|
```
```
$ hexdump -C BOOK.protobuf

00000000  0a 0f 4c 65 73 20 4d 69  73 c3 a9 72 61 62 6c 65  |..Les Mis..rable|
00000010  73 12 0b 56 69 63 74 6f  72 20 48 75 67 6f 1a 06  |s..Victor Hugo..|
00000020  46 72 61 6e 63 65 20 c6  0e 2a 0b 08 01 12 07 46  |France ..*.....F|
00000030  61 6e 74 69 6e 65 2a 0b  08 02 12 07 43 6f 73 65  |antine*.....Cose|
00000040  74 74 65 2a 0a 08 03 12  06 4d 61 72 69 75 73 2a  |tte*.....Marius*|
00000050  41 08 04 12 3d 54 68 65  20 49 64 79 6c 6c 20 69  |A...=The Idyll i|
00000060  6e 20 74 68 65 20 52 75  65 20 50 6c 75 6d 65 74  |n the Rue Plumet|
00000070  20 61 6e 64 20 74 68 65  20 45 70 69 63 20 69 6e  | and the Epic in|
00000080  20 74 68 65 20 52 75 65  20 53 74 2e 20 44 65 6e  | the Rue St. Den|
00000090  69 73 2a 10 08 05 12 0c  4a 65 61 6e 20 56 61 6c  |is*.....Jean Val|
000000a0  6a 65 61 6e 32 91 01 0a  0c 4a 65 61 6e 20 56 61  |jean2....Jean Va|
000000b0  6c 6a 65 61 6e 0a 06 4a  61 76 65 72 74 0a 07 46  |ljean..Javert..F|
000000c0  61 6e 74 69 6e 65 0a 07  43 6f 73 65 74 74 65 0a  |antine..Cosette.|
000000d0  10 4d 61 72 69 75 73 20  50 6f 6e 74 6d 65 72 63  |.Marius Pontmerc|
000000e0  79 0a 08 c3 89 70 6f 6e  69 6e 65 12 14 4d 6f 6e  |y....ponine..Mon|
000000f0  73 69 65 75 72 20 54 68  c3 a9 6e 61 72 64 69 65  |sieur Th..nardie|
00000100  72 12 12 4d 61 64 61 6d  65 20 54 68 c3 a9 6e 61  |r..Madame Th..na|
00000110  72 64 69 65 72 12 08 45  6e 6a 6f 6c 72 61 73 12  |rdier..Enjolras.|
00000120  08 47 61 76 72 6f 63 68  65 12 0d 42 69 73 68 6f  |.Gavroche..Bisho|
00000130  70 20 4d 79 72 69 65 6c  3a 0c c6 0e df 0e b8 0f  |p Myriel:.......|
00000140  c3 0f d8 0f dd 0f                                 |......|
```

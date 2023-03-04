# protoc-gen-go-ascii

An example repo showing how to build a simple protoc plugin for generating go code

<https://rotemtam.com/2021/03/22/creating-a-protoc-plugin-to-gen-go-code/>

```bash
# Install proto compilers
brew install protobuf
brew install protoc-gen-go

protoc --version
# libprotoc 3.21.12
protoc-gen-go --version
# protoc-gen-go v1.28.1


# Build to program into binary "protoc-gen-go-ascii"
go build

# Run
protoc \
  --plugin=protoc-gen-go-ascii \
  --go-ascii_out=. \
  --go-ascii_opt=paths=source_relative \
  --go_out=. \
  --go_opt=paths=source_relative \
  --proto_path=. \
  example/example.proto

# --go-ascii_out: protoc-gen-go-ascii: Plugin output is unparseable:
# going to print example.Hello\'s ASCII art representation:
# \n _   _        _  _\n| | | |      | || |\n| |_| |  ___ | || |  ___\n|  _  | / _ \\| || | / _ \\\n| | | ||  __/| || || (_) |\n\\_| |_/ \\___||_||_| \\___/\n\n
```

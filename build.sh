#MacOS
CGO_ENABLED=1 GOOS=darwin go build -o ./bin/goke-mac

#Linux
CGO_ENABLED=1 GOOS=linux go build -o ./bin/goke-linux

#NOTE: since we need CGO_ENABLED=1, if you want to build for windows, you need to install mingw-w64
#https://sourceforge.net/projects/mingw-w64/

#Windows x64 | CXX -> C++ compiler | CC -> C compiler
CGO_ENABLED=1 GOOS=windows CXX=i686-w64-mingw32-g++ CC=i686-w64-mingw32-gcc GOARCH=amd64 go build -o ./bin/goke-windows-x64.exe

#Windows x86 | CXX -> C++ compiler | CC -> C compiler
CGO_ENABLED=1 GOOS=windows CXX=i686-w64-mingw32-g++ CC=i686-w64-mingw32-gcc GOARCH=386 go build -o ./bin/goke-windows-x86.exe





all: 
#	go get --ldflags '-extldflags "-Wl,--allow-multiple-definition"' .
	go build -ldflags "-X 'main._version_=git.$(shell git log --pretty=format:"%h" -1)'"
	#go build -ldflags "-X 'main._version_=0adaf'"

exe: 
	CGO_ENABLED=0 GOOS=windows go build -installsuffix cgo -a -ldflags  "-X 'main._version_=git.$(shell git log --pretty=format:"%h" -1)'" -o firefly-hdwallet.exe
deps:
	git submodule update --recursive --init
	make -C extern/filecoin-ffi all

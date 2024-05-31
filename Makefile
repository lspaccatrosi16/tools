all: bupload chi releasetool toascii scaffold tsgo

bupload_windows: 
	@echo "bupload: windows"
	@GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o ./out/windows/ ./cmd/bupload/

bupload_linux: 
	@echo "bupload: linux"
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o ./out/linux/ ./cmd/bupload/

bupload: bupload_windows bupload_linux

chi_windows: 
	@echo "chi: windows"
	@GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o ./out/windows/ ./cmd/chi/

chi_linux: 
	@echo "chi: linux"
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o ./out/linux/ ./cmd/chi/

chi: chi_windows chi_linux

releasetool_windows: 
	@echo "releasetool: windows"
	@GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o ./out/windows/ ./cmd/releasetool/

releasetool_linux: 
	@echo "releasetool: linux"
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o ./out/linux/ ./cmd/releasetool/

releasetool: releasetool_windows releasetool_linux

toascii_windows: 
	@echo "toascii: windows"
	@GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o ./out/windows/ ./cmd/toascii/

toascii_linux: 
	@echo "toascii: linux"
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o ./out/linux/ ./cmd/toascii/

toascii: toascii_windows toascii_linux

scaffold_windows: 
	@echo "scaffold: windows"
	@GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o ./out/windows/ ./cmd/scaffold/

scaffold_linux: 
	@echo "scaffold: linux"
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o ./out/linux/ ./cmd/scaffold/

scaffold: scaffold_windows scaffold_linux

tsgo_windows: 
	@echo "tsgo: windows"
	@GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o ./out/windows/ ./cmd/tsgo/

tsgo_linux: 
	@echo "tsgo: linux"
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o ./out/linux/ ./cmd/tsgo/

tsgo: tsgo_windows tsgo_linux
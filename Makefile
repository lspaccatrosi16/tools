all: bupload chi releasetool toascii scaffold tsgo

bupload_windows: 
	@echo "bupload: windows"
	@GOOS=windows go build -ldflags="-s -w" -o ./out/windows/ ./cmd/bupload/

bupload_linux: 
	@echo "bupload: linux"
	@GOOS=linux go build -ldflags="-s -w" -o ./out/linux/ ./cmd/bupload/

bupload_darwin: 
	@echo "bupload: darwin"
	@GOOS=darwin go build -ldflags="-s -w" -o ./out/darwin/ ./cmd/bupload/

bupload: bupload_windows bupload_linux bupload_darwin

chi_windows: 
	@echo "chi: windows"
	@GOOS=windows go build -ldflags="-s -w" -o ./out/windows/ ./cmd/chi/

chi_linux: 
	@echo "chi: linux"
	@GOOS=linux go build -ldflags="-s -w" -o ./out/linux/ ./cmd/chi/

chi_darwin: 
	@echo "chi: darwin"
	@GOOS=darwin go build -ldflags="-s -w" -o ./out/darwin/ ./cmd/chi/

chi: chi_windows chi_linux chi_darwin

releasetool_windows: 
	@echo "releasetool: windows"
	@GOOS=windows go build -ldflags="-s -w" -o ./out/windows/ ./cmd/releasetool/

releasetool_linux: 
	@echo "releasetool: linux"
	@GOOS=linux go build -ldflags="-s -w" -o ./out/linux/ ./cmd/releasetool/

releasetool_darwin: 
	@echo "releasetool: darwin"
	@GOOS=darwin go build -ldflags="-s -w" -o ./out/darwin/ ./cmd/releasetool/

releasetool: releasetool_windows releasetool_linux releasetool_darwin

toascii_windows: 
	@echo "toascii: windows"
	@GOOS=windows go build -ldflags="-s -w" -o ./out/windows/ ./cmd/toascii/

toascii_linux: 
	@echo "toascii: linux"
	@GOOS=linux go build -ldflags="-s -w" -o ./out/linux/ ./cmd/toascii/

toascii_darwin: 
	@echo "toascii: darwin"
	@GOOS=darwin go build -ldflags="-s -w" -o ./out/darwin/ ./cmd/toascii/

toascii: toascii_windows toascii_linux toascii_darwin

scaffold_windows: 
	@echo "scaffold: windows"
	@GOOS=windows go build -ldflags="-s -w" -o ./out/windows/ ./cmd/scaffold/

scaffold_linux: 
	@echo "scaffold: linux"
	@GOOS=linux go build -ldflags="-s -w" -o ./out/linux/ ./cmd/scaffold/

scaffold_darwin: 
	@echo "scaffold: darwin"
	@GOOS=darwin go build -ldflags="-s -w" -o ./out/darwin/ ./cmd/scaffold/

scaffold: scaffold_windows scaffold_linux scaffold_darwin

tsgo_windows: 
	@echo "tsgo: windows"
	@GOOS=windows go build -ldflags="-s -w" -o ./out/windows/ ./cmd/tsgo/

tsgo_linux: 
	@echo "tsgo: linux"
	@GOOS=linux go build -ldflags="-s -w" -o ./out/linux/ ./cmd/tsgo/

tsgo_darwin: 
	@echo "tsgo: darwin"
	@GOOS=darwin go build -ldflags="-s -w" -o ./out/darwin/ ./cmd/tsgo/

tsgo: tsgo_windows tsgo_linux tsgo_darwin
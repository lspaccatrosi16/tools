all: bupload chi releasetool toAscii scaffold

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

toAscii_windows: 
	@echo "toAscii: windows"
	@GOOS=windows go build -ldflags="-s -w" -o ./out/windows/ ./cmd/toAscii/

toAscii_linux: 
	@echo "toAscii: linux"
	@GOOS=linux go build -ldflags="-s -w" -o ./out/linux/ ./cmd/toAscii/

toAscii_darwin: 
	@echo "toAscii: darwin"
	@GOOS=darwin go build -ldflags="-s -w" -o ./out/darwin/ ./cmd/toAscii/

toAscii: toAscii_windows toAscii_linux toAscii_darwin

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
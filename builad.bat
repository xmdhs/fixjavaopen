go install github.com/akavel/rsrc@latest
rsrc -manifest nac.manifest -o nac.syso
go build -trimpath -ldflags "-w -s"
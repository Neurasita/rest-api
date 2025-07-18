GO=go
CMD_DIR=cmd
OUT_DIR=dist


build:
	@${GO} build -o ${OUT_DIR}/server ${CMD_DIR}/server/main.go

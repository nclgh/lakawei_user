mkdir -p output/bin output/conf
export GO15VENDOREXPERIMENT="1"
go build -ldflags -o ./output/bin/${RUN_NAME}
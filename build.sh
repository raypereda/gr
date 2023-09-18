GOOS=darwin  GOARCH=arm64 go build -o gr--mac-apple-silicon
GOOS=darwin  GOARCH=amd64 go build -o gr--mac-intel  
GOOS=windows GOARCH=amd64 go build -o gr.exe--windows
GOOS=linux   GOARCH=amd64 go build -o gr--linux-intel


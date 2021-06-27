# sample-go-echo-api

### air
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
air -v


### gangsta/gin
go get github.com/codegangsta/gin
gin -h
gin --all -i run main.go



echo 'Getting goat'
go get -u github.com/yosssi/goat/...
echo 'Getting less-go lessc compiler'
go get -u github.com/kib357/less-go
echo 'Getting cobra dependency'
go get -u github.com/spf13/cobra
echo 'Installing lessc'
cd $GOPATH/src/github.com/kib357/less-go/lessc
go install
echo 'Installing goat'
cd $GOPATH/src/github.com/yosssi/goat
go install
echo 'Done'

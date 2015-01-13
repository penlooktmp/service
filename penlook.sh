TMP=`pwd`
cd $TMP/penlook/library
go build github.com/penlook/service/penlook/library
cd $TMP/penlook
go build github.com/penlook/service/penlook
cd $TMP

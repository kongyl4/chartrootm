module server

go 1.13

require github.com/common/message v0.0.0-incompatible

require github.com/utils v0.0.0-incompatible

replace github.com/common/message => ../common/message

replace github.com/utils => ../utils

require github.com/gomodule/redigo v2.0.0+incompatible

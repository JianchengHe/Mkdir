package main

import (
	"BtcGoWeb/btc"
	_ "BtcGoWeb/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	fmt.Println("hello  world !")
	result, err := btc.GetMsgByCommand("getblockhash", 0)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(result.Result)

	beego.Run()
	/**
	hello  world !
	{"jsonrpc":"2.0","id":"1608601306","method":"getblockhash","params":[0]}
	000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f
	 */

}


package main

import (
	"fmt"
	_ "goweb/httpRest/router"
)

func main() {
	var a string = "2b237ab94d384666a15ad9341cac87e7"
	b := make([]rune, len(a))
	for i := 0; i < len(a); i++ {
		b[i] = '0'
	}
	fmt.Println(string(b))
	//mem.Wg.Wait()
	//util.Test()
	//test()
	//token, err := util.SignJwt("xw","xw",60, "")
	//if err == nil {
	//	logrus.Info(token)
	//}
	//logrus.Info(len(viper.GetString("privateKey")))
	//logrus.Info(util.RandomHex(64))

	//defer config.InitLog().Close()
	//defer config.InitDB("postgres").Close()
	//http.ListenAndServe(":" + viper.GetString("port"), mem.G)

	//ethClient, err := config.InitEthClient()
	//if err != nil {
	//	panic(err)
	//}
	//defer ethClient.Close()

	//trans, err := ethereum.SendTrans("0x4022bbbf762c8ed4d60c2eaa1b67a19dcbfa0c15")
	//if err != nil {
	//	panic(err)
	//}
	//log.Info(trans)
	//
	//transList, err := ethereum.GetTransactions()
	//if err != nil {
	//	panic(err)
	//}
	//log.Info(transList)

	//head, _ := ethClient.HeaderByNumber(context.Background(), nil)
	//block, _ := ethClient.BlockByNumber(context.Background(), big.NewInt(head.Number.Int64()))
	//mysql, err := config.InitDB("mysql")
	//if err != nil {
	//	panic(err)
	//}
	//defer mysql.Close()
	//
	//postgres, err := config.InitDB("postgres")
	//if err != nil {
	//	panic(err)
	//}
	//defer postgres.Close()

	//a, _ := util.Bcrypt("aa")
	//log.Info(a)
	//b, _ := util.SignJwt(util.Context{ID: 123112, Username: "fsfafa"}, "")
	//log.Info(b)
	//c, _ := util.Parse(b, "")
	//log.Info(c)

	//
	//rs, err := ethereum.ListAccounts()
	//log.Info(rs)
	//
	//rs, err = ethereum.NodeInfo()
	//if err != nil {
	//	panic(err)
	//}
	//log.Info(rs)
	//
	//rs, err = ethereum.Peers()
	//if err != nil {
	//	panic(err)
	//}
	//log.Info(rs)

	//peerUrls := "enode://aa7eed134813387299e2dfc00c7fa664d3bd308caf1af87383c107052f3500de69ed1c452c42fdd06c6102132a5d76869c20aaa83413c0bd27c63a24e4961178@192.168.0.33:30303?discport=0"
	//rs, err = ethereum.AddPeer(peerUrls)
	//if err != nil {
	//	panic(err)
	//}
	//log.Info(rs)
	//crypto.HexToECDSA("")

	//sig, pub, err := util.SignECDSA("hello")
	//if err != nil {
	//	panic(err)
	//}
	//log.Info(sig)
	//log.Info(pub)
	//
	//matched, err:= util.VerifyECDSA("helo", sig, pub)
	//if err != nil {
	//	panic(err)
	//}
	//log.Info(matched)
	//
	//ct, err := util.EncryptEcies("hello", pub)
	//if err != nil {
	//	panic(err)
	//}
	//log.Info(ct)
	//
	//dect, err := util.DecryptEcies(ct)
	//if err != nil {
	//	panic(err)
	//}
	//log.Info(dect)
}

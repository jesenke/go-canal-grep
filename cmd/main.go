package main

import (
	"flag"
	"go-canal-grep/util"
	"runtime"
)

//
//package main
//
//import (
//"flag"
//"os"
//"os/signal"
//"runtime"
//"syscall"
//"time"
//
//"github.com/kpango/glg"
//"github.com/olachat/banban_server"
//"github.com/olachat/banban_server/cli_mysql_to_es/app"
//"github.com/olachat/banban_server/cli_mysql_to_es/conf"
//"github.com/olachat/banban_server/cli_mysql_to_es/instance/oversea/cities_index"
//"github.com/olachat/banban_server/cli_mysql_to_es/instance/oversea/room_new"
//"github.com/olachat/banban_server/cli_mysql_to_es/instance/oversea/song_index"
//"github.com/olachat/banban_server/cli_mysql_to_es/instance/oversea/user_matching"
//"github.com/olachat/banban_server/cli_mysql_to_es/instance/oversea/user_new"
//"github.com/olachat/banban_server/cli_mysql_to_es/model"
//_ "github.com/olachat/banban_server/cli_mysql_to_es/mq"
//)

func main() {
	util.Init()
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	//解析命令行参数
	var startWithFull string
	var dropIndexName string
	var env string
	var incremental bool
	
	flag.StringVar(&env, "env", "", "操作哪个配置文件数据")
	flag.StringVar(&startWithFull, "startWithFull", "", "哪些索引全量开始")
	flag.StringVar(&dropIndexName, "dropIndexName", "", "哪些索引先删除在重建")
	flag.BoolVar(&incremental, "incremental", false, "全量索引起始主键id")
	
	flag.Parse()
	util.Log.Info("开始服务env:%v,startWithFull:%v,dropIndexName:%v,incremental:%v", env, startWithFull, dropIndexName, incremental)
	//log.Fatal().Info("start parse config")
	//conf.ParseConfig(env)
	//
	//// print line number
	//if conf.IsDev {
	//	glg.Get().SetLineTraceMode(glg.TraceLineShort)
	//}
	//
	//glg.Info("start init db")
	//model.InitDb()
	////坚持是否已经有过从数据库初始化数据
	//
	//glg.Info("start new app")
	////记录当前时间
	//control := app.NewApp()
	//
	//song := song_index.GetIndex()
	//room := room_new.GetIndex()
	////roomTest := room_test.GetIndex()
	//user := user_new.GetIndex()
	//cities := cities_index.GetIndex()
	//userMatching := user_matching.GetIndex()
	//family := family_new.GetIndex()
	//
	//control.Add(song)
	//control.Add(user)
	//control.Add(room)
	//control.Add(cities)
	//control.Add(userMatching)
	//control.Add(family)
	////control.Add(roomTest)
	//go control.Start(startWithFull, dropIndexName, incremental)
	//
	////setupBatchImport(room, user, userMatching)
	//
	//sign := make(chan os.Signal, 1)
	//signal.Notify(sign, os.Interrupt, syscall.SIGTERM, syscall.SIGUSR1, syscall.SIGUSR2)
	//
	//glg.Infof("Adjust time to begin ticker: %#v | %#v", time.Now().Unix(), time.Now())
	//
	//for time.Now().Unix()%2 != 0 {
	//	time.Sleep(1 * time.Millisecond)
	//}
	//
	//glg.Infof("Begin ticker: %#v | %#v", time.Now().Unix(), time.Now())
	//timerLog := time.NewTicker(2 * time.Second)
	//
	//for {
	//	select {
	//	case time := <-timerLog.C:
	//		num := control.GetProcessNum()
	//		if time.Second()%10 == 0 {
	//			glg.Infof("timerLog.C:: time:%#v, ProcessNum:%d", time.Format("20060102_15_04_05.000"), num)
	//		}
	//	case s := <-sign:
	//		glg.Infof("receive sign data [%d]", s)
	//
	//		if control.IsFullComplete() {
	//			glg.Info("control stopped before")
	//			control.Stop()
	//			glg.Info("quit by auto")
	//		} else {
	//
	//			glg.Info("quitting by exit signal...")
	//		}
	//		glg.Info("quit wait group by exit")
	//
	//		time.Sleep(time.Second * 1)
	//		return
	//	}
	//}
}

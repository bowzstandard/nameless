package watcher

import(
    "os"
    "os/exec"
    "bytes"
    "github.com/go-fsnotify/fsnotify"
    log "github.com/Sirupsen/logrus"
    "bowz-util/interfaces"
)

type FsWatcher struct{
    dealer interfaces.WatchDealer
}

func (obj *FsWatcher)Watch(){
    var watcher *fsnotify.Watcher
    var err error

    if watcher,err = fsnotify.NewWatcher();err!=nil{
        log.Fatal(err)
    }
    defer watcher.Close()

    done := make(chan bool)

    log.Info("ディレクトリ監視スタート ...")

    go func(){
        for {
            select {
            case event := <- watcher.Events:
                if event.Op&fsnotify.Write != fsnotify.Write {
                    continue
                }
                obj.deal()
            case err := <-watcher.Errors:
                log.Fatal("error:watcher-0001")
                log.Fatal(err.Error())
            }
        }
    }()

    if err = watcher.Add(src);err!=nil{
        log.Fatal(err)
    }

    <-done
}

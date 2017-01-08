package watcher

import(
    "bowz-util/interfaces"
)

func GenFs(dealer interfaces.WatchDealer)*FsWatcher{
    var obj *FsWatcher
    obj = &FsWatcher{
        dealer:dealer,
    }

    return obj
}
package talker

import(
)

func GenCmd(cmd string,args ...string)(*CmdTalker){
    var tmp []string

    if len(args)>0{
        tmp = []string{}
        for _,val := range args {
            tmp = append(tmp,val)
        }
    }

    var obj *CmdTalker
    obj = &CmdTalker{
        cmdstr:cmd,
        args:args,
    }

    return obj
}

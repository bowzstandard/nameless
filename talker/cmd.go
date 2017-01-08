package talker

import(
    "os"
    "os/exec"
    "bytes"
    "errors"
    log "github.com/Sirupsen/logrus"
)

type CmdTalker struct{
    cmdstr string
    args []string
    out bytes.Buffer
    err bytes.Buffer
    cmd *exec.Cmd
}

func (obj *CmdTalker)Exec(){
    var err error

    if obj.args!=nil{
        obj.cmd = exec.Command(obj.cmdstr,obj.args...)
    }else{
        obj.cmd = exec.Command(obj.cmdstr)
    }

    obj.cmd.Stdout = &obj.out
    obj.cmd.Stderr = &obj.err

    log.Info("Now Executing => "+obj.cmdstr)

    if err = obj.cmd.Run();err!=nil {
        log.Warn("error occured! => cmdTalker:0001")
        log.Warn("exec:"+obj.cmdstr)
        log.Warn(err)
        log.Warn(obj.err.String())
        os.Exit(1)
        return
    }
}

func (obj *CmdTalker)Output()bytes.Buffer{
    return obj.out
}

func (obj *CmdTalker)Error()error{
    if(obj.err.String()==""){
        return nil
    }
    return errors.New(obj.err.String())
}
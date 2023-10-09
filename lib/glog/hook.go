package glog

type HookFieldFunc func([]Field)

func defaultHook(_ []Field) {

}

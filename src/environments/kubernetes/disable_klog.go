package kubernetes

import (
	"flag"
	"io/ioutil"
	"k8s.io/klog"
)

func DisableKlog() {
	klog.SetOutput(ioutil.Discard)
	klog.InitFlags(nil)
	flag.Set("logtostderr", "false")
	flag.Set("stderrthreshold", "4")
	flag.Parse()
	klog.Warningln("sth")
}

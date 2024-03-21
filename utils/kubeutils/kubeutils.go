package kubeutils

import (
	"krm-backend/utils/logs"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type KubeUtilsInterface interface {
	Create(string) error
	Delete(string, string, *int64) error
	DeleteList(string, []string, *int64) error
	Update(string) error
	Get(string, string) (interface{}, error)
	List(string, string, string) (interface{}, error)
}

type ResourceInstance struct {
	Kubeconfig string
	ClientSet  *kubernetes.Clientset
}

func (r *ResourceInstance) Init(kubeconfig string) {
	r.Kubeconfig = kubeconfig
	restConfig, err := clientcmd.RESTConfigFromKubeConfig([]byte(r.Kubeconfig))
	if err != nil {
		logs.Error(map[string]interface{}{"error": err.Error()}, "解析kubeconfig错误")
		msg := "解析kubeconfig错误" + err.Error()
		panic(msg)
	}
	clientSet, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		logs.Error(map[string]interface{}{"error": err.Error()}, "实例化clientSet错误")
		msg := "实例化clientSet错误" + err.Error()
		panic(msg)
	}
	r.ClientSet = clientSet
}

package initcontroller

import (
	"context"
	"krm-backend/config"
	"krm-backend/utils/logs"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func init() {
	logs.Deubg(nil, "初始化数据")
	restConfig, err := clientcmd.BuildConfigFromFlags("", "/Users/joe/.kube/config")
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		panic(err.Error())
	}

	_, err = clientset.CoreV1().Namespaces().Get(context.TODO(), config.MetaDataNamespace, metav1.GetOptions{})
	if err != nil {
		// 判断krm运行的命名空间是否存在, 不存在创建
		var metaNamespace corev1.Namespace
		metaNamespace.Name = "mgmt"
		_, err = clientset.CoreV1().Namespaces().Create(context.TODO(), &metaNamespace, metav1.CreateOptions{})
		if err != nil {
			logs.Error(map[string]interface{}{"message": "创建namespace失败"}, "")
			panic(err.Error())
		}
	}
	inClusterInfo, _ := clientset.Discovery().ServerVersion()
	logs.Info(map[string]interface{}{"Namespace": config.MetaDataNamespace, "Version": inClusterInfo.GitVersion}, "元数据命名空间初始化成功")

}

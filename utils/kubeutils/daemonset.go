package kubeutils

import (
	"context"
	"krm-backend/utils/logs"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	typedv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
)

type DaemonSet struct {
	InstanceInterface typedv1.AppsV1Interface
	Item              *appsv1.DaemonSet
}

func NewDaemonSet(kubecofig string, item *appsv1.DaemonSet) *DaemonSet {
	// 调用实例化kubeutils中的ResourceInstance并调用init, 配置默认值和生成clientset
	instance := ResourceInstance{}
	instance.Init(kubecofig)
	// 定义DaemonSet实例
	resource := DaemonSet{}
	resource.InstanceInterface = instance.ClientSet.AppsV1()
	resource.Item = item
	return &resource
}

// 创建DaemonSet资源
func (p *DaemonSet) Create(namespace string) error {
	logs.Info(map[string]interface{}{"名称": p.Item.Name, "命名空间": namespace}, "创建DaemonSet")
	_, err := p.InstanceInterface.DaemonSets(namespace).Create(context.TODO(), p.Item, metav1.CreateOptions{})
	return err
}

func (p *DaemonSet) Delete(namespace string, name string, gracePeriodSeconds *int64) error {
	logs.Info(map[string]interface{}{"名称": name, "命名空间": namespace}, "删除DaemonSet")
	deleteOption := metav1.DeleteOptions{}
	if gracePeriodSeconds != nil {
		deleteOption.GracePeriodSeconds = gracePeriodSeconds
	}
	err := p.InstanceInterface.DaemonSets(namespace).Delete(context.TODO(), name, deleteOption)
	return err
}

func (p *DaemonSet) DeleteList(namespace string, nameList []string, gracePeriodSeconds *int64) error {
	for _, name := range nameList {
		p.Delete(namespace, name, gracePeriodSeconds)
	}
	return nil
}

func (p *DaemonSet) Update(namespace string) error {
	logs.Info(map[string]interface{}{"名称": p.Item.Name, "命名空间": namespace}, "更新DaemonSet")
	_, err := p.InstanceInterface.DaemonSets(namespace).Update(context.TODO(), p.Item, metav1.UpdateOptions{})
	return err
}

func (p *DaemonSet) List(namespace, labelSelector, fieldSelector string) (items interface{}, err error) {
	logs.Info(map[string]interface{}{"名称": p.Item.Name, "命名空间": namespace}, "DaemonSet列表")
	var listOptions metav1.ListOptions
	listOptions.LabelSelector = labelSelector
	listOptions.FieldSelector = fieldSelector
	StatefulSetList, err := p.InstanceInterface.DaemonSets(namespace).List(context.TODO(), listOptions)
	items = StatefulSetList.Items
	return items, err
}

func (p *DaemonSet) Get(namespace, name string) (item interface{}, err error) {
	logs.Info(map[string]interface{}{"名称": name, "命名空间": namespace}, "DaemonSet对象")
	StatefulSet, err := p.InstanceInterface.DaemonSets(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	item = StatefulSet
	return item, err
}

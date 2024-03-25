package kubeutils

import (
	"context"
	"krm-backend/utils/logs"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	typedv1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

type ConfigMap struct {
	InstanceInterface typedv1.CoreV1Interface
	Item              *corev1.ConfigMap
}

func NewConfigMap(kubecofig string, item *corev1.ConfigMap) *ConfigMap {
	// 调用实例化kubeutils中的ResourceInstance并调用init, 配置默认值和生成clientset
	instance := ResourceInstance{}
	instance.Init(kubecofig)
	// 定义configmap实例
	resource := ConfigMap{}
	resource.InstanceInterface = instance.ClientSet.CoreV1()
	resource.Item = item
	return &resource
}

// 创建configmap资源
func (p *ConfigMap) Create(namespace string) error {
	logs.Info(map[string]interface{}{"名称": p.Item.Name, "命名空间": namespace}, "创建ConfigMap")
	_, err := p.InstanceInterface.ConfigMaps(namespace).Create(context.TODO(), p.Item, metav1.CreateOptions{})
	return err
}

func (p *ConfigMap) Delete(namespace string, name string, gracePeriodSeconds *int64) error {
	logs.Info(map[string]interface{}{"名称": name, "命名空间": namespace}, "删除ConfigMap")
	deleteOption := metav1.DeleteOptions{}
	if gracePeriodSeconds != nil {
		deleteOption.GracePeriodSeconds = gracePeriodSeconds
	}
	err := p.InstanceInterface.ConfigMaps(namespace).Delete(context.TODO(), name, deleteOption)
	return err
}

func (p *ConfigMap) DeleteList(namespace string, nameList []string, gracePeriodSeconds *int64) error {
	for _, name := range nameList {
		p.Delete(namespace, name, gracePeriodSeconds)
	}
	return nil
}

func (p *ConfigMap) Update(namespace string) error {
	logs.Info(map[string]interface{}{"名称": p.Item.Name, "命名空间": namespace}, "更新ConfigMap")
	_, err := p.InstanceInterface.ConfigMaps(namespace).Update(context.TODO(), p.Item, metav1.UpdateOptions{})
	return err
}

func (p *ConfigMap) List(namespace, labelSelector, fieldSelector string) (items interface{}, err error) {
	logs.Info(map[string]interface{}{"名称": p.Item.Name, "命名空间": namespace}, "ConfigMap列表")
	var listOptions metav1.ListOptions
	listOptions.LabelSelector = labelSelector
	listOptions.FieldSelector = fieldSelector
	configMapList, err := p.InstanceInterface.ConfigMaps(namespace).List(context.TODO(), listOptions)
	items = configMapList.Items
	return items, err
}

func (p *ConfigMap) Get(namespace, name string) (item interface{}, err error) {
	logs.Info(map[string]interface{}{"名称": name, "命名空间": namespace}, "ConfigMap对象")
	configMap, err := p.InstanceInterface.ConfigMaps(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	item = configMap
	return item, err
}

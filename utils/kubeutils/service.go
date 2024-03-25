package kubeutils

import (
	"context"
	"krm-backend/utils/logs"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	typedv1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

type Service struct {
	InstanceInterface typedv1.CoreV1Interface
	Item              *corev1.Service
}

func NewService(kubecofig string, item *corev1.Service) *Service {
	// 调用实例化kubeutils中的ResourceInstance并调用init, 配置默认值和生成clientset
	instance := ResourceInstance{}
	instance.Init(kubecofig)
	// 定义service实例
	resource := Service{}
	resource.InstanceInterface = instance.ClientSet.CoreV1()
	resource.Item = item
	return &resource
}

// 创建service资源
func (p *Service) Create(namespace string) error {
	logs.Info(map[string]interface{}{"名称": p.Item.Name, "命名空间": namespace}, "创建Service")
	_, err := p.InstanceInterface.Services(namespace).Create(context.TODO(), p.Item, metav1.CreateOptions{})
	return err
}

func (p *Service) Delete(namespace string, name string, gracePeriodSeconds *int64) error {
	logs.Info(map[string]interface{}{"名称": name, "命名空间": namespace}, "删除Service")
	deleteOption := metav1.DeleteOptions{}
	if gracePeriodSeconds != nil {
		deleteOption.GracePeriodSeconds = gracePeriodSeconds
	}
	err := p.InstanceInterface.Services(namespace).Delete(context.TODO(), name, deleteOption)
	return err
}

func (p *Service) DeleteList(namespace string, nameList []string, gracePeriodSeconds *int64) error {
	for _, name := range nameList {
		p.Delete(namespace, name, gracePeriodSeconds)
	}
	return nil
}

func (p *Service) Update(namespace string) error {
	logs.Info(map[string]interface{}{"名称": p.Item.Name, "命名空间": namespace}, "更新Service")
	_, err := p.InstanceInterface.Services(namespace).Update(context.TODO(), p.Item, metav1.UpdateOptions{})
	return err
}

func (p *Service) List(namespace, labelSelector, fieldSelector string) (items interface{}, err error) {
	logs.Info(map[string]interface{}{"名称": p.Item.Name, "命名空间": namespace}, "Service列表")
	var listOptions metav1.ListOptions
	listOptions.LabelSelector = labelSelector
	listOptions.FieldSelector = fieldSelector
	podList, err := p.InstanceInterface.Services(namespace).List(context.TODO(), listOptions)
	items = podList.Items
	return items, err
}

func (p *Service) Get(namespace, name string) (item interface{}, err error) {
	logs.Info(map[string]interface{}{"名称": name, "命名空间": namespace}, "Service对象")
	pod, err := p.InstanceInterface.Services(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	item = pod
	return item, err
}

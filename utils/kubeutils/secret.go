package kubeutils

import (
	"context"
	"krm-backend/utils/logs"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	typedv1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

type Secret struct {
	InstanceInterface typedv1.CoreV1Interface
	Item              *corev1.Secret
}

func NewSecret(kubecofig string, item *corev1.Secret) *Secret {
	// 调用实例化kubeutils中的ResourceInstance并调用init, 配置默认值和生成clientset
	instance := ResourceInstance{}
	instance.Init(kubecofig)
	// 定义secret实例
	resource := Secret{}
	resource.InstanceInterface = instance.ClientSet.CoreV1()
	resource.Item = item
	return &resource
}

// 创建secret资源
func (p *Secret) Create(namespace string) error {
	logs.Info(map[string]interface{}{"名称": p.Item.Name, "命名空间": namespace}, "创建Secret")
	_, err := p.InstanceInterface.Secrets(namespace).Create(context.TODO(), p.Item, metav1.CreateOptions{})
	return err
}

func (p *Secret) Delete(namespace string, name string, gracePeriodSeconds *int64) error {
	logs.Info(map[string]interface{}{"名称": name, "命名空间": namespace}, "删除Secret")
	deleteOption := metav1.DeleteOptions{}
	if gracePeriodSeconds != nil {
		deleteOption.GracePeriodSeconds = gracePeriodSeconds
	}
	err := p.InstanceInterface.Secrets(namespace).Delete(context.TODO(), name, deleteOption)
	return err
}

func (p *Secret) DeleteList(namespace string, nameList []string, gracePeriodSeconds *int64) error {
	for _, name := range nameList {
		p.Delete(namespace, name, gracePeriodSeconds)
	}
	return nil
}

func (p *Secret) Update(namespace string) error {
	logs.Info(map[string]interface{}{"名称": p.Item.Name, "命名空间": namespace}, "更新Secret")
	_, err := p.InstanceInterface.Secrets(namespace).Update(context.TODO(), p.Item, metav1.UpdateOptions{})
	return err
}

func (p *Secret) List(namespace, labelSelector, fieldSelector string) (items interface{}, err error) {
	logs.Info(map[string]interface{}{"名称": p.Item.Name, "命名空间": namespace}, "Secret列表")
	var listOptions metav1.ListOptions
	listOptions.LabelSelector = labelSelector
	listOptions.FieldSelector = fieldSelector
	secretList, err := p.InstanceInterface.Secrets(namespace).List(context.TODO(), listOptions)
	items = secretList.Items
	return items, err
}

func (p *Secret) Get(namespace, name string) (item interface{}, err error) {
	logs.Info(map[string]interface{}{"名称": name, "命名空间": namespace}, "Secret对象")
	secret, err := p.InstanceInterface.Secrets(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	item = secret
	return item, err
}

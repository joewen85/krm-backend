package kubeutils

import (
	"context"
	"krm-backend/utils/logs"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	typedv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
)

type Deployment struct {
	InstanceInterface typedv1.AppsV1Interface
	Item              *appsv1.Deployment
}

func NewDeployment(kubecofig string, item *appsv1.Deployment) *Deployment {
	// 调用实例化kubeutils中的ResourceInstance并调用init, 配置默认值和生成clientset
	instance := ResourceInstance{}
	instance.Init(kubecofig)
	// 定义deployment实例
	resource := Deployment{}
	resource.InstanceInterface = instance.ClientSet.AppsV1()
	resource.Item = item
	return &resource
}

// 创建deployment资源
func (p *Deployment) Create(namespace string) error {
	logs.Info(map[string]interface{}{"名称": p.Item.Name, "命名空间": namespace}, "创建Deployment")
	_, err := p.InstanceInterface.Deployments(namespace).Create(context.TODO(), p.Item, metav1.CreateOptions{})
	return err
}

func (p *Deployment) Delete(namespace string, name string, gracePeriodSeconds *int64) error {
	logs.Info(map[string]interface{}{"名称": name, "命名空间": namespace}, "删除Deployment")
	deleteOption := metav1.DeleteOptions{}
	if gracePeriodSeconds != nil {
		deleteOption.GracePeriodSeconds = gracePeriodSeconds
	}
	err := p.InstanceInterface.Deployments(namespace).Delete(context.TODO(), name, deleteOption)
	return err
}

func (p *Deployment) DeleteList(namespace string, nameList []string, gracePeriodSeconds *int64) error {
	for _, name := range nameList {
		p.Delete(namespace, name, gracePeriodSeconds)
	}
	return nil
}

func (p *Deployment) Update(namespace string) error {
	logs.Info(map[string]interface{}{"名称": p.Item.Name, "命名空间": namespace}, "更新Deployment")
	_, err := p.InstanceInterface.Deployments(namespace).Update(context.TODO(), p.Item, metav1.UpdateOptions{})
	return err
}

func (p *Deployment) List(namespace, labelSelector, fieldSelector string) (items interface{}, err error) {
	logs.Info(map[string]interface{}{"名称": p.Item.Name, "命名空间": namespace}, "Deployment列表")
	var listOptions metav1.ListOptions
	listOptions.LabelSelector = labelSelector
	listOptions.FieldSelector = fieldSelector
	deploymentList, err := p.InstanceInterface.Deployments(namespace).List(context.TODO(), listOptions)
	items = deploymentList.Items
	return items, err
}

func (p *Deployment) Get(namespace, name string) (item interface{}, err error) {
	logs.Info(map[string]interface{}{"名称": name, "命名空间": namespace}, "Deployment对象")
	deployment, err := p.InstanceInterface.Deployments(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	item = deployment
	return item, err
}

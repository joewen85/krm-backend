package kubeutils

import (
	"context"
	"krm-backend/utils/logs"

	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	typedv1 "k8s.io/client-go/kubernetes/typed/batch/v1"
)

type Cronjob struct {
	InstanceInterface typedv1.BatchV1Interface
	Item              *batchv1.CronJob
}

func NewCronjob(kubecofig string, item *batchv1.CronJob) *Cronjob {
	// 调用实例化kubeutils中的ResourceInstance并调用init, 配置默认值和生成clientset
	instance := ResourceInstance{}
	instance.Init(kubecofig)
	// 定义DaemonSet实例
	resource := Cronjob{}
	resource.InstanceInterface = instance.ClientSet.BatchV1()
	resource.Item = item
	return &resource
}

// 创建DaemonSet资源
func (p *Cronjob) Create(namespace string) error {
	logs.Info(map[string]interface{}{"名称": p.Item.Name, "命名空间": namespace}, "创建Cronjob")
	_, err := p.InstanceInterface.CronJobs(namespace).Create(context.TODO(), p.Item, metav1.CreateOptions{})
	return err
}

func (p *Cronjob) Delete(namespace string, name string, gracePeriodSeconds *int64) error {
	logs.Info(map[string]interface{}{"名称": name, "命名空间": namespace}, "删除Cronjob")
	deleteOption := metav1.DeleteOptions{}
	if gracePeriodSeconds != nil {
		deleteOption.GracePeriodSeconds = gracePeriodSeconds
	}
	err := p.InstanceInterface.CronJobs(namespace).Delete(context.TODO(), name, deleteOption)
	return err
}

func (p *Cronjob) DeleteList(namespace string, nameList []string, gracePeriodSeconds *int64) error {
	for _, name := range nameList {
		p.Delete(namespace, name, gracePeriodSeconds)
	}
	return nil
}

func (p *Cronjob) Update(namespace string) error {
	logs.Info(map[string]interface{}{"名称": p.Item.Name, "命名空间": namespace}, "更新Cronjob")
	_, err := p.InstanceInterface.CronJobs(namespace).Update(context.TODO(), p.Item, metav1.UpdateOptions{})
	return err
}

func (p *Cronjob) List(namespace, labelSelector, fieldSelector string) (items interface{}, err error) {
	logs.Info(map[string]interface{}{"名称": p.Item.Name, "命名空间": namespace}, "Cronjob列表")
	var listOptions metav1.ListOptions
	listOptions.LabelSelector = labelSelector
	listOptions.FieldSelector = fieldSelector
	StatefulSetList, err := p.InstanceInterface.CronJobs(namespace).List(context.TODO(), listOptions)
	items = StatefulSetList.Items
	return items, err
}

func (p *Cronjob) Get(namespace, name string) (item interface{}, err error) {
	logs.Info(map[string]interface{}{"名称": name, "命名空间": namespace}, "Cronjob对象")
	StatefulSet, err := p.InstanceInterface.CronJobs(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	item = StatefulSet
	return item, err
}

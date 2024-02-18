package main

import (
	"context"
	"encoding/json"
	"fmt"
	"krm-backend/utils/logs"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func mainbak() {
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/joe/.kube/config")
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	// get pods list
	pods, err := clientset.CoreV1().Pods("devops").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logs.Error(nil, "查询Pods列表失败")
	}
	fmt.Printf("there are %d pods in the cluster\n", len(pods.Items))

	// get deployments list
	deployments, _ := clientset.AppsV1().Deployments("").List(context.TODO(), metav1.ListOptions{})
	deploymentItems := deployments.Items
	for _, deply := range deploymentItems {
		fmt.Printf("当前资源名字是: %s, namesapce: %s\n", deply.Name, deply.Namespace)
	}

	// get pod resource object
	pod, err := clientset.CoreV1().Pods("devops").Get(context.TODO(), "devops-backend-56b4d66d89-9bbvt", metav1.GetOptions{})
	if err != nil {
		fmt.Printf("获取pod失败")
	}
	fmt.Print("获取pod对象", pod.Labels)

	// get deployment resourece object
	deploy, err := clientset.AppsV1().Deployments("default").Get(context.TODO(), "nginx", metav1.GetOptions{})
	if err != nil {
		fmt.Print("获取deployment失败")
	}
	fmt.Print("获取deployment对象", deploy.Spec.Template.Spec.Containers)

	// delete pod
	errr := clientset.CoreV1().Pods("default").Delete(context.TODO(), "nginx-748c667d99-jrb2b", metav1.DeleteOptions{})
	if errr != nil {
		fmt.Print("删除pod失败")
	}

	// update
	// deployDetail, err := clientset.AppsV1().Deployments("default").Get(context.TODO(), "nginx", metav1.GetOptions{})
	// newReplica := int32(2)
	// deployDetail.Spec.Replicas = &newReplica
	// labels := deployDetail.Labels
	// labels["testDev"] = "test_upgrade"
	// newdeploy, err := clientset.AppsV1().Deployments("default").Update(context.TODO(), deployDetail, metav1.UpdateOptions{})
	// if err != nil {
	// 	fmt.Print("修改deployment失败", err.Error())
	// } else {
	// 	fmt.Print(newdeploy.Labels)
	// }

	// create namespace

	var newNamespaece corev1.Namespace
	newNamespaece.Name = "test-dev"
	newNamespaceObj, err := clientset.CoreV1().Namespaces().Create(context.TODO(), &newNamespaece, metav1.CreateOptions{})
	if err != nil {
		fmt.Print("创建命名空间失败", err.Error())
	} else {
		fmt.Print("成功", newNamespaceObj.Name)
	}

	// create deployment
	var newdeploy appsv1.Deployment
	newdeploy.Name = "test-dev-deploy"
	label := make(map[string]string)
	label["app"] = "nginx"
	label["version"] = "v1"
	newdeploy.Labels = label

	newdeploy.Spec.Selector = &metav1.LabelSelector{}
	newdeploy.Spec.Selector.MatchLabels = label

	newdeploy.Spec.Template.Labels = label

	nreplica := int32(2)
	newdeploy.Spec.Replicas = &nreplica

	// 需要先创建
	// 后报index out of range,因为没有初始化Containers就赋值. 要先初始化Containers,再通过追加方式添加容器
	// newdeploy.Spec.Template.Spec.Containers[0].Image = "nginx"
	// newdeploy.Spec.Template.Spec.Containers[0].Name = "nginx"
	var containers []corev1.Container

	// 先声明一个pod,然后赋值给newdeploy
	// 创建容器
	var container corev1.Container
	container.Image = "redis"
	container.Name = "redis"
	containers = append(containers, container)
	container.Image = "nginx"
	container.Name = "nginx"
	containers = append(containers, container)

	newdeploy.Spec.Template.Spec.Containers = containers

	_, err = clientset.AppsV1().Deployments("test-dev").Create(context.TODO(), &newdeploy, metav1.CreateOptions{})
	if err != nil {
		fmt.Print("创建deployment失败", err.Error())
	} else {
		fmt.Print("创建deployment成功")
	}

	var newdeploy2 appsv1.Deployment
	deployJson := `{
		"kind": "Deployment",
		"apiVersion": "apps/v1",
		"metadata": {
			"name": "test-dev2",
			"creationTimestamp": null,
			"labels": {
				"app": "test-dev2"
			}
		},
		"spec": {
			"replicas": 1,
			"selector": {
				"matchLabels": {
					"app": "test-dev2"
				}
			},
			"template": {
				"metadata": {
					"creationTimestamp": null,
					"labels": {
						"app": "test-dev2"
					}
				},
				"spec": {
					"containers": [
						{
							"name": "redis",
							"image": "redis",
							"resources": {}
						}
					]
				}
			},
			"strategy": {}
		},
		"status": {}
	}`
	// 通过Unmarshal将json转换为结构体
	json.Unmarshal([]byte(deployJson), &newdeploy2)
	_, err = clientset.AppsV1().Deployments("default").Create(context.TODO(), &newdeploy2, metav1.CreateOptions{})
	if err != nil {
		fmt.Print("创建deployment失败", err.Error())
	} else {
		fmt.Print("创建deployment功能")
	}
	err = clientset.AppsV1().Deployments("default").Delete(context.TODO(), "test-dev2", metav1.DeleteOptions{})
	if err != nil {
		fmt.Print("删除delpoyment失败")
	}
}

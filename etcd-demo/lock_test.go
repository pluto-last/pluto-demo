package etcd_demo

import (
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"log"
	"testing"
	"time"
)

// 基于etcd的分布式锁
func Test_Lock(t *testing.T) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	//// 创建两个单独的会话用来演示锁竞争
	//s1, err := concurrency.NewSession(cli)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer s1.Close()
	//m1 := concurrency.NewMutex(s1, "/my-lock/")
	//
	//s2, err := concurrency.NewSession(cli)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer s2.Close()
	//m2 := concurrency.NewMutex(s2, "/my-lock/")
	//
	//// 会话s1获取锁
	//if err := m1.Lock(context.TODO()); err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("acquired lock for s1")
	//
	//m2Locked := make(chan struct{})
	//go func() {
	//	defer close(m2Locked)
	//	// 等待直到会话s1释放了/my-lock/的锁
	//	if err := m2.Lock(context.TODO()); err != nil {
	//		log.Fatal(err)
	//	}
	//}()
	//
	//if err := m1.Unlock(context.TODO()); err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("released lock for s1")
	//
	//<-m2Locked
	fmt.Println("acquired lock for s2")
}

package etcd_register

import (
	"context"
	"go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

type EtcdRegister struct {
	etcdCli *clientv3.Client // etcd连接
	leaseId clientv3.LeaseID // 租约ID
	ctx     context.Context
	cancel  context.CancelFunc
}

// 创建租约
// expire 有效期/秒
func (s *EtcdRegister) CreateLease(expire int64) error {

	res, err := s.etcdCli.Grant(s.ctx, expire)
	if err != nil {
		log.Printf("createLease failed,error %v \n", err)
		return err
	}

	s.leaseId = res.ID
	return nil
}

// 绑定租约
// 将租约和对应的KEY-VALUE绑定
func (s *EtcdRegister) BindLease(key string, value string) error {

	res, err := s.etcdCli.Put(s.ctx, key, value, clientv3.WithLease(s.leaseId))
	if err != nil {
		log.Printf("bindLease failed,error %v \n", err)
		return err
	}

	log.Printf("bindLease success %v \n", res)
	return nil
}

// 续租 发送心跳，表明服务正常

func (s *EtcdRegister) KeepAlive() (<-chan *clientv3.LeaseKeepAliveResponse, error) {

	resChan, err := s.etcdCli.KeepAlive(s.ctx, s.leaseId)
	if err != nil {
		log.Printf("keepAlive failed,error %v \n", resChan)
		return resChan, err
	}

	return resChan, nil
}

func (s *EtcdRegister) Watcher(key string, resChan <-chan *clientv3.LeaseKeepAliveResponse) {

	for {
		select {
		case l := <-resChan:
			log.Printf("续约成功,val:%+v \n", l)
		case <-s.ctx.Done():
			log.Printf("续约关闭")
			return
		}
	}
}

func (s *EtcdRegister) Close() error {

	s.cancel()

	log.Printf("closed...\n")

	// 撤销租约
	s.etcdCli.Revoke(s.ctx, s.leaseId)

	return s.etcdCli.Close()
}

func NewEtcdRegister() (*EtcdRegister, error) {

	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		log.Printf("new etcd client failed,error %v \n", err)
		return nil, err
	}

	ctx, cancelFunc := context.WithCancel(context.Background())

	svr := &EtcdRegister{
		etcdCli: client,
		ctx:     ctx,
		cancel:  cancelFunc,
	}
	return svr, nil
}

// 注册服务
// expire 过期时间
func (svr *EtcdRegister) RegisterServer(serviceName, addr string, expire int64) (err error) {

	// 创建租约
	err = svr.CreateLease(expire)
	if err != nil {
		return err
	}

	// 绑定租约
	err = svr.BindLease(serviceName, addr)
	if err != nil {
		return err
	}

	// 续租
	keepAliveChan, err := svr.KeepAlive()
	if err != nil {
		return err
	}

	// 监听续约
	go svr.Watcher(serviceName, keepAliveChan)

	return nil
}

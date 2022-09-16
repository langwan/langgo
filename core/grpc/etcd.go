package rpc

import (
	"context"
	"fmt"
	"github.com/langwan/langgo/core"
	"github.com/langwan/langgo/core/log"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
)

var etcdClient *clientv3.Client

func EtcdRegister(etcdHost, serviceName, addr string, ttl int64) error {

	etcdClient, err := clientv3.NewFromURL(etcdHost)

	if err != nil {
		return err
	}

	em, err := endpoints.NewManager(etcdClient, serviceName)
	if err != nil {
		return err
	}

	lease, _ := etcdClient.Grant(context.TODO(), ttl)

	err = em.AddEndpoint(context.TODO(), fmt.Sprintf("%s/%s", serviceName, addr), endpoints.Endpoint{Addr: addr}, clientv3.WithLease(lease.ID))
	if err != nil {
		return err
	}

	core.DeferAdd(func() {
		EtcdUnRegister(serviceName, addr)
	})

	alive, err := etcdClient.KeepAlive(context.TODO(), lease.ID)
	if err != nil {
		return err
	}

	go func() {
		for {
			<-alive
			log.Logger("grpc", "etcd").Debug().Msg("Keep Alive")
		}
	}()

	return nil
}

func EtcdUnRegister(serviceName, addr string) error {
	log.Logger("grpc", "etcd").Debug().Str("addr", addr).Msg("unregister")
	if etcdClient != nil {
		em, err := endpoints.NewManager(etcdClient, serviceName)
		if err != nil {
			return err
		}
		err = em.DeleteEndpoint(context.TODO(), fmt.Sprintf("%s/%s", serviceName, addr))
		if err != nil {
			return err
		}
		return err
	}
	return nil
}

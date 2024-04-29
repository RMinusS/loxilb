sudo su
apt-get update && apt-get install ipvsadm ipset -y
export WORKER_ADDR=$(ip a |grep global | grep -v '10.0.2.15' | grep '192.168.80' | awk '{print $2}' | cut -f1 -d '/')
export MASTER_ADDR=$(cat /vagrant/master-ip)
export NODE_TOKEN=$(cat /vagrant/node-token)
mkdir -p /etc/rancher/k3s
cp -f /vagrant/k3s.yaml /etc/rancher/k3s/k3s.yaml
curl -sfL https://get.k3s.io | K3S_TOKEN=${NODE_TOKEN} sh -s - agent --server https://192.168.80.80:6443 --node-ip=${WORKER_ADDR} --node-external-ip=${WORKER_ADDR} -t ${NODE_TOKEN} --flannel-iface=eth1 --kube-proxy-arg proxy-mode=ipvs --disable-apiserver-lb
/vagrant/wait_ready.sh
sysctl net.core.netdev_max_backlog=10000

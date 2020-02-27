vhost_lb_bridge=br-lb-data
vhost_lb_network=default
vhost_nest_dom=cobbler-vm
vhost_base_path=/home/arunt/lab/openstack
vhost_nest_path=$vhost_base_path/$vhost_nest_dom.qcow2
vhost_nest_path_ci=$vhost_base_path/$vhost_nest_dom-ud.qcow2
virt-install \
    --network=network:$vhost_lb_network \
    --name=$vhost_nest_dom \
    --disk path=$vhost_nest_path,format=qcow2,cache=none \
    --disk path=$vhost_nest_path_ci,format=qcow2,cache=none \
    --ram=256 \
    --vcpus=2 \
    --os-type linux \
    --autostart \
    --boot hd \
    --noautoconsole \
    --nographics


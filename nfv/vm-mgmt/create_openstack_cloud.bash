#Use a golden image to create multiple cloud images with 
#appropriate MAC addresses and IPs

#Cloud creation specific

LATEST_BUILD=openstack_latest.tar
VM_IMAGE_DIR=/var/lib/libvirt/images
VM_XML_DIR=/etc/libvirt/qemu
NET0=/etc/sysconfig/network-scripts/ifcfg-eth0
NET1=/etc/sysconfig/network-scripts/ifcfg-eth1
NET2=/etc/sysconfig/network-scripts/ifcfg-eth2
NET3=/etc/sysconfig/network-scripts/ifcfg-eth3

#BASE_VM=ovs1
BASE_VM=fed20_base
BASE_XML=$BASE_VM.xml
BASE_IMG=$VM_IMAGE_DIR/$BASE_VM.qcow2
BASE_MOUNT=/mnt/base
TARGET_MOUNT=/mnt/target

H_CLOUD_VM=h_cloud
H_CLOUD_HOST=control
H_CLOUD_XML=$H_CLOUD_VM.xml
H_CLOUD_MAC=52:54:00:0e:ae:e4
H_CLOUD_MAC_1=52:54:00:5c:81:5b
H_CLOUD_MAC_2=52:54:00:a6:3c:43
H_CLOUD_MAC_3=52:54:00:00:41:4c
H_CLOUD_MOUNT=/mnt/ostack/h_cloud
H_CLOUD_IMG=$VM_IMAGE_DIR/$H_CLOUD_VM.qcow2
H_CLOUD_IMG_1=$VM_IMAGE_DIR/$H_CLOUD_VM-1.qcow2

H_NETWORK_VM=h_network
H_NETWORK_HOST=network
H_NETWORK_XML=$H_NETWORK_VM.xml
H_NETWORK_MAC=52:54:00:de:ce:d1
H_NETWORK_MAC_1=52:54:00:eb:72:6b
H_NETWORK_MAC_2=52:54:00:4d:df:a1
H_NETWORK_MAC_3=52:54:00:22:a0:c3
H_NETWORK_MOUNT=/mnt/ostack/h_network
H_NETWORK_IMG=$VM_IMAGE_DIR/$H_NETWORK_VM.qcow2
H_NETWORK_IMG_1=$VM_IMAGE_DIR/$H_NETWORK_VM-1.qcow2

H_COMPUTE_VM=h_compute
H_COMPUTE_HOST=compute1
H_COMPUTE_XML=$H_COMPUTE_VM.xml
H_COMPUTE_MAC=52:54:00:2f:0d:20
H_COMPUTE_MAC_1=52:54:00:2c:00:5e
H_COMPUTE_MAC_2=52:54:00:90:9b:7b
H_COMPUTE_MAC_3=52:54:00:27:9a:29
H_COMPUTE_MOUNT=/mnt/ostack/h_compute
H_COMPUTE_IMG=$VM_IMAGE_DIR/$H_COMPUTE_VM.qcow2
H_COMPUTE_IMG_1=$VM_IMAGE_DIR/$H_COMPUTE_VM-1.qcow2

H_COMPUTE2_VM=h_compute2
H_COMPUTE2_HOST=compute2
H_COMPUTE2_XML=$H_COMPUTE2_VM.xml
H_COMPUTE2_MAC=52:54:00:72:7b:7a
H_COMPUTE2_MAC_1=52:54:00:84:89:1e
H_COMPUTE2_MAC_2=52:54:00:29:ff:66
H_COMPUTE2_MAC_3=52:54:00:6e:19:3f
H_COMPUTE2_MOUNT=/mnt/ostack/g_cloud
H_COMPUTE2_IMG=$VM_IMAGE_DIR/$H_COMPUTE2_VM.qcow2
H_COMPUTE2_IMG_1=$VM_IMAGE_DIR/$H_COMPUTE2_VM-1.qcow2


I_CLOUD_VM=i_cloud
I_CLOUD_HOST=block1
I_CLOUD_XML=$I_CLOUD_VM.xml
I_CLOUD_MAC=52:54:00:b6:c5:d2
I_CLOUD_MAC_1=52:54:00:04:28:39
I_CLOUD_MAC_2=52:54:00:63:a3:d0
I_CLOUD_MAC_3=52:54:00:58:0f:25
I_CLOUD_MOUNT=/mnt/ostack/i_cloud
I_CLOUD_IMG=$VM_IMAGE_DIR/$I_CLOUD_VM.qcow2
I_CLOUD_IMG_1=$VM_IMAGE_DIR/$I_CLOUD_VM-1.qcow2

I_NETWORK_VM=i_network
I_NETWORK_HOST=block2
I_NETWORK_XML=$I_NETWORK_VM.xml
I_NETWORK_MAC=52:54:00:7b:79:e1
I_NETWORK_MAC_1=52:54:00:4e:ad:2e
I_NETWORK_MAC_2=52:54:00:b1:5f:f1
I_NETWORK_MAC_3=52:54:00:67:ad:b3
I_NETWORK_MOUNT=/mnt/ostack/i_network
I_NETWORK_IMG=$VM_IMAGE_DIR/$I_NETWORK_VM.qcow2
I_NETWORK_IMG_1=$VM_IMAGE_DIR/$I_NETWORK_VM-1.qcow2

I_COMPUTE_VM=i_compute
I_COMPUTE_HOST=object1
I_COMPUTE_XML=$I_COMPUTE_VM.xml
I_COMPUTE_MAC=52:54:00:67:10:fe
I_COMPUTE_MAC_1=52:54:00:ea:bf:ef
I_COMPUTE_MAC_2=52:54:00:05:45:f4
I_COMPUTE_MAC_3=52:54:00:f0:f5:a1
I_COMPUTE_MOUNT=/mnt/ostack/i_compute
I_COMPUTE_IMG=$VM_IMAGE_DIR/$I_COMPUTE_VM.qcow2
I_COMPUTE_IMG_1=$VM_IMAGE_DIR/$I_COMPUTE_VM-1.qcow2

I_COMPUTE2_VM=i_compute2
I_COMPUTE2_HOST=object2
I_COMPUTE2_XML=$I_COMPUTE2_VM.xml
I_COMPUTE2_MAC=52:54:00:94:ea:7f
I_COMPUTE2_MAC_1=52:54:00:ae:a7:a0
I_COMPUTE2_MAC_2=52:54:00:76:b2:f6
I_COMPUTE2_MAC_3=52:54:00:a6:af:a9
I_COMPUTE2_MOUNT=/mnt/ostack/g_network
I_COMPUTE2_IMG=$VM_IMAGE_DIR/$I_COMPUTE2_VM.qcow2
I_COMPUTE2_IMG_1=$VM_IMAGE_DIR/$I_COMPUTE2_VM-1.qcow2



J_CLOUD_VM=j_cloud
J_CLOUD_HOST=holmes
J_CLOUD_XML=$I_CLOUD_VM.xml
J_CLOUD_MAC=52:54:00:c3:13:df
J_CLOUD_MAC_1=52:54:00:30:4e:7a
J_CLOUD_MAC_2=52:54:00:6b:19:92
J_CLOUD_MAC_3=52:54:00:87:c1:bd
J_CLOUD_MOUNT=/mnt/ostack/i_cloud
J_CLOUD_IMG=$VM_IMAGE_DIR/$J_CLOUD_VM.qcow2
J_CLOUD_IMG_1=$VM_IMAGE_DIR/$J_CLOUD_VM-1.qcow2

J_NETWORK_VM=j_network
J_NETWORK_HOST=watson
J_NETWORK_XML=$J_NETWORK_VM.xml
J_NETWORK_MAC=52:54:00:79:a2:b4
J_NETWORK_MAC_1=52:54:00:ce:cd:e1
J_NETWORK_MAC_2=52:54:00:a6:29:b2
J_NETWORK_MAC_3=52:54:00:95:11:8c
J_NETWORK_MOUNT=/mnt/ostack/j_network
J_NETWORK_IMG=$VM_IMAGE_DIR/$J_NETWORK_VM.qcow2
J_NETWORK_IMG_1=$VM_IMAGE_DIR/$J_NETWORK_VM-1.qcow2

J_COMPUTE_VM=j_compute
J_COMPUTE_HOST=adler
J_COMPUTE_XML=$J_COMPUTE_VM.xml
J_COMPUTE_MAC=52:54:00:59:bf:1c
J_COMPUTE_MAC_1=52:54:00:d9:b1:7e
J_COMPUTE_MAC_2=52:54:00:83:cb:c5
J_COMPUTE_MAC_3=52:54:00:27:4b:ca
J_COMPUTE_MOUNT=/mnt/ostack/j_compute
J_COMPUTE_IMG=$VM_IMAGE_DIR/$J_COMPUTE_VM.qcow2
J_COMPUTE_IMG_1=$VM_IMAGE_DIR/$J_COMPUTE_VM-1.qcow2

J_COMPUTE2_VM=j_compute2
J_COMPUTE2_HOST=hudson
J_COMPUTE2_XML=$J_COMPUTE2_VM.xml
J_COMPUTE2_MAC=52:54:00:21:ed:a4
J_COMPUTE2_MAC_1=52:54:00:08:31:a1
J_COMPUTE2_MAC_2=52:54:00:22:6c:1a
J_COMPUTE2_MAC_3=52:54:00:2d:83:41
J_COMPUTE2_MOUNT=/mnt/ostack/g_compute
J_COMPUTE2_IMG=$VM_IMAGE_DIR/$J_COMPUTE2_VM.qcow2
J_COMPUTE2_IMG_1=$VM_IMAGE_DIR/$J_COMPUTE2_VM-1.qcow2




DEV_VM=clouddev
DEV_XML=$DEV_VM.xml
DEV_MAC=52:54:00:91:99:74
DEV_MAC_1=52:54:00:6b:97:63
DEV_MAC_2=52:54:00:c9:63:7b
DEV_MAC_3=52:54:00:5c:a0:1b
DEV_MOUNT=/mnt/ostack/clouddev
DEV_IMG=$VM_IMAGE_DIR/$DEV_VM.qcow2
DEV_IMG_1=$VM_IMAGE_DIR/$DEV_VM-1.qcow2

OS_DEBUG_PAUSE=1

#Two (2) files for images - one for OS (vda), one for extra storage (vdb)
#Four (4) NICs for network - eth0 cloudnet eth1 extnet eth2 default eth4 macvtap

debug_openstack()
{

#arg1 : string to print
echo "###DEBUG### "
echo $1
echo "###DEBUG###"
if [ $OS_DEBUG_PAUSE -eq 1 ]
then
	echo "####DEBUG PAUSED### "
	echo "Exit to continue"
	echo "####DEBUG PAUSED### "
bash
fi




}

vm_change_mac_one()
{
#To clone one VM to another 
#arg1 - name of source VM
#arg2 - name of target VM
#arg3 - new MAC for eth0
#arg4 - new MAC for eth1
#arg5 - new MAC for eth2
#arg6 - new MAC for eth3
#arg7 - hostname	
echo "changing vm mac"
guestmount -a $VM_IMAGE_DIR/$1.qcow2 -i $BASE_MOUNT
guestmount -a $VM_IMAGE_DIR/$2.qcow2 -i $TARGET_MOUNT

echo "HWADDR=$3" >> $TARGET_MOUNT/$NET0
echo "HWADDR=$4" >> $TARGET_MOUNT/$NET1
echo "HWADDR=$5" >> $TARGET_MOUNT/$NET2
echo "HWADDR=$6" >> $TARGET_MOUNT/$NET3

echo $7 > $TARGET_MOUNT/etc/hostname

#echo "verify information ifcfg mac"

#echo "Trying to check mount values; Exit to umount and continue"
#echo "Check the ifconfig files to ensure IP address values are accurate"
#echo "Check for hostname"
#grep HWADDR $TARGET_MOUNT/$NET0
#grep HWADDR $TARGET_MOUNT/$NET1
#grep HWADDR $TARGET_MOUNT/$NET2
#grep HWADDR $TARGET_MOUNT/$NET3
#bash


umount $BASE_MOUNT
umount $TARGET_MOUNT


}


vm_change_mac_one_sed()
{
#To clone one VM to another 
#arg1 - name of source VM
#arg2 - name of target VM
#arg3 - new MAC for eth0
#arg4 - new MAC for eth1
#arg5 - new MAC for eth2
#arg6 - new MAC for eth3
#arg7 - hostname	
echo "changing vm mac"
guestmount -a $VM_IMAGE_DIR/$1.qcow2 -i $BASE_MOUNT
guestmount -a $VM_IMAGE_DIR/$2.qcow2 -i $TARGET_MOUNT

base_hw_addr=`grep HWADDR $BASE_MOUNT/$NET0`
sed -i  "s/$base_hw_addr/HWADDR=$3/g" $TARGET_MOUNT/$NET0
base_hw_addr=`grep HWADDR $BASE_MOUNT/$NET1`
sed -i  "s/$base_hw_addr/HWADDR=$4/g" $TARGET_MOUNT/$NET1
base_hw_addr=`grep HWADDR $BASE_MOUNT/$NET2`
sed -i  "s/$base_hw_addr/HWADDR=$5/g" $TARGET_MOUNT/$NET2
base_hw_addr=`grep HWADDR $BASE_MOUNT/$NET3`
sed -i  "s/$base_hw_addr/HWADDR=$6/g" $TARGET_MOUNT/$NET3

echo $7 > $TARGET_MOUNT/etc/hostname

#echo "Trying to check mount values; Exit to umount and continue"
#echo "Check the ifconfig files to ensure IP address values are accurate"
#echo "Check for hostname"
#grep HWADDR $TARGET_MOUNT/$NET0
#grep HWADDR $TARGET_MOUNT/$NET1
#grep HWADDR $TARGET_MOUNT/$NET2
#grep HWADDR $TARGET_MOUNT/$NET3
#bash


umount $BASE_MOUNT
umount $TARGET_MOUNT


}

vm_clone_one()
{

#To clone one VM to another 
#arg1 - name of source VM
#arg2 - name of target VM
#arg3 - new MAC for eth0
#arg4 - new MAC for eth1
#arg5 - new MAC for eth2
#arg6 - new MAC for eth3
virsh destroy $1
virt-clone --original $1 --name $2 --file $VM_IMAGE_DIR/$2.qcow2 --file $VM_IMAGE_DIR/$2-1.qcow2 --mac $3 --mac $4 --mac $5 --mac $6

}

vm_startup_one()
{
#1 - name of base VM
#2 - name of target vm
#3 - Mac 0
#4 - Mac 1
#5 - Mac 2
#6 - Mac 3
#7 - hostname
#8 - number of disks (two for cinder hosts, one for everyone else)
virsh destroy $2
virsh undefine $2 --remove-all-storage 
#virsh destroy $1
virt-clone --original $1 --name $2 --file $VM_IMAGE_DIR/$2.qcow2 --file $VM_IMAGE_DIR/$2-1.qcow2 --mac $3 --mac $4 --mac $5 --mac $6
virsh start $2
#sleep 30
#virsh destroy $2
vm_change_mac_one $1 $2 $3 $4 $5 $6 $7
virsh destroy $2
sleep 30
virsh start $2
}

vm_startup_one_no_mac()
{
#1 - name of base VM
#2 - name of target vm
virsh destroy $2
virsh undefine $2 --remove-all-storage 
virsh destroy $1
virt-clone --original $1 --name $2 --file $VM_IMAGE_DIR/$2.qcow2 --file $VM_IMAGE_DIR/$2-1.qcow2 
#vm_change_mac_one $1 $2 $3 $4 $5 $6 $7
virsh start $2

}

vm_startup_havana()
{
#echo "nothing here"
#havana cloud is used for control and compute platforms
vm_startup_one $BASE_VM $H_CLOUD_VM $H_CLOUD_MAC $H_CLOUD_MAC_1 $H_CLOUD_MAC_2 $H_CLOUD_MAC_3 $H_CLOUD_HOST
sleep 2
vm_startup_one $BASE_VM $H_NETWORK_VM $H_NETWORK_MAC $H_NETWORK_MAC_1 $H_NETWORK_MAC_2 $H_NETWORK_MAC_3 $H_NETWORK_HOST
sleep 2
vm_startup_one $BASE_VM $H_COMPUTE_VM $H_COMPUTE_MAC $H_COMPUTE_MAC_1 $H_COMPUTE_MAC_2 $H_COMPUTE_MAC_3 $H_COMPUTE_HOST
sleep 2
vm_startup_one $BASE_VM $H_COMPUTE2_VM $H_COMPUTE2_MAC $H_COMPUTE2_MAC_1 $H_COMPUTE2_MAC_2 $H_COMPUTE2_MAC_3 $H_COMPUTE2_HOST 
sleep 2

}

vm_startup_icehouse()
{
#icehouse cloud is exclusively for storage
vm_startup_one $BASE_VM $I_CLOUD_VM $I_CLOUD_MAC $I_CLOUD_MAC_1 $I_CLOUD_MAC_2 $I_CLOUD_MAC_3 $I_CLOUD_HOST
sleep 2
vm_startup_one $BASE_VM $I_NETWORK_VM $I_NETWORK_MAC $I_NETWORK_MAC_1 $I_NETWORK_MAC_2 $I_NETWORK_MAC_3 $I_NETWORK_HOST
sleep 2
vm_startup_one $BASE_VM $I_COMPUTE_VM $I_COMPUTE_MAC $I_COMPUTE_MAC_1 $I_COMPUTE_MAC_2 $I_COMPUTE_MAC_3 $I_COMPUTE_HOST
sleep 2
vm_startup_one $BASE_VM $I_COMPUTE2_VM $I_COMPUTE2_MAC $I_COMPUTE2_MAC_1 $I_COMPUTE2_MAC_2 $I_COMPUTE2_MAC_3 $I_COMPUTE2_HOST 
sleep 2
}

vm_startup_juno()
{
vm_startup_one $BASE_VM $J_CLOUD_VM $J_CLOUD_MAC $J_CLOUD_MAC_1 $J_CLOUD_MAC_2 $J_CLOUD_MAC_3 $J_CLOUD_HOST
vm_startup_one $BASE_VM $J_NETWORK_VM $J_NETWORK_MAC $J_NETWORK_MAC_1 $J_NETWORK_MAC_2 $J_NETWORK_MAC_3 $J_NETWORK_HOST
vm_startup_one $BASE_VM $J_COMPUTE_VM $J_COMPUTE_MAC $J_COMPUTE_MAC_1 $J_COMPUTE_MAC_2 $J_COMPUTE_MAC_3 $J_COMPUTE_HOST
vm_startup_one $BASE_VM $J_COMPUTE2_VM $J_COMPUTE2_MAC $J_COMPUTE2_MAC_1 $J_COMPUTE2_MAC_2 $J_COMPUTE2_MAC_3 $J_COMPUTE2_HOST
}

virsh destroy $BASE_VM
virt-sysprep -d $BASE_VM --enable udev-persistent-net,machine-id
vm_startup_havana
vm_startup_icehouse
#vm_startup_juno

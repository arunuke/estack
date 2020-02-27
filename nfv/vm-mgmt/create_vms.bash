VM_IMAGE_DIR=/home/arunt/basevms
BASE_MOUNT=/mnt/guestmnt
OS_DEBUG_PAUSE=0

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

filesystem_vm_clone_one()
{

#To clone one VM to another
#arg1 - name of source VM
#arg2 - name of target VM
#arg3 - new MAC for eth0
#arg4 - new MAC for eth1
#arg5 - hostname

virsh shutdown $1
virt-sysprep -d $1 --enable udev-persistent-net,machine-id
virsh destroy $2

virsh undefine $2 --remove-all-storage

virt-clone --original $1 --name $2 --file $VM_IMAGE_DIR/$2.qcow2 --mac $3 --mac $4
virt-sysprep --hostname $5 --enable hostname -d $2

virsh start $2


}

#filesystem_vm_clone_one filesysbase ctrl1 52:54:00:af:23:58 52:54:00:f8:0e:2b ctrl1
#filesystem_vm_clone_one filesysbase ctrl2 52:54:00:0b:06:0c 52:54:00:ad:9f:4e ctrl2
#filesystem_vm_clone_one filesysbase ctrl3 52:54:00:6d:f9:d1 52:54:00:5d:80:34 ctrl3
#filesystem_vm_clone_one filesysbase ctrl4 52:54:00:24:ff:81 52:54:00:6a:64:b9 ctrl4

#filesystem_vm_clone_one filesysbase vsa1 52:54:00:83:fd:05 52:54:00:b5:71:90 vsa1
#filesystem_vm_clone_one filesysbase vsa2 52:54:00:e1:1c:2e 52:54:00:8a:2d:8b vsa2
#filesystem_vm_clone_one filesysbase vsa3 52:54:00:af:ed:da 52:54:00:1b:67:de vsa3


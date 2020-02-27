#For a typical VM, we will provide 5 interfaces
#mgmt : all management networks
#data : underlay for openstack networks
#external : external network
#default : default that uses the dhcp provided by linux bridge
#direct : a macvtap interface plugged in, just in case 

#Need a new UUID
echo "*********************"
echo "$1 MAC addresses to generate"
echo "*********************"
for i in `seq 1 $1`
do
od -An -N6 -tx1 /dev/urandom | sed -e 's/^  *//' -e 's/  */:/g' -e 's/:$//' -e 's/^\(.\)[13579bdf]/\10/'
done
echo "*********************"
uuidgen -r
echo "*********************"

#common tasks for all services

#1. name of service - glance, nova
#2. name of domain - default
#3. service user password - nfv
#4. service description - "Image Service", "Compute Service"
#5. service name - image, compute
#6. region name  - RegionOne
#7. controller name - controller
#8. service port - 9292

#ingest admin credentials
#create a user for the service
#add the user as admin role 
#create the service
#create endpoints for public, internal and admin
  
. ./adminrc

openstack user create --domain $2 --password $3 $1

openstack role add --project service --user $1 admin

openstack service create --name $1 --description $4 $5

openstack endpoint create --region $6 $5 public http://$7:$8 
openstack endpoint create --region $6 $5 internal http://$7:$8 
openstack endpoint create --region $6 $5 admin http://$7:$8 

#call db_sync based on service being configured, captured in $1
if [ "$1" == "keystone" ]
then
  keystone-manage db_sync
elif [ "$1" == "glance" ]
then
  glance-manage db_sync
elif [ "$1" == "nova" ]
then
  nova-manage api_db sync
#  nova-manage cell_v2 map_cell0
#  nova-manage cell_v2 create_cell --name=cell1 --verbose
  nova-manage db sync
else
  echo "unknown service"
fi

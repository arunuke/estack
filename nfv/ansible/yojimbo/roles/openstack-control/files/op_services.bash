#script to perform tasks that do not have ansible modules
#arg1 : service (ks_bootstrap)
#arg2 : admin pass (for keystone bootstrap)
#arg3 : hostname of controller
#arg4 : region name

export OS_USERNAME=admin
export OS_PASSWORD=nfv
export OS_PROJECT_NAME=admin
export OS_USER_DOMAIN_NAME=Default
export OS_PROJECT_DOMAIN_NAME=Default
export OS_AUTH_URL=http://controller:35357/v3
export OS_IDENTITY_API_VERSION=3
#populate keystone database
keystone-manage db_sync


#Initialize Fernet key repos
keystone-manage fernet_setup --keystone-user keystone --keystone-group keystone
keystone-manage credential_setup --keystone-user keystone --keystone-group keystone

#keystone-manage bootstrap --bootstrap-password nfv --bootstrap-admin-url http://controller:35357/v3/ --bootstrap-internal-url http://controller:5000/v3/ --bootstrap-public-url http://controller:5000/v3/ --bootstrap-region-id RegionOne

keystone-manage bootstrap --bootstrap-password $2 \
--bootstrap-admin-url http://$3:35357/v3/ \
--bootstrap-internal-url http://$3:5000/v3/ \
--bootstrap-public-url http://$3:5000/v3/ \
--bootstrap-region-id $4

openstack project create --domain default --description "Service Project" service

openstack project create --domain default --description "Demo Project" demo

openstack user create --domain default --password $2 demo

openstack role create user

openstack role add --project demo --user demo user

export OS_USERNAME=admin
export OS_PASSWORD=nfv
export OS_PROJECT_NAME=admin
export OS_USER_DOMAIN_NAME=Default
export OS_PROJECT_DOMAIN_NAME=Default
export OS_AUTH_URL=http://controller:35357/v3
export OS_IDENTITY_API_VERSION=3
VAR_HOST=controller
VAR_REGION=RegionOne
VAR_BSTRAP_PWD=nfv


keystone-manage bootstrap --bootstrap-password $VAR_BSTRAP_PWD \
--bootstrap-admin-url http://$VAR_HOST:35357/v3/ \
--bootstrap-internal-url http://$VAR_HOST:5000/v3/ \
--bootstrap-public-url http://$VAR_HOST:5000/v3/ \
--bootstrap-region-id $VAR_REGION


IPADDR=10.8.0.20
for i in 1 2 3 4
do 
  curl https://$IPADDR/redfish/v1/Systems/1/BaseNetworkAdapters/$i/ --insecure -u admin:HP1nvent -L | json_pp > pci_base_file_$i.json
done

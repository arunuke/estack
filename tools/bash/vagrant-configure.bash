#configure vagrant vms with ubuntu user (to avoid ansible become user)
#copy authorized_keys file

useradd -G vagrant ubuntu


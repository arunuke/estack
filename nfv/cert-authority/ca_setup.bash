# !!!!Do not run as sudo. local keys get created as root and that's messy !!! #
CA_ID=nfv-system-ca
CA_USER_ID=nfv-system-userkey-ca
CA_HOST=ca-host-key
SRV2_HOST=srv2-host-key
HOST_RSA_PATH=/etc/ssh
HOST_RSA_KEY=ssh_host_rsa_key
HOST_RSA_FILE=$HOST_RSA_PATH/$HOST_RSA_KEY
SRV1=ctrl1.hpe.com
SRV2=ctrl2
SRV3=ctrl3
USER=arunt
NFV_USER=arunt

SSH_CONFIG_FILE=sshd_config
SSH_CONFIG_DIR=/etc/ssh
SSH_CONFIG_PATH=$SSH_CONFIG_DIR/$SSH_CONFIG_FILE

ca_setup()
{

#setup host signing keys
#uses empty pass phrase, change it to actual pass phrase 
if [[ -f $CA_ID ]]
then
	echo "$CA_ID key exists"
else
	ssh-keygen -b 4096 -t rsa -f $CA_ID -C "NFV System CA" -N ""
fi

#sign host keys on ca
#Always sign a public key (host_rsa_key.pub) using the private key of
#the host
ssh-keygen -s $CA_ID -I $CA_HOST -h -n $1 -V +52w  $HOST_RSA_FILE.pub -C "signed host key for CA"
ls -lart $HOST_RSA_FILE*
cat $HOST_RSA_FILE-cert.pub

#setup sshd to use the right host certificate
#Check if file has an entry for Host Certificate with the file path. If not, go ahead add
grep "HostCertificate $HOST_RSA_FILE-cert.pub" $SSH_CONFIG_PATH
if [ $? -eq 1 ]
then
	echo  "HostCertificate /etc/ssh/ssh_host_rsa_key-cert.pub" >>  $SSH_CONFIG_PATH
	echo "ssh config information changed"
else
	echo "ssh config information present already, hence unchanged"

fi


service ssh restart
}

ca_user_key_setup()
{


#setup user signing keys
#uses empty pass phrase, change it to actual pass phrase 
ssh-keygen -b 4096 -t rsa -f $CA_USER_ID -C "NFV System CA User Key" -N ""

#setup sshd to use the right user certificate
#Check if file has an entry for Trusted CA Key with the file path. If not, go ahead add
grep "TrustedUserCAKeys /etc/ssh/$CA_USER_ID.pub" $SSH_CONFIG_PATH
if [ $? -eq 1 ]
then
	echo  "TrustedUserCAKeys /etc/ssh/$CA_USER_ID.pub" >>  $SSH_CONFIG_PATH
	echo "ssh config information changed"
else
	echo "ssh config information present already, hence unchanged"

fi

#copy public key of user signing key to all server nodes through server_user_setup()

#setup all ssh clients to use this certificate


}

server_user_setup()
{
#arg1 is host

#copy the user signing key to /etc/ssh on the ssh server node
#copy it to the home directory, then move it to up /etc/ssh 
scp $CA_USER_ID.pub $USER@$1:~/
scp $SSH_CONFIG_PATH $USER@$1:~/
ssh $USER@$1 "sudo cp ~/$CA_USER_ID.pub $SSH_CONFIG_DIR/"
ssh $USER@$1 "sudo cp ~/$SSH_CONFIG_FILE $SSH_CONFIG_DIR/"
ssh $USER@$1 "sudo service ssh restart"

}

client_user_setup()
{
#arg1 is host
#generate keys if not available already
ssh $1 "if [[ -f ~/.ssh/$2 ]]; then echo "found"; else ssh-keygen -f ~/.ssh/id_rsa -t rsa -N \"\" ; fi"

#pickup user key of remote host
scp $USER@$1:~/.ssh/id_rsa.pub .

#sign user's ssh key with signing key (user key) of CA
ssh-keygen -s $CA_USER_ID -I user_$NFV_USER -n $NFV_USER -V +52w id_rsa.pub

scp id_rsa-cert.pub $USER@$1:~/.ssh/

}

server_host_setup()
{
#arg1 is host 

#copy the host key, sign it with signing key
#copy the signed host key back to the home directory of user, then copy it to
#ssh directory (since using scp directly doesn't work)
#also modify the ssh config file to include the host informaton

#pickup host key of remote host
scp $USER@$1:$HOST_RSA_FILE.pub .

#sign remote host's host key with signing key
ssh-keygen -s $CA_ID -I $1_HOST -h -n $1 -V +52w  ./$HOST_RSA_KEY.pub
ls -lart ./$HOST_RSA_KEY-cert.pub
cat $HOST_RSA_KEY-cert.pub

#copy signed host key back to home directory of user
	
scp $SSH_CONFIG_PATH $USER@$1:~/
scp $HOST_RSA_KEY-cert.pub $USER@$1:~/

#move the files from the home directory of remote host to the right ssh configuration directories
#restart ssh service

ssh $USER@$1 "sudo cp ~/$HOST_RSA_KEY-cert.pub $SSH_CONFIG_DIR/"
ssh $USER@$1 "sudo cp ~/$SSH_CONFIG_FILE $SSH_CONFIG_DIR/"
ssh $USER@$1 "sudo service ssh restart"

#remove local copy
rm ./$HOST_RSA_KEY.pub ./$HOST_RSA_KEY-cert.pub

#sign user keys on srv2

}

client_host_setup()
{
#TODO : Automate setup of the client's known hosts file
#arg1 is the client which will login to the ssh server using the public signing key
#of the CA included as a known host

echo "client host key setup"
#add an entry to known_hosts file for @cert-authority <all hosts, ie '*', followed
#by the public signing key of the CA



}

cleanup()
{

rm $CA_ID*
rm $HOST_RSA_KEY-cert*
ssh $SRV2 "rm $HOST_RSA_PATH/$HOST_RSA_KEY-cert"

}

#ca_setup ctrl1
#server_host_setup ctrl2
#ca_user_key_setup
#server_user_setup ctrl2
#client_user_setup ctrl3
server_host_setup ctrl3
#server_user_setup ctrl3
#client_user_setup ctrl2
#cleanup

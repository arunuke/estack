#script courtesy: http://stackoverflow.com/questions/10929453/read-a-file-line-by-line-assigning-the-value-to-a-variable
USER=arunt
SESAME_RSA=/home/$USER/nfv-tools-git/files/build/sesame_rsa

while IFS='' read -r line || [[ -n "$line" ]]; do
    rm $line
    ln -s $SESAME_RSA $line
done < "$1"

set_string="$USER ALL=(ALL) NOPASSWD: /usr/bin/pmset"
sudoers_contents=$(sudo grep "$set_string" /etc/sudoers)
if [[ "$sudoers_contents" == "" ]]
then
  echo $set_string | sudo EDITOR="tee -a" visudo
fi

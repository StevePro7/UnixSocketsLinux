https://www.journaldev.com/34113/opening-a-port-on-linux
netstat -lntu

sudo ufw allow 4000

firewall-cmd --add-port=4000/tcp


iptables -A INPUT -p tcp --dport 4000 -j ACCEPT

sudo iptables -v --numeric --table nat --list

sudo service iptables restart
Failed to restart iptables.service: Unit iptables.service not found.

ls | nc -l -p 4000

telnet localhost 4000
nmap localhost -p 4000


OR
sudo iptables-save | sudo tee -a /etc/iptables.conf

sudo iptables-restore < /etc/iptables.conf



https://www.cyberciti.biz/faq/ubuntu-start-stop-iptables-service

sudo ufw status

sudo ufw enable


sudo iptables -L -n -v

sudo iptables-restore

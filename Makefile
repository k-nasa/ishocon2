prebench:
	sudo cp /dev/null /var/log/nginx/access.log
	sudo systemctl restart nginx
	sudo systemctl restart mariadb

nginx:
	sudo cp -f ~/nginx.conf /etc/nginx.conf
	sudo systemctl restart nginx
mysql:
	sudo cp ~/my.cnf /etc/mysql/my.cnf
	sudo systemctl restart mariadb

alp:
	alp -f /var/log/nginx/access.log --aggregates "/political_parties/, /candidates/" --avg --excludes "/initialize"

pt:
	sudo pt-query-digest /var/log/mysql/slow.log

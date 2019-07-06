prebench:
	sudo cp /dev/null /var/log/nginx/access.log
	sudo systemctl restart nginx
	sudo systemctl restart mysql

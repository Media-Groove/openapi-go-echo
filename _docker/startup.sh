# supervisor起動
cd /usr/src || exit
supervisord -n -c /etc/supervisor/supervisord.conf

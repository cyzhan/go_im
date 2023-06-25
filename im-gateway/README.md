# im-gateway
```.env=
#golang字串保留字元，參數不使用字元: %
#---------------------
# RabbitMQ       
#---------------------
AMQP_DIAL='amqp://developer:1234qwer@172.28.40.70:5672/'

#---------------------
# MariaDB        
#---------------------
MARIA='root:qwer@tcp(172.28.41.222:3306)/rabbit?charset=utf8'
MARIA_URL=172.28.41.222:3306
MARIA_ACCOUNT=root
MARIA_PW=qwer
MARIA_DATABASE=im
MARIA_LOG_MOD=true

#---------------------
# Redis        
#---------------------
REDIS_URL='inno.jim.com:6379'
REDIS_PW=''
REDIS_POOL_SIZE=4 # 默認為4倍CPU數 
REDIS_MIN_IDLE=2

#---------------------
# MINIO       
#---------------------
MINIO_ENDPOINT='im.innodev.site:9000'
MINIO_ACCESS_KEY_ID='mzhzKHQmACPeim5x'
MINIO_SECRET_ACCESS_KEY='tUwZvvJgDTPPd6ncWDgv7nmd59VaQKmy'
MINIO_SECURE='false'

#---------------------
# Application       
#---------------------
SERVER_PORT=':8080'
LOGIC_SERVER='localhost:8070'
SALT1='fepwhgZeiTVpeugDkYc63T' # used for password encrypt
JWT_SECRET='pwojgpfj4i2@xSfjI4387SlHJ'
ALLOW_ORIGINS='http://172.28.40.157,http://localhost,http://im.innodev.site'
SNOWFLAKE_NODE=33

```



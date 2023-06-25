# .env example
```.env=
#---------------------
# RabbitMQ       
#---------------------
AMQP_DIAL='amqp://developer:1234qwer@172.28.40.70:5672/'

#---------------------
# MariaDB        
#---------------------
MARIA='root:qwer@tcp(172.28.41.222:3306)/im?charset=utf8'


#---------------------
# Redis        
#---------------------
REDIS_URL='inno.jim.com:6379'
REDIS_PW=''
REDIS_POOL_SIZE=4  # 默認為4倍CPU數 
REDIS_MIN_IDLE=2

#---------------------
# Elastic Search        
#---------------------
ELASTIC_ADDRESS='https://172.28.40.70:9200'
ELASTIC_USER='elastic'
ELASTIC_PASSWORD='1234qwer'

#---------------------
# Application       
#---------------------
AVG_MSG_ID_DIFF=2
CHAT_MSG_CACHE_KEEP=5
MESSAGE_LENGTH_MAX=100
SERVER_PORT=':8091'
# snowflake不能重複: 0-31
SNOWFLAKE_NODE=1
```
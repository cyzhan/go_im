# .env example
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
MARIA_URL='172.28.41.222:3306'
MARIA_ACCOUNT='root'
MARIA_PW='qwer'
MARIA_DATABASE='im'
MARIA_LOG_MOD='true'


#---------------------
# Redis
#---------------------
REDIS_URL='inno.jim.com:6379'
REDIS_PW=''
REDIS_POOL_SIZE=4 # 默認為4倍CPU數
REDIS_MIN_IDLE=2

#---------------------
# Application
#---------------------
SERVER_PORT=':8080'
SALT1='fepwhgZeiTVpeugDkYc63T' # used for password encrypt
JWT_SECRET='pwojgpfj4i2@xSfjI4387SlHJ'
WS_URL='ws://localhost:8090/im/ws/chat?token='
SNOWFLAKE_NODE='56'
CACHE_DEFAULT_EXPIRATION_SEC='300'
SERVER_PORT=

#---------------------
# Grpc
#---------------------
RPC_HOST=localhost
RPC_PORT=8070

#---------------------
# Elastic Search
#---------------------
ELASTIC_ADDRESS='https://172.28.40.70:9200'
ELASTIC_USER='elastic'
ELASTIC_PASSWORD='1234qwer'

```
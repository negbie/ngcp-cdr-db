# ngcp-cdr-db

Is a simple program which scans a configurable folder for new sipwise cdr's and sends them as a batch to timescaledb.

## Init Superset
sudo docker-compose up  
sudo docker exec -it superset superset-init  
browse to localhost:8088  


# ngcp-cdr-db
Is a simple program which scans a configurable folder for new sipwise cdr's and sends them as a batch to timescaledb.

## Init Container
cd docker  
sudo docker-compose up -d  
sudo docker exec -it superset superset-init  


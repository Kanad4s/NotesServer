# PostgreSQL
## Get PostgreSQL

### Download postgres
```bash
git clone git@github.com:postgres/postgres.git --depth=1 --branch=REL_17_STABLE
```
### Set dependencies
```bash
sudo apt install -y \
	gcc \
	bison \
	flex \
	libreadline-dev \
	make \
	libperl-dev \
	libipc-run-perl \
	libicu-dev 
```

### Configure
Go to "postgres"
```bash
./configure --enable-debug --enable-cassert --prefix="$PWD/install" --enable-tap-tests
```
### Install
```bash
make -j $(nproc) world-bin
make -j $(nproc) install-world-bin
```
### Launch
add ```postgres/install/bin``` to PATH  
```vim ~/.bashrc```  
add ```export PATH=$PATH:/your/pass/to/postgres/install/bin```  
close vim  

## Use PostgreSQL
### Create db
```bash
initdb -D data1
```
### Start/Stop db
```bash
pg_ctl start -D data1 -l data1/postgres.log 
pg_ctl stop  -D data1 -l data1/postgres.log 
```
if something went wrong try:
```bash
pg_ctl stop  -D data1 -l data1/postgres.log -m immediate
```
### Connect
```bash
psql --host=localhost --dbname=postgres
```
Originally postgres starts on 5432 port, each next db starts on 5433, 5434,...

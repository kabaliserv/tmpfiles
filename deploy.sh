systemctl stop tmpfile
cd ./client
npm install
npm run build
cd ..
go build -o /usr/bin/tmpfile
systemctl start tmpfile
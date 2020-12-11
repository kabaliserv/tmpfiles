systemctl stop tmpfiles
cd ./client
npm install
npm run build
cd ..
rm /usr/bin/tmpfiles
go build -o /usr/bin/tmpfiles
systemctl start tmpfiles
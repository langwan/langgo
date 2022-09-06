echo "Rotating..."
mv ./logs/app.log ./logs/app.log-old
killall -USR2 load
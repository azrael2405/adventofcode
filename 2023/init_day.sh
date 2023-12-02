
cp -r day_X day_$1
sed -i 's/dayX/day'$1'/' day_$1/go.mod
sed -i 's/dayX/day'$1'/' day_$1/main.go
sed -i 's/use (/use (\n\t.\/day_'$1'/' go.work

cp -r day_X day_$1
mv day_$1/dayX day_$1/day$1
mv day_$1/dayX_test.go day_$1/day${1}_test.go
sed -i 's/use (/use (\n\t.\/day_'$1'/' go.work
sed -i 's/dayX/day'$1'/' day_$1/go.mod
sed -i 's/dayX/day'$1'/g' day_$1/main.go
sed -i 's/dayX/day'$1'/g' day_$1/day$1/Part1.go
sed -i 's/dayX/day'$1'/g' day_$1/day$1/Part2.go
sed -i 's/dayX/day'$1'/g' day_$1/day$1/Funcs.go
sed -i 's/dayX/day'$1'/g' day_$1/day${1}_test.go
sed -i 's/day_X/day_'$1'/g' day_$1/day${1}_test.go
git add .
git commit -m "feat: starting day "$1
cd day_$1

cp -r day_X day_$1
sed -i 's/dayX/day'$1'/' day_$1/go.mod
sed -i 's/use (/use (\n\t.\/day_'$1'/' go.work
git add .
git commit -m "feat: chore staring day "$1
git push
cd day_$1
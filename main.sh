#!/bin/bash

query="$@"

parseDate ()
{
	query=$@
		case $query in
		[0-9]*年[0-9]*月[0-9]*日 )
			echo $query | sed -E 's/([0-9]+)年([0-9]+)月([0-9]+)日/\1-\2-\3/' ;;
		[0-9]*月[0-9]*日 )
			echo $query | sed -E 's/([0-9]+)月([0-9]+)日/\1\/\2/' ;;
		[0-9]*\/[0-9]*\/[0-9]* )
			echo $query | sed -E 's/([0-9]+)\/([0-9]+)\/([0-9]+)/\3-\2-\1/' ;;
		[0-9]*\/[0-9]* )
			echo $query | sed -E 's/([0-9]+)\/([0-9]+)/\2\/\1/' ;;
		[0-9]*\.[0-9]*\.[0-9]* )
			echo $query | sed -E 's/([0-9]+)\.([0-9]+)\.([0-9]+)/\1-\2-\3/' ;;
		[0-9]*\.[0-9]* )
			echo $query | sed -E 's/([0-9]+)\.([0-9]+)/\1\/\2/' ;;
		*)
		  echo $query ;;
	esac
}

if [ -z "$query" ]; then
	# if empty imput, will try to parse string in clipboard. return null if fail
	# will send stderr to /dev/null
	datestring=$(date --date="pbpaste")
else
	datestring=$(parseDate $query)
fi

# `gdate` is same as `date` in linux
# convert day value to chinese day
cday=$(./gdate --date="$datestring" +%u|
xargs -I@ sh -c 'days=(一 二 三 四 五 六 日); echo ${days[@-1]}')

# convert year value to chinese year
cyear=$(./gdate --date="$datestring" +%Y |
	sed -E 's/([0-9])([0-9])([0-9])([0-9])/\1 \2 \3 \4/' |
	xargs -n1 -I@ sh -c 'years=(零 一 二 三 四 五 六 七 八 九); printf ${years[@]}')

# convert month value to chinese month
cmonth=$(./gdate --date="$datestring" +%-m |
	xargs -n1 -I@ sh -c 'months=(零 一 二 三 四 五 六 七 八 九 十 十一 十二); printf ${months[@]}')

# convert date value to chinese date
cdate=$(./gdate --date="$datestring" +%d |
	sed -E 's/([0-9])([0-9])/\1 \2/' |
	xargs sh -c 'tens=("" "十" "二十" "三十"); digits=("" 一 二 三 四 五 六 七 八 九);
	echo ${tens[$0]}${digits[$1]}')

# list of formats for listing
formats=(
	"%Y-%m-%d"
	"%Y-%m-%d_(%a)"
	"%d/%-m/%y_(%a)"
	"%a,_%-d_%b_%y"
	"%A,_%-d_%B_%Y"
	"%-m月%-d日"
	"%Y年%-m月%-d日"
	"%Y年%-m月%-d日（星期${cday}）"
	"${cyear}年${cmonth}月${cdate}日"
	"${cyear}年${cmonth}月${cdate}日（星期${cday}）"
)


# formating result string
items=$(echo \[\"$( echo $echo ${formats[@]} |
	xargs -I@ ./gdate --date="$datestring" +@ |
 # convert list of format string e.g."format_A format_B" to ["format A", "format B"]
	sed -E 's/ /", "/g')\"\] | sed -E 's/_/ /g' |
# using jq as alfred parse script filter in json format
# https://www.alfredapp.com/help/workflows/inputs/script-filter/json/
	./jq '{items: [.[] | .|{title: . ,arg: .}]}')

echo $items

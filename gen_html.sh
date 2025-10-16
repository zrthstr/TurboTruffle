echo "======================================="
echo [+] Running gen_html
echo "======================================="

HL="highlighter.js"

for file in results/*.trufflog;
do
	./gen_html $file $file.html $HL ;
done

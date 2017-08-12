mkdir -p dist-sketchground/www
cp -r css dist-sketchground/www/css
cp -r img dist-sketchground/www/img
cp -r open dist-sketchground/www/open
cp -r fonts dist-sketchground/www/fonts
cp -r index.html dist-sketchground/www/index.html
cp -r contact.html dist-sketchground/www/contact.html
go build -o dist-sketchground/sketchground *.go

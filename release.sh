mkdir -p dist-sketchground
cp -r css dist-sketchground/css
cp -r img dist-sketchground/img
cp -r open dist-sketchground/open
cp -r fonts dist-sketchground/fonts
cp -r index.html dist-sketchground/index.html
go build -o dist-sketchground/sketchground *.go

try: get load upload edit build

load:
	lua main.lua load github.com/ishiikurisu/mcr

upload:
	lua main.lua upload github.com/ishiikurisu/mcr
	lua main.lua upload github.com/ishiikurisu/doodle

upload-no-args:
	lua main.lua load
	lua main.lua upload

edit:
	lua main.lua edit github.com/ishiikurisu/doodle

edit-itself:
	lua main.lua edit -e atom

get:
	lua main.lua get github.com/ishiikurisu/doodle

build:
	lua main.lua build github.com/ishiikurisu/doodle

try-no-args: upload-no-args

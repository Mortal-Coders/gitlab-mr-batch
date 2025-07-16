

build:
	podman build -t gmb:v1 .

run:
	podman run -p 8080:8080 gmb:v1
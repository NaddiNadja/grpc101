# Setup of repository

1. make ``go.mod`` file with:

    ``$ go mod init``

2. make a ``.proto`` file in a subfolder, for example ``time/time.proto``.
3. run command:

    ``$ protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative time/time.proto``

    which should create the two ``pb.go`` files.
4. run command:

    ``$ go mod tidy``

    to install dependencies and create the ``go.sum`` file.

# Email Alias Manager
This is a very simple tool that uses Golang and standard web technologies. It allows me to create and manage email aliases for a catch-all domain. I can create aliases in my desired format (`servicename.<randomhex>@domain`) and I can manage them so I don't forget aliases I have created.

Its a simple tool which means there are some oddities in place to save time. For example, its a single user system. A user can be created on the `/setup` page and then they can log in via the `/login` page. Once someone has setup an account then nobody else can make an account.

Import/Export capabilities coming soon, the API is in place and working just need to design the page.

## Usage
You can use this container via the GitHub container repo, the image link is: `ghcr.io/tywil04/email-alias-manager`

### Ports
By default the container exposes port 4041 and this is the web interface. You can change this port using the `SERVER_ADDRESS` environment variable. For example, to change the port to `12345` then `SERVER_ADDRESS` should be `:12345`. 

### Volumes
By default, to store the data produced by the container you will need to set up a volume or a path on your host to /database.db (its a file, not a folder). You can, however, alter the location with the `SQLITE_DATABASE_PATH` environment variable.

### Environment Variables
- `CRYPTO_PEPPER`: a random string added to all operations, its like a salt but global to the whole database. Not required but heavily recommended.
- `SQLITE_DATABASE_PATH`: the path for the sqlite database path, not required default is /database.db.
- `SERVER_ADDRESS`: the address of the server, by default its `:4041` but can be anything.

### Example
Heres an example `docker run` command:
```
docker run --name email-alias-manager -p 4041:4041 -v /somepath/database.db:/database.db -e CRYPTO_PEPPER="totallyrandom" ghcr.io/tywil04/email-alias-manager
```

## Screenshot
![Demo Screenshot](screenshot.png)

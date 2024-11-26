# Golang Load Balancer

This program is used to run any aplication using golang. But in this version is just for run production file that build on singlepage application such as vue etc.

## How to run?

 ```
go run . -num <total_ports>
 ```

Total ports is num of ports or total service that you want to use

## Install

If you are using linux, you can this command using root:

```
make install
```

Than you can use 

```
goload -num <total_server>
```

Total ports is num of ports or total service that you want to use

## Contributing

I'm open to anyone who wants to contribute to this project! Thanks in advance, and let's build something awesome together!

### Todo
Add more features to the load balancer
- [ ] Can Run Multipage App
- [ ] Make proxy to specify ports that want to use
- [ ] Improve the documentation
- [ ] Add more examples
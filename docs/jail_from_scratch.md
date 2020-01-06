# Jail from scratch

## Resources

- [FreeBSD Handbook: Jails](https://www.freebsd.org/doc/handbook/jails-build.html)

## Steps

As `root`:

### Tools

```sh
cp -v /usr/src/share/examples/jails/jng /usr/sbin/
cp -v /usr/src/share/examples/jails/jib /usr/sbin/
```

### ZFS

```sh
zfs create -o mountpoint=/jailer zroot/jailer
zfs create zroot/jailer/base_jail
```

### Base Jail

```sh
export URL_BASE=https://download.freebsd.org/ftp/releases/amd64/
cd /tmp

wget $URL_BASE/12.0-RELEASE/base.txz --no-check-certificate
wget $URL_BASE/12.0-RELEASE/lib32.txz --no-check-certificate

tar -zxvf /tmp/base.txz -C /jailer/base_jail
tar -zxvf /tmp/lib32.txz -C /jailer/base_jail

freebsd-update -b /jailer/base_jail fetch install
freebsd-update -b /jailer/base_jail IDS

zfs snapshot zroot/jailer/base_jail@12.1-RELEASE
```

### Create Jail

```sh
export JAILNAME=test_jail
zfs send -R zroot/jailer/base_jail@12.1-RELEASE | zfs receive zroot/jailer/$JAILNAME
```

Edit `/etc/jail.conf`:

```conf
# VNET disabled for now

test_jail {
    host.hostname = "test_jail.tobante.local";
    path = "/jailer/test_jail";

    exec.clean;
    exec.system_user = "root";
    exec.jail_user = "root";

    #vnet;
    #vnet.interface = "ng0_test_jail";               # vnet interface(s)

    #exec.prestart += "jng bridge test_jail igb1";   # bridge interface(s)
    #exec.poststop += "jng shutdown test_jail";      # destroy interface(s)

    # Standard stuff
    exec.start += "/bin/sh /etc/rc";
    exec.stop = "/bin/sh /etc/rc.shutdown";
    exec.consolelog = "/var/log/jail_test_jail_console.log";

    mount.devfs;          #mount devfs
    allow.raw_sockets;    #allow ping-pong
    devfs_ruleset="5";    #devfs ruleset for this jail
}
```

